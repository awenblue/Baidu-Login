package baidulogin

import (
	"database/sql"
        "fmt"
      _ "github.com/go-sql-driver/mysql"
        "time"
)

//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK = "tcp"
	SERVER = "mysql"
	PORT = 3306
	DATABASE = "db_aus"
)

//插入数据
func InsertData(cookie string) {
    conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
    DB, err := sql.Open("mysql", conn)
    if err != nil {
	fmt.Println("connection to mysql failed:", err)
	return
    }

    DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超时的连接就close
    DB.SetMaxOpenConns(100)                //设置最大连接数
    timeStr:=time.Now().Format("2006-01-02 15:04:05")
    result,err := DB.Exec("insert INTO tb_cookie(phone,cookie,create_date) values(?,?,?)","123", cookie)
    if err != nil{
        fmt.Printf("Insert data failed,err:%v", err)
        return
    }
    lastInsertID,err := result.LastInsertId()    //获取插入数据的自增ID
    if err != nil {
        fmt.Printf("Get insert id failed,err:%v", err)
        return
    }
    fmt.Println("Insert data id:", lastInsertID)

    rowsaffected,err := result.RowsAffected()  //通过RowsAffected获取受影响的行数
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("Affected rows:", rowsaffected)
}

