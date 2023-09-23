package main

import (
	"net/http"
)

type goodResponse struct{
	Name string `json:"message"`
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, goodResponse{
		Name: "Hello Server",
	})
}
