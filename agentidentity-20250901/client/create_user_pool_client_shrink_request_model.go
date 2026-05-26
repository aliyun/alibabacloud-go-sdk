// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolClientShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenValidity(v string) *CreateUserPoolClientShrinkRequest
	GetAccessTokenValidity() *string
	SetClientName(v string) *CreateUserPoolClientShrinkRequest
	GetClientName() *string
	SetEnforcePKCE(v bool) *CreateUserPoolClientShrinkRequest
	GetEnforcePKCE() *bool
	SetRedirectURIsShrink(v string) *CreateUserPoolClientShrinkRequest
	GetRedirectURIsShrink() *string
	SetRefreshTokenValidity(v string) *CreateUserPoolClientShrinkRequest
	GetRefreshTokenValidity() *string
	SetSecretRequired(v bool) *CreateUserPoolClientShrinkRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *CreateUserPoolClientShrinkRequest
	GetUserPoolName() *string
}

type CreateUserPoolClientShrinkRequest struct {
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
	EnforcePKCE        *bool   `json:"EnforcePKCE,omitempty" xml:"EnforcePKCE,omitempty"`
	RedirectURIsShrink *string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty"`
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

func (s CreateUserPoolClientShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolClientShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserPoolClientShrinkRequest) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *CreateUserPoolClientShrinkRequest) GetClientName() *string {
	return s.ClientName
}

func (s *CreateUserPoolClientShrinkRequest) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *CreateUserPoolClientShrinkRequest) GetRedirectURIsShrink() *string {
	return s.RedirectURIsShrink
}

func (s *CreateUserPoolClientShrinkRequest) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *CreateUserPoolClientShrinkRequest) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *CreateUserPoolClientShrinkRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserPoolClientShrinkRequest) SetAccessTokenValidity(v string) *CreateUserPoolClientShrinkRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) SetClientName(v string) *CreateUserPoolClientShrinkRequest {
	s.ClientName = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) SetEnforcePKCE(v bool) *CreateUserPoolClientShrinkRequest {
	s.EnforcePKCE = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) SetRedirectURIsShrink(v string) *CreateUserPoolClientShrinkRequest {
	s.RedirectURIsShrink = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) SetRefreshTokenValidity(v string) *CreateUserPoolClientShrinkRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) SetSecretRequired(v bool) *CreateUserPoolClientShrinkRequest {
	s.SecretRequired = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) SetUserPoolName(v string) *CreateUserPoolClientShrinkRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserPoolClientShrinkRequest) Validate() error {
	return dara.Validate(s)
}
