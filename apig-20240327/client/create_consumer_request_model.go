// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkSkIdentityConfigs(v []*AkSkIdentityConfig) *CreateConsumerRequest
	GetAkSkIdentityConfigs() []*AkSkIdentityConfig
	SetApikeyIdentityConfig(v *ApiKeyIdentityConfig) *CreateConsumerRequest
	GetApikeyIdentityConfig() *ApiKeyIdentityConfig
	SetDescription(v string) *CreateConsumerRequest
	GetDescription() *string
	SetEnable(v bool) *CreateConsumerRequest
	GetEnable() *bool
	SetExpireAt(v int64) *CreateConsumerRequest
	GetExpireAt() *int64
	SetGatewayType(v string) *CreateConsumerRequest
	GetGatewayType() *string
	SetJwtIdentityConfig(v *JwtIdentityConfig) *CreateConsumerRequest
	GetJwtIdentityConfig() *JwtIdentityConfig
	SetName(v string) *CreateConsumerRequest
	GetName() *string
	SetClientToken(v string) *CreateConsumerRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateConsumerRequest
	GetDryRun() *bool
}

type CreateConsumerRequest struct {
	// The list of AK/SK identity configurations.
	AkSkIdentityConfigs []*AkSkIdentityConfig `json:"akSkIdentityConfigs,omitempty" xml:"akSkIdentityConfigs,omitempty" type:"Repeated"`
	// The identity configuration for API key authentication.
	ApikeyIdentityConfig *ApiKeyIdentityConfig `json:"apikeyIdentityConfig,omitempty" xml:"apikeyIdentityConfig,omitempty"`
	// The description of the consumer.
	//
	// example:
	//
	// consumer for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable the consumer.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The expiration time of the consumer. The value must be a positive integer that represents a UNIX timestamp in seconds and is later than the current time. If this parameter is left empty, the consumer never expires. The API key, JWT, and HMAC of the same consumer share this expiration time. Milliseconds are not automatically converted to seconds.
	//
	// example:
	//
	// 1893456000
	ExpireAt *int64 `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The gateway type. Valid values:
	//
	// - API: the consumer of a cloud-native API gateway.
	//
	// - AI: the consumer of an AI gateway.
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The JWT identity configuration.
	JwtIdentityConfig *JwtIdentityConfig `json:"jwtIdentityConfig,omitempty" xml:"jwtIdentityConfig,omitempty"`
	// The name of the consumer.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The idempotence identity. You can invoke the API with a globally unique value (a UUID is recommended) that does not exceed 64 characters in length. If you invoke a duplicate request with the same ClientToken and identical request parameters within 24 hours after the first request succeeds, the system directly returns the consumer ID created by the first request without creating a new consumer. If you invoke a request with the same ClientToken but different request parameters, the system returns IdempotentParameterMismatch. If the first request is still being processed, the system returns IdempotentProcessing. If this parameter is left empty, idempotence control is disabled, and the behavior is the same as that of the existing version.
	//
	// example:
	//
	// 5f7a2c1e-9b3d-4e8f-a1c6-0d2b8e4f7a13
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to perform only a dry run. If you set this parameter to true, the system performs the same parameter, permission, IAM, resource policy, resource existence, and business rule checks as a real creation request, but does not create the consumer, write data to the database, create or delete a SecretRef, or generate audit logs. If this parameter is left empty or set to false, the behavior is the same as that of the previous version.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerRequest) GetAkSkIdentityConfigs() []*AkSkIdentityConfig {
	return s.AkSkIdentityConfigs
}

func (s *CreateConsumerRequest) GetApikeyIdentityConfig() *ApiKeyIdentityConfig {
	return s.ApikeyIdentityConfig
}

func (s *CreateConsumerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConsumerRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateConsumerRequest) GetExpireAt() *int64 {
	return s.ExpireAt
}

func (s *CreateConsumerRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateConsumerRequest) GetJwtIdentityConfig() *JwtIdentityConfig {
	return s.JwtIdentityConfig
}

func (s *CreateConsumerRequest) GetName() *string {
	return s.Name
}

func (s *CreateConsumerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateConsumerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateConsumerRequest) SetAkSkIdentityConfigs(v []*AkSkIdentityConfig) *CreateConsumerRequest {
	s.AkSkIdentityConfigs = v
	return s
}

func (s *CreateConsumerRequest) SetApikeyIdentityConfig(v *ApiKeyIdentityConfig) *CreateConsumerRequest {
	s.ApikeyIdentityConfig = v
	return s
}

func (s *CreateConsumerRequest) SetDescription(v string) *CreateConsumerRequest {
	s.Description = &v
	return s
}

func (s *CreateConsumerRequest) SetEnable(v bool) *CreateConsumerRequest {
	s.Enable = &v
	return s
}

func (s *CreateConsumerRequest) SetExpireAt(v int64) *CreateConsumerRequest {
	s.ExpireAt = &v
	return s
}

func (s *CreateConsumerRequest) SetGatewayType(v string) *CreateConsumerRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateConsumerRequest) SetJwtIdentityConfig(v *JwtIdentityConfig) *CreateConsumerRequest {
	s.JwtIdentityConfig = v
	return s
}

func (s *CreateConsumerRequest) SetName(v string) *CreateConsumerRequest {
	s.Name = &v
	return s
}

func (s *CreateConsumerRequest) SetClientToken(v string) *CreateConsumerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConsumerRequest) SetDryRun(v bool) *CreateConsumerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateConsumerRequest) Validate() error {
	if s.AkSkIdentityConfigs != nil {
		for _, item := range s.AkSkIdentityConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ApikeyIdentityConfig != nil {
		if err := s.ApikeyIdentityConfig.Validate(); err != nil {
			return err
		}
	}
	if s.JwtIdentityConfig != nil {
		if err := s.JwtIdentityConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
