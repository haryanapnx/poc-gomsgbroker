package main

import (
	"Golang/src/test-is-uji/broker"
	"fmt"
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

	queue, err := ch.QueueDeclare(
		"test-is-uji", // name
		false,         // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // args
	)
	if err != nil {
		panic(err)
	}

	// declaring consumer with its properties over channel opened
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan interface{})
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
