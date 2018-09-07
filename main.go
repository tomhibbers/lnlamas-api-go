package main

import (
	"lnlamas-api-go/catalog"
	"lnlamas-api-go/server"
	"log"
	"os"

	"github.com/gorilla/mux"
)

var CertFile = "certs/server.crt"
var KeyFile = "certs/server.key"
var ServiceAddr = ":8009"

func main() {

	logger := log.New(os.Stdout, "lnlamas", log.LstdFlags|log.Lshortfile)
	h := catalog.NewHandlers(logger)
	router := mux.NewRouter()
	h.SetupRoutes(router)

	logger.Println("server starting")
	srv := server.New(router, ServiceAddr)
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}

}
