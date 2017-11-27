package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	port := flag.Int("port", 6100, "The port to listen on.")
	path := flag.String("path", ".", "The path to serve.")
	flag.Parse()

	http.Handle("/", handlers.LoggingHandler(
		os.Stdout,
		http.FileServer(http.Dir(*path))),
	)

	log.Printf("Starting on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
