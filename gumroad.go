package license

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func Gumroad(productUrl, key string) (License, error) {
	lic := License{
		Active: false,
	}
	form := url.Values{
		"product_permalink": {productUrl},
		"license_key":       {key},
	}
	resp, err := http.PostForm("https://api.gumroad.com/v2/licenses/verify", form)
	if err != nil {
		return lic, err
	}

	r := &gumroadResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return lic, err
	}

	var variants []string = []string{}
	if r.Purchase.Variants != "" {
		variants = strings.Split(r.Purchase.Variants[1:len(r.Purchase.Variants)-1], ",")
	}

	return License{
		Active: r.Success && !(r.Purchase.Refunded || r.Purchase.Chargebacked),
		Context: map[string]interface{}{
			"variants": variants,
			"response": r,
		},
	}, nil
}

type gumroadResponse struct {
	Success  bool  `json:"success"`
	Uses     int64 `json:"uses"`
	Purchase struct {
		id                      string
		ProductName             string                 `json:"product_name"`
		CreatedAt               string                 `json:"created_at"`
		FullName                string                 `json:"full_name"`
		Variants                string                 `json:"variants"`
		Refunded                bool                   `json:"refunded"`
		Chargebacked            bool                   `json:"chargebacked"`
		SubscriptionCancelledAt string                 `json:"subscription_cancelled_at"`
		SubscriptionFailedAt    string                 `json:"subscription_failed_at"`
		CustomFields            map[string]interface{} `json:"custom_fields"`
		Email                   string                 `json:"email"`
	} `json:"purchase"`
}
