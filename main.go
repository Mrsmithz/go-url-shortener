package main

import (
	"net/http"
	"server/src/repository"
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func main() {

	repository.Init()

	shortenerService := service.New()

	r := gin.Default()
	r.POST("/shorten", shortenerService.ShortenUrl)
	r.GET("/:url", shortenerService.GetOriginalUrl)
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	s.ListenAndServe()
}
