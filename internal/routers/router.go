package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wangys891019/blog-service/docs"
	"github.com/wangys891019/blog-service/internal/middleware"
	v1 "github.com/wangys891019/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()

	v1 := r.Group("/api/v1")
	{
		// 文章标签
		v1.POST("/tags", tag.Create)
		v1.DELETE("/tags/:id", tag.Delete)
		v1.PUT("/tags/:id", tag.Update)
		v1.PATCH("/tags/:id/state", tag.Update)
		v1.GET("/tags", tag.List)

		// 文章
		v1.POST("/articles", article.Create)
		v1.DELETE("/articles/:id", article.Delete)
		v1.PUT("/articles/:id", article.Update)
		v1.PATCH("/articles/:id/state", article.Update)
		v1.GET("/articles", article.List)
	}
	return r
}
