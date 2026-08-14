package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)

	s := server.NewServer(logger)
	if err := s.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
