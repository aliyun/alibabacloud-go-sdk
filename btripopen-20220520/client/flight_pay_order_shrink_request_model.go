// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpPayPrice(v int64) *FlightPayOrderShrinkRequest
	GetCorpPayPrice() *int64
	SetDisOrderId(v string) *FlightPayOrderShrinkRequest
	GetDisOrderId() *string
	SetExtraShrink(v string) *FlightPayOrderShrinkRequest
	GetExtraShrink() *string
	SetPersonalPayPrice(v int64) *FlightPayOrderShrinkRequest
	GetPersonalPayPrice() *int64
	SetTotalPayPrice(v int64) *FlightPayOrderShrinkRequest
	GetTotalPayPrice() *int64
}

type FlightPayOrderShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	CorpPayPrice *int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId  *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	ExtraShrink *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PersonalPayPrice *int64 `json:"personal_pay_price,omitempty" xml:"personal_pay_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalPayPrice *int64 `json:"total_pay_price,omitempty" xml:"total_pay_price,omitempty"`
}

func (s FlightPayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightPayOrderShrinkRequest) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *FlightPayOrderShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightPayOrderShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *FlightPayOrderShrinkRequest) GetPersonalPayPrice() *int64 {
	return s.PersonalPayPrice
}

func (s *FlightPayOrderShrinkRequest) GetTotalPayPrice() *int64 {
	return s.TotalPayPrice
}

func (s *FlightPayOrderShrinkRequest) SetCorpPayPrice(v int64) *FlightPayOrderShrinkRequest {
	s.CorpPayPrice = &v
	return s
}

func (s *FlightPayOrderShrinkRequest) SetDisOrderId(v string) *FlightPayOrderShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightPayOrderShrinkRequest) SetExtraShrink(v string) *FlightPayOrderShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *FlightPayOrderShrinkRequest) SetPersonalPayPrice(v int64) *FlightPayOrderShrinkRequest {
	s.PersonalPayPrice = &v
	return s
}

func (s *FlightPayOrderShrinkRequest) SetTotalPayPrice(v int64) *FlightPayOrderShrinkRequest {
	s.TotalPayPrice = &v
	return s
}

func (s *FlightPayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
