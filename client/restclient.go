package client

type client struct {
}

type Client interface {
}

func NewClient() Client {
	return &client{}
}
