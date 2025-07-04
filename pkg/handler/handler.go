package handler

import (
	"druna_server/pkg/service"

	"github.com/gin-gonic/gin"

	_ "druna_server/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()


	ping := router.Group("/ping")
	{
		ping.GET("/", h.ping)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/renew-token", h.renewToken)
	}

	api := router.Group("/api", h.userIdentity)
	{
		friends := api.Group("/friends")
		{
			friends.GET("/list", h.getFriendList)
			friends.GET("/request-list", h.getFriendRequestList)
			friends.POST("/request", h.sendFriendRequest)
			friends.POST("/accept", h.acceptFriendRequest)
			friends.POST("/reject", h.rejectFriendRequest)
			friends.DELETE("/", h.deleteFriend)
		}

		events := api.Group("/events")
		{
			events.GET("/", h.getEventList)
			events.POST("/", h.addEvent)
			events.DELETE("/:id", h.deleteEvent)

			events.POST("/free-time", h.getFreeTime)
		}

		groups := api.Group("/groups")
		{
			groups.POST("/create", h.createGroup)
		}
	}
	return router
}
