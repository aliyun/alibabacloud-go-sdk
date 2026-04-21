// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketAuthCodeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationEndpoint(v string) *HiMarketAuthCodeConfig
	GetAuthorizationEndpoint() *string
	SetClientId(v string) *HiMarketAuthCodeConfig
	GetClientId() *string
	SetClientSecret(v string) *HiMarketAuthCodeConfig
	GetClientSecret() *string
	SetIssuer(v string) *HiMarketAuthCodeConfig
	GetIssuer() *string
	SetJwkSetUri(v string) *HiMarketAuthCodeConfig
	GetJwkSetUri() *string
	SetRedirectUri(v string) *HiMarketAuthCodeConfig
	GetRedirectUri() *string
	SetScopes(v string) *HiMarketAuthCodeConfig
	GetScopes() *string
	SetTokenEndpoint(v string) *HiMarketAuthCodeConfig
	GetTokenEndpoint() *string
	SetUserInfoEndpoint(v string) *HiMarketAuthCodeConfig
	GetUserInfoEndpoint() *string
}

type HiMarketAuthCodeConfig struct {
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" xml:"authorizationEndpoint,omitempty"`
	ClientId              *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	ClientSecret          *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	Issuer                *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	JwkSetUri             *string `json:"jwkSetUri,omitempty" xml:"jwkSetUri,omitempty"`
	RedirectUri           *string `json:"redirectUri,omitempty" xml:"redirectUri,omitempty"`
	Scopes                *string `json:"scopes,omitempty" xml:"scopes,omitempty"`
	TokenEndpoint         *string `json:"tokenEndpoint,omitempty" xml:"tokenEndpoint,omitempty"`
	UserInfoEndpoint      *string `json:"userInfoEndpoint,omitempty" xml:"userInfoEndpoint,omitempty"`
}

func (s HiMarketAuthCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketAuthCodeConfig) GoString() string {
	return s.String()
}

func (s *HiMarketAuthCodeConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *HiMarketAuthCodeConfig) GetClientId() *string {
	return s.ClientId
}

func (s *HiMarketAuthCodeConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *HiMarketAuthCodeConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *HiMarketAuthCodeConfig) GetJwkSetUri() *string {
	return s.JwkSetUri
}

func (s *HiMarketAuthCodeConfig) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *HiMarketAuthCodeConfig) GetScopes() *string {
	return s.Scopes
}

func (s *HiMarketAuthCodeConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *HiMarketAuthCodeConfig) GetUserInfoEndpoint() *string {
	return s.UserInfoEndpoint
}

func (s *HiMarketAuthCodeConfig) SetAuthorizationEndpoint(v string) *HiMarketAuthCodeConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetClientId(v string) *HiMarketAuthCodeConfig {
	s.ClientId = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetClientSecret(v string) *HiMarketAuthCodeConfig {
	s.ClientSecret = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetIssuer(v string) *HiMarketAuthCodeConfig {
	s.Issuer = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetJwkSetUri(v string) *HiMarketAuthCodeConfig {
	s.JwkSetUri = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetRedirectUri(v string) *HiMarketAuthCodeConfig {
	s.RedirectUri = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetScopes(v string) *HiMarketAuthCodeConfig {
	s.Scopes = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetTokenEndpoint(v string) *HiMarketAuthCodeConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *HiMarketAuthCodeConfig) SetUserInfoEndpoint(v string) *HiMarketAuthCodeConfig {
	s.UserInfoEndpoint = &v
	return s
}

func (s *HiMarketAuthCodeConfig) Validate() error {
	return dara.Validate(s)
}
