package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	DB *sql.DB
	ID string
	Name string
	Description string
}

func NewCategory(db *sql.DB) *Category{
	return &Category{
		DB: db,
	}
}

func(c *Category) Create(name,description string) (Category,error) {
	id := uuid.New().String()

	_,err := c.DB.Exec("insert into categories (id,name,description) values ($1,$2,$3)", id,name,description)
	if err != nil {
		return Category{},err
	}
	return Category{ID: id,Name: name,Description: description},nil
}

func(c *Category) FindAll() ([]Category,error) {
	rows,err := c.DB.Query("select id,name,description from categories")
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id,name,description string
		if err := rows.Scan(&id,&name,&description); err != nil {
			return nil,err
		}
		categories = append(categories, Category{ID: id,Name: name,Description: description})
	}
	return categories,nil
}

func(c *Category) FindByCourseID(courseID string) (Category,error) {
	var id,name,description string
	query := "select c.id,c.name,c.description from categories c join courses co on c.id=co.category_id where co.id=$1"
	err := c.DB.QueryRow(query,courseID).Scan(&id,&name,&description)
	if err != nil {
		return Category{},err
	}
	
	return Category{
		ID: id,
		Name: name,
		Description: description,
	},nil
}