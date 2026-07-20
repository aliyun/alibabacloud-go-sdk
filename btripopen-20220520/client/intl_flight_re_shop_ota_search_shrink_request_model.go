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
	// The business travel order ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1017089206331988332
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// The external order ID.
	//
	// example:
	//
	// JP2024071800000002
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// Specifies whether to enable external polling. When enabled, the external frontend controls polling, which reduces the response time (RT) of a single search. Default value: false.
	//
	// example:
	//
	// true
	OutWheelSearch *bool `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// The rebooking group key returned by the consultation operation.
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
	// The search journeys. Only pass the journeys to be rebooked.
	//
	// This parameter is required.
	SearchJourneysShrink *string `json:"search_journeys,omitempty" xml:"search_journeys,omitempty"`
	// The list of selected passengers for rebooking.
	//
	// This parameter is required.
	SelectedPassengersShrink *string `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty"`
	// The query record token used for external polling.
	//
	// example:
	//
	// 9960b412-cc05-4d10-b570-93372d816807
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
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
