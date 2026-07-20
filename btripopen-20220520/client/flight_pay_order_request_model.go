// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpPayPrice(v int64) *FlightPayOrderRequest
	GetCorpPayPrice() *int64
	SetDisOrderId(v string) *FlightPayOrderRequest
	GetDisOrderId() *string
	SetExtra(v map[string]*string) *FlightPayOrderRequest
	GetExtra() map[string]*string
	SetPersonalPayPrice(v int64) *FlightPayOrderRequest
	GetPersonalPayPrice() *int64
	SetTotalPayPrice(v int64) *FlightPayOrderRequest
	GetTotalPayPrice() *int64
}

type FlightPayOrderRequest struct {
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
	DisOrderId *string            `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	Extra      map[string]*string `json:"extra,omitempty" xml:"extra,omitempty"`
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

func (s FlightPayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderRequest) GoString() string {
	return s.String()
}

func (s *FlightPayOrderRequest) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *FlightPayOrderRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightPayOrderRequest) GetExtra() map[string]*string {
	return s.Extra
}

func (s *FlightPayOrderRequest) GetPersonalPayPrice() *int64 {
	return s.PersonalPayPrice
}

func (s *FlightPayOrderRequest) GetTotalPayPrice() *int64 {
	return s.TotalPayPrice
}

func (s *FlightPayOrderRequest) SetCorpPayPrice(v int64) *FlightPayOrderRequest {
	s.CorpPayPrice = &v
	return s
}

func (s *FlightPayOrderRequest) SetDisOrderId(v string) *FlightPayOrderRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightPayOrderRequest) SetExtra(v map[string]*string) *FlightPayOrderRequest {
	s.Extra = v
	return s
}

func (s *FlightPayOrderRequest) SetPersonalPayPrice(v int64) *FlightPayOrderRequest {
	s.PersonalPayPrice = &v
	return s
}

func (s *FlightPayOrderRequest) SetTotalPayPrice(v int64) *FlightPayOrderRequest {
	s.TotalPayPrice = &v
	return s
}

func (s *FlightPayOrderRequest) Validate() error {
	return dara.Validate(s)
}
