package main

import (
	"github.com/ivandi1980/my-goservice/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	// Call the handler directly
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Create the handlers
	helloHandler := handlers.NewHello(logger)
	goodbyeHandler := handlers.NewGoodbye(logger)

	// Register the handler with the server
	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	// Start the server
	http.ListenAndServe(":8888", sm)
}
