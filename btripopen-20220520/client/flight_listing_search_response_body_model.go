// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightListingSearchResponseBody
	GetCode() *string
	SetMessage(v string) *FlightListingSearchResponseBody
	GetMessage() *string
	SetModule(v *FlightListingSearchResponseBodyModule) *FlightListingSearchResponseBody
	GetModule() *FlightListingSearchResponseBodyModule
	SetRequestId(v string) *FlightListingSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightListingSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightListingSearchResponseBody
	GetTraceId() *string
}

type FlightListingSearchResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightListingSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightListingSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBody) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightListingSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightListingSearchResponseBody) GetModule() *FlightListingSearchResponseBodyModule {
	return s.Module
}

func (s *FlightListingSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightListingSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightListingSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightListingSearchResponseBody) SetCode(v string) *FlightListingSearchResponseBody {
	s.Code = &v
	return s
}

func (s *FlightListingSearchResponseBody) SetMessage(v string) *FlightListingSearchResponseBody {
	s.Message = &v
	return s
}

func (s *FlightListingSearchResponseBody) SetModule(v *FlightListingSearchResponseBodyModule) *FlightListingSearchResponseBody {
	s.Module = v
	return s
}

func (s *FlightListingSearchResponseBody) SetRequestId(v string) *FlightListingSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightListingSearchResponseBody) SetSuccess(v bool) *FlightListingSearchResponseBody {
	s.Success = &v
	return s
}

func (s *FlightListingSearchResponseBody) SetTraceId(v string) *FlightListingSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightListingSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModule struct {
	FlightList []*FlightListingSearchResponseBodyModuleFlightList `json:"flight_list,omitempty" xml:"flight_list,omitempty" type:"Repeated"`
}

func (s FlightListingSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModule) GetFlightList() []*FlightListingSearchResponseBodyModuleFlightList {
	return s.FlightList
}

func (s *FlightListingSearchResponseBodyModule) SetFlightList(v []*FlightListingSearchResponseBodyModuleFlightList) *FlightListingSearchResponseBodyModule {
	s.FlightList = v
	return s
}

func (s *FlightListingSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightList struct {
	AirlineInfo    *FlightListingSearchResponseBodyModuleFlightListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
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
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// 2
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// FM
	CarrierAirline *string `json:"carrier_airline,omitempty" xml:"carrier_airline,omitempty"`
	// example:
	//
	// FM9152
	CarrierNo      *string                                                        `json:"carrier_no,omitempty" xml:"carrier_no,omitempty"`
	DepAirportInfo *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
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
	FlightNo       *string                                                          `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightRuleList []*FlightListingSearchResponseBodyModuleFlightListFlightRuleList `json:"flight_rule_list,omitempty" xml:"flight_rule_list,omitempty" type:"Repeated"`
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
	// 4ec61b13fc9746f99c072a16bfc265af_0
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

func (s FlightListingSearchResponseBodyModuleFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightList) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetAirlineInfo() *FlightListingSearchResponseBodyModuleFlightListAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetArrAirportInfo() *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetArrDate() *string {
	return s.ArrDate
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetCarrierAirline() *string {
	return s.CarrierAirline
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetCarrierNo() *string {
	return s.CarrierNo
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetDepAirportInfo() *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetDiscount() *int32 {
	return s.Discount
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetFlightRuleList() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	return s.FlightRuleList
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetFlightRuleListStr() *string {
	return s.FlightRuleListStr
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetIsShare() *bool {
	return s.IsShare
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetIsStop() *bool {
	return s.IsStop
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetIsTransfer() *bool {
	return s.IsTransfer
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetPrice() *int32 {
	return s.Price
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetProductType() *int64 {
	return s.ProductType
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetProductTypeDesc() *string {
	return s.ProductTypeDesc
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetPromotionPrice() *string {
	return s.PromotionPrice
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetRemainedSeatCount() *string {
	return s.RemainedSeatCount
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetSecretParams() *string {
	return s.SecretParams
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetSegmentNumber() *string {
	return s.SegmentNumber
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetTotalPrice() *string {
	return s.TotalPrice
}

func (s *FlightListingSearchResponseBodyModuleFlightList) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetAirlineInfo(v *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) *FlightListingSearchResponseBodyModuleFlightList {
	s.AirlineInfo = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetArrAirportInfo(v *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) *FlightListingSearchResponseBodyModuleFlightList {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetArrDate(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.ArrDate = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetBasicCabinPrice(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.BasicCabinPrice = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetBuildPrice(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.BuildPrice = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetCabin(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.Cabin = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetCabinClass(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetCarrierAirline(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.CarrierAirline = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetCarrierNo(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.CarrierNo = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetDepAirportInfo(v *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) *FlightListingSearchResponseBodyModuleFlightList {
	s.DepAirportInfo = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetDepCityCode(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.DepCityCode = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetDepDate(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.DepDate = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetDiscount(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.Discount = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetFlightNo(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetFlightRuleList(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleList) *FlightListingSearchResponseBodyModuleFlightList {
	s.FlightRuleList = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetFlightRuleListStr(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.FlightRuleListStr = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetFlightSize(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.FlightSize = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetFlightType(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.FlightType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetInvoiceType(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.InvoiceType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetIsProtocol(v bool) *FlightListingSearchResponseBodyModuleFlightList {
	s.IsProtocol = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetIsShare(v bool) *FlightListingSearchResponseBodyModuleFlightList {
	s.IsShare = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetIsStop(v bool) *FlightListingSearchResponseBodyModuleFlightList {
	s.IsStop = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetIsTransfer(v bool) *FlightListingSearchResponseBodyModuleFlightList {
	s.IsTransfer = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetMealDesc(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.MealDesc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetOilPrice(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.OilPrice = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetOtaItemId(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.OtaItemId = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetPrice(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.Price = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetProductType(v int64) *FlightListingSearchResponseBodyModuleFlightList {
	s.ProductType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetProductTypeDesc(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.ProductTypeDesc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetPromotionPrice(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.PromotionPrice = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetRemainedSeatCount(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.RemainedSeatCount = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetSecretParams(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.SecretParams = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetSegmentNumber(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.SegmentNumber = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetStopArrTime(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.StopArrTime = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetStopCity(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.StopCity = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetStopDepTime(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.StopDepTime = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetTicketPrice(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.TicketPrice = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetTotalPrice(v string) *FlightListingSearchResponseBodyModuleFlightList {
	s.TotalPrice = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) SetTripType(v int32) *FlightListingSearchResponseBodyModuleFlightList {
	s.TripType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) SetAirlineCode(v string) *FlightListingSearchResponseBodyModuleFlightListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) SetAirlineName(v string) *FlightListingSearchResponseBodyModuleFlightListAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) SetAirlineSimpleName(v string) *FlightListingSearchResponseBodyModuleFlightListAirlineInfo {
	s.AirlineSimpleName = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListArrAirportInfo struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) SetAirportCode(v string) *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) SetAirportName(v string) *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) SetCityCode(v string) *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) SetCityName(v string) *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) SetTerminal(v string) *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListDepAirportInfo struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) SetAirportCode(v string) *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) SetAirportName(v string) *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) SetCityCode(v string) *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) SetCityName(v string) *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) SetTerminal(v string) *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleList struct {
	// example:
	//
	// demo
	BaggageInfo    *string                                                                      `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	BaggageItem    *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem    `json:"baggage_item,omitempty" xml:"baggage_item,omitempty" type:"Struct"`
	ChangeRule     *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule     `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Struct"`
	ChangeRuleItem *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Extra          *string                                                                      `json:"extra,omitempty" xml:"extra,omitempty"`
	RefundRule     *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Struct"`
	RefundRuleItem *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty" type:"Struct"`
	SignRule       *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule       `json:"sign_rule,omitempty" xml:"sign_rule,omitempty" type:"Struct"`
	TuigaiqianInfo *string                                                                      `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	UpgradeRule    *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule    `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty" type:"Struct"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleList) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetBaggageInfo() *string {
	return s.BaggageInfo
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetBaggageItem() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	return s.BaggageItem
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetChangeRule() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule {
	return s.ChangeRule
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetChangeRuleItem() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	return s.ChangeRuleItem
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetExtra() *string {
	return s.Extra
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetRefundRule() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule {
	return s.RefundRule
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetRefundRuleItem() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	return s.RefundRuleItem
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetSignRule() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule {
	return s.SignRule
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) GetUpgradeRule() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	return s.UpgradeRule
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetBaggageInfo(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.BaggageInfo = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetBaggageItem(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.BaggageItem = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetChangeRule(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.ChangeRule = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetChangeRuleItem(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.ChangeRuleItem = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetExtra(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.Extra = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetRefundRule(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.RefundRule = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetRefundRuleItem(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.RefundRuleItem = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetSignRule(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.SignRule = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetTuigaiqianInfo(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) SetUpgradeRule(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) *FlightListingSearchResponseBodyModuleFlightListFlightRuleList {
	s.UpgradeRule = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleList) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem struct {
	BaggageSubItems []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index     *int32                                                                        `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                       `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                       `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetBaggageSubItems() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTips() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	return s.Tips
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) GetType() *int32 {
	return s.Type
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetBaggageSubItems(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.BaggageSubItems = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetIndex(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Index = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTableHead(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.TableHead = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTips(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Tips = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) SetType(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Type = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItem) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                                         `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetIsStruct(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetPtc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BaggageSubContentType *int32                                                                                                                         `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetLogo(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.Logo = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetTipsDesc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetTipsImage(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.TipsImage = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListBaggageItemTips) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule struct {
	// example:
	//
	// false
	Able *bool                                                                          `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) GetInfo() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	return s.Info
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) SetAble(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule {
	s.Able = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) SetInfo(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule {
	s.Info = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetContent(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetCost(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetCostPercent(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTimeStamp(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTimeType(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem struct {
	ExtraContents []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                       `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                    `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                      `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                      `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetExtraContents() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetRefundSubItems() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetExtraContents(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetIndex(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Index = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetRefundSubItems(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetSubTableHead(v []*string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetTableHead(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetType(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Type = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) SetContent(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems struct {
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
	Ptc               *string                                                                                                       `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                       `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetRefundSubContents() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetIsStruct(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetPtc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetRefundSubContents(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule struct {
	// example:
	//
	// false
	Able *bool                                                                          `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) GetAble() *bool {
	return s.Able
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) GetInfo() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	return s.Info
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) SetAble(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule {
	s.Able = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) SetInfo(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule {
	s.Info = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetContent(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetCost(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetCostPercent(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTimeStamp(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTimeType(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem struct {
	ExtraContents []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index          *int32                                                                                       `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                    `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                      `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                      `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetExtraContents() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetRefundSubItems() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetExtraContents(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetIndex(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Index = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetRefundSubItems(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetSubTableHead(v []*string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetTableHead(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetType(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Type = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) SetContent(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                       `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                       `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetRefundSubContents() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetIsStruct(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetPtc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetRefundSubContents(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule struct {
	// example:
	//
	// false
	Able *bool                                                                        `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) GetAble() *bool {
	return s.Able
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) GetInfo() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	return s.Info
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) SetAble(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule {
	s.Able = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) SetInfo(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule {
	s.Info = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRule) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetContent(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetCost(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetCostPercent(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTimeStamp(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTimeType(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListSignRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule struct {
	// example:
	//
	// false
	Able *bool                                                                           `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) GetInfo() []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	return s.Info
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) SetAble(v bool) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	s.Able = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) SetInfo(v []*FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	s.Info = v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRule) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo struct {
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

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetContent(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetCost(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetCostPercent(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTimeStamp(v int32) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTimeType(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTitle(v string) *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightListingSearchResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) Validate() error {
	return dara.Validate(s)
}
