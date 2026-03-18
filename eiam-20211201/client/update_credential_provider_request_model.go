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
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 认证令牌提供商的配置。
	CredentialProviderConfig *UpdateCredentialProviderRequestCredentialProviderConfig `json:"CredentialProviderConfig,omitempty" xml:"CredentialProviderConfig,omitempty" type:"Struct"`
	// 认证令牌提供商ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// 认证令牌提供商名称。
	//
	// example:
	//
	// test_example_name
	CredentialProviderName *string `json:"CredentialProviderName,omitempty" xml:"CredentialProviderName,omitempty"`
	// IDaaS EIAM实例的ID。
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
	// JWT身份提供商配置。
	JwtProviderConfig *UpdateCredentialProviderRequestCredentialProviderConfigJwtProviderConfig `json:"JwtProviderConfig,omitempty" xml:"JwtProviderConfig,omitempty" type:"Struct"`
	// OAuth 2LO机用类型的提供商的配置。
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
	// OAuth协议中的client_secret，客户端密钥。
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
