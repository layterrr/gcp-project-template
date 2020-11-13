package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blendle/zapdriver"
	server "github.com/frontedxyz/synthwave/services/{{.ServiceName}}/server"
	pb "github.com/frontedxyz/synthwave/services/{{.ServiceName}}/proto"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	var err error
	ctx := context.Background()

	// Get GCP project id
	projectID, ok := os.LookupEnv("PROJECT_ID")
	if !ok {
		log.Fatalf("PROJECT_ID env variable isn't set\n")
	}

	// Initialise logger
	var logger *zap.Logger
	if isLocal(projectID) {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zapdriver.NewProduction()
	}
	if err != nil {
		log.Fatalf("Unable to configure logger: %v\n", err)
	}
	zap.ReplaceGlobals(logger)

	// Load .env file in development
	if isLocal(projectID) {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Unable to load .env file: %v\n", err)
		}
	}

	// Create twirp server handler
	server := &server.Server{
		Logger: logger,
	}
	handler := pb.NewLoanServiceServer(server, nil)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// Run the service
	log.Printf("Running service on http://localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}

func isLocal(projectID string) bool {
	return projectID == "local"
}
