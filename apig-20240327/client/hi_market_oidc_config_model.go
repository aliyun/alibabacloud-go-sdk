// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketOidcConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCodeConfig(v *HiMarketOidcConfigAuthCodeConfig) *HiMarketOidcConfig
	GetAuthCodeConfig() *HiMarketOidcConfigAuthCodeConfig
	SetEnabled(v bool) *HiMarketOidcConfig
	GetEnabled() *bool
	SetGrantType(v string) *HiMarketOidcConfig
	GetGrantType() *string
	SetIdentityMapping(v *HiMarketOidcConfigIdentityMapping) *HiMarketOidcConfig
	GetIdentityMapping() *HiMarketOidcConfigIdentityMapping
	SetLogoUrl(v string) *HiMarketOidcConfig
	GetLogoUrl() *string
	SetName(v string) *HiMarketOidcConfig
	GetName() *string
	SetProvider(v string) *HiMarketOidcConfig
	GetProvider() *string
}

type HiMarketOidcConfig struct {
	AuthCodeConfig  *HiMarketOidcConfigAuthCodeConfig  `json:"authCodeConfig,omitempty" xml:"authCodeConfig,omitempty" type:"Struct"`
	Enabled         *bool                              `json:"enabled,omitempty" xml:"enabled,omitempty"`
	GrantType       *string                            `json:"grantType,omitempty" xml:"grantType,omitempty"`
	IdentityMapping *HiMarketOidcConfigIdentityMapping `json:"identityMapping,omitempty" xml:"identityMapping,omitempty" type:"Struct"`
	LogoUrl         *string                            `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	Name            *string                            `json:"name,omitempty" xml:"name,omitempty"`
	Provider        *string                            `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s HiMarketOidcConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketOidcConfig) GoString() string {
	return s.String()
}

func (s *HiMarketOidcConfig) GetAuthCodeConfig() *HiMarketOidcConfigAuthCodeConfig {
	return s.AuthCodeConfig
}

func (s *HiMarketOidcConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *HiMarketOidcConfig) GetGrantType() *string {
	return s.GrantType
}

func (s *HiMarketOidcConfig) GetIdentityMapping() *HiMarketOidcConfigIdentityMapping {
	return s.IdentityMapping
}

func (s *HiMarketOidcConfig) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *HiMarketOidcConfig) GetName() *string {
	return s.Name
}

func (s *HiMarketOidcConfig) GetProvider() *string {
	return s.Provider
}

func (s *HiMarketOidcConfig) SetAuthCodeConfig(v *HiMarketOidcConfigAuthCodeConfig) *HiMarketOidcConfig {
	s.AuthCodeConfig = v
	return s
}

func (s *HiMarketOidcConfig) SetEnabled(v bool) *HiMarketOidcConfig {
	s.Enabled = &v
	return s
}

func (s *HiMarketOidcConfig) SetGrantType(v string) *HiMarketOidcConfig {
	s.GrantType = &v
	return s
}

func (s *HiMarketOidcConfig) SetIdentityMapping(v *HiMarketOidcConfigIdentityMapping) *HiMarketOidcConfig {
	s.IdentityMapping = v
	return s
}

func (s *HiMarketOidcConfig) SetLogoUrl(v string) *HiMarketOidcConfig {
	s.LogoUrl = &v
	return s
}

func (s *HiMarketOidcConfig) SetName(v string) *HiMarketOidcConfig {
	s.Name = &v
	return s
}

func (s *HiMarketOidcConfig) SetProvider(v string) *HiMarketOidcConfig {
	s.Provider = &v
	return s
}

func (s *HiMarketOidcConfig) Validate() error {
	if s.AuthCodeConfig != nil {
		if err := s.AuthCodeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IdentityMapping != nil {
		if err := s.IdentityMapping.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketOidcConfigAuthCodeConfig struct {
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

func (s HiMarketOidcConfigAuthCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketOidcConfigAuthCodeConfig) GoString() string {
	return s.String()
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetClientId() *string {
	return s.ClientId
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetJwkSetUri() *string {
	return s.JwkSetUri
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetScopes() *string {
	return s.Scopes
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *HiMarketOidcConfigAuthCodeConfig) GetUserInfoEndpoint() *string {
	return s.UserInfoEndpoint
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetAuthorizationEndpoint(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetClientId(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.ClientId = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetClientSecret(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.ClientSecret = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetIssuer(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.Issuer = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetJwkSetUri(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.JwkSetUri = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetRedirectUri(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.RedirectUri = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetScopes(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.Scopes = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetTokenEndpoint(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) SetUserInfoEndpoint(v string) *HiMarketOidcConfigAuthCodeConfig {
	s.UserInfoEndpoint = &v
	return s
}

func (s *HiMarketOidcConfigAuthCodeConfig) Validate() error {
	return dara.Validate(s)
}

type HiMarketOidcConfigIdentityMapping struct {
	CustomFields  map[string]*string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	EmailField    *string            `json:"emailField,omitempty" xml:"emailField,omitempty"`
	UserIdField   *string            `json:"userIdField,omitempty" xml:"userIdField,omitempty"`
	UserNameField *string            `json:"userNameField,omitempty" xml:"userNameField,omitempty"`
}

func (s HiMarketOidcConfigIdentityMapping) String() string {
	return dara.Prettify(s)
}

func (s HiMarketOidcConfigIdentityMapping) GoString() string {
	return s.String()
}

func (s *HiMarketOidcConfigIdentityMapping) GetCustomFields() map[string]*string {
	return s.CustomFields
}

func (s *HiMarketOidcConfigIdentityMapping) GetEmailField() *string {
	return s.EmailField
}

func (s *HiMarketOidcConfigIdentityMapping) GetUserIdField() *string {
	return s.UserIdField
}

func (s *HiMarketOidcConfigIdentityMapping) GetUserNameField() *string {
	return s.UserNameField
}

func (s *HiMarketOidcConfigIdentityMapping) SetCustomFields(v map[string]*string) *HiMarketOidcConfigIdentityMapping {
	s.CustomFields = v
	return s
}

func (s *HiMarketOidcConfigIdentityMapping) SetEmailField(v string) *HiMarketOidcConfigIdentityMapping {
	s.EmailField = &v
	return s
}

func (s *HiMarketOidcConfigIdentityMapping) SetUserIdField(v string) *HiMarketOidcConfigIdentityMapping {
	s.UserIdField = &v
	return s
}

func (s *HiMarketOidcConfigIdentityMapping) SetUserNameField(v string) *HiMarketOidcConfigIdentityMapping {
	s.UserNameField = &v
	return s
}

func (s *HiMarketOidcConfigIdentityMapping) Validate() error {
	return dara.Validate(s)
}
