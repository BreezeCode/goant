package controller

import "github.com/gin-gonic/gin"

type AdminController struct {
}

func (t *AdminController) Index(c *gin.Context) {
	c.Header("token", "GinServer")
	c.HTML(200, "index.html", gin.H{
		"title": "Gin 测试模板",
	})
}

func (t *AdminController) Json(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
