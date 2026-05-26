// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPoolClientShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenValidity(v string) *UpdateUserPoolClientShrinkRequest
	GetAccessTokenValidity() *string
	SetClientName(v string) *UpdateUserPoolClientShrinkRequest
	GetClientName() *string
	SetEnforcePKCE(v bool) *UpdateUserPoolClientShrinkRequest
	GetEnforcePKCE() *bool
	SetRedirectURIsShrink(v string) *UpdateUserPoolClientShrinkRequest
	GetRedirectURIsShrink() *string
	SetRefreshTokenValidity(v string) *UpdateUserPoolClientShrinkRequest
	GetRefreshTokenValidity() *string
	SetSecretRequired(v bool) *UpdateUserPoolClientShrinkRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *UpdateUserPoolClientShrinkRequest
	GetUserPoolName() *string
}

type UpdateUserPoolClientShrinkRequest struct {
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

func (s UpdateUserPoolClientShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPoolClientShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPoolClientShrinkRequest) GetAccessTokenValidity() *string {
	return s.AccessTokenValidity
}

func (s *UpdateUserPoolClientShrinkRequest) GetClientName() *string {
	return s.ClientName
}

func (s *UpdateUserPoolClientShrinkRequest) GetEnforcePKCE() *bool {
	return s.EnforcePKCE
}

func (s *UpdateUserPoolClientShrinkRequest) GetRedirectURIsShrink() *string {
	return s.RedirectURIsShrink
}

func (s *UpdateUserPoolClientShrinkRequest) GetRefreshTokenValidity() *string {
	return s.RefreshTokenValidity
}

func (s *UpdateUserPoolClientShrinkRequest) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *UpdateUserPoolClientShrinkRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateUserPoolClientShrinkRequest) SetAccessTokenValidity(v string) *UpdateUserPoolClientShrinkRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) SetClientName(v string) *UpdateUserPoolClientShrinkRequest {
	s.ClientName = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) SetEnforcePKCE(v bool) *UpdateUserPoolClientShrinkRequest {
	s.EnforcePKCE = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) SetRedirectURIsShrink(v string) *UpdateUserPoolClientShrinkRequest {
	s.RedirectURIsShrink = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) SetRefreshTokenValidity(v string) *UpdateUserPoolClientShrinkRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) SetSecretRequired(v bool) *UpdateUserPoolClientShrinkRequest {
	s.SecretRequired = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) SetUserPoolName(v string) *UpdateUserPoolClientShrinkRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateUserPoolClientShrinkRequest) Validate() error {
	return dara.Validate(s)
}
