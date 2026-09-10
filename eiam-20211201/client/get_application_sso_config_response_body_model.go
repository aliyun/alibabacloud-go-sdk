// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationSsoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) *GetApplicationSsoConfigResponseBody
	GetApplicationSsoConfig() *GetApplicationSsoConfigResponseBodyApplicationSsoConfig
	SetRequestId(v string) *GetApplicationSsoConfigResponseBody
	GetRequestId() *string
}

type GetApplicationSsoConfigResponseBody struct {
	// The single sign-on (SSO) configuration information of the application.
	ApplicationSsoConfig *GetApplicationSsoConfigResponseBodyApplicationSsoConfig `json:"ApplicationSsoConfig,omitempty" xml:"ApplicationSsoConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationSsoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBody) GetApplicationSsoConfig() *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	return s.ApplicationSsoConfig
}

func (s *GetApplicationSsoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationSsoConfigResponseBody) SetApplicationSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) *GetApplicationSsoConfigResponseBody {
	s.ApplicationSsoConfig = v
	return s
}

func (s *GetApplicationSsoConfigResponseBody) SetRequestId(v string) *GetApplicationSsoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBody) Validate() error {
	if s.ApplicationSsoConfig != nil {
		if err := s.ApplicationSsoConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfig struct {
	// The initialization single sign-on (SSO) method. Valid values:
	//
	// - only_app_init_sso: Only application-initiated SSO. This is the default value for OIDC protocol applications. When a SAML application specifies this method, InitLoginUrl must be specified.
	//
	// - idaas_or_app_init_sso: IDaaS portal-initiated or application-initiated SSO. This is the default value for SAML protocol applications. When an OIDC application specifies this method, InitLoginUrl must be specified.
	//
	// example:
	//
	// only_app_init_sso
	InitLoginType *string `json:"InitLoginType,omitempty" xml:"InitLoginType,omitempty"`
	// The initialization single sign-on (SSO) trigger URL. This parameter is required when the InitLoginType of an OIDC protocol application is set to idaas_or_app_init_sso, or when the InitLoginType of a SAML protocol application is set to only_app_init_sso.
	//
	// example:
	//
	// http://127.0.0.1:8000/start_login?enterprise_code=ABCDEF
	InitLoginUrl *string `json:"InitLoginUrl,omitempty" xml:"InitLoginUrl,omitempty"`
	// The SSO configuration parameters for OIDC protocol applications. This parameter is returned only when the application SSO protocol is OIDC.
	OidcSsoConfig *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	// The metadata endpoint configuration provided by the application.
	ProtocolEndpointDomain *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain `json:"ProtocolEndpointDomain,omitempty" xml:"ProtocolEndpointDomain,omitempty" type:"Struct"`
	// The SSO configuration parameters for SAML protocol applications. This parameter is returned only when the application SSO protocol is SAML 2.0.
	SamlSsoConfig *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig `json:"SamlSsoConfig,omitempty" xml:"SamlSsoConfig,omitempty" type:"Struct"`
	// The SSO status of the application. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	SsoStatus *string `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GetInitLoginType() *string {
	return s.InitLoginType
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GetInitLoginUrl() *string {
	return s.InitLoginUrl
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GetOidcSsoConfig() *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	return s.OidcSsoConfig
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GetProtocolEndpointDomain() *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	return s.ProtocolEndpointDomain
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GetSamlSsoConfig() *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	return s.SamlSsoConfig
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetInitLoginType(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.InitLoginType = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetInitLoginUrl(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.InitLoginUrl = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetOidcSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.OidcSsoConfig = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetProtocolEndpointDomain(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.ProtocolEndpointDomain = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetSamlSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.SamlSsoConfig = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetSsoStatus(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.SsoStatus = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) Validate() error {
	if s.OidcSsoConfig != nil {
		if err := s.OidcSsoConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ProtocolEndpointDomain != nil {
		if err := s.ProtocolEndpointDomain.Validate(); err != nil {
			return err
		}
	}
	if s.SamlSsoConfig != nil {
		if err := s.SamlSsoConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig struct {
	// The validity period of the issued access token. Unit: seconds. Default value: 1200 (20 minutes).
	//
	// example:
	//
	// 1200
	AccessTokenEffectiveTime *int64 `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	// Specifies whether the application is allowed to request the IDaaS EIAM authorization server as a public client. This parameter can be enabled only in authorization code mode and device mode. Default value: false.
	//
	// example:
	//
	// true
	AllowedPublicClient *string `json:"AllowedPublicClient,omitempty" xml:"AllowedPublicClient,omitempty"`
	// The validity period of the issued code. Unit: seconds. Default value: 60 (1 minute).
	//
	// example:
	//
	// 60
	CodeEffectiveTime *int64 `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	// The custom user information included in the ID token response.
	CustomClaims []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims `json:"CustomClaims,omitempty" xml:"CustomClaims,omitempty" type:"Repeated"`
	// The OIDC standard parameter scope, which specifies the range of user attributes that can be returned by the userinfo endpoint or ID token.
	//
	// example:
	//
	// profile，email
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The list of supported OIDC protocol grant types.
	//
	// example:
	//
	// authorization_code
	GrantTypes []*string `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	// The validity period of the issued ID token. Unit: seconds. Default value: 300 (5 minutes).
	//
	// example:
	//
	// 1200
	IdTokenEffectiveTime *int64 `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	// The ID of the identity authentication source used in password mode. This parameter takes effect only when the GrantTypes specified for the OIDC protocol application include the password mode.
	//
	// example:
	//
	// ia_password
	PasswordAuthenticationSourceId *string `json:"PasswordAuthenticationSourceId,omitempty" xml:"PasswordAuthenticationSourceId,omitempty"`
	// Specifies whether TOTP-based secondary authentication is required in password mode. This parameter takes effect only when the GrantTypes specified for the OIDC protocol application include the password mode.
	//
	// example:
	//
	// true
	PasswordTotpMfaRequired *bool `json:"PasswordTotpMfaRequired,omitempty" xml:"PasswordTotpMfaRequired,omitempty"`
	// The algorithm used to calculate the Code Challenge in PKCE.
	//
	// example:
	//
	// S256
	PkceChallengeMethods []*string `json:"PkceChallengeMethods,omitempty" xml:"PkceChallengeMethods,omitempty" type:"Repeated"`
	// Specifies whether the application SSO requires PKCE (RFC 7636).
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
	// The list of logout callback addresses supported by the application.
	PostLogoutRedirectUris []*string `json:"PostLogoutRedirectUris,omitempty" xml:"PostLogoutRedirectUris,omitempty" type:"Repeated"`
	// The list of redirect URIs supported by the application.
	RedirectUris []*string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	// The validity period of the issued refresh token. Unit: seconds. Default value: 86400 (1 day).
	//
	// example:
	//
	// 86400
	RefreshTokenEffective *int64 `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
	// The response types supported by the application when OidcSsoConfig.GrantTypes includes the implicit mode.
	//
	// example:
	//
	// token id_token
	ResponseTypes []*string `json:"ResponseTypes,omitempty" xml:"ResponseTypes,omitempty" type:"Repeated"`
	// The custom expression for the sub value returned in the ID token.
	//
	// example:
	//
	// user.userid
	SubjectIdExpression *string `json:"SubjectIdExpression,omitempty" xml:"SubjectIdExpression,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetAccessTokenEffectiveTime() *int64 {
	return s.AccessTokenEffectiveTime
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetAllowedPublicClient() *string {
	return s.AllowedPublicClient
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetCodeEffectiveTime() *int64 {
	return s.CodeEffectiveTime
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetCustomClaims() []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims {
	return s.CustomClaims
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetGrantTypes() []*string {
	return s.GrantTypes
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetIdTokenEffectiveTime() *int64 {
	return s.IdTokenEffectiveTime
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetPasswordAuthenticationSourceId() *string {
	return s.PasswordAuthenticationSourceId
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetPasswordTotpMfaRequired() *bool {
	return s.PasswordTotpMfaRequired
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetPkceChallengeMethods() []*string {
	return s.PkceChallengeMethods
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetPostLogoutRedirectUris() []*string {
	return s.PostLogoutRedirectUris
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetRedirectUris() []*string {
	return s.RedirectUris
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetRefreshTokenEffective() *int64 {
	return s.RefreshTokenEffective
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetResponseTypes() []*string {
	return s.ResponseTypes
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GetSubjectIdExpression() *string {
	return s.SubjectIdExpression
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetAccessTokenEffectiveTime(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.AccessTokenEffectiveTime = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetAllowedPublicClient(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.AllowedPublicClient = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetCodeEffectiveTime(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.CodeEffectiveTime = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetCustomClaims(v []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.CustomClaims = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetGrantScopes(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.GrantScopes = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetGrantTypes(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.GrantTypes = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetIdTokenEffectiveTime(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.IdTokenEffectiveTime = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPasswordAuthenticationSourceId(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PasswordAuthenticationSourceId = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPasswordTotpMfaRequired(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PasswordTotpMfaRequired = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPkceChallengeMethods(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PkceChallengeMethods = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPkceRequired(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PkceRequired = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPostLogoutRedirectUris(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PostLogoutRedirectUris = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetRedirectUris(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.RedirectUris = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetRefreshTokenEffective(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.RefreshTokenEffective = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetResponseTypes(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.ResponseTypes = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetSubjectIdExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.SubjectIdExpression = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) Validate() error {
	if s.CustomClaims != nil {
		for _, item := range s.CustomClaims {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims struct {
	// The name of the returned claim.
	//
	// example:
	//
	// userOuIds
	ClaimName *string `json:"ClaimName,omitempty" xml:"ClaimName,omitempty"`
	// The value expression of the returned claim.
	//
	// example:
	//
	// ObjectToJsonString(user.organizationalUnits)
	ClaimValueExpression *string `json:"ClaimValueExpression,omitempty" xml:"ClaimValueExpression,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) GetClaimName() *string {
	return s.ClaimName
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) GetClaimValueExpression() *string {
	return s.ClaimValueExpression
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) SetClaimName(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims {
	s.ClaimName = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) SetClaimValueExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims {
	s.ClaimValueExpression = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) Validate() error {
	return dara.Validate(s)
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain struct {
	// The OAuth 2.0 authorization endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/login/app/app_mltta64q65enci54slingvvsgq/oauth2/authorize
	Oauth2AuthorizationEndpoint *string `json:"Oauth2AuthorizationEndpoint,omitempty" xml:"Oauth2AuthorizationEndpoint,omitempty"`
	// The OAuth 2.0 device authorization endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/device/code
	Oauth2DeviceAuthorizationEndpoint *string `json:"Oauth2DeviceAuthorizationEndpoint,omitempty" xml:"Oauth2DeviceAuthorizationEndpoint,omitempty"`
	// The OAuth 2.0 token revocation endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/revoke
	Oauth2RevokeEndpoint *string `json:"Oauth2RevokeEndpoint,omitempty" xml:"Oauth2RevokeEndpoint,omitempty"`
	// The OAuth 2.0 token endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/token
	Oauth2TokenEndpoint *string `json:"Oauth2TokenEndpoint,omitempty" xml:"Oauth2TokenEndpoint,omitempty"`
	// The OIDC user information endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/userinfo
	Oauth2UserinfoEndpoint *string `json:"Oauth2UserinfoEndpoint,omitempty" xml:"Oauth2UserinfoEndpoint,omitempty"`
	// The OIDC issuer information. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oidc
	OidcIssuer *string `json:"OidcIssuer,omitempty" xml:"OidcIssuer,omitempty"`
	// The OIDC JWKS endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oidc/jwks
	OidcJwksEndpoint *string `json:"OidcJwksEndpoint,omitempty" xml:"OidcJwksEndpoint,omitempty"`
	// The OIDC RP-initiated logout endpoint. This parameter is returned only when the application SSO protocol is OIDC.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/login/app/app_mltta64q65enci54slingvvsgq/oauth2/logout
	OidcLogoutEndpoint *string `json:"OidcLogoutEndpoint,omitempty" xml:"OidcLogoutEndpoint,omitempty"`
	// The SAML protocol metadata endpoint URL. This parameter is returned only when the application SSO protocol is SAML 2.0.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/api/v2/app_mltuxdwd4lq4eer6tmtlmaxm5e/saml2/meta
	SamlMetaEndpoint *string `json:"SamlMetaEndpoint,omitempty" xml:"SamlMetaEndpoint,omitempty"`
	// The SAML single logout URL (SLO URL) on the IdP side. The SP redirects the user to this URL to initiate single logout.
	//
	// example:
	//
	// https://example.com/saml/slo
	SamlSloEndpoint *string `json:"SamlSloEndpoint,omitempty" xml:"SamlSloEndpoint,omitempty"`
	// The SAML protocol AuthnRequest receiving endpoint. This parameter is returned only when the application SSO protocol is SAML 2.0.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/login/app/app_mltuxdwd4lq4eer6tmtlmaxm5e/saml2/sso
	SamlSsoEndpoint *string `json:"SamlSsoEndpoint,omitempty" xml:"SamlSsoEndpoint,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOauth2AuthorizationEndpoint() *string {
	return s.Oauth2AuthorizationEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOauth2DeviceAuthorizationEndpoint() *string {
	return s.Oauth2DeviceAuthorizationEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOauth2RevokeEndpoint() *string {
	return s.Oauth2RevokeEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOauth2TokenEndpoint() *string {
	return s.Oauth2TokenEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOauth2UserinfoEndpoint() *string {
	return s.Oauth2UserinfoEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOidcIssuer() *string {
	return s.OidcIssuer
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOidcJwksEndpoint() *string {
	return s.OidcJwksEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetOidcLogoutEndpoint() *string {
	return s.OidcLogoutEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetSamlMetaEndpoint() *string {
	return s.SamlMetaEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetSamlSloEndpoint() *string {
	return s.SamlSloEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GetSamlSsoEndpoint() *string {
	return s.SamlSsoEndpoint
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2AuthorizationEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2AuthorizationEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2DeviceAuthorizationEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2DeviceAuthorizationEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2RevokeEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2RevokeEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2TokenEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2TokenEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2UserinfoEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2UserinfoEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOidcIssuer(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.OidcIssuer = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOidcJwksEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.OidcJwksEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOidcLogoutEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.OidcLogoutEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetSamlMetaEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.SamlMetaEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetSamlSloEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.SamlSloEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetSamlSsoEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.SamlSsoEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) Validate() error {
	return dara.Validate(s)
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig struct {
	// Specifies whether the assertion needs to be signed. ResponseSigned and AssertionSigned cannot both be set to false.
	//
	// - true: Signed.
	//
	// - false: Not signed.
	//
	// example:
	//
	// true
	AssertionSigned *bool `json:"AssertionSigned,omitempty" xml:"AssertionSigned,omitempty"`
	// The additional user attribute configuration included in the SAML assertion.
	AttributeStatements []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements `json:"AttributeStatements,omitempty" xml:"AttributeStatements,omitempty" type:"Repeated"`
	// The default RelayState value. When the single sign-on (SSO) request is initiated by EIAM, the SAML Response provided by EIAM specifies the RelayState as this value. This applies when the user logon request is initiated by EIAM.
	//
	// example:
	//
	// https://home.console.aliyun.com
	DefaultRelayState *string `json:"DefaultRelayState,omitempty" xml:"DefaultRelayState,omitempty"`
	// The Entity ID that represents the IdP identity in the SAML protocol.
	//
	// example:
	//
	// https://example.com/
	IdPEntityId *string `json:"IdPEntityId,omitempty" xml:"IdPEntityId,omitempty"`
	// The NameID format defined by the SAML protocol standard. Valid values:
	//
	// - urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified: Unspecified. The application determines how to parse the NameID.
	//
	// - urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress: Email address format.
	//
	// - urn:oasis:names:tc:SAML:2.0:nameid-format:persistent: Persistent NameID.
	//
	// - urn:oasis:names:tc:SAML:2.0:nameid-format:transient: Transient NameID.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified
	NameIdFormat *string `json:"NameIdFormat,omitempty" xml:"NameIdFormat,omitempty"`
	// The expression used to generate the actual NameID value in the SAML protocol.
	//
	// example:
	//
	// user.username
	NameIdValueExpression *string `json:"NameIdValueExpression,omitempty" xml:"NameIdValueExpression,omitempty"`
	// The optional RelayState values. In the application portal, the application card displays multiple optional redirect addresses with display names. After a user clicks an address and completes SSO, the user is automatically redirected to the corresponding address. You can specify optional redirect addresses only after you specify a default redirect address.
	OptionalRelayStates []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates `json:"OptionalRelayStates,omitempty" xml:"OptionalRelayStates,omitempty" type:"Repeated"`
	// Indicates whether SSO AuthnRequest signature verification is enabled.
	//
	// example:
	//
	// true
	RequireAuthnRequestSigned *bool `json:"RequireAuthnRequestSigned,omitempty" xml:"RequireAuthnRequestSigned,omitempty"`
	// Indicates whether the Response needs to be signed. ResponseSigned and AssertionSigned cannot both be set to false. Valid values:
	//
	// - true: Signing is required.
	//
	// - false: Signing is not required.
	//
	// example:
	//
	// true
	ResponseSigned *bool `json:"ResponseSigned,omitempty" xml:"ResponseSigned,omitempty"`
	// The signature algorithm for the SAML assertion.
	//
	// example:
	//
	// RSA-SHA256
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The SAML EntityId of the application (SP).
	//
	// example:
	//
	// urn:alibaba:cloudcomputing
	SpEntityId *string `json:"SpEntityId,omitempty" xml:"SpEntityId,omitempty"`
	// The configured SP signing verification certificates in PEM format. A maximum of two certificates are returned for the console or API caller to read and display.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIC0jCCAbqgAwIBAgIQXXXXX -----END CERTIFICATE-----
	SpSigningCertificates []*string `json:"SpSigningCertificates,omitempty" xml:"SpSigningCertificates,omitempty" type:"Repeated"`
	// The configured SP SLO response URL.
	//
	// example:
	//
	// https://example.com/api/slo/response
	SpSloResponseUrl *string `json:"SpSloResponseUrl,omitempty" xml:"SpSloResponseUrl,omitempty"`
	// The SAML Assertion Consumer Service (ACS) URL of the application (SP).
	//
	// example:
	//
	// https://signin.aliyun.com/saml-role/sso
	SpSsoAcsUrl *string `json:"SpSsoAcsUrl,omitempty" xml:"SpSsoAcsUrl,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetAssertionSigned() *bool {
	return s.AssertionSigned
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetAttributeStatements() []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements {
	return s.AttributeStatements
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetDefaultRelayState() *string {
	return s.DefaultRelayState
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetIdPEntityId() *string {
	return s.IdPEntityId
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetNameIdFormat() *string {
	return s.NameIdFormat
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetNameIdValueExpression() *string {
	return s.NameIdValueExpression
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetOptionalRelayStates() []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates {
	return s.OptionalRelayStates
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetRequireAuthnRequestSigned() *bool {
	return s.RequireAuthnRequestSigned
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetResponseSigned() *bool {
	return s.ResponseSigned
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetSpEntityId() *string {
	return s.SpEntityId
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetSpSigningCertificates() []*string {
	return s.SpSigningCertificates
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetSpSloResponseUrl() *string {
	return s.SpSloResponseUrl
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GetSpSsoAcsUrl() *string {
	return s.SpSsoAcsUrl
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetAssertionSigned(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.AssertionSigned = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetAttributeStatements(v []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.AttributeStatements = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetDefaultRelayState(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.DefaultRelayState = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetIdPEntityId(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.IdPEntityId = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetNameIdFormat(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.NameIdFormat = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetNameIdValueExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.NameIdValueExpression = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetOptionalRelayStates(v []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.OptionalRelayStates = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetRequireAuthnRequestSigned(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.RequireAuthnRequestSigned = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetResponseSigned(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.ResponseSigned = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSignatureAlgorithm(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSpEntityId(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SpEntityId = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSpSigningCertificates(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SpSigningCertificates = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSpSloResponseUrl(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SpSloResponseUrl = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSpSsoAcsUrl(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SpSsoAcsUrl = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) Validate() error {
	if s.AttributeStatements != nil {
		for _, item := range s.AttributeStatements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OptionalRelayStates != nil {
		for _, item := range s.OptionalRelayStates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements struct {
	// The Name of the attribute in the SAML assertion.
	//
	// example:
	//
	// https://www.aliyun.com/SAML-Role/Attributes/RoleSessionName
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The attribute value expression in the SAML assertion.
	//
	// example:
	//
	// user.username
	AttributeValueExpression *string `json:"AttributeValueExpression,omitempty" xml:"AttributeValueExpression,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) GetAttributeName() *string {
	return s.AttributeName
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) GetAttributeValueExpression() *string {
	return s.AttributeValueExpression
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) SetAttributeName(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements {
	s.AttributeName = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) SetAttributeValueExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements {
	s.AttributeValueExpression = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) Validate() error {
	return dara.Validate(s)
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates struct {
	// The display name of the RelayState.
	//
	// example:
	//
	// Ram Account SSO
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The optional RelayState value. In the application portal, the application card displays multiple optional redirect addresses with display names. After a user clicks an address and completes SSO, the user is automatically redirected to the corresponding address.
	//
	// example:
	//
	// https://home.console.aliyun.com
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) GetRelayState() *string {
	return s.RelayState
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) SetDisplayName(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates {
	s.DisplayName = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) SetRelayState(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates {
	s.RelayState = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigOptionalRelayStates) Validate() error {
	return dara.Validate(s)
}
