package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kalyanKumarPokkula/Go-jwt/controllers"
	"github.com/kalyanKumarPokkula/Go-jwt/initializers"
	"github.com/kalyanKumarPokkula/Go-jwt/middlewares"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.Connect()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.POST("/signup" , controllers.Signup)
	r.POST("/login" , controllers.Login)
	r.GET("/validate" ,middlewares.AuthenticateJwt , controllers.Vaildate)

	r.Run()
}

