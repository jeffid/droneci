package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
)

func TestEcho(t *testing.T) {
	Echo()
	t.Log("test OK")
}

func TestRedis(t *testing.T) {
	ctx:=context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res:=rdb.Incr(ctx,"after")
	fmt.Println("redis res: ",*res)
}