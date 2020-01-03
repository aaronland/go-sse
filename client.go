package sse

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
)

type Client struct {
	client   *http.Client
	endpoint string
}

type ClientFunc func([]byte) error

func NewClient(endpoint string) (*Client, error) {

	client := &http.Client{}

	l := Client{
		client:   client,
		endpoint: endpoint,
	}

	return &l, nil
}

func (l *Client) Listen(callback ClientFunc) error {

	req, err := http.NewRequest("GET", l.endpoint, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "text/event-stream")

	res, err := l.client.Do(req)

	if err != nil {
		return err
	}

	br := bufio.NewReader(res.Body)
	defer res.Body.Close()

	delim := []byte{':', ' '}

	for {
		bs, err := br.ReadBytes('\n')

		if err != nil && err != io.EOF {
			return err
		}

		if len(bs) < 2 {
			continue
		}

		spl := bytes.Split(bs, delim)

		if len(spl) < 2 {
			continue
		}

		callback(spl[1])
	}
}
