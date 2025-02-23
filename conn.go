package main

import (
	"net"
	"strings"

	"github.com/ydv-ankit/redis-go/utils"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		utils.Logger.ErrorLogger.Println("unable to read from connection")
		return
	}
	if strings.TrimSpace(string(buffer[:n])) == "PING" {

		data := []byte("+PONG\r\n")
		_, err = conn.Write(data)
	}
	if err != nil {
		utils.Logger.ErrorLogger.Println("unable to send response")
	}
}
