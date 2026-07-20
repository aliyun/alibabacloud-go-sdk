// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncApplyKey(v string) *IntlFlightReShopCreateRequest
	GetAsyncApplyKey() *string
	SetAsyncApplyMode(v bool) *IntlFlightReShopCreateRequest
	GetAsyncApplyMode() *bool
	SetOrderId(v string) *IntlFlightReShopCreateRequest
	GetOrderId() *string
	SetOtaItemId(v string) *IntlFlightReShopCreateRequest
	GetOtaItemId() *string
	SetOutOrderId(v string) *IntlFlightReShopCreateRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopCreateRequest
	GetOutReShopApplyId() *string
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopCreateRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopCreateRequest
	GetReShopReasonCode() *string
	SetSelectedPassengers(v []*IntlFlightReShopCreateRequestSelectedPassengers) *IntlFlightReShopCreateRequest
	GetSelectedPassengers() []*IntlFlightReShopCreateRequestSelectedPassengers
	SetTotalReShopFee(v int64) *IntlFlightReShopCreateRequest
	GetTotalReShopFee() *int64
}

type IntlFlightReShopCreateRequest struct {
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
	SelectedPassengers []*IntlFlightReShopCreateRequestSelectedPassengers `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty" type:"Repeated"`
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

func (s IntlFlightReShopCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateRequest) GetAsyncApplyKey() *string {
	return s.AsyncApplyKey
}

func (s *IntlFlightReShopCreateRequest) GetAsyncApplyMode() *bool {
	return s.AsyncApplyMode
}

func (s *IntlFlightReShopCreateRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopCreateRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *IntlFlightReShopCreateRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopCreateRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopCreateRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopCreateRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopCreateRequest) GetSelectedPassengers() []*IntlFlightReShopCreateRequestSelectedPassengers {
	return s.SelectedPassengers
}

func (s *IntlFlightReShopCreateRequest) GetTotalReShopFee() *int64 {
	return s.TotalReShopFee
}

func (s *IntlFlightReShopCreateRequest) SetAsyncApplyKey(v string) *IntlFlightReShopCreateRequest {
	s.AsyncApplyKey = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetAsyncApplyMode(v bool) *IntlFlightReShopCreateRequest {
	s.AsyncApplyMode = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetOrderId(v string) *IntlFlightReShopCreateRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetOtaItemId(v string) *IntlFlightReShopCreateRequest {
	s.OtaItemId = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetOutOrderId(v string) *IntlFlightReShopCreateRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetOutReShopApplyId(v string) *IntlFlightReShopCreateRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopCreateRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetReShopReasonCode(v string) *IntlFlightReShopCreateRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetSelectedPassengers(v []*IntlFlightReShopCreateRequestSelectedPassengers) *IntlFlightReShopCreateRequest {
	s.SelectedPassengers = v
	return s
}

func (s *IntlFlightReShopCreateRequest) SetTotalReShopFee(v int64) *IntlFlightReShopCreateRequest {
	s.TotalReShopFee = &v
	return s
}

func (s *IntlFlightReShopCreateRequest) Validate() error {
	if s.SelectedPassengers != nil {
		for _, item := range s.SelectedPassengers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopCreateRequestSelectedPassengers struct {
	// The full name of the passenger.
	//
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// The ID of the passenger.
	//
	// example:
	//
	// 1000001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightReShopCreateRequestSelectedPassengers) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateRequestSelectedPassengers) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateRequestSelectedPassengers) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightReShopCreateRequestSelectedPassengers) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopCreateRequestSelectedPassengers) SetFullName(v string) *IntlFlightReShopCreateRequestSelectedPassengers {
	s.FullName = &v
	return s
}

func (s *IntlFlightReShopCreateRequestSelectedPassengers) SetPassengerId(v int64) *IntlFlightReShopCreateRequestSelectedPassengers {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopCreateRequestSelectedPassengers) Validate() error {
	return dara.Validate(s)
}
