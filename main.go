package main

import (
	"github.com/gin-gonic/gin"
)

func homepage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
func main() {
	r := gin.Default()
	r.GET("/", homepage)
	r.Run()
}
