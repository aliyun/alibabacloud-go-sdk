// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncApplyKey(v string) *IntlFlightReShopApplyRequest
	GetAsyncApplyKey() *string
	SetAsyncApplyMode(v bool) *IntlFlightReShopApplyRequest
	GetAsyncApplyMode() *bool
	SetOrderId(v string) *IntlFlightReShopApplyRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopApplyRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopApplyRequest
	GetOutReShopApplyId() *string
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopApplyRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopApplyRequest
	GetReShopReasonCode() *string
	SetSelectedJourneys(v []*IntlFlightReShopApplyRequestSelectedJourneys) *IntlFlightReShopApplyRequest
	GetSelectedJourneys() []*IntlFlightReShopApplyRequestSelectedJourneys
	SetSelectedPassengers(v []*IntlFlightReShopApplyRequestSelectedPassengers) *IntlFlightReShopApplyRequest
	GetSelectedPassengers() []*IntlFlightReShopApplyRequestSelectedPassengers
	SetUserIntentionMemo(v string) *IntlFlightReShopApplyRequest
	GetUserIntentionMemo() *string
}

type IntlFlightReShopApplyRequest struct {
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
	SelectedJourneys []*IntlFlightReShopApplyRequestSelectedJourneys `json:"selected_journeys,omitempty" xml:"selected_journeys,omitempty" type:"Repeated"`
	// This parameter is required.
	SelectedPassengers []*IntlFlightReShopApplyRequestSelectedPassengers `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty" type:"Repeated"`
	UserIntentionMemo  *string                                           `json:"user_intention_memo,omitempty" xml:"user_intention_memo,omitempty"`
}

func (s IntlFlightReShopApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyRequest) GetAsyncApplyKey() *string {
	return s.AsyncApplyKey
}

func (s *IntlFlightReShopApplyRequest) GetAsyncApplyMode() *bool {
	return s.AsyncApplyMode
}

func (s *IntlFlightReShopApplyRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopApplyRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopApplyRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopApplyRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopApplyRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopApplyRequest) GetSelectedJourneys() []*IntlFlightReShopApplyRequestSelectedJourneys {
	return s.SelectedJourneys
}

func (s *IntlFlightReShopApplyRequest) GetSelectedPassengers() []*IntlFlightReShopApplyRequestSelectedPassengers {
	return s.SelectedPassengers
}

func (s *IntlFlightReShopApplyRequest) GetUserIntentionMemo() *string {
	return s.UserIntentionMemo
}

func (s *IntlFlightReShopApplyRequest) SetAsyncApplyKey(v string) *IntlFlightReShopApplyRequest {
	s.AsyncApplyKey = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetAsyncApplyMode(v bool) *IntlFlightReShopApplyRequest {
	s.AsyncApplyMode = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetOrderId(v string) *IntlFlightReShopApplyRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetOutOrderId(v string) *IntlFlightReShopApplyRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetOutReShopApplyId(v string) *IntlFlightReShopApplyRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopApplyRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetReShopReasonCode(v string) *IntlFlightReShopApplyRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetSelectedJourneys(v []*IntlFlightReShopApplyRequestSelectedJourneys) *IntlFlightReShopApplyRequest {
	s.SelectedJourneys = v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetSelectedPassengers(v []*IntlFlightReShopApplyRequestSelectedPassengers) *IntlFlightReShopApplyRequest {
	s.SelectedPassengers = v
	return s
}

func (s *IntlFlightReShopApplyRequest) SetUserIntentionMemo(v string) *IntlFlightReShopApplyRequest {
	s.UserIntentionMemo = &v
	return s
}

func (s *IntlFlightReShopApplyRequest) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopApplyRequestSelectedJourneys struct {
	// This parameter is required.
	//
	// example:
	//
	// DLC
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TSN
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-10-10
	IntentDate *string `json:"intent_date,omitempty" xml:"intent_date,omitempty"`
	// This parameter is required.
	SelectedFlights []*IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopApplyRequestSelectedJourneys) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyRequestSelectedJourneys) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) GetIntentDate() *string {
	return s.IntentDate
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) GetSelectedFlights() []*IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) SetArrCityCode(v string) *IntlFlightReShopApplyRequestSelectedJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) SetDepCityCode(v string) *IntlFlightReShopApplyRequestSelectedJourneys {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) SetIntentDate(v string) *IntlFlightReShopApplyRequestSelectedJourneys {
	s.IntentDate = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) SetSelectedFlights(v []*IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) *IntlFlightReShopApplyRequestSelectedJourneys {
	s.SelectedFlights = v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneys) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights struct {
	// This parameter is required.
	//
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KN6728HGHPKX0725
	SegmentKey *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
}

func (s IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) SetArrCityCode(v string) *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) SetDepCityCode(v string) *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) SetSegmentKey(v string) *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopApplyRequestSelectedPassengers struct {
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightReShopApplyRequestSelectedPassengers) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyRequestSelectedPassengers) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyRequestSelectedPassengers) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightReShopApplyRequestSelectedPassengers) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopApplyRequestSelectedPassengers) SetFullName(v string) *IntlFlightReShopApplyRequestSelectedPassengers {
	s.FullName = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedPassengers) SetPassengerId(v int64) *IntlFlightReShopApplyRequestSelectedPassengers {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopApplyRequestSelectedPassengers) Validate() error {
	return dara.Validate(s)
}
