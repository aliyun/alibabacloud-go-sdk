// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *CreateOAuthTokenRequest
	GetClientId() *string
	SetClientSecret(v string) *CreateOAuthTokenRequest
	GetClientSecret() *string
	SetCode(v string) *CreateOAuthTokenRequest
	GetCode() *string
	SetGrantType(v string) *CreateOAuthTokenRequest
	GetGrantType() *string
	SetLogin(v string) *CreateOAuthTokenRequest
	GetLogin() *string
	SetScope(v string) *CreateOAuthTokenRequest
	GetScope() *string
}

type CreateOAuthTokenRequest struct {
	// clientId
	//
	// This parameter is required.
	//
	// example:
	//
	// dc7e0b3c00a3e58f46
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// client_secret
	//
	// This parameter is required.
	//
	// example:
	//
	// a433294edea39cae7e7870
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	// example:
	//
	// 86df532f74454e189740d100ac97f4b9
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// code
	GrantType *string `json:"grantType,omitempty" xml:"grantType,omitempty"`
	// example:
	//
	// username
	Login *string `json:"login,omitempty" xml:"login,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// read:repo
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s CreateOAuthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CreateOAuthTokenRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateOAuthTokenRequest) GetCode() *string {
	return s.Code
}

func (s *CreateOAuthTokenRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *CreateOAuthTokenRequest) GetLogin() *string {
	return s.Login
}

func (s *CreateOAuthTokenRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateOAuthTokenRequest) SetClientId(v string) *CreateOAuthTokenRequest {
	s.ClientId = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetClientSecret(v string) *CreateOAuthTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetCode(v string) *CreateOAuthTokenRequest {
	s.Code = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetGrantType(v string) *CreateOAuthTokenRequest {
	s.GrantType = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetLogin(v string) *CreateOAuthTokenRequest {
	s.Login = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetScope(v string) *CreateOAuthTokenRequest {
	s.Scope = &v
	return s
}

func (s *CreateOAuthTokenRequest) Validate() error {
	return dara.Validate(s)
}
