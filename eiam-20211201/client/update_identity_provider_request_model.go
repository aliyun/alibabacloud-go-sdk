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
	SetWeComConfig(v *UpdateIdentityProviderRequestWeComConfig) *UpdateIdentityProviderRequest
	GetWeComConfig() *UpdateIdentityProviderRequestWeComConfig
}

type UpdateIdentityProviderRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 钉钉出基本信息
	DingtalkAppConfig *UpdateIdentityProviderRequestDingtalkAppConfig `json:"DingtalkAppConfig,omitempty" xml:"DingtalkAppConfig,omitempty" type:"Struct"`
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// 身份提供方名称
	//
	// example:
	//
	// test
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 飞书配置
	LarkConfig *UpdateIdentityProviderRequestLarkConfig `json:"LarkConfig,omitempty" xml:"LarkConfig,omitempty" type:"Struct"`
	// AD/LDAP基本信息
	LdapConfig *UpdateIdentityProviderRequestLdapConfig `json:"LdapConfig,omitempty" xml:"LdapConfig,omitempty" type:"Struct"`
	LogoUrl    *string                                  `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// 网络端点ID
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// OIDC IdP配置。
	OidcConfig *UpdateIdentityProviderRequestOidcConfig `json:"OidcConfig,omitempty" xml:"OidcConfig,omitempty" type:"Struct"`
	// 企业微信基本信息
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

func (s *UpdateIdentityProviderRequest) SetWeComConfig(v *UpdateIdentityProviderRequestWeComConfig) *UpdateIdentityProviderRequest {
	s.WeComConfig = v
	return s
}

func (s *UpdateIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestDingtalkAppConfig struct {
	// 钉钉一方应用的AppKey
	//
	// example:
	//
	// 49nyeaqumk7f
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 钉钉一方应用的AppSecret
	//
	// example:
	//
	// 86nozWFL2CxgwnhKiXaG8dN4keLPkUNc5xxxx
	AppSecret         *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	EncryptKey        *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
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
	// example:
	//
	// cli_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// KiiLzh5Dueh4wbLxxxx
	AppSecret         *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	EncryptKey        *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
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
	// 管理员密码
	//
	// example:
	//
	// xxxxxx
	AdministratorPassword *string `json:"AdministratorPassword,omitempty" xml:"AdministratorPassword,omitempty"`
	// 管理员账号
	//
	// example:
	//
	// DC=example,DC=com
	AdministratorUsername *string `json:"AdministratorUsername,omitempty" xml:"AdministratorUsername,omitempty"`
	// 是否验证指纹证书
	//
	// example:
	//
	// enabled
	CertificateFingerprintStatus *string `json:"CertificateFingerprintStatus,omitempty" xml:"CertificateFingerprintStatus,omitempty"`
	// 证书指纹列表
	CertificateFingerprints []*string `json:"CertificateFingerprints,omitempty" xml:"CertificateFingerprints,omitempty" type:"Repeated"`
	// 通信协议
	//
	// example:
	//
	// ldap
	LdapProtocol *string `json:"LdapProtocol,omitempty" xml:"LdapProtocol,omitempty"`
	// ad/ldap 服务器地址
	//
	// example:
	//
	// 123.xx.xx.89
	LdapServerHost *string `json:"LdapServerHost,omitempty" xml:"LdapServerHost,omitempty"`
	// 端口号
	//
	// example:
	//
	// 636
	LdapServerPort *int32 `json:"LdapServerPort,omitempty" xml:"LdapServerPort,omitempty"`
	// startTls是否开启
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
	// OIDC客户端认证配置。
	AuthnParam *UpdateIdentityProviderRequestOidcConfigAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// OIDC 端点配置。
	EndpointConfig *UpdateIdentityProviderRequestOidcConfigEndpointConfig `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty" type:"Struct"`
	// OIDC标准参数，如profile、email等
	//
	// example:
	//
	// openid
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// OIDC授权类型。
	//
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// 支持的PKCE算法类型。
	//
	// example:
	//
	// S256
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// AuthorizationCode授权模式下是否使用PKCE。
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
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestOidcConfigAuthnParam struct {
	// OIDC/oAuth2 认证方法。
	//
	// example:
	//
	// client_secret_post
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// OIDC/oAuth2 客户端密钥。
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
	// oAuth2 授权端点。
	//
	// example:
	//
	// https://example.com/oauth/authorize
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// OIDC issuer信息。
	//
	// example:
	//
	// https://example.com/oauth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// OIDC jwks地址。
	//
	// example:
	//
	// https://example.com/oauth/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// oAuth2 Token端点。
	//
	// example:
	//
	// https://example.com/oauth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	// OIDC 用户信息端点。
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

type UpdateIdentityProviderRequestWeComConfig struct {
	// 企业微信自建应用的Id
	//
	// example:
	//
	// 1237403
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 授权回调域
	//
	// example:
	//
	// https://xxx.aliyunidaas.com/xxxxx
	AuthorizeCallbackDomain *string `json:"AuthorizeCallbackDomain,omitempty" xml:"AuthorizeCallbackDomain,omitempty"`
	// 企业微信自建应用的corpSecret
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	CorpSecret *string `json:"CorpSecret,omitempty" xml:"CorpSecret,omitempty"`
	// 可信域名
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
