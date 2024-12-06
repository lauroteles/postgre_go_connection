package main

import (
	"fmt"
	"log"
	"net/http"
	"postgre_go/router"
)

func main() {

	r := router.Router()
	fmt.Println("Starting server on the port 8008...")

	log.Fatal(http.ListenAndServe(":8080", r))

}
