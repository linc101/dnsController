package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/kataras/iris"
)

func main() {

	config, err := ReadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Etcd.Protocol)

	cfg := client.Config{
		Endpoints: []string{config.GetEndpoint()},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	app := iris.Default()
	app.Get("/values", func(ctx iris.Context) {

		path := ctx.URLParamDefault("key", "/skydns")
		resp, err := kapi.Get(context.Background(), path, nil)
		if err != nil {
			log.Println(err)
			ctx.JSON(iris.Map{
				"error": err,
			})
		} else {
			ctx.JSON(iris.Map{
				"nodes": resp.Node.Nodes,
				"value": resp.Node.Value,
			})
		}

	})
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":" + config.GetWebPort()))
}
