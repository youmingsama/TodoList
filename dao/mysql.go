package dao

import (
	"ToDoList/setting"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitMySQL(cfg *setting.MySQLConfig) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return DB.DB()
}
