// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncApplyKey(v string) *IntlFlightReShopApplyShrinkRequest
	GetAsyncApplyKey() *string
	SetAsyncApplyMode(v bool) *IntlFlightReShopApplyShrinkRequest
	GetAsyncApplyMode() *bool
	SetOrderId(v string) *IntlFlightReShopApplyShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopApplyShrinkRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopApplyShrinkRequest
	GetOutReShopApplyId() *string
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopApplyShrinkRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopApplyShrinkRequest
	GetReShopReasonCode() *string
	SetSelectedJourneysShrink(v string) *IntlFlightReShopApplyShrinkRequest
	GetSelectedJourneysShrink() *string
	SetSelectedPassengersShrink(v string) *IntlFlightReShopApplyShrinkRequest
	GetSelectedPassengersShrink() *string
	SetUserIntentionMemo(v string) *IntlFlightReShopApplyShrinkRequest
	GetUserIntentionMemo() *string
}

type IntlFlightReShopApplyShrinkRequest struct {
	// example:
	//
	// asyncKey_2390u230slgw023
	AsyncApplyKey *string `json:"async_apply_key,omitempty" xml:"async_apply_key,omitempty"`
	// example:
	//
	// true
	AsyncApplyMode *bool `json:"async_apply_mode,omitempty" xml:"async_apply_mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1017035199907040165
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// D1736316966048SC4877
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// JPM20241024354
	OutReShopApplyId *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
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
	ReShopReasonCode *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// This parameter is required.
	SelectedJourneysShrink *string `json:"selected_journeys,omitempty" xml:"selected_journeys,omitempty"`
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
	UserIntentionMemo        *string `json:"user_intention_memo,omitempty" xml:"user_intention_memo,omitempty"`
}

func (s IntlFlightReShopApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyShrinkRequest) GetAsyncApplyKey() *string {
	return s.AsyncApplyKey
}

func (s *IntlFlightReShopApplyShrinkRequest) GetAsyncApplyMode() *bool {
	return s.AsyncApplyMode
}

func (s *IntlFlightReShopApplyShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopApplyShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopApplyShrinkRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopApplyShrinkRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopApplyShrinkRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopApplyShrinkRequest) GetSelectedJourneysShrink() *string {
	return s.SelectedJourneysShrink
}

func (s *IntlFlightReShopApplyShrinkRequest) GetSelectedPassengersShrink() *string {
	return s.SelectedPassengersShrink
}

func (s *IntlFlightReShopApplyShrinkRequest) GetUserIntentionMemo() *string {
	return s.UserIntentionMemo
}

func (s *IntlFlightReShopApplyShrinkRequest) SetAsyncApplyKey(v string) *IntlFlightReShopApplyShrinkRequest {
	s.AsyncApplyKey = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetAsyncApplyMode(v bool) *IntlFlightReShopApplyShrinkRequest {
	s.AsyncApplyMode = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetOrderId(v string) *IntlFlightReShopApplyShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetOutOrderId(v string) *IntlFlightReShopApplyShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetOutReShopApplyId(v string) *IntlFlightReShopApplyShrinkRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopApplyShrinkRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetReShopReasonCode(v string) *IntlFlightReShopApplyShrinkRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetSelectedJourneysShrink(v string) *IntlFlightReShopApplyShrinkRequest {
	s.SelectedJourneysShrink = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetSelectedPassengersShrink(v string) *IntlFlightReShopApplyShrinkRequest {
	s.SelectedPassengersShrink = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) SetUserIntentionMemo(v string) *IntlFlightReShopApplyShrinkRequest {
	s.UserIntentionMemo = &v
	return s
}

func (s *IntlFlightReShopApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
