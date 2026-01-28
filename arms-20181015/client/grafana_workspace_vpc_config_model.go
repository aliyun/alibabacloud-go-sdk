// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceVpcConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFcConfig(v string) *GrafanaWorkspaceVpcConfig
	GetFcConfig() *string
	SetId(v int64) *GrafanaWorkspaceVpcConfig
	GetId() *int64
	SetInstallMsg(v string) *GrafanaWorkspaceVpcConfig
	GetInstallMsg() *string
	SetInstallStatus(v string) *GrafanaWorkspaceVpcConfig
	GetInstallStatus() *string
	SetName(v string) *GrafanaWorkspaceVpcConfig
	GetName() *string
	SetRegionId(v string) *GrafanaWorkspaceVpcConfig
	GetRegionId() *string
	SetSecurityGroupId(v string) *GrafanaWorkspaceVpcConfig
	GetSecurityGroupId() *string
	SetUserId(v string) *GrafanaWorkspaceVpcConfig
	GetUserId() *string
	SetVSwitchId(v string) *GrafanaWorkspaceVpcConfig
	GetVSwitchId() *string
	SetVpcId(v string) *GrafanaWorkspaceVpcConfig
	GetVpcId() *string
}

type GrafanaWorkspaceVpcConfig struct {
	// example:
	//
	// fc、ansm
	FcConfig *string `json:"fcConfig,omitempty" xml:"fcConfig,omitempty"`
	// example:
	//
	// 1
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	InstallMsg *string `json:"installMsg,omitempty" xml:"installMsg,omitempty"`
	// example:
	//
	// CreateSucceed
	InstallStatus *string `json:"installStatus,omitempty" xml:"installStatus,omitempty"`
	// example:
	//
	// 北京VPC-A通道
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// sg-6we94uvybteyc******
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// 10983***********
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// vsw-6we3**********
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// vpc-6wehr2x**********
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GrafanaWorkspaceVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceVpcConfig) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceVpcConfig) GetFcConfig() *string {
	return s.FcConfig
}

func (s *GrafanaWorkspaceVpcConfig) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceVpcConfig) GetInstallMsg() *string {
	return s.InstallMsg
}

func (s *GrafanaWorkspaceVpcConfig) GetInstallStatus() *string {
	return s.InstallStatus
}

func (s *GrafanaWorkspaceVpcConfig) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceVpcConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *GrafanaWorkspaceVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GrafanaWorkspaceVpcConfig) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceVpcConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GrafanaWorkspaceVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *GrafanaWorkspaceVpcConfig) SetFcConfig(v string) *GrafanaWorkspaceVpcConfig {
	s.FcConfig = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetId(v int64) *GrafanaWorkspaceVpcConfig {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetInstallMsg(v string) *GrafanaWorkspaceVpcConfig {
	s.InstallMsg = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetInstallStatus(v string) *GrafanaWorkspaceVpcConfig {
	s.InstallStatus = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetName(v string) *GrafanaWorkspaceVpcConfig {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetRegionId(v string) *GrafanaWorkspaceVpcConfig {
	s.RegionId = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetSecurityGroupId(v string) *GrafanaWorkspaceVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetUserId(v string) *GrafanaWorkspaceVpcConfig {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetVSwitchId(v string) *GrafanaWorkspaceVpcConfig {
	s.VSwitchId = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) SetVpcId(v string) *GrafanaWorkspaceVpcConfig {
	s.VpcId = &v
	return s
}

func (s *GrafanaWorkspaceVpcConfig) Validate() error {
	return dara.Validate(s)
}
