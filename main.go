package main

import (
	"hngcv/name"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Data struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}

func main() {
	
	r := mux.NewRouter()

	name.PrintName()

	r.Handle("/", http.FileServer(http.Dir("./client")))
	r.HandleFunc("/form", HandleForm).Methods("POST")

	fmt.Println("listening...")

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
   err := http.ListenAndServe(":"+port, r)
   if err != nil{
	   log.Fatal("server error: ", err)
   }
}

func HandleForm( w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/json")
	var data Data 
	_ = json.NewDecoder(r.Body).Decode(&data) 

	go WriteToJSON(data)
	json.NewEncoder(w).Encode(data) 
}

func  WriteToJSON (data Data) {
	file, err := json.MarshalIndent(data, "", " ") 
	if err != nil {
		log.Fatal("ERROR while writing to file", err)
	}
	_ = ioutil.WriteFile("formresponse.json", file, 0644)
}