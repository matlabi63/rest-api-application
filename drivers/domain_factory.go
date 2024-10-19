package drivers

import (
	contentDomain "Echo/businesses/contents"
	contentDB "Echo/drivers/mysql/contents"

	categoryDomain "Echo/businesses/categories"
	categoryDB "Echo/drivers/mysql/categories"

	userDomain "Echo/businesses/users"
	userDB "Echo/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewContentRepository(conn *gorm.DB) contentDomain.Repository {
	return contentDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}