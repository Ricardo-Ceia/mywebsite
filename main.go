package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/",fileServer)

	log.Println("port running on http://localhost:8000")
	if err:=http.ListenAndServe(":8000",nil);err!=nil{
		log.Fatal(err)
	}
	
}

