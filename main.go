package main

import (
	"log"

	"github.com/tzey/twittorcgs/bd"
	"github.com/tzey/twittorcgs/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
	}
	handlers.Manejadores()
}
