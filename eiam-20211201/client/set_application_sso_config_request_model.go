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
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// client-examplexxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The SSO initiation method. Valid values:
	//
	// - only_app_init_sso: SSO is initiated only by the application. This is the default value for OIDC applications. If you set this parameter to this value for a SAML application, you must specify InitLoginUrl.
	//
	// - idaas_or_app_init_sso: SSO can be initiated by the IDaaS console or the application. This is the default value for SAML applications. If you set this parameter to this value for an OIDC application, you must specify InitLoginUrl.
	//
	// example:
	//
	// only_app_init_sso
	InitLoginType *string `json:"InitLoginType,omitempty" xml:"InitLoginType,omitempty"`
	// The URL that is used to initiate SSO. You must specify this parameter if you set InitLoginType to idaas_or_app_init_sso for an OIDC application. You must specify this parameter if you set InitLoginType to only_app_init_sso for a SAML application.
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
	// The SSO properties for an application that uses the OIDC protocol.
	OidcSsoConfig *SetApplicationSsoConfigRequestOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	// The SSO properties for an application that uses the SAML protocol.
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
	// The validity period of the access token. Unit: seconds. Default value: 1200 (20 minutes).
	//
	// example:
	//
	// 1200
	AccessTokenEffectiveTime *int64 `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	// Specifies whether the application is allowed to act as a public client to request the IDaaS authorization server. This parameter can be enabled only for the authorization code grant type and the device authorization grant type. Default value: false.
	//
	// example:
	//
	// true
	AllowedPublicClient *bool `json:"AllowedPublicClient,omitempty" xml:"AllowedPublicClient,omitempty"`
	// The validity period of the authorization code. Unit: seconds. Default value: 60 (1 minute).
	//
	// example:
	//
	// 60
	CodeEffectiveTime *int64 `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	// The custom claims that are returned in the ID token.
	CustomClaims []*SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims `json:"CustomClaims,omitempty" xml:"CustomClaims,omitempty" type:"Repeated"`
	// The scope parameter in the OIDC protocol. This parameter specifies the scope of user information that can be returned by the userinfo endpoint or included in the ID token.
	//
	// example:
	//
	// profile，email
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The list of OIDC grant types that are supported.
	//
	// example:
	//
	// authorization_code
	GrantTypes []*string `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	// The validity period of the ID token. Unit: seconds. Default value: 300 (5 minutes).
	//
	// example:
	//
	// 300
	IdTokenEffectiveTime *int64 `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	// The ID of the identity source for the resource owner password credentials grant type. This parameter is valid only when the GrantTypes for the OIDC application is set to password.
	//
	// example:
	//
	// ia_password
	PasswordAuthenticationSourceId *string `json:"PasswordAuthenticationSourceId,omitempty" xml:"PasswordAuthenticationSourceId,omitempty"`
	// Specifies whether Time-based One-time Password (TOTP) multi-factor authentication (MFA) is required for the resource owner password credentials grant type. This parameter is valid only when the GrantTypes for the OIDC application is set to password.
	//
	// example:
	//
	// true
	PasswordTotpMfaRequired *bool `json:"PasswordTotpMfaRequired,omitempty" xml:"PasswordTotpMfaRequired,omitempty"`
	// The algorithm used to compute the code challenge in PKCE.
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
	// The list of post-logout redirect URIs that the application supports.
	PostLogoutRedirectUris []*string `json:"PostLogoutRedirectUris,omitempty" xml:"PostLogoutRedirectUris,omitempty" type:"Repeated"`
	// The list of redirect URIs that the application supports.
	RedirectUris []*string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	// The validity period of the refresh token. Unit: seconds. Default value: 86400 (1 day).
	//
	// example:
	//
	// 86400
	RefreshTokenEffective *int64 `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
	// The response type supported by the application when OidcSsoConfig.GrantTypes is set to implicit.
	//
	// example:
	//
	// token id_token
	ResponseTypes []*string `json:"ResponseTypes,omitempty" xml:"ResponseTypes,omitempty" type:"Repeated"`
	// The expression used to generate the value of the sub claim in the ID token.
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
	// The name of the claim.
	//
	// example:
	//
	// "Role"
	ClaimName *string `json:"ClaimName,omitempty" xml:"ClaimName,omitempty"`
	// The expression used to generate the value of the claim.
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
	// Specifies whether the assertion must be signed. ResponseSigned and AssertionSigned cannot both be false.
	//
	// - true: The assertion must be signed.
	//
	// - false: The assertion does not need to be signed.
	//
	// example:
	//
	// true
	AssertionSigned *bool `json:"AssertionSigned,omitempty" xml:"AssertionSigned,omitempty"`
	// The configurations of additional user attributes in the SAML assertion.
	AttributeStatements []*SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements `json:"AttributeStatements,omitempty" xml:"AttributeStatements,omitempty" type:"Repeated"`
	// The default value of RelayState. When an SSO request is initiated by IDaaS, the SAML response provided by IDaaS contains this value for RelayState.
	//
	// example:
	//
	// https://home.console.aliyun.com
	DefaultRelayState *string `json:"DefaultRelayState,omitempty" xml:"DefaultRelayState,omitempty"`
	// The entity ID of the identity provider (IdP) in the SAML protocol. The value can be in a URL or URN format.
	//
	// example:
	//
	// https://example.com/
	IdPEntityId *string `json:"IdPEntityId,omitempty" xml:"IdPEntityId,omitempty"`
	// The format of the NameID in the SAML protocol. Valid values:
	//
	// - urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified: The format is not specified. The application determines how to parse the NameID.
	//
	// - urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress: The email address format.
	//
	// - urn:oasis:names:tc:SAML:2.0:nameid-format:persistent: The persistent NameID.
	//
	// - urn:oasis:names:tc:SAML:2.0:nameid-format:transient: The transient NameID.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified
	NameIdFormat *string `json:"NameIdFormat,omitempty" xml:"NameIdFormat,omitempty"`
	// The expression used to generate the value of the NameID in the SAML protocol.
	//
	// example:
	//
	// user.email
	NameIdValueExpression *string `json:"NameIdValueExpression,omitempty" xml:"NameIdValueExpression,omitempty"`
	// The optional RelayState configurations.
	OptionalRelayStates []*SetApplicationSsoConfigRequestSamlSsoConfigOptionalRelayStates `json:"OptionalRelayStates,omitempty" xml:"OptionalRelayStates,omitempty" type:"Repeated"`
	// Specifies whether the response must be signed. ResponseSigned and AssertionSigned cannot both be false.
	//
	// - true: The response must be signed.
	//
	// - false: The response does not need to be signed.
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
	// The entity ID of the application (service provider) that uses SAML.
	//
	// example:
	//
	// urn:alibaba:cloudcomputing
	SpEntityId *string `json:"SpEntityId,omitempty" xml:"SpEntityId,omitempty"`
	// The SAML assertion consumer service (ACS) URL of the application (service provider).
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

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetResponseSigned() *bool {
	return s.ResponseSigned
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) GetSpEntityId() *string {
	return s.SpEntityId
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
	// The expression used to generate the value of the attribute in the SAML assertion.
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
	// The value of RelayState.
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
