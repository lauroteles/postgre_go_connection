package router

import (
	"postgre_go/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.Router()

	router.HandleFunc("/api/stocks{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/API/stocks", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

}
