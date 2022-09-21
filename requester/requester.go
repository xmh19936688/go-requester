// Package requester Usage:
//
// urlStr := "https://domain.com/api/path"
//
//	var data = Data{
//	    ID: "id",
//	}
//
//	req := requester.New().URL(urlStr).POST().AddHeaders([][2]string{
//	    {"Content-Type", "application/json;charset=UTF-8"},
//	}).RequestJson(data).Do()
//
//	if req.Err() != nil {
//	    fmt.Println(req.Err().Error())
//	    return
//	}
//
// res := string(req.Result())
package requester

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type requester struct {
	url     string
	method  string
	reqData io.Reader
	headers [][2]string
	cookies []*http.Cookie
	res     []byte
	err     error
}

func New() *requester {
	return &requester{}
}

func (h *requester) URL(url string) *requester {
	h.url = url
	return h
}

func (h *requester) POST() *requester {
	h.method = "POST"
	return h
}

func (h *requester) GET() *requester {
	h.method = "GET"
	return h
}

func (h *requester) AddHeader(key, value string) *requester {
	h.headers = append(h.headers, [2]string{key, value})
	return h
}

func (h *requester) AddHeaders(pairs [][2]string) *requester {
	h.headers = append(h.headers, pairs...)
	return h
}

func (h *requester) AddCookie(cookie *http.Cookie) *requester {
	h.cookies = append(h.cookies, cookie)
	return h
}

func (h *requester) AddCookies(cookies []*http.Cookie) *requester {
	h.cookies = append(h.cookies, cookies...)
	return h
}

// RequestData RequestJson 只能用一个，后调用的会覆盖先调用的
func (h *requester) RequestData(r io.Reader) *requester {
	h.reqData = r
	return h
}

// RequestJson RequestData 只能用一个，后调用的会覆盖先调用的
func (h *requester) RequestJson(data any) *requester {
	bs, err := json.Marshal(data)
	if err != nil {
		h.err = err
		return h
	}

	h.reqData = bytes.NewReader(bs)
	return h
}

func (h *requester) Err() error {
	return h.err
}

func (h *requester) Result() []byte {
	return h.res
}

func (h *requester) Do() *requester {
	if h.err != nil {
		return h
	}

	request, err := http.NewRequest(h.method, h.url, h.reqData)
	if err != nil {
		h.err = err
		return h
	}

	for _, pair := range h.headers {
		request.Header.Add(pair[0], pair[1])
	}

	for _, cookie := range h.cookies {
		request.AddCookie(cookie)
	}

	cli := http.Client{}
	resp, err := cli.Do(request)
	if err != nil {
		h.err = err
		return h
	}

	h.res, h.err = io.ReadAll(resp.Body)
	return h
}
