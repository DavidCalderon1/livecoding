package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	UrlApiRandom = "https://api.chucknorris.io/jokes/random"
)

type Categories struct {
	CreatedAt string `json:"created_at"`
	IconUrl   string `json:"icon_url"`
	Id        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
	Url       string `json:"url"`
	Value     string `json:"value"`
}

type CategoriesController struct {
}

func (c *CategoriesController) GetRandomCategory(w http.ResponseWriter, r *http.Request) {
	var categories Categories

	response, err := http.Get(UrlApiRandom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(categories)

	return
}

func NewCategoriesController() *CategoriesController {
	return &CategoriesController{}
}
