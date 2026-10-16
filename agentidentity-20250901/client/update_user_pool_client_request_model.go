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
	SetScopes(v []*string) *UpdateUserPoolClientRequest
	GetScopes() []*string
	SetSecretRequired(v bool) *UpdateUserPoolClientRequest
	GetSecretRequired() *bool
	SetUserPoolName(v string) *UpdateUserPoolClientRequest
	GetUserPoolName() *string
}

type UpdateUserPoolClientRequest struct {
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
	RedirectURIs []*string `json:"RedirectURIs,omitempty" xml:"RedirectURIs,omitempty" type:"Repeated"`
	// The validity period of the refresh token, in seconds. Valid values: 7200 to 31536000 (90 days). Default value: 2592000 (30 days).
	//
	// example:
	//
	// 86400
	RefreshTokenValidity *string `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The scopes granted to the application. You can specify multiple scopes.
	Scopes []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
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

func (s *UpdateUserPoolClientRequest) GetScopes() []*string {
	return s.Scopes
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

func (s *UpdateUserPoolClientRequest) SetScopes(v []*string) *UpdateUserPoolClientRequest {
	s.Scopes = v
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
