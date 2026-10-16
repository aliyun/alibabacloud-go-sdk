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
	SetScopesShrink(v string) *CreateUserPoolClientShrinkRequest
	GetScopesShrink() *string
	SetSecretRequired(v bool) *CreateUserPoolClientShrinkRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *CreateUserPoolClientShrinkRequest
	GetUserPoolName() *string
}

type CreateUserPoolClientShrinkRequest struct {
	// The validity period of the access token, in seconds. Valid values: 900 to 10800. Default value: 3600.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *string `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The name of the client. The name must be unique within the same user pool and must be 1 to 64 characters in length.
	//
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// Specifies whether to enforce Proof Key for Code Exchange (PKCE). Valid values:
	//
	// - true: Enforced. The code_challenge parameter is required in the authorization code flow.
	//
	// - false: Not enforced.
	//
	// example:
	//
	// true
	EnforcePKCE *bool `json:"EnforcePKCE,omitempty" xml:"EnforcePKCE,omitempty"`
	// The list of OAuth redirect URIs. After authorization is complete, the user\\"s browser is redirected to one of the URIs in the list.
	RedirectURIsShrink *string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty"`
	// The validity period of the refresh token, in seconds. Valid values: 7200 to 31536000 (90 days). Default value: 2592000 (30 days).
	//
	// example:
	//
	// 86400
	RefreshTokenValidity *string `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The scopes granted to the application. You can specify multiple scopes.
	ScopesShrink *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
	// Specifies whether an application secret is required. Valid values:
	//
	// - true: Required. This value applies to confidential clients such as server-side applications. You must call CreateClientSecret to create a secret.
	//
	// - false: Not required. This value applies to public clients such as single-page applications (SPAs) and mobile applications.
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

func (s *CreateUserPoolClientShrinkRequest) GetScopesShrink() *string {
	return s.ScopesShrink
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

func (s *CreateUserPoolClientShrinkRequest) SetScopesShrink(v string) *CreateUserPoolClientShrinkRequest {
	s.ScopesShrink = &v
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
