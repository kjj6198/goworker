package goworker

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/redis.v5"
)

type redisConn struct {
	Client    *redis.Client
	NameSpace string
}

func NewClient() *redis.Client {
	logrus.Info("starting redis server")
	client := redis.NewClient(&redis.Options{
		// TODO: replace with constant
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func (client *redisConn) Close() {
	if client.Client != nil {
		client.Client.Close()
	}
}
