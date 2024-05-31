// used https://medium.com/@pradityadhitama/simple-logger-in-golang-f72dadf2c8c6

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(file, "Logger:", log.Ltime)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	 }

	for {
		logger.Println("Current log entry")
		time.Sleep(time.Second)
	}
}
