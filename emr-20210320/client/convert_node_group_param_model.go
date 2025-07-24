// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *ConvertNodeGroupParam
	GetNodeGroupId() *string
	SetPaymentDuration(v int64) *ConvertNodeGroupParam
	GetPaymentDuration() *int64
	SetPaymentDurationUnit(v string) *ConvertNodeGroupParam
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *ConvertNodeGroupParam
	GetPaymentType() *string
}

type ConvertNodeGroupParam struct {
	// This parameter is required.
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// This parameter is required.
	PaymentDuration     *int64  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// This parameter is required.
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s ConvertNodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s ConvertNodeGroupParam) GoString() string {
	return s.String()
}

func (s *ConvertNodeGroupParam) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ConvertNodeGroupParam) GetPaymentDuration() *int64 {
	return s.PaymentDuration
}

func (s *ConvertNodeGroupParam) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *ConvertNodeGroupParam) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ConvertNodeGroupParam) SetNodeGroupId(v string) *ConvertNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *ConvertNodeGroupParam) SetPaymentDuration(v int64) *ConvertNodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *ConvertNodeGroupParam) SetPaymentDurationUnit(v string) *ConvertNodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ConvertNodeGroupParam) SetPaymentType(v string) *ConvertNodeGroupParam {
	s.PaymentType = &v
	return s
}

func (s *ConvertNodeGroupParam) Validate() error {
	return dara.Validate(s)
}
