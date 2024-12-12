package api

import (
    "encoding/json"
    "errors"
    "net/http"
)

type ZenQuoteResponse struct {
    Q string `json:"q"`
    A string `json:"a"`
}

func GetMotivationalQuote() (string, error) {
    resp, err := http.Get("https://zenquotes.io/api/random")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("failed to fetch quote")
    }

    var quotes []ZenQuoteResponse
    if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
        return "", err
    }
    if len(quotes) == 0 {
        return "", errors.New("no quotes found")
    }
    quote := quotes[0]
    return quote.Q + " - " + quote.A, nil
}