// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundFeeData interface {
	dara.Model
	String() string
	GoString() string
	SetMaxRefundFee(v int64) *RefundFeeData
	GetMaxRefundFee() *int64
	SetMinRefundFee(v int64) *RefundFeeData
	GetMinRefundFee() *int64
}

type RefundFeeData struct {
	// example:
	//
	// 100
	MaxRefundFee *int64 `json:"maxRefundFee,omitempty" xml:"maxRefundFee,omitempty"`
	// example:
	//
	// 1
	MinRefundFee *int64 `json:"minRefundFee,omitempty" xml:"minRefundFee,omitempty"`
}

func (s RefundFeeData) String() string {
	return dara.Prettify(s)
}

func (s RefundFeeData) GoString() string {
	return s.String()
}

func (s *RefundFeeData) GetMaxRefundFee() *int64 {
	return s.MaxRefundFee
}

func (s *RefundFeeData) GetMinRefundFee() *int64 {
	return s.MinRefundFee
}

func (s *RefundFeeData) SetMaxRefundFee(v int64) *RefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *RefundFeeData) SetMinRefundFee(v int64) *RefundFeeData {
	s.MinRefundFee = &v
	return s
}

func (s *RefundFeeData) Validate() error {
	return dara.Validate(s)
}
