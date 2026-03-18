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
	CredentialProvider *GetCredentialProviderResponseBodyCredentialProvider `json:"CredentialProvider,omitempty" xml:"CredentialProvider,omitempty" type:"Struct"`
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
	// 认证令牌提供商的创建时间，Unix时间戳。
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 认证令牌提供商的配置。
	CredentialProviderConfig *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// 认证令牌提供商的创建类型。
	//
	// example:
	//
	// user_custom
	CredentialProviderCreationType *string `json:"CredentialProviderCreationType,omitempty" xml:"CredentialProviderCreationType,omitempty"`
	// 认证令牌提供商ID。
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// 认证令牌提供商的业务标识。
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"CredentialProviderIdentifier,omitempty" xml:"CredentialProviderIdentifier,omitempty"`
	// 认证令牌提供商名称。
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// 认证令牌提供商的类型。
	//
	// example:
	//
	// oauth
	CredentialProviderType *string `json:"CredentialProviderType,omitempty" xml:"CredentialProviderType,omitempty"`
	// 描述。
	//
	// example:
	//
	// This is an example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// EIAM实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 认证令牌提供商的状态。
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 认证令牌提供商的更新时间，Unix时间戳。
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
	// JWT身份提供商配置。
	JwtProviderConfig *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// OAuth 2LO机用类型的提供商的配置。
	OAuthProviderConfig *GetCredentialProviderResponseBodyCredentialProviderCredentialProviderConfigOAuthProviderConfig `json:"OAuthProviderConfig,omitempty" xml:"OAuthProviderConfig,omitempty" type:"Struct"`
	// 认证令牌提供商的敏感配置对应的凭据ID列表。
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
	// 签发出的JWT中的issuer字段的允许列表。
	AllowedTokenIssuers []*string `json:"AllowedTokenIssuers,omitempty" xml:"AllowedTokenIssuers,omitempty" type:"Repeated"`
	// 是否开启JWT派生短令牌能力。
	//
	// example:
	//
	// false
	DerivedShortTokenEnabled *bool `json:"DerivedShortTokenEnabled,omitempty" xml:"DerivedShortTokenEnabled,omitempty"`
	// JWT的有效时长，单位秒。
	//
	// example:
	//
	// 900
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 是否开启JWT过期清理。
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
	// JWKs端点地址。
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
	// OAuth协议中的client_id，客户端ID。
	//
	// example:
	//
	// client_id_example_xxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// OAuth协议中的scope，权限范围。
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// OAuth协议的Token端点。
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
