package entity

import "fmt"

type Student struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Name      string
	Address   string
	FacultyID int `gorm:"collumn:faculty_id"`
	Faculty   Faculty
}

func (student *Student) TableName() string {
	return "student"
}

func (student *Student) ToString() string {
	return fmt.Sprintf("id: %d \n name %s \n address %s \n faculty %s", student.Id, student.Name, student.Address, student.Faculty.ToString())
}
