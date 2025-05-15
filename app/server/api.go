package server

import "github.com/gin-gonic/gin"

// 注册所有API路由
func SetupRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.GET("/ping", ping)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
