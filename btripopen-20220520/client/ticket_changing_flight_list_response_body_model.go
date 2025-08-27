// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingFlightListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TicketChangingFlightListResponseBody
	GetCode() *string
	SetMessage(v string) *TicketChangingFlightListResponseBody
	GetMessage() *string
	SetModule(v *TicketChangingFlightListResponseBodyModule) *TicketChangingFlightListResponseBody
	GetModule() *TicketChangingFlightListResponseBodyModule
	SetRequestId(v string) *TicketChangingFlightListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketChangingFlightListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TicketChangingFlightListResponseBody
	GetTraceId() *string
}

type TicketChangingFlightListResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TicketChangingFlightListResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TicketChangingFlightListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBody) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBody) GetCode() *string {
	return s.Code
}

func (s *TicketChangingFlightListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TicketChangingFlightListResponseBody) GetModule() *TicketChangingFlightListResponseBodyModule {
	return s.Module
}

func (s *TicketChangingFlightListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketChangingFlightListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketChangingFlightListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TicketChangingFlightListResponseBody) SetCode(v string) *TicketChangingFlightListResponseBody {
	s.Code = &v
	return s
}

func (s *TicketChangingFlightListResponseBody) SetMessage(v string) *TicketChangingFlightListResponseBody {
	s.Message = &v
	return s
}

func (s *TicketChangingFlightListResponseBody) SetModule(v *TicketChangingFlightListResponseBodyModule) *TicketChangingFlightListResponseBody {
	s.Module = v
	return s
}

func (s *TicketChangingFlightListResponseBody) SetRequestId(v string) *TicketChangingFlightListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketChangingFlightListResponseBody) SetSuccess(v bool) *TicketChangingFlightListResponseBody {
	s.Success = &v
	return s
}

func (s *TicketChangingFlightListResponseBody) SetTraceId(v string) *TicketChangingFlightListResponseBody {
	s.TraceId = &v
	return s
}

func (s *TicketChangingFlightListResponseBody) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModule struct {
	FlightInfoList []*TicketChangingFlightListResponseBodyModuleFlightInfoList `json:"flight_info_list,omitempty" xml:"flight_info_list,omitempty" type:"Repeated"`
}

func (s TicketChangingFlightListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModule) GetFlightInfoList() []*TicketChangingFlightListResponseBodyModuleFlightInfoList {
	return s.FlightInfoList
}

func (s *TicketChangingFlightListResponseBodyModule) SetFlightInfoList(v []*TicketChangingFlightListResponseBodyModuleFlightInfoList) *TicketChangingFlightListResponseBodyModule {
	s.FlightInfoList = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoList struct {
	AirlineInfo    *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	CabinList      []*TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList    `json:"cabin_list,omitempty" xml:"cabin_list,omitempty" type:"Repeated"`
	// example:
	//
	// CA1704
	CarrierAirline *string `json:"carrier_airline,omitempty" xml:"carrier_airline,omitempty"`
	// example:
	//
	// CA1704
	CarrierNo      *string                                                                 `json:"carrier_no,omitempty" xml:"carrier_no,omitempty"`
	DepAirportInfo *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// CA1704
	FlightNo   *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightSize *string `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// true
	IsProtocol *bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	// example:
	//
	// false
	IsShare *bool `json:"is_share,omitempty" xml:"is_share,omitempty"`
	// example:
	//
	// false
	IsStop *bool `json:"is_stop,omitempty" xml:"is_stop,omitempty"`
	// example:
	//
	// Y
	LowestCabin *string `json:"lowest_cabin,omitempty" xml:"lowest_cabin,omitempty"`
	// example:
	//
	// Y
	LowestCabinClass *string                                                                     `json:"lowest_cabin_class,omitempty" xml:"lowest_cabin_class,omitempty"`
	LowestCabinDesc  *string                                                                     `json:"lowest_cabin_desc,omitempty" xml:"lowest_cabin_desc,omitempty"`
	LowestCabinNum   *string                                                                     `json:"lowest_cabin_num,omitempty" xml:"lowest_cabin_num,omitempty"`
	LowestCabinPrice []*TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice `json:"lowest_cabin_price,omitempty" xml:"lowest_cabin_price,omitempty" type:"Repeated"`
	MealDesc         *string                                                                     `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ModifyFlightArrTime *string `json:"modify_flight_arr_time,omitempty" xml:"modify_flight_arr_time,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ModifyFlightDepDate *string `json:"modify_flight_dep_date,omitempty" xml:"modify_flight_dep_date,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ModifyFlightDepTime *string `json:"modify_flight_dep_time,omitempty" xml:"modify_flight_dep_time,omitempty"`
	// example:
	//
	// d1fb9e0a794f45e1b762d36ff1d17zz
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// HGH
	StopCity *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetAirlineInfo() *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo {
	return s.AirlineInfo
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetArrAirportInfo() *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetCabinList() []*TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	return s.CabinList
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetCarrierAirline() *string {
	return s.CarrierAirline
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetCarrierNo() *string {
	return s.CarrierNo
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetDepAirportInfo() *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetFlightType() *string {
	return s.FlightType
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetIsShare() *bool {
	return s.IsShare
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetIsStop() *bool {
	return s.IsStop
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetLowestCabin() *string {
	return s.LowestCabin
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetLowestCabinClass() *string {
	return s.LowestCabinClass
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetLowestCabinDesc() *string {
	return s.LowestCabinDesc
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetLowestCabinNum() *string {
	return s.LowestCabinNum
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetLowestCabinPrice() []*TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice {
	return s.LowestCabinPrice
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetModifyFlightArrTime() *string {
	return s.ModifyFlightArrTime
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetModifyFlightDepDate() *string {
	return s.ModifyFlightDepDate
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetModifyFlightDepTime() *string {
	return s.ModifyFlightDepTime
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetSessionId() *string {
	return s.SessionId
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetStopCity() *string {
	return s.StopCity
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetAirlineInfo(v *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.AirlineInfo = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetArrAirportInfo(v *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.ArrAirportInfo = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetCabinList(v []*TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.CabinList = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetCarrierAirline(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.CarrierAirline = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetCarrierNo(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.CarrierNo = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetDepAirportInfo(v *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.DepAirportInfo = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetFlightNo(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.FlightNo = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetFlightSize(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.FlightSize = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetFlightType(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.FlightType = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetIsProtocol(v bool) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.IsProtocol = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetIsShare(v bool) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.IsShare = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetIsStop(v bool) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.IsStop = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetLowestCabin(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.LowestCabin = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetLowestCabinClass(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.LowestCabinClass = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetLowestCabinDesc(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.LowestCabinDesc = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetLowestCabinNum(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.LowestCabinNum = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetLowestCabinPrice(v []*TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.LowestCabinPrice = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetMealDesc(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.MealDesc = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetModifyFlightArrTime(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.ModifyFlightArrTime = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetModifyFlightDepDate(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.ModifyFlightDepDate = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetModifyFlightDepTime(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.ModifyFlightDepTime = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetSessionId(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.SessionId = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetStopArrTime(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.StopArrTime = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetStopCity(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.StopCity = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) SetStopDepTime(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoList {
	s.StopDepTime = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) SetAirlineCode(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) SetAirlineName(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) SetAirlineSimpleName(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo {
	s.AirlineSimpleName = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo struct {
	// example:
	//
	// HGH
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// HGH
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// T4
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) SetAirportCode(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) SetAirportName(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) SetCityCode(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo {
	s.CityCode = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) SetCityName(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo {
	s.CityName = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) SetTerminal(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList struct {
	// example:
	//
	// G
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinDesc  *string `json:"cabin_desc,omitempty" xml:"cabin_desc,omitempty"`
	// example:
	//
	// 80
	CabinDiscount *int32 `json:"cabin_discount,omitempty" xml:"cabin_discount,omitempty"`
	// example:
	//
	// G
	ChildCabin *string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
	// example:
	//
	// A
	LeftNum         *string                                                                             `json:"left_num,omitempty" xml:"left_num,omitempty"`
	ModifyPriceList []*TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList `json:"modify_price_list,omitempty" xml:"modify_price_list,omitempty" type:"Repeated"`
	// example:
	//
	// "360379a11ee84e9aa011baa41b758fe6"
	OtaItemid *string `json:"ota_itemid,omitempty" xml:"ota_itemid,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetCabin() *string {
	return s.Cabin
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetCabinDesc() *string {
	return s.CabinDesc
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetCabinDiscount() *int32 {
	return s.CabinDiscount
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetChildCabin() *string {
	return s.ChildCabin
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetLeftNum() *string {
	return s.LeftNum
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetModifyPriceList() []*TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	return s.ModifyPriceList
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) GetOtaItemid() *string {
	return s.OtaItemid
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetCabin(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.Cabin = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetCabinClass(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.CabinClass = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetCabinDesc(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.CabinDesc = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetCabinDiscount(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.CabinDiscount = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetChildCabin(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.ChildCabin = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetLeftNum(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.LeftNum = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetModifyPriceList(v []*TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.ModifyPriceList = v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) SetOtaItemid(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList {
	s.OtaItemid = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList struct {
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 100
	UpgradeFee *int32 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// example:
	//
	// 100
	UpgradePrice *int32 `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetUpgradeFee() *int32 {
	return s.UpgradeFee
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetUpgradePrice() *int32 {
	return s.UpgradePrice
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetPassengerType(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.PassengerType = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetTicketPrice(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.TicketPrice = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetUpgradeFee(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.UpgradeFee = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetUpgradePrice(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.UpgradePrice = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListCabinListModifyPriceList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo struct {
	// example:
	//
	// PEK
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// BJS
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) SetAirportCode(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) SetAirportName(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) SetCityCode(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo {
	s.CityCode = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) SetCityName(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo {
	s.CityName = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) SetTerminal(v string) *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice struct {
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 100
	UpgradeFee *int32 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// example:
	//
	// 100
	UpgradePrice *int32 `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) GetUpgradeFee() *int32 {
	return s.UpgradeFee
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) GetUpgradePrice() *int32 {
	return s.UpgradePrice
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) SetPassengerType(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.PassengerType = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) SetTicketPrice(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.TicketPrice = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) SetUpgradeFee(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.UpgradeFee = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) SetUpgradePrice(v int32) *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.UpgradePrice = &v
	return s
}

func (s *TicketChangingFlightListResponseBodyModuleFlightInfoListLowestCabinPrice) Validate() error {
	return dara.Validate(s)
}
