package main


import "net/http"
import "log"





func main(){

	router := NewRouter()
	server := http.ListenAndServe(":8080",router)
	log.Fatal(server)
}

