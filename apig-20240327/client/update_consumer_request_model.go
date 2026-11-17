// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkSkIdentityConfigs(v []*AkSkIdentityConfig) *UpdateConsumerRequest
	GetAkSkIdentityConfigs() []*AkSkIdentityConfig
	SetApikeyIdentityConfig(v *ApiKeyIdentityConfig) *UpdateConsumerRequest
	GetApikeyIdentityConfig() *ApiKeyIdentityConfig
	SetClearExpireAt(v bool) *UpdateConsumerRequest
	GetClearExpireAt() *bool
	SetDescription(v string) *UpdateConsumerRequest
	GetDescription() *string
	SetEnable(v bool) *UpdateConsumerRequest
	GetEnable() *bool
	SetEnableAfterRenewal(v bool) *UpdateConsumerRequest
	GetEnableAfterRenewal() *bool
	SetExpireAt(v int64) *UpdateConsumerRequest
	GetExpireAt() *int64
	SetJwtIdentityConfig(v *JwtIdentityConfig) *UpdateConsumerRequest
	GetJwtIdentityConfig() *JwtIdentityConfig
	SetRenew(v bool) *UpdateConsumerRequest
	GetRenew() *bool
}

type UpdateConsumerRequest struct {
	// The list of AK/SK identity authentication configurations.
	AkSkIdentityConfigs []*AkSkIdentityConfig `json:"akSkIdentityConfigs,omitempty" xml:"akSkIdentityConfigs,omitempty" type:"Repeated"`
	// The API key identity authentication configuration.
	ApikeyIdentityConfig *ApiKeyIdentityConfig `json:"apikeyIdentityConfig,omitempty" xml:"apikeyIdentityConfig,omitempty"`
	// Specifies whether to clear the expiration time and restore permanent validity. If this parameter is set to true, expireAt cannot be specified at the same time. The default value is false, which indicates that the expiration time is not cleared.
	//
	// example:
	//
	// false
	ClearExpireAt *bool `json:"clearExpireAt,omitempty" xml:"clearExpireAt,omitempty"`
	// The description.
	//
	// example:
	//
	// Test consumer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable the consumer.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Specifies whether to enable the consumer after renewal. This parameter can be set to true only when renew=true. The default value is false, which retains the current enabled or disabled state and does not implicitly enable a manually disabled consumer.
	//
	// example:
	//
	// false
	EnableAfterRenewal *bool `json:"enableAfterRenewal,omitempty" xml:"enableAfterRenewal,omitempty"`
	// The expiration time of the consumer, in UTC UNIX seconds. It must be a positive integer and later than the current time. If this parameter is not specified, the original value is retained. This parameter cannot be set at the same time as clearExpireAt=true. If renew=true, the value must be later than the original unexpired expiration time. If the configuration is saved but fails to be delivered, a retry with the original expiration time is allowed, but shortening the expiration time is still prohibited.
	//
	// example:
	//
	// 1893456000
	ExpireAt *int64 `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The JWT identity authentication configuration.
	JwtIdentityConfig *JwtIdentityConfig `json:"jwtIdentityConfig,omitempty" xml:"jwtIdentityConfig,omitempty"`
	// Specifies whether to only renew the consumer. If this parameter is set to true, only the validity period is updated and the consumer is explicitly enabled at the same time. The latest credentials, description, and authorizations on the server are retained. You must set expireAt or clearExpireAt=true, and you cannot specify identity configurations or description. The default value is false, which retains the original consumer update semantics.
	//
	// example:
	//
	// false
	Renew *bool `json:"renew,omitempty" xml:"renew,omitempty"`
}

func (s UpdateConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerRequest) GetAkSkIdentityConfigs() []*AkSkIdentityConfig {
	return s.AkSkIdentityConfigs
}

func (s *UpdateConsumerRequest) GetApikeyIdentityConfig() *ApiKeyIdentityConfig {
	return s.ApikeyIdentityConfig
}

func (s *UpdateConsumerRequest) GetClearExpireAt() *bool {
	return s.ClearExpireAt
}

func (s *UpdateConsumerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConsumerRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateConsumerRequest) GetEnableAfterRenewal() *bool {
	return s.EnableAfterRenewal
}

func (s *UpdateConsumerRequest) GetExpireAt() *int64 {
	return s.ExpireAt
}

func (s *UpdateConsumerRequest) GetJwtIdentityConfig() *JwtIdentityConfig {
	return s.JwtIdentityConfig
}

func (s *UpdateConsumerRequest) GetRenew() *bool {
	return s.Renew
}

func (s *UpdateConsumerRequest) SetAkSkIdentityConfigs(v []*AkSkIdentityConfig) *UpdateConsumerRequest {
	s.AkSkIdentityConfigs = v
	return s
}

func (s *UpdateConsumerRequest) SetApikeyIdentityConfig(v *ApiKeyIdentityConfig) *UpdateConsumerRequest {
	s.ApikeyIdentityConfig = v
	return s
}

func (s *UpdateConsumerRequest) SetClearExpireAt(v bool) *UpdateConsumerRequest {
	s.ClearExpireAt = &v
	return s
}

func (s *UpdateConsumerRequest) SetDescription(v string) *UpdateConsumerRequest {
	s.Description = &v
	return s
}

func (s *UpdateConsumerRequest) SetEnable(v bool) *UpdateConsumerRequest {
	s.Enable = &v
	return s
}

func (s *UpdateConsumerRequest) SetEnableAfterRenewal(v bool) *UpdateConsumerRequest {
	s.EnableAfterRenewal = &v
	return s
}

func (s *UpdateConsumerRequest) SetExpireAt(v int64) *UpdateConsumerRequest {
	s.ExpireAt = &v
	return s
}

func (s *UpdateConsumerRequest) SetJwtIdentityConfig(v *JwtIdentityConfig) *UpdateConsumerRequest {
	s.JwtIdentityConfig = v
	return s
}

func (s *UpdateConsumerRequest) SetRenew(v bool) *UpdateConsumerRequest {
	s.Renew = &v
	return s
}

func (s *UpdateConsumerRequest) Validate() error {
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
