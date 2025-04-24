package repositories

import (
	"api/config"
	"api/models"
)

func GetUsersByFilter() []models.User {
	var users []models.User
	config.DB.Find(&users)

	return users
}

func GetUserById(userId string) models.User {
	var user models.User
	config.DB.Where("id = ?", userId).First(&user)

	return user
}

func CreateUser(user models.User) models.User {
	config.DB.Create(&user)

	return user
}

func UpdateUser(user models.User) models.User {
	config.DB.Save(&user)

	return user
}

func DeleteUser(userId int) {
	var user models.User
	config.DB.Where("id = ?", userId).Delete(user)
}
