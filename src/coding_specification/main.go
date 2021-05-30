package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayhello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message": "Hello golang!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/hello",sayhello)
	//原版风格
	//r.GET("/book",...)
	//r.GET("/create_book",...)
	//r.GET("update_ book",...)
	//r.GET("delete_book",...)

	//restful风格-->我们使用restful风格
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"method":"GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})
	r.Run(":9090")
	//可利用postman软件向服务器无线发包

}