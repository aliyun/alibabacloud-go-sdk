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
	SetScopes(v []*string) *CreateUserPoolClientRequest
	GetScopes() []*string
	SetSecretRequired(v bool) *CreateUserPoolClientRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *CreateUserPoolClientRequest
	GetUserPoolName() *string
}

type CreateUserPoolClientRequest struct {
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
	RedirectURIs []*string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty" type:"Repeated"`
	// The validity period of the refresh token, in seconds. Valid values: 7200 to 31536000 (90 days). Default value: 2592000 (30 days).
	//
	// example:
	//
	// 86400
	RefreshTokenValidity *string `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The scopes granted to the application. You can specify multiple scopes.
	Scopes []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
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

func (s *CreateUserPoolClientRequest) GetScopes() []*string {
	return s.Scopes
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

func (s *CreateUserPoolClientRequest) SetScopes(v []*string) *CreateUserPoolClientRequest {
	s.Scopes = v
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
