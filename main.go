package main

import (
	"ginadmin/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.LoadHTMLGlob("views/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}