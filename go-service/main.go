package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type MessageQueue struct {
	Body    string `json:"body"`
	Pattern string `json:"pattern"`
	Age     string `json:"age"`
	Data    string `json:"data"`
}

func NewMessageQueue(body, pattern, age string, data string) *MessageQueue {
	return &MessageQueue{
		body, pattern, age, data,
	}
}

func (m *MessageQueue) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(m)

	if err != nil {
		return nil, err
	}
	return bytes, err
}

func main() {
	conn, amqError := amqp.Dial("amqp://localhost:5672/")
	if amqError != nil {
		panic(amqError)
	}

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare(
		"default", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	message := NewMessageQueue("Hello...", "test", "20", "data...")
	body, _ := message.Marshal()

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
