// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientSecurityGroupIds(v []*string) *StartInstanceRequest
	GetClientSecurityGroupIds() []*string
	SetEnablePortalPrivateAccess(v bool) *StartInstanceRequest
	GetEnablePortalPrivateAccess() *bool
	SetInstanceId(v string) *StartInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *StartInstanceRequest
	GetRegionId() *string
	SetSecurityGroupIds(v []*string) *StartInstanceRequest
	GetSecurityGroupIds() []*string
	SetSlaveVswitchId(v string) *StartInstanceRequest
	GetSlaveVswitchId() *string
	SetVswitchId(v string) *StartInstanceRequest
	GetVswitchId() *string
}

type StartInstanceRequest struct {
	ClientSecurityGroupIds    []*string `json:"ClientSecurityGroupIds,omitempty" xml:"ClientSecurityGroupIds,omitempty" type:"Repeated"`
	EnablePortalPrivateAccess *bool     `json:"EnablePortalPrivateAccess,omitempty" xml:"EnablePortalPrivateAccess,omitempty"`
	// The ID of the bastion host that you want to enable.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// An array consisting of the IDs of security groups to which the bastion host is added.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp1aiupc4yjqgmm****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SlaveVswitchId   *string   `json:"SlaveVswitchId,omitempty" xml:"SlaveVswitchId,omitempty"`
	// The ID of the vSwitch to which the bastion host belongs.
	//
	// example:
	//
	// vsw-bp1xfwzzfti0kjbf****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetClientSecurityGroupIds() []*string {
	return s.ClientSecurityGroupIds
}

func (s *StartInstanceRequest) GetEnablePortalPrivateAccess() *bool {
	return s.EnablePortalPrivateAccess
}

func (s *StartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstanceRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *StartInstanceRequest) GetSlaveVswitchId() *string {
	return s.SlaveVswitchId
}

func (s *StartInstanceRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *StartInstanceRequest) SetClientSecurityGroupIds(v []*string) *StartInstanceRequest {
	s.ClientSecurityGroupIds = v
	return s
}

func (s *StartInstanceRequest) SetEnablePortalPrivateAccess(v bool) *StartInstanceRequest {
	s.EnablePortalPrivateAccess = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) SetSecurityGroupIds(v []*string) *StartInstanceRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *StartInstanceRequest) SetSlaveVswitchId(v string) *StartInstanceRequest {
	s.SlaveVswitchId = &v
	return s
}

func (s *StartInstanceRequest) SetVswitchId(v string) *StartInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
