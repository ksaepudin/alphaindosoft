package service

import (
	"net/http"

	"github.com/ksaepudin/alphaindosoft/app/entity"
	"github.com/ksaepudin/alphaindosoft/app/helper"
	"github.com/ksaepudin/alphaindosoft/app/usecase"
	"github.com/labstack/echo/v4"
)

type ArticlesService interface {
	AddAtricles(c echo.Context) error
}

type articles struct {
	uc usecase.ArticlesUsecase
}

func NewArticle(uc usecase.ArticlesUsecase) ArticlesService {
	return &articles{uc: uc}
}

func (m *articles) AddAtricles(c echo.Context) error {
	req := new(entity.Articles)
	if err := c.Bind(&req); err != nil {
		return helper.Response(c, http.StatusBadRequest, err.Error())
	}
	input := &entity.ArticlesListRequest{
		Author: req.Author,
		Title:  req.Title,
		Body:   req.Body,
	}
	res, errRes := m.uc.AddAtricles(input)
	if errRes != nil {
		return helper.Response(c, http.StatusFailedDependency, "Usecase Error", errRes)
	}
	return helper.Response(c, http.StatusOK, "Yey Berhasil", res)
}
