package main

import (
	"example.com/mod/albums"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albums.GetAlbums)
	router.GET("/albums/:id", albums.GetAlbumByID)
	router.POST("/albums", albums.AddAlbums)

	router.Run("localhost:8080")
}
