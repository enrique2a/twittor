package main

import (
	"log"

	"github.com/enrique2a/twittor/bd"
	"github.com/enrique2a/twittor/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Error to connect to the datbase")
		return
	}
	handlers.Manejadores()
}
