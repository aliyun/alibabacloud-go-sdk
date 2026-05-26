// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPoolClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenValidity(v string) *UpdateUserPoolClientRequest
	GetAccessTokenValidity() *string
	SetClientName(v string) *UpdateUserPoolClientRequest
	GetClientName() *string
	SetEnforcePKCE(v bool) *UpdateUserPoolClientRequest
	GetEnforcePKCE() *bool
	SetRedirectURIs(v []*string) *UpdateUserPoolClientRequest
	GetRedirectURIs() []*string
	SetRefreshTokenValidity(v string) *UpdateUserPoolClientRequest
	GetRefreshTokenValidity() *string
	SetSecretRequired(v bool) *UpdateUserPoolClientRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *UpdateUserPoolClientRequest
	GetUserPoolName() *string
}

type UpdateUserPoolClientRequest struct {
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

func (s UpdateUserPoolClientRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPoolClientRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPoolClientRequest) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *UpdateUserPoolClientRequest) GetClientName() *string {
	return s.ClientName
}

func (s *UpdateUserPoolClientRequest) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *UpdateUserPoolClientRequest) GetRedirectURIs() []*string {
	return s.RedirectURIs
}

func (s *UpdateUserPoolClientRequest) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *UpdateUserPoolClientRequest) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *UpdateUserPoolClientRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateUserPoolClientRequest) SetAccessTokenValidity(v string) *UpdateUserPoolClientRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *UpdateUserPoolClientRequest) SetClientName(v string) *UpdateUserPoolClientRequest {
	s.ClientName = &v
	return s
}

func (s *UpdateUserPoolClientRequest) SetEnforcePKCE(v bool) *UpdateUserPoolClientRequest {
	s.EnforcePKCE = &v
	return s
}

func (s *UpdateUserPoolClientRequest) SetRedirectURIs(v []*string) *UpdateUserPoolClientRequest {
	s.RedirectURIs = v
	return s
}

func (s *UpdateUserPoolClientRequest) SetRefreshTokenValidity(v string) *UpdateUserPoolClientRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *UpdateUserPoolClientRequest) SetSecretRequired(v bool) *UpdateUserPoolClientRequest {
	s.SecretRequired = &v
	return s
}

func (s *UpdateUserPoolClientRequest) SetUserPoolName(v string) *UpdateUserPoolClientRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateUserPoolClientRequest) Validate() error {
	return dara.Validate(s)
}
