package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ApiServer struct {
	ListenAddr string
	store      Storage
}

//	func printX() any {
//		os.Getenv("PORT")
//		return fmt.Errorf("method not allowed %s", r.Method)
//	}
func NewApiServer(listenAddr string, store Storage) *ApiServer {
	return &ApiServer{
		ListenAddr: listenAddr,
		//store:      store,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	log.Println("JSON API server running on post: ", s.ListenAddr)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	//account := main.NewAccount("Eman", "xx")
	return writeJSON(w, http.StatusOK, "")

}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error
type ApiError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}

	}
}
