// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightRefundApplyRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightRefundApplyRequest
	GetOutOrderId() *string
	SetOutRefundApplyId(v string) *IntlFlightRefundApplyRequest
	GetOutRefundApplyId() *string
	SetPassengerJourneyGroupKey(v string) *IntlFlightRefundApplyRequest
	GetPassengerJourneyGroupKey() *string
	SetRefundReasonCode(v string) *IntlFlightRefundApplyRequest
	GetRefundReasonCode() *string
	SetRefundSegmentList(v []*IntlFlightRefundApplyRequestRefundSegmentList) *IntlFlightRefundApplyRequest
	GetRefundSegmentList() []*IntlFlightRefundApplyRequestRefundSegmentList
	SetSelectedPassengers(v []*IntlFlightRefundApplyRequestSelectedPassengers) *IntlFlightRefundApplyRequest
	GetSelectedPassengers() []*IntlFlightRefundApplyRequestSelectedPassengers
}

type IntlFlightRefundApplyRequest struct {
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
	RefundSegmentList []*IntlFlightRefundApplyRequestRefundSegmentList `json:"refund_segment_list,omitempty" xml:"refund_segment_list,omitempty" type:"Repeated"`
	// This parameter is required.
	SelectedPassengers []*IntlFlightRefundApplyRequestSelectedPassengers `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty" type:"Repeated"`
}

func (s IntlFlightRefundApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightRefundApplyRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightRefundApplyRequest) GetOutRefundApplyId() *string {
	return s.OutRefundApplyId
}

func (s *IntlFlightRefundApplyRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightRefundApplyRequest) GetRefundReasonCode() *string {
	return s.RefundReasonCode
}

func (s *IntlFlightRefundApplyRequest) GetRefundSegmentList() []*IntlFlightRefundApplyRequestRefundSegmentList {
	return s.RefundSegmentList
}

func (s *IntlFlightRefundApplyRequest) GetSelectedPassengers() []*IntlFlightRefundApplyRequestSelectedPassengers {
	return s.SelectedPassengers
}

func (s *IntlFlightRefundApplyRequest) SetOrderId(v string) *IntlFlightRefundApplyRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightRefundApplyRequest) SetOutOrderId(v string) *IntlFlightRefundApplyRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightRefundApplyRequest) SetOutRefundApplyId(v string) *IntlFlightRefundApplyRequest {
	s.OutRefundApplyId = &v
	return s
}

func (s *IntlFlightRefundApplyRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightRefundApplyRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightRefundApplyRequest) SetRefundReasonCode(v string) *IntlFlightRefundApplyRequest {
	s.RefundReasonCode = &v
	return s
}

func (s *IntlFlightRefundApplyRequest) SetRefundSegmentList(v []*IntlFlightRefundApplyRequestRefundSegmentList) *IntlFlightRefundApplyRequest {
	s.RefundSegmentList = v
	return s
}

func (s *IntlFlightRefundApplyRequest) SetSelectedPassengers(v []*IntlFlightRefundApplyRequestSelectedPassengers) *IntlFlightRefundApplyRequest {
	s.SelectedPassengers = v
	return s
}

func (s *IntlFlightRefundApplyRequest) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundApplyRequestRefundSegmentList struct {
	// This parameter is required.
	//
	// example:
	//
	// CZ5009PKXHKG0616
	SegmentKey *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
}

func (s IntlFlightRefundApplyRequestRefundSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyRequestRefundSegmentList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyRequestRefundSegmentList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightRefundApplyRequestRefundSegmentList) SetSegmentKey(v string) *IntlFlightRefundApplyRequestRefundSegmentList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightRefundApplyRequestRefundSegmentList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundApplyRequestSelectedPassengers struct {
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightRefundApplyRequestSelectedPassengers) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyRequestSelectedPassengers) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyRequestSelectedPassengers) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightRefundApplyRequestSelectedPassengers) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightRefundApplyRequestSelectedPassengers) SetFullName(v string) *IntlFlightRefundApplyRequestSelectedPassengers {
	s.FullName = &v
	return s
}

func (s *IntlFlightRefundApplyRequestSelectedPassengers) SetPassengerId(v int64) *IntlFlightRefundApplyRequestSelectedPassengers {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightRefundApplyRequestSelectedPassengers) Validate() error {
	return dara.Validate(s)
}
