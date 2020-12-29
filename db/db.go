package db

import (
	"members/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:password@tcp(db:3306)/development?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                            // default size for string fields
		DisableDatetimePrecision:  true,                                                                           // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                           // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                           // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                          // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
	user := models.User{
		ID:    1,
		Name:  "aoki",
		Posts: []models.Post{{ID: 1, Content: "tweet1"}, {ID: 2, Content: "tweet2"}},
	}
	db.Create(&user)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	// if err := db.Close(); err != nil {
	// 	panic(err)
	// }
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
