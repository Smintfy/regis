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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to bind to port: ", PORT)
			os.Exit(1)
		}

		// concurrency for handling multiple connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// read data
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading buffer:", err)
			return
		}

		log.Printf("Received %d bytes", n)
		log.Println("Received data: ", buffer[:n])
		log.Println("Data: ", string(buffer[:n]))

		// write data
		msg := []byte("+PONG\r\n")
		conn.Write(msg)
	}
}