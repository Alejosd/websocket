package main

import (
	"log"
	"net/http"
	"./chat"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port=":8080"
	}

	log.SetFlags(log.Lshortfile)

	// websocket server
	server := chat.NewServer("/entry")
	go server.Listen()
	// static files
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Fatal(http.ListenAndServe(port, nil))




}
