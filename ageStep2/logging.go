package main

import (
	"golang.org/x/net/context"
  "github.com/go-kit/kit/endpoint"
  "github.com/go-kit/kit/log"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func loggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
