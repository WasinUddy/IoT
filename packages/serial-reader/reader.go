package serial_reader

import (
	"bufio"
	"fmt"
	"time"

	"github.com/tarm/serial"
)

type Reader struct {
	port    string
	config  *serial.Config
	s       *serial.Port
	scanner *bufio.Scanner
}

func NewReader(port string, baud int) (*Reader, error) {
	cfg := &serial.Config{Name: port, Baud: baud, ReadTimeout: time.Second * 1}
	s, err := serial.OpenPort(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to open serial port %s: %w", port, err)
	}

	r := &Reader{
		port:    port,
		config:  cfg,
		s:       s,
		scanner: bufio.NewScanner(s),
	}
	return r, nil
}

func (r *Reader) ReadLine() (string, error) {
	if r.scanner.Scan() {
		return r.scanner.Text(), nil
	}
	if err := r.scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no data available")
}

func (r *Reader) Close() error {
	return r.s.Close()
}
