package main

import (
"fmt"
"log"
"github.com/buaazp/fasthttprouter"
"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
  fmt.Fprint(ctx, "Hi there, I love !")
}

func Hello(ctx *fasthttp.RequestCtx) {
  fmt.Fprintf(ctx, "Hi there, I love %s !", ctx.UserValue("name"))
}

func main() {
  router := fasthttprouter.New()
  router.GET("/", Index)
  router.GET("/:name", Hello)

  log.Fatal(fasthttp.ListenAndServe(":4000", router.Handler))
}
