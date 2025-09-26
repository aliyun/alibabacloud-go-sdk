// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLLMDeployConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDeploy(v bool) *LLMDeployConfig
	GetAutoDeploy() *bool
	SetBackendScene(v string) *LLMDeployConfig
	GetBackendScene() *string
	SetCustomDomainIds(v []*string) *LLMDeployConfig
	GetCustomDomainIds() []*string
	SetGatewayType(v string) *LLMDeployConfig
	GetGatewayType() *string
	SetPolicyConfigs(v []*PolicyConfig) *LLMDeployConfig
	GetPolicyConfigs() []*PolicyConfig
	SetServiceConfigs(v []*TargetServiceConfig) *LLMDeployConfig
	GetServiceConfigs() []*TargetServiceConfig
}

type LLMDeployConfig struct {
	AutoDeploy      *bool                  `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	BackendScene    *string                `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	CustomDomainIds []*string              `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	GatewayType     *string                `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	PolicyConfigs   []*PolicyConfig        `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
	ServiceConfigs  []*TargetServiceConfig `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s LLMDeployConfig) String() string {
	return dara.Prettify(s)
}

func (s LLMDeployConfig) GoString() string {
	return s.String()
}

func (s *LLMDeployConfig) GetAutoDeploy() *bool {
	return s.AutoDeploy
}

func (s *LLMDeployConfig) GetBackendScene() *string {
	return s.BackendScene
}

func (s *LLMDeployConfig) GetCustomDomainIds() []*string {
	return s.CustomDomainIds
}

func (s *LLMDeployConfig) GetGatewayType() *string {
	return s.GatewayType
}

func (s *LLMDeployConfig) GetPolicyConfigs() []*PolicyConfig {
	return s.PolicyConfigs
}

func (s *LLMDeployConfig) GetServiceConfigs() []*TargetServiceConfig {
	return s.ServiceConfigs
}

func (s *LLMDeployConfig) SetAutoDeploy(v bool) *LLMDeployConfig {
	s.AutoDeploy = &v
	return s
}

func (s *LLMDeployConfig) SetBackendScene(v string) *LLMDeployConfig {
	s.BackendScene = &v
	return s
}

func (s *LLMDeployConfig) SetCustomDomainIds(v []*string) *LLMDeployConfig {
	s.CustomDomainIds = v
	return s
}

func (s *LLMDeployConfig) SetGatewayType(v string) *LLMDeployConfig {
	s.GatewayType = &v
	return s
}

func (s *LLMDeployConfig) SetPolicyConfigs(v []*PolicyConfig) *LLMDeployConfig {
	s.PolicyConfigs = v
	return s
}

func (s *LLMDeployConfig) SetServiceConfigs(v []*TargetServiceConfig) *LLMDeployConfig {
	s.ServiceConfigs = v
	return s
}

func (s *LLMDeployConfig) Validate() error {
	return dara.Validate(s)
}
