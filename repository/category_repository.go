package repository

import "golang-mocking-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}