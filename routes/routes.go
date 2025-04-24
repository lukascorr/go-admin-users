package routes

import (
	"api/config"
	"api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(config.Authorize)

	userGroup := r.Group("/users")
	{
		userGroup.GET("", services.IndexUsers)
		userGroup.GET("/:userId", services.ShowUser)
		userGroup.POST("", services.StoreUser)
		userGroup.PATCH("/:userId", services.UpdateUser)
		userGroup.DELETE("/:userId", services.DeleteUser)
	}
}
