package service

import (
	"reddit123/models"
	"reddit123/package/repository"
)

type Post interface {
	GetById(id string) (*models.Post, error)
	GetList(page int, limit int) (*models.OutputPostList, error)
	Create(post *models.InputPost) (*models.OutputPost, error)
	Update(post *models.InputUpdatePost) error
	Delete(id string) error
}

type Service struct {
	Post Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post: NewPostService(repos.Post),
	}
}
