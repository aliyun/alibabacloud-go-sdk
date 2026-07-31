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
	// The idempotency token that ensures the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see References [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration of the credential provider.
	CredentialProviderConfig *CreateCredentialProviderRequestCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// The identifier of the credential provider.
	//
	// > Allowed characters include uppercase and lowercase letters, digits, and the special characters `.-_`. The length cannot exceed 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"CredentialProviderIdentifier,omitempty" xml:"CredentialProviderIdentifier,omitempty"`
	// The name of the credential provider.
	//
	// > The length cannot exceed 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// The type of the credential provider. Valid values:
	//
	// - oauth: OAuth credential provider.
	//
	// - jwt: JWT credential provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// oauth
	CredentialProviderType *string `json:"CredentialProviderType,omitempty" xml:"CredentialProviderType,omitempty"`
	// The description.
	//
	// > The length cannot exceed 128 characters.
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
	// > The list length cannot exceed 200.
	AllowedTokenIssuers []*string `json:"AllowedTokenIssuers,omitempty" xml:"AllowedTokenIssuers,omitempty" type:"Repeated"`
	// Specifies whether to enable the JWT derived short token capability.
	//
	// example:
	//
	// false
	DerivedShortTokenEnabled *bool `json:"DerivedShortTokenEnabled,omitempty" xml:"DerivedShortTokenEnabled,omitempty"`
	// The validity duration of the JWT. Unit: seconds.
	//
	// example:
	//
	// 900
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// Specifies whether to enable JWT expiration cleanup.
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
	// The endpoint address used to guide users through authorization. Conditionally required: required when AuthorizationFlow=user_federation and ProviderVendor=custom. For preset vendors, this can be automatically populated through DiscoveryUrl.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The OAuth authorization flow type. Valid values: m2m: machine-to-machine (2LO, Client Credentials). user_federation: user federation (3LO, Authorization Code).
	AuthorizationFlow *string `json:"AuthorizationFlow,omitempty" xml:"AuthorizationFlow,omitempty"`
	// The client_id in the OAuth protocol.
	//
	// > The length cannot exceed 128 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_id_example_xxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client_secret in the OAuth protocol.
	//
	// > The length cannot exceed 1024 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_secret_example_xxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The Discovery document URL used to automatically retrieve OAuth endpoint configurations. Conditionally optional: used when AuthorizationFlow=user_federation. If DiscoveryUrl is not provided, you must manually configure fields such as TokenEndpoint and AuthorizationEndpoint.
	DiscoveryUrl *string `json:"DiscoveryUrl,omitempty" xml:"DiscoveryUrl,omitempty"`
	Issuer       *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The PKCE code_challenge generation method. Default value: s256.
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Specifies whether to use the PKCE extension to enhance security. We recommend that you always enable this feature.
	PkceEnabled *bool `json:"PkceEnabled,omitempty" xml:"PkceEnabled,omitempty"`
	// The preset vendor or custom configuration. Optional. Default value: custom.
	ProviderVendor *string `json:"ProviderVendor,omitempty" xml:"ProviderVendor,omitempty"`
	// The scope in the OAuth protocol, which defines the permission range.
	//
	// > The Scope configuration on the credential provider serves as the fallback value. If the scope parameter is not specified when calling the DeveloperAPI to obtain an OAuth Access Token, the Scope configuration on the credential provider is used for issuance.
	//
	// 	Notice: Separate multiple Scope values with spaces.
	//
	// Restrictions for each individual Scope value:
	//
	// 1. Allowed characters: lowercase letters, digits, and the special characters `|/:_-.`
	//
	// 2. Must contain at least one lowercase letter or digit.
	//
	// 3. Must start with the special character `.`, a lowercase letter, or a digit.
	//
	// 4. The length cannot exceed 1024 characters.
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The token endpoint of the OAuth protocol.
	//
	// > Must start with `http://` or `https://`, and the length cannot exceed 1024 characters.
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
