// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingEnquiryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TicketChangingEnquiryResponseBody
	GetCode() *string
	SetMessage(v string) *TicketChangingEnquiryResponseBody
	GetMessage() *string
	SetModule(v *TicketChangingEnquiryResponseBodyModule) *TicketChangingEnquiryResponseBody
	GetModule() *TicketChangingEnquiryResponseBodyModule
	SetRequestId(v string) *TicketChangingEnquiryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketChangingEnquiryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TicketChangingEnquiryResponseBody
	GetTraceId() *string
}

type TicketChangingEnquiryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TicketChangingEnquiryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
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

func (s TicketChangingEnquiryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBody) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TicketChangingEnquiryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TicketChangingEnquiryResponseBody) GetModule() *TicketChangingEnquiryResponseBodyModule {
	return s.Module
}

func (s *TicketChangingEnquiryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketChangingEnquiryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketChangingEnquiryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TicketChangingEnquiryResponseBody) SetCode(v string) *TicketChangingEnquiryResponseBody {
	s.Code = &v
	return s
}

func (s *TicketChangingEnquiryResponseBody) SetMessage(v string) *TicketChangingEnquiryResponseBody {
	s.Message = &v
	return s
}

func (s *TicketChangingEnquiryResponseBody) SetModule(v *TicketChangingEnquiryResponseBodyModule) *TicketChangingEnquiryResponseBody {
	s.Module = v
	return s
}

func (s *TicketChangingEnquiryResponseBody) SetRequestId(v string) *TicketChangingEnquiryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketChangingEnquiryResponseBody) SetSuccess(v bool) *TicketChangingEnquiryResponseBody {
	s.Success = &v
	return s
}

func (s *TicketChangingEnquiryResponseBody) SetTraceId(v string) *TicketChangingEnquiryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TicketChangingEnquiryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModule struct {
	FlightInfoList []*TicketChangingEnquiryResponseBodyModuleFlightInfoList `json:"flight_info_list,omitempty" xml:"flight_info_list,omitempty" type:"Repeated"`
}

func (s TicketChangingEnquiryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModule) GetFlightInfoList() []*TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	return s.FlightInfoList
}

func (s *TicketChangingEnquiryResponseBodyModule) SetFlightInfoList(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoList) *TicketChangingEnquiryResponseBodyModule {
	s.FlightInfoList = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoList struct {
	AirlineInfo    *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	CabinList      []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList    `json:"cabin_list,omitempty" xml:"cabin_list,omitempty" type:"Repeated"`
	// example:
	//
	// CA1704
	CarrierAirline *string `json:"carrier_airline,omitempty" xml:"carrier_airline,omitempty"`
	// example:
	//
	// CA1704
	CarrierNo      *string                                                              `json:"carrier_no,omitempty" xml:"carrier_no,omitempty"`
	DepAirportInfo *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// CA1351
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// false
	IsShare *bool `json:"is_share,omitempty" xml:"is_share,omitempty"`
	// example:
	//
	// Y
	LowestCabin *string `json:"lowest_cabin,omitempty" xml:"lowest_cabin,omitempty"`
	// example:
	//
	// Y
	LowestCabinClass *string `json:"lowest_cabin_class,omitempty" xml:"lowest_cabin_class,omitempty"`
	LowestCabinDesc  *string `json:"lowest_cabin_desc,omitempty" xml:"lowest_cabin_desc,omitempty"`
	// example:
	//
	// 0
	LowestCabinNum   *string                                                                  `json:"lowest_cabin_num,omitempty" xml:"lowest_cabin_num,omitempty"`
	LowestCabinPrice []*TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice `json:"lowest_cabin_price,omitempty" xml:"lowest_cabin_price,omitempty" type:"Repeated"`
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
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetAirlineInfo() *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo {
	return s.AirlineInfo
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetArrAirportInfo() *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetCabinList() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	return s.CabinList
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetCarrierAirline() *string {
	return s.CarrierAirline
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetCarrierNo() *string {
	return s.CarrierNo
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetDepAirportInfo() *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetIsShare() *bool {
	return s.IsShare
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetLowestCabin() *string {
	return s.LowestCabin
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetLowestCabinClass() *string {
	return s.LowestCabinClass
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetLowestCabinDesc() *string {
	return s.LowestCabinDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetLowestCabinNum() *string {
	return s.LowestCabinNum
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetLowestCabinPrice() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice {
	return s.LowestCabinPrice
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetModifyFlightArrTime() *string {
	return s.ModifyFlightArrTime
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetModifyFlightDepDate() *string {
	return s.ModifyFlightDepDate
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetModifyFlightDepTime() *string {
	return s.ModifyFlightDepTime
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) GetSessionId() *string {
	return s.SessionId
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetAirlineInfo(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.AirlineInfo = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetArrAirportInfo(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.ArrAirportInfo = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetCabinList(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.CabinList = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetCarrierAirline(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.CarrierAirline = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetCarrierNo(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.CarrierNo = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetDepAirportInfo(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.DepAirportInfo = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetDepCityCode(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.DepCityCode = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetFlightNo(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.FlightNo = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetIsShare(v bool) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.IsShare = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetLowestCabin(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.LowestCabin = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetLowestCabinClass(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.LowestCabinClass = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetLowestCabinDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.LowestCabinDesc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetLowestCabinNum(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.LowestCabinNum = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetLowestCabinPrice(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.LowestCabinPrice = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetModifyFlightArrTime(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.ModifyFlightArrTime = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetModifyFlightDepDate(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.ModifyFlightDepDate = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetModifyFlightDepTime(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.ModifyFlightDepTime = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) SetSessionId(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoList {
	s.SessionId = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) SetAirlineCode(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) SetAirlineName(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) SetAirlineSimpleName(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo {
	s.AirlineSimpleName = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo struct {
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

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) SetAirportCode(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) SetAirportName(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) SetCityCode(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo {
	s.CityCode = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) SetCityName(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo {
	s.CityName = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) SetTerminal(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList struct {
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
	CabinDiscount       *int32                                                                             `json:"cabin_discount,omitempty" xml:"cabin_discount,omitempty"`
	ChangeOtaItemRuleRq *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq `json:"change_ota_item_rule_rq,omitempty" xml:"change_ota_item_rule_rq,omitempty" type:"Struct"`
	// example:
	//
	// G
	ChildCabin *string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
	// example:
	//
	// 0
	LeftNum         *string                                                                          `json:"left_num,omitempty" xml:"left_num,omitempty"`
	ModifyPriceList []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList `json:"modify_price_list,omitempty" xml:"modify_price_list,omitempty" type:"Repeated"`
	// example:
	//
	// 360379a11ee84e9aa011baa41b758fe6
	OtaItemid *string `json:"ota_itemid,omitempty" xml:"ota_itemid,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetCabin() *string {
	return s.Cabin
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetCabinDesc() *string {
	return s.CabinDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetCabinDiscount() *int32 {
	return s.CabinDiscount
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetChangeOtaItemRuleRq() *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq {
	return s.ChangeOtaItemRuleRq
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetChildCabin() *string {
	return s.ChildCabin
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetLeftNum() *string {
	return s.LeftNum
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetModifyPriceList() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	return s.ModifyPriceList
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) GetOtaItemid() *string {
	return s.OtaItemid
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetCabin(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.Cabin = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetCabinClass(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.CabinClass = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetCabinDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.CabinDesc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetCabinDiscount(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.CabinDiscount = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetChangeOtaItemRuleRq(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.ChangeOtaItemRuleRq = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetChildCabin(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.ChildCabin = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetLeftNum(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.LeftNum = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetModifyPriceList(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.ModifyPriceList = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) SetOtaItemid(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList {
	s.OtaItemid = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq struct {
	BaggageDetails []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails `json:"baggage_details,omitempty" xml:"baggage_details,omitempty" type:"Repeated"`
	ChangeDetails  []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails  `json:"change_details,omitempty" xml:"change_details,omitempty" type:"Repeated"`
	RefundDetails  []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails  `json:"refund_details,omitempty" xml:"refund_details,omitempty" type:"Repeated"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) GetBaggageDetails() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	return s.BaggageDetails
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) GetChangeDetails() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	return s.ChangeDetails
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) GetRefundDetails() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	return s.RefundDetails
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) SetBaggageDetails(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq {
	s.BaggageDetails = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) SetChangeDetails(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq {
	s.ChangeDetails = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) SetRefundDetails(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq {
	s.RefundDetails = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRq) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails struct {
	BaggageSubItems []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index     *int32                                                                                               `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                                              `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                                              `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GetBaggageSubItems() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	return s.BaggageSubItems
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GetIndex() *int32 {
	return s.Index
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GetTableHead() *string {
	return s.TableHead
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GetTips() *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips {
	return s.Tips
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) GetType() *int32 {
	return s.Type
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) SetBaggageSubItems(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	s.BaggageSubItems = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) SetIndex(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	s.Index = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) SetTableHead(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	s.TableHead = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) SetTips(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	s.Tips = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) SetType(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails {
	s.Type = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetails) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems struct {
	// attributes
	Attributes                  map[string]interface{}                                                                                                                       `json:"attributes,omitempty" xml:"attributes,omitempty"`
	BaggageSubContentVisualizes []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	BaggageSubContents          []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents          `json:"baggage_sub_contents,omitempty" xml:"baggage_sub_contents,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// example:
	//
	// ADT
	Ptc   *string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetBaggageSubContentVisualizes() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetBaggageSubContents() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents {
	return s.BaggageSubContents
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetContent() *string {
	return s.Content
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetAttributes(v map[string]interface{}) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.Attributes = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetBaggageSubContentVisualizes(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetBaggageSubContents(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.BaggageSubContents = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetContent(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.Content = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetIsStruct(v bool) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetPtc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BaggageSubContentType *int32                                                                                                                                                `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"imageDO,omitempty" xml:"imageDO,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"subTitle,omitempty" xml:"subTitle,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// demo
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// demo
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01kLt3m923XsUs6WVif_!!6000000007266-2-tps-280-300.png
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// example:
	//
	// 50
	Largest *string `json:"largest,omitempty" xml:"largest,omitempty"`
	// example:
	//
	// 40
	Middle *string `json:"middle,omitempty" xml:"middle,omitempty"`
	// example:
	//
	// 20
	Smallest *string `json:"smallest,omitempty" xml:"smallest,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents struct {
	BaggageDesc *string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 0
	Style    *int32  `json:"style,omitempty" xml:"style,omitempty"`
	SubTitle *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) GetIcon() *string {
	return s.Icon
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) GetSubTitle() *string {
	return s.SubTitle
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) SetBaggageDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents {
	s.BaggageDesc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) SetIcon(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents {
	s.Icon = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) SetStyle(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents {
	s.Style = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) SetSubTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents {
	s.SubTitle = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsBaggageSubItemsBaggageSubContents) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips struct {
	// example:
	//
	// https://gw.alicdn.com/imgextra/i1/O1CN019zl3WZ22fNLxzx2cR_!!6000000007147-2-tps-125-45.png
	Logo     *string `json:"logo,omitempty" xml:"logo,omitempty"`
	TipsDesc *string `json:"tips_desc,omitempty" xml:"tips_desc,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01rJxjw61f3bXNHAmlk_!!6000000003951-2-tps-1050-675.png
	TipsImage *string `json:"tips_image,omitempty" xml:"tips_image,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) GetLogo() *string {
	return s.Logo
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) SetLogo(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips {
	s.Logo = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) SetTipsDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips {
	s.TipsDesc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) SetTipsImage(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips {
	s.TipsImage = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqBaggageDetailsTips) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails struct {
	ExtraContents []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                                          `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	TableHead      *string                                                                                                         `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                                         `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GetExtraContents() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents {
	return s.ExtraContents
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GetIndex() *int32 {
	return s.Index
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GetRefundSubItems() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems {
	return s.RefundSubItems
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GetTableHead() *string {
	return s.TableHead
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) GetType() *int32 {
	return s.Type
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) SetExtraContents(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	s.ExtraContents = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) SetIndex(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	s.Index = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) SetRefundSubItems(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	s.RefundSubItems = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) SetTableHead(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	s.TableHead = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) SetType(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails {
	s.Type = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetails) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Icon    *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) GetContent() *string {
	return s.Content
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) GetIcon() *string {
	return s.Icon
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) SetContent(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents {
	s.Content = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) SetIcon(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents {
	s.Icon = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsExtraContents) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                          `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                          `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) GetContent() *string {
	return s.Content
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) GetRefundSubContents() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) SetContent(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems {
	s.Content = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) SetIsStruct(v bool) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) SetPtc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) SetRefundSubContents(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 0
	Style *string `json:"style,omitempty" xml:"style,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) GetStyle() *string {
	return s.Style
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) SetFeeDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) SetFeeRange(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) SetStyle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqChangeDetailsRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails struct {
	ExtraContents []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                                          `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	TableHead      *string                                                                                                         `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                                         `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GetExtraContents() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents {
	return s.ExtraContents
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GetIndex() *int32 {
	return s.Index
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GetRefundSubItems() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems {
	return s.RefundSubItems
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GetTableHead() *string {
	return s.TableHead
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) GetType() *int32 {
	return s.Type
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) SetExtraContents(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	s.ExtraContents = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) SetIndex(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	s.Index = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) SetRefundSubItems(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	s.RefundSubItems = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) SetTableHead(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	s.TableHead = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) SetType(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails {
	s.Type = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetails) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Icon    *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) GetContent() *string {
	return s.Content
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) GetIcon() *string {
	return s.Icon
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) SetContent(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents {
	s.Content = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) SetIcon(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents {
	s.Icon = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsExtraContents) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                          `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                          `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) GetContent() *string {
	return s.Content
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) GetRefundSubContents() []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) SetContent(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems {
	s.Content = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) SetIsStruct(v bool) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) SetPtc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) SetRefundSubContents(v []*TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) SetTitle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems {
	s.Title = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 0
	Style *string `json:"style,omitempty" xml:"style,omitempty"`
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) GetStyle() *string {
	return s.Style
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) SetFeeDesc(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) SetFeeRange(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) SetStyle(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListChangeOtaItemRuleRqRefundDetailsRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList struct {
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

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetUpgradeFee() *int32 {
	return s.UpgradeFee
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) GetUpgradePrice() *int32 {
	return s.UpgradePrice
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetPassengerType(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.PassengerType = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetTicketPrice(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.TicketPrice = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetUpgradeFee(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.UpgradeFee = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) SetUpgradePrice(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList {
	s.UpgradePrice = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListCabinListModifyPriceList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo struct {
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

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) SetAirportCode(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) SetAirportName(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) SetCityCode(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo {
	s.CityCode = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) SetCityName(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo {
	s.CityName = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) SetTerminal(v string) *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice struct {
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

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) GetUpgradeFee() *int32 {
	return s.UpgradeFee
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) GetUpgradePrice() *int32 {
	return s.UpgradePrice
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) SetPassengerType(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.PassengerType = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) SetTicketPrice(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.TicketPrice = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) SetUpgradeFee(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.UpgradeFee = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) SetUpgradePrice(v int32) *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice {
	s.UpgradePrice = &v
	return s
}

func (s *TicketChangingEnquiryResponseBodyModuleFlightInfoListLowestCabinPrice) Validate() error {
	return dara.Validate(s)
}
