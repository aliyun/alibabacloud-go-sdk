// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderDetailQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IFlightOrderDetailQueryResponseBody
	GetCode() *string
	SetMessage(v string) *IFlightOrderDetailQueryResponseBody
	GetMessage() *string
	SetModule(v *IFlightOrderDetailQueryResponseBodyModule) *IFlightOrderDetailQueryResponseBody
	GetModule() *IFlightOrderDetailQueryResponseBodyModule
	SetRequestId(v string) *IFlightOrderDetailQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IFlightOrderDetailQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IFlightOrderDetailQueryResponseBody
	GetTraceId() *string
}

type IFlightOrderDetailQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IFlightOrderDetailQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IFlightOrderDetailQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *IFlightOrderDetailQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IFlightOrderDetailQueryResponseBody) GetModule() *IFlightOrderDetailQueryResponseBodyModule {
	return s.Module
}

func (s *IFlightOrderDetailQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IFlightOrderDetailQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IFlightOrderDetailQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IFlightOrderDetailQueryResponseBody) SetCode(v string) *IFlightOrderDetailQueryResponseBody {
	s.Code = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBody) SetMessage(v string) *IFlightOrderDetailQueryResponseBody {
	s.Message = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBody) SetModule(v *IFlightOrderDetailQueryResponseBodyModule) *IFlightOrderDetailQueryResponseBody {
	s.Module = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBody) SetRequestId(v string) *IFlightOrderDetailQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBody) SetSuccess(v bool) *IFlightOrderDetailQueryResponseBody {
	s.Success = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBody) SetTraceId(v string) *IFlightOrderDetailQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModule struct {
	FlightModifyOrderList []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList `json:"flight_modify_order_list,omitempty" xml:"flight_modify_order_list,omitempty" type:"Repeated"`
	FlightRefundOrderList []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList `json:"flight_refund_order_list,omitempty" xml:"flight_refund_order_list,omitempty" type:"Repeated"`
	FlightSaleOrder       *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder         `json:"flight_sale_order,omitempty" xml:"flight_sale_order,omitempty" type:"Struct"`
}

func (s IFlightOrderDetailQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModule) GetFlightModifyOrderList() []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	return s.FlightModifyOrderList
}

func (s *IFlightOrderDetailQueryResponseBodyModule) GetFlightRefundOrderList() []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	return s.FlightRefundOrderList
}

func (s *IFlightOrderDetailQueryResponseBodyModule) GetFlightSaleOrder() *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	return s.FlightSaleOrder
}

func (s *IFlightOrderDetailQueryResponseBodyModule) SetFlightModifyOrderList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) *IFlightOrderDetailQueryResponseBodyModule {
	s.FlightModifyOrderList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModule) SetFlightRefundOrderList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) *IFlightOrderDetailQueryResponseBodyModule {
	s.FlightRefundOrderList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModule) SetFlightSaleOrder(v *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) *IFlightOrderDetailQueryResponseBodyModule {
	s.FlightSaleOrder = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList struct {
	CorpPayPrice                *int32                                                                                       `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	FlightModifySegmentList     []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList     `json:"flight_modify_segment_list,omitempty" xml:"flight_modify_segment_list,omitempty" type:"Repeated"`
	FlightOrderModifyTicketList []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList `json:"flight_order_modify_ticket_list,omitempty" xml:"flight_order_modify_ticket_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1005200138736000
	ModifyApplyId *int64                                                                     `json:"modify_apply_id,omitempty" xml:"modify_apply_id,omitempty"`
	ModifyFee     []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee `json:"modify_fee,omitempty" xml:"modify_fee,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ModifyOrderStatus *int32    `json:"modify_order_status,omitempty" xml:"modify_order_status,omitempty"`
	ModifyTotalFee    *int32    `json:"modify_total_fee,omitempty" xml:"modify_total_fee,omitempty"`
	PassengerList     []*string `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PersonPayPrice    *int32    `json:"person_pay_price,omitempty" xml:"person_pay_price,omitempty"`
	// example:
	//
	// 1005200138736028
	RelateModifyApplyId *int64 `json:"relate_modify_apply_id,omitempty" xml:"relate_modify_apply_id,omitempty"`
	ServiceFee          *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// 2024-10-26 11:25:00
	SubmitModifyTime *string `json:"submit_modify_time,omitempty" xml:"submit_modify_time,omitempty"`
	Voluntary        *bool   `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetCorpPayPrice() *int32 {
	return s.CorpPayPrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetFlightModifySegmentList() []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	return s.FlightModifySegmentList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetFlightOrderModifyTicketList() []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	return s.FlightOrderModifyTicketList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetModifyApplyId() *int64 {
	return s.ModifyApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetModifyFee() []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee {
	return s.ModifyFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetModifyOrderStatus() *int32 {
	return s.ModifyOrderStatus
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetModifyTotalFee() *int32 {
	return s.ModifyTotalFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetPassengerList() []*string {
	return s.PassengerList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetPersonPayPrice() *int32 {
	return s.PersonPayPrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetRelateModifyApplyId() *int64 {
	return s.RelateModifyApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetSubmitModifyTime() *string {
	return s.SubmitModifyTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetCorpPayPrice(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.CorpPayPrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetFlightModifySegmentList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.FlightModifySegmentList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetFlightOrderModifyTicketList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.FlightOrderModifyTicketList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetModifyApplyId(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.ModifyApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetModifyFee(v []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.ModifyFee = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetModifyOrderStatus(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.ModifyOrderStatus = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetModifyTotalFee(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.ModifyTotalFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetPassengerList(v []*string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.PassengerList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetPersonPayPrice(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.PersonPayPrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetRelateModifyApplyId(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.RelateModifyApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetServiceFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.ServiceFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetSubmitModifyTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.SubmitModifyTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) SetVoluntary(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList {
	s.Voluntary = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetShare() *bool {
	return s.Share
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopAptCode() *string {
	return s.StopAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetAirlineCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.AirlineCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetAirlineName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.AirlineName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrApt(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrApt = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrTerminal(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetArrTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.ArrTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetCarrierAirlineCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.CarrierAirlineCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetCarrierAirlineName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.CarrierAirlineName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepApt(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepApt = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepTerminal(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepTerminal = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetDepTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.DepTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetJourneyIndex(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetSegmentIndex(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetShare(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.Share = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopArrTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopArrTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) SetStopDepTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList {
	s.StopDepTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightModifySegmentList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList struct {
	CabinClass []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	// example:
	//
	// MU5236
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 781-6605285563
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 0132
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetCabinClass() []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	return s.CabinClass
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetCabinClass(v []*IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.CabinClass = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetTicketNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.TicketNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GetCabin() *string {
	return s.Cabin
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) SetCabin(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	s.Cabin = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) SetCabinClass(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	s.CabinClass = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListFlightOrderModifyTicketListCabinClass) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee struct {
	ModifyHandFee    *int64 `json:"modify_hand_fee,omitempty" xml:"modify_hand_fee,omitempty"`
	ModifyUpgradeFee *int64 `json:"modify_upgrade_fee,omitempty" xml:"modify_upgrade_fee,omitempty"`
	TaxGap           *int64 `json:"tax_gap,omitempty" xml:"tax_gap,omitempty"`
	// example:
	//
	// 0132
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) GetModifyHandFee() *int64 {
	return s.ModifyHandFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) GetModifyUpgradeFee() *int64 {
	return s.ModifyUpgradeFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) GetTaxGap() *int64 {
	return s.TaxGap
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) SetModifyHandFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee {
	s.ModifyHandFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) SetModifyUpgradeFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee {
	s.ModifyUpgradeFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) SetTaxGap(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee {
	s.TaxGap = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightModifyOrderListModifyFee) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList struct {
	CorpRefundAmount            *int32                                                                                       `json:"corp_refund_Amount,omitempty" xml:"corp_refund_Amount,omitempty"`
	FlightOrderRefundTicketList []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList `json:"flight_order_refund_ticket_list,omitempty" xml:"flight_order_refund_ticket_list,omitempty" type:"Repeated"`
	FlightPassengerFee          []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee          `json:"flight_passenger_fee,omitempty" xml:"flight_passenger_fee,omitempty" type:"Repeated"`
	FlightRefundSegmentList     []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList     `json:"flight_refund_segment_list,omitempty" xml:"flight_refund_segment_list,omitempty" type:"Repeated"`
	PassengerList               []*string                                                                                    `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PersonRefundAmount          *int32                                                                                       `json:"person_refund_Amount,omitempty" xml:"person_refund_Amount,omitempty"`
	// example:
	//
	// 1006200138737069
	RefundApplyId *int64 `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	RefundHandFee *int32 `json:"refund_hand_fee,omitempty" xml:"refund_hand_fee,omitempty"`
	// example:
	//
	// 3
	RefundOrderStatus   *int32   `json:"refund_order_status,omitempty" xml:"refund_order_status,omitempty"`
	RefundServiceFee    *int64   `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
	RefundTotalAmount   *int32   `json:"refund_total_Amount,omitempty" xml:"refund_total_Amount,omitempty"`
	RelateModifyApplyId []*int64 `json:"relate_modify_apply_id,omitempty" xml:"relate_modify_apply_id,omitempty" type:"Repeated"`
	// example:
	//
	// 10062001387370
	RelateRefundApplyId *int64 `json:"relate_refund_apply_id,omitempty" xml:"relate_refund_apply_id,omitempty"`
	RepeatRefund        *bool  `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	// example:
	//
	// 2024-10-26 11:25:00
	SubmitRefundTime *string `json:"submit_refund_time,omitempty" xml:"submit_refund_time,omitempty"`
	Voluntary        *bool   `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetCorpRefundAmount() *int32 {
	return s.CorpRefundAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetFlightOrderRefundTicketList() []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	return s.FlightOrderRefundTicketList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetFlightPassengerFee() []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	return s.FlightPassengerFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetFlightRefundSegmentList() []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	return s.FlightRefundSegmentList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetPassengerList() []*string {
	return s.PassengerList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetPersonRefundAmount() *int32 {
	return s.PersonRefundAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRefundApplyId() *int64 {
	return s.RefundApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRefundHandFee() *int32 {
	return s.RefundHandFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRefundOrderStatus() *int32 {
	return s.RefundOrderStatus
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRefundServiceFee() *int64 {
	return s.RefundServiceFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRefundTotalAmount() *int32 {
	return s.RefundTotalAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRelateModifyApplyId() []*int64 {
	return s.RelateModifyApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRelateRefundApplyId() *int64 {
	return s.RelateRefundApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetRepeatRefund() *bool {
	return s.RepeatRefund
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetSubmitRefundTime() *string {
	return s.SubmitRefundTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetCorpRefundAmount(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.CorpRefundAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetFlightOrderRefundTicketList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.FlightOrderRefundTicketList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetFlightPassengerFee(v []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.FlightPassengerFee = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetFlightRefundSegmentList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.FlightRefundSegmentList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetPassengerList(v []*string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.PassengerList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetPersonRefundAmount(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.PersonRefundAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRefundApplyId(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRefundHandFee(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundHandFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRefundOrderStatus(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundOrderStatus = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRefundServiceFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundServiceFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRefundTotalAmount(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RefundTotalAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRelateModifyApplyId(v []*int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RelateModifyApplyId = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRelateRefundApplyId(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RelateRefundApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetRepeatRefund(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.RepeatRefund = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetSubmitRefundTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.SubmitRefundTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) SetVoluntary(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList {
	s.Voluntary = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList struct {
	CabinClass []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	// example:
	//
	// MU5236
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 018-6605785754
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 0132
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetCabinClass() []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	return s.CabinClass
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetCabinClass(v []*IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.CabinClass = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetTicketNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.TicketNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GetCabin() *string {
	return s.Cabin
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) SetCabin(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	s.Cabin = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) SetCabinClass(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	s.CabinClass = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightOrderRefundTicketListCabinClass) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee struct {
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
	// 0132
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetNoRefundModifyHandFee() *int64 {
	return s.NoRefundModifyHandFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetNoRefundModifyUpgradeFee() *int64 {
	return s.NoRefundModifyUpgradeFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetRefundAmount() *int64 {
	return s.RefundAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetRefundHandFee() *int64 {
	return s.RefundHandFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetRefundModifyAmount() *int64 {
	return s.RefundModifyAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetRefundModifyHandAmount() *int64 {
	return s.RefundModifyHandAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetRefundModifyUpgradeAmount() *int64 {
	return s.RefundModifyUpgradeAmount
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetRefundTaxHandFee() *int64 {
	return s.RefundTaxHandFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetNoRefundModifyHandFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.NoRefundModifyHandFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetNoRefundModifyUpgradeFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.NoRefundModifyUpgradeFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetRefundAmount(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.RefundAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetRefundHandFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.RefundHandFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetRefundModifyAmount(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.RefundModifyAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetRefundModifyHandAmount(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.RefundModifyHandAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetRefundModifyUpgradeAmount(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.RefundModifyUpgradeAmount = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetRefundTaxHandFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.RefundTaxHandFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightPassengerFee) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetShare() *bool {
	return s.Share
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopAptCode() *string {
	return s.StopAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetAirlineCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetAirlineName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.AirlineName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrApt(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrApt = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrTerminal(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetArrTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.ArrTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetCarrierAirlineCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.CarrierAirlineCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetCarrierAirlineName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.CarrierAirlineName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepApt(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepApt = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepTerminal(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetDepTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.DepTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetJourneyIndex(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetSegmentIndex(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetShare(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.Share = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopArrTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopArrTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) SetStopDepTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList {
	s.StopDepTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightRefundOrderListFlightRefundSegmentList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder struct {
	// example:
	//
	// 82587500
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 0
	BookType     *int32                                                              `json:"book_type,omitempty" xml:"book_type,omitempty"`
	BookerInfo   *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo `json:"booker_info,omitempty" xml:"booker_info,omitempty" type:"Struct"`
	CorpPayPrice *int64                                                              `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// example:
	//
	// 123
	ExceedApplyId         *string                                                                          `json:"exceed_apply_id,omitempty" xml:"exceed_apply_id,omitempty"`
	FlightOrderInsureList []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList `json:"flight_order_insure_list,omitempty" xml:"flight_order_insure_list,omitempty" type:"Repeated"`
	FlightOrderTicketList []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList `json:"flight_order_ticket_list,omitempty" xml:"flight_order_ticket_list,omitempty" type:"Repeated"`
	FlightSegmentList     []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList     `json:"flight_segment_list,omitempty" xml:"flight_segment_list,omitempty" type:"Repeated"`
	MixPay                *bool                                                                            `json:"mix_pay,omitempty" xml:"mix_pay,omitempty"`
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
	OrderType     *int32                                                                   `json:"order_type,omitempty" xml:"order_type,omitempty"`
	PassengerList []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetBookType() *int32 {
	return s.BookType
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetBookerInfo() *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo {
	return s.BookerInfo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetExceedApplyId() *string {
	return s.ExceedApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetFlightOrderInsureList() []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	return s.FlightOrderInsureList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetFlightOrderTicketList() []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	return s.FlightOrderTicketList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetFlightSegmentList() []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	return s.FlightSegmentList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetMixPay() *bool {
	return s.MixPay
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderCreateTime() *string {
	return s.OrderCreateTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderPayTime() *string {
	return s.OrderPayTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderReservePrice() *int64 {
	return s.OrderReservePrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetOrderType() *int32 {
	return s.OrderType
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetPassengerList() []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	return s.PassengerList
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetPayType() *int32 {
	return s.PayType
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetPersonPayPrice() *int64 {
	return s.PersonPayPrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) GetTripType() *int32 {
	return s.TripType
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetApplyId(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.ApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetBookType(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.BookType = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetBookerInfo(v *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.BookerInfo = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetCorpPayPrice(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.CorpPayPrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetExceedApplyId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.ExceedApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetFlightOrderInsureList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.FlightOrderInsureList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetFlightOrderTicketList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.FlightOrderTicketList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetFlightSegmentList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.FlightSegmentList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetMixPay(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.MixPay = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderCreateTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderCreateTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderPayTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderPayTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderReservePrice(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderReservePrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderStatus(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderStatus = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderStatusDesc(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderStatusDesc = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetOrderType(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.OrderType = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetPassengerList(v []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.PassengerList = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetPayType(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.PayType = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetPersonPayPrice(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.PersonPayPrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetServiceFee(v int64) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.ServiceFee = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetThirdPartApplyId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.ThirdPartApplyId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) SetTripType(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder {
	s.TripType = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrder) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo struct {
	// example:
	//
	// 01323
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) GetUserName() *string {
	return s.UserName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) SetUserName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo {
	s.UserName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderBookerInfo) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetInsPayType() *string {
	return s.InsPayType
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetInsTotalPrice() *int32 {
	return s.InsTotalPrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) GetTradeAction() *string {
	return s.TradeAction
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetInsOrderId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.InsOrderId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetInsPayType(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.InsPayType = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetInsTotalPrice(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.InsTotalPrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) SetTradeAction(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList {
	s.TradeAction = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderInsureList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList struct {
	CabinClass []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	Tax        *int32                                                                                     `json:"tax,omitempty" xml:"tax,omitempty"`
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetCabinClass() []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	return s.CabinClass
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetTax() *int32 {
	return s.Tax
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetCabinClass(v []*IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.CabinClass = v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetTax(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.Tax = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetTicketNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.TicketNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetTicketPrice(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.TicketPrice = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GetCabin() *string {
	return s.Cabin
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) SetCabin(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	s.Cabin = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) SetCabinClass(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	s.CabinClass = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightOrderTicketListCabinClass) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetShare() *bool {
	return s.Share
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopAptCode() *string {
	return s.StopAptCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetAirlineCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetAirlineName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.AirlineName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrApt(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrApt = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrTerminal(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetArrTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.ArrTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetCarrierAirlineCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.CarrierAirlineCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetCarrierAirlineName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.CarrierAirlineName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepApt(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepApt = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepTerminal(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetDepTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.DepTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetFlightNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetJourneyIndex(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetSegmentIndex(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetShare(v bool) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.Share = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopAptCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopAptCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopArrTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopArrTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopCity(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopCity = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopCityCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopCityCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) SetStopDepTime(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList {
	s.StopDepTime = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderFlightSegmentList) Validate() error {
	return dara.Validate(s)
}

type IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList struct {
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

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GoString() string {
	return s.String()
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetCostCenterId() *string {
	return s.CostCenterId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetInvoiceId() *string {
	return s.InvoiceId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) GetUserName() *string {
	return s.UserName
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetCostCenterId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.CostCenterId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetCostCenterName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.CostCenterName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetDepartmentId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.DepartmentId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetDepartmentName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.DepartmentName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetInvoiceId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.InvoiceId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetInvoiceTitle(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.InvoiceTitle = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetJobNo(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.JobNo = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetPassengerType(v int32) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.PassengerType = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetProjectCode(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.ProjectCode = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetProjectTitle(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.ProjectTitle = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetUserId(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.UserId = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) SetUserName(v string) *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList {
	s.UserName = &v
	return s
}

func (s *IFlightOrderDetailQueryResponseBodyModuleFlightSaleOrderPassengerList) Validate() error {
	return dara.Validate(s)
}
