// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPayOrder(v bool) *NodeGroupParam
	GetAutoPayOrder() *bool
	SetAutoRenew(v bool) *NodeGroupParam
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *NodeGroupParam
	GetAutoRenewDuration() *int32
	SetAutoRenewDurationUnit(v string) *NodeGroupParam
	GetAutoRenewDurationUnit() *string
	SetDataDisks(v []*DiskInfo) *NodeGroupParam
	GetDataDisks() []*DiskInfo
	SetDescription(v string) *NodeGroupParam
	GetDescription() *string
	SetInstanceTypes(v []*string) *NodeGroupParam
	GetInstanceTypes() []*string
	SetNodeCount(v int32) *NodeGroupParam
	GetNodeCount() *int32
	SetNodeGroupIndex(v int32) *NodeGroupParam
	GetNodeGroupIndex() *int32
	SetNodeGroupName(v string) *NodeGroupParam
	GetNodeGroupName() *string
	SetNodeGroupType(v string) *NodeGroupParam
	GetNodeGroupType() *string
	SetPaymentDuration(v int32) *NodeGroupParam
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *NodeGroupParam
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *NodeGroupParam
	GetPaymentType() *string
	SetSystemDisk(v *SystemDiskParam) *NodeGroupParam
	GetSystemDisk() *SystemDiskParam
	SetVSwitchIds(v []*string) *NodeGroupParam
	GetVSwitchIds() []*string
	SetZoneId(v string) *NodeGroupParam
	GetZoneId() *string
}

type NodeGroupParam struct {
	AutoPayOrder          *bool            `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	AutoRenew             *bool            `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *int32           `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string          `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	DataDisks             []*DiskInfo      `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	Description           *string          `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceTypes         []*string        `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	NodeCount             *int32           `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupIndex        *int32           `json:"NodeGroupIndex,omitempty" xml:"NodeGroupIndex,omitempty"`
	NodeGroupName         *string          `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupType         *string          `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	PaymentDuration       *int32           `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit   *string          `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType           *string          `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	SystemDisk            *SystemDiskParam `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VSwitchIds            []*string        `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	ZoneId                *string          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s NodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s NodeGroupParam) GoString() string {
	return s.String()
}

func (s *NodeGroupParam) GetAutoPayOrder() *bool {
	return s.AutoPayOrder
}

func (s *NodeGroupParam) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *NodeGroupParam) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *NodeGroupParam) GetAutoRenewDurationUnit() *string {
	return s.AutoRenewDurationUnit
}

func (s *NodeGroupParam) GetDataDisks() []*DiskInfo {
	return s.DataDisks
}

func (s *NodeGroupParam) GetDescription() *string {
	return s.Description
}

func (s *NodeGroupParam) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *NodeGroupParam) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *NodeGroupParam) GetNodeGroupIndex() *int32 {
	return s.NodeGroupIndex
}

func (s *NodeGroupParam) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *NodeGroupParam) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *NodeGroupParam) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *NodeGroupParam) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *NodeGroupParam) GetPaymentType() *string {
	return s.PaymentType
}

func (s *NodeGroupParam) GetSystemDisk() *SystemDiskParam {
	return s.SystemDisk
}

func (s *NodeGroupParam) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *NodeGroupParam) GetZoneId() *string {
	return s.ZoneId
}

func (s *NodeGroupParam) SetAutoPayOrder(v bool) *NodeGroupParam {
	s.AutoPayOrder = &v
	return s
}

func (s *NodeGroupParam) SetAutoRenew(v bool) *NodeGroupParam {
	s.AutoRenew = &v
	return s
}

func (s *NodeGroupParam) SetAutoRenewDuration(v int32) *NodeGroupParam {
	s.AutoRenewDuration = &v
	return s
}

func (s *NodeGroupParam) SetAutoRenewDurationUnit(v string) *NodeGroupParam {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *NodeGroupParam) SetDataDisks(v []*DiskInfo) *NodeGroupParam {
	s.DataDisks = v
	return s
}

func (s *NodeGroupParam) SetDescription(v string) *NodeGroupParam {
	s.Description = &v
	return s
}

func (s *NodeGroupParam) SetInstanceTypes(v []*string) *NodeGroupParam {
	s.InstanceTypes = v
	return s
}

func (s *NodeGroupParam) SetNodeCount(v int32) *NodeGroupParam {
	s.NodeCount = &v
	return s
}

func (s *NodeGroupParam) SetNodeGroupIndex(v int32) *NodeGroupParam {
	s.NodeGroupIndex = &v
	return s
}

func (s *NodeGroupParam) SetNodeGroupName(v string) *NodeGroupParam {
	s.NodeGroupName = &v
	return s
}

func (s *NodeGroupParam) SetNodeGroupType(v string) *NodeGroupParam {
	s.NodeGroupType = &v
	return s
}

func (s *NodeGroupParam) SetPaymentDuration(v int32) *NodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *NodeGroupParam) SetPaymentDurationUnit(v string) *NodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *NodeGroupParam) SetPaymentType(v string) *NodeGroupParam {
	s.PaymentType = &v
	return s
}

func (s *NodeGroupParam) SetSystemDisk(v *SystemDiskParam) *NodeGroupParam {
	s.SystemDisk = v
	return s
}

func (s *NodeGroupParam) SetVSwitchIds(v []*string) *NodeGroupParam {
	s.VSwitchIds = v
	return s
}

func (s *NodeGroupParam) SetZoneId(v string) *NodeGroupParam {
	s.ZoneId = &v
	return s
}

func (s *NodeGroupParam) Validate() error {
	return dara.Validate(s)
}
