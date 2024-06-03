// used https://medium.com/@pradityadhitama/simple-logger-in-golang-f72dadf2c8c6
//this file appends log entries to logfile.log (which will remain empty until the script is ran for the first time)

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//open file and logger
	file, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(file, "Logger:", log.Ltime)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	 }
	//add the current time every second as a log entry
	for {
		logger.Println("Current log entry")
		time.Sleep(time.Second)
	}
}
