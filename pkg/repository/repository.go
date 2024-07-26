package repository

import (
	TODO_app "TODO-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TODO_app.User) (int, error)
	GetUser(username, password string) (TODO_app.User, error)
}

type TodoList interface {
	Create(userId int, list TODO_app.TodoList) (int, error)
	GetAll(userId int) ([]TODO_app.TodoList, error)
	GetById(userId, listId int) (TODO_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TODO_app.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item TODO_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TODO_app.TodoItem, error)
	GetById(userId, itemId int) (TODO_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input TODO_app.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
