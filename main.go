package main

import (
	"fmt"
	"github.com/learnviper/dao/mysql"
	"github.com/learnviper/dao/redis"
	"github.com/learnviper/setting"
)

func main() {
	// viper读取配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed ,err:%v\n", err)
		return
	}

	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("Init Mysql failed ,err :%v", err)
	}
	defer mysql.Close()

	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("Init Redis failed ,err :%v", err)
	}
	defer redis.Close()

}
