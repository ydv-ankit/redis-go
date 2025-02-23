package main

import (
	"net"
	"os"

	"github.com/ydv-ankit/redis-go/utils"
)

func main() {
	host := "0.0.0.0"
	port := "6379"

	listener, err := net.Listen(("tcp"), host+":"+port)

	if err != nil {
		utils.Logger.ErrorLogger.Println("unable to start server on port: 6379")
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.Logger.ErrorLogger.Println("unable to accept connection")
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}
