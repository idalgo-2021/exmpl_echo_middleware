package main

import (
	"log"

	"echo_middleware/internal/pkg/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
