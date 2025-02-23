package main

import (
	"log"
	"net"
	"os"

	"github.com/ydv-ankit/redis-go/app/utils"
)

func main() {
	listener, err := net.Listen(("tcp"), "0.0.0.0:6379")

	infoLogger := log.New(os.Stdout, "INFO: ", log.Lshortfile)
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Lshortfile)

	logger := utils.Logger{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}

	if err != nil {
		logger.ErrorLogger.Println("unable to create tcp server")
	}

}
