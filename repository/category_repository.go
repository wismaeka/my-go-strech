package repository

import "my-go-strech/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
