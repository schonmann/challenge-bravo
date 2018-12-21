package external

import (
	"fmt"
	"github.com/schonmann/challenge-bravo/config"
	"github.com/schonmann/challenge-bravo/util"
)

type RetrieveQuotasStrategy interface {
	RetrieveRates() (*QuotasResponse, error)
}

type QuotasResponse struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Quotas    map[string]float64 `json:"rates"`
}

type OpenExchangeRatesStrategy struct{}

func (o OpenExchangeRatesStrategy) RetrieveRates() (*QuotasResponse, error) {
	apiConfig := config.Get().Worker.ExternalAPIs.OpenExchangeRates
	apiUrl := fmt.Sprintf("%slatest.json?app_id=%s&show_alternative=1", apiConfig.URL, apiConfig.APIKey)
	response := QuotasResponse{}
	if err := util.GetAndParseJSON(apiUrl, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

/**
  Strategy implemented to give API flexibility. Just returning
  OpenExchangeRatesStrategy for now.
*/

func GetRetrieveRatesStrategy() RetrieveQuotasStrategy {
	return OpenExchangeRatesStrategy{}
}