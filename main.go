package main

import (
	"bufio"
	"log"
	"os"

	"github.com/ydkulks/asm-lsp/rpc"
)

func main() {
	logger := getLogger("/home/yd/Projects/asm-lsp/log.txt")
	logger.Println("asm-lsp started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Unhandled file type")
	}

	return log.New(logfile, "[asm-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
