// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceVpcRegionDetail interface {
	dara.Model
	String() string
	GoString() string
	SetFcConfig(v string) *GrafanaWorkspaceVpcRegionDetail
	GetFcConfig() *string
	SetId(v int64) *GrafanaWorkspaceVpcRegionDetail
	GetId() *int64
	SetInstallStatus(v string) *GrafanaWorkspaceVpcRegionDetail
	GetInstallStatus() *string
	SetName(v string) *GrafanaWorkspaceVpcRegionDetail
	GetName() *string
	SetRegionId(v string) *GrafanaWorkspaceVpcRegionDetail
	GetRegionId() *string
	SetSecurityGroupId(v string) *GrafanaWorkspaceVpcRegionDetail
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *GrafanaWorkspaceVpcRegionDetail
	GetSecurityGroupIds() []*string
	SetUserId(v string) *GrafanaWorkspaceVpcRegionDetail
	GetUserId() *string
	SetVSwitchId(v string) *GrafanaWorkspaceVpcRegionDetail
	GetVSwitchId() *string
	SetVSwitchIds(v []*string) *GrafanaWorkspaceVpcRegionDetail
	GetVSwitchIds() []*string
	SetVpcId(v string) *GrafanaWorkspaceVpcRegionDetail
	GetVpcId() *string
}

type GrafanaWorkspaceVpcRegionDetail struct {
	// example:
	//
	// fc、ansm
	FcConfig *string `json:"fcConfig,omitempty" xml:"fcConfig,omitempty"`
	// example:
	//
	// 配置ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// Uninitialized
	InstallStatus *string `json:"installStatus,omitempty" xml:"installStatus,omitempty"`
	// example:
	//
	// vpc-abc*****
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId         *string   `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SecurityGroupId  *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 10988**********
	UserId     *string   `json:"userId,omitempty" xml:"userId,omitempty"`
	VSwitchId  *string   `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2ze4siu98**********
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GrafanaWorkspaceVpcRegionDetail) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceVpcRegionDetail) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetFcConfig() *string {
	return s.FcConfig
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetInstallStatus() *string {
	return s.InstallStatus
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GrafanaWorkspaceVpcRegionDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetFcConfig(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.FcConfig = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetId(v int64) *GrafanaWorkspaceVpcRegionDetail {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetInstallStatus(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.InstallStatus = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetName(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetRegionId(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.RegionId = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetSecurityGroupId(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.SecurityGroupId = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetSecurityGroupIds(v []*string) *GrafanaWorkspaceVpcRegionDetail {
	s.SecurityGroupIds = v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetUserId(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetVSwitchId(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.VSwitchId = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetVSwitchIds(v []*string) *GrafanaWorkspaceVpcRegionDetail {
	s.VSwitchIds = v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) SetVpcId(v string) *GrafanaWorkspaceVpcRegionDetail {
	s.VpcId = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegionDetail) Validate() error {
	return dara.Validate(s)
}
