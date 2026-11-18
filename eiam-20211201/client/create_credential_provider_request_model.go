// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCredentialProviderRequest
	GetClientToken() *string
	SetCredentialProviderConfig(v *CreateCredentialProviderRequestCredentialProviderConfig) *CreateCredentialProviderRequest
	GetCredentialProviderConfig() *CreateCredentialProviderRequestCredentialProviderConfig
	SetCredentialProviderIdentifier(v string) *CreateCredentialProviderRequest
	GetCredentialProviderIdentifier() *string
	SetCredentialProviderName(v string) *CreateCredentialProviderRequest
	GetCredentialProviderName() *string
	SetCredentialProviderType(v string) *CreateCredentialProviderRequest
	GetCredentialProviderType() *string
	SetDescription(v string) *CreateCredentialProviderRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCredentialProviderRequest
	GetInstanceId() *string
}

type CreateCredentialProviderRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The credential provider configuration.
	CredentialProviderConfig *CreateCredentialProviderRequestCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// The credential provider identifier.
	//
	// > Valid characters: uppercase and lowercase letters, digits, and the following special characters: `.-_`. The identifier can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"CredentialProviderIdentifier,omitempty" xml:"CredentialProviderIdentifier,omitempty"`
	// The credential provider name.
	//
	// > The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// The credential provider type. Valid values:
	//
	// - oauth: OAuth credential provider
	//
	// - jwt: JWT credential provider
	//
	// This parameter is required.
	//
	// example:
	//
	// oauth
	CredentialProviderType *string `json:"CredentialProviderType,omitempty" xml:"CredentialProviderType,omitempty"`
	// The description.
	//
	// > The description can be up to 128 characters in length.
	//
	// example:
	//
	// This is an example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCredentialProviderRequest) GetCredentialProviderConfig() *CreateCredentialProviderRequestCredentialProviderConfig {
	return s.CredentialProviderConfig
}

func (s *CreateCredentialProviderRequest) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *CreateCredentialProviderRequest) GetCredentialProviderName() *string {
	return s.CredentialProviderName
}

func (s *CreateCredentialProviderRequest) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *CreateCredentialProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCredentialProviderRequest) SetClientToken(v string) *CreateCredentialProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCredentialProviderRequest) SetCredentialProviderConfig(v *CreateCredentialProviderRequestCredentialProviderConfig) *CreateCredentialProviderRequest {
	s.CredentialProviderConfig = v
	return s
}

func (s *CreateCredentialProviderRequest) SetCredentialProviderIdentifier(v string) *CreateCredentialProviderRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *CreateCredentialProviderRequest) SetCredentialProviderName(v string) *CreateCredentialProviderRequest {
	s.CredentialProviderName = &v
	return s
}

func (s *CreateCredentialProviderRequest) SetCredentialProviderType(v string) *CreateCredentialProviderRequest {
	s.CredentialProviderType = &v
	return s
}

func (s *CreateCredentialProviderRequest) SetDescription(v string) *CreateCredentialProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateCredentialProviderRequest) SetInstanceId(v string) *CreateCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCredentialProviderRequest) Validate() error {
	if s.CredentialProviderConfig != nil {
		if err := s.CredentialProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialProviderRequestCredentialProviderConfig struct {
	// The configuration of the JWT credential provider.
	JwtProviderConfig *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// The configuration of the OAuth credential provider.
	OAuthProviderConfig *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig `json:"OAuthProviderConfig,omitempty" xml:"OAuthProviderConfig,omitempty" type:"Struct"`
}

func (s CreateCredentialProviderRequestCredentialProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialProviderRequestCredentialProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateCredentialProviderRequestCredentialProviderConfig) GetJwtProviderConfig() *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	return s.JwtProviderConfig
}

func (s *CreateCredentialProviderRequestCredentialProviderConfig) GetOAuthProviderConfig() *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	return s.OAuthProviderConfig
}

func (s *CreateCredentialProviderRequestCredentialProviderConfig) SetJwtProviderConfig(v *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) *CreateCredentialProviderRequestCredentialProviderConfig {
	s.JwtProviderConfig = v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfig) SetOAuthProviderConfig(v *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) *CreateCredentialProviderRequestCredentialProviderConfig {
	s.OAuthProviderConfig = v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfig) Validate() error {
	if s.JwtProviderConfig != nil {
		if err := s.JwtProviderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OAuthProviderConfig != nil {
		if err := s.OAuthProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig struct {
	// The list of allowed JWT issuers.
	//
	// > The list can contain up to 200 items.
	AllowedTokenIssuers []*string `json:"AllowedTokenIssuers,omitempty" xml:"AllowedTokenIssuers,omitempty" type:"Repeated"`
	// Specifies whether to enable the capability to derive short tokens from JWTs.
	//
	// example:
	//
	// false
	DerivedShortTokenEnabled *bool `json:"DerivedShortTokenEnabled,omitempty" xml:"DerivedShortTokenEnabled,omitempty"`
	// The validity period of the JWT, in seconds.
	//
	// example:
	//
	// 900
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// Specifies whether to enable cleanup of expired JWTs.
	//
	// example:
	//
	// true
	ExpirationCleanupEnabled *bool `json:"ExpirationCleanupEnabled,omitempty" xml:"ExpirationCleanupEnabled,omitempty"`
}

func (s CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetAllowedTokenIssuers() []*string {
	return s.AllowedTokenIssuers
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetDerivedShortTokenEnabled() *bool {
	return s.DerivedShortTokenEnabled
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetExpiration() *int32 {
	return s.Expiration
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetExpirationCleanupEnabled() *bool {
	return s.ExpirationCleanupEnabled
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetAllowedTokenIssuers(v []*string) *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.AllowedTokenIssuers = v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetDerivedShortTokenEnabled(v bool) *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.DerivedShortTokenEnabled = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetExpiration(v int32) *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.Expiration = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetExpirationCleanupEnabled(v bool) *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.ExpirationCleanupEnabled = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig struct {
	// The authorization endpoint in the OAuth protocol.
	//
	// example:
	//
	// https://example.com/authorize
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The OAuth authorization flow type. Valid values:
	//
	// - m2m: Machine-to-machine
	//
	// - user_federation: User federation
	//
	// example:
	//
	// m2m
	AuthorizationFlow *string `json:"AuthorizationFlow,omitempty" xml:"AuthorizationFlow,omitempty"`
	// The client ID in the OAuth protocol.
	//
	// > The client ID can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_id_example_xxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret in the OAuth protocol.
	//
	// > The client secret can be up to 1,024 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_secret_example_xxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The URL of the OAuth discovery document.
	//
	// example:
	//
	// https://example.com/.well-known/openid-configuration
	DiscoveryUrl *string `json:"DiscoveryUrl,omitempty" xml:"DiscoveryUrl,omitempty"`
	// The issuer endpoint in the OAuth protocol.
	//
	// example:
	//
	// https://example.com/issuer
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The algorithm used to calculate the code challenge in PKCE. Valid values:
	//
	// - S256
	//
	// - plain
	//
	// example:
	//
	// S256
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Specifies whether to enable Proof Key for Code Exchange (PKCE).
	//
	// example:
	//
	// true
	PkceEnabled *bool `json:"PkceEnabled,omitempty" xml:"PkceEnabled,omitempty"`
	// The OAuth provider vendor type. Valid values:
	//
	// - custom: Custom
	//
	// - dingtalk: DingTalk
	//
	// - feishu: Lark
	//
	// - github: GitHub
	//
	// - microsoft: Microsoft
	//
	// - google: Google
	//
	// example:
	//
	// custom
	ProviderVendor *string `json:"ProviderVendor,omitempty" xml:"ProviderVendor,omitempty"`
	// The scope in the OAuth protocol.
	//
	// > The scope configuration of the OAuth credential provider serves as a fallback value. If the scope parameter is not specified when you call the Developer API to obtain an OAuth access token, the scope configuration of the credential provider is used to issue the token.
	//
	// 	Notice: Separate multiple scope values with spaces.</notice>
	//
	// Limits on a single scope value:
	//
	// 1. Valid characters: lowercase letters, digits, and the following special characters: `|/:_-.`.
	//
	// 2. The value must contain lowercase letters or digits.
	//
	// 3. The value must start with a period (`.`), a lowercase letter, or a digit.
	//
	// 4. The value can be up to 1,024 characters in length.
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The token endpoint in the OAuth protocol.
	//
	// > The endpoint must start with `http://` or `https://` and can be up to 1,024 characters in length.
	//
	// example:
	//
	// https://example.com/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetAuthorizationFlow() *string {
	return s.AuthorizationFlow
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetClientId() *string {
	return s.ClientId
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetDiscoveryUrl() *string {
	return s.DiscoveryUrl
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetPkceChallengeMethod() *string {
	return s.PkceChallengeMethod
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetPkceEnabled() *bool {
	return s.PkceEnabled
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetProviderVendor() *string {
	return s.ProviderVendor
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetScope() *string {
	return s.Scope
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetAuthorizationEndpoint(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetAuthorizationFlow(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.AuthorizationFlow = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetClientId(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.ClientId = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetClientSecret(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.ClientSecret = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetDiscoveryUrl(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.DiscoveryUrl = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetIssuer(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.Issuer = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetPkceChallengeMethod(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.PkceChallengeMethod = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetPkceEnabled(v bool) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.PkceEnabled = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetProviderVendor(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.ProviderVendor = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetScope(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.Scope = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetTokenEndpoint(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) Validate() error {
	return dara.Validate(s)
}
