package main

import (
	"fmt"
)

func main() {
	//store, err := NewPostgressStore()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err := store.Init; err != nil {
	//	log.Fatal(err)
	//}

	fmt.Printf("hello")

	//apiService := NewApiServer("8070")
	//router := mux.NewRouter()
	//router.HandleFunc("/account", makeHTTPHandlerFunc(apiService.handleAccount))
	//log.Println("JSON API server running on post: ", apiService.ListenAddr)
	server := NewApiServer(":8070", nil)
	server.Run()

}
