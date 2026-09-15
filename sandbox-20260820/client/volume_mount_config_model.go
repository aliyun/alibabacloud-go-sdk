// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVolumeMountConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *VolumeMountConfig
	GetRole() *string
	SetVpcConfig(v *VolumeMountConfigVpcConfig) *VolumeMountConfig
	GetVpcConfig() *VolumeMountConfigVpcConfig
}

type VolumeMountConfig struct {
	Role      *string                     `json:"role,omitempty" xml:"role,omitempty"`
	VpcConfig *VolumeMountConfigVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s VolumeMountConfig) String() string {
	return dara.Prettify(s)
}

func (s VolumeMountConfig) GoString() string {
	return s.String()
}

func (s *VolumeMountConfig) GetRole() *string {
	return s.Role
}

func (s *VolumeMountConfig) GetVpcConfig() *VolumeMountConfigVpcConfig {
	return s.VpcConfig
}

func (s *VolumeMountConfig) SetRole(v string) *VolumeMountConfig {
	s.Role = &v
	return s
}

func (s *VolumeMountConfig) SetVpcConfig(v *VolumeMountConfigVpcConfig) *VolumeMountConfig {
	s.VpcConfig = v
	return s
}

func (s *VolumeMountConfig) Validate() error {
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VolumeMountConfigVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s VolumeMountConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s VolumeMountConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *VolumeMountConfigVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *VolumeMountConfigVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *VolumeMountConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *VolumeMountConfigVpcConfig) SetSecurityGroupId(v string) *VolumeMountConfigVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *VolumeMountConfigVpcConfig) SetVSwitchIds(v []*string) *VolumeMountConfigVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *VolumeMountConfigVpcConfig) SetVpcId(v string) *VolumeMountConfigVpcConfig {
	s.VpcId = &v
	return s
}

func (s *VolumeMountConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}
