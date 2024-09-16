package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.uber.org/zap"
)

const LOG_PATH = "./logs/go.log"

func main() {
	os.OpenFile(LOG_PATH, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", "./logs/go.log"}
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	i := 0
	for {
		i++
		time.Sleep(time.Second * 3)
		if rand.Intn(10) == 1 {
			l.Error("test error", zap.Error(fmt.Errorf("error because test: %d", i)))
		} else {
			l.Info(fmt.Sprintf("test log: %d", i))
		}
	}
}
