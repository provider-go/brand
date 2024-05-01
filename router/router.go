package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/brand/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// brand 表操作
		Router.POST("add", api.CreateBrand)
		Router.POST("delete", api.DeleteBrand)
		Router.POST("list", api.ListBrand)
		Router.POST("view", api.ViewBrand)
	}
}
