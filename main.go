package main

import (
	"fmt"
	"log"
	"flag"
	"github.com/tarm/serial"
)

var (
	targetComPort string
	baudRate int
)

// init flags
func init() {
	flag.StringVar(&targetComPort, "target", "/dev/ttyAMA0", "specify com port to sniff")
	flag.IntVar(&baudRate, "baud", 9600, "specify baud rate")
}

func main() {
	config := &serial.Config{
	Name: targetComPort,
	Baud: baudRate,
	ReadTimeout: 1,
	Size: 8
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)

	for {
		n, err := stream.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[:n])
		fmt.Println(s) // Replace with multi logger (stdout & text file)
	}
}
