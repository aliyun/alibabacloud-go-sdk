// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *RefundApplyShrinkRequest
	GetOrderNum() *int64
	SetRefundJourneysShrink(v string) *RefundApplyShrinkRequest
	GetRefundJourneysShrink() *string
	SetRefundPassengerListShrink(v string) *RefundApplyShrinkRequest
	GetRefundPassengerListShrink() *string
	SetRefundTypeShrink(v string) *RefundApplyShrinkRequest
	GetRefundTypeShrink() *string
}

type RefundApplyShrinkRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// journeys for which a refund is being requested
	//
	// This parameter is required.
	RefundJourneysShrink *string `json:"refund_journeys,omitempty" xml:"refund_journeys,omitempty"`
	// passengers that applying for a refund
	//
	// This parameter is required.
	RefundPassengerListShrink *string `json:"refund_passenger_list,omitempty" xml:"refund_passenger_list,omitempty"`
	// refund type and attachments
	//
	// This parameter is required.
	RefundTypeShrink *string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

func (s RefundApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefundApplyShrinkRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *RefundApplyShrinkRequest) GetRefundJourneysShrink() *string {
	return s.RefundJourneysShrink
}

func (s *RefundApplyShrinkRequest) GetRefundPassengerListShrink() *string {
	return s.RefundPassengerListShrink
}

func (s *RefundApplyShrinkRequest) GetRefundTypeShrink() *string {
	return s.RefundTypeShrink
}

func (s *RefundApplyShrinkRequest) SetOrderNum(v int64) *RefundApplyShrinkRequest {
	s.OrderNum = &v
	return s
}

func (s *RefundApplyShrinkRequest) SetRefundJourneysShrink(v string) *RefundApplyShrinkRequest {
	s.RefundJourneysShrink = &v
	return s
}

func (s *RefundApplyShrinkRequest) SetRefundPassengerListShrink(v string) *RefundApplyShrinkRequest {
	s.RefundPassengerListShrink = &v
	return s
}

func (s *RefundApplyShrinkRequest) SetRefundTypeShrink(v string) *RefundApplyShrinkRequest {
	s.RefundTypeShrink = &v
	return s
}

func (s *RefundApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
