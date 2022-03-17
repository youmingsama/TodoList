package main

import (
	"ToDoList/dao"
	models "ToDoList/models"
	"ToDoList/routers"
	"ToDoList/setting"
	"fmt"
)

func main() {
	// 加载配置文件
	if err := setting.Init("conf/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	_, err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// 模型绑定
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}