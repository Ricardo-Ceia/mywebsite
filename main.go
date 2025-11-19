package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      nil, // Uses DefaultServeMux
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("port running on http://localhost:8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
