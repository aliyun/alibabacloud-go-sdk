// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviders(v []*ListCredentialProvidersResponseBodyCredentialProviders) *ListCredentialProvidersResponseBody
	GetCredentialProviders() []*ListCredentialProvidersResponseBodyCredentialProviders
	SetMaxResults(v int32) *ListCredentialProvidersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCredentialProvidersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCredentialProvidersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCredentialProvidersResponseBody
	GetTotalCount() *int32
}

type ListCredentialProvidersResponseBody struct {
	// The list of credential providers.
	CredentialProviders []*ListCredentialProvidersResponseBodyCredentialProviders `json:"CredentialProviders,omitempty" xml:"CredentialProviders,omitempty" type:"Repeated"`
	// The maximum number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCredentialProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersResponseBody) GetCredentialProviders() []*ListCredentialProvidersResponseBodyCredentialProviders {
	return s.CredentialProviders
}

func (s *ListCredentialProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCredentialProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCredentialProvidersResponseBody) SetCredentialProviders(v []*ListCredentialProvidersResponseBodyCredentialProviders) *ListCredentialProvidersResponseBody {
	s.CredentialProviders = v
	return s
}

func (s *ListCredentialProvidersResponseBody) SetMaxResults(v int32) *ListCredentialProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialProvidersResponseBody) SetNextToken(v string) *ListCredentialProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCredentialProvidersResponseBody) SetRequestId(v string) *ListCredentialProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialProvidersResponseBody) SetTotalCount(v int32) *ListCredentialProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCredentialProvidersResponseBody) Validate() error {
	if s.CredentialProviders != nil {
		for _, item := range s.CredentialProviders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCredentialProvidersResponseBodyCredentialProviders struct {
	// The time when the credential provider was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The configuration of the credential provider.
	CredentialProviderConfig *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// The creation type of the credential provider. Valid values:
	//
	// example:
	//
	// user_custom
	CredentialProviderCreationType *string `json:"CredentialProviderCreationType,omitempty" xml:"CredentialProviderCreationType,omitempty"`
	// The credential provider ID.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// The business identifier of the credential provider.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"CredentialProviderIdentifier,omitempty" xml:"CredentialProviderIdentifier,omitempty"`
	// The name of the credential provider.
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// The credential provider type. Valid values:
	//
	// example:
	//
	// oauth
	CredentialProviderType *string `json:"CredentialProviderType,omitempty" xml:"CredentialProviderType,omitempty"`
	// The description of the credential provider.
	//
	// example:
	//
	// This is an example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The credential provider status. Valid values:
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the credential provider was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCredentialProvidersResponseBodyCredentialProviders) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersResponseBodyCredentialProviders) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCredentialProviderConfig() *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig {
	return s.CredentialProviderConfig
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCredentialProviderCreationType() *string {
	return s.CredentialProviderCreationType
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCredentialProviderName() *string {
	return s.CredentialProviderName
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetDescription() *string {
	return s.Description
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetStatus() *string {
	return s.Status
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCreateTime(v int64) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CreateTime = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCredentialProviderConfig(v *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CredentialProviderConfig = v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCredentialProviderCreationType(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CredentialProviderCreationType = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCredentialProviderId(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CredentialProviderId = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCredentialProviderIdentifier(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCredentialProviderName(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CredentialProviderName = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetCredentialProviderType(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.CredentialProviderType = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetDescription(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.Description = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetInstanceId(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.InstanceId = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetStatus(v string) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.Status = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) SetUpdateTime(v int64) *ListCredentialProvidersResponseBodyCredentialProviders {
	s.UpdateTime = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProviders) Validate() error {
	if s.CredentialProviderConfig != nil {
		if err := s.CredentialProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig struct {
	// The configuration of the JWT-type credential provider.
	JwtProviderConfig *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// The configuration of the OAuth-type credential provider.
	OAuthProviderConfig *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig `json:"OAuthProviderConfig,omitempty" xml:"OAuthProviderConfig,omitempty" type:"Struct"`
	// The list of credential IDs that correspond to the sensitive configurations of the credential provider.
	ProviderCredentialIds []*string `json:"ProviderCredentialIds,omitempty" xml:"ProviderCredentialIds,omitempty" type:"Repeated"`
}

func (s ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) GetJwtProviderConfig() *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	return s.JwtProviderConfig
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) GetOAuthProviderConfig() *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	return s.OAuthProviderConfig
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) GetProviderCredentialIds() []*string {
	return s.ProviderCredentialIds
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) SetJwtProviderConfig(v *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig {
	s.JwtProviderConfig = v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) SetOAuthProviderConfig(v *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig {
	s.OAuthProviderConfig = v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) SetProviderCredentialIds(v []*string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig {
	s.ProviderCredentialIds = v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfig) Validate() error {
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

type ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig struct {
	// The list of allowed JWT issuers.
	AllowedTokenIssuers []*string `json:"AllowedTokenIssuers,omitempty" xml:"AllowedTokenIssuers,omitempty" type:"Repeated"`
	// Specifies whether to enable the JWT-derived short token capability.
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
	// JWT issuer。
	//
	// example:
	//
	// https://test.issuer.com
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The JWKs endpoint URL.
	//
	// example:
	//
	// https://example123456.aliyunidaas.com/api/v2/auths_ngz2wj35ixxxdyat55nexxxxxx/oauth2/jwks
	JwksEndpoint *string `json:"JwksEndpoint,omitempty" xml:"JwksEndpoint,omitempty"`
}

func (s ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GetAllowedTokenIssuers() []*string {
	return s.AllowedTokenIssuers
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GetDerivedShortTokenEnabled() *bool {
	return s.DerivedShortTokenEnabled
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GetExpiration() *int32 {
	return s.Expiration
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GetExpirationCleanupEnabled() *bool {
	return s.ExpirationCleanupEnabled
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) GetJwksEndpoint() *string {
	return s.JwksEndpoint
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) SetAllowedTokenIssuers(v []*string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	s.AllowedTokenIssuers = v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) SetDerivedShortTokenEnabled(v bool) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	s.DerivedShortTokenEnabled = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) SetExpiration(v int32) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	s.Expiration = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) SetExpirationCleanupEnabled(v bool) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	s.ExpirationCleanupEnabled = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) SetIssuer(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	s.Issuer = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) SetJwksEndpoint(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig {
	s.JwksEndpoint = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigJwtProviderConfig) Validate() error {
	return dara.Validate(s)
}

type ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig struct {
	// The endpoint URL used to guide users through authorization. Conditionally required: this parameter is required when AuthorizationFlow is set to user_federation and ProviderVendor is set to custom. For preset vendors, this value can be automatically populated through DiscoveryUrl.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The OAuth authorization flow type. Valid values:
	//
	// - m2m: machine-to-machine (2LO, Client Credentials).
	//
	// - user_federation: user federation (3LO, Authorization Code).
	AuthorizationFlow *string `json:"AuthorizationFlow,omitempty" xml:"AuthorizationFlow,omitempty"`
	// The client_id in the OAuth protocol.
	//
	// example:
	//
	// client_id_example_xxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The URL of the discovery document used to automatically obtain OAuth endpoint configurations. Conditionally optional: used when AuthorizationFlow is set to user_federation. If DiscoveryUrl is not provided, you must manually configure fields such as TokenEndpoint and AuthorizationEndpoint.
	DiscoveryUrl *string `json:"DiscoveryUrl,omitempty" xml:"DiscoveryUrl,omitempty"`
	Issuer       *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The method used to generate the PKCE code_challenge. Default value: s256.
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Specifies whether to use the PKCE extension to enhance security. We recommend that you always enable this feature.
	PkceEnabled *bool `json:"PkceEnabled,omitempty" xml:"PkceEnabled,omitempty"`
	// The preset vendor or custom configuration. Optional. Default value: custom.
	ProviderVendor *string `json:"ProviderVendor,omitempty" xml:"ProviderVendor,omitempty"`
	// The scope in the OAuth protocol, which specifies the permission scope.
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The redirect URI automatically generated by the system when the credential provider is created. Configure this value as the redirect_uri in the OAuth provider.
	SystemRedirectUri *string `json:"SystemRedirectUri,omitempty" xml:"SystemRedirectUri,omitempty"`
	// The token endpoint of the OAuth protocol.
	//
	// example:
	//
	// https://example.com/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetAuthorizationFlow() *string {
	return s.AuthorizationFlow
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetClientId() *string {
	return s.ClientId
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetDiscoveryUrl() *string {
	return s.DiscoveryUrl
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetPkceChallengeMethod() *string {
	return s.PkceChallengeMethod
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetPkceEnabled() *bool {
	return s.PkceEnabled
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetProviderVendor() *string {
	return s.ProviderVendor
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetScope() *string {
	return s.Scope
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetSystemRedirectUri() *string {
	return s.SystemRedirectUri
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetAuthorizationEndpoint(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetAuthorizationFlow(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.AuthorizationFlow = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetClientId(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.ClientId = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetDiscoveryUrl(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.DiscoveryUrl = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetIssuer(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.Issuer = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetPkceChallengeMethod(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.PkceChallengeMethod = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetPkceEnabled(v bool) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.PkceEnabled = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetProviderVendor(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.ProviderVendor = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetScope(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.Scope = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetSystemRedirectUri(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.SystemRedirectUri = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) SetTokenEndpoint(v string) *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *ListCredentialProvidersResponseBodyCredentialProvidersCredentialProviderConfigOAuthProviderConfig) Validate() error {
	return dara.Validate(s)
}
