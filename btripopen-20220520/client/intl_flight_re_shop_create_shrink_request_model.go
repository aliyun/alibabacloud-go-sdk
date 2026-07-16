// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCreateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncApplyKey(v string) *IntlFlightReShopCreateShrinkRequest
	GetAsyncApplyKey() *string
	SetAsyncApplyMode(v bool) *IntlFlightReShopCreateShrinkRequest
	GetAsyncApplyMode() *bool
	SetOrderId(v string) *IntlFlightReShopCreateShrinkRequest
	GetOrderId() *string
	SetOtaItemId(v string) *IntlFlightReShopCreateShrinkRequest
	GetOtaItemId() *string
	SetOutOrderId(v string) *IntlFlightReShopCreateShrinkRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopCreateShrinkRequest
	GetOutReShopApplyId() *string
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopCreateShrinkRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopCreateShrinkRequest
	GetReShopReasonCode() *string
	SetSelectedPassengersShrink(v string) *IntlFlightReShopCreateShrinkRequest
	GetSelectedPassengersShrink() *string
	SetTotalReShopFee(v int64) *IntlFlightReShopCreateShrinkRequest
	GetTotalReShopFee() *int64
}

type IntlFlightReShopCreateShrinkRequest struct {
	AsyncApplyKey  *string `json:"async_apply_key,omitempty" xml:"async_apply_key,omitempty"`
	AsyncApplyMode *bool   `json:"async_apply_mode,omitempty" xml:"async_apply_mode,omitempty"`
	// This parameter is required.
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	OtaItemId        *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	OutOrderId       *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutReShopApplyId *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	// This parameter is required.
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	// This parameter is required.
	ReShopReasonCode *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
	TotalReShopFee           *int64  `json:"total_re_shop_fee,omitempty" xml:"total_re_shop_fee,omitempty"`
}

func (s IntlFlightReShopCreateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateShrinkRequest) GetAsyncApplyKey() *string {
	return s.AsyncApplyKey
}

func (s *IntlFlightReShopCreateShrinkRequest) GetAsyncApplyMode() *bool {
	return s.AsyncApplyMode
}

func (s *IntlFlightReShopCreateShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopCreateShrinkRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *IntlFlightReShopCreateShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopCreateShrinkRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopCreateShrinkRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopCreateShrinkRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopCreateShrinkRequest) GetSelectedPassengersShrink() *string {
	return s.SelectedPassengersShrink
}

func (s *IntlFlightReShopCreateShrinkRequest) GetTotalReShopFee() *int64 {
	return s.TotalReShopFee
}

func (s *IntlFlightReShopCreateShrinkRequest) SetAsyncApplyKey(v string) *IntlFlightReShopCreateShrinkRequest {
	s.AsyncApplyKey = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetAsyncApplyMode(v bool) *IntlFlightReShopCreateShrinkRequest {
	s.AsyncApplyMode = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetOrderId(v string) *IntlFlightReShopCreateShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetOtaItemId(v string) *IntlFlightReShopCreateShrinkRequest {
	s.OtaItemId = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetOutOrderId(v string) *IntlFlightReShopCreateShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetOutReShopApplyId(v string) *IntlFlightReShopCreateShrinkRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopCreateShrinkRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetReShopReasonCode(v string) *IntlFlightReShopCreateShrinkRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetSelectedPassengersShrink(v string) *IntlFlightReShopCreateShrinkRequest {
	s.SelectedPassengersShrink = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) SetTotalReShopFee(v int64) *IntlFlightReShopCreateShrinkRequest {
	s.TotalReShopFee = &v
	return s
}

func (s *IntlFlightReShopCreateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
