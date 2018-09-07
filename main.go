package main

import (
	"log"
	"os"

	"github.com/tomhibbers/lnlamas-api-go/catalog"
	"github.com/tomhibbers/lnlamas-api-go/server"

	"github.com/gorilla/mux"
)

var CertFile = "certs/server.crt"
var KeyFile = "certs/server.key"
var ServiceAddr = ":8080"

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
