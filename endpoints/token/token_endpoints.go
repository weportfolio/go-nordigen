package token

import (
	"github.com/weportfolio/go-nordigen/consts"
)

func (c Client) New() (*Token, error) {
	var token Token
	accessCreds := map[string]string{
		"secret_id":  c.HTTP.APISecretID,
		"secret_key": c.HTTP.APISecretKey,
	}

	err := c.HTTP.Post(consts.TokenNewPath, nil, accessCreds, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (c Client) Refresh(refreshToken string) (*Token, error) {
	var token Token
	refreshCreds := map[string]string{
		"refresh": refreshToken,
	}

	err := c.HTTP.Post(consts.TokenRefreshPath, nil, refreshCreds, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}