// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProvider(v *GetCredentialProviderResponseBodyCredentialProvider) *GetCredentialProviderResponseBody
	GetCredentialProvider() *GetCredentialProviderResponseBodyCredentialProvider
	SetRequestId(v string) *GetCredentialProviderResponseBody
	GetRequestId() *string
}

type GetCredentialProviderResponseBody struct {
	// Credential provider.
	CredentialProvider *GetCredentialProviderResponseBodyCredentialProvider `json:"CredentialProvider,omitempty" xml:"CredentialProvider,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderResponseBody) GetCredentialProvider() *GetCredentialProviderResponseBodyCredentialProvider {
	return s.CredentialProvider
}

func (s *GetCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialProviderResponseBody) SetCredentialProvider(v *GetCredentialProviderResponseBodyCredentialProvider) *GetCredentialProviderResponseBody {
	s.CredentialProvider = v
	return s
}

func (s *GetCredentialProviderResponseBody) SetRequestId(v string) *GetCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialProviderResponseBody) Validate() error {
	if s.CredentialProvider != nil {
		if err := s.CredentialProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialProviderResponseBodyCredentialProvider struct {
	// Creation time of the credential provider, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Credential provider configuration.
	CredentialProviderConfig *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// Credential provider creation type. Valid values:
	//
	// - system_init: Created by the system
	//
	// - user_custom: Created by a user
	//
	// example:
	//
	// user_custom
	CredentialProviderCreationType *string `json:"CredentialProviderCreationType,omitempty" xml:"CredentialProviderCreationType,omitempty"`
	// Credential provider ID.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// Credential provider identifier.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"CredentialProviderIdentifier,omitempty" xml:"CredentialProviderIdentifier,omitempty"`
	// Credential provider name.
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// Credential provider type. Valid values:
	//
	// - oauth: OAuth credential provider
	//
	// - jwt: JWT credential provider
	//
	// example:
	//
	// oauth
	CredentialProviderType *string `json:"CredentialProviderType,omitempty" xml:"CredentialProviderType,omitempty"`
	// Description.
	//
	// example:
	//
	// This is an example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Credential provider status. Valid values:
	//
	// - enabled: Enabled
	//
	// - disabled: Disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Update time of the credential provider, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCredentialProviderResponseBodyCredentialProvider) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderResponseBodyCredentialProvider) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCredentialProviderConfig() *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig {
	return s.CredentialProviderConfig
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCredentialProviderCreationType() *string {
	return s.CredentialProviderCreationType
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCredentialProviderName() *string {
	return s.CredentialProviderName
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetStatus() *string {
	return s.Status
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCreateTime(v int64) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CreateTime = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCredentialProviderConfig(v *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CredentialProviderConfig = v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCredentialProviderCreationType(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CredentialProviderCreationType = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCredentialProviderId(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CredentialProviderId = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCredentialProviderIdentifier(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCredentialProviderName(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CredentialProviderName = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetCredentialProviderType(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.CredentialProviderType = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetDescription(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.Description = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetInstanceId(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.InstanceId = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetStatus(v string) *GetCredentialProviderResponseBodyCredentialProvider {
	s.Status = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) SetUpdateTime(v int64) *GetCredentialProviderResponseBodyCredentialProvider {
	s.UpdateTime = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProvider) Validate() error {
	if s.CredentialProviderConfig != nil {
		if err := s.CredentialProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig struct {
	// Configuration for a JWT credential provider.
	JwtProviderConfig *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// Configuration for an OAuth credential provider.
	OAuthProviderConfig *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig `json:"OAuthProviderConfig,omitempty" xml:"OAuthProviderConfig,omitempty" type:"Struct"`
	// List of credential IDs for sensitive configurations of the credential provider.
	//
	// > The system securely stores sensitive configuration information as credentials.
	ProviderCredentialIds []*string `json:"ProviderCredentialIds,omitempty" xml:"ProviderCredentialIds,omitempty" type:"Repeated"`
}

func (s GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) GetJwtProviderConfig() *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	return s.JwtProviderConfig
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) GetOAuthProviderConfig() *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig {
	return s.OAuthProviderConfig
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) GetProviderCredentialIds() []*string {
	return s.ProviderCredentialIds
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) SetJwtProviderConfig(v *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig {
	s.JwtProviderConfig = v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) SetOAuthProviderConfig(v *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig {
	s.OAuthProviderConfig = v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) SetProviderCredentialIds(v []*string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig {
	s.ProviderCredentialIds = v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig) Validate() error {
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

type GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig struct {
	// List of allowed JWT issuers.
	AllowedTokenIssuers []*string `json:"AllowedTokenIssuers,omitempty" xml:"AllowedTokenIssuers,omitempty" type:"Repeated"`
	// Enable JWT derived short token.
	//
	// example:
	//
	// false
	DerivedShortTokenEnabled *bool `json:"DerivedShortTokenEnabled,omitempty" xml:"DerivedShortTokenEnabled,omitempty"`
	// Validity period of the JWT. Unit: seconds.
	//
	// example:
	//
	// 900
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// Enable JWT expiration cleanup.
	//
	// example:
	//
	// true
	ExpirationCleanupEnabled *bool `json:"ExpirationCleanupEnabled,omitempty" xml:"ExpirationCleanupEnabled,omitempty"`
	// JWT issuer.
	//
	// example:
	//
	// https://test.issuer.com
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// JWKs endpoint URL.
	//
	// example:
	//
	// https://example123456.aliyunidaas.com/api/v2/auths_ngz2wj35ixxxdyat55nexxxxxx/oauth2/jwks
	JwksEndpoint *string `json:"JwksEndpoint,omitempty" xml:"JwksEndpoint,omitempty"`
}

func (s GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GetAllowedTokenIssuers() []*string {
	return s.AllowedTokenIssuers
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GetDerivedShortTokenEnabled() *bool {
	return s.DerivedShortTokenEnabled
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GetExpiration() *int32 {
	return s.Expiration
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GetExpirationCleanupEnabled() *bool {
	return s.ExpirationCleanupEnabled
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) GetJwksEndpoint() *string {
	return s.JwksEndpoint
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) SetAllowedTokenIssuers(v []*string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	s.AllowedTokenIssuers = v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) SetDerivedShortTokenEnabled(v bool) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	s.DerivedShortTokenEnabled = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) SetExpiration(v int32) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	s.Expiration = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) SetExpirationCleanupEnabled(v bool) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	s.ExpirationCleanupEnabled = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) SetIssuer(v string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	s.Issuer = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) SetJwksEndpoint(v string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig {
	s.JwksEndpoint = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig) Validate() error {
	return dara.Validate(s)
}

type GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig struct {
	// Client ID, corresponding to client_id in the OAuth protocol.
	//
	// example:
	//
	// client_id_example_xxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Scope, corresponding to scope in the OAuth protocol.
	//
	// > The Scope value configured for the OAuth credential provider serves as the default. If you do not specify the scope parameter when calling the Developer API to obtain an OAuth access token, the system uses this default Scope value.
	//
	// 	Notice:
	//
	// Separate multiple Scope values with spaces.
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Token endpoint, corresponding to the OAuth protocol.
	//
	// example:
	//
	// https://example.com/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) GetClientId() *string {
	return s.ClientId
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) GetScope() *string {
	return s.Scope
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) SetClientId(v string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig {
	s.ClientId = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) SetScope(v string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig {
	s.Scope = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) SetTokenEndpoint(v string) *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig) Validate() error {
	return dara.Validate(s)
}
