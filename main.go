package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello magic stream movies!")

	})

	if err := router.Run(":3000"); err != nil {
		fmt.Println("Failed to start server ", err)
	}

}
