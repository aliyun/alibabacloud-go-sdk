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
	// This parameter is required.
	//
	// example:
	//
	// 1
	ActionType    *int32  `json:"action_type,omitempty" xml:"action_type,omitempty"`
	ArrCityCode   *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName   *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	CarScenesCode *string `json:"car_scenes_code,omitempty" xml:"car_scenes_code,omitempty"`
	DepCityCode   *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName   *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepDate       *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// 460e254b5a5b4bd0801744a2790e5d78
	ItineraryId       *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	MiddlePage        *int32  `json:"middle_page,omitempty" xml:"middle_page,omitempty"`
	OrderId           *string `json:"order_Id,omitempty" xml:"order_Id,omitempty"`
	Phone             *string `json:"phone,omitempty" xml:"phone,omitempty"`
	SessionParameters *string `json:"session_parameters,omitempty" xml:"session_parameters,omitempty"`
	SubCorpId         *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// example:
	//
	// https://alibtrip.open.com
	TaobaoCallbackUrl *string `json:"taobao_callback_url,omitempty" xml:"taobao_callback_url,omitempty"`
	ThirdpartApplyId  *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	TravelerId        *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// 1
	Type            *int32 `json:"type,omitempty" xml:"type,omitempty"`
	UseBookingProxy *int32 `json:"use_booking_proxy,omitempty" xml:"use_booking_proxy,omitempty"`
	// This parameter is required.
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
