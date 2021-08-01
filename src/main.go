package main

import (
	"github.com/gin-gonic/gin"
	"src/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", routes.Home)

	router.Run(":8080")
}