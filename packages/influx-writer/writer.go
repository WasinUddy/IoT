package influx_writer

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Writer struct {
	client influxdb2.Client
	org    string
	bucket string
}

func NewWriter(url, token, org, bucket string) *Writer {
	client := influxdb2.NewClient(url, token)
	return &Writer{
		client: client,
		org:    org,
		bucket: bucket,
	}
}

func (w *Writer) Write(ctx context.Context, measurement string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) error {
	writeAPI := w.client.WriteAPIBlocking(w.org, w.bucket)

	point := influxdb2.NewPoint(
		measurement,
		tags,
		fields,
		timestamp,
	)

	return writeAPI.WritePoint(ctx, point)
}

func (w *Writer) Close() {
	w.client.Close()
}
