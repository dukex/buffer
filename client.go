package buffer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	URL = "https://api.bufferapp.com/1"
)

type Client struct {
	AccessToken string
	Url         string
}

type Update struct {
	Id        string
	Text      string
	ProfileId string
}

}

func (c *Client) CreateUpdate(text string, profileIds []string, options map[string]interface{}) []Update {

	params := url.Values{}
	params.Set("text", text)
	for _, p := range profileIds {
		params.Add("profile_ids[]", p)
	}

	bufferResponse := c.send("updates/create", params)

	var response struct {
		Success          bool
		BufferCount      int
		BufferPercentage int
		Updates          []Update
	}

	err := json.Unmarshal(bufferResponse, &response)

	if err != nil {
		panic(err)
	}

	return response.Updates
}

func (c *Client) send(resource string, params url.Values) []byte {
	urlEndpoint := c.Url + "/" + resource + ".json?access_token=" + c.AccessToken
	request, err := http.PostForm(urlEndpoint, params)
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()
	requestBodyByte, _ := ioutil.ReadAll(request.Body)

	return requestBodyByte
}

func NewClient(accessToken string) *Client {
	return &Client{Url: URL, AccessToken: accessToken}
}
