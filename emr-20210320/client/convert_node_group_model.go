// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertNodeGroup interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *ConvertNodeGroup
	GetNodeGroupId() *string
	SetPaymentDuration(v int32) *ConvertNodeGroup
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *ConvertNodeGroup
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *ConvertNodeGroup
	GetPaymentType() *string
}

type ConvertNodeGroup struct {
	// 节点组ID。
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
	// 付费类型。
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s ConvertNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s ConvertNodeGroup) GoString() string {
	return s.String()
}

func (s *ConvertNodeGroup) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ConvertNodeGroup) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *ConvertNodeGroup) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *ConvertNodeGroup) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ConvertNodeGroup) SetNodeGroupId(v string) *ConvertNodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *ConvertNodeGroup) SetPaymentDuration(v int32) *ConvertNodeGroup {
	s.PaymentDuration = &v
	return s
}

func (s *ConvertNodeGroup) SetPaymentDurationUnit(v string) *ConvertNodeGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ConvertNodeGroup) SetPaymentType(v string) *ConvertNodeGroup {
	s.PaymentType = &v
	return s
}

func (s *ConvertNodeGroup) Validate() error {
	return dara.Validate(s)
}
