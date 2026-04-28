// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssertion(v string) *TokenRequest
	GetAssertion() *string
	SetClientId(v string) *TokenRequest
	GetClientId() *string
	SetClientSecret(v string) *TokenRequest
	GetClientSecret() *string
	SetCode(v string) *TokenRequest
	GetCode() *string
	SetGrantType(v string) *TokenRequest
	GetGrantType() *string
	SetRedirectUri(v string) *TokenRequest
	GetRedirectUri() *string
	SetRefreshToken(v string) *TokenRequest
	GetRefreshToken() *string
}

type TokenRequest struct {
	// The JWT assertion that is signed by using the JWT private key. The JWT assertion contains the information about the user to be authorized and the authorization parameters. For more information about the structure of the JWT assertion, see JWTPayload. This parameter is required if grant_type is set to urn:ietf:params:oauth:grant-type:jwt-bearer.
	//
	// example:
	//
	// ey***asd
	Assertion *string `json:"assertion,omitempty" xml:"assertion,omitempty"`
	// The AppId of the application that is created in the Drive and Photo Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Zu***flH
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The AppSecret of the application that is created in the Drive and Photo Service console. This parameter is required if the application is of the WebServer type.
	//
	// example:
	//
	// 80D***3i5
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The authorization code in the redirect URI that is specified after the authorization process is complete. This parameter is required if grant_type is set to authorization_code.
	//
	// example:
	//
	// 0045157fa8e24f4f9a0d9e3ff158c1e0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The method that is used to generate an access token. Valid values:
	//
	// authorization_code: generates an access token by using the authorization code that is returned after the authorization process is complete.
	//
	// refresh_token: generates an access token by using the refresh token that is returned after the authorization process is complete.
	//
	// urn:ietf:params:oauth:grant-type:jwt-bearer: generates an access token by using a JWT assertion.
	//
	// This parameter is required.
	//
	// example:
	//
	// refresh_token
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// The redirect URI that is specified when you initiate the authorization request. This parameter is required if grant_type is set to authorization_code.
	//
	// example:
	//
	// https://aliyun.com/pds
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The refresh token that is used to refresh the access token. This parameter is required if grant_type is set to refresh_token.
	//
	// example:
	//
	// 399623e13353490391266c7d48a13ed1
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}

func (s TokenRequest) String() string {
	return dara.Prettify(s)
}

func (s TokenRequest) GoString() string {
	return s.String()
}

func (s *TokenRequest) GetAssertion() *string {
	return s.Assertion
}

func (s *TokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *TokenRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *TokenRequest) GetCode() *string {
	return s.Code
}

func (s *TokenRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *TokenRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *TokenRequest) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *TokenRequest) SetAssertion(v string) *TokenRequest {
	s.Assertion = &v
	return s
}

func (s *TokenRequest) SetClientId(v string) *TokenRequest {
	s.ClientId = &v
	return s
}

func (s *TokenRequest) SetClientSecret(v string) *TokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *TokenRequest) SetCode(v string) *TokenRequest {
	s.Code = &v
	return s
}

func (s *TokenRequest) SetGrantType(v string) *TokenRequest {
	s.GrantType = &v
	return s
}

func (s *TokenRequest) SetRedirectUri(v string) *TokenRequest {
	s.RedirectUri = &v
	return s
}

func (s *TokenRequest) SetRefreshToken(v string) *TokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *TokenRequest) Validate() error {
	return dara.Validate(s)
}
