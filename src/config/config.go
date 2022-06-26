package config

import "github.com/jinzhu/gorm"

func GetDB() (*gorm.DB, error) {
	dbDriver := "mysql"
	dbName := "learnGorm"
	dbUser := "root"
	dbPassword := ""
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}

	return db, nil
}
