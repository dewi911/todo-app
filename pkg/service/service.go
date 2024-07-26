package service

import (
	TODO_app "TODO-app"
	repository "TODO-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user TODO_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list TODO_app.TodoList) (int, error)
	GetAll(userId int) ([]TODO_app.TodoList, error)
	GetById(userId, listId int) (TODO_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TODO_app.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item TODO_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TODO_app.TodoItem, error)
	GetById(userId, itemId int) (TODO_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input TODO_app.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
