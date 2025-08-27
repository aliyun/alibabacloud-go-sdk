// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightOrderDetailResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightOrderDetailResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightOrderDetailResponseBodyModule) *IntlFlightOrderDetailResponseBody
	GetModule() *IntlFlightOrderDetailResponseBodyModule
	SetRequestId(v string) *IntlFlightOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightOrderDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightOrderDetailResponseBody
	GetTraceId() *string
}

type IntlFlightOrderDetailResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    *IntlFlightOrderDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOrderDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightOrderDetailResponseBody) GetModule() *IntlFlightOrderDetailResponseBodyModule {
	return s.Module
}

func (s *IntlFlightOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightOrderDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightOrderDetailResponseBody) SetCode(v string) *IntlFlightOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBody) SetMessage(v string) *IntlFlightOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBody) SetModule(v *IntlFlightOrderDetailResponseBodyModule) *IntlFlightOrderDetailResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightOrderDetailResponseBody) SetRequestId(v string) *IntlFlightOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBody) SetSuccess(v bool) *IntlFlightOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBody) SetTraceId(v string) *IntlFlightOrderDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModule struct {
	BookerUserId   *string `json:"booker_user_id,omitempty" xml:"booker_user_id,omitempty"`
	BookerUserName *string `json:"booker_user_name,omitempty" xml:"booker_user_name,omitempty"`
	CloseReason    *string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// example:
	//
	// 1709708165000
	CloseTime   *string                                               `json:"close_time,omitempty" xml:"close_time,omitempty"`
	ContactInfo *IntlFlightOrderDetailResponseBodyModuleContactInfo   `json:"contact_info,omitempty" xml:"contact_info,omitempty" type:"Struct"`
	JourneyList []*IntlFlightOrderDetailResponseBodyModuleJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1012000000000000
	OrderId       *string                                                 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderItemList []*IntlFlightOrderDetailResponseBodyModuleOrderItemList `json:"order_item_list,omitempty" xml:"order_item_list,omitempty" type:"Repeated"`
	OrderStatus   *int32                                                  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// F11374007131319304192
	OutOrderId          *string                                                       `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerList       []*IntlFlightOrderDetailResponseBodyModulePassengerList       `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PassengerTicketList []*IntlFlightOrderDetailResponseBodyModulePassengerTicketList `json:"passenger_ticket_list,omitempty" xml:"passenger_ticket_list,omitempty" type:"Repeated"`
	PayLatestTime       *string                                                       `json:"pay_latest_time,omitempty" xml:"pay_latest_time,omitempty"`
	// example:
	//
	// 1
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 2024-03-06 15:00:35
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 8
	PayType     *int32  `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	SuccessTime *string `json:"success_time,omitempty" xml:"success_time,omitempty"`
	// example:
	//
	// 21300
	TotalPrice       *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	TotalTaxPrice    *int64 `json:"total_tax_price,omitempty" xml:"total_tax_price,omitempty"`
	TotalTicketPrice *int64 `json:"total_ticket_price,omitempty" xml:"total_ticket_price,omitempty"`
	// example:
	//
	// 0
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetBookerUserId() *string {
	return s.BookerUserId
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetBookerUserName() *string {
	return s.BookerUserName
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetCloseReason() *string {
	return s.CloseReason
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetCloseTime() *string {
	return s.CloseTime
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetContactInfo() *IntlFlightOrderDetailResponseBodyModuleContactInfo {
	return s.ContactInfo
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetJourneyList() []*IntlFlightOrderDetailResponseBodyModuleJourneyList {
	return s.JourneyList
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetOrderItemList() []*IntlFlightOrderDetailResponseBodyModuleOrderItemList {
	return s.OrderItemList
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetPassengerList() []*IntlFlightOrderDetailResponseBodyModulePassengerList {
	return s.PassengerList
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetPassengerTicketList() []*IntlFlightOrderDetailResponseBodyModulePassengerTicketList {
	return s.PassengerTicketList
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetPayLatestTime() *string {
	return s.PayLatestTime
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetPayType() *int32 {
	return s.PayType
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetSuccessTime() *string {
	return s.SuccessTime
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetTotalTaxPrice() *int64 {
	return s.TotalTaxPrice
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetTotalTicketPrice() *int64 {
	return s.TotalTicketPrice
}

func (s *IntlFlightOrderDetailResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetBookerUserId(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.BookerUserId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetBookerUserName(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.BookerUserName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetCloseReason(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.CloseReason = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetCloseTime(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.CloseTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetContactInfo(v *IntlFlightOrderDetailResponseBodyModuleContactInfo) *IntlFlightOrderDetailResponseBodyModule {
	s.ContactInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetJourneyList(v []*IntlFlightOrderDetailResponseBodyModuleJourneyList) *IntlFlightOrderDetailResponseBodyModule {
	s.JourneyList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetOrderId(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetOrderItemList(v []*IntlFlightOrderDetailResponseBodyModuleOrderItemList) *IntlFlightOrderDetailResponseBodyModule {
	s.OrderItemList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetOrderStatus(v int32) *IntlFlightOrderDetailResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetOutOrderId(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetPassengerList(v []*IntlFlightOrderDetailResponseBodyModulePassengerList) *IntlFlightOrderDetailResponseBodyModule {
	s.PassengerList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetPassengerTicketList(v []*IntlFlightOrderDetailResponseBodyModulePassengerTicketList) *IntlFlightOrderDetailResponseBodyModule {
	s.PassengerTicketList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetPayLatestTime(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.PayLatestTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetPayStatus(v int32) *IntlFlightOrderDetailResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetPayTime(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetPayType(v int32) *IntlFlightOrderDetailResponseBodyModule {
	s.PayType = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetSuccessTime(v string) *IntlFlightOrderDetailResponseBodyModule {
	s.SuccessTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetTotalPrice(v int64) *IntlFlightOrderDetailResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetTotalTaxPrice(v int64) *IntlFlightOrderDetailResponseBodyModule {
	s.TotalTaxPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetTotalTicketPrice(v int64) *IntlFlightOrderDetailResponseBodyModule {
	s.TotalTicketPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) SetTripType(v int32) *IntlFlightOrderDetailResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleContactInfo struct {
	ContactName *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleContactInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleContactInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleContactInfo) GetContactName() *string {
	return s.ContactName
}

func (s *IntlFlightOrderDetailResponseBodyModuleContactInfo) SetContactName(v string) *IntlFlightOrderDetailResponseBodyModuleContactInfo {
	s.ContactName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleContactInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyList struct {
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime     *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// SHA
	DepCityCode        *string                                                                 `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName        *string                                                                 `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime            *string                                                                 `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Duration           *int32                                                                  `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetFlightSegmentInfos() []*IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetArrCityCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetArrCityName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetArrTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetDepCityCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetDepCityName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetDepTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.DepTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetDuration(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.Duration = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetFlightSegmentInfos(v []*IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetJourneyIndex(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) SetTransferTime(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyList {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos struct {
	AirlineInfo        *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo          `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo     *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo       `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	ArrCityCode        *string                                                                                   `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName        *string                                                                                   `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime            *string                                                                                   `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportInfo     *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo       `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	DepCityCode        *string                                                                                   `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName        *string                                                                                   `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime            *string                                                                                   `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Duration           *int32                                                                                    `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightNo           *string                                                                                   `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                                   `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	FlightType         *string                                                                                   `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	LuggageDirectInfo  *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo    `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer       *string                                                                                   `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc           *string                                                                                   `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	OneMore            *int32                                                                                    `json:"one_more,omitempty" xml:"one_more,omitempty"`
	OneMoreShow        *string                                                                                   `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	SegmentIndex       *int32                                                                                    `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	SegmentKey         *string                                                                                   `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark  *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark    `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	Share              *bool                                                                                     `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize    *string                                                                                   `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	Stop               *bool                                                                                     `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime          *string                                                                                   `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetAirlineInfo() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrAirportInfo() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepAirportInfo() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightShareInfo() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetAirlineInfo(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrCityName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepCityName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDuration(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightNo(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightSize(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightType(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetManufacturer(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetMealDesc(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetOneMore(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetShare(v bool) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetStop(v bool) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetTotalTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo struct {
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo struct {
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	Terminal         *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo struct {
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	Terminal         *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	OperatingFlightNo    *string                                                                                                  `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList struct {
	StopAirport     *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	StopArrTerm     *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	StopArrTime     *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCityCode    *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName    *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopDepTerm     *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	StopDepTime     *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	StopTime        *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo struct {
	DepCityLuggageDirect  *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark struct {
	DepCityVisaRemark   *string   `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleOrderItemList struct {
	BaggageRule        *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule          `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty" type:"Struct"`
	PassengerPriceList []*IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList `json:"passenger_price_list,omitempty" xml:"passenger_price_list,omitempty" type:"Repeated"`
	RefundChangeRule   *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule     `json:"refund_change_rule,omitempty" xml:"refund_change_rule,omitempty" type:"Struct"`
	SegmentKeyList     []*string                                                                 `json:"segment_key_list,omitempty" xml:"segment_key_list,omitempty" type:"Repeated"`
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) GetBaggageRule() *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule {
	return s.BaggageRule
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) GetPassengerPriceList() []*IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList {
	return s.PassengerPriceList
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) GetRefundChangeRule() *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	return s.RefundChangeRule
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) GetSegmentKeyList() []*string {
	return s.SegmentKeyList
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) SetBaggageRule(v *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) *IntlFlightOrderDetailResponseBodyModuleOrderItemList {
	s.BaggageRule = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) SetPassengerPriceList(v []*IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) *IntlFlightOrderDetailResponseBodyModuleOrderItemList {
	s.PassengerPriceList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) SetRefundChangeRule(v *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) *IntlFlightOrderDetailResponseBodyModuleOrderItemList {
	s.RefundChangeRule = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) SetSegmentKeyList(v []*string) *IntlFlightOrderDetailResponseBodyModuleOrderItemList {
	s.SegmentKeyList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule struct {
	BaggageDigest       *string                                                              `json:"baggage_digest,omitempty" xml:"baggage_digest,omitempty"`
	OfferBaggageInfoMap map[string][]*ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue `json:"offer_baggage_info_map,omitempty" xml:"offer_baggage_info_map,omitempty"`
	// example:
	//
	// true
	StructuredBaggage *bool `json:"structured_baggage,omitempty" xml:"structured_baggage,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) GetBaggageDigest() *string {
	return s.BaggageDigest
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) GetOfferBaggageInfoMap() map[string][]*ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue {
	return s.OfferBaggageInfoMap
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) GetStructuredBaggage() *bool {
	return s.StructuredBaggage
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) SetBaggageDigest(v string) *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule {
	s.BaggageDigest = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) SetOfferBaggageInfoMap(v map[string][]*ModuleOrderItemListBaggageRuleOfferBaggageInfoMapValue) *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule {
	s.OfferBaggageInfoMap = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) SetStructuredBaggage(v bool) *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule {
	s.StructuredBaggage = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListBaggageRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList struct {
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// example:
	//
	// 21300
	SellPrice *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// example:
	//
	// 19300
	Tax *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
	// example:
	//
	// 2000
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) GetTax() *int32 {
	return s.Tax
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) SetPassengerType(v int32) *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList {
	s.PassengerType = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) SetSellPrice(v int32) *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList {
	s.SellPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) SetTax(v int32) *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList {
	s.Tax = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) SetTicketPrice(v int32) *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList {
	s.TicketPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListPassengerPriceList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule struct {
	// example:
	//
	// false
	CancelFeeInd *bool `json:"cancel_fee_ind,omitempty" xml:"cancel_fee_ind,omitempty"`
	// example:
	//
	// false
	ChangeFeeInd        *bool                                                                     `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	ChangeRuleDesc      *string                                                                   `json:"change_rule_desc,omitempty" xml:"change_rule_desc,omitempty"`
	OfferPenaltyInfoMap map[string][]*ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue `json:"offer_penalty_info_map,omitempty" xml:"offer_penalty_info_map,omitempty"`
	RefundChangeDigest  *string                                                                   `json:"refund_change_digest,omitempty" xml:"refund_change_digest,omitempty"`
	RefundRuleDesc      *string                                                                   `json:"refund_rule_desc,omitempty" xml:"refund_rule_desc,omitempty"`
	// example:
	//
	// true
	StructuredRefund *bool `json:"structured_refund,omitempty" xml:"structured_refund,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetChangeRuleDesc() *string {
	return s.ChangeRuleDesc
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetOfferPenaltyInfoMap() map[string][]*ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue {
	return s.OfferPenaltyInfoMap
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetRefundChangeDigest() *string {
	return s.RefundChangeDigest
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetRefundRuleDesc() *string {
	return s.RefundRuleDesc
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) GetStructuredRefund() *bool {
	return s.StructuredRefund
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetCancelFeeInd(v bool) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.CancelFeeInd = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetChangeFeeInd(v bool) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.ChangeFeeInd = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetChangeRuleDesc(v string) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.ChangeRuleDesc = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetOfferPenaltyInfoMap(v map[string][]*ModuleOrderItemListRefundChangeRuleOfferPenaltyInfoMapValue) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.OfferPenaltyInfoMap = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetRefundChangeDigest(v string) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.RefundChangeDigest = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetRefundRuleDesc(v string) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.RefundRuleDesc = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) SetStructuredRefund(v bool) *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule {
	s.StructuredRefund = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModuleOrderItemListRefundChangeRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModulePassengerList struct {
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1001101
	JobNo       *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// example:
	//
	// CN
	NationalityCode *string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// example:
	//
	// 8432002
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 12292812036903456
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetFullName(v string) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetGender(v int32) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.Gender = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetJobNo(v string) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.JobNo = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetNationality(v string) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.Nationality = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetNationalityCode(v string) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.NationalityCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetPassengerId(v int64) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetType(v int32) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetUserId(v string) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.UserId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) SetUserType(v int32) *IntlFlightOrderDetailResponseBodyModulePassengerList {
	s.UserType = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModulePassengerTicketList struct {
	PassengerId *int64                                                                  `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	TicketList  []*IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList `json:"ticket_list,omitempty" xml:"ticket_list,omitempty" type:"Repeated"`
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketList) GetTicketList() []*IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	return s.TicketList
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketList) SetPassengerId(v int64) *IntlFlightOrderDetailResponseBodyModulePassengerTicketList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketList) SetTicketList(v []*IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) *IntlFlightOrderDetailResponseBodyModulePassengerTicketList {
	s.TicketList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList struct {
	IssueTime         *string                                                                                  `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	PnrNo             *string                                                                                  `json:"pnr_no,omitempty" xml:"pnr_no,omitempty"`
	PriceInfo         *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo           `json:"price_info,omitempty" xml:"price_info,omitempty" type:"Struct"`
	SegmentKeyList    []*string                                                                                `json:"segment_key_list,omitempty" xml:"segment_key_list,omitempty" type:"Repeated"`
	TicketNo          *string                                                                                  `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketSegmentList []*IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList `json:"ticket_segment_list,omitempty" xml:"ticket_segment_list,omitempty" type:"Repeated"`
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GetIssueTime() *string {
	return s.IssueTime
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GetPnrNo() *string {
	return s.PnrNo
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GetPriceInfo() *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo {
	return s.PriceInfo
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GetSegmentKeyList() []*string {
	return s.SegmentKeyList
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) GetTicketSegmentList() []*IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	return s.TicketSegmentList
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) SetIssueTime(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	s.IssueTime = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) SetPnrNo(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	s.PnrNo = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) SetPriceInfo(v *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	s.PriceInfo = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) SetSegmentKeyList(v []*string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	s.SegmentKeyList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) SetTicketNo(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	s.TicketNo = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) SetTicketSegmentList(v []*IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList {
	s.TicketSegmentList = v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo struct {
	SellPrice   *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	Tax         *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) GetTax() *int32 {
	return s.Tax
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) SetSellPrice(v int32) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo {
	s.SellPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) SetTax(v int32) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo {
	s.Tax = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) SetTicketPrice(v int32) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo {
	s.TicketPrice = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListPriceInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList struct {
	Cabin            *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass       *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	Modified         *bool   `json:"modified,omitempty" xml:"modified,omitempty"`
	OpenTicketStatus *string `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
	Refunded         *bool   `json:"refunded,omitempty" xml:"refunded,omitempty"`
	SegmentKey       *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetModified() *bool {
	return s.Modified
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetOpenTicketStatus() *string {
	return s.OpenTicketStatus
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetRefunded() *bool {
	return s.Refunded
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetCabin(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.Cabin = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetCabinClass(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.CabinClass = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetModified(v bool) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.Modified = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetOpenTicketStatus(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.OpenTicketStatus = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetRefunded(v bool) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.Refunded = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetSegmentKey(v string) *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightOrderDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) Validate() error {
	return dara.Validate(s)
}
