package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var PORT = "0.0.0.0:6379"

func main() {
	// create a TCP connection
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Failed to bind to port: ", PORT)
        os.Exit(1)
	}

	// close the listener once exit
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to bind to port: ", PORT)
        os.Exit(1)
	}
	
	defer conn.Close()

	// Read data
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading buffer:", err)
		os.Exit(1)
	}

	log.Printf("Received %d bytes", n)
	log.Println("Received data: ", buffer[:n])
	log.Println("Data: ", string(buffer[:n]))

	// Write data
	msg := []byte("+PONG\r\n")
	conn.Write(msg)
}