package entity

import "fmt"

type Faculty struct {
	Id      int `gorm:"primary_key, AUTO_INCREMENT"`
	Name    string
	Student []Student `gorm:"foreignkey:FacultyID"`
}

func (faculty *Faculty) TableName() string {
	return "faculty"
}

func (faculty *Faculty) ToString() string {
	return fmt.Sprintf("id: %d \n name %s", faculty.Id, faculty.Name)
}
