package main

import (
	"log"

	"github.com/streadway/amqp"
)

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

	body := "{ \"body\":\"Hello World!\", \"pattern\":\"test\"}"
	//y := x{data: "AAAAAAA"}
	//body, _ := json.Marshal(y)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
