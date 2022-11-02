package main

import (
	"Golang/src/test-is-uji/broker"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
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

	// declaring queue with its properties over the the channel opened
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

	// publishing a message
	err = ch.Publish(
		"",         // exchange
		queue.Name, // key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(os.Args[1]),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)

	fmt.Println("Successfully published message", os.Args[1])
}
