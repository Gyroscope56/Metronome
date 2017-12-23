package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Metronome Server")
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))

}
