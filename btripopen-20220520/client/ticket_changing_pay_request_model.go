// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpPayPrice(v int64) *TicketChangingPayRequest
	GetCorpPayPrice() *int64
	SetDisOrderId(v string) *TicketChangingPayRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *TicketChangingPayRequest
	GetDisSubOrderId() *string
	SetExtra(v map[string]*string) *TicketChangingPayRequest
	GetExtra() map[string]*string
	SetPersonalPayPrice(v int64) *TicketChangingPayRequest
	GetPersonalPayPrice() *int64
	SetTotalPayPrice(v int64) *TicketChangingPayRequest
	GetTotalPayPrice() *int64
}

type TicketChangingPayRequest struct {
	// example:
	//
	// 100
	CorpPayPrice *int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// refun123
	DisSubOrderId *string            `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	Extra         map[string]*string `json:"extra,omitempty" xml:"extra,omitempty"`
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

func (s TicketChangingPayRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingPayRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingPayRequest) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *TicketChangingPayRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingPayRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingPayRequest) GetExtra() map[string]*string {
	return s.Extra
}

func (s *TicketChangingPayRequest) GetPersonalPayPrice() *int64 {
	return s.PersonalPayPrice
}

func (s *TicketChangingPayRequest) GetTotalPayPrice() *int64 {
	return s.TotalPayPrice
}

func (s *TicketChangingPayRequest) SetCorpPayPrice(v int64) *TicketChangingPayRequest {
	s.CorpPayPrice = &v
	return s
}

func (s *TicketChangingPayRequest) SetDisOrderId(v string) *TicketChangingPayRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingPayRequest) SetDisSubOrderId(v string) *TicketChangingPayRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingPayRequest) SetExtra(v map[string]*string) *TicketChangingPayRequest {
	s.Extra = v
	return s
}

func (s *TicketChangingPayRequest) SetPersonalPayPrice(v int64) *TicketChangingPayRequest {
	s.PersonalPayPrice = &v
	return s
}

func (s *TicketChangingPayRequest) SetTotalPayPrice(v int64) *TicketChangingPayRequest {
	s.TotalPayPrice = &v
	return s
}

func (s *TicketChangingPayRequest) Validate() error {
	return dara.Validate(s)
}
