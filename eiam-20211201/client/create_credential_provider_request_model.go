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
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 认证令牌提供商的配置。
	CredentialProviderConfig *CreateCredentialProviderRequestCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// 认证令牌提供商的业务标识。是一个具备可读性的唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"CredentialProviderIdentifier,omitempty" xml:"CredentialProviderIdentifier,omitempty"`
	// 认证令牌提供商名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// 认证令牌提供商的类型。
	//
	// This parameter is required.
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
	// IDaaS EIAM实例的ID。
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
	// JWT身份提供商配置。
	JwtProviderConfig *CreateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// OAuth 2LO机用类型的提供商的配置。
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
	// OAuth协议中的client_id，客户端ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// client_id_example_xxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// OAuth协议中的client_secret，客户端密钥。
	//
	// This parameter is required.
	//
	// example:
	//
	// client_secret_example_xxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// OAuth协议中的scope，权限范围。
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// OAuth协议的Token端点。
	//
	// This parameter is required.
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

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetClientId() *string {
	return s.ClientId
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetScope() *string {
	return s.Scope
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetClientId(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.ClientId = &v
	return s
}

func (s *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig) SetClientSecret(v string) *CreateCredentialProviderRequestCredentialProviderConfigOAuthProviderConfig {
	s.ClientSecret = &v
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
