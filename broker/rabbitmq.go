package broker

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get channel")
	}
	return conn, ch, nil
}