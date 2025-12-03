// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *GenerateTokenRequest
	GetClientId() *string
	SetClientSecret(v string) *GenerateTokenRequest
	GetClientSecret() *string
	SetCode(v string) *GenerateTokenRequest
	GetCode() *string
	SetCodeVerifier(v string) *GenerateTokenRequest
	GetCodeVerifier() *string
	SetDeviceCode(v string) *GenerateTokenRequest
	GetDeviceCode() *string
	SetExclusiveTag(v string) *GenerateTokenRequest
	GetExclusiveTag() *string
	SetGrantType(v string) *GenerateTokenRequest
	GetGrantType() *string
	SetPassword(v string) *GenerateTokenRequest
	GetPassword() *string
	SetRedirectUri(v string) *GenerateTokenRequest
	GetRedirectUri() *string
	SetRefreshToken(v string) *GenerateTokenRequest
	GetRefreshToken() *string
	SetScope(v string) *GenerateTokenRequest
	GetScope() *string
	SetUsername(v string) *GenerateTokenRequest
	GetUsername() *string
}

type GenerateTokenRequest struct {
	// The client ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The client secret. This parameter is required if grant_type is set to client_credentials.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The authorization code. This parameter is required if grant_type is set to authorization_code.
	//
	// example:
	//
	// xxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The verification code.
	//
	// example:
	//
	// xxx
	CodeVerifier *string `json:"code_verifier,omitempty" xml:"code_verifier,omitempty"`
	// The device code. This parameter is required if grant_type is set to authorization_code.urn:ietf:params:oauth:grant-type:device_code.
	//
	// example:
	//
	// xxxx
	DeviceCode *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// The excluded tags.
	//
	// example:
	//
	// ATxxx
	ExclusiveTag *string `json:"exclusive_tag,omitempty" xml:"exclusive_tag,omitempty"`
	// The supported authorization types are as follows:
	//
	// - client_credentials:Client credentials flow, requires client_id and client_secret.
	//
	// - refresh_token:Refresh token flow.
	//
	// - authorization_code:Authorization code flow.
	//
	// - urn:ietf:params:oauth:grant-type:device_code:Device authorization flow.
	//
	// - password:Password (Resource Owner Password Credentials) flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_credentials
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// The username. This parameter is required if grant_type is set to password. The password authentication type is not supported.
	//
	// example:
	//
	// xxxxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The redirect URI. This parameter is required if grant_type is set to authorization_code. The value of this parameter must be the same as the redirect URI in the request to obtain the authorization code.
	//
	// example:
	//
	// xxx
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The refreshed token. This parameter is required if grant_type is set to refresh_token.
	//
	// example:
	//
	// ATxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The authorization scope. Valid values:
	//
	// 	- openid
	//
	// 	- email
	//
	// 	- phone
	//
	// 	- profile
	//
	// example:
	//
	// xxxx
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The username. This parameter is required if grant_type is set to password. The password authentication type is not supported.
	//
	// example:
	//
	// uesrname_001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GenerateTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GenerateTokenRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *GenerateTokenRequest) GetCode() *string {
	return s.Code
}

func (s *GenerateTokenRequest) GetCodeVerifier() *string {
	return s.CodeVerifier
}

func (s *GenerateTokenRequest) GetDeviceCode() *string {
	return s.DeviceCode
}

func (s *GenerateTokenRequest) GetExclusiveTag() *string {
	return s.ExclusiveTag
}

func (s *GenerateTokenRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *GenerateTokenRequest) GetPassword() *string {
	return s.Password
}

func (s *GenerateTokenRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *GenerateTokenRequest) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *GenerateTokenRequest) GetScope() *string {
	return s.Scope
}

func (s *GenerateTokenRequest) GetUsername() *string {
	return s.Username
}

func (s *GenerateTokenRequest) SetClientId(v string) *GenerateTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GenerateTokenRequest) SetClientSecret(v string) *GenerateTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *GenerateTokenRequest) SetCode(v string) *GenerateTokenRequest {
	s.Code = &v
	return s
}

func (s *GenerateTokenRequest) SetCodeVerifier(v string) *GenerateTokenRequest {
	s.CodeVerifier = &v
	return s
}

func (s *GenerateTokenRequest) SetDeviceCode(v string) *GenerateTokenRequest {
	s.DeviceCode = &v
	return s
}

func (s *GenerateTokenRequest) SetExclusiveTag(v string) *GenerateTokenRequest {
	s.ExclusiveTag = &v
	return s
}

func (s *GenerateTokenRequest) SetGrantType(v string) *GenerateTokenRequest {
	s.GrantType = &v
	return s
}

func (s *GenerateTokenRequest) SetPassword(v string) *GenerateTokenRequest {
	s.Password = &v
	return s
}

func (s *GenerateTokenRequest) SetRedirectUri(v string) *GenerateTokenRequest {
	s.RedirectUri = &v
	return s
}

func (s *GenerateTokenRequest) SetRefreshToken(v string) *GenerateTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenRequest) SetScope(v string) *GenerateTokenRequest {
	s.Scope = &v
	return s
}

func (s *GenerateTokenRequest) SetUsername(v string) *GenerateTokenRequest {
	s.Username = &v
	return s
}

func (s *GenerateTokenRequest) Validate() error {
	return dara.Validate(s)
}
