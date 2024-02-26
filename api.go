package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)

	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccout))
	router.HandleFunc("/account/", makeHTTPHandleFunc(s.handleAccout))
	log.Println("Api running on ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)

}

func (s *APIServer) handleAccout(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccout(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccout(w, r)
	}
	if r.Method == "PUT" {
		return s.handleEditAccout(w, r)
	}

	return nil
}
func (s *APIServer) handleGetAccout(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	log.Println(vars["id"])
	account := NewAccount("Victor", "Rezende")
	return writeJson(w, 200, account)
}

func (s *APIServer) handleCreateAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleEditAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
