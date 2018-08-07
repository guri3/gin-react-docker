package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "poyo",
		})
	})

	r.Run(":3030")
}
