// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightRefundDetailV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightRefundDetailV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightRefundDetailV2ResponseBodyModule) *FlightRefundDetailV2ResponseBody
	GetModule() *FlightRefundDetailV2ResponseBodyModule
	SetRequestId(v string) *FlightRefundDetailV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightRefundDetailV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightRefundDetailV2ResponseBody
	GetTraceId() *string
}

type FlightRefundDetailV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightRefundDetailV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210e845f16785007404904300ddc92
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightRefundDetailV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightRefundDetailV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightRefundDetailV2ResponseBody) GetModule() *FlightRefundDetailV2ResponseBodyModule {
	return s.Module
}

func (s *FlightRefundDetailV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightRefundDetailV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightRefundDetailV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightRefundDetailV2ResponseBody) SetCode(v string) *FlightRefundDetailV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBody) SetMessage(v string) *FlightRefundDetailV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBody) SetModule(v *FlightRefundDetailV2ResponseBodyModule) *FlightRefundDetailV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightRefundDetailV2ResponseBody) SetRequestId(v string) *FlightRefundDetailV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBody) SetSuccess(v bool) *FlightRefundDetailV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBody) SetTraceId(v string) *FlightRefundDetailV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailV2ResponseBodyModule struct {
	// example:
	//
	// 2023-08-19 17:18:19
	ApplyTime                       *string                                                 `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	ContactInfoDTO                  *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO   `json:"contact_info_d_t_o,omitempty" xml:"contact_info_d_t_o,omitempty" type:"Struct"`
	FlightInfoDTOS                  []*FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS `json:"flight_info_d_t_o_s,omitempty" xml:"flight_info_d_t_o_s,omitempty" type:"Repeated"`
	NonRefundableChangeServicePrice *int64                                                  `json:"non_refundable_change_service_price,omitempty" xml:"non_refundable_change_service_price,omitempty"`
	NonRefundableChangeUpgradePrice *int64                                                  `json:"non_refundable_change_upgrade_price,omitempty" xml:"non_refundable_change_upgrade_price,omitempty"`
	// example:
	//
	// 1002039195025156700
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1002039195025156700
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195836916039
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	Reason        *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// 1
	ReasonCode       *string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	RefundFailReason *string `json:"refund_fail_reason,omitempty" xml:"refund_fail_reason,omitempty"`
	// example:
	//
	// 100
	RefundHandlingFee *int64 `json:"refund_handling_fee,omitempty" xml:"refund_handling_fee,omitempty"`
	// example:
	//
	// 10000
	RefundMoney *int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1000000000297003
	SubOrderId       *int64                                                    `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	TravelerInfoDTOS []*FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS `json:"traveler_info_d_t_o_s,omitempty" xml:"traveler_info_d_t_o_s,omitempty" type:"Repeated"`
}

func (s FlightRefundDetailV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetContactInfoDTO() *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO {
	return s.ContactInfoDTO
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetFlightInfoDTOS() []*FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	return s.FlightInfoDTOS
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetNonRefundableChangeServicePrice() *int64 {
	return s.NonRefundableChangeServicePrice
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetNonRefundableChangeUpgradePrice() *int64 {
	return s.NonRefundableChangeUpgradePrice
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetReason() *string {
	return s.Reason
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetRefundFailReason() *string {
	return s.RefundFailReason
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetRefundHandlingFee() *int64 {
	return s.RefundHandlingFee
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetRefundMoney() *int64 {
	return s.RefundMoney
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *FlightRefundDetailV2ResponseBodyModule) GetTravelerInfoDTOS() []*FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	return s.TravelerInfoDTOS
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetApplyTime(v string) *FlightRefundDetailV2ResponseBodyModule {
	s.ApplyTime = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetContactInfoDTO(v *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) *FlightRefundDetailV2ResponseBodyModule {
	s.ContactInfoDTO = v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetFlightInfoDTOS(v []*FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) *FlightRefundDetailV2ResponseBodyModule {
	s.FlightInfoDTOS = v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetNonRefundableChangeServicePrice(v int64) *FlightRefundDetailV2ResponseBodyModule {
	s.NonRefundableChangeServicePrice = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetNonRefundableChangeUpgradePrice(v int64) *FlightRefundDetailV2ResponseBodyModule {
	s.NonRefundableChangeUpgradePrice = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetOrderId(v int64) *FlightRefundDetailV2ResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetOutOrderId(v string) *FlightRefundDetailV2ResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetOutSubOrderId(v string) *FlightRefundDetailV2ResponseBodyModule {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetReason(v string) *FlightRefundDetailV2ResponseBodyModule {
	s.Reason = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetReasonCode(v string) *FlightRefundDetailV2ResponseBodyModule {
	s.ReasonCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetRefundFailReason(v string) *FlightRefundDetailV2ResponseBodyModule {
	s.RefundFailReason = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetRefundHandlingFee(v int64) *FlightRefundDetailV2ResponseBodyModule {
	s.RefundHandlingFee = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetRefundMoney(v int64) *FlightRefundDetailV2ResponseBodyModule {
	s.RefundMoney = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetStatus(v int32) *FlightRefundDetailV2ResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetSubOrderId(v int64) *FlightRefundDetailV2ResponseBodyModule {
	s.SubOrderId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) SetTravelerInfoDTOS(v []*FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) *FlightRefundDetailV2ResponseBodyModule {
	s.TravelerInfoDTOS = v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailV2ResponseBodyModuleContactInfoDTO struct {
	// example:
	//
	// 17685764353@163.com
	ContactEmail *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	ContactName  *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// example:
	//
	// 17685764353
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// true
	SendMsgToPassenger *bool `json:"send_msg_to_passenger,omitempty" xml:"send_msg_to_passenger,omitempty"`
}

func (s FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) GetContactName() *string {
	return s.ContactName
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) GetSendMsgToPassenger() *bool {
	return s.SendMsgToPassenger
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) SetContactEmail(v string) *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactEmail = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) SetContactName(v string) *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) SetContactPhone(v string) *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactPhone = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) SetSendMsgToPassenger(v bool) *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO {
	s.SendMsgToPassenger = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleContactInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS struct {
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
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// MU5193
	FlightNo   *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	MealDesc   *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 1194012
	SegmentId       *string                                                              `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	SegmentPosition *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	StopArrTime     *string                                                              `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCity        *string                                                              `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	StopDepTime     *string                                                              `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetAirlineIconUrl() *string {
	return s.AirlineIconUrl
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCabin() *string {
	return s.Cabin
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCabinDiscount() *int64 {
	return s.CabinDiscount
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCarrierAirlineCode() *string {
	return s.CarrierAirlineCode
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCarrierAirlineIconUrl() *string {
	return s.CarrierAirlineIconUrl
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCarrierAirlineName() *string {
	return s.CarrierAirlineName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetCarrierFlightNo() *string {
	return s.CarrierFlightNo
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetSegmentId() *string {
	return s.SegmentId
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetSegmentPosition() *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition {
	return s.SegmentPosition
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetAirlineCode(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.AirlineCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetAirlineIconUrl(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.AirlineIconUrl = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetAirlineName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.AirlineName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetArrAirportCode(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetArrAirportName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.ArrAirportName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetArrCityCode(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.ArrCityCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetArrCityName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.ArrCityName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetArrTerminal(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.ArrTerminal = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetArrTime(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.ArrTime = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCabin(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.Cabin = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCabinClass(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CabinClass = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCabinClassName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CabinClassName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCabinDiscount(v int64) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CabinDiscount = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCarrierAirlineCode(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CarrierAirlineCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCarrierAirlineIconUrl(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CarrierAirlineIconUrl = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCarrierAirlineName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CarrierAirlineName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetCarrierFlightNo(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.CarrierFlightNo = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetDepAirportCode(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.DepAirportCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetDepAirportName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.DepAirportName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetDepCityCode(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.DepCityCode = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetDepCityName(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.DepCityName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetDepTerminal(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.DepTerminal = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetDepTime(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.DepTime = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetFlightNo(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.FlightNo = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetFlightType(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.FlightType = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetMealDesc(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.MealDesc = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetSegmentId(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.SegmentId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetSegmentPosition(v *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.SegmentPosition = v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetStopArrTime(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.StopArrTime = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetStopCity(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.StopCity = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) SetStopDepTime(v string) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS {
	s.StopDepTime = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOS) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) SetJourneyIndex(v int32) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) SetSegmentIndex(v int32) *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleFlightInfoDTOSSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS struct {
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
	CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
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
	Pid       *int64    `json:"pid,omitempty" xml:"pid,omitempty"`
	TicketNos []*string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
}

func (s FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetBirthDate() *string {
	return s.BirthDate
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetCertNo() *string {
	return s.CertNo
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetCertType() *int32 {
	return s.CertType
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetGender() *int32 {
	return s.Gender
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetOriginTicketNos() []*string {
	return s.OriginTicketNos
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPhone() *string {
	return s.Phone
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetPid() *int64 {
	return s.Pid
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) GetTicketNos() []*string {
	return s.TicketNos
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetBirthDate(v string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.BirthDate = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetCertNo(v string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.CertNo = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetCertType(v int32) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.CertType = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetGender(v int32) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.Gender = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetOriginTicketNos(v []*string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.OriginTicketNos = v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPassengerId(v string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.PassengerId = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPassengerName(v string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.PassengerName = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPassengerType(v int32) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.PassengerType = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPhone(v string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.Phone = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetPid(v int64) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.Pid = &v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) SetTicketNos(v []*string) *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS {
	s.TicketNos = v
	return s
}

func (s *FlightRefundDetailV2ResponseBodyModuleTravelerInfoDTOS) Validate() error {
	return dara.Validate(s)
}
