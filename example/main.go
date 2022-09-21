package main

import (
	"fmt"

	"github.com/xmh19936688/go-requester/requester"
)

type Data struct {
	ID string `json:"id"`
}

func main() {
	urlStr := "https://domain.com/api/path"
	var data = Data{
		ID: "id",
	}

	req := requester.New().URL(urlStr).POST().AddHeaders([][2]string{
		{"Content-Type", "application/json;charset=UTF-8"},
	}).RequestJson(data).Do()

	if req.Err() != nil {
		fmt.Println(req.Err().Error())
		return
	}
	fmt.Println(string(req.Result()))
}
