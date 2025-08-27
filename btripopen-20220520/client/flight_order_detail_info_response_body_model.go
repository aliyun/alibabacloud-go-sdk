// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOrderDetailInfoResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOrderDetailInfoResponseBody
	GetMessage() *string
	SetModule(v *FlightOrderDetailInfoResponseBodyModule) *FlightOrderDetailInfoResponseBody
	GetModule() *FlightOrderDetailInfoResponseBodyModule
	SetRequestId(v string) *FlightOrderDetailInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOrderDetailInfoResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOrderDetailInfoResponseBody
	GetTraceId() *string
}

type FlightOrderDetailInfoResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightOrderDetailInfoResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightOrderDetailInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOrderDetailInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOrderDetailInfoResponseBody) GetModule() *FlightOrderDetailInfoResponseBodyModule {
	return s.Module
}

func (s *FlightOrderDetailInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOrderDetailInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOrderDetailInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOrderDetailInfoResponseBody) SetCode(v string) *FlightOrderDetailInfoResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBody) SetMessage(v string) *FlightOrderDetailInfoResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBody) SetModule(v *FlightOrderDetailInfoResponseBodyModule) *FlightOrderDetailInfoResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderDetailInfoResponseBody) SetRequestId(v string) *FlightOrderDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBody) SetSuccess(v bool) *FlightOrderDetailInfoResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBody) SetTraceId(v string) *FlightOrderDetailInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailInfoResponseBodyModule struct {
	// example:
	//
	// 2389927372772
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// 1223
	BookUserId *string `json:"book_user_id,omitempty" xml:"book_user_id,omitempty"`
	// example:
	//
	// 123
	BtripOrderId *int64  `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	ContactName  *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// example:
	//
	// 12345678909
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// {}
	Extra          *string                                                  `json:"extra,omitempty" xml:"extra,omitempty"`
	FlightInfoList []*FlightOrderDetailInfoResponseBodyModuleFlightInfoList `json:"flight_info_list,omitempty" xml:"flight_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2000-00-00 00:00:00
	LastPayTime *string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// example:
	//
	// 0
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 1000
	PromotionPrice *int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// example:
	//
	// 100
	SettleAmount *int64 `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	// example:
	//
	// 0
	SettleType *int32 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// example:
	//
	// 1
	Status         *int32                                                   `json:"status,omitempty" xml:"status,omitempty"`
	TicketInfoList []*FlightOrderDetailInfoResponseBodyModuleTicketInfoList `json:"ticket_info_list,omitempty" xml:"ticket_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	TotalBuildPrice *int64 `json:"total_build_price,omitempty" xml:"total_build_price,omitempty"`
	// example:
	//
	// 1000
	TotalOilPrice *int64 `json:"total_oil_price,omitempty" xml:"total_oil_price,omitempty"`
	// example:
	//
	// 1000
	TotalOrderPrice  *int64                                                     `json:"total_order_price,omitempty" xml:"total_order_price,omitempty"`
	TravelerInfoList []*FlightOrderDetailInfoResponseBodyModuleTravelerInfoList `json:"traveler_info_list,omitempty" xml:"traveler_info_list,omitempty" type:"Repeated"`
}

func (s FlightOrderDetailInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetBookUserId() *string {
	return s.BookUserId
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetContactName() *string {
	return s.ContactName
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetExtra() *string {
	return s.Extra
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetFlightInfoList() []*FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	return s.FlightInfoList
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetLastPayTime() *string {
	return s.LastPayTime
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetPromotionPrice() *int64 {
	return s.PromotionPrice
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetSettleAmount() *int64 {
	return s.SettleAmount
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetSettleType() *int32 {
	return s.SettleType
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetTicketInfoList() []*FlightOrderDetailInfoResponseBodyModuleTicketInfoList {
	return s.TicketInfoList
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetTotalBuildPrice() *int64 {
	return s.TotalBuildPrice
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetTotalOilPrice() *int64 {
	return s.TotalOilPrice
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetTotalOrderPrice() *int64 {
	return s.TotalOrderPrice
}

func (s *FlightOrderDetailInfoResponseBodyModule) GetTravelerInfoList() []*FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	return s.TravelerInfoList
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetAlipayTradeNo(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.AlipayTradeNo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetBookUserId(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.BookUserId = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetBtripOrderId(v int64) *FlightOrderDetailInfoResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetContactName(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetContactPhone(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.ContactPhone = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetDisOrderId(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetExtra(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.Extra = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetFlightInfoList(v []*FlightOrderDetailInfoResponseBodyModuleFlightInfoList) *FlightOrderDetailInfoResponseBodyModule {
	s.FlightInfoList = v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetLastPayTime(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.LastPayTime = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetPayStatus(v int32) *FlightOrderDetailInfoResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetPayTime(v string) *FlightOrderDetailInfoResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetPromotionPrice(v int64) *FlightOrderDetailInfoResponseBodyModule {
	s.PromotionPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetSettleAmount(v int64) *FlightOrderDetailInfoResponseBodyModule {
	s.SettleAmount = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetSettleType(v int32) *FlightOrderDetailInfoResponseBodyModule {
	s.SettleType = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetStatus(v int32) *FlightOrderDetailInfoResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetTicketInfoList(v []*FlightOrderDetailInfoResponseBodyModuleTicketInfoList) *FlightOrderDetailInfoResponseBodyModule {
	s.TicketInfoList = v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetTotalBuildPrice(v int64) *FlightOrderDetailInfoResponseBodyModule {
	s.TotalBuildPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetTotalOilPrice(v int64) *FlightOrderDetailInfoResponseBodyModule {
	s.TotalOilPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetTotalOrderPrice(v int64) *FlightOrderDetailInfoResponseBodyModule {
	s.TotalOrderPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) SetTravelerInfoList(v []*FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) *FlightOrderDetailInfoResponseBodyModule {
	s.TravelerInfoList = v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailInfoResponseBodyModuleFlightInfoList struct {
	// example:
	//
	// CA
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
	ArrAirport        *string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// example:
	//
	// LHW
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
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// “”
	Baggage *string `json:"baggage,omitempty" xml:"baggage,omitempty"`
	// example:
	//
	// 100
	BuildPrice *int64  `json:"build_price,omitempty" xml:"build_price,omitempty"`
	Cabin      *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// "CA1351_PEK_CAN_2000-00-00 08:00
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
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2000-01-01 00:00:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// CA1351
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// Y
	LastCabin *string `json:"last_cabin,omitempty" xml:"last_cabin,omitempty"`
	// example:
	//
	// CA1351
	LastFlightNo *string `json:"last_flight_no,omitempty" xml:"last_flight_no,omitempty"`
	Meal         *string `json:"meal,omitempty" xml:"meal,omitempty"`
	// example:
	//
	// 100
	OilPrice *int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// example:
	//
	// 0
	SegmentType *int32 `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// CAN
	StopCity *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// “”
	TuigaiqianInfo *string `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
}

func (s FlightOrderDetailInfoResponseBodyModuleFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrAirport() *string {
	return s.ArrAirport
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrAirportCodeName() *string {
	return s.ArrAirportCodeName
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetBaggage() *string {
	return s.Baggage
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetBuildPrice() *int64 {
	return s.BuildPrice
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetCarrier() *string {
	return s.Carrier
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepAirport() *string {
	return s.DepAirport
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepAirportCodeName() *string {
	return s.DepAirportCodeName
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetLastCabin() *string {
	return s.LastCabin
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetLastFlightNo() *string {
	return s.LastFlightNo
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetMeal() *string {
	return s.Meal
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetOilPrice() *int64 {
	return s.OilPrice
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetSegmentType() *int32 {
	return s.SegmentType
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetAirlineCode(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.AirlineCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetAirlineName(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.AirlineName = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetAirlineSimpleName(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.AirlineSimpleName = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrAirport(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrAirport = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrAirportCode(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrAirportCodeName(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrAirportCodeName = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrCity(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrCityCode(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrTerminal(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrTerminal = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetArrTime(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetBaggage(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.Baggage = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetBuildPrice(v int64) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.BuildPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetCabin(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetCabinClass(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetCarrier(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.Carrier = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepAirport(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepAirport = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepAirportCode(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepAirportCodeName(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepAirportCodeName = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepCity(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepCityCode(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepTerminal(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepTerminal = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetDepTime(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetFlightNo(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetLastCabin(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.LastCabin = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetLastFlightNo(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.LastFlightNo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetMeal(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.Meal = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetOilPrice(v int64) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.OilPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetSegmentType(v int32) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.SegmentType = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetStopArrTime(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.StopArrTime = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetStopCity(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.StopCity = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetStopDepTime(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.StopDepTime = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetTicketPrice(v int64) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.TicketPrice = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) SetTuigaiqianInfo(v string) *FlightOrderDetailInfoResponseBodyModuleFlightInfoList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleFlightInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailInfoResponseBodyModuleTicketInfoList struct {
	// example:
	//
	// 0
	OpenTicketStatus *string `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
	// example:
	//
	// KF0528
	PnrCode *string `json:"pnr_code,omitempty" xml:"pnr_code,omitempty"`
	// example:
	//
	// 444-000000000
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 1
	TicketStatus *string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
}

func (s FlightOrderDetailInfoResponseBodyModuleTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoResponseBodyModuleTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) GetOpenTicketStatus() *string {
	return s.OpenTicketStatus
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) GetPnrCode() *string {
	return s.PnrCode
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) SetOpenTicketStatus(v string) *FlightOrderDetailInfoResponseBodyModuleTicketInfoList {
	s.OpenTicketStatus = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) SetPnrCode(v string) *FlightOrderDetailInfoResponseBodyModuleTicketInfoList {
	s.PnrCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) SetTicketNo(v string) *FlightOrderDetailInfoResponseBodyModuleTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) SetTicketStatus(v string) *FlightOrderDetailInfoResponseBodyModuleTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailInfoResponseBodyModuleTravelerInfoList struct {
	// example:
	//
	// 2000-01-01
	BirthDate *string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// example:
	//
	// 1234
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
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetBirthDate() *string {
	return s.BirthDate
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetCertNo() *string {
	return s.CertNo
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetCertType() *string {
	return s.CertType
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetOpenTicketStatus() *int32 {
	return s.OpenTicketStatus
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetPassengerType() *string {
	return s.PassengerType
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetPhone() *string {
	return s.Phone
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetBirthDate(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.BirthDate = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetCertNo(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.CertNo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetCertType(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.CertType = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetOpenTicketStatus(v int32) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.OpenTicketStatus = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetPassengerName(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetPassengerType(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.PassengerType = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetPhone(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.Phone = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetTicketNo(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) SetUserId(v string) *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList {
	s.UserId = &v
	return s
}

func (s *FlightOrderDetailInfoResponseBodyModuleTravelerInfoList) Validate() error {
	return dara.Validate(s)
}
