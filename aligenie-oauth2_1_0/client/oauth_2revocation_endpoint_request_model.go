// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2RevocationEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *OAuth2RevocationEndpointRequest
	GetToken() *string
	SetTokenTypeHint(v string) *OAuth2RevocationEndpointRequest
	GetTokenTypeHint() *string
}

type OAuth2RevocationEndpointRequest struct {
	// example:
	//
	// UJMiksSwuMJvwXrJLULMykSw6qZ6VqaxOkN4qd5cW1Q4HhsLxvUR5xVOIv1WB3br5LoP20lPa8xiYLSMbt8JqHACXdSdw7fNkhRTIHnadxWW5jfDg7BELUB0FcFfPiv0
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// refresh_token
	TokenTypeHint *string `json:"TokenTypeHint,omitempty" xml:"TokenTypeHint,omitempty"`
}

func (s OAuth2RevocationEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s OAuth2RevocationEndpointRequest) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointRequest) GetToken() *string {
	return s.Token
}

func (s *OAuth2RevocationEndpointRequest) GetTokenTypeHint() *string {
	return s.TokenTypeHint
}

func (s *OAuth2RevocationEndpointRequest) SetToken(v string) *OAuth2RevocationEndpointRequest {
	s.Token = &v
	return s
}

func (s *OAuth2RevocationEndpointRequest) SetTokenTypeHint(v string) *OAuth2RevocationEndpointRequest {
	s.TokenTypeHint = &v
	return s
}

func (s *OAuth2RevocationEndpointRequest) Validate() error {
	return dara.Validate(s)
}
