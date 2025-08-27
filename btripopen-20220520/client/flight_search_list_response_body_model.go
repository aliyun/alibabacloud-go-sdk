// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightSearchListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightSearchListResponseBody
	GetCode() *string
	SetMessage(v string) *FlightSearchListResponseBody
	GetMessage() *string
	SetModule(v *FlightSearchListResponseBodyModule) *FlightSearchListResponseBody
	GetModule() *FlightSearchListResponseBodyModule
	SetRequestId(v string) *FlightSearchListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightSearchListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightSearchListResponseBody
	GetTraceId() *string
}

type FlightSearchListResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightSearchListResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightSearchListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBody) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightSearchListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightSearchListResponseBody) GetModule() *FlightSearchListResponseBodyModule {
	return s.Module
}

func (s *FlightSearchListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightSearchListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightSearchListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightSearchListResponseBody) SetCode(v string) *FlightSearchListResponseBody {
	s.Code = &v
	return s
}

func (s *FlightSearchListResponseBody) SetMessage(v string) *FlightSearchListResponseBody {
	s.Message = &v
	return s
}

func (s *FlightSearchListResponseBody) SetModule(v *FlightSearchListResponseBodyModule) *FlightSearchListResponseBody {
	s.Module = v
	return s
}

func (s *FlightSearchListResponseBody) SetRequestId(v string) *FlightSearchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightSearchListResponseBody) SetSuccess(v bool) *FlightSearchListResponseBody {
	s.Success = &v
	return s
}

func (s *FlightSearchListResponseBody) SetTraceId(v string) *FlightSearchListResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightSearchListResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModule struct {
	FlightList []*FlightSearchListResponseBodyModuleFlightList `json:"flight_list,omitempty" xml:"flight_list,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsReplacePnr *bool `json:"is_replace_pnr,omitempty" xml:"is_replace_pnr,omitempty"`
}

func (s FlightSearchListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModule) GetFlightList() []*FlightSearchListResponseBodyModuleFlightList {
	return s.FlightList
}

func (s *FlightSearchListResponseBodyModule) GetIsReplacePnr() *bool {
	return s.IsReplacePnr
}

func (s *FlightSearchListResponseBodyModule) SetFlightList(v []*FlightSearchListResponseBodyModuleFlightList) *FlightSearchListResponseBodyModule {
	s.FlightList = v
	return s
}

func (s *FlightSearchListResponseBodyModule) SetIsReplacePnr(v bool) *FlightSearchListResponseBodyModule {
	s.IsReplacePnr = &v
	return s
}

func (s *FlightSearchListResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightList struct {
	AirlineInfo    *FlightSearchListResponseBodyModuleFlightListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightSearchListResponseBodyModuleFlightListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	ArrDate         *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	BasicCabinPrice *int32  `json:"basic_cabin_price,omitempty" xml:"basic_cabin_price,omitempty"`
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
	CabinClass    *string                                                      `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinInfoList []*FlightSearchListResponseBodyModuleFlightListCabinInfoList `json:"cabin_info_list,omitempty" xml:"cabin_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// FM
	CarrierAirline *string `json:"carrier_airline,omitempty" xml:"carrier_airline,omitempty"`
	// example:
	//
	// FM9152
	CarrierNo      *string                                                     `json:"carrier_no,omitempty" xml:"carrier_no,omitempty"`
	ClassRule      *string                                                     `json:"class_rule,omitempty" xml:"class_rule,omitempty"`
	DepAirportInfo *FlightSearchListResponseBodyModuleFlightListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
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
	FlightNo       *string                                                       `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightRuleList []*FlightSearchListResponseBodyModuleFlightListFlightRuleList `json:"flight_rule_list,omitempty" xml:"flight_rule_list,omitempty" type:"Repeated"`
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
	// {"key":"value"}
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 100
	OilPrice *int32 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// example:
	//
	// wisdiii2ii22ii2
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
	TotalPrice   *string                                                   `json:"total_price,omitempty" xml:"total_price,omitempty"`
	TransferInfo *FlightSearchListResponseBodyModuleFlightListTransferInfo `json:"transfer_info,omitempty" xml:"transfer_info,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightList) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetAirlineInfo() *FlightSearchListResponseBodyModuleFlightListAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetArrAirportInfo() *FlightSearchListResponseBodyModuleFlightListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetArrDate() *string {
	return s.ArrDate
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetCabinInfoList() []*FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	return s.CabinInfoList
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetCarrierAirline() *string {
	return s.CarrierAirline
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetCarrierNo() *string {
	return s.CarrierNo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetClassRule() *string {
	return s.ClassRule
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetDepAirportInfo() *FlightSearchListResponseBodyModuleFlightListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetDiscount() *int32 {
	return s.Discount
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetFlightRuleList() []*FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	return s.FlightRuleList
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetFlightRuleListStr() *string {
	return s.FlightRuleListStr
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetIsShare() *bool {
	return s.IsShare
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetIsStop() *bool {
	return s.IsStop
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetIsTransfer() *bool {
	return s.IsTransfer
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetMemo() *string {
	return s.Memo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetPrice() *int32 {
	return s.Price
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetProductType() *int64 {
	return s.ProductType
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetProductTypeDesc() *string {
	return s.ProductTypeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetPromotionPrice() *string {
	return s.PromotionPrice
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetRemainedSeatCount() *string {
	return s.RemainedSeatCount
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetSecretParams() *string {
	return s.SecretParams
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetSegmentNumber() *string {
	return s.SegmentNumber
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetTotalPrice() *string {
	return s.TotalPrice
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetTransferInfo() *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	return s.TransferInfo
}

func (s *FlightSearchListResponseBodyModuleFlightList) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetAirlineInfo(v *FlightSearchListResponseBodyModuleFlightListAirlineInfo) *FlightSearchListResponseBodyModuleFlightList {
	s.AirlineInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetArrAirportInfo(v *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) *FlightSearchListResponseBodyModuleFlightList {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetArrDate(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.ArrDate = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetBasicCabinPrice(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.BasicCabinPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetBuildPrice(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.BuildPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetCabin(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.Cabin = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetCabinClass(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetCabinInfoList(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoList) *FlightSearchListResponseBodyModuleFlightList {
	s.CabinInfoList = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetCarrierAirline(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.CarrierAirline = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetCarrierNo(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.CarrierNo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetClassRule(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.ClassRule = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetDepAirportInfo(v *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) *FlightSearchListResponseBodyModuleFlightList {
	s.DepAirportInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetDepCityCode(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.DepCityCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetDepDate(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.DepDate = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetDiscount(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.Discount = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetFlightNo(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetFlightRuleList(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleList) *FlightSearchListResponseBodyModuleFlightList {
	s.FlightRuleList = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetFlightRuleListStr(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.FlightRuleListStr = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetFlightSize(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.FlightSize = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetFlightType(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.FlightType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetInvoiceType(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.InvoiceType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetIsProtocol(v bool) *FlightSearchListResponseBodyModuleFlightList {
	s.IsProtocol = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetIsShare(v bool) *FlightSearchListResponseBodyModuleFlightList {
	s.IsShare = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetIsStop(v bool) *FlightSearchListResponseBodyModuleFlightList {
	s.IsStop = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetIsTransfer(v bool) *FlightSearchListResponseBodyModuleFlightList {
	s.IsTransfer = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetMealDesc(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.MealDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetMemo(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.Memo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetOilPrice(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.OilPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetOtaItemId(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.OtaItemId = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetPrice(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.Price = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetProductType(v int64) *FlightSearchListResponseBodyModuleFlightList {
	s.ProductType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetProductTypeDesc(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.ProductTypeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetPromotionPrice(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.PromotionPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetRemainedSeatCount(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.RemainedSeatCount = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetSecretParams(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.SecretParams = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetSegmentNumber(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.SegmentNumber = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetStopArrTime(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.StopArrTime = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetStopCity(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.StopCity = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetStopDepTime(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.StopDepTime = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetTicketPrice(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.TicketPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetTotalPrice(v string) *FlightSearchListResponseBodyModuleFlightList {
	s.TotalPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetTransferInfo(v *FlightSearchListResponseBodyModuleFlightListTransferInfo) *FlightSearchListResponseBodyModuleFlightList {
	s.TransferInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) SetTripType(v int32) *FlightSearchListResponseBodyModuleFlightList {
	s.TripType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) SetAirlineCode(v string) *FlightSearchListResponseBodyModuleFlightListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) SetAirlineName(v string) *FlightSearchListResponseBodyModuleFlightListAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) SetAirlineSimpleName(v string) *FlightSearchListResponseBodyModuleFlightListAirlineInfo {
	s.AirlineSimpleName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListArrAirportInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) SetAirportCode(v string) *FlightSearchListResponseBodyModuleFlightListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) SetAirportName(v string) *FlightSearchListResponseBodyModuleFlightListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) SetCityCode(v string) *FlightSearchListResponseBodyModuleFlightListArrAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) SetCityName(v string) *FlightSearchListResponseBodyModuleFlightListArrAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) SetTerminal(v string) *FlightSearchListResponseBodyModuleFlightListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoList struct {
	AgentId         *int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
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
	Discount       *string                                                                    `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightRuleList []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList `json:"flight_rule_list,omitempty" xml:"flight_rule_list,omitempty" type:"Repeated"`
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
	// {"key":"value"}
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
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
	// wisdiii2ii22ii2
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoList) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetBasicCabinPrice() *int32 {
	return s.BasicCabinPrice
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetBuildPrice() *int32 {
	return s.BuildPrice
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetChildCabin() *string {
	return s.ChildCabin
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetClassName() *string {
	return s.ClassName
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetClassRule() *string {
	return s.ClassRule
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetDiscount() *string {
	return s.Discount
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetFlightRuleList() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	return s.FlightRuleList
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetFlightRuleListStr() *string {
	return s.FlightRuleListStr
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetMemo() *string {
	return s.Memo
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetOilPrice() *int32 {
	return s.OilPrice
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetOrderParams() *string {
	return s.OrderParams
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetPrice() *int32 {
	return s.Price
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetProductType() *int64 {
	return s.ProductType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetProductTypeDesc() *string {
	return s.ProductTypeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetPromotionPrice() *string {
	return s.PromotionPrice
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetRemainedSeatCount() *string {
	return s.RemainedSeatCount
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) GetTotalPrice() *int32 {
	return s.TotalPrice
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetAgentId(v int64) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.AgentId = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetBasicCabinPrice(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.BasicCabinPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetBuildPrice(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.BuildPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetCabin(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.Cabin = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetCabinClass(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.CabinClass = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetCabinClassName(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.CabinClassName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetChildCabin(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.ChildCabin = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetClassName(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.ClassName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetClassRule(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.ClassRule = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetDiscount(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.Discount = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetFlightRuleList(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.FlightRuleList = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetFlightRuleListStr(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.FlightRuleListStr = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetInvoiceType(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.InvoiceType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetIsProtocol(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.IsProtocol = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetMemo(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.Memo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetOilPrice(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.OilPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetOrderParams(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.OrderParams = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetOtaItemId(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.OtaItemId = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetPrice(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.Price = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetProductType(v int64) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.ProductType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetProductTypeDesc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.ProductTypeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetPromotionPrice(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.PromotionPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetRemainedSeatCount(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.RemainedSeatCount = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetTicketPrice(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.TicketPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) SetTotalPrice(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoList {
	s.TotalPrice = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList struct {
	// example:
	//
	// demo
	BaggageInfo    *string                                                                                `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	BaggageItem    *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem    `json:"baggage_item,omitempty" xml:"baggage_item,omitempty" type:"Struct"`
	ChangeRule     *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule     `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Struct"`
	ChangeRuleItem *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Extra          *string                                                                                `json:"extra,omitempty" xml:"extra,omitempty"`
	RefundRule     *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Struct"`
	RefundRuleItem *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty" type:"Struct"`
	SignRule       *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule       `json:"sign_rule,omitempty" xml:"sign_rule,omitempty" type:"Struct"`
	TuigaiqianInfo *string                                                                                `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	UpgradeRule    *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule    `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty" type:"Struct"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetBaggageInfo() *string {
	return s.BaggageInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetBaggageItem() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	return s.BaggageItem
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetChangeRule() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule {
	return s.ChangeRule
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetChangeRuleItem() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	return s.ChangeRuleItem
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetExtra() *string {
	return s.Extra
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetRefundRule() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule {
	return s.RefundRule
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetRefundRuleItem() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	return s.RefundRuleItem
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetSignRule() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule {
	return s.SignRule
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) GetUpgradeRule() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule {
	return s.UpgradeRule
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetBaggageInfo(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.BaggageInfo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetBaggageItem(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.BaggageItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetChangeRule(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.ChangeRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetChangeRuleItem(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.ChangeRuleItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetExtra(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.Extra = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetRefundRule(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.RefundRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetRefundRuleItem(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.RefundRuleItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetSignRule(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.SignRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetTuigaiqianInfo(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) SetUpgradeRule(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList {
	s.UpgradeRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleList) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem struct {
	BaggageSubItems []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index     *int32                                                                                  `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                                 `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                                 `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetBaggageSubItems() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetTips() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	return s.Tips
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetBaggageSubItems(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.BaggageSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetTips(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Tips = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                                                   `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	BaggageSubContentType *int32                                                                                                                                   `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) SetLogo(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	s.Logo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) SetTipsDesc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) SetTipsImage(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips {
	s.TipsImage = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListBaggageItemTips) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule struct {
	// example:
	//
	// true
	Able *bool                                                                                    `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem struct {
	ExtraContents []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                                 `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                              `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                                `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                                `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetExtraContents() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetRefundSubItems() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetExtraContents(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetRefundSubItems(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetSubTableHead(v []*string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                 `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                 `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetRefundSubContents() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetRefundSubContents(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule struct {
	// example:
	//
	// true
	Able *bool                                                                                    `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem struct {
	ExtraContents []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index          *int32                                                                                                 `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                              `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                                `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetExtraContents() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetRefundSubItems() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetExtraContents(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetRefundSubItems(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetSubTableHead(v []*string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                 `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                 `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetRefundSubContents() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetRefundSubContents(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule struct {
	// example:
	//
	// true
	Able *bool                                                                                  `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListSignRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule struct {
	// example:
	//
	// true
	Able *bool                                                                                     `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListCabinInfoListFlightRuleListUpgradeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListDepAirportInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) SetAirportCode(v string) *FlightSearchListResponseBodyModuleFlightListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) SetAirportName(v string) *FlightSearchListResponseBodyModuleFlightListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) SetCityCode(v string) *FlightSearchListResponseBodyModuleFlightListDepAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) SetCityName(v string) *FlightSearchListResponseBodyModuleFlightListDepAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) SetTerminal(v string) *FlightSearchListResponseBodyModuleFlightListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleList struct {
	// example:
	//
	// demo
	BaggageInfo    *string                                                                   `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	BaggageItem    *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem    `json:"baggage_item,omitempty" xml:"baggage_item,omitempty" type:"Struct"`
	ChangeRule     *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule     `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Struct"`
	ChangeRuleItem *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Extra          *string                                                                   `json:"extra,omitempty" xml:"extra,omitempty"`
	RefundRule     *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Struct"`
	RefundRuleItem *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty" type:"Struct"`
	SignRule       *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule       `json:"sign_rule,omitempty" xml:"sign_rule,omitempty" type:"Struct"`
	TuigaiqianInfo *string                                                                   `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	UpgradeRule    *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule    `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty" type:"Struct"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleList) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleList) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetBaggageInfo() *string {
	return s.BaggageInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetBaggageItem() *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	return s.BaggageItem
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetChangeRule() *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule {
	return s.ChangeRule
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetChangeRuleItem() *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	return s.ChangeRuleItem
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetExtra() *string {
	return s.Extra
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetRefundRule() *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule {
	return s.RefundRule
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetRefundRuleItem() *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	return s.RefundRuleItem
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetSignRule() *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule {
	return s.SignRule
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) GetUpgradeRule() *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	return s.UpgradeRule
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetBaggageInfo(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.BaggageInfo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetBaggageItem(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.BaggageItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetChangeRule(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.ChangeRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetChangeRuleItem(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.ChangeRuleItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetExtra(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.Extra = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetRefundRule(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.RefundRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetRefundRuleItem(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.RefundRuleItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetSignRule(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.SignRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetTuigaiqianInfo(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) SetUpgradeRule(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) *FlightSearchListResponseBodyModuleFlightListFlightRuleList {
	s.UpgradeRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleList) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem struct {
	BaggageSubItems []*FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index     *int32                                                                     `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                    `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                    `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GetBaggageSubItems() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTips() *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	return s.Tips
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) SetBaggageSubItems(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.BaggageSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTips(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Tips = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                                      `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BaggageSubContentType *int32                                                                                                                      `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetLogo(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.Logo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetTipsDesc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) SetTipsImage(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips {
	s.TipsImage = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListBaggageItemTips) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule struct {
	// example:
	//
	// false
	Able *bool                                                                       `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem struct {
	ExtraContents []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                    `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                 `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                   `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                   `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetExtraContents() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetRefundSubItems() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetExtraContents(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetRefundSubItems(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetSubTableHead(v []*string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                    `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                    `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetRefundSubContents() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetRefundSubContents(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule struct {
	// example:
	//
	// false
	Able *bool                                                                       `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem struct {
	ExtraContents []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index          *int32                                                                                    `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                 `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                   `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                   `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetExtraContents() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetRefundSubItems() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetExtraContents(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetRefundSubItems(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetSubTableHead(v []*string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                    `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                    `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetRefundSubContents() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetRefundSubContents(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule struct {
	// example:
	//
	// false
	Able *bool                                                                     `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListSignRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule struct {
	// example:
	//
	// false
	Able *bool                                                                        `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListFlightRuleListUpgradeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfo struct {
	FlightSize *string `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	// example:
	//
	// demo
	FlightType             *string                                                                         `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	TransferAirlineInfo    *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo    `json:"transfer_airline_info,omitempty" xml:"transfer_airline_info,omitempty" type:"Struct"`
	TransferArrAirportInfo *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo `json:"transfer_arr_airport_info,omitempty" xml:"transfer_arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	TransferArrDate        *string                                                                         `json:"transfer_arr_date,omitempty" xml:"transfer_arr_date,omitempty"`
	TransferDepAirportInfo *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo `json:"transfer_dep_airport_info,omitempty" xml:"transfer_dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	TransferDepDate *string `json:"transfer_dep_date,omitempty" xml:"transfer_dep_date,omitempty"`
	// example:
	//
	// CA1234
	TransferFlightNo       *string                                                                           `json:"transfer_flight_no,omitempty" xml:"transfer_flight_no,omitempty"`
	TransferFlightRuleList []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList `json:"transfer_flight_rule_list,omitempty" xml:"transfer_flight_rule_list,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferAirlineInfo() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo {
	return s.TransferAirlineInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferArrAirportInfo() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo {
	return s.TransferArrAirportInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferArrDate() *string {
	return s.TransferArrDate
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferDepAirportInfo() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo {
	return s.TransferDepAirportInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferDepDate() *string {
	return s.TransferDepDate
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferFlightNo() *string {
	return s.TransferFlightNo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) GetTransferFlightRuleList() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	return s.TransferFlightRuleList
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetFlightSize(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.FlightSize = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetFlightType(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.FlightType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferAirlineInfo(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferAirlineInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferArrAirportInfo(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferArrAirportInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferArrDate(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferArrDate = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferDepAirportInfo(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferDepAirportInfo = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferDepDate(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferDepDate = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferFlightNo(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferFlightNo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) SetTransferFlightRuleList(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) *FlightSearchListResponseBodyModuleFlightListTransferInfo {
	s.TransferFlightRuleList = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo struct {
	// example:
	//
	// ZH
	AirlineCode       *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName       *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	AirlineSimpleName *string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) GetAirlineSimpleName() *string {
	return s.AirlineSimpleName
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) SetAirlineCode(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) SetAirlineName(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) SetAirlineSimpleName(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo {
	s.AirlineSimpleName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) SetAirportCode(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) SetAirportName(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) SetCityCode(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) SetCityName(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) SetTerminal(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) GetCityName() *string {
	return s.CityName
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) SetAirportCode(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) SetAirportName(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) SetCityCode(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo {
	s.CityCode = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) SetCityName(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo {
	s.CityName = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) SetTerminal(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList struct {
	// example:
	//
	// demo
	BaggageInfo    *string                                                                                       `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	BaggageItem    *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem    `json:"baggage_item,omitempty" xml:"baggage_item,omitempty" type:"Struct"`
	ChangeRule     *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule     `json:"change_rule,omitempty" xml:"change_rule,omitempty" type:"Struct"`
	ChangeRuleItem *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Extra          *string                                                                                       `json:"extra,omitempty" xml:"extra,omitempty"`
	RefundRule     *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" type:"Struct"`
	RefundRuleItem *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty" type:"Struct"`
	SignRule       *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule       `json:"sign_rule,omitempty" xml:"sign_rule,omitempty" type:"Struct"`
	// example:
	//
	// demo
	TuigaiqianInfo *string                                                                                    `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	UpgradeRule    *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty" type:"Struct"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetBaggageInfo() *string {
	return s.BaggageInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetBaggageItem() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	return s.BaggageItem
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetChangeRule() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule {
	return s.ChangeRule
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetChangeRuleItem() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	return s.ChangeRuleItem
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetExtra() *string {
	return s.Extra
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetRefundRule() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule {
	return s.RefundRule
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetRefundRuleItem() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	return s.RefundRuleItem
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetSignRule() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule {
	return s.SignRule
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetTuigaiqianInfo() *string {
	return s.TuigaiqianInfo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) GetUpgradeRule() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule {
	return s.UpgradeRule
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetBaggageInfo(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.BaggageInfo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetBaggageItem(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.BaggageItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetChangeRule(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.ChangeRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetChangeRuleItem(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.ChangeRuleItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetExtra(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.Extra = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetRefundRule(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.RefundRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetRefundRuleItem(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.RefundRuleItem = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetSignRule(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.SignRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetTuigaiqianInfo(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.TuigaiqianInfo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) SetUpgradeRule(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList {
	s.UpgradeRule = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleList) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem struct {
	BaggageSubItems []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index     *int32                                                                                         `json:"index,omitempty" xml:"index,omitempty"`
	TableHead *string                                                                                        `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Tips      *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips `json:"tips,omitempty" xml:"tips,omitempty" type:"Struct"`
	Title     *string                                                                                        `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GetBaggageSubItems() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems {
	return s.BaggageSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GetTips() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips {
	return s.Tips
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) SetBaggageSubItems(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	s.BaggageSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) SetTips(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	s.Tips = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems struct {
	BaggageSubContentVisualizes []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes,omitempty" type:"Repeated"`
	ExtraContentVisualizes      []interface{}                                                                                                                          `json:"extra_content_visualizes,omitempty" xml:"extra_content_visualizes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc *string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	// example:
	//
	// 1
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) GetBaggageSubContentVisualizes() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	return s.BaggageSubContentVisualizes
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) GetExtraContentVisualizes() []interface{} {
	return s.ExtraContentVisualizes
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) SetBaggageSubContentVisualizes(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems {
	s.BaggageSubContentVisualizes = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) SetExtraContentVisualizes(v []interface{}) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems {
	s.ExtraContentVisualizes = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes struct {
	BaggageDesc []*string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BaggageSubContentType *int32                                                                                                                                          `json:"baggage_sub_content_type,omitempty" xml:"baggage_sub_content_type,omitempty"`
	Description           *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	ImageDO               *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO     `json:"image_d_o,omitempty" xml:"image_d_o,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsHighlight *bool   `json:"is_highlight,omitempty" xml:"is_highlight,omitempty"`
	SubTitle    *string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageDesc() []*string {
	return s.BaggageDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetBaggageSubContentType() *int32 {
	return s.BaggageSubContentType
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetDescription() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	return s.Description
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetImageDO() *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	return s.ImageDO
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetIsHighlight() *bool {
	return s.IsHighlight
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) GetSubTitle() *string {
	return s.SubTitle
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageDesc(v []*string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageDesc = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetBaggageSubContentType(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.BaggageSubContentType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetDescription(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.Description = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetImageDO(v *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.ImageDO = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetIsHighlight(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.IsHighlight = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) SetSubTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes {
	s.SubTitle = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizes) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetDesc() *string {
	return s.Desc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetIcon() *string {
	return s.Icon
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetImage() *string {
	return s.Image
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetDesc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Desc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetIcon(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Icon = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetImage(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Image = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesDescription) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetImage() *string {
	return s.Image
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetLargest() *string {
	return s.Largest
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetMiddle() *string {
	return s.Middle
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) GetSmallest() *string {
	return s.Smallest
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetImage(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Image = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetLargest(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Largest = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetMiddle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Middle = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) SetSmallest(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO {
	s.Smallest = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemBaggageSubItemsBaggageSubContentVisualizesImageDO) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) GetLogo() *string {
	return s.Logo
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) GetTipsDesc() *string {
	return s.TipsDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) GetTipsImage() *string {
	return s.TipsImage
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) SetLogo(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips {
	s.Logo = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) SetTipsDesc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips {
	s.TipsDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) SetTipsImage(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips {
	s.TipsImage = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListBaggageItemTips) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule struct {
	// example:
	//
	// true
	Able *bool                                                                                           `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 100
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem struct {
	ExtraContents []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Index          *int32                                                                                                        `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                                     `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                                       `json:"table_head,omitempty" xml:"table_head,omitempty"`
	Title          *string                                                                                                       `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetExtraContents() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetRefundSubItems() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetExtraContents(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetRefundSubItems(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetSubTableHead(v []*string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                        `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                        `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) GetRefundSubContents() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) SetRefundSubContents(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListChangeRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule struct {
	// example:
	//
	// true
	Able *bool                                                                                           `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 100
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 100
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem struct {
	ExtraContents []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents `json:"extra_contents,omitempty" xml:"extra_contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Index          *int32                                                                                                        `json:"index,omitempty" xml:"index,omitempty"`
	RefundSubItems []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems `json:"refund_sub_items,omitempty" xml:"refund_sub_items,omitempty" type:"Repeated"`
	SubTableHead   []*string                                                                                                     `json:"sub_table_head,omitempty" xml:"sub_table_head,omitempty" type:"Repeated"`
	TableHead      *string                                                                                                       `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetExtraContents() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents {
	return s.ExtraContents
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetIndex() *int32 {
	return s.Index
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetRefundSubItems() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems {
	return s.RefundSubItems
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetSubTableHead() []*string {
	return s.SubTableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetTableHead() *string {
	return s.TableHead
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) GetType() *int32 {
	return s.Type
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetExtraContents(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.ExtraContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetIndex(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.Index = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetRefundSubItems(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.RefundSubItems = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetSubTableHead(v []*string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.SubTableHead = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetTableHead(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.TableHead = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) SetType(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem {
	s.Type = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItem) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents struct {
	// example:
	//
	// demo
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemExtraContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems struct {
	// example:
	//
	// true
	IsStruct *bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
	// PTC
	//
	// example:
	//
	// ADT
	Ptc               *string                                                                                                                        `json:"ptc,omitempty" xml:"ptc,omitempty"`
	RefundSubContents []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents,omitempty" type:"Repeated"`
	Title             *string                                                                                                                        `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) GetIsStruct() *bool {
	return s.IsStruct
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) GetPtc() *string {
	return s.Ptc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) GetRefundSubContents() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	return s.RefundSubContents
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) SetIsStruct(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems {
	s.IsStruct = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) SetPtc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems {
	s.Ptc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) SetRefundSubContents(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems {
	s.RefundSubContents = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents struct {
	FeeDesc  *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeRange *string `json:"fee_range,omitempty" xml:"fee_range,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetFeeRange() *string {
	return s.FeeRange
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) GetStyle() *int32 {
	return s.Style
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeDesc(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeDesc = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetFeeRange(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.FeeRange = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) SetStyle(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents {
	s.Style = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListRefundRuleItemRefundSubItemsRefundSubContents) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule struct {
	// example:
	//
	// true
	Able *bool                                                                                         `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListSignRuleInfo) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule struct {
	// example:
	//
	// true
	Able *bool                                                                                            `json:"able,omitempty" xml:"able,omitempty"`
	Info []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo `json:"info,omitempty" xml:"info,omitempty" type:"Repeated"`
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) GetAble() *bool {
	return s.Able
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) GetInfo() []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	return s.Info
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) SetAble(v bool) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule {
	s.Able = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) SetInfo(v []*FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule {
	s.Info = v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRule) Validate() error {
	return dara.Validate(s)
}

type FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo struct {
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

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GetContent() *string {
	return s.Content
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GetCost() *int32 {
	return s.Cost
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GetCostPercent() *int32 {
	return s.CostPercent
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GetTimeStamp() *int32 {
	return s.TimeStamp
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) SetContent(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	s.Content = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) SetCost(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	s.Cost = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) SetCostPercent(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	s.CostPercent = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) SetTimeStamp(v int32) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	s.TimeStamp = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) SetTimeType(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	s.TimeType = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) SetTitle(v string) *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo {
	s.Title = &v
	return s
}

func (s *FlightSearchListResponseBodyModuleFlightListTransferInfoTransferFlightRuleListUpgradeRuleInfo) Validate() error {
	return dara.Validate(s)
}
