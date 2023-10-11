package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course{
	return &Course{DB: db}
}

func (c *Course) Create(name, description, categoryID string) (*Course,error)  {
	id := uuid.New().String()
	query := `insert into courses (id,name,description,category_id) values($1,$2,$3,$4)`
	_,err := c.DB.Exec(query,id,name,description,categoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID: id,
		Name: name,
		Description: description,
		CategoryID: categoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course,error) {
	rows,err := c.DB.Query("select id,name,description,category_id from courses")
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id,name,description,categoryID string
		if err := rows.Scan(&id,&name,&description,&categoryID); err != nil {
			return nil,err
		}
		courses = append(courses, Course{ID: id,Name: name,Description: description,CategoryID: categoryID})
	}
	return courses,nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course,error) {
	rows,err := c.DB.Query("select id,name,description,category_id from courses where category_id=$1",categoryID)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id,name,description,categoryID string
		if err := rows.Scan(&id,&name,&description,&categoryID); err != nil {
			return nil,err
		}
		courses = append(courses, Course{ID: id,Name: name,Description: description,CategoryID: categoryID})
	}
	return courses,nil
}