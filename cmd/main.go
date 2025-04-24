package main

import (
	"api/config"
	"api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDatabase()
}

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8000")
}
