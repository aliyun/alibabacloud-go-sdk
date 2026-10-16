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
	SetScopesShrink(v string) *UpdateUserPoolClientShrinkRequest
	GetScopesShrink() *string
	SetSecretRequired(v bool) *UpdateUserPoolClientShrinkRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *UpdateUserPoolClientShrinkRequest
	GetUserPoolName() *string
}

type UpdateUserPoolClientShrinkRequest struct {
	// The validity period of the access token, in seconds. Valid values: 900 to 10800. Default value: 3600.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *string `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The name of the client.
	//
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// Specifies whether Proof Key for Code Exchange (PKCE) is enforced.
	//
	// example:
	//
	// true
	EnforcePKCE *bool `json:"EnforcePKCE,omitempty" xml:"EnforcePKCE,omitempty"`
	// The new list of OAuth callback URLs. Specifying this parameter replaces all existing callback URLs.
	RedirectURIsShrink *string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty"`
	// The validity period of the refresh token, in seconds. Valid values: 7200 to 31536000 (90 days). Default value: 2592000 (30 days).
	//
	// example:
	//
	// 86400
	RefreshTokenValidity *string `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The scopes granted to the application. You can specify multiple scopes.
	ScopesShrink *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
	// Specifies whether an application secret is required. A value of true indicates a confidential client, and a value of false indicates a public client.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The name of the user pool.
	//
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

func (s *UpdateUserPoolClientShrinkRequest) GetScopesShrink() *string {
	return s.ScopesShrink
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

func (s *UpdateUserPoolClientShrinkRequest) SetScopesShrink(v string) *UpdateUserPoolClientShrinkRequest {
	s.ScopesShrink = &v
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
