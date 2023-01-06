package router

import (
	"github.com/ksaepudin/alphaindosoft/app/service"
	"github.com/ksaepudin/alphaindosoft/app/usecase"
	"github.com/labstack/echo/v4"
)

func Articles(e *echo.Echo) {
	articlesUsecase := usecase.NewArticleUsecase()
	articleService := service.NewArticle(articlesUsecase)
	grpArticle := e.Group("/article")
	grpArticle.POST("", articleService.AddAtricles)
}
