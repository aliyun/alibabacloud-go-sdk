// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenValidity(v string) *CreateUserPoolClientRequest
	GetAccessTokenValidity() *string
	SetClientName(v string) *CreateUserPoolClientRequest
	GetClientName() *string
	SetEnforcePKCE(v bool) *CreateUserPoolClientRequest
	GetEnforcePKCE() *bool
	SetRedirectURIs(v []*string) *CreateUserPoolClientRequest
	GetRedirectURIs() []*string
	SetRefreshTokenValidity(v string) *CreateUserPoolClientRequest
	GetRefreshTokenValidity() *string
	SetSecretRequired(v bool) *CreateUserPoolClientRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *CreateUserPoolClientRequest
	GetUserPoolName() *string
}

type CreateUserPoolClientRequest struct {
	// example:
	//
	// 3600
	AccessTokenValidity *string `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// example:
	//
	// true
	EnforcePKCE  *bool     `json:"EnforcePKCE,omitempty" xml:"EnforcePKCE,omitempty"`
	RedirectURIs []*string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty" type:"Repeated"`
	// example:
	//
	// 86400
	RefreshTokenValidity *string `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateUserPoolClientRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolClientRequest) GoString() string {
	return s.String()
}

func (s *CreateUserPoolClientRequest) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *CreateUserPoolClientRequest) GetClientName() *string {
	return s.ClientName
}

func (s *CreateUserPoolClientRequest) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *CreateUserPoolClientRequest) GetRedirectURIs() []*string {
	return s.RedirectURIs
}

func (s *CreateUserPoolClientRequest) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *CreateUserPoolClientRequest) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *CreateUserPoolClientRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserPoolClientRequest) SetAccessTokenValidity(v string) *CreateUserPoolClientRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateUserPoolClientRequest) SetClientName(v string) *CreateUserPoolClientRequest {
	s.ClientName = &v
	return s
}

func (s *CreateUserPoolClientRequest) SetEnforcePKCE(v bool) *CreateUserPoolClientRequest {
	s.EnforcePKCE = &v
	return s
}

func (s *CreateUserPoolClientRequest) SetRedirectURIs(v []*string) *CreateUserPoolClientRequest {
	s.RedirectURIs = v
	return s
}

func (s *CreateUserPoolClientRequest) SetRefreshTokenValidity(v string) *CreateUserPoolClientRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateUserPoolClientRequest) SetSecretRequired(v bool) *CreateUserPoolClientRequest {
	s.SecretRequired = &v
	return s
}

func (s *CreateUserPoolClientRequest) SetUserPoolName(v string) *CreateUserPoolClientRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserPoolClientRequest) Validate() error {
	return dara.Validate(s)
}
