package routes

import (
	"github.com/citiaps/GOLANG-Malla-backend/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1/plan")
	{
		v1.GET("/", controller.GetAllPlanes)
	}
}
