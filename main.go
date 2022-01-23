package main

import (
	"api"
	_ "docs"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumByID)
	router.POST("/albums", api.PostAlbums)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run(":" + getPort())
	if err != nil {
		return
	}
}
