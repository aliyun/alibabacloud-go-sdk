// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetNodeCount(v int64) *IncreaseNodeGroupParam
	GetNodeCount() *int64
	SetNodeGroupId(v string) *IncreaseNodeGroupParam
	GetNodeGroupId() *string
	SetPaymentDuration(v int32) *IncreaseNodeGroupParam
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *IncreaseNodeGroupParam
	GetPaymentDurationUnit() *string
	SetVSwitchId(v string) *IncreaseNodeGroupParam
	GetVSwitchId() *string
}

type IncreaseNodeGroupParam struct {
	NodeCount           *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupId         *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PaymentDuration     *int32  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s IncreaseNodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s IncreaseNodeGroupParam) GoString() string {
	return s.String()
}

func (s *IncreaseNodeGroupParam) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *IncreaseNodeGroupParam) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *IncreaseNodeGroupParam) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *IncreaseNodeGroupParam) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *IncreaseNodeGroupParam) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *IncreaseNodeGroupParam) SetNodeCount(v int64) *IncreaseNodeGroupParam {
	s.NodeCount = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetNodeGroupId(v string) *IncreaseNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetPaymentDuration(v int32) *IncreaseNodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetPaymentDurationUnit(v string) *IncreaseNodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetVSwitchId(v string) *IncreaseNodeGroupParam {
	s.VSwitchId = &v
	return s
}

func (s *IncreaseNodeGroupParam) Validate() error {
	return dara.Validate(s)
}
