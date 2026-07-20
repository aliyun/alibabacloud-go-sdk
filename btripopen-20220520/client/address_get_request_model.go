// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressGetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *AddressGetRequest
	GetActionType() *int32
	SetArrCityCode(v string) *AddressGetRequest
	GetArrCityCode() *string
	SetArrCityName(v string) *AddressGetRequest
	GetArrCityName() *string
	SetCarScenesCode(v string) *AddressGetRequest
	GetCarScenesCode() *string
	SetDepCityCode(v string) *AddressGetRequest
	GetDepCityCode() *string
	SetDepCityName(v string) *AddressGetRequest
	GetDepCityName() *string
	SetDepDate(v string) *AddressGetRequest
	GetDepDate() *string
	SetItineraryId(v string) *AddressGetRequest
	GetItineraryId() *string
	SetMiddlePage(v int32) *AddressGetRequest
	GetMiddlePage() *int32
	SetOrderId(v string) *AddressGetRequest
	GetOrderId() *string
	SetPhone(v string) *AddressGetRequest
	GetPhone() *string
	SetSessionParameters(v string) *AddressGetRequest
	GetSessionParameters() *string
	SetSubCorpId(v string) *AddressGetRequest
	GetSubCorpId() *string
	SetTaobaoCallbackUrl(v string) *AddressGetRequest
	GetTaobaoCallbackUrl() *string
	SetThirdpartApplyId(v string) *AddressGetRequest
	GetThirdpartApplyId() *string
	SetTravelerId(v string) *AddressGetRequest
	GetTravelerId() *string
	SetType(v int32) *AddressGetRequest
	GetType() *int32
	SetUseBookingProxy(v int32) *AddressGetRequest
	GetUseBookingProxy() *int32
	SetUserId(v string) *AddressGetRequest
	GetUserId() *string
}

type AddressGetRequest struct {
	// The redirect page type. For illustrations of each page, refer to [How to implement SSO redirection - Appendix](https://openapi.alibtrip.com/doc/toDocDetail?docId=4746411).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ActionType *int32 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// The three-letter code of the arrival city.
	//
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// The arrival city name.
	//
	// example:
	//
	// 北京
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// The car service scenario.
	//
	// example:
	//
	// TRAVEL
	CarScenesCode *string `json:"car_scenes_code,omitempty" xml:"car_scenes_code,omitempty"`
	// The three-letter code of the departure city.
	//
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// The departure city name.
	//
	// example:
	//
	// 杭州
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// The departure date.
	//
	// example:
	//
	// 2023-02-26
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// The itinerary ID.
	//
	// - When the redirect page is the business travel booking page (`action_type = 1`), you can optionally pass this parameter to quickly redirect to the booking page of the category associated with the itinerary.
	//
	// - The itinerary ID must have been submitted to the Alibaba Business Travel system through the [Create a business trip approval](https://openapi.alibtrip.com/doc/toDocDetail?docId=4929938) operation.
	//
	// example:
	//
	// 460e********5d78
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// Specifies whether to skip the booking intermediate page.
	//
	// 1. Set this parameter to 2 to skip the booking intermediate page. When skipping the intermediate page, the **itinerary_id*	- parameter is required. If this parameter is empty or set to a value other than 2, the intermediate page is not skipped.
	//
	// 2. This parameter is available when the redirect page is the **H5 booking page*	- (`action_type = 1`) and the category is **flight*	- (`type = 1`) or **train*	- (`type = 2`).
	//
	// example:
	//
	// 1
	MiddlePage *int32 `json:"middle_page,omitempty" xml:"middle_page,omitempty"`
	// The order ID. This parameter is required when the redirect page type is the specified order details page on either platform (`action_type = 11 or 12`).
	//
	// example:
	//
	// 1002************464
	OrderId *string `json:"order_Id,omitempty" xml:"order_Id,omitempty"`
	// The contact phone number, typically used for car service scenarios.
	//
	// example:
	//
	// 131****8888
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// Session parameters. The format must be a JSON string where both keys and values are strings.
	//
	// Example: "{\\"returnURL\\":\\"https://open.alibtrip.com/\\"}"
	SessionParameters *string `json:"session_parameters,omitempty" xml:"session_parameters,omitempty"`
	// The sub-enterprise ID. Pass this parameter to redirect to the business page of the specified sub-enterprise.
	//
	// - **View permissions**: Only enterprise administrators have view permissions.
	//
	// - **Path to obtain**: Enterprise management console > Parent-child account management > Account management > Sub-account management > Company ID.
	//
	// example:
	//
	// btrip01******00
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// The redirect URL after Taobao account binding.
	//
	// example:
	//
	// https://example.com
	TaobaoCallbackUrl *string `json:"taobao_callback_url,omitempty" xml:"taobao_callback_url,omitempty"`
	// The third-party approval ID.
	//
	// example:
	//
	// TP00097732
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// The ID of the actual traveler (the person being booked for).
	//
	// example:
	//
	// user_1234
	TravelerId *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// The business type. This parameter is required when the redirect page is the **booking page*	- (`action_type = 1`) or the **order view page*	- (`action_type = 2`).
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to use proxy booking mode.
	//
	// - The proxy booking page is accessible only when this parameter is set to 1.
	//
	// example:
	//
	// 1
	UseBookingProxy *int32 `json:"use_booking_proxy,omitempty" xml:"use_booking_proxy,omitempty"`
	// The employee ID. The employee must be registered in the business travel system before you pass this parameter. Otherwise, the call fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_1234
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s AddressGetRequest) String() string {
	return dara.Prettify(s)
}

func (s AddressGetRequest) GoString() string {
	return s.String()
}

func (s *AddressGetRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *AddressGetRequest) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *AddressGetRequest) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *AddressGetRequest) GetCarScenesCode() *string {
	return s.CarScenesCode
}

func (s *AddressGetRequest) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *AddressGetRequest) GetDepCityName() *string {
	return s.DepCityName
}

func (s *AddressGetRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *AddressGetRequest) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *AddressGetRequest) GetMiddlePage() *int32 {
	return s.MiddlePage
}

func (s *AddressGetRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *AddressGetRequest) GetPhone() *string {
	return s.Phone
}

func (s *AddressGetRequest) GetSessionParameters() *string {
	return s.SessionParameters
}

func (s *AddressGetRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *AddressGetRequest) GetTaobaoCallbackUrl() *string {
	return s.TaobaoCallbackUrl
}

func (s *AddressGetRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *AddressGetRequest) GetTravelerId() *string {
	return s.TravelerId
}

func (s *AddressGetRequest) GetType() *int32 {
	return s.Type
}

func (s *AddressGetRequest) GetUseBookingProxy() *int32 {
	return s.UseBookingProxy
}

func (s *AddressGetRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddressGetRequest) SetActionType(v int32) *AddressGetRequest {
	s.ActionType = &v
	return s
}

func (s *AddressGetRequest) SetArrCityCode(v string) *AddressGetRequest {
	s.ArrCityCode = &v
	return s
}

func (s *AddressGetRequest) SetArrCityName(v string) *AddressGetRequest {
	s.ArrCityName = &v
	return s
}

func (s *AddressGetRequest) SetCarScenesCode(v string) *AddressGetRequest {
	s.CarScenesCode = &v
	return s
}

func (s *AddressGetRequest) SetDepCityCode(v string) *AddressGetRequest {
	s.DepCityCode = &v
	return s
}

func (s *AddressGetRequest) SetDepCityName(v string) *AddressGetRequest {
	s.DepCityName = &v
	return s
}

func (s *AddressGetRequest) SetDepDate(v string) *AddressGetRequest {
	s.DepDate = &v
	return s
}

func (s *AddressGetRequest) SetItineraryId(v string) *AddressGetRequest {
	s.ItineraryId = &v
	return s
}

func (s *AddressGetRequest) SetMiddlePage(v int32) *AddressGetRequest {
	s.MiddlePage = &v
	return s
}

func (s *AddressGetRequest) SetOrderId(v string) *AddressGetRequest {
	s.OrderId = &v
	return s
}

func (s *AddressGetRequest) SetPhone(v string) *AddressGetRequest {
	s.Phone = &v
	return s
}

func (s *AddressGetRequest) SetSessionParameters(v string) *AddressGetRequest {
	s.SessionParameters = &v
	return s
}

func (s *AddressGetRequest) SetSubCorpId(v string) *AddressGetRequest {
	s.SubCorpId = &v
	return s
}

func (s *AddressGetRequest) SetTaobaoCallbackUrl(v string) *AddressGetRequest {
	s.TaobaoCallbackUrl = &v
	return s
}

func (s *AddressGetRequest) SetThirdpartApplyId(v string) *AddressGetRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *AddressGetRequest) SetTravelerId(v string) *AddressGetRequest {
	s.TravelerId = &v
	return s
}

func (s *AddressGetRequest) SetType(v int32) *AddressGetRequest {
	s.Type = &v
	return s
}

func (s *AddressGetRequest) SetUseBookingProxy(v int32) *AddressGetRequest {
	s.UseBookingProxy = &v
	return s
}

func (s *AddressGetRequest) SetUserId(v string) *AddressGetRequest {
	s.UserId = &v
	return s
}

func (s *AddressGetRequest) Validate() error {
	return dara.Validate(s)
}
