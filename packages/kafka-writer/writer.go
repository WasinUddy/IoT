package kafka_writer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Writer struct {
	writer *kafka.Writer
}

func NewWriter(brokers []string, topic string) *Writer {
	return &Writer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll,
		},
	}
}

func (w *Writer) Publish(ctx context.Context, key, value []byte) error {
	return w.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	})
}

func (w *Writer) Close() error {
	return w.writer.Close()
}
