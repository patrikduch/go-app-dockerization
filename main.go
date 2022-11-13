package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default() //new gin router initialization
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World from Patrik Duch :)"})
	}) // first endpoint returns Hello World
	router.Run(":8000") //running application, Default port is 8080
}
