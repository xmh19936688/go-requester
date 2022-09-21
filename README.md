# go-requester

go-requester用于发送http请求，避免写重复代码

使用`go-requester`前：

```go
type Data struct {
	ID string `json:"id"`
}

func main() {
	urlStr := "https://domain.com/api/path"
	var data = Data{
		ID: "id",
	}
	bs, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewReader(bs))
	if err != nil {
		return
	}

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return
	}

	bs, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bs))
}
```

使用`go-requester`后：

```go
// go get github.com/xmh19936688/go-requester
import "github.com/xmh19936688/go-requester/requester"

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
```
