package buffer

import "github.com/franela/goreq"

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

type ResponseUpdate struct {
	Success          bool
	BufferCount      int
	BufferPercentage int
	Updates          []Update
}

func (c *Client) CreateUpdate(text string, profileIds []string, options map[string]interface{}) []Update {

	params := options
	params["Text"] = text
	params["ProfileIds"] = profileIds

	request, err := goreq.Request{
		Method:      "POST",
		Uri:         c.Url + "/updates/create.json",
		Body:        params,
		Accept:      "application/json",
		ContentType: "application/json",
	}.Do()

	if err != nil {
		panic(err)
	}

	response := new(ResponseUpdate)
	err = request.Body.FromJsonTo(&response)

	if err != nil {
		panic(err)
	}

	return response.Updates
}

func NewClient(accessToken string) *Client {
	return &Client{Url: URL, AccessToken: accessToken}
}
