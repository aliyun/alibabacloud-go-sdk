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
	// Configuration settings for the authorization code grant type.
	AuthCodeConfig *HiMarketOidcConfigAuthCodeConfig `json:"authCodeConfig,omitempty" xml:"authCodeConfig,omitempty" type:"Struct"`
	// Enables or disables this identity provider. If set to `false`, users cannot sign in with this provider.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The OAuth 2.0 grant type. For OIDC, this must be `authorization_code`.
	GrantType *string `json:"grantType,omitempty" xml:"grantType,omitempty"`
	// Specifies how to map claims from an ID token to user attributes in your system.
	IdentityMapping *HiMarketOidcConfigIdentityMapping `json:"identityMapping,omitempty" xml:"identityMapping,omitempty" type:"Struct"`
	// The URL for the provider\\"s logo. This logo appears on the sign-in page.
	LogoUrl *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	// The provider\\"s display name. This name appears on the sign-in page.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier for the identity provider.
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
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
	// The URL of the identity provider\\"s authorization endpoint.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" xml:"authorizationEndpoint,omitempty"`
	// The client ID obtained from the identity provider after registering your application.
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The client secret obtained from the identity provider after registering your application.
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	// The identity provider\\"s unique issuer URL, used to validate ID tokens.
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// The provider\\"s JWK Set URI. This URI provides the public keys needed to verify ID token signatures.
	JwkSetUri *string `json:"jwkSetUri,omitempty" xml:"jwkSetUri,omitempty"`
	// The application\\"s redirect URI. The provider sends the authorization code to this URI after successful authentication. You must register this URI with the identity provider.
	RedirectUri *string `json:"redirectUri,omitempty" xml:"redirectUri,omitempty"`
	// A space-separated list of scopes to request from the provider. The `openid` scope is required for OIDC authentication. For example: `openid profile email`.
	Scopes *string `json:"scopes,omitempty" xml:"scopes,omitempty"`
	// The URL of the identity provider\\"s token endpoint.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" xml:"tokenEndpoint,omitempty"`
	// The URL of the identity provider\\"s user info endpoint.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" xml:"userInfoEndpoint,omitempty"`
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
	// Maps additional claims from the ID token to custom user attributes. For each mapping, the key is the target attribute in your system, and the value is the name of the claim from the ID token.
	CustomFields map[string]*string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// The ID token claim that maps to the user\\"s email address. The `email` claim is a common choice.
	EmailField *string `json:"emailField,omitempty" xml:"emailField,omitempty"`
	// The ID token claim that maps to the user\\"s unique ID. The `sub` claim is a common choice.
	UserIdField *string `json:"userIdField,omitempty" xml:"userIdField,omitempty"`
	// The ID token claim that maps to the user\\"s display name. Common choices include `name` and `preferred_username`.
	UserNameField *string `json:"userNameField,omitempty" xml:"userNameField,omitempty"`
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
