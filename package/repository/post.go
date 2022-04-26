package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"reddit123/models"
	"time"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (s *PostPostgres) GetById(id string) (*models.Post, error) {
	var post models.Post

	if err := s.db.Get(&post, `SELECT * FROM "post" WHERE id=$1`, id); err != nil {
		return nil, err
	}

	return &post, nil
}
func (s *PostPostgres) GetList(page int, limit int) (*models.OutputPostList, error) {
	var output models.OutputPostList

	if err := s.db.Select(&output, `SELECT * FROM "post" WHERE deleted=FALSE ORDER BY create_date = DESC limit $1 OFFSET $2`,
		limit, (page-1)*limit); err != nil {
		return nil, err
	}

	if err := s.db.Get(&output.TotalCount, `SELECT count(*) FROM "post" WHERE deleted=FALSE`); err != nil {
		return nil, err
	}
	return &output, nil
}

func (s *PostPostgres) Create(post *models.InputPost) (*models.OutputPost, error) {
	id := uuid.New().String()
	if id == "" {
		return nil, errors.New("generate uuid invalid")
	}

	timeNow := time.Now()
	_, err := s.db.Query(`insert into "post"(id, author, caption, body, create_date, deleted)
   values($1, $2, $3, $4, $5, $6)`,
		id, post.Author, post.Caption, post.Body, timeNow, false)
	if err != nil {
		return nil, err
	}

	return &models.OutputPost{
		Id:         id,
		CreateDate: timeNow,
	}, nil
}
func (s *PostPostgres) Update(post *models.InputUpdatePost) error {

}
func (s *PostPostgres) Delete(id string) error {

}
