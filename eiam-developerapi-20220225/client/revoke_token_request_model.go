// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *RevokeTokenRequest
	GetClientId() *string
	SetClientSecret(v string) *RevokeTokenRequest
	GetClientSecret() *string
	SetToken(v string) *RevokeTokenRequest
	GetToken() *string
	SetTokenTypeHint(v string) *RevokeTokenRequest
	GetTokenTypeHint() *string
}

type RevokeTokenRequest struct {
	// The client ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The client secret.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The token to be revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// ATxxxx
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The type of the token. Valid values: access_token refresh_token
	//
	// example:
	//
	// access_token
	TokenTypeHint *string `json:"token_type_hint,omitempty" xml:"token_type_hint,omitempty"`
}

func (s RevokeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RevokeTokenRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *RevokeTokenRequest) GetToken() *string {
	return s.Token
}

func (s *RevokeTokenRequest) GetTokenTypeHint() *string {
	return s.TokenTypeHint
}

func (s *RevokeTokenRequest) SetClientId(v string) *RevokeTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RevokeTokenRequest) SetClientSecret(v string) *RevokeTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *RevokeTokenRequest) SetToken(v string) *RevokeTokenRequest {
	s.Token = &v
	return s
}

func (s *RevokeTokenRequest) SetTokenTypeHint(v string) *RevokeTokenRequest {
	s.TokenTypeHint = &v
	return s
}

func (s *RevokeTokenRequest) Validate() error {
	return dara.Validate(s)
}
