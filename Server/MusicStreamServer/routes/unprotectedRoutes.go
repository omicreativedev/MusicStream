package routes

import (
	controller "github.com/omicreativedev/MusicStream/Server/MusicStreamServer/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnProtectedRoutes(router *gin.Engine, client *mongo.Client) {

	router.GET("/musics", controller.GetMusics(client))
	router.POST("/register", controller.RegisterUser(client))
	router.POST("/login", controller.LoginUser(client))
	router.POST("/logout", controller.LogoutHandler(client))
	router.GET("/genres", controller.GetGenres(client))
	router.POST("/refresh", controller.RefreshTokenHandler(client))
}
