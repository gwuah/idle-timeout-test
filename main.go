package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("Hello World!"))
	})
	r.GET("/longconnection", func(c *gin.Context) {
		time.Sleep(11 * time.Minute)
		c.Writer.Write([]byte("longconnection"))
	})
	http.ListenAndServe("0.0.0.0:8192", r)
}
