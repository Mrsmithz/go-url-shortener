package repository

import (
	"context"
	"fmt"
	"server/src/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type db struct {
	Redis redis.UniversalClient
}

var DB db

func Init() {
	redisConnection()
}

func redisConnection() {
	opt := redis.UniversalOptions{
		Addrs: []string{"redis:6379"},
	}
	DB.Redis = redis.NewUniversalClient(&opt)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	result, err := DB.Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("redis connected with result :%s \n", result)

}

func (d *db) SaveUrl(ctx context.Context, url model.Url) {
	err := d.Redis.Set(ctx, url.Shorten, url.Original, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (d *db) GetOriginalUrl(ctx context.Context, url string) (string, error) {
	value, err := d.Redis.Get(ctx, url).Result()

	return value, err
}
