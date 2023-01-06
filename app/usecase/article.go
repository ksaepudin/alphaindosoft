package usecase

import (
	"github.com/ksaepudin/alphaindosoft/app/entity"
)

type ArticlesUsecase interface {
	AddAtricles(data *entity.ArticlesListRequest) (interface{}, error)
}

type articles struct {
}

func NewArticleUsecase() ArticlesUsecase {
	return &articles{}
}

func (m *articles) AddAtricles(data *entity.ArticlesListRequest) (interface{}, error) {
	return data, nil
}
