package main

import (
	"fmt"
	"log"
	"flag"
	"io"
	"log"
	"os"
	"github.com/tarm/serial"
)

var (
	targetComPort string
	baudRate int
	logfile string
)

// init flags
func init() {
	flag.StringVar(&targetComPort, "target", "/dev/ttyAMA0", "com port to sniff")
	flag.StringVar(&targetComPort, "logfile", "./comd.log", "logfile path")
	flag.IntVar(&baudRate, "baud", 9600, "baud rate")
}

func main() {
	// Setup serial config
	config := &serial.Config{
		Name: targetComPort,
		Baud: baudRate,
		ReadTimeout: 1,
		Size: 8
	}

	// Setup logfile
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening logfile: %v", err)
	}
	defer f.Close()

	// Create multiwriter, so we can log to both stdout AND a logfile
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)
	
	// Open serial port connection
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
		log.Println(fmt.Sprintf("%s", s))
	}
}
