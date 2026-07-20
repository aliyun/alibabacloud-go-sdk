// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightRefundApplyShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightRefundApplyShrinkRequest
	GetOutOrderId() *string
	SetOutRefundApplyId(v string) *IntlFlightRefundApplyShrinkRequest
	GetOutRefundApplyId() *string
	SetPassengerJourneyGroupKey(v string) *IntlFlightRefundApplyShrinkRequest
	GetPassengerJourneyGroupKey() *string
	SetRefundReasonCode(v string) *IntlFlightRefundApplyShrinkRequest
	GetRefundReasonCode() *string
	SetRefundSegmentListShrink(v string) *IntlFlightRefundApplyShrinkRequest
	GetRefundSegmentListShrink() *string
	SetSelectedPassengersShrink(v string) *IntlFlightRefundApplyShrinkRequest
	GetSelectedPassengersShrink() *string
}

type IntlFlightRefundApplyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 3750656668336001024
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 202503251022170001
	OutRefundApplyId *string `json:"out_refund_apply_id,omitempty" xml:"out_refund_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edcac4f4c79d40ccb141ddb6da567e65
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RefundReasonCode *string `json:"refund_reason_code,omitempty" xml:"refund_reason_code,omitempty"`
	// This parameter is required.
	RefundSegmentListShrink *string `json:"refund_segment_list,omitempty" xml:"refund_segment_list,omitempty"`
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
}

func (s IntlFlightRefundApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightRefundApplyShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightRefundApplyShrinkRequest) GetOutRefundApplyId() *string {
	return s.OutRefundApplyId
}

func (s *IntlFlightRefundApplyShrinkRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightRefundApplyShrinkRequest) GetRefundReasonCode() *string {
	return s.RefundReasonCode
}

func (s *IntlFlightRefundApplyShrinkRequest) GetRefundSegmentListShrink() *string {
	return s.RefundSegmentListShrink
}

func (s *IntlFlightRefundApplyShrinkRequest) GetSelectedPassengersShrink() *string {
	return s.SelectedPassengersShrink
}

func (s *IntlFlightRefundApplyShrinkRequest) SetOrderId(v string) *IntlFlightRefundApplyShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) SetOutOrderId(v string) *IntlFlightRefundApplyShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) SetOutRefundApplyId(v string) *IntlFlightRefundApplyShrinkRequest {
	s.OutRefundApplyId = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightRefundApplyShrinkRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) SetRefundReasonCode(v string) *IntlFlightRefundApplyShrinkRequest {
	s.RefundReasonCode = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) SetRefundSegmentListShrink(v string) *IntlFlightRefundApplyShrinkRequest {
	s.RefundSegmentListShrink = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) SetSelectedPassengersShrink(v string) *IntlFlightRefundApplyShrinkRequest {
	s.SelectedPassengersShrink = &v
	return s
}

func (s *IntlFlightRefundApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
