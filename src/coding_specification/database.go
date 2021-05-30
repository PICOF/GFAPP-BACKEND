package main

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //导入gorm库
	_ "github.com/jinzhu/gorm/dialects/mysql" //导入本地MySQL数据库驱动
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}
//手动创建数据库：CREATE DATABASE db1;
func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3307)/monitor?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		fmt.Println("database connection failed")
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "cxk1", "male", "唱"}
	u2 := UserInfo{2, "cxk2", "female", "跳"}
	u3 := UserInfo{2, "cxk3", "male", "rap"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "rap")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "篮球")
	// 删除
	db.Delete(&u)
}
