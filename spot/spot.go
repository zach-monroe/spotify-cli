package spot

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetAuth() (string, error) {
	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")
	if client_secret == "" || client_id == "" {
		log.Print("no environment files found in .env")
		return "", errors.New("No Environment variables")
	}
	tokenURL := "https://accounts.spotify.com/api/token"
	data := url.Values{}

	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", tokenURL, nil)
	if err != nil {
		return "", errors.New("Invalid HTTP request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(client_id, client_secret)
	req.Body = io.NopCloser(strings.NewReader(data.Encode()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response
	body, _ := io.ReadAll(resp.Body)
	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}
	return tokenResp.AccessToken, nil

}
