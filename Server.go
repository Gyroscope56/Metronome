package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Metronome Server")
	panic(http.ListenAndServe(":80", http.FileServer(http.Dir("."))))

}
