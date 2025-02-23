package utils

import (
	"log"
	"os"
)

// Declare a global logger
var Logger = &LoggerStruct{
	InfoLogger:  log.New(os.Stdout, "INFO: ", log.Lshortfile),
	ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Lshortfile),
}

type LoggerStruct struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}
