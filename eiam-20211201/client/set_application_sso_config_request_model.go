// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationSsoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetApplicationSsoConfigRequest
	GetApplicationId() *string
	SetClientToken(v string) *SetApplicationSsoConfigRequest
	GetClientToken() *string
	SetInitLoginType(v string) *SetApplicationSsoConfigRequest
	GetInitLoginType() *string
	SetInitLoginUrl(v string) *SetApplicationSsoConfigRequest
	GetInitLoginUrl() *string
	SetInstanceId(v string) *SetApplicationSsoConfigRequest
	GetInstanceId() *string
	SetOidcSsoConfig(v *SetApplicationSsoConfigRequestOidcSsoConfig) *SetApplicationSsoConfigRequest
	GetOidcSsoConfig() *SetApplicationSsoConfigRequestOidcSsoConfig
	SetSamlSsoConfig(v *SetApplicationSsoConfigRequestSamlSsoConfig) *SetApplicationSsoConfigRequest
	GetSamlSsoConfig() *SetApplicationSsoConfigRequestSamlSsoConfig
}

type SetApplicationSsoConfigRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique value from your client. The value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// client-examplexxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The initial single sign-on (SSO) logon method. Valid values:
	//
	// - only_app_init_sso: Only application-initiated SSO is supported. This is the default value for OIDC protocol applications. When a SAML application uses this method, InitLoginUrl must be specified.
	//
	// - idaas_or_app_init_sso: Both IDaaS portal-initiated and application-initiated SSO are supported. This is the default value for SAML protocol applications. When an OIDC protocol application uses this method, InitLoginUrl must be specified.
	//
	// example:
	//
	// only_app_init_sso
	InitLoginType *string `json:"InitLoginType,omitempty" xml:"InitLoginType,omitempty"`
	// The initial single sign-on (SSO) logon trigger URL.
	//
	// This parameter is required when an OIDC protocol application sets InitLoginType to idaas_or_app_init_sso.
	//
	// This parameter is required when a SAML protocol application sets InitLoginType to only_app_init_sso.
	//
	// example:
	//
	// http://127.0.0.1:8000/start_login?enterprise_code=ABCDEF
	InitLoginUrl *string `json:"InitLoginUrl,omitempty" xml:"InitLoginUrl,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The SSO configuration parameters for an OIDC-based application.
	OidcSsoConfig *SetApplicationSsoConfigRequestOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	// The SSO configuration parameters for a SAML-based application.
	SamlSsoConfig *SetApplicationSsoConfigRequestSamlSsoConfig `json:"SamlSsoConfig,omitempty" xml:"SamlSsoConfig,omitempty" type:"Struct"`
}

func (s SetApplicationSsoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetApplicationSsoConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetApplicationSsoConfigRequest) GetInitLoginType() *string {
	return s.InitLoginType
}

func (s *SetApplicationSsoConfigRequest) GetInitLoginUrl() *string {
	return s.InitLoginUrl
}

func (s *SetApplicationSsoConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetApplicationSsoConfigRequest) GetOidcSsoConfig() *SetApplicationSsoConfigRequestOidcSsoConfig {
	return s.OidcSsoConfig
}

func (s *SetApplicationSsoConfigRequest) GetSamlSsoConfig() *SetApplicationSsoConfigRequestSamlSsoConfig {
	return s.SamlSsoConfig
}

func (s *SetApplicationSsoConfigRequest) SetApplicationId(v string) *SetApplicationSsoConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetClientToken(v string) *SetApplicationSsoConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetInitLoginType(v string) *SetApplicationSsoConfigRequest {
	s.InitLoginType = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetInitLoginUrl(v string) *SetApplicationSsoConfigRequest {
	s.InitLoginUrl = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetInstanceId(v string) *SetApplicationSsoConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetOidcSsoConfig(v *SetApplicationSsoConfigRequestOidcSsoConfig) *SetApplicationSsoConfigRequest {
	s.OidcSsoConfig = v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetSamlSsoConfig(v *SetApplicationSsoConfigRequestSamlSsoConfig) *SetApplicationSsoConfigRequest {
	s.SamlSsoConfig = v
	return s
}

func (s *SetApplicationSsoConfigRequest) Validate() error {
	if s.OidcSsoConfig != nil {
		if err := s.OidcSsoConfig.Validate(); err != nil {
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

type SetApplicationSsoConfigRequestOidcSsoConfig struct {
	// The validity period of the issued access token. Unit: seconds. Default value: 1200 (20 minutes).
	//
	// example:
	//
	// 1200
	AccessTokenEffectiveTime *int64 `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	// Specifies whether the application is allowed to act as a public client to request the IDaaS EIAM authorization server. This parameter can be enabled only in authorization code mode and device mode. Default value: false.
	//
	// example:
	//
	// true
	AllowedPublicClient *bool `json:"AllowedPublicClient,omitempty" xml:"AllowedPublicClient,omitempty"`
	// The validity period of the issued code. Unit: seconds. Default value: 60 (1 minute).
	//
	// example:
	//
	// 60
	CodeEffectiveTime *int64 `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	// The custom user information included in the returned ID token.
	CustomClaims []*SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims `json:"CustomClaims,omitempty" xml:"CustomClaims,omitempty" type:"Repeated"`
	// The OIDC standard parameter scope, which specifies the range of user attributes that can be returned by the userinfo endpoint or the id_token.
	//
	// example:
	//
	// profile，email
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The list of supported OIDC grant types.
	//
	// example:
	//
	// authorization_code
	GrantTypes []*string `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	// The validity period of the issued ID token. Unit: seconds. Default value: 300 (5 minutes).
	//
	// example:
	//
	// 300
	IdTokenEffectiveTime *int64 `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	// The ID of the authentication source used in password mode. This parameter takes effect only when the GrantTypes specified for the OIDC application include the password mode.
	//
	// example:
	//
	// ia_password
	PasswordAuthenticationSourceId *string `json:"PasswordAuthenticationSourceId,omitempty" xml:"PasswordAuthenticationSourceId,omitempty"`
	// Specifies whether TOTP-based secondary authentication is required for password mode. This parameter takes effect only when the GrantTypes specified for the OIDC application include the password mode.
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
	// Specifies whether Proof Key for Code Exchange (PKCE) (RFC 7636) is required for application SSO.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
	// The list of logout callback URIs supported by the application.
	PostLogoutRedirectUris []*string `json:"PostLogoutRedirectUris,omitempty" xml:"PostLogoutRedirectUris,omitempty" type:"Repeated"`
	// The list of redirect URIs supported by the application.
	RedirectUris []*string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	// The validity period of the issued refresh token. Unit: seconds. Default value: 86400 (1 day).
	//
	// example:
	//
	// 86400
	RefreshTokenEffective *int64 `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
	// The response types supported by the application when OidcSsoConfig.GrantTypes includes the implicit grant type.
	//
	// example:
	//
	// token id_token
	ResponseTypes []*string `json:"ResponseTypes,omitempty" xml:"ResponseTypes,omitempty" type:"Repeated"`
	// The custom expression for the sub claim value returned in the ID token.
	//
	// example:
	//
	// user.userid
	SubjectIdExpression *string `json:"SubjectIdExpression,omitempty" xml:"SubjectIdExpression,omitempty"`
}

func (s SetApplicationSsoConfigRequestOidcSsoConfig) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigRequestOidcSsoConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetAccessTokenEffectiveTime() *int64 {
	return s.AccessTokenEffectiveTime
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetAllowedPublicClient() *bool {
	return s.AllowedPublicClient
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetCodeEffectiveTime() *int64 {
	return s.CodeEffectiveTime
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetCustomClaims() []*SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims {
	return s.CustomClaims
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetGrantTypes() []*string {
	return s.GrantTypes
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetIdTokenEffectiveTime() *int64 {
	return s.IdTokenEffectiveTime
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetPasswordAuthenticationSourceId() *string {
	return s.PasswordAuthenticationSourceId
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetPasswordTotpMfaRequired() *bool {
	return s.PasswordTotpMfaRequired
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetPkceChallengeMethods() []*string {
	return s.PkceChallengeMethods
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetPostLogoutRedirectUris() []*string {
	return s.PostLogoutRedirectUris
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetRedirectUris() []*string {
	return s.RedirectUris
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetRefreshTokenEffective() *int64 {
	return s.RefreshTokenEffective
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetResponseTypes() []*string {
	return s.ResponseTypes
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) GetSubjectIdExpression() *string {
	return s.SubjectIdExpression
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetAccessTokenEffectiveTime(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.AccessTokenEffectiveTime = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetAllowedPublicClient(v bool) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.AllowedPublicClient = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetCodeEffectiveTime(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.CodeEffectiveTime = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetCustomClaims(v []*SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.CustomClaims = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetGrantScopes(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.GrantScopes = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetGrantTypes(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.GrantTypes = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetIdTokenEffectiveTime(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.IdTokenEffectiveTime = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPasswordAuthenticationSourceId(v string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PasswordAuthenticationSourceId = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPasswordTotpMfaRequired(v bool) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PasswordTotpMfaRequired = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPkceChallengeMethods(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PkceChallengeMethods = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPkceRequired(v bool) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PkceRequired = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPostLogoutRedirectUris(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PostLogoutRedirectUris = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetRedirectUris(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.RedirectUris = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetRefreshTokenEffective(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.RefreshTokenEffective = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetResponseTypes(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.ResponseTypes = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetSubjectIdExpression(v string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.SubjectIdExpression = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) Validate() error {
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

type SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims struct {
	// The name of the returned claim.
	//
	// example:
	//
	// "Role"
	ClaimName *string `json:"ClaimName,omitempty" xml:"ClaimName,omitempty"`
	// The value expression of the returned claim.
	//
	// example:
	//
	// user.dict.applicationRole
	ClaimValueExpression *string `json:"ClaimValueExpression,omitempty" xml:"ClaimValueExpression,omitempty"`
}

func (s SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) GetClaimName() *string {
	return s.ClaimName
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) GetClaimValueExpression() *string {
	return s.ClaimValueExpression
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) SetClaimName(v string) *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims {
	s.ClaimName = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) SetClaimValueExpression(v string) *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims {
	s.ClaimValueExpression = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) Validate() error {
	return dara.Validate(s)
}

type SetApplicationSsoConfigRequestSamlSsoConfig struct {
	// Specifies whether the assertion needs to be signed. ResponseSigned and AssertionSigned cannot both be set to false.
	//
	// - true: The assertion is signed.
	//
	// - false: The assertion is not signed.
	//
	// example:
	//
	// true
	AssertionSigned *bool `json:"AssertionSigned,omitempty" xml:"AssertionSigned,omitempty"`
	// The additional user attribute configurations included in the SAML assertion.
	AttributeStatements []*SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements `json:"AttributeStatements,omitempty" xml:"AttributeStatements,omitempty" type:"Repeated"`
	// The default RelayState value. When a single sign-on (SSO) logon request is initiated by EIAM, the SAML Response provided by EIAM specifies the RelayState as this value.
	//
	// example:
	//
	// https://home.console.aliyun.com
	DefaultRelayState *string `json:"DefaultRelayState,omitempty" xml:"DefaultRelayState,omitempty"`
	// The Entity ID that represents the IdP identity in the SAML protocol. URL format and URN format are supported.
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
	// user.email
	NameIdValueExpression *string `json:"NameIdValueExpression,omitempty" xml:"NameIdValueExpression,omitempty"`
	// The optional RelayState configurations.
	OptionalRelayStates []*SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates `json:"OptionalRelayStates,omitempty" xml:"OptionalRelayStates,omitempty" type:"Repeated"`
	// Specifies whether SSO AuthnRequest signature verification is enabled. Default value: false. If set to true, spSigningCertificates must be configured (the array must not be empty).
	//
	// example:
	//
	// false
	RequireAuthnRequestSigned *bool `json:"RequireAuthnRequestSigned,omitempty" xml:"RequireAuthnRequestSigned,omitempty"`
	// Specifies whether the response needs to be signed. ResponseSigned and AssertionSigned cannot both be set to false.
	//
	// - true: The response is signed.
	//
	// - false: The response is not signed.
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
	// The array of SP signature verification certificates in PEM format. A maximum of two certificates are allowed and are shared by SSO and SLO. Each certificate is validated for format and validity upon write. Requests with more than two certificates are rejected.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIC0jCCAbqgAwIBAgIQXXXXX -----END CERTIFICATE-----
	SpSigningCertificates []*string `json:"SpSigningCertificates,omitempty" xml:"SpSigningCertificates,omitempty" type:"Repeated"`
	// The URL on the SP side that receives the LogoutResponse. This parameter is optional.
	//
	// example:
	//
	// https://example.com/api/slo/response
	SpSloResponseUrl *string `json:"SpSloResponseUrl,omitempty" xml:"SpSloResponseUrl,omitempty"`
	// The SAML assertion consumer service URL of the application (SP).
	//
	// example:
	//
	// https://signin.aliyun.com/saml-role/sso
	SpSsoAcsUrl *string `json:"SpSsoAcsUrl,omitempty" xml:"SpSsoAcsUrl,omitempty"`
}

func (s SetApplicationSsoConfigRequestSamlSsoConfig) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigRequestSamlSsoConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetAssertionSigned() *bool {
	return s.AssertionSigned
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetAttributeStatements() []*SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements {
	return s.AttributeStatements
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetDefaultRelayState() *string {
	return s.DefaultRelayState
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetIdPEntityId() *string {
	return s.IdPEntityId
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetNameIdFormat() *string {
	return s.NameIdFormat
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetNameIdValueExpression() *string {
	return s.NameIdValueExpression
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetOptionalRelayStates() []*SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates {
	return s.OptionalRelayStates
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetRequireAuthnRequestSigned() *bool {
	return s.RequireAuthnRequestSigned
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetResponseSigned() *bool {
	return s.ResponseSigned
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSpEntityId() *string {
	return s.SpEntityId
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSpSigningCertificates() []*string {
	return s.SpSigningCertificates
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSpSloResponseUrl() *string {
	return s.SpSloResponseUrl
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSpSsoAcsUrl() *string {
	return s.SpSsoAcsUrl
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetAssertionSigned(v bool) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.AssertionSigned = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetAttributeStatements(v []*SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.AttributeStatements = v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetDefaultRelayState(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.DefaultRelayState = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetIdPEntityId(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.IdPEntityId = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetNameIdFormat(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.NameIdFormat = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetNameIdValueExpression(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.NameIdValueExpression = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetOptionalRelayStates(v []*SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.OptionalRelayStates = v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetRequireAuthnRequestSigned(v bool) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.RequireAuthnRequestSigned = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetResponseSigned(v bool) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.ResponseSigned = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSignatureAlgorithm(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SignatureAlgorithm = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSpEntityId(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SpEntityId = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSpSigningCertificates(v []*string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SpSigningCertificates = v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSpSloResponseUrl(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SpSloResponseUrl = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSpSsoAcsUrl(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SpSsoAcsUrl = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) Validate() error {
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

type SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements struct {
	// The name of the attribute in the SAML assertion.
	//
	// example:
	//
	// https://www.aliyun.com/SAML-Role/Attributes/RoleSessionName
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The value expression of the attribute in the SAML assertion.
	//
	// example:
	//
	// user.username
	AttributeValueExpression *string `json:"AttributeValueExpression,omitempty" xml:"AttributeValueExpression,omitempty"`
}

func (s SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) GetAttributeName() *string {
	return s.AttributeName
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) GetAttributeValueExpression() *string {
	return s.AttributeValueExpression
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) SetAttributeName(v string) *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements {
	s.AttributeName = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) SetAttributeValueExpression(v string) *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements {
	s.AttributeValueExpression = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) Validate() error {
	return dara.Validate(s)
}

type SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates struct {
	// The display name of the RelayState.
	//
	// example:
	//
	// Ram
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The RelayState value.
	//
	// example:
	//
	// https://ram.console.aliyun.com/
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
}

func (s SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) GetDisplayName() *string {
	return s.DisplayName
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) GetRelayState() *string {
	return s.RelayState
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) SetDisplayName(v string) *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates {
	s.DisplayName = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) SetRelayState(v string) *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates {
	s.RelayState = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates) Validate() error {
	return dara.Validate(s)
}
