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
	// The key for the asynchronous application.
	//
	// example:
	//
	// asyncKey_2390u230slgw023
	AsyncApplyKey *string `json:"async_apply_key,omitempty" xml:"async_apply_key,omitempty"`
	// Specifies whether to use the asynchronous commit pattern. If asynchronous commit is used, only a key is returned before the application result is available.
	//
	// example:
	//
	// true
	AsyncApplyMode *bool `json:"async_apply_mode,omitempty" xml:"async_apply_mode,omitempty"`
	// The business travel order ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1017035199702438072
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// The ID of the rebooking product.
	//
	// This parameter is required.
	//
	// example:
	//
	// d01eb358456b4ba38eb4d8f1499186da_0
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// The external order ID.
	//
	// example:
	//
	// 3881646464813400064
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// The external rebooking application ID.
	//
	// example:
	//
	// JPM20241024354
	OutReShopApplyId *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	// The rebooking group key returned by the inquiry operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// edcac4f4c79d40ccb141ddb6da567e65
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	// The rebooking reason code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ReShopReasonCode *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// The list of passengers selected for rebooking.
	//
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
	// The total rebooking fee (excluding the service fee), in cents.
	//
	//      	- Total rebooking fee = cabin upgrade fee + handling fee + tax difference (applicable to international flights).
	//
	//      	- Pass in this parameter when fees are incurred to verify whether the price has changed.
	//
	// example:
	//
	// 1400
	TotalReShopFee *int64 `json:"total_re_shop_fee,omitempty" xml:"total_re_shop_fee,omitempty"`
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
