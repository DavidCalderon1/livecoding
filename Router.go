package main

import (
	"github.com/gorilla/mux"
	"livecoding/controllers"
)

func categoriesControllerProvider() *controllers.CategoriesController {
	return controllers.NewCategoriesController()
}

func newRouter() *mux.Router {
	categoriesController := categoriesControllerProvider()

	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/random", categoriesController.GetRandomCategory).Methods("GET")

	return router
}
