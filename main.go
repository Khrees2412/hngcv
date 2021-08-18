package main

import (
	"fmt"
	"hngcv/name"
	"log"
	"os"

	"net/http"
)


func main() {
	
	name.PrintName()

	http.Handle("/", http.FileServer(http.Dir("./client")))
	fmt.Println("listening...")
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
   err := http.ListenAndServe(":"+port, nil)
   if err != nil{
	   log.Fatal("server error: ", err)
   }
}