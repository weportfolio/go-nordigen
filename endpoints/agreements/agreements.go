package agreements

import "time"

type AgreementRequestBody struct {
	InstitutionID      string   `json:"institution_id"`
	MaxHistoricalDays  string   `json:"max_historical_days"`
	AccessValidForDays string   `json:"access_valid_for_days"`
	AccessScope        []string `json:"access_scope"`
}

type Agreement struct {
	ID                 string    `json:"id"`
	Created            string    `json:"created"`
	InstitutionID      string    `json:"institution_id"`
	MaxHistoricalDays  int       `json:"max_historical_days"`
	AccessValidForDays int       `json:"access_valid_for_days"`
	AccessScope        []string  `json:"access_scope"`
	Accepted           time.Time `json:"accepted"`
}