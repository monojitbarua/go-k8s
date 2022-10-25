package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/monojitbarua/go-util-lib/zlogger"
)

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", Handler)
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/readiness", ReadinessHandler)
	r.HandleFunc("/customers", GetCustomersHandler)

	server := &http.Server{
		Handler:      r,
		Addr:         ":4040",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		zlogger.Info("Starting Server ...")
		if err := server.ListenAndServe(); err != nil {
			zlogger.Error(err.Error())
		}
	}()

	// Graceful Shutdown
	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	zlogger.Info("Shutting down")
	os.Exit(0)
}
