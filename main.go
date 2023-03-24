package main

import (
	"context"
	"github.com/ivandi1980/my-goservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Call the handler directly
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Create the handlers
	phandler := handlers.NewProducts(logger)

	// Register the handler with the server
	sm := http.NewServeMux()
	sm.Handle("/", phandler)

	// Create a new server
	server := &http.Server{
		Addr:         ":8888",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start the server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
