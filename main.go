package main

import (
	"log"
	"logger"

	_ "github.com/ClorisMoQi/logger/logger"
)

func main() {
	log.Fatal("1111111")
    logger := logger.NewLogger("debug")
    logger.DEBUG("identifier", "hahaha!!!")
}