package main

import (
	"github.com/KasimDilsiz/Golang-go-gin/controller"
	"github.com/KasimDilsiz/Golang-go-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService = service.New()

	videoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))

	})
	server.Run(":8080")

}
