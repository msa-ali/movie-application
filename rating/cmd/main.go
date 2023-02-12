package main

import (
	rating "github.com/Altamashattari/movieapplication/rating/internal/controller"
	httpHandler "github.com/Altamashattari/movieapplication/rating/internal/handler/http"
	"github.com/Altamashattari/movieapplication/rating/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
