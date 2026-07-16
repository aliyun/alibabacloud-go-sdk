// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopOtaSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetOutOrderId() *string
	SetOutWheelSearch(v bool) *IntlFlightReShopOtaSearchShrinkRequest
	GetOutWheelSearch() *bool
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetReShopReasonCode() *string
	SetSearchJourneysShrink(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetSearchJourneysShrink() *string
	SetSelectedPassengersShrink(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetSelectedPassengersShrink() *string
	SetToken(v string) *IntlFlightReShopOtaSearchShrinkRequest
	GetToken() *string
}

type IntlFlightReShopOtaSearchShrinkRequest struct {
	// This parameter is required.
	OrderId        *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId     *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutWheelSearch *bool   `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// This parameter is required.
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	ReShopReasonCode         *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// This parameter is required.
	SearchJourneysShrink *string `json:"search_journeys,omitempty" xml:"search_journeys,omitempty"`
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
	Token                    *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightReShopOtaSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetOutWheelSearch() *bool {
	return s.OutWheelSearch
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetSearchJourneysShrink() *string {
	return s.SearchJourneysShrink
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetSelectedPassengersShrink() *string {
	return s.SelectedPassengersShrink
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetOrderId(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetOutOrderId(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetOutWheelSearch(v bool) *IntlFlightReShopOtaSearchShrinkRequest {
	s.OutWheelSearch = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetReShopReasonCode(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetSearchJourneysShrink(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.SearchJourneysShrink = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetSelectedPassengersShrink(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.SelectedPassengersShrink = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) SetToken(v string) *IntlFlightReShopOtaSearchShrinkRequest {
	s.Token = &v
	return s
}

func (s *IntlFlightReShopOtaSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
