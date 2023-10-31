package service

import (
	"errors"
	"jobportal-graphql/graph/model"
	"jobportal-graphql/repo"
)

type UserService interface {
	UserSignup(model.NewUser) (*model.User, error)
	CreateCompany(model.NewCompany) (*model.Company, error)
	ViewAllCompanies() ([]*model.Company, error)
}

type Service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		userRepo: userRepo,
	}, nil
}
