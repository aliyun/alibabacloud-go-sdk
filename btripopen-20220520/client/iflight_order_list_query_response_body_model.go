// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IFlightOrderListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *IFlightOrderListQueryResponseBody
	GetMessage() *string
	SetModule(v []*IFlightOrderListQueryResponseBodyModule) *IFlightOrderListQueryResponseBody
	GetModule() []*IFlightOrderListQueryResponseBodyModule
	SetPageInfo(v *IFlightOrderListQueryResponseBodyPageInfo) *IFlightOrderListQueryResponseBody
	GetPageInfo() *IFlightOrderListQueryResponseBodyPageInfo
	SetRequestId(v string) *IFlightOrderListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IFlightOrderListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IFlightOrderListQueryResponseBody
	GetTraceId() *string
}

type IFlightOrderListQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code     *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message  *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module   []*IFlightOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo *IFlightOrderListQueryResponseBodyPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-****-****-****-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce********056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IFlightOrderListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *IFlightOrderListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IFlightOrderListQueryResponseBody) GetModule() []*IFlightOrderListQueryResponseBodyModule {
	return s.Module
}

func (s *IFlightOrderListQueryResponseBody) GetPageInfo() *IFlightOrderListQueryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *IFlightOrderListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IFlightOrderListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IFlightOrderListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IFlightOrderListQueryResponseBody) SetCode(v string) *IFlightOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *IFlightOrderListQueryResponseBody) SetMessage(v string) *IFlightOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *IFlightOrderListQueryResponseBody) SetModule(v []*IFlightOrderListQueryResponseBodyModule) *IFlightOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *IFlightOrderListQueryResponseBody) SetPageInfo(v *IFlightOrderListQueryResponseBodyPageInfo) *IFlightOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *IFlightOrderListQueryResponseBody) SetRequestId(v string) *IFlightOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBody) SetSuccess(v bool) *IFlightOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *IFlightOrderListQueryResponseBody) SetTraceId(v string) *IFlightOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModule struct {
	FlightModifyOrderList []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList `json:"flight_modify_order_list,omitempty" xml:"flight_modify_order_list,omitempty" type:"Repeated"`
	FlightRefundOrderList []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList `json:"flight_refund_order_list,omitempty" xml:"flight_refund_order_list,omitempty" type:"Repeated"`
	FlightSaleOrder       *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder         `json:"flight_sale_order,omitempty" xml:"flight_sale_order,omitempty" type:"Struct"`
}

func (s IFlightOrderListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModule) GetFlightModifyOrderList() []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	return s.FlightModifyOrderList
}

func (s *IFlightOrderListQueryResponseBodyModule) GetFlightRefundOrderList() []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	return s.FlightRefundOrderList
}

func (s *IFlightOrderListQueryResponseBodyModule) GetFlightSaleOrder() *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	return s.FlightSaleOrder
}

func (s *IFlightOrderListQueryResponseBodyModule) SetFlightModifyOrderList(v []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) *IFlightOrderListQueryResponseBodyModule {
	s.FlightModifyOrderList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModule) SetFlightRefundOrderList(v []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) *IFlightOrderListQueryResponseBodyModule {
	s.FlightRefundOrderList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModule) SetFlightSaleOrder(v *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) *IFlightOrderListQueryResponseBodyModule {
	s.FlightSaleOrder = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList struct {
	CorpPayPrice                *int32                                                                                     `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	FlightModifySegmentList     []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList     `json:"flight_modify_segment_list,omitempty" xml:"flight_modify_segment_list,omitempty" type:"Repeated"`
	FlightOrderModifyTicketList []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList `json:"flight_order_modify_ticket_list,omitempty" xml:"flight_order_modify_ticket_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1005200138736000
	ModifyApplyId  *int64                                                                      `json:"modify_apply_id,omitempty" xml:"modify_apply_id,omitempty"`
	PassengerFee   []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee `json:"passenger_fee,omitempty" xml:"passenger_fee,omitempty" type:"Repeated"`
	PassengerList  []*string                                                                   `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PersonPayPrice *int32                                                                      `json:"person_pay_price,omitempty" xml:"person_pay_price,omitempty"`
	// example:
	//
	// 1005200138736028
	RelateModifyApplyId *int64 `json:"relate_modify_apply_id,omitempty" xml:"relate_modify_apply_id,omitempty"`
	ServiceFee          *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	TotalFee            *int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetCorpPayPrice() *int32 {
	return s.CorpPayPrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetFlightModifySegmentList() []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	return s.FlightModifySegmentList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetFlightOrderModifyTicketList() []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	return s.FlightOrderModifyTicketList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetModifyApplyId() *int64 {
	return s.ModifyApplyId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetPassengerFee() []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee {
	return s.PassengerFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetPassengerList() []*string {
	return s.PassengerList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetPersonPayPrice() *int32 {
	return s.PersonPayPrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetRelateModifyApplyId() *int64 {
	return s.RelateModifyApplyId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) GetTotalFee() *int64 {
	return s.TotalFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetCorpPayPrice(v int32) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.CorpPayPrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetFlightModifySegmentList(v []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.FlightModifySegmentList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetFlightOrderModifyTicketList(v []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.FlightOrderModifyTicketList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetModifyApplyId(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.ModifyApplyId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetPassengerFee(v []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.PassengerFee = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetPassengerList(v []*string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.PassengerList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetPersonPayPrice(v int32) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.PersonPayPrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetRelateModifyApplyId(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.RelateModifyApplyId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetServiceFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.ServiceFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) SetTotalFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList {
	s.TotalFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrApt      *string `json:"arr_apt,omitempty" xml:"arr_apt,omitempty"`
	// example:
	//
	// HGH
	ArrAptCode *string `json:"arr_apt_code,omitempty" xml:"arr_apt_code,omitempty"`
	ArrCity    *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T2
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2024-10-28 14:26:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// CA
	CarrierAirlineCode *string `json:"carrier_airline_code,omitempty" xml:"carrier_airline_code,omitempty"`
	CarrierAirlineName *string `json:"carrier_airline_name,omitempty" xml:"carrier_airline_name,omitempty"`
	DepApt             *string `json:"dep_apt,omitempty" xml:"dep_apt,omitempty"`
	// example:
	//
	// HGH
	DepAptCode *string `json:"dep_apt_code,omitempty" xml:"dep_apt_code,omitempty"`
	DepCity    *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T1
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2024-10-25 12:25:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// MU7384
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	Share        *bool  `json:"share,omitempty" xml:"share,omitempty"`
	// example:
	//
	// KIX
	StopAptCode *string `json:"stop_apt_code,omitempty" xml:"stop_apt_code,omitempty"`
	// example:
	//
	// 2024-10-26 11:25:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCity    *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// OSA
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	// example:
	//
	// 2024-10-27 11:26:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetShare() *bool {
	return s.Share
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopAptCode() *string {
	return s.StopAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetAirlineCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.AirlineCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetAirlineName(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.AirlineName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrApt(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrApt = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrTerminal(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetCarrierAirlineCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.CarrierAirlineCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetCarrierAirlineName(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.CarrierAirlineName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepApt(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepApt = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepTerminal(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepTerminal = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetJourneyIndex(v int32) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetSegmentIndex(v int32) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetShare(v bool) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.Share = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopArrTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopArrTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopDepTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopDepTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList struct {
	CabinClass []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	// example:
	//
	// MU5236
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 781-6605714721
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 01332
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetCabinClass() []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	return s.CabinClass
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetCabinClass(v []*IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.CabinClass = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetTicketNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.TicketNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass struct {
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// MU5236
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GetCabin() *string {
	return s.Cabin
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) SetCabin(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	s.Cabin = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) SetCabinClass(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	s.CabinClass = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee struct {
	ModifyHandFee    *int64 `json:"modify_hand_fee,omitempty" xml:"modify_hand_fee,omitempty"`
	ModifyUpgradeFee *int64 `json:"modify_upgrade_fee,omitempty" xml:"modify_upgrade_fee,omitempty"`
	TaxGap           *int64 `json:"tax_gap,omitempty" xml:"tax_gap,omitempty"`
	// example:
	//
	// 01332
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) GetModifyHandFee() *int64 {
	return s.ModifyHandFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) GetModifyUpgradeFee() *int64 {
	return s.ModifyUpgradeFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) GetTaxGap() *int64 {
	return s.TaxGap
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) SetModifyHandFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee {
	s.ModifyHandFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) SetModifyUpgradeFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee {
	s.ModifyUpgradeFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) SetTaxGap(v int64) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee {
	s.TaxGap = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightModifyOrderListPassengerFee) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList struct {
	CorpRefundAmount            *int32                                                                                     `json:"corp_refund_Amount,omitempty" xml:"corp_refund_Amount,omitempty"`
	FlightOrderRefundTicketList []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList `json:"flight_order_refund_ticket_list,omitempty" xml:"flight_order_refund_ticket_list,omitempty" type:"Repeated"`
	FlightRefundSegmentList     []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList     `json:"flight_refund_segment_list,omitempty" xml:"flight_refund_segment_list,omitempty" type:"Repeated"`
	PassengerFee                []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee                `json:"passenger_fee,omitempty" xml:"passenger_fee,omitempty" type:"Repeated"`
	PassengerInfo               []*string                                                                                  `json:"passenger_info,omitempty" xml:"passenger_info,omitempty" type:"Repeated"`
	PersonRefundAmount          *int32                                                                                     `json:"person_refund_Amount,omitempty" xml:"person_refund_Amount,omitempty"`
	RefundAmount                *int32                                                                                     `json:"refund_Amount,omitempty" xml:"refund_Amount,omitempty"`
	// example:
	//
	// 1006200138737069
	RefundApplyId *int64 `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	RefundHandFee *int32 `json:"refund_hand_fee,omitempty" xml:"refund_hand_fee,omitempty"`
	ServiceFee    *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetCorpRefundAmount() *int32 {
	return s.CorpRefundAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetFlightOrderRefundTicketList() []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	return s.FlightOrderRefundTicketList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetFlightRefundSegmentList() []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	return s.FlightRefundSegmentList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetPassengerFee() []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	return s.PassengerFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetPassengerInfo() []*string {
	return s.PassengerInfo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetPersonRefundAmount() *int32 {
	return s.PersonRefundAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetRefundAmount() *int32 {
	return s.RefundAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetRefundApplyId() *int64 {
	return s.RefundApplyId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetRefundHandFee() *int32 {
	return s.RefundHandFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetCorpRefundAmount(v int32) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.CorpRefundAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetFlightOrderRefundTicketList(v []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.FlightOrderRefundTicketList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetFlightRefundSegmentList(v []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.FlightRefundSegmentList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetPassengerFee(v []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.PassengerFee = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetPassengerInfo(v []*string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.PassengerInfo = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetPersonRefundAmount(v int32) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.PersonRefundAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetRefundAmount(v int32) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetRefundApplyId(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundApplyId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetRefundHandFee(v int32) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundHandFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) SetServiceFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList {
	s.ServiceFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList struct {
	CabinClass []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	// example:
	//
	// BK3162
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 999-6605133193
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 01332
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetCabinClass() []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	return s.CabinClass
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetCabinClass(v []*IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.CabinClass = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetTicketNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.TicketNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass struct {
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// BK3162
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GetCabin() *string {
	return s.Cabin
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) SetCabin(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	s.Cabin = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) SetCabinClass(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	s.CabinClass = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrApt      *string `json:"arr_apt,omitempty" xml:"arr_apt,omitempty"`
	// example:
	//
	// HKG
	ArrAptCode *string `json:"arr_apt_code,omitempty" xml:"arr_apt_code,omitempty"`
	ArrCity    *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T1
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2024-10-28 14:26:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// CA
	CarrierAirlineCode *string `json:"carrier_airline_code,omitempty" xml:"carrier_airline_code,omitempty"`
	CarrierAirlineName *string `json:"carrier_airline_name,omitempty" xml:"carrier_airline_name,omitempty"`
	DepApt             *string `json:"dep_apt,omitempty" xml:"dep_apt,omitempty"`
	// example:
	//
	// HGH
	DepAptCode *string `json:"dep_apt_code,omitempty" xml:"dep_apt_code,omitempty"`
	DepCity    *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T1
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2024-10-25 11:24:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// MU5334
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	Share        *bool  `json:"share,omitempty" xml:"share,omitempty"`
	// example:
	//
	// KIX
	StopAptCode *string `json:"stop_apt_code,omitempty" xml:"stop_apt_code,omitempty"`
	// example:
	//
	// 2024-10-26 11:25:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCity    *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// OSA
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	// example:
	//
	// 2024-10-27 11:26:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetShare() *bool {
	return s.Share
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopAptCode() *string {
	return s.StopAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetAirlineCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetAirlineName(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.AirlineName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrApt(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrApt = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrTerminal(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetCarrierAirlineCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.CarrierAirlineCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetCarrierAirlineName(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.CarrierAirlineName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepApt(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepApt = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepTerminal(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetJourneyIndex(v int32) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetSegmentIndex(v int32) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetShare(v bool) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.Share = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopArrTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopArrTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopDepTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopDepTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee struct {
	NoRefundModifyHandFee     *int64 `json:"no_refund_modify_hand_fee,omitempty" xml:"no_refund_modify_hand_fee,omitempty"`
	NoRefundModifyUpgradeFee  *int64 `json:"no_refund_modify_upgrade_fee,omitempty" xml:"no_refund_modify_upgrade_fee,omitempty"`
	RefundAmount              *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	RefundHandFee             *int64 `json:"refund_hand_fee,omitempty" xml:"refund_hand_fee,omitempty"`
	RefundModifyAmount        *int64 `json:"refund_modify_amount,omitempty" xml:"refund_modify_amount,omitempty"`
	RefundModifyHandAmount    *int64 `json:"refund_modify_hand_amount,omitempty" xml:"refund_modify_hand_amount,omitempty"`
	RefundModifyUpgradeAmount *int64 `json:"refund_modify_upgrade_amount,omitempty" xml:"refund_modify_upgrade_amount,omitempty"`
	RefundTaxHandFee          *int64 `json:"refund_tax_hand_fee,omitempty" xml:"refund_tax_hand_fee,omitempty"`
	// example:
	//
	// 01332
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetNoRefundModifyHandFee() *int64 {
	return s.NoRefundModifyHandFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetNoRefundModifyUpgradeFee() *int64 {
	return s.NoRefundModifyUpgradeFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetRefundAmount() *int64 {
	return s.RefundAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetRefundHandFee() *int64 {
	return s.RefundHandFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetRefundModifyAmount() *int64 {
	return s.RefundModifyAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetRefundModifyHandAmount() *int64 {
	return s.RefundModifyHandAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetRefundModifyUpgradeAmount() *int64 {
	return s.RefundModifyUpgradeAmount
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetRefundTaxHandFee() *int64 {
	return s.RefundTaxHandFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetNoRefundModifyHandFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.NoRefundModifyHandFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetNoRefundModifyUpgradeFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.NoRefundModifyUpgradeFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetRefundAmount(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.RefundAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetRefundHandFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.RefundHandFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetRefundModifyAmount(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.RefundModifyAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetRefundModifyHandAmount(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.RefundModifyHandAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetRefundModifyUpgradeAmount(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.RefundModifyUpgradeAmount = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetRefundTaxHandFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.RefundTaxHandFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightRefundOrderListPassengerFee) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrder struct {
	// example:
	//
	// 82587500
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 0
	BookType     *int32                                                            `json:"book_type,omitempty" xml:"book_type,omitempty"`
	BookerInfo   *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo `json:"booker_info,omitempty" xml:"booker_info,omitempty" type:"Struct"`
	CorpPayPrice *int64                                                            `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// example:
	//
	// 123
	ExceedApplyId         *string                                                                        `json:"exceed_apply_id,omitempty" xml:"exceed_apply_id,omitempty"`
	FlightOrderInsureList []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList `json:"flight_order_insure_list,omitempty" xml:"flight_order_insure_list,omitempty" type:"Repeated"`
	FlightOrderTicketList []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList `json:"flight_order_ticket_list,omitempty" xml:"flight_order_ticket_list,omitempty" type:"Repeated"`
	FlightSegmentList     []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList     `json:"flight_segment_list,omitempty" xml:"flight_segment_list,omitempty" type:"Repeated"`
	MixPay                *bool                                                                          `json:"mix_pay,omitempty" xml:"mix_pay,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	OrderCreateTime *string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// example:
	//
	// 1003038200110661039
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	OrderPayTime      *string `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
	OrderReservePrice *int64  `json:"order_reserve_price,omitempty" xml:"order_reserve_price,omitempty"`
	// example:
	//
	// 5
	OrderStatus     *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 1
	OrderType     *int32                                                                 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	PassengerList []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PayType        *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	PersonPayPrice *int64 `json:"person_pay_price,omitempty" xml:"person_pay_price,omitempty"`
	ServiceFee     *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// business_trip_api_000001
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetBookType() *int32 {
	return s.BookType
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetBookerInfo() *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo {
	return s.BookerInfo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetExceedApplyId() *string {
	return s.ExceedApplyId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetFlightOrderInsureList() []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	return s.FlightOrderInsureList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetFlightOrderTicketList() []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	return s.FlightOrderTicketList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetFlightSegmentList() []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	return s.FlightSegmentList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetMixPay() *bool {
	return s.MixPay
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderCreateTime() *string {
	return s.OrderCreateTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderPayTime() *string {
	return s.OrderPayTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderReservePrice() *int64 {
	return s.OrderReservePrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetOrderType() *int32 {
	return s.OrderType
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetPassengerList() []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	return s.PassengerList
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetPayType() *int32 {
	return s.PayType
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetPersonPayPrice() *int64 {
	return s.PersonPayPrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) GetTripType() *int32 {
	return s.TripType
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetApplyId(v int64) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.ApplyId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetBookType(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.BookType = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetBookerInfo(v *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.BookerInfo = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetCorpPayPrice(v int64) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.CorpPayPrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetExceedApplyId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.ExceedApplyId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetFlightOrderInsureList(v []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.FlightOrderInsureList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetFlightOrderTicketList(v []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.FlightOrderTicketList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetFlightSegmentList(v []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.FlightSegmentList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetMixPay(v bool) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.MixPay = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderCreateTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderCreateTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderPayTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderPayTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderReservePrice(v int64) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderReservePrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderStatus(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderStatus = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderStatusDesc(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderStatusDesc = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetOrderType(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.OrderType = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetPassengerList(v []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.PassengerList = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetPayType(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.PayType = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetPersonPayPrice(v int64) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.PersonPayPrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetServiceFee(v int64) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.ServiceFee = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetThirdPartApplyId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.ThirdPartApplyId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) SetTripType(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder {
	s.TripType = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrder) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo struct {
	// example:
	//
	// 01323
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) GetUserName() *string {
	return s.UserName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) SetUserName(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo {
	s.UserName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderBookerInfo) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList struct {
	// example:
	//
	// 17060573244016310
	InsOrderId    *string `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	InsPayType    *string `json:"ins_pay_type,omitempty" xml:"ins_pay_type,omitempty"`
	InsTotalPrice *int32  `json:"ins_total_price,omitempty" xml:"ins_total_price,omitempty"`
	// example:
	//
	// 1
	TradeAction *string `json:"trade_action,omitempty" xml:"trade_action,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetInsPayType() *string {
	return s.InsPayType
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetInsTotalPrice() *int32 {
	return s.InsTotalPrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetTradeAction() *string {
	return s.TradeAction
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetInsOrderId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.InsOrderId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetInsPayType(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.InsPayType = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetInsTotalPrice(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.InsTotalPrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetTradeAction(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.TradeAction = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList struct {
	CabinClass []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	Tax        *int32                                                                                   `json:"tax,omitempty" xml:"tax,omitempty"`
	// example:
	//
	// 018-6605785754
	TicketNo    *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketPrice *int32  `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 01323
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetCabinClass() []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	return s.CabinClass
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetTax() *int32 {
	return s.Tax
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetCabinClass(v []*IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.CabinClass = v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetTax(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.Tax = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetTicketNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.TicketNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetTicketPrice(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.TicketPrice = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass struct {
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// MF8765
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GetCabin() *string {
	return s.Cabin
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) SetCabin(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	s.Cabin = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) SetCabinClass(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	s.CabinClass = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrApt      *string `json:"arr_apt,omitempty" xml:"arr_apt,omitempty"`
	// example:
	//
	// HKG
	ArrAptCode *string `json:"arr_apt_code,omitempty" xml:"arr_apt_code,omitempty"`
	ArrCity    *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T2
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2024-10-25 15:26:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// CA
	CarrierAirlineCode *string `json:"carrier_airline_code,omitempty" xml:"carrier_airline_code,omitempty"`
	CarrierAirlineName *string `json:"carrier_airline_name,omitempty" xml:"carrier_airline_name,omitempty"`
	DepApt             *string `json:"dep_apt,omitempty" xml:"dep_apt,omitempty"`
	// example:
	//
	// HGH
	DepAptCode *string `json:"dep_apt_code,omitempty" xml:"dep_apt_code,omitempty"`
	DepCity    *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T1
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2024-10-24 15:26:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// MU5925
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	Share        *bool  `json:"share,omitempty" xml:"share,omitempty"`
	// example:
	//
	// KIX
	StopAptCode *string `json:"stop_apt_code,omitempty" xml:"stop_apt_code,omitempty"`
	// example:
	//
	// 2024-10-26 11:25:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCity    *string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// example:
	//
	// OSA
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	// example:
	//
	// 2024-10-27 11:26:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetShare() *bool {
	return s.Share
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopAptCode() *string {
	return s.StopAptCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetAirlineCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetAirlineName(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.AirlineName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrApt(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrApt = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrTerminal(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetCarrierAirlineCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.CarrierAirlineCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetCarrierAirlineName(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.CarrierAirlineName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepApt(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepApt = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepTerminal(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetFlightNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetJourneyIndex(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetSegmentIndex(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetShare(v bool) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.Share = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopAptCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopAptCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopArrTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopArrTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopCity(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopCity = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopCityCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopCityCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopDepTime(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopDepTime = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList struct {
	// example:
	//
	// 0111
	CostCenterId   *string `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// 1002
	DepartmentId   *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DepartmentName *string `json:"department_name,omitempty" xml:"department_name,omitempty"`
	// example:
	//
	// 01112
	InvoiceId    *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// 100757
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// example:
	//
	// 1111
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 01323
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetCostCenterId() *string {
	return s.CostCenterId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetInvoiceId() *string {
	return s.InvoiceId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) GetUserName() *string {
	return s.UserName
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetCostCenterId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.CostCenterId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetCostCenterName(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.CostCenterName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetDepartmentId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.DepartmentId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetDepartmentName(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.DepartmentName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetInvoiceId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.InvoiceId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetInvoiceTitle(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.InvoiceTitle = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetJobNo(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.JobNo = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetPassengerType(v int32) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.PassengerType = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetProjectCode(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.ProjectCode = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetProjectTitle(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.ProjectTitle = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetUserId(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) SetUserName(v string) *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.UserName = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyModuleFlightSaleOrderPassengerList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderListQueryResponseBodyPageInfo struct {
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4ZGViODFlYmM3MYzM4
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 100
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s IFlightOrderListQueryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) GetNumber() *int32 {
	return s.Number
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) GetScrollId() *string {
	return s.ScrollId
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) SetNumber(v int32) *IFlightOrderListQueryResponseBodyPageInfo {
	s.Number = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) SetScrollId(v string) *IFlightOrderListQueryResponseBodyPageInfo {
	s.ScrollId = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *IFlightOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *IFlightOrderListQueryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
