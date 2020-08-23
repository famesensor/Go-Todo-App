package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBconfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBconfig {
	dbConfig := DBconfig{
		Host:     "0.0.0.0",
		Port:     3306,
		DBName:   "todos",
		Password: "root",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBconfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charst=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
}
