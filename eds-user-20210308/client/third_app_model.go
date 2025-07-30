// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThirdApp interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *ThirdApp
	GetAppKey() *string
	SetName(v string) *ThirdApp
	GetName() *string
	SetOidcSsoConfig(v *ThirdAppOidcSsoConfig) *ThirdApp
	GetOidcSsoConfig() *ThirdAppOidcSsoConfig
	SetSecrets(v []*ThirdAppSecrets) *ThirdApp
	GetSecrets() []*ThirdAppSecrets
}

type ThirdApp struct {
	AppKey        *string                `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Name          *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OidcSsoConfig *ThirdAppOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	Secrets       []*ThirdAppSecrets     `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s ThirdApp) String() string {
	return dara.Prettify(s)
}

func (s ThirdApp) GoString() string {
	return s.String()
}

func (s *ThirdApp) GetAppKey() *string {
	return s.AppKey
}

func (s *ThirdApp) GetName() *string {
	return s.Name
}

func (s *ThirdApp) GetOidcSsoConfig() *ThirdAppOidcSsoConfig {
	return s.OidcSsoConfig
}

func (s *ThirdApp) GetSecrets() []*ThirdAppSecrets {
	return s.Secrets
}

func (s *ThirdApp) SetAppKey(v string) *ThirdApp {
	s.AppKey = &v
	return s
}

func (s *ThirdApp) SetName(v string) *ThirdApp {
	s.Name = &v
	return s
}

func (s *ThirdApp) SetOidcSsoConfig(v *ThirdAppOidcSsoConfig) *ThirdApp {
	s.OidcSsoConfig = v
	return s
}

func (s *ThirdApp) SetSecrets(v []*ThirdAppSecrets) *ThirdApp {
	s.Secrets = v
	return s
}

func (s *ThirdApp) Validate() error {
	return dara.Validate(s)
}

type ThirdAppOidcSsoConfig struct {
	AccessTokenEffectiveTime *int32                          `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	CodeEffectiveTime        *int32                          `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	EnableAuthLogin          *bool                           `json:"EnableAuthLogin,omitempty" xml:"EnableAuthLogin,omitempty"`
	Endpoints                *ThirdAppOidcSsoConfigEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	GrantScopes              []*string                       `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	GrantTypes               []*string                       `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	IdTokenAlgorithmType     *string                         `json:"IdTokenAlgorithmType,omitempty" xml:"IdTokenAlgorithmType,omitempty"`
	IdTokenEffectiveTime     *int32                          `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	RedirectUris             []*string                       `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	RefreshTokenEffective    *int32                          `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
}

func (s ThirdAppOidcSsoConfig) String() string {
	return dara.Prettify(s)
}

func (s ThirdAppOidcSsoConfig) GoString() string {
	return s.String()
}

func (s *ThirdAppOidcSsoConfig) GetAccessTokenEffectiveTime() *int32 {
	return s.AccessTokenEffectiveTime
}

func (s *ThirdAppOidcSsoConfig) GetCodeEffectiveTime() *int32 {
	return s.CodeEffectiveTime
}

func (s *ThirdAppOidcSsoConfig) GetEnableAuthLogin() *bool {
	return s.EnableAuthLogin
}

func (s *ThirdAppOidcSsoConfig) GetEndpoints() *ThirdAppOidcSsoConfigEndpoints {
	return s.Endpoints
}

func (s *ThirdAppOidcSsoConfig) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *ThirdAppOidcSsoConfig) GetGrantTypes() []*string {
	return s.GrantTypes
}

func (s *ThirdAppOidcSsoConfig) GetIdTokenAlgorithmType() *string {
	return s.IdTokenAlgorithmType
}

func (s *ThirdAppOidcSsoConfig) GetIdTokenEffectiveTime() *int32 {
	return s.IdTokenEffectiveTime
}

func (s *ThirdAppOidcSsoConfig) GetRedirectUris() []*string {
	return s.RedirectUris
}

func (s *ThirdAppOidcSsoConfig) GetRefreshTokenEffective() *int32 {
	return s.RefreshTokenEffective
}

func (s *ThirdAppOidcSsoConfig) SetAccessTokenEffectiveTime(v int32) *ThirdAppOidcSsoConfig {
	s.AccessTokenEffectiveTime = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetCodeEffectiveTime(v int32) *ThirdAppOidcSsoConfig {
	s.CodeEffectiveTime = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetEnableAuthLogin(v bool) *ThirdAppOidcSsoConfig {
	s.EnableAuthLogin = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetEndpoints(v *ThirdAppOidcSsoConfigEndpoints) *ThirdAppOidcSsoConfig {
	s.Endpoints = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetGrantScopes(v []*string) *ThirdAppOidcSsoConfig {
	s.GrantScopes = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetGrantTypes(v []*string) *ThirdAppOidcSsoConfig {
	s.GrantTypes = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetIdTokenAlgorithmType(v string) *ThirdAppOidcSsoConfig {
	s.IdTokenAlgorithmType = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetIdTokenEffectiveTime(v int32) *ThirdAppOidcSsoConfig {
	s.IdTokenEffectiveTime = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetRedirectUris(v []*string) *ThirdAppOidcSsoConfig {
	s.RedirectUris = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetRefreshTokenEffective(v int32) *ThirdAppOidcSsoConfig {
	s.RefreshTokenEffective = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) Validate() error {
	return dara.Validate(s)
}

type ThirdAppOidcSsoConfigEndpoints struct {
	AuthorizationEndpoint      *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	DiscoveryEndpoint          *string `json:"DiscoveryEndpoint,omitempty" xml:"DiscoveryEndpoint,omitempty"`
	GuestAuthorizationEndpoint *string `json:"GuestAuthorizationEndpoint,omitempty" xml:"GuestAuthorizationEndpoint,omitempty"`
	Issuer                     *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	JwksEndpoint               *string `json:"JwksEndpoint,omitempty" xml:"JwksEndpoint,omitempty"`
	LogoutEndpoint             *string `json:"LogoutEndpoint,omitempty" xml:"LogoutEndpoint,omitempty"`
	RevokeEndpoint             *string `json:"RevokeEndpoint,omitempty" xml:"RevokeEndpoint,omitempty"`
	TokenEndpoint              *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	UserinfoEndpoint           *string `json:"UserinfoEndpoint,omitempty" xml:"UserinfoEndpoint,omitempty"`
}

func (s ThirdAppOidcSsoConfigEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ThirdAppOidcSsoConfigEndpoints) GoString() string {
	return s.String()
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetDiscoveryEndpoint() *string {
	return s.DiscoveryEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetGuestAuthorizationEndpoint() *string {
	return s.GuestAuthorizationEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetIssuer() *string {
	return s.Issuer
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetJwksEndpoint() *string {
	return s.JwksEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetLogoutEndpoint() *string {
	return s.LogoutEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetRevokeEndpoint() *string {
	return s.RevokeEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetAuthorizationEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetDiscoveryEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetGuestAuthorizationEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.GuestAuthorizationEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetIssuer(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.Issuer = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetJwksEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.JwksEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetLogoutEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.LogoutEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetRevokeEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.RevokeEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetTokenEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.TokenEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetUserinfoEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.UserinfoEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) Validate() error {
	return dara.Validate(s)
}

type ThirdAppSecrets struct {
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
}

func (s ThirdAppSecrets) String() string {
	return dara.Prettify(s)
}

func (s ThirdAppSecrets) GoString() string {
	return s.String()
}

func (s *ThirdAppSecrets) GetEnable() *bool {
	return s.Enable
}

func (s *ThirdAppSecrets) GetSecret() *string {
	return s.Secret
}

func (s *ThirdAppSecrets) SetEnable(v bool) *ThirdAppSecrets {
	s.Enable = &v
	return s
}

func (s *ThirdAppSecrets) SetSecret(v string) *ThirdAppSecrets {
	s.Secret = &v
	return s
}

func (s *ThirdAppSecrets) Validate() error {
	return dara.Validate(s)
}
