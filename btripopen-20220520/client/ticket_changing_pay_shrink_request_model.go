// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingPayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpPayPrice(v int64) *TicketChangingPayShrinkRequest
	GetCorpPayPrice() *int64
	SetDisOrderId(v string) *TicketChangingPayShrinkRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *TicketChangingPayShrinkRequest
	GetDisSubOrderId() *string
	SetExtraShrink(v string) *TicketChangingPayShrinkRequest
	GetExtraShrink() *string
	SetPersonalPayPrice(v int64) *TicketChangingPayShrinkRequest
	GetPersonalPayPrice() *int64
	SetTotalPayPrice(v int64) *TicketChangingPayShrinkRequest
	GetTotalPayPrice() *int64
}

type TicketChangingPayShrinkRequest struct {
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
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	ExtraShrink   *string `json:"extra,omitempty" xml:"extra,omitempty"`
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

func (s TicketChangingPayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingPayShrinkRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingPayShrinkRequest) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *TicketChangingPayShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingPayShrinkRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingPayShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *TicketChangingPayShrinkRequest) GetPersonalPayPrice() *int64 {
	return s.PersonalPayPrice
}

func (s *TicketChangingPayShrinkRequest) GetTotalPayPrice() *int64 {
	return s.TotalPayPrice
}

func (s *TicketChangingPayShrinkRequest) SetCorpPayPrice(v int64) *TicketChangingPayShrinkRequest {
	s.CorpPayPrice = &v
	return s
}

func (s *TicketChangingPayShrinkRequest) SetDisOrderId(v string) *TicketChangingPayShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingPayShrinkRequest) SetDisSubOrderId(v string) *TicketChangingPayShrinkRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingPayShrinkRequest) SetExtraShrink(v string) *TicketChangingPayShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *TicketChangingPayShrinkRequest) SetPersonalPayPrice(v int64) *TicketChangingPayShrinkRequest {
	s.PersonalPayPrice = &v
	return s
}

func (s *TicketChangingPayShrinkRequest) SetTotalPayPrice(v int64) *TicketChangingPayShrinkRequest {
	s.TotalPayPrice = &v
	return s
}

func (s *TicketChangingPayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
