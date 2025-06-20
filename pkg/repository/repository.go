package repository

import (
	"druna_server/pkg/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, passwordHash string) (model.User, error)
}

type User interface {
}

type Event interface {
}

type Friendship interface {
}

type Group interface {
}

type Repository struct {
	Authorization
	User
	Event
	Friendship
	Group
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
