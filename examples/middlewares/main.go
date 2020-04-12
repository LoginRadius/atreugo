package main

import (
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	// Register before middlewares
	server.UseBefore(beforeGlobal)

	// Register after middlewares
	server.UseAfter(afterGlobal)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Middlewares example")
	}).UseBefore(beforeView).UseAfter(afterView)

	// Run
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
