package main

import (
	"Golang/src/test-is-uji/broker"
	"bytes"
	"log"
	"time"

	"github.com/pkg/errors"
)

func main() {
	conn, ch, err := broker.RabbitMQ()
	if err != nil {
		panic(err)
	}

	defer func() {
		ch.Close()
		conn.Close()
	}()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to declare queue"))
	}

	// berfungsi untuk mengatur berapa banyak jumlah tugas yang boleh dikerjakan
	// oleh worker tersebut dalam waktu yang bersamaan.
	// Angka 1 menunjukkan bahwa worker tersebut hanya boleh mengerjakan 1 tugas
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to set QoS"))
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to consume queue"))
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Println("Done!")
			// menandakan bahwa task sudah kelar
			d.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for messages. To Exit press CTRL+C\n")
	<-forever
}
