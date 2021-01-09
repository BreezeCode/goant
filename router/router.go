package router

import (
	"ginadmin/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	GroupV1 := r.Group("/admin")
	{
		GroupV1.GET("/index", (&controller.AdminController{}).Index)
	}
	r.GET("/index", (&controller.AdminController{}).Json)
}
