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
	// The client secret. This parameter is required when \\`grant_type\\` is \\`client_credentials\\` and the \\`client_secret_post\\` method is used.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The authorization code. This parameter is required when \\`grant_type\\` is \\`authorization_code\\`.
	//
	// example:
	//
	// xxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The code verifier. This is used in the authorization code grant type when PKCE is enabled.
	//
	// example:
	//
	// xxx
	CodeVerifier *string `json:"code_verifier,omitempty" xml:"code_verifier,omitempty"`
	// The device code. This parameter is required when \\`grant_type\\` is \\`urn:ietf:params:oauth:grant-type:device_code\\` (device flow).
	//
	// example:
	//
	// xxxx
	DeviceCode *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// The excluded tag.
	//
	// example:
	//
	// ATxxx
	ExclusiveTag *string `json:"exclusive_tag,omitempty" xml:"exclusive_tag,omitempty"`
	// The authorization grant type. The following types are supported:
	//
	// - \\`client_credentials\\`: Client credentials grant. Requires \\`client_id\\` and \\`client_secret\\`.
	//
	// - \\`refresh_token\\`: Refresh token grant.
	//
	// - \\`authorization_code\\`: Authorization code grant.
	//
	// - \\`urn:ietf:params:oauth:grant-type:device_code\\`: Device flow.
	//
	// - \\`password\\`: Password grant.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_credentials
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// The username. This parameter is required for password mode.
	//
	// example:
	//
	// xxxxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The redirection URI. This parameter is required for the authorization code grant type. It must match the redirection URI in the request to get the authorization code.
	//
	// example:
	//
	// xxx
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The refresh token. This parameter is required when \\`grant_type\\` is \\`refresh_token\\` (refresh token grant).
	//
	// example:
	//
	// ATxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The scope. This parameter is optional. Multiple values are supported. Separate multiple values with spaces.
	//
	// Valid values:
	//
	// - openid
	//
	// - email
	//
	// - phone
	//
	// - profile
	//
	// example:
	//
	// xxxx
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The username. This parameter is required for the password grant type.
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
