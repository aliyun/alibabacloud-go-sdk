// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v bool) *ModifyRCInstanceAttributeRequest
	GetDeletionProtection() *bool
	SetEnableJumboFrame(v bool) *ModifyRCInstanceAttributeRequest
	GetEnableJumboFrame() *bool
	SetHostName(v string) *ModifyRCInstanceAttributeRequest
	GetHostName() *string
	SetInstanceId(v string) *ModifyRCInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *ModifyRCInstanceAttributeRequest
	GetInstanceIds() []*string
	SetInstanceName(v string) *ModifyRCInstanceAttributeRequest
	GetInstanceName() *string
	SetPassword(v string) *ModifyRCInstanceAttributeRequest
	GetPassword() *string
	SetReboot(v bool) *ModifyRCInstanceAttributeRequest
	GetReboot() *bool
	SetRegionId(v string) *ModifyRCInstanceAttributeRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *ModifyRCInstanceAttributeRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *ModifyRCInstanceAttributeRequest
	GetSecurityGroupIds() []*string
}

type ModifyRCInstanceAttributeRequest struct {
	DeletionProtection *bool     `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	EnableJumboFrame   *bool     `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	HostName           *string   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId         *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// k8s-node
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Password         *string   `json:"Password,omitempty" xml:"Password,omitempty"`
	Reboot           *bool     `json:"Reboot,omitempty" xml:"Reboot,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId  *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s ModifyRCInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceAttributeRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyRCInstanceAttributeRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *ModifyRCInstanceAttributeRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyRCInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceAttributeRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyRCInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyRCInstanceAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyRCInstanceAttributeRequest) GetReboot() *bool {
	return s.Reboot
}

func (s *ModifyRCInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceAttributeRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyRCInstanceAttributeRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyRCInstanceAttributeRequest) SetDeletionProtection(v bool) *ModifyRCInstanceAttributeRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetEnableJumboFrame(v bool) *ModifyRCInstanceAttributeRequest {
	s.EnableJumboFrame = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetHostName(v string) *ModifyRCInstanceAttributeRequest {
	s.HostName = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetInstanceId(v string) *ModifyRCInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetInstanceIds(v []*string) *ModifyRCInstanceAttributeRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetInstanceName(v string) *ModifyRCInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetPassword(v string) *ModifyRCInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetReboot(v bool) *ModifyRCInstanceAttributeRequest {
	s.Reboot = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetRegionId(v string) *ModifyRCInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetSecurityGroupId(v string) *ModifyRCInstanceAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) SetSecurityGroupIds(v []*string) *ModifyRCInstanceAttributeRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyRCInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
