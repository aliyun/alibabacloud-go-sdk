// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseNodeGroup interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *IncreaseNodeGroup
	GetDescription() *string
	SetNodeCount(v int32) *IncreaseNodeGroup
	GetNodeCount() *int32
	SetNodeGroupId(v string) *IncreaseNodeGroup
	GetNodeGroupId() *string
	SetPaymentDuration(v int32) *IncreaseNodeGroup
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *IncreaseNodeGroup
	GetPaymentDurationUnit() *string
	SetVSwitchId(v string) *IncreaseNodeGroup
	GetVSwitchId() *string
}

type IncreaseNodeGroup struct {
	// 描述。
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 节点数量。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// 节点组ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// 付费时长。
	//
	// example:
	//
	// 12
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// 付费时长单位。
	//
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// 虚拟机交换机ID。
	//
	// example:
	//
	// vsw-hp35g7ya5ymw68mmg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s IncreaseNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s IncreaseNodeGroup) GoString() string {
	return s.String()
}

func (s *IncreaseNodeGroup) GetDescription() *string {
	return s.Description
}

func (s *IncreaseNodeGroup) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *IncreaseNodeGroup) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *IncreaseNodeGroup) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *IncreaseNodeGroup) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *IncreaseNodeGroup) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *IncreaseNodeGroup) SetDescription(v string) *IncreaseNodeGroup {
	s.Description = &v
	return s
}

func (s *IncreaseNodeGroup) SetNodeCount(v int32) *IncreaseNodeGroup {
	s.NodeCount = &v
	return s
}

func (s *IncreaseNodeGroup) SetNodeGroupId(v string) *IncreaseNodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *IncreaseNodeGroup) SetPaymentDuration(v int32) *IncreaseNodeGroup {
	s.PaymentDuration = &v
	return s
}

func (s *IncreaseNodeGroup) SetPaymentDurationUnit(v string) *IncreaseNodeGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *IncreaseNodeGroup) SetVSwitchId(v string) *IncreaseNodeGroup {
	s.VSwitchId = &v
	return s
}

func (s *IncreaseNodeGroup) Validate() error {
	return dara.Validate(s)
}
