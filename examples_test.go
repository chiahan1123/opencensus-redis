package opencensus_redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/kit-x/opencensus-redis/ochook"
)

func Example_NewClient() {
	redisOpts := redis.Options{
		Addr: "127.0.0.1:6379",
	}
	traceOptions := []ochook.TraceOption{
		ochook.WithAllowRoot(true),
	}

	client := redis.NewClient(&redisOpts)
	client.AddHook(ochook.New(traceOptions...))
	client.Ping(context.Background())
}
