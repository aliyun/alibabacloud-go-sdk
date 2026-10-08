// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v bool) *ModifyRCInstanceAttributeShrinkRequest
	GetDeletionProtection() *bool
	SetEnableJumboFrame(v bool) *ModifyRCInstanceAttributeShrinkRequest
	GetEnableJumboFrame() *bool
	SetHostName(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetHostName() *string
	SetInstanceId(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetInstanceId() *string
	SetInstanceIdsShrink(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetInstanceIdsShrink() *string
	SetInstanceName(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetInstanceName() *string
	SetPassword(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetPassword() *string
	SetReboot(v bool) *ModifyRCInstanceAttributeShrinkRequest
	GetReboot() *bool
	SetRegionId(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIdsShrink(v string) *ModifyRCInstanceAttributeShrinkRequest
	GetSecurityGroupIdsShrink() *string
}

type ModifyRCInstanceAttributeShrinkRequest struct {
	DeletionProtection *bool   `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	EnableJumboFrame   *bool   `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIdsShrink  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// k8s-node
	InstanceName           *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Reboot                 *bool   `json:"Reboot,omitempty" xml:"Reboot,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId        *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIdsShrink *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
}

func (s ModifyRCInstanceAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetReboot() *bool {
	return s.Reboot
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyRCInstanceAttributeShrinkRequest) GetSecurityGroupIdsShrink() *string {
	return s.SecurityGroupIdsShrink
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetDeletionProtection(v bool) *ModifyRCInstanceAttributeShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetEnableJumboFrame(v bool) *ModifyRCInstanceAttributeShrinkRequest {
	s.EnableJumboFrame = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetHostName(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.HostName = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetInstanceId(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetInstanceIdsShrink(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetInstanceName(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetPassword(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.Password = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetReboot(v bool) *ModifyRCInstanceAttributeShrinkRequest {
	s.Reboot = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetRegionId(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetSecurityGroupId(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) SetSecurityGroupIdsShrink(v string) *ModifyRCInstanceAttributeShrinkRequest {
	s.SecurityGroupIdsShrink = &v
	return s
}

func (s *ModifyRCInstanceAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
