// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCredentialProviderRequest
	GetClientToken() *string
	SetCredentialProviderConfig(v *UpdateCredentialProviderRequestCredentialProviderConfig) *UpdateCredentialProviderRequest
	GetCredentialProviderConfig() *UpdateCredentialProviderRequestCredentialProviderConfig
	SetCredentialProviderId(v string) *UpdateCredentialProviderRequest
	GetCredentialProviderId() *string
	SetCredentialProviderName(v string) *UpdateCredentialProviderRequest
	GetCredentialProviderName() *string
	SetInstanceId(v string) *UpdateCredentialProviderRequest
	GetInstanceId() *string
}

type UpdateCredentialProviderRequest struct {
	// The idempotency token that ensures the idempotence of the request.
	//
	// Generate a unique parameter value from your client to ensure that the value is unique among different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The credential provider configuration.
	CredentialProviderConfig *UpdateCredentialProviderRequestCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// The credential provider ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// The credential provider name.
	//
	// > The name cannot exceed 64 characters in length.
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialProviderRequest) GetCredentialProviderConfig() *UpdateCredentialProviderRequestCredentialProviderConfig {
	return s.CredentialProviderConfig
}

func (s *UpdateCredentialProviderRequest) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *UpdateCredentialProviderRequest) GetCredentialProviderName() *string {
	return s.CredentialProviderName
}

func (s *UpdateCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCredentialProviderRequest) SetClientToken(v string) *UpdateCredentialProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialProviderRequest) SetCredentialProviderConfig(v *UpdateCredentialProviderRequestCredentialProviderConfig) *UpdateCredentialProviderRequest {
	s.CredentialProviderConfig = v
	return s
}

func (s *UpdateCredentialProviderRequest) SetCredentialProviderId(v string) *UpdateCredentialProviderRequest {
	s.CredentialProviderId = &v
	return s
}

func (s *UpdateCredentialProviderRequest) SetCredentialProviderName(v string) *UpdateCredentialProviderRequest {
	s.CredentialProviderName = &v
	return s
}

func (s *UpdateCredentialProviderRequest) SetInstanceId(v string) *UpdateCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCredentialProviderRequest) Validate() error {
	if s.CredentialProviderConfig != nil {
		if err := s.CredentialProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialProviderRequestCredentialProviderConfig struct {
	// The configuration of the JWT credential provider.
	JwtProviderConfig *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// The configuration of the OAuth credential provider.
	OAuthProviderConfig *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig `json:"OAuthProviderConfig,omitempty" xml:"OAuthProviderConfig,omitempty" type:"Struct"`
}

func (s UpdateCredentialProviderRequestCredentialProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderRequestCredentialProviderConfig) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfig) GetJwtProviderConfig() *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	return s.JwtProviderConfig
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfig) GetOAuthProviderConfig() *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	return s.OAuthProviderConfig
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfig) SetJwtProviderConfig(v *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) *UpdateCredentialProviderRequestCredentialProviderConfig {
	s.JwtProviderConfig = v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfig) SetOAuthProviderConfig(v *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) *UpdateCredentialProviderRequestCredentialProviderConfig {
	s.OAuthProviderConfig = v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfig) Validate() error {
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

type UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig struct {
	// The list of allowed JWT issuers.
	//
	// > The list cannot contain more than 200 entries.
	//
	// 	Notice: To clear the issuer list, pass an empty list or an empty string.
	AllowedTokenIssuers []*string `json:"AllowedTokenIssuers,omitempty" xml:"AllowedTokenIssuers,omitempty" type:"Repeated"`
	// Specifies whether to enable the JWT derived short token feature.
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
	// Specifies whether to enable JWT expiration cleanup.
	//
	// example:
	//
	// true
	ExpirationCleanupEnabled *bool `json:"ExpirationCleanupEnabled,omitempty" xml:"ExpirationCleanupEnabled,omitempty"`
}

func (s UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetAllowedTokenIssuers() []*string {
	return s.AllowedTokenIssuers
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetDerivedShortTokenEnabled() *bool {
	return s.DerivedShortTokenEnabled
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetExpiration() *int32 {
	return s.Expiration
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) GetExpirationCleanupEnabled() *bool {
	return s.ExpirationCleanupEnabled
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetAllowedTokenIssuers(v []*string) *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.AllowedTokenIssuers = v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetDerivedShortTokenEnabled(v bool) *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.DerivedShortTokenEnabled = &v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetExpiration(v int32) *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.Expiration = &v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) SetExpirationCleanupEnabled(v bool) *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig {
	s.ExpirationCleanupEnabled = &v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig struct {
	// The client_secret in the OAuth protocol, which is the client secret.
	//
	// > The value cannot exceed 1024 characters in length.
	//
	// example:
	//
	// client_secret_example_xxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The scope in the OAuth protocol, which specifies the permission scope.
	//
	// > The Scope configuration at the credential provider serves as the default value. If the scope parameter is not specified when calling the DeveloperAPI to obtain an OAuth Access Token, the Scope configuration at the credential provider is used for issuance.
	//
	// 	Notice: Separate multiple Scope values with spaces. To clear the Scope configuration, pass an empty string.
	//
	// Restrictions on a single Scope value:
	//
	// 1. Allowed characters: lowercase letters, digits, and special characters `|/:_-.`
	//
	// 2. Must contain at least one lowercase letter or digit.
	//
	// 3. Must start with a special character `.`, a lowercase letter, or a digit.
	//
	// 4. Cannot exceed 1024 characters in length.
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The token endpoint of the OAuth protocol.
	//
	// > The value must start with `http://` or `https://` and cannot exceed 1024 characters in length.
	//
	// example:
	//
	// https://example.com/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetScope() *string {
	return s.Scope
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetClientSecret(v string) *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.ClientSecret = &v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetScope(v string) *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.Scope = &v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetTokenEndpoint(v string) *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *UpdateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) Validate() error {
	return dara.Validate(s)
}
