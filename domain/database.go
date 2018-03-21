package domain

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	var err error
	err = godotenv.Load("mapic.env")
	if err != nil {
		panic(err)
	}
	// db, err = gorm.Open("mysql", "root@/mapic_api_development?charset=utf8&parseTime=True")
	db, err = gorm.Open("mysql", os.Getenv("CLOUD_SQL_PATH"))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Location{}, &User{})
}
