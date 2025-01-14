// Copyright 2020 Akiomi Kamakura
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vimeo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/xup6m6fu04/vimeo-dl/config"
)

type Client struct {
	Client    *http.Client
	UserAgent string
	Referer   string
}

func NewClient() *Client {
	client := Client{}
	client.Client = http.DefaultClient
	client.UserAgent = "vimeo-dl/" + config.Version
	client.Referer = "https://www.yes588.com.tw/"

	return &client
}

func (c *Client) get(url *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Referer", c.Referer)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetMasterJson(url *url.URL) (*MasterJson, error) {
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	jsonBlob, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	masterJson := new(MasterJson)
	err = json.Unmarshal(jsonBlob, &masterJson)
	if err != nil {
		return nil, err
	}

	return masterJson, nil
}

func (c *Client) Download(url *url.URL, output io.Writer) error {
	res, err := c.get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(output, res.Body)
	if err != nil {
		return err
	}

	return nil
}
