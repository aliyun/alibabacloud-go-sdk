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
	SetGatewayType(v string) *CreateConsumerRequest
	GetGatewayType() *string
	SetJwtIdentityConfig(v *JwtIdentityConfig) *CreateConsumerRequest
	GetJwtIdentityConfig() *JwtIdentityConfig
	SetName(v string) *CreateConsumerRequest
	GetName() *string
}

type CreateConsumerRequest struct {
	AkSkIdentityConfigs  []*AkSkIdentityConfig `json:"akSkIdentityConfigs,omitempty" xml:"akSkIdentityConfigs,omitempty" type:"Repeated"`
	ApikeyIdentityConfig *ApiKeyIdentityConfig `json:"apikeyIdentityConfig,omitempty" xml:"apikeyIdentityConfig,omitempty"`
	// example:
	//
	// consumer for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// API
	GatewayType       *string            `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	JwtIdentityConfig *JwtIdentityConfig `json:"jwtIdentityConfig,omitempty" xml:"jwtIdentityConfig,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *CreateConsumerRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateConsumerRequest) GetJwtIdentityConfig() *JwtIdentityConfig {
	return s.JwtIdentityConfig
}

func (s *CreateConsumerRequest) GetName() *string {
	return s.Name
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

func (s *CreateConsumerRequest) Validate() error {
	return dara.Validate(s)
}
