package services

import (
	"api/models"
	"api/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexUsers(c *gin.Context) {
	users := repositories.GetUsersByFilter()
	c.JSON(http.StatusOK, users)
}

func ShowUser(c *gin.Context) {
	userId := c.Param("userId")
	user := repositories.GetUserById(userId)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func StoreUser(c *gin.Context) {
	var userRequest models.User

	if err := c.BindJSON(&userRequest); err != nil {
		return
	}
	user := repositories.CreateUser(userRequest)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	user := repositories.GetUserById(userId)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}

	var userRequest models.User
	if err := c.BindJSON(&userRequest); err != nil {
		return
	}

	userRequest.Id = user.Id
	res := repositories.UpdateUser(userRequest)
	c.JSON(http.StatusCreated, res)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	user := repositories.GetUserById(userId)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	} else {
		repositories.DeleteUser(user.Id)
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted",
		})
	}
}
