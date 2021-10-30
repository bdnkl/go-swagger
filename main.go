package main

import (
	"github.com/bdnkl/go-swagger/api"
	_ "github.com/bdnkl/go-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumByID)
	router.POST("/albums", api.PostAlbums)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
