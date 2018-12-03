package qiita

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Article is a data fetching from qiita
type Article struct {
	URL            string `json:"url"`
	Title          string `json:"title"`
	LikesCount     int    `json:"likes_count"`
	ReactionsCount int    `json:"reactions_count"`
	PageViewsCount int    `json:"page_views_count"`
}

// GetUserArticles get articles of the user specified by id as argment
func GetUserArticles(accessToken string, id string) ([]Article, error) {
	if len(accessToken) == 0 {
		return nil, errors.New("access token is missing")
	}

	baseURL := "https://qiita.com/api/v2/"
	action := "items"
	varParam := "?query=user:" + id

	endpointURL, err := url.Parse(baseURL + action + varParam)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse url")
	}

	b, err := json.Marshal(Article{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal blank Article")
	}

	// HTTPリクエストの作成
	resp, err := http.DefaultClient.Do(&http.Request{
		URL:    endpointURL,
		Method: "GET",
		Header: http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + accessToken},
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to receive the response from qiita")
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read data from a response from qiita")
	}

	var articles []Article

	if err := json.Unmarshal(b, &articles); err != nil {
		return nil, errors.Wrap(err, "JSON Unmarshal error")
	}

	return articles, nil
}
