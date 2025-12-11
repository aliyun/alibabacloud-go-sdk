// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *CreateSourceRequest
	GetGatewayId() *string
	SetK8sSourceConfig(v *CreateSourceRequestK8sSourceConfig) *CreateSourceRequest
	GetK8sSourceConfig() *CreateSourceRequestK8sSourceConfig
	SetNacosSourceConfig(v *CreateSourceRequestNacosSourceConfig) *CreateSourceRequest
	GetNacosSourceConfig() *CreateSourceRequestNacosSourceConfig
	SetResourceGroupId(v string) *CreateSourceRequest
	GetResourceGroupId() *string
	SetType(v string) *CreateSourceRequest
	GetType() *string
}

type CreateSourceRequest struct {
	// example:
	//
	// gw-cq7l5s5lhtgi6q***
	GatewayId         *string                               `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	K8sSourceConfig   *CreateSourceRequestK8sSourceConfig   `json:"k8sSourceConfig,omitempty" xml:"k8sSourceConfig,omitempty" type:"Struct"`
	NacosSourceConfig *CreateSourceRequestNacosSourceConfig `json:"nacosSourceConfig,omitempty" xml:"nacosSourceConfig,omitempty" type:"Struct"`
	ResourceGroupId   *string                               `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// MSE_NACOS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateSourceRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateSourceRequest) GetK8sSourceConfig() *CreateSourceRequestK8sSourceConfig {
	return s.K8sSourceConfig
}

func (s *CreateSourceRequest) GetNacosSourceConfig() *CreateSourceRequestNacosSourceConfig {
	return s.NacosSourceConfig
}

func (s *CreateSourceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateSourceRequest) SetGatewayId(v string) *CreateSourceRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateSourceRequest) SetK8sSourceConfig(v *CreateSourceRequestK8sSourceConfig) *CreateSourceRequest {
	s.K8sSourceConfig = v
	return s
}

func (s *CreateSourceRequest) SetNacosSourceConfig(v *CreateSourceRequestNacosSourceConfig) *CreateSourceRequest {
	s.NacosSourceConfig = v
	return s
}

func (s *CreateSourceRequest) SetResourceGroupId(v string) *CreateSourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSourceRequest) SetType(v string) *CreateSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateSourceRequest) Validate() error {
	if s.K8sSourceConfig != nil {
		if err := s.K8sSourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NacosSourceConfig != nil {
		if err := s.NacosSourceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSourceRequestK8sSourceConfig struct {
	AuthorizeSecurityGroupRules []*CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules `json:"authorizeSecurityGroupRules,omitempty" xml:"authorizeSecurityGroupRules,omitempty" type:"Repeated"`
	// example:
	//
	// c3fbe6caaaece4062b*****
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s CreateSourceRequestK8sSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceRequestK8sSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateSourceRequestK8sSourceConfig) GetAuthorizeSecurityGroupRules() []*CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules {
	return s.AuthorizeSecurityGroupRules
}

func (s *CreateSourceRequestK8sSourceConfig) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateSourceRequestK8sSourceConfig) SetAuthorizeSecurityGroupRules(v []*CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) *CreateSourceRequestK8sSourceConfig {
	s.AuthorizeSecurityGroupRules = v
	return s
}

func (s *CreateSourceRequestK8sSourceConfig) SetClusterId(v string) *CreateSourceRequestK8sSourceConfig {
	s.ClusterId = &v
	return s
}

func (s *CreateSourceRequestK8sSourceConfig) Validate() error {
	if s.AuthorizeSecurityGroupRules != nil {
		for _, item := range s.AuthorizeSecurityGroupRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules struct {
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	PortRanges  []*string `json:"portRanges,omitempty" xml:"portRanges,omitempty" type:"Repeated"`
	// example:
	//
	// sg-bp14w4fa4j***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) GoString() string {
	return s.String()
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) GetDescription() *string {
	return s.Description
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) GetPortRanges() []*string {
	return s.PortRanges
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) SetDescription(v string) *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules {
	s.Description = &v
	return s
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) SetPortRanges(v []*string) *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules {
	s.PortRanges = v
	return s
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) SetSecurityGroupId(v string) *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateSourceRequestK8sSourceConfigAuthorizeSecurityGroupRules) Validate() error {
	return dara.Validate(s)
}

type CreateSourceRequestNacosSourceConfig struct {
	// example:
	//
	// mse-cn-0dw3w***
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateSourceRequestNacosSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceRequestNacosSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateSourceRequestNacosSourceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSourceRequestNacosSourceConfig) SetInstanceId(v string) *CreateSourceRequestNacosSourceConfig {
	s.InstanceId = &v
	return s
}

func (s *CreateSourceRequestNacosSourceConfig) Validate() error {
	return dara.Validate(s)
}
