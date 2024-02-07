package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kalyanKumarPokkula/Go-jwt/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.Connect()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.GET("/ping" , func(ctx *gin.Context) {
		ctx.JSON(200 , gin.H{
			"message" : "pong",
		})
	})

	r.Run()
}

