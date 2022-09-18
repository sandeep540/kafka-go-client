package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Brokers:  []string{"localhost:49838"},
		GroupID:  "consumer-group-id",
		Topic:    "test",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	fmt.Printf("Starting Consumer... \n")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error inside loop %s", err)
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		fmt.Printf("Error outside loop %s", err)
		log.Fatal("failed to close reader:", err)
	}
}
