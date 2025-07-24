// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateNodeGroupParam
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *CreateNodeGroupParam
	GetAutoRenewDuration() *int32
	SetAutoRenewDurationUnit(v string) *CreateNodeGroupParam
	GetAutoRenewDurationUnit() *string
	SetDataDisks(v []*DiskInfo) *CreateNodeGroupParam
	GetDataDisks() []*DiskInfo
	SetInstanceTypes(v []*string) *CreateNodeGroupParam
	GetInstanceTypes() []*string
	SetNodeCount(v int32) *CreateNodeGroupParam
	GetNodeCount() *int32
	SetNodeGroupName(v string) *CreateNodeGroupParam
	GetNodeGroupName() *string
	SetNodeGroupType(v string) *CreateNodeGroupParam
	GetNodeGroupType() *string
	SetNodeKeyPairName(v string) *CreateNodeGroupParam
	GetNodeKeyPairName() *string
	SetNodeRamRole(v string) *CreateNodeGroupParam
	GetNodeRamRole() *string
	SetNodeRootPassword(v string) *CreateNodeGroupParam
	GetNodeRootPassword() *string
	SetPaymentDuration(v int32) *CreateNodeGroupParam
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *CreateNodeGroupParam
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *CreateNodeGroupParam
	GetPaymentType() *string
	SetSecurityGroupId(v string) *CreateNodeGroupParam
	GetSecurityGroupId() *string
	SetSpotStrategy(v string) *CreateNodeGroupParam
	GetSpotStrategy() *string
	SetSystemDisk(v *SystemDiskParam) *CreateNodeGroupParam
	GetSystemDisk() *SystemDiskParam
	SetVSwitchIds(v []*string) *CreateNodeGroupParam
	GetVSwitchIds() []*string
	SetWithPublicIp(v bool) *CreateNodeGroupParam
	GetWithPublicIp() *bool
	SetZoneId(v string) *CreateNodeGroupParam
	GetZoneId() *string
}

type CreateNodeGroupParam struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// example:
	//
	// Monthly
	AutoRenewDurationUnit *string     `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	DataDisks             []*DiskInfo `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	InstanceTypes         []*string   `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// core-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// example:
	//
	// sshkey
	NodeKeyPairName *string `json:"NodeKeyPairName,omitempty" xml:"NodeKeyPairName,omitempty"`
	// example:
	//
	// AliyunEmrEcsDefaultRole
	NodeRamRole *string `json:"NodeRamRole,omitempty" xml:"NodeRamRole,omitempty"`
	// example:
	//
	// *******
	NodeRootPassword *string `json:"NodeRootPassword,omitempty" xml:"NodeRootPassword,omitempty"`
	// example:
	//
	// 1
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// example:
	//
	// Monthly
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// sg-hp3abbae8lb6lmb1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string          `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk   *SystemDiskParam `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VSwitchIds   []*string        `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	WithPublicIp *bool `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateNodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupParam) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupParam) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateNodeGroupParam) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *CreateNodeGroupParam) GetAutoRenewDurationUnit() *string {
	return s.AutoRenewDurationUnit
}

func (s *CreateNodeGroupParam) GetDataDisks() []*DiskInfo {
	return s.DataDisks
}

func (s *CreateNodeGroupParam) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateNodeGroupParam) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateNodeGroupParam) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateNodeGroupParam) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *CreateNodeGroupParam) GetNodeKeyPairName() *string {
	return s.NodeKeyPairName
}

func (s *CreateNodeGroupParam) GetNodeRamRole() *string {
	return s.NodeRamRole
}

func (s *CreateNodeGroupParam) GetNodeRootPassword() *string {
	return s.NodeRootPassword
}

func (s *CreateNodeGroupParam) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *CreateNodeGroupParam) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *CreateNodeGroupParam) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateNodeGroupParam) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateNodeGroupParam) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateNodeGroupParam) GetSystemDisk() *SystemDiskParam {
	return s.SystemDisk
}

func (s *CreateNodeGroupParam) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateNodeGroupParam) GetWithPublicIp() *bool {
	return s.WithPublicIp
}

func (s *CreateNodeGroupParam) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateNodeGroupParam) SetAutoRenew(v bool) *CreateNodeGroupParam {
	s.AutoRenew = &v
	return s
}

func (s *CreateNodeGroupParam) SetAutoRenewDuration(v int32) *CreateNodeGroupParam {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateNodeGroupParam) SetAutoRenewDurationUnit(v string) *CreateNodeGroupParam {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *CreateNodeGroupParam) SetDataDisks(v []*DiskInfo) *CreateNodeGroupParam {
	s.DataDisks = v
	return s
}

func (s *CreateNodeGroupParam) SetInstanceTypes(v []*string) *CreateNodeGroupParam {
	s.InstanceTypes = v
	return s
}

func (s *CreateNodeGroupParam) SetNodeCount(v int32) *CreateNodeGroupParam {
	s.NodeCount = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeGroupName(v string) *CreateNodeGroupParam {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeGroupType(v string) *CreateNodeGroupParam {
	s.NodeGroupType = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeKeyPairName(v string) *CreateNodeGroupParam {
	s.NodeKeyPairName = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeRamRole(v string) *CreateNodeGroupParam {
	s.NodeRamRole = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeRootPassword(v string) *CreateNodeGroupParam {
	s.NodeRootPassword = &v
	return s
}

func (s *CreateNodeGroupParam) SetPaymentDuration(v int32) *CreateNodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *CreateNodeGroupParam) SetPaymentDurationUnit(v string) *CreateNodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateNodeGroupParam) SetPaymentType(v string) *CreateNodeGroupParam {
	s.PaymentType = &v
	return s
}

func (s *CreateNodeGroupParam) SetSecurityGroupId(v string) *CreateNodeGroupParam {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateNodeGroupParam) SetSpotStrategy(v string) *CreateNodeGroupParam {
	s.SpotStrategy = &v
	return s
}

func (s *CreateNodeGroupParam) SetSystemDisk(v *SystemDiskParam) *CreateNodeGroupParam {
	s.SystemDisk = v
	return s
}

func (s *CreateNodeGroupParam) SetVSwitchIds(v []*string) *CreateNodeGroupParam {
	s.VSwitchIds = v
	return s
}

func (s *CreateNodeGroupParam) SetWithPublicIp(v bool) *CreateNodeGroupParam {
	s.WithPublicIp = &v
	return s
}

func (s *CreateNodeGroupParam) SetZoneId(v string) *CreateNodeGroupParam {
	s.ZoneId = &v
	return s
}

func (s *CreateNodeGroupParam) Validate() error {
	return dara.Validate(s)
}
