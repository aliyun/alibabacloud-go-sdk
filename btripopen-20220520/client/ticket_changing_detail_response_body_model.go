// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TicketChangingDetailResponseBody
	GetCode() *string
	SetMessage(v string) *TicketChangingDetailResponseBody
	GetMessage() *string
	SetModule(v *TicketChangingDetailResponseBodyModule) *TicketChangingDetailResponseBody
	GetModule() *TicketChangingDetailResponseBodyModule
	SetRequestId(v string) *TicketChangingDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketChangingDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TicketChangingDetailResponseBody
	GetTraceId() *string
}

type TicketChangingDetailResponseBody struct {
	// example:
	//
	// 0
	Code    *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TicketChangingDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TicketChangingDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailResponseBody) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *TicketChangingDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TicketChangingDetailResponseBody) GetModule() *TicketChangingDetailResponseBodyModule {
	return s.Module
}

func (s *TicketChangingDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketChangingDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketChangingDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TicketChangingDetailResponseBody) SetCode(v string) *TicketChangingDetailResponseBody {
	s.Code = &v
	return s
}

func (s *TicketChangingDetailResponseBody) SetMessage(v string) *TicketChangingDetailResponseBody {
	s.Message = &v
	return s
}

func (s *TicketChangingDetailResponseBody) SetModule(v *TicketChangingDetailResponseBodyModule) *TicketChangingDetailResponseBody {
	s.Module = v
	return s
}

func (s *TicketChangingDetailResponseBody) SetRequestId(v string) *TicketChangingDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketChangingDetailResponseBody) SetSuccess(v bool) *TicketChangingDetailResponseBody {
	s.Success = &v
	return s
}

func (s *TicketChangingDetailResponseBody) SetTraceId(v string) *TicketChangingDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *TicketChangingDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type TicketChangingDetailResponseBodyModule struct {
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// 1002
	BtripOrderId *int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// 37772
	BtripSubOrderId *int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// chang123
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// example:
	//
	// {}
	Extra          *string                                                 `json:"extra,omitempty" xml:"extra,omitempty"`
	FlightInfoList []*TicketChangingDetailResponseBodyModuleFlightInfoList `json:"flight_info_list,omitempty" xml:"flight_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-12-30 11:30:00
	LastPayTime *string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// example:
	//
	// 0
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 2022-12-30 11:30:00
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 100
	SettlePrice *int64 `json:"settle_price,omitempty" xml:"settle_price,omitempty"`
	SettleType  *int32 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	Status      *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 100
	TotalChangePrice *int64 `json:"total_change_price,omitempty" xml:"total_change_price,omitempty"`
	// example:
	//
	// 100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// example:
	//
	// 100
	TotalUpgradePrice *int64                                                    `json:"total_upgrade_price,omitempty" xml:"total_upgrade_price,omitempty"`
	TravelerInfoList  []*TicketChangingDetailResponseBodyModuleTravelerInfoList `json:"traveler_info_list,omitempty" xml:"traveler_info_list,omitempty" type:"Repeated"`
}

func (s TicketChangingDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailResponseBodyModule) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *TicketChangingDetailResponseBodyModule) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *TicketChangingDetailResponseBodyModule) GetBtripSubOrderId() *int64 {
	return s.BtripSubOrderId
}

func (s *TicketChangingDetailResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingDetailResponseBodyModule) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingDetailResponseBodyModule) GetExtra() *string {
	return s.Extra
}

func (s *TicketChangingDetailResponseBodyModule) GetFlightInfoList() []*TicketChangingDetailResponseBodyModuleFlightInfoList {
	return s.FlightInfoList
}

func (s *TicketChangingDetailResponseBodyModule) GetLastPayTime() *string {
	return s.LastPayTime
}

func (s *TicketChangingDetailResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *TicketChangingDetailResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *TicketChangingDetailResponseBodyModule) GetSettlePrice() *int64 {
	return s.SettlePrice
}

func (s *TicketChangingDetailResponseBodyModule) GetSettleType() *int32 {
	return s.SettleType
}

func (s *TicketChangingDetailResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *TicketChangingDetailResponseBodyModule) GetTotalChangePrice() *int64 {
	return s.TotalChangePrice
}

func (s *TicketChangingDetailResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *TicketChangingDetailResponseBodyModule) GetTotalUpgradePrice() *int64 {
	return s.TotalUpgradePrice
}

func (s *TicketChangingDetailResponseBodyModule) GetTravelerInfoList() []*TicketChangingDetailResponseBodyModuleTravelerInfoList {
	return s.TravelerInfoList
}

func (s *TicketChangingDetailResponseBodyModule) SetAlipayTradeNo(v string) *TicketChangingDetailResponseBodyModule {
	s.AlipayTradeNo = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetBtripOrderId(v int64) *TicketChangingDetailResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetBtripSubOrderId(v int64) *TicketChangingDetailResponseBodyModule {
	s.BtripSubOrderId = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetDisOrderId(v string) *TicketChangingDetailResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetDisSubOrderId(v string) *TicketChangingDetailResponseBodyModule {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetExtra(v string) *TicketChangingDetailResponseBodyModule {
	s.Extra = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetFlightInfoList(v []*TicketChangingDetailResponseBodyModuleFlightInfoList) *TicketChangingDetailResponseBodyModule {
	s.FlightInfoList = v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetLastPayTime(v string) *TicketChangingDetailResponseBodyModule {
	s.LastPayTime = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetPayStatus(v int32) *TicketChangingDetailResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetPayTime(v string) *TicketChangingDetailResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetSettlePrice(v int64) *TicketChangingDetailResponseBodyModule {
	s.SettlePrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetSettleType(v int32) *TicketChangingDetailResponseBodyModule {
	s.SettleType = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetStatus(v int32) *TicketChangingDetailResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetTotalChangePrice(v int64) *TicketChangingDetailResponseBodyModule {
	s.TotalChangePrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetTotalPrice(v int64) *TicketChangingDetailResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetTotalUpgradePrice(v int64) *TicketChangingDetailResponseBodyModule {
	s.TotalUpgradePrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) SetTravelerInfoList(v []*TicketChangingDetailResponseBodyModuleTravelerInfoList) *TicketChangingDetailResponseBodyModule {
	s.TravelerInfoList = v
	return s
}

func (s *TicketChangingDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TicketChangingDetailResponseBodyModuleFlightInfoList struct {
	// example:
	//
	// MU
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
	ArrAirport        *string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// example:
	//
	// HGH
	ArrAirportCode     *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrAirportCodeName *string `json:"arr_airport_code_name,omitempty" xml:"arr_airport_code_name,omitempty"`
	// example:
	//
	// HGH
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// LHW
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T3
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	Baggage *string `json:"baggage,omitempty" xml:"baggage,omitempty"`
	// example:
	//
	// 100
	BuildPrice *int64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// example:
	//
	// V
	Cabin      *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// MU3849
	Carrier    *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	DepAirport *string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// example:
	//
	// NGB
	DepAirportCode     *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepAirportCodeName *string `json:"dep_airport_code_name,omitempty" xml:"dep_airport_code_name,omitempty"`
	// example:
	//
	// BJS
	DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// NGB
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T4
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2000-01-01 00:00:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// MU3849
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// Y
	LastCabin *string `json:"last_cabin,omitempty" xml:"last_cabin,omitempty"`
	// example:
	//
	// CA1982
	LastFlightNo *string `json:"last_flight_no,omitempty" xml:"last_flight_no,omitempty"`
	// example:
	//
	// 3
	Meal *string `json:"meal,omitempty" xml:"meal,omitempty"`
	// example:
	//
	// 100
	OilPrice    *int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	SegmentType *int32 `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// NGB
	StopCity *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 100
	TicketPrice    *int64  `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	TuigaiqianInfo *string `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
}

func (s TicketChangingDetailResponseBodyModuleFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailResponseBodyModuleFlightInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrAirport() *string {
	return s.ArrAirport
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrAirportCodeName() *string {
	return s.ArrAirportCodeName
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetArrTime() *string {
	return s.ArrTime
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetBaggage() *string {
	return s.Baggage
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetBuildPrice() *int64 {
	return s.BuildPrice
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetCabin() *string {
	return s.Cabin
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetCarrier() *string {
	return s.Carrier
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepAirport() *string {
	return s.DepAirport
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepAirportCodeName() *string {
	return s.DepAirportCodeName
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetDepTime() *string {
	return s.DepTime
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetLastCabin() *string {
	return s.LastCabin
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetLastFlightNo() *string {
	return s.LastFlightNo
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetMeal() *string {
	return s.Meal
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetOilPrice() *int64 {
	return s.OilPrice
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetSegmentType() *int32 {
	return s.SegmentType
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetStopCity() *string {
	return s.StopCity
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetAirlineCode(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.AirlineCode = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetAirlineName(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.AirlineName = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetAirlineSimpleName(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.AirlineSimpleName = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrAirport(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrAirport = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrAirportCode(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrAirportCodeName(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrAirportCodeName = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrCity(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrCity = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrCityCode(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrTerminal(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrTerminal = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetArrTime(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.ArrTime = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetBaggage(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.Baggage = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetBuildPrice(v int64) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.BuildPrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetCabin(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.Cabin = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetCabinClass(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.CabinClass = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetCarrier(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.Carrier = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepAirport(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepAirport = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepAirportCode(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepAirportCodeName(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepAirportCodeName = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepCity(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepCity = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepCityCode(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepCityCode = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepTerminal(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepTerminal = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetDepTime(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.DepTime = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetFlightNo(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.FlightNo = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetLastCabin(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.LastCabin = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetLastFlightNo(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.LastFlightNo = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetMeal(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.Meal = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetOilPrice(v int64) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.OilPrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetSegmentType(v int32) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.SegmentType = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetStopArrTime(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.StopArrTime = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetStopCity(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.StopCity = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetStopDepTime(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.StopDepTime = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetTicketPrice(v int64) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.TicketPrice = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) SetTuigaiqianInfo(v string) *TicketChangingDetailResponseBodyModuleFlightInfoList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleFlightInfoList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingDetailResponseBodyModuleTravelerInfoList struct {
	// example:
	//
	// 2000-01-01
	BirthDate *string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// example:
	//
	// 12345
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 0
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// example:
	//
	// 0
	OpenTicketStatus *int32  `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
	PassengerName    *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 0
	PassengerType *string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// example:
	//
	// 12345678909
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 444-000000000
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 012992
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TicketChangingDetailResponseBodyModuleTravelerInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailResponseBodyModuleTravelerInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetBirthDate() *string {
	return s.BirthDate
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetCertNo() *string {
	return s.CertNo
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetCertType() *string {
	return s.CertType
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetOpenTicketStatus() *int32 {
	return s.OpenTicketStatus
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetPassengerType() *string {
	return s.PassengerType
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetPhone() *string {
	return s.Phone
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetBirthDate(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.BirthDate = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetCertNo(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.CertNo = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetCertType(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.CertType = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetOpenTicketStatus(v int32) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.OpenTicketStatus = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetPassengerName(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.PassengerName = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetPassengerType(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.PassengerType = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetPhone(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.Phone = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetTicketNo(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.TicketNo = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) SetUserId(v string) *TicketChangingDetailResponseBodyModuleTravelerInfoList {
	s.UserId = &v
	return s
}

func (s *TicketChangingDetailResponseBodyModuleTravelerInfoList) Validate() error {
	return dara.Validate(s)
}
