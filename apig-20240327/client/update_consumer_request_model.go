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
	SetDescription(v string) *UpdateConsumerRequest
	GetDescription() *string
	SetEnable(v bool) *UpdateConsumerRequest
	GetEnable() *bool
	SetJwtIdentityConfig(v *JwtIdentityConfig) *UpdateConsumerRequest
	GetJwtIdentityConfig() *JwtIdentityConfig
}

type UpdateConsumerRequest struct {
	AkSkIdentityConfigs  []*AkSkIdentityConfig `json:"akSkIdentityConfigs,omitempty" xml:"akSkIdentityConfigs,omitempty" type:"Repeated"`
	ApikeyIdentityConfig *ApiKeyIdentityConfig `json:"apikeyIdentityConfig,omitempty" xml:"apikeyIdentityConfig,omitempty"`
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Enable            *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	JwtIdentityConfig *JwtIdentityConfig `json:"jwtIdentityConfig,omitempty" xml:"jwtIdentityConfig,omitempty"`
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

func (s *UpdateConsumerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConsumerRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateConsumerRequest) GetJwtIdentityConfig() *JwtIdentityConfig {
	return s.JwtIdentityConfig
}

func (s *UpdateConsumerRequest) SetAkSkIdentityConfigs(v []*AkSkIdentityConfig) *UpdateConsumerRequest {
	s.AkSkIdentityConfigs = v
	return s
}

func (s *UpdateConsumerRequest) SetApikeyIdentityConfig(v *ApiKeyIdentityConfig) *UpdateConsumerRequest {
	s.ApikeyIdentityConfig = v
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

func (s *UpdateConsumerRequest) SetJwtIdentityConfig(v *JwtIdentityConfig) *UpdateConsumerRequest {
	s.JwtIdentityConfig = v
	return s
}

func (s *UpdateConsumerRequest) Validate() error {
	return dara.Validate(s)
}
