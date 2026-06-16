// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateIdentityProviderRequest
	GetClientToken() *string
	SetDingtalkAppConfig(v *UpdateIdentityProviderRequestDingtalkAppConfig) *UpdateIdentityProviderRequest
	GetDingtalkAppConfig() *UpdateIdentityProviderRequestDingtalkAppConfig
	SetIdentityProviderId(v string) *UpdateIdentityProviderRequest
	GetIdentityProviderId() *string
	SetIdentityProviderName(v string) *UpdateIdentityProviderRequest
	GetIdentityProviderName() *string
	SetInstanceId(v string) *UpdateIdentityProviderRequest
	GetInstanceId() *string
	SetLarkConfig(v *UpdateIdentityProviderRequestLarkConfig) *UpdateIdentityProviderRequest
	GetLarkConfig() *UpdateIdentityProviderRequestLarkConfig
	SetLdapConfig(v *UpdateIdentityProviderRequestLdapConfig) *UpdateIdentityProviderRequest
	GetLdapConfig() *UpdateIdentityProviderRequestLdapConfig
	SetLogoUrl(v string) *UpdateIdentityProviderRequest
	GetLogoUrl() *string
	SetNetworkAccessEndpointId(v string) *UpdateIdentityProviderRequest
	GetNetworkAccessEndpointId() *string
	SetOidcConfig(v *UpdateIdentityProviderRequestOidcConfig) *UpdateIdentityProviderRequest
	GetOidcConfig() *UpdateIdentityProviderRequestOidcConfig
	SetSamlConfig(v *UpdateIdentityProviderRequestSamlConfig) *UpdateIdentityProviderRequest
	GetSamlConfig() *UpdateIdentityProviderRequestSamlConfig
	SetWeComConfig(v *UpdateIdentityProviderRequestWeComConfig) *UpdateIdentityProviderRequest
	GetWeComConfig() *UpdateIdentityProviderRequestWeComConfig
}

type UpdateIdentityProviderRequest struct {
	// A client-generated token to ensure request idempotence. This value must be unique across requests.
	//
	// example:
	//
	// client-examplexxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration for the DingTalk identity provider.
	DingtalkAppConfig *UpdateIdentityProviderRequestDingtalkAppConfig `json:"DingtalkAppConfig,omitempty" xml:"DingtalkAppConfig,omitempty" type:"Struct"`
	// The ID of the identity provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The name of the identity provider.
	//
	// example:
	//
	// test
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration for the Lark identity provider.
	LarkConfig *UpdateIdentityProviderRequestLarkConfig `json:"LarkConfig,omitempty" xml:"LarkConfig,omitempty" type:"Struct"`
	// The configuration for the Active Directory (AD) or Lightweight Directory Access Protocol (LDAP) identity provider.
	LdapConfig *UpdateIdentityProviderRequestLdapConfig `json:"LdapConfig,omitempty" xml:"LdapConfig,omitempty" type:"Struct"`
	// The URL of the application logo.
	//
	// example:
	//
	// idaas-image://idaas_23aqr2ye554csg33dqpch5exxxx/tmp/d17d9adc-a943-45e7-ba0c-2838dddea678xxxx
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The ID of the network access endpoint.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The OpenID Connect (OIDC) configuration.
	OidcConfig *UpdateIdentityProviderRequestOidcConfig `json:"OidcConfig,omitempty" xml:"OidcConfig,omitempty" type:"Struct"`
	// The configuration for the SAML identity provider.
	SamlConfig *UpdateIdentityProviderRequestSamlConfig `json:"SamlConfig,omitempty" xml:"SamlConfig,omitempty" type:"Struct"`
	// The configuration for the WeCom identity provider.
	WeComConfig *UpdateIdentityProviderRequestWeComConfig `json:"WeComConfig,omitempty" xml:"WeComConfig,omitempty" type:"Struct"`
}

func (s UpdateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIdentityProviderRequest) GetDingtalkAppConfig() *UpdateIdentityProviderRequestDingtalkAppConfig {
	return s.DingtalkAppConfig
}

func (s *UpdateIdentityProviderRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *UpdateIdentityProviderRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *UpdateIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateIdentityProviderRequest) GetLarkConfig() *UpdateIdentityProviderRequestLarkConfig {
	return s.LarkConfig
}

func (s *UpdateIdentityProviderRequest) GetLdapConfig() *UpdateIdentityProviderRequestLdapConfig {
	return s.LdapConfig
}

func (s *UpdateIdentityProviderRequest) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *UpdateIdentityProviderRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *UpdateIdentityProviderRequest) GetOidcConfig() *UpdateIdentityProviderRequestOidcConfig {
	return s.OidcConfig
}

func (s *UpdateIdentityProviderRequest) GetSamlConfig() *UpdateIdentityProviderRequestSamlConfig {
	return s.SamlConfig
}

func (s *UpdateIdentityProviderRequest) GetWeComConfig() *UpdateIdentityProviderRequestWeComConfig {
	return s.WeComConfig
}

func (s *UpdateIdentityProviderRequest) SetClientToken(v string) *UpdateIdentityProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetDingtalkAppConfig(v *UpdateIdentityProviderRequestDingtalkAppConfig) *UpdateIdentityProviderRequest {
	s.DingtalkAppConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) SetIdentityProviderId(v string) *UpdateIdentityProviderRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetIdentityProviderName(v string) *UpdateIdentityProviderRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetInstanceId(v string) *UpdateIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetLarkConfig(v *UpdateIdentityProviderRequestLarkConfig) *UpdateIdentityProviderRequest {
	s.LarkConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) SetLdapConfig(v *UpdateIdentityProviderRequestLdapConfig) *UpdateIdentityProviderRequest {
	s.LdapConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) SetLogoUrl(v string) *UpdateIdentityProviderRequest {
	s.LogoUrl = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetNetworkAccessEndpointId(v string) *UpdateIdentityProviderRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetOidcConfig(v *UpdateIdentityProviderRequestOidcConfig) *UpdateIdentityProviderRequest {
	s.OidcConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) SetSamlConfig(v *UpdateIdentityProviderRequestSamlConfig) *UpdateIdentityProviderRequest {
	s.SamlConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) SetWeComConfig(v *UpdateIdentityProviderRequestWeComConfig) *UpdateIdentityProviderRequest {
	s.WeComConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) Validate() error {
	if s.DingtalkAppConfig != nil {
		if err := s.DingtalkAppConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LarkConfig != nil {
		if err := s.LarkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LdapConfig != nil {
		if err := s.LdapConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OidcConfig != nil {
		if err := s.OidcConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SamlConfig != nil {
		if err := s.SamlConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WeComConfig != nil {
		if err := s.WeComConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderRequestDingtalkAppConfig struct {
	// The AppKey of the DingTalk application.
	//
	// example:
	//
	// 49nyeaqumk7f
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The AppSecret of the DingTalk application.
	//
	// example:
	//
	// 86nozWFL2CxgwnhKiXaG8dN4keLPkUNc5xxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The version of the DingTalk QR code login.
	//
	// example:
	//
	// new_version
	DingtalkLoginVersion *string `json:"DingtalkLoginVersion,omitempty" xml:"DingtalkLoginVersion,omitempty"`
	// The EncryptKey of the DingTalk application.
	//
	// example:
	//
	// VkdWw91mdkrjVFr3ObNwefap21dfxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// The verification token of the DingTalk application.
	//
	// example:
	//
	// myDingApp_VerifyTokenxxxxx
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s UpdateIdentityProviderRequestDingtalkAppConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestDingtalkAppConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) GetAppKey() *string {
	return s.AppKey
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) GetDingtalkLoginVersion() *string {
	return s.DingtalkLoginVersion
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) SetAppKey(v string) *UpdateIdentityProviderRequestDingtalkAppConfig {
	s.AppKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) SetAppSecret(v string) *UpdateIdentityProviderRequestDingtalkAppConfig {
	s.AppSecret = &v
	return s
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) SetDingtalkLoginVersion(v string) *UpdateIdentityProviderRequestDingtalkAppConfig {
	s.DingtalkLoginVersion = &v
	return s
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) SetEncryptKey(v string) *UpdateIdentityProviderRequestDingtalkAppConfig {
	s.EncryptKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) SetVerificationToken(v string) *UpdateIdentityProviderRequestDingtalkAppConfig {
	s.VerificationToken = &v
	return s
}

func (s *UpdateIdentityProviderRequestDingtalkAppConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestLarkConfig struct {
	// The application ID of the custom application in Lark.
	//
	// example:
	//
	// cli_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application secret of the custom application in Lark.
	//
	// example:
	//
	// KiiLzh5Dueh4wbLxxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The EncryptKey of the custom application in Lark.
	//
	// example:
	//
	// VkdWw91mdkrjVFr3ObNwefap21dfbZbKxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// The verification token of the custom application in Lark.
	//
	// example:
	//
	// feishuVerifyTokenxxxxx
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s UpdateIdentityProviderRequestLarkConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestLarkConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestLarkConfig) GetAppId() *string {
	return s.AppId
}

func (s *UpdateIdentityProviderRequestLarkConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *UpdateIdentityProviderRequestLarkConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *UpdateIdentityProviderRequestLarkConfig) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *UpdateIdentityProviderRequestLarkConfig) SetAppId(v string) *UpdateIdentityProviderRequestLarkConfig {
	s.AppId = &v
	return s
}

func (s *UpdateIdentityProviderRequestLarkConfig) SetAppSecret(v string) *UpdateIdentityProviderRequestLarkConfig {
	s.AppSecret = &v
	return s
}

func (s *UpdateIdentityProviderRequestLarkConfig) SetEncryptKey(v string) *UpdateIdentityProviderRequestLarkConfig {
	s.EncryptKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestLarkConfig) SetVerificationToken(v string) *UpdateIdentityProviderRequestLarkConfig {
	s.VerificationToken = &v
	return s
}

func (s *UpdateIdentityProviderRequestLarkConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestLdapConfig struct {
	// The password for the administrator account.
	//
	// example:
	//
	// xxxxxx
	AdministratorPassword *string `json:"AdministratorPassword,omitempty" xml:"AdministratorPassword,omitempty"`
	// The administrator account.
	//
	// example:
	//
	// DC=example,DC=com
	AdministratorUsername *string `json:"AdministratorUsername,omitempty" xml:"AdministratorUsername,omitempty"`
	// Specifies whether to enable certificate fingerprint verification. Valid values:
	//
	// - `disabled`: Verification is disabled.
	//
	// - `enabled`: Verification is enabled.
	//
	// example:
	//
	// enabled
	CertificateFingerprintStatus *string `json:"CertificateFingerprintStatus,omitempty" xml:"CertificateFingerprintStatus,omitempty"`
	// The list of certificate fingerprints.
	CertificateFingerprints []*string `json:"CertificateFingerprints,omitempty" xml:"CertificateFingerprints,omitempty" type:"Repeated"`
	// The communication protocol.
	//
	// example:
	//
	// ldap
	LdapProtocol *string `json:"LdapProtocol,omitempty" xml:"LdapProtocol,omitempty"`
	// The server address.
	//
	// example:
	//
	// 123.xx.xx.89
	LdapServerHost *string `json:"LdapServerHost,omitempty" xml:"LdapServerHost,omitempty"`
	// The port number.
	//
	// example:
	//
	// 636
	LdapServerPort *int32 `json:"LdapServerPort,omitempty" xml:"LdapServerPort,omitempty"`
	// Specifies whether to enable StartTLS. Valid values:
	//
	// - `disabled`: StartTLS is disabled.
	//
	// - `enabled`: StartTLS is enabled.
	//
	// example:
	//
	// enabled
	StartTlsStatus *string `json:"StartTlsStatus,omitempty" xml:"StartTlsStatus,omitempty"`
}

func (s UpdateIdentityProviderRequestLdapConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestLdapConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetAdministratorPassword() *string {
	return s.AdministratorPassword
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetAdministratorUsername() *string {
	return s.AdministratorUsername
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetCertificateFingerprintStatus() *string {
	return s.CertificateFingerprintStatus
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetCertificateFingerprints() []*string {
	return s.CertificateFingerprints
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetLdapProtocol() *string {
	return s.LdapProtocol
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetLdapServerHost() *string {
	return s.LdapServerHost
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetLdapServerPort() *int32 {
	return s.LdapServerPort
}

func (s *UpdateIdentityProviderRequestLdapConfig) GetStartTlsStatus() *string {
	return s.StartTlsStatus
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetAdministratorPassword(v string) *UpdateIdentityProviderRequestLdapConfig {
	s.AdministratorPassword = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetAdministratorUsername(v string) *UpdateIdentityProviderRequestLdapConfig {
	s.AdministratorUsername = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetCertificateFingerprintStatus(v string) *UpdateIdentityProviderRequestLdapConfig {
	s.CertificateFingerprintStatus = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetCertificateFingerprints(v []*string) *UpdateIdentityProviderRequestLdapConfig {
	s.CertificateFingerprints = v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetLdapProtocol(v string) *UpdateIdentityProviderRequestLdapConfig {
	s.LdapProtocol = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetLdapServerHost(v string) *UpdateIdentityProviderRequestLdapConfig {
	s.LdapServerHost = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetLdapServerPort(v int32) *UpdateIdentityProviderRequestLdapConfig {
	s.LdapServerPort = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) SetStartTlsStatus(v string) *UpdateIdentityProviderRequestLdapConfig {
	s.StartTlsStatus = &v
	return s
}

func (s *UpdateIdentityProviderRequestLdapConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestOidcConfig struct {
	// The OIDC client authentication configuration.
	AuthnParam *UpdateIdentityProviderRequestOidcConfigAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// The OIDC endpoint configuration.
	EndpointConfig *UpdateIdentityProviderRequestOidcConfigEndpointConfig `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty" type:"Struct"`
	// The OIDC authorization scopes.
	//
	// example:
	//
	// openid
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The OIDC grant type.
	//
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The Proof Key for Code Exchange (PKCE) method. Valid values:
	//
	// - `S256`: The SHA-256 algorithm.
	//
	// - `plain`: The plaintext format.
	//
	// example:
	//
	// S256
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Specifies whether PKCE is required for the authorization code grant type.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
}

func (s UpdateIdentityProviderRequestOidcConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestOidcConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestOidcConfig) GetAuthnParam() *UpdateIdentityProviderRequestOidcConfigAuthnParam {
	return s.AuthnParam
}

func (s *UpdateIdentityProviderRequestOidcConfig) GetEndpointConfig() *UpdateIdentityProviderRequestOidcConfigEndpointConfig {
	return s.EndpointConfig
}

func (s *UpdateIdentityProviderRequestOidcConfig) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *UpdateIdentityProviderRequestOidcConfig) GetGrantType() *string {
	return s.GrantType
}

func (s *UpdateIdentityProviderRequestOidcConfig) GetPkceChallengeMethod() *string {
	return s.PkceChallengeMethod
}

func (s *UpdateIdentityProviderRequestOidcConfig) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *UpdateIdentityProviderRequestOidcConfig) SetAuthnParam(v *UpdateIdentityProviderRequestOidcConfigAuthnParam) *UpdateIdentityProviderRequestOidcConfig {
	s.AuthnParam = v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfig) SetEndpointConfig(v *UpdateIdentityProviderRequestOidcConfigEndpointConfig) *UpdateIdentityProviderRequestOidcConfig {
	s.EndpointConfig = v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfig) SetGrantScopes(v []*string) *UpdateIdentityProviderRequestOidcConfig {
	s.GrantScopes = v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfig) SetGrantType(v string) *UpdateIdentityProviderRequestOidcConfig {
	s.GrantType = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfig) SetPkceChallengeMethod(v string) *UpdateIdentityProviderRequestOidcConfig {
	s.PkceChallengeMethod = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfig) SetPkceRequired(v bool) *UpdateIdentityProviderRequestOidcConfig {
	s.PkceRequired = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfig) Validate() error {
	if s.AuthnParam != nil {
		if err := s.AuthnParam.Validate(); err != nil {
			return err
		}
	}
	if s.EndpointConfig != nil {
		if err := s.EndpointConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderRequestOidcConfigAuthnParam struct {
	// The OIDC client authentication method. Valid values:
	//
	// - `client_secret_basic`
	//
	// - `client_secret_post`
	//
	// example:
	//
	// client_secret_post
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The OIDC client secret.
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s UpdateIdentityProviderRequestOidcConfigAuthnParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestOidcConfigAuthnParam) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestOidcConfigAuthnParam) GetAuthnMethod() *string {
	return s.AuthnMethod
}

func (s *UpdateIdentityProviderRequestOidcConfigAuthnParam) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateIdentityProviderRequestOidcConfigAuthnParam) SetAuthnMethod(v string) *UpdateIdentityProviderRequestOidcConfigAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigAuthnParam) SetClientSecret(v string) *UpdateIdentityProviderRequestOidcConfigAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigAuthnParam) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestOidcConfigEndpointConfig struct {
	// The OIDC authorization endpoint.
	//
	// example:
	//
	// https://example.com/oauth/authorize
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The OIDC issuer.
	//
	// example:
	//
	// https://example.com/oauth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The JSON Web Key Set (JWKS) URI.
	//
	// example:
	//
	// https://example.com/oauth/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// The OIDC token endpoint.
	//
	// example:
	//
	// https://example.com/oauth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	// The OIDC userinfo endpoint.
	//
	// example:
	//
	// https://example.com/oauth/userinfo
	UserinfoEndpoint *string `json:"UserinfoEndpoint,omitempty" xml:"UserinfoEndpoint,omitempty"`
}

func (s UpdateIdentityProviderRequestOidcConfigEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestOidcConfigEndpointConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) SetAuthorizationEndpoint(v string) *UpdateIdentityProviderRequestOidcConfigEndpointConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) SetIssuer(v string) *UpdateIdentityProviderRequestOidcConfigEndpointConfig {
	s.Issuer = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) SetJwksUri(v string) *UpdateIdentityProviderRequestOidcConfigEndpointConfig {
	s.JwksUri = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) SetTokenEndpoint(v string) *UpdateIdentityProviderRequestOidcConfigEndpointConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) SetUserinfoEndpoint(v string) *UpdateIdentityProviderRequestOidcConfigEndpointConfig {
	s.UserinfoEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderRequestOidcConfigEndpointConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestSamlConfig struct {
	// The SAML binding method for the SSO request. Valid values are `HTTP-POST` and `HTTP-REDIRECT`.
	//
	// example:
	//
	// HTTP-REDIRECT
	BindingMethod *string `json:"BindingMethod,omitempty" xml:"BindingMethod,omitempty"`
	// The signing certificates from the SAML identity provider.
	Certificates []*UpdateIdentityProviderRequestSamlConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The entity ID of the SAML identity provider.
	//
	// example:
	//
	// http://dc.test.com/adfs/services/trust
	IdPEntityId *string `json:"IdPEntityId,omitempty" xml:"IdPEntityId,omitempty"`
	// The single sign-on (SSO) URL of the SAML identity provider.
	//
	// example:
	//
	// https://dc.test.com/adfs/ls/
	IdPSsoUrl *string `json:"IdPSsoUrl,omitempty" xml:"IdPSsoUrl,omitempty"`
	// The maximum allowed clock skew, in seconds.
	//
	// example:
	//
	// 180
	MaxClockSkew *int64 `json:"MaxClockSkew,omitempty" xml:"MaxClockSkew,omitempty"`
	// Specifies whether the SAML authentication request must be signed.
	//
	// example:
	//
	// true
	RequireRequestSigned *bool `json:"RequireRequestSigned,omitempty" xml:"RequireRequestSigned,omitempty"`
	// Specifies whether the assertions in the SAML response must be signed.
	WantAssertionsSigned *bool `json:"WantAssertionsSigned,omitempty" xml:"WantAssertionsSigned,omitempty"`
	// Specifies whether the SAML response must be signed.
	WantResponseSigned *bool `json:"WantResponseSigned,omitempty" xml:"WantResponseSigned,omitempty"`
}

func (s UpdateIdentityProviderRequestSamlConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestSamlConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetBindingMethod() *string {
	return s.BindingMethod
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetCertificates() []*UpdateIdentityProviderRequestSamlConfigCertificates {
	return s.Certificates
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetIdPEntityId() *string {
	return s.IdPEntityId
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetIdPSsoUrl() *string {
	return s.IdPSsoUrl
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetMaxClockSkew() *int64 {
	return s.MaxClockSkew
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetRequireRequestSigned() *bool {
	return s.RequireRequestSigned
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetWantAssertionsSigned() *bool {
	return s.WantAssertionsSigned
}

func (s *UpdateIdentityProviderRequestSamlConfig) GetWantResponseSigned() *bool {
	return s.WantResponseSigned
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetBindingMethod(v string) *UpdateIdentityProviderRequestSamlConfig {
	s.BindingMethod = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetCertificates(v []*UpdateIdentityProviderRequestSamlConfigCertificates) *UpdateIdentityProviderRequestSamlConfig {
	s.Certificates = v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetIdPEntityId(v string) *UpdateIdentityProviderRequestSamlConfig {
	s.IdPEntityId = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetIdPSsoUrl(v string) *UpdateIdentityProviderRequestSamlConfig {
	s.IdPSsoUrl = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetMaxClockSkew(v int64) *UpdateIdentityProviderRequestSamlConfig {
	s.MaxClockSkew = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetRequireRequestSigned(v bool) *UpdateIdentityProviderRequestSamlConfig {
	s.RequireRequestSigned = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetWantAssertionsSigned(v bool) *UpdateIdentityProviderRequestSamlConfig {
	s.WantAssertionsSigned = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) SetWantResponseSigned(v bool) *UpdateIdentityProviderRequestSamlConfig {
	s.WantResponseSigned = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfig) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateIdentityProviderRequestSamlConfigCertificates struct {
	// The content of the signing certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIC0jCCAbqgAwIBAgIQXXXXX-----END CERTIFICATE-----
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateIdentityProviderRequestSamlConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestSamlConfigCertificates) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestSamlConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *UpdateIdentityProviderRequestSamlConfigCertificates) SetContent(v string) *UpdateIdentityProviderRequestSamlConfigCertificates {
	s.Content = &v
	return s
}

func (s *UpdateIdentityProviderRequestSamlConfigCertificates) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestWeComConfig struct {
	// The agent ID of the custom application in WeCom.
	//
	// example:
	//
	// 1237403
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The authorized callback domain.
	//
	// example:
	//
	// https://xxx.aliyunidaas.com/xxxxx
	AuthorizeCallbackDomain *string `json:"AuthorizeCallbackDomain,omitempty" xml:"AuthorizeCallbackDomain,omitempty"`
	// The CorpSecret of the custom application in WeCom.
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	CorpSecret *string `json:"CorpSecret,omitempty" xml:"CorpSecret,omitempty"`
	// The trusted domain.
	//
	// example:
	//
	// https://xxx.aliyunidaas.com
	TrustableDomain *string `json:"TrustableDomain,omitempty" xml:"TrustableDomain,omitempty"`
}

func (s UpdateIdentityProviderRequestWeComConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestWeComConfig) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestWeComConfig) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateIdentityProviderRequestWeComConfig) GetAuthorizeCallbackDomain() *string {
	return s.AuthorizeCallbackDomain
}

func (s *UpdateIdentityProviderRequestWeComConfig) GetCorpSecret() *string {
	return s.CorpSecret
}

func (s *UpdateIdentityProviderRequestWeComConfig) GetTrustableDomain() *string {
	return s.TrustableDomain
}

func (s *UpdateIdentityProviderRequestWeComConfig) SetAgentId(v string) *UpdateIdentityProviderRequestWeComConfig {
	s.AgentId = &v
	return s
}

func (s *UpdateIdentityProviderRequestWeComConfig) SetAuthorizeCallbackDomain(v string) *UpdateIdentityProviderRequestWeComConfig {
	s.AuthorizeCallbackDomain = &v
	return s
}

func (s *UpdateIdentityProviderRequestWeComConfig) SetCorpSecret(v string) *UpdateIdentityProviderRequestWeComConfig {
	s.CorpSecret = &v
	return s
}

func (s *UpdateIdentityProviderRequestWeComConfig) SetTrustableDomain(v string) *UpdateIdentityProviderRequestWeComConfig {
	s.TrustableDomain = &v
	return s
}

func (s *UpdateIdentityProviderRequestWeComConfig) Validate() error {
	return dara.Validate(s)
}
