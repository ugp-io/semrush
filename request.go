package semrush

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	netURL "net/url"
	"strings"
)

var accessCode, refreshAccessCode string

type Client struct {
	APIKey     string
	APISecret  string
	AccessCode string
	Locations  LocationsService
}

type TokenRefreshResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func NewClient(apiKey, apiSecret, apiAccessCode string) *Client {

	c := &Client{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		AccessCode: apiAccessCode,
	}

	c.Locations = &LocationsServiceOp{client: c}

	return c

}

func (c *Client) TokenAccess(ctx context.Context) error {

	if refreshAccessCode == "" {
		refreshAccessCode = c.AccessCode
	}

	data := netURL.Values{}
	data.Set("client_id", c.APIKey)
	data.Set("client_secret", c.APISecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", refreshAccessCode)
	data.Set("redirect_uri", "%2Foauth2%2Fsuccess")

	req, err := http.NewRequest(
		"POST",
		"https://oauth.semrush.com/oauth2/access_token",
		strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var token TokenRefreshResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return err
	}

	refreshAccessCode = token.RefreshToken
	accessCode = token.AccessToken

	return nil
}

func (c *Client) Request(method string, url string, bodyJSON io.Reader, response *[]byte) error {

	httpReq, errNewRequest := http.NewRequest(method, url, bodyJSON)
	if errNewRequest != nil {
		return errNewRequest
	}

	httpReq.Header.Add("Authorization", "Bearer "+accessCode)

	client := &http.Client{}
	res, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	*response = bodyBytes

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errResp map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &errResp); err != nil {
			return fmt.Errorf("request failed with status %d: %s", res.StatusCode, string(bodyBytes))
		}
		return fmt.Errorf("request failed with status %d: %v", res.StatusCode, errResp)
	}

	return nil

}
