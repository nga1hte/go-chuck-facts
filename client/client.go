package client

import (
	"encoding/json"
	"go-chuck-facts/model"
	_ "io"
	"net/http"
	_ "os"
	"time"
)

const (
	BaseURL              string        = "https://api.chucknorris.io/jokes/random"
	DefaultClientTimeout time.Duration = 30 * time.Second
)

type CNClient struct {
	client  *http.Client
	baseURL string
}

func NewCNClient() *CNClient {
	return &CNClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

func (c *CNClient) SetTimeout(d time.Duration) {
	c.client.Timeout = d
}

func (c *CNClient) Fetch(category string) (model.Fact, error) {
	url := BaseURL
	if category != "" {
		url = BaseURL + "?category=" + category
	}
	resp, err := c.client.Get(url)
	if err != nil {
		return model.Fact{}, err
	}
	defer resp.Body.Close()
	var factResp model.Fact
	if err := json.NewDecoder(resp.Body).Decode(&factResp); err != nil {
		return model.Fact{}, err
	}
	return factResp, err
}
