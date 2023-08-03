package main

import (
	"github.com/damyan.dimitrov/gogardenofapples/initializers"
	"github.com/damyan.dimitrov/gogardenofapples/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
