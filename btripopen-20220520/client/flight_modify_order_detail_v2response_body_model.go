// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOrderDetailV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightModifyOrderDetailV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightModifyOrderDetailV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightModifyOrderDetailV2ResponseBodyModule) *FlightModifyOrderDetailV2ResponseBody
	GetModule() *FlightModifyOrderDetailV2ResponseBodyModule
	SetRequestId(v string) *FlightModifyOrderDetailV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightModifyOrderDetailV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightModifyOrderDetailV2ResponseBody
	GetTraceId() *string
}

type FlightModifyOrderDetailV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightModifyOrderDetailV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210bc2dc16839612026565712dcbe6
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightModifyOrderDetailV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightModifyOrderDetailV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightModifyOrderDetailV2ResponseBody) GetModule() *FlightModifyOrderDetailV2ResponseBodyModule {
	return s.Module
}

func (s *FlightModifyOrderDetailV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightModifyOrderDetailV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightModifyOrderDetailV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightModifyOrderDetailV2ResponseBody) SetCode(v string) *FlightModifyOrderDetailV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBody) SetMessage(v string) *FlightModifyOrderDetailV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBody) SetModule(v *FlightModifyOrderDetailV2ResponseBodyModule) *FlightModifyOrderDetailV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBody) SetRequestId(v string) *FlightModifyOrderDetailV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBody) SetSuccess(v bool) *FlightModifyOrderDetailV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBody) SetTraceId(v string) *FlightModifyOrderDetailV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModule struct {
	// example:
	//
	// 2023-08-14 11:28:01
	ApplyTime  *string                                                `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	Attributes *FlightModifyOrderDetailV2ResponseBodyModuleAttributes `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Struct"`
	// example:
	//
	// 17635462345@163.com
	BookUserEmail *string `json:"book_user_email,omitempty" xml:"book_user_email,omitempty"`
	BookUserName  *string `json:"book_user_name,omitempty" xml:"book_user_name,omitempty"`
	// example:
	//
	// 17635462345
	BookuserPhone      *string                                                          `json:"bookuser_phone,omitempty" xml:"bookuser_phone,omitempty"`
	ChangeFailReason   *string                                                          `json:"change_fail_reason,omitempty" xml:"change_fail_reason,omitempty"`
	ContactInfoDTO     *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO       `json:"contact_info_d_t_o,omitempty" xml:"contact_info_d_t_o,omitempty" type:"Struct"`
	DestFlightInfoDTOS []*FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS `json:"dest_flight_info_d_t_o_s,omitempty" xml:"dest_flight_info_d_t_o_s,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-08-14 11:48:01
	LastPayTime *string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// example:
	//
	// 1017124195788186048
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017124195788186048
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195836916029
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// example:
	//
	// 2023-08-14 11:38:01
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	Reason  *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1019195836916029
	SubOrderId *int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// 2023-08-19 15:24:08
	TicketTime *string `json:"ticket_time,omitempty" xml:"ticket_time,omitempty"`
	// example:
	//
	// 100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// example:
	//
	// 0
	TotalServiceFeePrice *int64                                                         `json:"total_service_fee_price,omitempty" xml:"total_service_fee_price,omitempty"`
	TravelerInfoDTOS     []*FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS `json:"traveler_info_d_t_o_s,omitempty" xml:"traveler_info_d_t_o_s,omitempty" type:"Repeated"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetAttributes() *FlightModifyOrderDetailV2ResponseBodyModuleAttributes {
	return s.Attributes
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetBookUserEmail() *string {
	return s.BookUserEmail
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetBookUserName() *string {
	return s.BookUserName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetBookuserPhone() *string {
	return s.BookuserPhone
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetChangeFailReason() *string {
	return s.ChangeFailReason
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetContactInfoDTO() *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO {
	return s.ContactInfoDTO
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetDestFlightInfoDTOS() []*FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	return s.DestFlightInfoDTOS
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetLastPayTime() *string {
	return s.LastPayTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetReason() *string {
	return s.Reason
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetTicketTime() *string {
	return s.TicketTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetTotalServiceFeePrice() *int64 {
	return s.TotalServiceFeePrice
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) GetTravelerInfoDTOS() []*FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	return s.TravelerInfoDTOS
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetApplyTime(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.ApplyTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetAttributes(v *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.Attributes = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetBookUserEmail(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.BookUserEmail = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetBookUserName(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.BookUserName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetBookuserPhone(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.BookuserPhone = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetChangeFailReason(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.ChangeFailReason = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetContactInfoDTO(v *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.ContactInfoDTO = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetDestFlightInfoDTOS(v []*FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.DestFlightInfoDTOS = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetLastPayTime(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.LastPayTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetOrderId(v int64) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetOutOrderId(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetOutSubOrderId(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetPayTime(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetReason(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.Reason = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetStatus(v int32) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetSubOrderId(v int64) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.SubOrderId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetTicketTime(v string) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.TicketTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetTotalPrice(v int64) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetTotalServiceFeePrice(v int64) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.TotalServiceFeePrice = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) SetTravelerInfoDTOS(v []*FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) *FlightModifyOrderDetailV2ResponseBodyModule {
	s.TravelerInfoDTOS = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleAttributes struct {
	BaggageRule      *string     `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty"`
	ChangeRule       *string     `json:"change_rule,omitempty" xml:"change_rule,omitempty"`
	LatestPayTime    interface{} `json:"latest_pay_time,omitempty" xml:"latest_pay_time,omitempty"`
	LatestPayTimeStr *string     `json:"latest_pay_time_str,omitempty" xml:"latest_pay_time_str,omitempty"`
	RefundRule       *string     `json:"refund_rule,omitempty" xml:"refund_rule,omitempty"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleAttributes) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleAttributes) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) GetBaggageRule() *string {
	return s.BaggageRule
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) GetChangeRule() *string {
	return s.ChangeRule
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) GetLatestPayTime() interface{} {
	return s.LatestPayTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) GetLatestPayTimeStr() *string {
	return s.LatestPayTimeStr
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) GetRefundRule() *string {
	return s.RefundRule
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) SetBaggageRule(v string) *FlightModifyOrderDetailV2ResponseBodyModuleAttributes {
	s.BaggageRule = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) SetChangeRule(v string) *FlightModifyOrderDetailV2ResponseBodyModuleAttributes {
	s.ChangeRule = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) SetLatestPayTime(v interface{}) *FlightModifyOrderDetailV2ResponseBodyModuleAttributes {
	s.LatestPayTime = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) SetLatestPayTimeStr(v string) *FlightModifyOrderDetailV2ResponseBodyModuleAttributes {
	s.LatestPayTimeStr = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) SetRefundRule(v string) *FlightModifyOrderDetailV2ResponseBodyModuleAttributes {
	s.RefundRule = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleAttributes) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO struct {
	// example:
	//
	// 17816963077@163.com
	ContactEmail *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	ContactName  *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// example:
	//
	// 17816963077
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// false
	SendMsgToPassenger *bool `json:"send_msg_to_passenger,omitempty" xml:"send_msg_to_passenger,omitempty"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) GetContactName() *string {
	return s.ContactName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) GetSendMsgToPassenger() *bool {
	return s.SendMsgToPassenger
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) SetContactEmail(v string) *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactEmail = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) SetContactName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) SetContactPhone(v string) *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactPhone = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) SetSendMsgToPassenger(v bool) *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.SendMsgToPassenger = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleContactInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS struct {
	// example:
	//
	// MU
	AirlineCode    *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineIconUrl *string `json:"airline_icon_url,omitempty" xml:"airline_icon_url,omitempty"`
	AirlineName    *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// example:
	//
	// HGH
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrAirportName *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2023-10-03 09:30:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// 10
	CabinDiscount         *int64  `json:"cabin_discount,omitempty" xml:"cabin_discount,omitempty"`
	CarrierAirlineCode    *string `json:"carrier_airline_code,omitempty" xml:"carrier_airline_code,omitempty"`
	CarrierAirlineIconUrl *string `json:"carrier_airline_icon_url,omitempty" xml:"carrier_airline_icon_url,omitempty"`
	CarrierAirlineName    *string `json:"carrier_airline_name,omitempty" xml:"carrier_airline_name,omitempty"`
	CarrierFlightNo       *string `json:"carrier_flight_no,omitempty" xml:"carrier_flight_no,omitempty"`
	// example:
	//
	// PKX
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepAirportName *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2023-10-03 07:30:00
	DepTime      *string                                                                    `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightChange *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange `json:"flight_change,omitempty" xml:"flight_change,omitempty" type:"Struct"`
	// example:
	//
	// MU5193
	FlightNo   *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	MealDesc   *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 1194012
	SegmentIId      *string                                                                       `json:"segmentI_id,omitempty" xml:"segmentI_id,omitempty"`
	SegmentPosition *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	StopArrTime     *string                                                                       `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCity        *string                                                                       `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	StopDepTime     *string                                                                       `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetAirlineIconUrl() *string {
	return s.AirlineIconUrl
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCabin() *string {
	return s.Cabin
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCabinDiscount() *int64 {
	return s.CabinDiscount
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCarrierAirlineIconUrl() *string {
	return s.CarrierAirlineIconUrl
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetCarrierFlightNo() *string {
	return s.CarrierFlightNo
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetFlightChange() *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange {
	return s.FlightChange
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetSegmentIId() *string {
	return s.SegmentIId
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetSegmentPosition() *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition {
	return s.SegmentPosition
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetAirlineCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetAirlineIconUrl(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.AirlineIconUrl = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetAirlineName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.AirlineName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetArrAirportCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetArrAirportName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.ArrAirportName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetArrCityCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.ArrCityCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetArrCityName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.ArrCityName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetArrTerminal(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.ArrTerminal = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetArrTime(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.ArrTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCabin(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.Cabin = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCabinClass(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CabinClass = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCabinClassName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CabinClassName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCabinDiscount(v int64) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CabinDiscount = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCarrierAirlineCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CarrierAirlineCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCarrierAirlineIconUrl(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CarrierAirlineIconUrl = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCarrierAirlineName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CarrierAirlineName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetCarrierFlightNo(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.CarrierFlightNo = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetDepAirportCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.DepAirportCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetDepAirportName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.DepAirportName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetDepCityCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.DepCityCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetDepCityName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.DepCityName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetDepTerminal(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.DepTerminal = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetDepTime(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.DepTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetFlightChange(v *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.FlightChange = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetFlightNo(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.FlightNo = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetFlightType(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.FlightType = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetMealDesc(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.MealDesc = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetSegmentIId(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.SegmentIId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetSegmentPosition(v *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.SegmentPosition = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetStopArrTime(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.StopArrTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetStopCity(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.StopCity = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) SetStopDepTime(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS {
	s.StopDepTime = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOS) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange struct {
	ChangeDesc       *string     `json:"change_desc,omitempty" xml:"change_desc,omitempty"`
	ChangeStatus     *string     `json:"change_status,omitempty" xml:"change_status,omitempty"`
	ChangeStatusCode *string     `json:"change_status_code,omitempty" xml:"change_status_code,omitempty"`
	NewSegment       interface{} `json:"new_segment,omitempty" xml:"new_segment,omitempty"`
	PassengerNames   []*string   `json:"passenger_names,omitempty" xml:"passenger_names,omitempty" type:"Repeated"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) GetChangeDesc() *string {
	return s.ChangeDesc
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) GetChangeStatus() *string {
	return s.ChangeStatus
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) GetChangeStatusCode() *string {
	return s.ChangeStatusCode
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) GetNewSegment() interface{} {
	return s.NewSegment
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) GetPassengerNames() []*string {
	return s.PassengerNames
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) SetChangeDesc(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange {
	s.ChangeDesc = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) SetChangeStatus(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange {
	s.ChangeStatus = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) SetChangeStatusCode(v string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange {
	s.ChangeStatusCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) SetNewSegment(v interface{}) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange {
	s.NewSegment = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) SetPassengerNames(v []*string) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange {
	s.PassengerNames = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSFlightChange) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) SetJourneyIndex(v int32) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) SetSegmentIndex(v int32) *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleDestFlightInfoDTOSSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS struct {
	// example:
	//
	// 2000-08-19
	BirthDate *string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// example:
	//
	// 430131413423435353
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 0
	CertType  *int32                                                                `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	ChangeFee *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee `json:"change_fee,omitempty" xml:"change_fee,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Gender          *int32    `json:"gender,omitempty" xml:"gender,omitempty"`
	OriginTicketNos []*string `json:"origin_ticket_nos,omitempty" xml:"origin_ticket_nos,omitempty" type:"Repeated"`
	// example:
	//
	// 12172819047252004460056
	PassengerId   *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// example:
	//
	// 17635462345
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 3243028
	Pid                *int64                 `json:"pid,omitempty" xml:"pid,omitempty"`
	TicketNoSegmentMap map[string]interface{} `json:"ticket_no_segment_map,omitempty" xml:"ticket_no_segment_map,omitempty"`
	TicketNos          []*string              `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetBirthDate() *string {
	return s.BirthDate
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetCertNo() *string {
	return s.CertNo
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetCertType() *int32 {
	return s.CertType
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetChangeFee() *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee {
	return s.ChangeFee
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetGender() *int32 {
	return s.Gender
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetOriginTicketNos() []*string {
	return s.OriginTicketNos
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPhone() *string {
	return s.Phone
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPid() *int64 {
	return s.Pid
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetTicketNoSegmentMap() map[string]interface{} {
	return s.TicketNoSegmentMap
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) GetTicketNos() []*string {
	return s.TicketNos
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetBirthDate(v string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.BirthDate = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetCertNo(v string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.CertNo = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetCertType(v int32) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.CertType = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetChangeFee(v *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.ChangeFee = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetGender(v int32) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.Gender = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetOriginTicketNos(v []*string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.OriginTicketNos = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPassengerId(v string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.PassengerId = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPassengerName(v string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.PassengerName = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPassengerType(v int32) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.PassengerType = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPhone(v string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.Phone = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPid(v int64) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.Pid = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetTicketNoSegmentMap(v map[string]interface{}) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.TicketNoSegmentMap = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) SetTicketNos(v []*string) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.TicketNos = v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOS) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee struct {
	ChangeFee    *int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ServiceFee   *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	UpgradePrice *int64 `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) GetChangeFee() *int64 {
	return s.ChangeFee
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) GetUpgradePrice() *int64 {
	return s.UpgradePrice
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) SetChangeFee(v int64) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee {
	s.ChangeFee = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) SetServiceFee(v int64) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee {
	s.ServiceFee = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) SetUpgradePrice(v int64) *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee {
	s.UpgradePrice = &v
	return s
}

func (s *FlightModifyOrderDetailV2ResponseBodyModuleTravelerInfoDTOSChangeFee) Validate() error {
	return dara.Validate(s)
}
