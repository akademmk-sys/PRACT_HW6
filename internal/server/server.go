package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Server *http.Server
	logger *log.Logger
}

func NewServer(logger *log.Logger) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HTMLGet)
	mux.HandleFunc("/upload", handlers.HTMLParser)

	mySrv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}
	return &Server{
		Server: mySrv,
		logger: logger,
	}
}
