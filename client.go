package goworker

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	"gopkg.in/redis.v5"
)

type Runnable interface {
	Run(args ...interface{}) error
}

type item struct {
	job Runnable
	at  time.Time
}

type client struct {
	conn  *redisConn
	items []item
}

func (c *client) push(it *item) error {
	jobDump, err := json.Marshal(it.job)

	if err != nil {
		logrus.Error(err)
		return err
	}

	z := redis.Z{
		Score:  float64(it.at.Unix()),
		Member: string(jobDump),
	}

	c.conn.Client.ZAdd("schedule", z)
	return nil
}
