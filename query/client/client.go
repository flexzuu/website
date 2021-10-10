package client

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

type authedTransport struct {
	token   string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.wrapped.RoundTrip(req)
}

type client struct {
	graphql.Client
}

func New(endpoint, key string) client {
	httpClient := http.Client{
		Transport: &authedTransport{
			token:   key,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient(endpoint, &httpClient)
	return client{graphqlClient}
}

func (c client) Posts(ctx context.Context) (*PostsResponse, error) {
	return Posts(ctx, c)
}
