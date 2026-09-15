// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInnerSandboxRuntimeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *InnerSandboxRuntimeConfig
	GetRole() *string
	SetVpcConfig(v *InnerSandboxRuntimeConfigVpcConfig) *InnerSandboxRuntimeConfig
	GetVpcConfig() *InnerSandboxRuntimeConfigVpcConfig
}

type InnerSandboxRuntimeConfig struct {
	Role      *string                             `json:"role,omitempty" xml:"role,omitempty"`
	VpcConfig *InnerSandboxRuntimeConfigVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s InnerSandboxRuntimeConfig) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxRuntimeConfig) GoString() string {
	return s.String()
}

func (s *InnerSandboxRuntimeConfig) GetRole() *string {
	return s.Role
}

func (s *InnerSandboxRuntimeConfig) GetVpcConfig() *InnerSandboxRuntimeConfigVpcConfig {
	return s.VpcConfig
}

func (s *InnerSandboxRuntimeConfig) SetRole(v string) *InnerSandboxRuntimeConfig {
	s.Role = &v
	return s
}

func (s *InnerSandboxRuntimeConfig) SetVpcConfig(v *InnerSandboxRuntimeConfigVpcConfig) *InnerSandboxRuntimeConfig {
	s.VpcConfig = v
	return s
}

func (s *InnerSandboxRuntimeConfig) Validate() error {
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InnerSandboxRuntimeConfigVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s InnerSandboxRuntimeConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s InnerSandboxRuntimeConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *InnerSandboxRuntimeConfigVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *InnerSandboxRuntimeConfigVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *InnerSandboxRuntimeConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *InnerSandboxRuntimeConfigVpcConfig) SetSecurityGroupId(v string) *InnerSandboxRuntimeConfigVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *InnerSandboxRuntimeConfigVpcConfig) SetVSwitchIds(v []*string) *InnerSandboxRuntimeConfigVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *InnerSandboxRuntimeConfigVpcConfig) SetVpcId(v string) *InnerSandboxRuntimeConfigVpcConfig {
	s.VpcId = &v
	return s
}

func (s *InnerSandboxRuntimeConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}
