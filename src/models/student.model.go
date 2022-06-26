package models

import (
	"github.com/xvbnm48/learnGorm/src/config"
	entity "github.com/xvbnm48/learnGorm/src/entities"
)

type StudentModel struct{}

func (student StudentModel) FindALl() ([]entity.Student, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var student []entity.Student
		db.Preload("Faculty").Find(&student)
		return student, nil
	}
}
