// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOtaSearchResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOtaSearchResponseBody
	GetMessage() *string
	SetModule(v *FlightOtaSearchResponseBodyModule) *FlightOtaSearchResponseBody
	GetModule() *FlightOtaSearchResponseBodyModule
	SetRequestId(v string) *FlightOtaSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOtaSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOtaSearchResponseBody
	GetTraceId() *string
}

type FlightOtaSearchResponseBody struct {
	// example:
	//
	// 0
	Code    *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightOtaSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightOtaSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOtaSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOtaSearchResponseBody) GetModule() *FlightOtaSearchResponseBodyModule {
	return s.Module
}

func (s *FlightOtaSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOtaSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOtaSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOtaSearchResponseBody) SetCode(v string) *FlightOtaSearchResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOtaSearchResponseBody) SetMessage(v string) *FlightOtaSearchResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOtaSearchResponseBody) SetModule(v *FlightOtaSearchResponseBodyModule) *FlightOtaSearchResponseBody {
	s.Module = v
	return s
}

func (s *FlightOtaSearchResponseBody) SetRequestId(v string) *FlightOtaSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOtaSearchResponseBody) SetSuccess(v bool) *FlightOtaSearchResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOtaSearchResponseBody) SetTraceId(v string) *FlightOtaSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOtaSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModule struct {
	FlightList []*FlightOtaSearchResponseBodyModuleFlightList `json:"flight_list,omitempty" xml:"flight_list,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModule) GetFlightList() []*FlightOtaSearchResponseBodyModuleFlightList {
	return s.FlightList
}

func (s *FlightOtaSearchResponseBodyModule) SetFlightList(v []*FlightOtaSearchResponseBodyModuleFlightList) *FlightOtaSearchResponseBodyModule {
	s.FlightList = v
	return s
}

func (s *FlightOtaSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightList struct {
	AirlineInfo    *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// 100
	BuildPrice *int32 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// 2
	CabinClass    *string                                                     `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinInfoList []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoList `json:"cabin_info_list,omitempty" xml:"cabin_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// FM
	CarrierAirline *string `json:"carrier_airline,omitempty" xml:"carrier_airline,omitempty"`
	// example:
	//
	// FM9152
	CarrierNo      *string                                                    `json:"carrier_no,omitempty" xml:"carrier_no,omitempty"`
	DepAirportInfo *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// 1
	Discount *int32 `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// CA1234
	FlightNo       *string                                                      `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightRuleList []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleList `json:"flight_rule_list,omitempty" xml:"flight_rule_list,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	FlightRuleListStr *string `json:"flight_rule_list_str,omitempty" xml:"flight_rule_list_str,omitempty"`
	FlightSize        *string `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	// example:
	//
	// demo
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 1
	InvoiceType *int32 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
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
	// false
	IsTransfer *bool   `json:"is_transfer,omitempty" xml:"is_transfer,omitempty"`
	MealDesc   *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 100
	OilPrice *int32 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// example:
	//
	// 6669c8e53b684105b8687bad0331988a_41
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// example:
	//
	// 100
	Price *int32 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 1
	ProductType     *int64  `json:"product_type,omitempty" xml:"product_type,omitempty"`
	ProductTypeDesc *string `json:"product_type_desc,omitempty" xml:"product_type_desc,omitempty"`
	// example:
	//
	// 100
	PromotionPrice *string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// example:
	//
	// 1
	RemainedSeatCount *string `json:"remained_seat_count,omitempty" xml:"remained_seat_count,omitempty"`
	// example:
	//
	// 1000_1_0
	SecretParams *string `json:"secret_params,omitempty" xml:"secret_params,omitempty"`
	// example:
	//
	// 1
	SegmentNumber *string `json:"segment_number,omitempty" xml:"segment_number,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// HGH
	StopCity *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 100
	TotalPrice *string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightList) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetAirlineInfo() *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetArrAirportInfo() *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetArrDate() *string {
	return s.ArrDate
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetCabinInfoList() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	return s.CabinInfoList
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetCarrierAirline() *string {
	return s.CarrierAirline
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetCarrierNo() *string {
	return s.CarrierNo
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetDepAirportInfo() *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetDiscount() *int32 {
	return s.Discount
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetFlightRuleList() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	return s.FlightRuleList
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetFlightRuleListStr() *string {
	return s.FlightRuleListStr
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetIsShare() *bool {
	return s.IsShare
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetIsStop() *bool {
	return s.IsStop
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetIsTransfer() *bool {
	return s.IsTransfer
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetPrice() *int32 {
	return s.Price
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetProductType() *int64 {
	return s.ProductType
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetProductTypeDesc() *string {
	return s.ProductTypeDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetPromotionPrice() *string {
	return s.PromotionPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetRemainedSeatCount() *string {
	return s.RemainedSeatCount
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetSecretParams() *string {
	return s.SecretParams
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetSegmentNumber() *string {
	return s.SegmentNumber
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetTotalPrice() *string {
	return s.TotalPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetAirlineInfo(v *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) *FlightOtaSearchResponseBodyModuleFlightList {
	s.AirlineInfo = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetArrAirportInfo(v *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) *FlightOtaSearchResponseBodyModuleFlightList {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetArrDate(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.ArrDate = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetBuildPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.BuildPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetCabin(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.Cabin = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetCabinClass(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetCabinInfoList(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) *FlightOtaSearchResponseBodyModuleFlightList {
	s.CabinInfoList = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetCarrierAirline(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.CarrierAirline = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetCarrierNo(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.CarrierNo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetDepAirportInfo(v *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) *FlightOtaSearchResponseBodyModuleFlightList {
	s.DepAirportInfo = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetDepCityCode(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetDepDate(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.DepDate = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetDiscount(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.Discount = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetFlightNo(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetFlightRuleList(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) *FlightOtaSearchResponseBodyModuleFlightList {
	s.FlightRuleList = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetFlightRuleListStr(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.FlightRuleListStr = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetFlightSize(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.FlightSize = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetFlightType(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.FlightType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetInvoiceType(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.InvoiceType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetIsProtocol(v bool) *FlightOtaSearchResponseBodyModuleFlightList {
	s.IsProtocol = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetIsShare(v bool) *FlightOtaSearchResponseBodyModuleFlightList {
	s.IsShare = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetIsStop(v bool) *FlightOtaSearchResponseBodyModuleFlightList {
	s.IsStop = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetIsTransfer(v bool) *FlightOtaSearchResponseBodyModuleFlightList {
	s.IsTransfer = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetMealDesc(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.MealDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetOilPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.OilPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetOtaItemId(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.OtaItemId = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.Price = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetProductType(v int64) *FlightOtaSearchResponseBodyModuleFlightList {
	s.ProductType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetProductTypeDesc(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.ProductTypeDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetPromotionPrice(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.PromotionPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetRemainedSeatCount(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.RemainedSeatCount = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetSecretParams(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.SecretParams = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetSegmentNumber(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.SegmentNumber = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetStopArrTime(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.StopArrTime = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetStopCity(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.StopCity = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetStopDepTime(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.StopDepTime = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetTicketPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.TicketPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetTotalPrice(v string) *FlightOtaSearchResponseBodyModuleFlightList {
	s.TotalPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) SetTripType(v int32) *FlightOtaSearchResponseBodyModuleFlightList {
	s.TripType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) SetAirlineCode(v string) *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) SetAirlineName(v string) *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) SetAirlineSimpleName(v string) *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo {
	s.AirlineSimpleName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo struct {
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
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) SetAirportCode(v string) *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) SetAirportName(v string) *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) SetCityCode(v string) *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) SetCityName(v string) *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) SetTerminal(v string) *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoList struct {
	AgentId *int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// example:
	//
	// 12456
	BasicCabinPrice *int32 `json:"basic_cabin_price,omitempty" xml:"basic_cabin_price,omitempty"`
	// example:
	//
	// 100
	BuildPrice *int32 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// example:
	//
	// G
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// G
	ChildCabin *string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
	ClassName  *string `json:"class_name,omitempty" xml:"class_name,omitempty"`
	// example:
	//
	// demo
	ClassRule *string `json:"class_rule,omitempty" xml:"class_rule,omitempty"`
	// example:
	//
	// 10
	Discount       *string                                                                   `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightRuleList []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList `json:"flight_rule_list,omitempty" xml:"flight_rule_list,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	FlightRuleListStr *string `json:"flight_rule_list_str,omitempty" xml:"flight_rule_list_str,omitempty"`
	// example:
	//
	// 1
	InvoiceType *int32 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// example:
	//
	// true
	IsProtocol *bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	// example:
	//
	// 100
	OilPrice *int32 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// example:
	//
	// 1000_1_0
	OrderParams *string `json:"order_params,omitempty" xml:"order_params,omitempty"`
	// example:
	//
	// 97f64e2d6f61408a827dd523817fefd6_0
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// example:
	//
	// 100
	Price *int32 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 1
	ProductType *int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// example:
	//
	// demo
	ProductTypeDesc *string `json:"product_type_desc,omitempty" xml:"product_type_desc,omitempty"`
	// example:
	//
	// 100
	PromotionPrice *string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// example:
	//
	// 1
	RemainedSeatCount *string `json:"remained_seat_count,omitempty" xml:"remained_seat_count,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 100
	TotalPrice *int32 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetChildCabin() *string {
	return s.ChildCabin
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetClassName() *string {
	return s.ClassName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetClassRule() *string {
	return s.ClassRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetDiscount() *string {
	return s.Discount
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetFlightRuleList() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	return s.FlightRuleList
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetFlightRuleListStr() *string {
	return s.FlightRuleListStr
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetOrderParams() *string {
	return s.OrderParams
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetPrice() *int32 {
	return s.Price
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetProductType() *int64 {
	return s.ProductType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetProductTypeDesc() *string {
	return s.ProductTypeDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetPromotionPrice() *string {
	return s.PromotionPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetRemainedSeatCount() *string {
	return s.RemainedSeatCount
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) GetTotalPrice() *int32 {
	return s.TotalPrice
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetAgentId(v int64) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.AgentId = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetBasicCabinPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.BasicCabinPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetBuildPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.BuildPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetCabin(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.Cabin = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetCabinClass(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.CabinClass = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetCabinClassName(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.CabinClassName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetChildCabin(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.ChildCabin = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetClassName(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.ClassName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetClassRule(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.ClassRule = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetDiscount(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.Discount = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetFlightRuleList(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.FlightRuleList = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetFlightRuleListStr(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.FlightRuleListStr = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetInvoiceType(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.InvoiceType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetIsProtocol(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.IsProtocol = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetOilPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.OilPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetOrderParams(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.OrderParams = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetOtaItemId(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.OtaItemId = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.Price = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetProductType(v int64) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.ProductType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetProductTypeDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.ProductTypeDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetPromotionPrice(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.PromotionPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetRemainedSeatCount(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.RemainedSeatCount = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetTicketPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.TicketPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) SetTotalPrice(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList {
	s.TotalPrice = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList struct {
	// example:
	//
	// demo
	BaggageInfo    *string                                                                               `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	BaggageItem    *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem    `json:"baggage_item,omitempty" xml:"baggage_item,omitempty" type:"Struct"`
	ChangeRule     *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule     `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Struct"`
	ChangeRuleItem *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Extra          *string                                                                               `json:"extra,omitempty" xml:"extra,omitempty"`
	RefundRule     *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Struct"`
	RefundRuleItem *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty" type:"Struct"`
	SignRule       *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule       `json:"sign_rule,omitempty" xml:"sign_rule,omitempty" type:"Struct"`
	TuigaiqianInfo *string                                                                               `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	UpgradeRule    *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule    `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty" type:"Struct"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetBaggageInfo() *string {
	return s.BaggageInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetBaggageItem() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	return s.BaggageItem
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetChangeRule() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule {
	return s.ChangeRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetChangeRuleItem() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	return s.ChangeRuleItem
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetExtra() *string {
	return s.Extra
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetRefundRule() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule {
	return s.RefundRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetRefundRuleItem() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	return s.RefundRuleItem
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetSignRule() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule {
	return s.SignRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetUpgradeRule() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule {
	return s.UpgradeRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetBaggageInfo(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.BaggageInfo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetBaggageItem(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.BaggageItem = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetChangeRule(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.ChangeRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetChangeRuleItem(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.ChangeRuleItem = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetExtra(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.Extra = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetRefundRule(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.RefundRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetRefundRuleItem(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.RefundRuleItem = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetSignRule(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.SignRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetTuigaiqianInfo(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetUpgradeRule(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.UpgradeRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleList) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem struct {
	BaggageSubItems []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index     *int32                                                                                 `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                                `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                                `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetBaggageSubItems() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetTips() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	return s.Tips
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetBaggageSubItems(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.BaggageSubItems = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetIndex(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Index = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetTableHead(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.TableHead = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetTips(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Tips = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetType(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Type = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                                                  `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc   *string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetIsStruct(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetPtc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	BaggageSubContentType *int32                                                                                                                                  `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i4/O1CN01UynXG31pjsEtA3tMF_!!6000000005397-2-tps-36-36.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i1/O1CN01qe7wL21gJ0SmEXXL7_!!6000000004120-2-tps-1206-768.png
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
	// example:
	//
	// demo
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

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips struct {
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

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) SetLogo(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	s.Logo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) SetTipsDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) SetTipsImage(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	s.TipsImage = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule struct {
	// example:
	//
	// true
	Able *bool                                                                                   `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem struct {
	ExtraContents []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                                `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                             `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                               `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                               `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetExtraContents() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetRefundSubItems() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetExtraContents(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetIndex(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.Index = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetRefundSubItems(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetSubTableHead(v []*string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetTableHead(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetType(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.Type = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetRefundSubContents() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetIsStruct(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetPtc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetRefundSubContents(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule struct {
	// example:
	//
	// true
	Able *bool                                                                                   `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem struct {
	ExtraContents []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index          *int32                                                                                                `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                             `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                               `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetExtraContents() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetRefundSubItems() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetExtraContents(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetIndex(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.Index = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetRefundSubItems(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetSubTableHead(v []*string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetTableHead(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetType(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.Type = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetRefundSubContents() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetIsStruct(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetPtc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetRefundSubContents(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule struct {
	// example:
	//
	// true
	Able *bool                                                                                 `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule struct {
	// example:
	//
	// true
	Able *bool                                                                                    `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo struct {
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

func (s FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) SetAirportCode(v string) *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) SetAirportName(v string) *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) SetCityCode(v string) *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) SetCityName(v string) *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) SetTerminal(v string) *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleList struct {
	// example:
	//
	// demo
	BaggageInfo    *string                                                                  `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	BaggageItem    *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem    `json:"baggage_item,omitempty" xml:"baggage_item,omitempty" type:"Struct"`
	ChangeRule     *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule     `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Struct"`
	ChangeRuleItem *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Extra          *string                                                                  `json:"extra,omitempty" xml:"extra,omitempty"`
	RefundRule     *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Struct"`
	RefundRuleItem *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty" type:"Struct"`
	SignRule       *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule       `json:"sign_rule,omitempty" xml:"sign_rule,omitempty" type:"Struct"`
	TuigaiqianInfo *string                                                                  `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	UpgradeRule    *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule    `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty" type:"Struct"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetBaggageInfo() *string {
	return s.BaggageInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetBaggageItem() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	return s.BaggageItem
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetChangeRule() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule {
	return s.ChangeRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetChangeRuleItem() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	return s.ChangeRuleItem
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetExtra() *string {
	return s.Extra
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetRefundRule() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule {
	return s.RefundRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetRefundRuleItem() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	return s.RefundRuleItem
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetSignRule() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule {
	return s.SignRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) GetUpgradeRule() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	return s.UpgradeRule
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetBaggageInfo(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.BaggageInfo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetBaggageItem(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.BaggageItem = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetChangeRule(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.ChangeRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetChangeRuleItem(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.ChangeRuleItem = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetExtra(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.Extra = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetRefundRule(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.RefundRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetRefundRuleItem(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.RefundRuleItem = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetSignRule(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.SignRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetTuigaiqianInfo(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) SetUpgradeRule(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList {
	s.UpgradeRule = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleList) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem struct {
	BaggageSubItems []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index     *int32                                                                    `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                   `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                   `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetBaggageSubItems() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTips() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	return s.Tips
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetBaggageSubItems(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.BaggageSubItems = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetIndex(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Index = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTableHead(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.TableHead = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTips(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Tips = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetType(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Type = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                                     `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc   *string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetIsStruct(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetPtc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BaggageSubContentType *int32                                                                                                                     `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i4/O1CN01UynXG31pjsEtA3tMF_!!6000000005397-2-tps-36-36.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i1/O1CN01qe7wL21gJ0SmEXXL7_!!6000000004120-2-tps-1206-768.png
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
	// example:
	//
	// demo
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

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips struct {
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

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetLogo(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.Logo = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetTipsDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetTipsImage(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.TipsImage = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule struct {
	// example:
	//
	// false
	Able *bool                                                                      `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem struct {
	ExtraContents []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                   `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                  `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                  `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetExtraContents() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetRefundSubItems() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetExtraContents(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetIndex(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Index = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetRefundSubItems(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetSubTableHead(v []*string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetTableHead(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetType(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Type = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems struct {
	// isStruct : true
	//
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                   `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetRefundSubContents() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetIsStruct(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetPtc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetRefundSubContents(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule struct {
	// example:
	//
	// false
	Able *bool                                                                      `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem struct {
	ExtraContents []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index          *int32                                                                                   `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                  `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                  `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetExtraContents() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetRefundSubItems() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetExtraContents(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetIndex(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Index = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetRefundSubItems(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetSubTableHead(v []*string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetTableHead(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetType(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Type = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                   `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetRefundSubContents() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetIsStruct(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetPtc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetRefundSubContents(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule struct {
	// example:
	//
	// false
	Able *bool                                                                    `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule struct {
	// example:
	//
	// false
	Able *bool                                                                       `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) GetInfo() []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	return s.Info
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) SetAble(v bool) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	s.Able = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) SetInfo(v []*FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	s.Info = v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 10
	CostPercent *int32 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
	// example:
	//
	// 1675036500000
	TimeStamp *int32 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// example:
	//
	// demo
	TimeType *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetContent(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetCost(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetCostPercent(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTimeStamp(v int32) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTimeType(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTitle(v string) *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightOtaSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) Validate() error {
	return dara.Validate(s)
}
