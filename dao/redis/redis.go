package redis

import (
	"context"
	"fmt"
	"github.com/learnviper/setting"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	//Nil    = redis.Nil
)
var ctx = context.Background()

//type SliceCmd = redis.SliceCmd
//type StringStringMapCmd = redis.StringStringMapCmd

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	} else {
		fmt.Println("Redis init success!")
	}
	return err
}

func Close() {
	_ = client.Close()
}
