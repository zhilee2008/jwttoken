package utils

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var Cache cache.Cache

func Init() {
	// collectionName := beego.AppConfig.String("redis::collectionName")
	conn := beego.AppConfig.String("redis::conn")
	dbNum := beego.AppConfig.String("redis::dbNum")
	//password := beego.AppConfig.String("cache.password")
	// 设置配置参数
	config := orm.Params{
		"key":      "testa",
		"conn":     conn,
		"dbNum":    dbNum,
		"password": password,
	}
	configStr, err := json.Marshal(config)
	logs.Debug(string(configStr))
	if err != nil {
		logs.Error("redis配置模型转换失败")
		return
	}
	Cache, err = cache.NewCache("redis", string(configStr))
	if err != nil {
		logs.Error("redis初始化失败")
		return
	}
	Cache.Put("aaa", "aa", 1000)
	logs.Info("******************************************************************************")
	logs.Info("********************************redis连接成功**********************************")
	logs.Info("******************************************************************************")

	// c, err := redis.Dial("tcp", "127.0.0.1:6379")
	// 	if err != nil {
	// 		fmt.Println("Connect to redis error", err)
	// 		return
	// 	}
	// 	defer c.Close()

	// 	_, err = c.Do("SET", "testtoken", token)
	// 	if err != nil {
	// 		fmt.Println("redis set failed:", err)
	// 	}
}
