package catalog

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
	logger *log.Logger
}

type CatalogItem struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}

var catalog []CatalogItem

func (h *Handlers) Catalog(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request processed")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("catalog"))
}
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

// Display all from the catalog var
func GetCatalog(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(catalog)
}

// Display a single data
func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range catalog {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&CatalogItem{})
}

// create a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var catalogItem CatalogItem
	_ = json.NewDecoder(r.Body).Decode(&catalogItem)
	catalogItem.ID = params["id"]
	catalog = append(catalog, catalogItem)
	json.NewEncoder(w).Encode(catalog)
}

// Delete an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range catalog {
		if item.ID == params["id"] {
			catalog = append(catalog[:index], catalog[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(catalog)
	}
}

func (h *Handlers) SetupRoutes(router *mux.Router) {
	catalog = append(catalog, CatalogItem{ID: "1", Title: "Arifureta"})
	catalog = append(catalog, CatalogItem{ID: "2", Title: "Zectas"})
	router.HandleFunc("/catalog", GetCatalog).Methods("GET")
	router.HandleFunc("/catalog/{id}", GetItem).Methods("GET")
	router.HandleFunc("/catalog/{id}", CreateItem).Methods("POST")
	router.HandleFunc("/catalog/{id}", DeleteItem).Methods("DELETE")
}
