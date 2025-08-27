// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOrderQueryResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOrderQueryResponseBody
	GetMessage() *string
	SetModule(v *FlightOrderQueryResponseBodyModule) *FlightOrderQueryResponseBody
	GetModule() *FlightOrderQueryResponseBodyModule
	SetRequestId(v string) *FlightOrderQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOrderQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOrderQueryResponseBody
	GetTraceId() *string
}

type FlightOrderQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightOrderQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s FlightOrderQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOrderQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOrderQueryResponseBody) GetModule() *FlightOrderQueryResponseBodyModule {
	return s.Module
}

func (s *FlightOrderQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOrderQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOrderQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOrderQueryResponseBody) SetCode(v string) *FlightOrderQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetMessage(v string) *FlightOrderQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetModule(v *FlightOrderQueryResponseBodyModule) *FlightOrderQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderQueryResponseBody) SetRequestId(v string) *FlightOrderQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetSuccess(v bool) *FlightOrderQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderQueryResponseBody) SetTraceId(v string) *FlightOrderQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOrderQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModule struct {
	FlightChangeTicketInfoList []*FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList `json:"flight_change_ticket_info_list,omitempty" xml:"flight_change_ticket_info_list,omitempty" type:"Repeated"`
	FlightInfoList             []*FlightOrderQueryResponseBodyModuleFlightInfoList             `json:"flight_info_list,omitempty" xml:"flight_info_list,omitempty" type:"Repeated"`
	FlightRefundTicketInfoList []*FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList `json:"flight_refund_ticket_info_list,omitempty" xml:"flight_refund_ticket_info_list,omitempty" type:"Repeated"`
	FlightTicketInfoList       []*FlightOrderQueryResponseBodyModuleFlightTicketInfoList       `json:"flight_ticket_info_list,omitempty" xml:"flight_ticket_info_list,omitempty" type:"Repeated"`
	InsuranceInfoList          []*FlightOrderQueryResponseBodyModuleInsuranceInfoList          `json:"insurance_info_list,omitempty" xml:"insurance_info_list,omitempty" type:"Repeated"`
	InvoiceInfo                *FlightOrderQueryResponseBodyModuleInvoiceInfo                  `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo              *FlightOrderQueryResponseBodyModuleOrderBaseInfo                `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerInfoList          []*FlightOrderQueryResponseBodyModulePassengerInfoList          `json:"passenger_info_list,omitempty" xml:"passenger_info_list,omitempty" type:"Repeated"`
	PriceInfoList              []*FlightOrderQueryResponseBodyModulePriceInfoList              `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
}

func (s FlightOrderQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModule) GetFlightChangeTicketInfoList() []*FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	return s.FlightChangeTicketInfoList
}

func (s *FlightOrderQueryResponseBodyModule) GetFlightInfoList() []*FlightOrderQueryResponseBodyModuleFlightInfoList {
	return s.FlightInfoList
}

func (s *FlightOrderQueryResponseBodyModule) GetFlightRefundTicketInfoList() []*FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	return s.FlightRefundTicketInfoList
}

func (s *FlightOrderQueryResponseBodyModule) GetFlightTicketInfoList() []*FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	return s.FlightTicketInfoList
}

func (s *FlightOrderQueryResponseBodyModule) GetInsuranceInfoList() []*FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	return s.InsuranceInfoList
}

func (s *FlightOrderQueryResponseBodyModule) GetInvoiceInfo() *FlightOrderQueryResponseBodyModuleInvoiceInfo {
	return s.InvoiceInfo
}

func (s *FlightOrderQueryResponseBodyModule) GetOrderBaseInfo() *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	return s.OrderBaseInfo
}

func (s *FlightOrderQueryResponseBodyModule) GetPassengerInfoList() []*FlightOrderQueryResponseBodyModulePassengerInfoList {
	return s.PassengerInfoList
}

func (s *FlightOrderQueryResponseBodyModule) GetPriceInfoList() []*FlightOrderQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightChangeTicketInfoList(v []*FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightChangeTicketInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightInfoList(v []*FlightOrderQueryResponseBodyModuleFlightInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightRefundTicketInfoList(v []*FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightRefundTicketInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetFlightTicketInfoList(v []*FlightOrderQueryResponseBodyModuleFlightTicketInfoList) *FlightOrderQueryResponseBodyModule {
	s.FlightTicketInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetInsuranceInfoList(v []*FlightOrderQueryResponseBodyModuleInsuranceInfoList) *FlightOrderQueryResponseBodyModule {
	s.InsuranceInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetInvoiceInfo(v *FlightOrderQueryResponseBodyModuleInvoiceInfo) *FlightOrderQueryResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetOrderBaseInfo(v *FlightOrderQueryResponseBodyModuleOrderBaseInfo) *FlightOrderQueryResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetPassengerInfoList(v []*FlightOrderQueryResponseBodyModulePassengerInfoList) *FlightOrderQueryResponseBodyModule {
	s.PassengerInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) SetPriceInfoList(v []*FlightOrderQueryResponseBodyModulePriceInfoList) *FlightOrderQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList struct {
	ApplyId        *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrAirport     *string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// V
	ChangeCabin *string `json:"change_cabin,omitempty" xml:"change_cabin,omitempty"`
	// example:
	//
	// Y
	ChangeCabinLevel *string `json:"change_cabin_level,omitempty" xml:"change_cabin_level,omitempty"`
	// example:
	//
	// 100
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// example:
	//
	// MU7767
	ChangeFlightNo *string `json:"change_flight_no,omitempty" xml:"change_flight_no,omitempty"`
	// example:
	//
	// 33576
	ChangeOrderId *int64  `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	ChangeReason  *string `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	// example:
	//
	// 0
	ChangeType     *int32  `json:"change_type,omitempty" xml:"change_type,omitempty"`
	DepAirport     *string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime  *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Discount *string `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtModify    *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	JourneyIndex *int32  `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 00-123123
	OriginTicketNo *string `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
	OutApplyId     *string `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
	SegmentIndex   *int32  `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	StopCity       *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// 000-123123
	TicketNo         *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketStatus     *string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	TicketStatusCode *int32  `json:"ticket_status_code,omitempty" xml:"ticket_status_code,omitempty"`
	// example:
	//
	// 100
	UpgradeFee *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetApplyId() *string {
	return s.ApplyId
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetArrAirport() *string {
	return s.ArrAirport
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeCabin() *string {
	return s.ChangeCabin
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeCabinLevel() *string {
	return s.ChangeCabinLevel
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeFlightNo() *string {
	return s.ChangeFlightNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeOrderId() *int64 {
	return s.ChangeOrderId
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeReason() *string {
	return s.ChangeReason
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetDepAirport() *string {
	return s.DepAirport
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetDiscount() *string {
	return s.Discount
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetOriginTicketNo() *string {
	return s.OriginTicketNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetTicketStatusCode() *int32 {
	return s.TicketStatusCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) GetUpgradeFee() *float64 {
	return s.UpgradeFee
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetApplyId(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetArrAirport(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ArrAirport = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetArrAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetArrCity(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetArrCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetArrTime(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeCabin(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeCabin = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeCabinLevel(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeCabinLevel = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeFee(v float64) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeFlightNo(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeFlightNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeOrderId(v int64) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeOrderId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeReason(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeReason = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetChangeType(v int32) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.ChangeType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDepAirport(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.DepAirport = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDepAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDepCity(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDepCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDepTime(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetDiscount(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.Discount = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetJourneyIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetOriginTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.OriginTicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetOutApplyId(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.OutApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetSegmentIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetStopCity(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.StopCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetTicketStatus(v string) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetTicketStatusCode(v int32) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.TicketStatusCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) SetUpgradeFee(v float64) *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList {
	s.UpgradeFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightChangeTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightInfoList struct {
	// example:
	//
	// MU
	AirlineCode          *string                                                               `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName          *string                                                               `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrAirportCityCounty *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty `json:"arr_airport_city_county,omitempty" xml:"arr_airport_city_county,omitempty" type:"Struct"`
	// example:
	//
	// LHW
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrAirportName *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	ArrCityAdCode  *string `json:"arr_city_ad_code,omitempty" xml:"arr_city_ad_code,omitempty"`
	// example:
	//
	// LHW
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinLevel           *string                                                               `json:"cabin_level,omitempty" xml:"cabin_level,omitempty"`
	DepAirportCityCounty *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty `json:"dep_airport_city_county,omitempty" xml:"dep_airport_city_county,omitempty" type:"Struct"`
	// example:
	//
	// NGB
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepAirportName *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	DepCityAdCode  *string `json:"dep_city_ad_code,omitempty" xml:"dep_city_ad_code,omitempty"`
	// example:
	//
	// NGB
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 900
	FlightMile *int32 `json:"flight_mile,omitempty" xml:"flight_mile,omitempty"`
	// example:
	//
	// MU3849
	FlightNo         *string                                                             `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	JourneyIndex     *int32                                                              `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex     *int32                                                              `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	StopCity         []*string                                                           `json:"stop_city,omitempty" xml:"stop_city,omitempty" type:"Repeated"`
	StopCityInfoList []*FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList `json:"stop_city_info_list,omitempty" xml:"stop_city_info_list,omitempty" type:"Repeated"`
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrAirportCityCounty() *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	return s.ArrAirportCityCounty
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrCityAdCode() *string {
	return s.ArrCityAdCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetCabinLevel() *string {
	return s.CabinLevel
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepAirportCityCounty() *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	return s.DepAirportCityCounty
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepCityAdCode() *string {
	return s.DepCityAdCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetFlightMile() *int32 {
	return s.FlightMile
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetStopCity() []*string {
	return s.StopCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) GetStopCityInfoList() []*FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList {
	return s.StopCityInfoList
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetAirlineCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.AirlineCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetAirlineName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.AirlineName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrAirportCityCounty(v *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrAirportCityCounty = v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrAirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrCityAdCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrCityAdCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrTerminal(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrTerminal = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetArrTime(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetCabin(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetCabinLevel(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.CabinLevel = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepAirportCityCounty(v *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepAirportCityCounty = v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepAirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepCityAdCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepCityAdCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepTerminal(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepTerminal = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetDepTime(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetFlightMile(v int32) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.FlightMile = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetFlightNo(v string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetJourneyIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetSegmentIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetStopCity(v []*string) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.StopCity = v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) SetStopCityInfoList(v []*FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) *FlightOrderQueryResponseBodyModuleFlightInfoList {
	s.StopCityInfoList = v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty struct {
	Adcode                *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	AirportCityCode       *string `json:"airport_city_code,omitempty" xml:"airport_city_code,omitempty"`
	AirportCityName       *string `json:"airport_city_name,omitempty" xml:"airport_city_name,omitempty"`
	AirportCode           *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName           *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportParentCityName *string `json:"airport_parent_city_name,omitempty" xml:"airport_parent_city_name,omitempty"`
	CountyCityAdcode      *string `json:"county_city_adcode,omitempty" xml:"county_city_adcode,omitempty"`
	CountyCityName        *string `json:"county_city_name,omitempty" xml:"county_city_name,omitempty"`
	PrefectureCityAdcode  *string `json:"prefecture_city_adcode,omitempty" xml:"prefecture_city_adcode,omitempty"`
	PrefectureCityName    *string `json:"prefecture_city_name,omitempty" xml:"prefecture_city_name,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetAdcode() *string {
	return s.Adcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetAirportCityCode() *string {
	return s.AirportCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetAirportCityName() *string {
	return s.AirportCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetAirportParentCityName() *string {
	return s.AirportParentCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetCountyCityAdcode() *string {
	return s.CountyCityAdcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetCountyCityName() *string {
	return s.CountyCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetPrefectureCityAdcode() *string {
	return s.PrefectureCityAdcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) GetPrefectureCityName() *string {
	return s.PrefectureCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.Adcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetAirportCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.AirportCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetAirportCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.AirportCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.AirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.AirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetAirportParentCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.AirportParentCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetCountyCityAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.CountyCityAdcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetCountyCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.CountyCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetPrefectureCityAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.PrefectureCityAdcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) SetPrefectureCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty {
	s.PrefectureCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListArrAirportCityCounty) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty struct {
	Adcode                *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	AirportCityCode       *string `json:"airport_city_code,omitempty" xml:"airport_city_code,omitempty"`
	AirportCityName       *string `json:"airport_city_name,omitempty" xml:"airport_city_name,omitempty"`
	AirportCode           *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName           *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportParentCityName *string `json:"airport_parent_city_name,omitempty" xml:"airport_parent_city_name,omitempty"`
	CountyCityAdcode      *string `json:"county_city_adcode,omitempty" xml:"county_city_adcode,omitempty"`
	CountyCityName        *string `json:"county_city_name,omitempty" xml:"county_city_name,omitempty"`
	PrefectureCityAdcode  *string `json:"prefecture_city_adcode,omitempty" xml:"prefecture_city_adcode,omitempty"`
	PrefectureCityName    *string `json:"prefecture_city_name,omitempty" xml:"prefecture_city_name,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetAdcode() *string {
	return s.Adcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetAirportCityCode() *string {
	return s.AirportCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetAirportCityName() *string {
	return s.AirportCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetAirportParentCityName() *string {
	return s.AirportParentCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetCountyCityAdcode() *string {
	return s.CountyCityAdcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetCountyCityName() *string {
	return s.CountyCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetPrefectureCityAdcode() *string {
	return s.PrefectureCityAdcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) GetPrefectureCityName() *string {
	return s.PrefectureCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.Adcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetAirportCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.AirportCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetAirportCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.AirportCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.AirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.AirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetAirportParentCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.AirportParentCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetCountyCityAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.CountyCityAdcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetCountyCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.CountyCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetPrefectureCityAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.PrefectureCityAdcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) SetPrefectureCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty {
	s.PrefectureCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListDepAirportCityCounty) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList struct {
	StopAirportCityCounty *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty `json:"stop_airport_city_county,omitempty" xml:"stop_airport_city_county,omitempty" type:"Struct"`
	StopAirportCode       *string                                                                                `json:"stop_airport_code,omitempty" xml:"stop_airport_code,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) GetStopAirportCityCounty() *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	return s.StopAirportCityCounty
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) GetStopAirportCode() *string {
	return s.StopAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) SetStopAirportCityCounty(v *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList {
	s.StopAirportCityCounty = v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) SetStopAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList {
	s.StopAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty struct {
	Adcode                *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	AirportCityCode       *string `json:"airport_city_code,omitempty" xml:"airport_city_code,omitempty"`
	AirportCityName       *string `json:"airport_city_name,omitempty" xml:"airport_city_name,omitempty"`
	AirportCode           *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName           *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportParentCityName *string `json:"airport_parent_city_name,omitempty" xml:"airport_parent_city_name,omitempty"`
	CountyCityAdcode      *string `json:"county_city_adcode,omitempty" xml:"county_city_adcode,omitempty"`
	CountyCityName        *string `json:"county_city_name,omitempty" xml:"county_city_name,omitempty"`
	PrefectureCityAdcode  *string `json:"prefecture_city_adcode,omitempty" xml:"prefecture_city_adcode,omitempty"`
	PrefectureCityName    *string `json:"prefecture_city_name,omitempty" xml:"prefecture_city_name,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetAdcode() *string {
	return s.Adcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetAirportCityCode() *string {
	return s.AirportCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetAirportCityName() *string {
	return s.AirportCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetAirportParentCityName() *string {
	return s.AirportParentCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetCountyCityAdcode() *string {
	return s.CountyCityAdcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetCountyCityName() *string {
	return s.CountyCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetPrefectureCityAdcode() *string {
	return s.PrefectureCityAdcode
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) GetPrefectureCityName() *string {
	return s.PrefectureCityName
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.Adcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetAirportCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.AirportCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetAirportCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.AirportCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.AirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetAirportName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.AirportName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetAirportParentCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.AirportParentCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetCountyCityAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.CountyCityAdcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetCountyCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.CountyCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetPrefectureCityAdcode(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.PrefectureCityAdcode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) SetPrefectureCityName(v string) *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty {
	s.PrefectureCityName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightInfoListStopCityInfoListStopAirportCityCounty) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList struct {
	ApplyId                *string  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrAirport             *string  `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	ArrAirportCode         *string  `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity                *string  `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode            *string  `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	CompanyRefundTicketFee *float64 `json:"company_refund_ticket_fee,omitempty" xml:"company_refund_ticket_fee,omitempty"`
	DepAirport             *string  `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	DepAirportCode         *string  `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepCity                *string  `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode            *string  `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	FlightNo               *string  `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtModify               *string  `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	JourneyIndex            *int32   `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	OutApplyId              *string  `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
	PersonalRefundTicketFee *float64 `json:"personal_refund_ticket_fee,omitempty" xml:"personal_refund_ticket_fee,omitempty"`
	// example:
	//
	// 43667
	RefundOrderId *int64  `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	RefundReason  *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// example:
	//
	// 100
	RefundTicketFee *float64 `json:"refund_ticket_fee,omitempty" xml:"refund_ticket_fee,omitempty"`
	// example:
	//
	// 0
	RefundType   *int32 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// 000-13232
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetApplyId() *string {
	return s.ApplyId
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetArrAirport() *string {
	return s.ArrAirport
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetCompanyRefundTicketFee() *float64 {
	return s.CompanyRefundTicketFee
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetDepAirport() *string {
	return s.DepAirport
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetPersonalRefundTicketFee() *float64 {
	return s.PersonalRefundTicketFee
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetRefundOrderId() *int64 {
	return s.RefundOrderId
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetRefundReason() *string {
	return s.RefundReason
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetRefundTicketFee() *float64 {
	return s.RefundTicketFee
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetRefundType() *int32 {
	return s.RefundType
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetApplyId(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetArrAirport(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.ArrAirport = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetArrAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetArrCity(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetArrCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetCompanyRefundTicketFee(v float64) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.CompanyRefundTicketFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetDepAirport(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.DepAirport = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetDepAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetDepCity(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetDepCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetFlightNo(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetJourneyIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetOutApplyId(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.OutApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetPersonalRefundTicketFee(v float64) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.PersonalRefundTicketFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundOrderId(v int64) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundOrderId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundReason(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundReason = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundTicketFee(v float64) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundTicketFee = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetRefundType(v int32) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.RefundType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetSegmentIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) SetTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightRefundTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleFlightTicketInfoList struct {
	ArrAirport     *string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 100
	BuildPrice *float64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// example:
	//
	// false
	Changed        *bool   `json:"changed,omitempty" xml:"changed,omitempty"`
	DepAirport     *string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 10
	Discount *int32  `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtModify    *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	JourneyIndex *int32  `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 100
	OilPrice *float64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// example:
	//
	// 1
	PayType       *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	PersonalPrice *float64 `json:"personal_price,omitempty" xml:"personal_price,omitempty"`
	SegmentIndex  *int32   `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// 100
	SettlePrice *float64 `json:"settle_price,omitempty" xml:"settle_price,omitempty"`
	// example:
	//
	// 000-123
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 100
	TicketPrice *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// OPEN
	TicketStatus *string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	// example:
	//
	// 1
	TicketStatusCode *int32  `json:"ticket_status_code,omitempty" xml:"ticket_status_code,omitempty"`
	UserId           *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleFlightTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetArrAirport() *string {
	return s.ArrAirport
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetBuildPrice() *float64 {
	return s.BuildPrice
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetChanged() *bool {
	return s.Changed
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetDepAirport() *string {
	return s.DepAirport
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetDiscount() *int32 {
	return s.Discount
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetOilPrice() *float64 {
	return s.OilPrice
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetPersonalPrice() *float64 {
	return s.PersonalPrice
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetSettlePrice() *float64 {
	return s.SettlePrice
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetTicketPrice() *float64 {
	return s.TicketPrice
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetTicketStatusCode() *int32 {
	return s.TicketStatusCode
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetArrAirport(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.ArrAirport = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetArrAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetArrCity(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetArrCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetBuildPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.BuildPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetChanged(v bool) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.Changed = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetDepAirport(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.DepAirport = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetDepAirportCode(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetDepCity(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetDepCityCode(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetDiscount(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.Discount = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetFlightNo(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetJourneyIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetOilPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.OilPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetPayType(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetPersonalPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.PersonalPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetSegmentIndex(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetSettlePrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.SettlePrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketNo(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketPrice(v float64) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketPrice = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketStatus(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetTicketStatusCode(v int32) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.TicketStatusCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) SetUserId(v string) *FlightOrderQueryResponseBodyModuleFlightTicketInfoList {
	s.UserId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleFlightTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleInsuranceInfoList struct {
	// example:
	//
	// 100
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 345-987-098
	InsuranceNo *string `json:"insurance_no,omitempty" xml:"insurance_no,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleInsuranceInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleInsuranceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) GetAmount() *float64 {
	return s.Amount
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) GetInsuranceNo() *string {
	return s.InsuranceNo
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) GetName() *string {
	return s.Name
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) GetType() *string {
	return s.Type
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetAmount(v float64) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Amount = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetInsuranceNo(v string) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.InsuranceNo = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetName(v string) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Name = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetStatus(v int32) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Status = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) SetType(v string) *FlightOrderQueryResponseBodyModuleInsuranceInfoList {
	s.Type = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInsuranceInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleInvoiceInfo struct {
	// example:
	//
	// 43316
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) GetTitle() *string {
	return s.Title
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) SetId(v int64) *FlightOrderQueryResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) SetTitle(v string) *FlightOrderQueryResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModuleOrderBaseInfo struct {
	// example:
	//
	// 175634
	ApplyId                *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle             *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName            *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId                 *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName               *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId               *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName             *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExceedApplyId          *string `json:"exceed_apply_id,omitempty" xml:"exceed_apply_id,omitempty"`
	ExceedThirdPartApplyId *string `json:"exceed_third_part_apply_id,omitempty" xml:"exceed_third_part_apply_id,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// 13628
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// 146178
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 5
	OrderStatus *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	Supplier    *string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// example:
	//
	// CS-FLIGHT
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	ThirdpartCorpId     *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// example:
	//
	// CS-FLIGHT
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 0
	TripType *int32  `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderQueryResponseBodyModuleOrderBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetApplyId() *string {
	return s.ApplyId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetContactName() *string {
	return s.ContactName
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpName() *string {
	return s.CorpName
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetExceedApplyId() *string {
	return s.ExceedApplyId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetExceedThirdPartApplyId() *string {
	return s.ExceedThirdPartApplyId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtModify() *string {
	return s.GmtModify
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetSupplier() *string {
	return s.Supplier
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartCorpId() *string {
	return s.ThirdpartCorpId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetContactName(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ContactName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetExceedApplyId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ExceedApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetExceedThirdPartApplyId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ExceedThirdPartApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtModify(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtModify = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderId(v int64) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetSupplier(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.Supplier = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartBusinessId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartCorpId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartCorpId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetTripType(v int32) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.TripType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) SetUserId(v string) *FlightOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModuleOrderBaseInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModulePassengerInfoList struct {
	// example:
	//
	// 13446
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// CS-3345
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// example:
	//
	// CS-PROJECT
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 133576
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// CS-THIRDPROJECT
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	UserId             *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s FlightOrderQueryResponseBodyModulePassengerInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModulePassengerInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetUserName() *string {
	return s.UserName
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) GetUserType() *int32 {
	return s.UserType
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterId(v int64) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterName(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterNumber(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterNumber = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetProjectCode(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetProjectId(v int64) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetProjectTitle(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectTitle = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetThirdpartProjectId(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetUserId(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.UserId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetUserName(v string) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.UserName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) SetUserType(v int32) *FlightOrderQueryResponseBodyModulePassengerInfoList {
	s.UserType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePassengerInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate     *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 4
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 100
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// CS73290
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) SetType(v int32) *FlightOrderQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *FlightOrderQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}
