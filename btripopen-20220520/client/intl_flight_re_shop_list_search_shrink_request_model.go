// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopListSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopListSearchShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopListSearchShrinkRequest
	GetOutOrderId() *string
	SetOutWheelSearch(v bool) *IntlFlightReShopListSearchShrinkRequest
	GetOutWheelSearch() *bool
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopListSearchShrinkRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopListSearchShrinkRequest
	GetReShopReasonCode() *string
	SetSearchJourneysShrink(v string) *IntlFlightReShopListSearchShrinkRequest
	GetSearchJourneysShrink() *string
	SetSelectedPassengersShrink(v string) *IntlFlightReShopListSearchShrinkRequest
	GetSelectedPassengersShrink() *string
	SetToken(v string) *IntlFlightReShopListSearchShrinkRequest
	GetToken() *string
}

type IntlFlightReShopListSearchShrinkRequest struct {
	// The business travel order ID. Required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1002027205317939247
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// The external order ID.
	//
	// example:
	//
	// 3985893777358602240
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// Specifies whether external polling is enabled. When enabled, the external frontend controls polling to reduce the response time of a single search. Default value: false.
	//
	// example:
	//
	// false
	OutWheelSearch *bool `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// The rebooking group key returned by the consultation API.
	//
	// This parameter is required.
	//
	// example:
	//
	// edcac4f4c79d40ccb141ddb6da567e65
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	// The rebooking reason code.
	//
	// example:
	//
	// 0
	ReShopReasonCode *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// The list of journeys selected for rebooking.
	//
	// This parameter is required.
	SearchJourneysShrink *string `json:"search_journeys,omitempty" xml:"search_journeys,omitempty"`
	// The list of passengers selected for rebooking.
	//
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
	// The query record token for external polling.
	//
	// example:
	//
	// 0305b8203a7767626f911d97a91a9473
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightReShopListSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetOutWheelSearch() *bool {
	return s.OutWheelSearch
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetSearchJourneysShrink() *string {
	return s.SearchJourneysShrink
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetSelectedPassengersShrink() *string {
	return s.SelectedPassengersShrink
}

func (s *IntlFlightReShopListSearchShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetOrderId(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetOutOrderId(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetOutWheelSearch(v bool) *IntlFlightReShopListSearchShrinkRequest {
	s.OutWheelSearch = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetReShopReasonCode(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetSearchJourneysShrink(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.SearchJourneysShrink = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetSelectedPassengersShrink(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.SelectedPassengersShrink = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) SetToken(v string) *IntlFlightReShopListSearchShrinkRequest {
	s.Token = &v
	return s
}

func (s *IntlFlightReShopListSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
