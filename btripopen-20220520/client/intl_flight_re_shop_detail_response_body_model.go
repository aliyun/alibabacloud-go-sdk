// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopDetailResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopDetailResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopDetailResponseBodyModule) *IntlFlightReShopDetailResponseBody
	GetModule() *IntlFlightReShopDetailResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopDetailResponseBody
	GetTraceId() *string
}

type IntlFlightReShopDetailResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightReShopDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopDetailResponseBody) GetModule() *IntlFlightReShopDetailResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopDetailResponseBody) SetCode(v string) *IntlFlightReShopDetailResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBody) SetMessage(v string) *IntlFlightReShopDetailResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBody) SetModule(v *IntlFlightReShopDetailResponseBodyModule) *IntlFlightReShopDetailResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopDetailResponseBody) SetRequestId(v string) *IntlFlightReShopDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBody) SetSuccess(v bool) *IntlFlightReShopDetailResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBody) SetTraceId(v string) *IntlFlightReShopDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModule struct {
	BaggageRule *IntlFlightReShopDetailResponseBodyModuleBaggageRule `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty" type:"Struct"`
	CloseReason *string                                              `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// example:
	//
	// 2024-03-06 15:00:35
	CloseTime *string `json:"close_time,omitempty" xml:"close_time,omitempty"`
	// example:
	//
	// 1
	CloseType   *int32                                                 `json:"close_type,omitempty" xml:"close_type,omitempty"`
	JourneyList []*IntlFlightReShopDetailResponseBodyModuleJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1017124195788186048
	OrderId           *int64                                                       `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OriginJourneyList []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyList `json:"origin_journey_list,omitempty" xml:"origin_journey_list,omitempty" type:"Repeated"`
	// example:
	//
	// F11552194294228713472
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// JPM20241024354
	OutReShopApplyId       *string                                                           `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	PassengerList          []*IntlFlightReShopDetailResponseBodyModulePassengerList          `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PassengerPriceInfoList []*IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList `json:"passenger_price_info_list,omitempty" xml:"passenger_price_info_list,omitempty" type:"Repeated"`
	PassengerTicketList    []*IntlFlightReShopDetailResponseBodyModulePassengerTicketList    `json:"passenger_ticket_list,omitempty" xml:"passenger_ticket_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-03-06 15:35:00
	PayLatestTime *string `json:"pay_latest_time,omitempty" xml:"pay_latest_time,omitempty"`
	// example:
	//
	// 0
	PayStatus *int32                                             `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	PriceInfo *IntlFlightReShopDetailResponseBodyModulePriceInfo `json:"price_info,omitempty" xml:"price_info,omitempty" type:"Struct"`
	// example:
	//
	// 1017035199374643191
	ReShopApplyId *int64 `json:"re_shop_apply_id,omitempty" xml:"re_shop_apply_id,omitempty"`
	// example:
	//
	// 0
	ReShopReasonCode *string                                                   `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	ReShopReasonDesc *string                                                   `json:"re_shop_reason_desc,omitempty" xml:"re_shop_reason_desc,omitempty"`
	RefundChangeRule *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule `json:"refund_change_rule,omitempty" xml:"refund_change_rule,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2024-03-06 15:01:35
	SuccessTime       *string `json:"success_time,omitempty" xml:"success_time,omitempty"`
	UserIntentionMemo *string `json:"user_intention_memo,omitempty" xml:"user_intention_memo,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetBaggageRule() *IntlFlightReShopDetailResponseBodyModuleBaggageRule {
	return s.BaggageRule
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetCloseReason() *string {
	return s.CloseReason
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetCloseTime() *string {
	return s.CloseTime
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetCloseType() *int32 {
	return s.CloseType
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetJourneyList() []*IntlFlightReShopDetailResponseBodyModuleJourneyList {
	return s.JourneyList
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetOrderId() *int64 {
	return s.OrderId
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetOriginJourneyList() []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	return s.OriginJourneyList
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetPassengerList() []*IntlFlightReShopDetailResponseBodyModulePassengerList {
	return s.PassengerList
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetPassengerPriceInfoList() []*IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList {
	return s.PassengerPriceInfoList
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetPassengerTicketList() []*IntlFlightReShopDetailResponseBodyModulePassengerTicketList {
	return s.PassengerTicketList
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetPayLatestTime() *string {
	return s.PayLatestTime
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetPriceInfo() *IntlFlightReShopDetailResponseBodyModulePriceInfo {
	return s.PriceInfo
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetReShopApplyId() *int64 {
	return s.ReShopApplyId
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetReShopReasonDesc() *string {
	return s.ReShopReasonDesc
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetRefundChangeRule() *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule {
	return s.RefundChangeRule
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetSuccessTime() *string {
	return s.SuccessTime
}

func (s *IntlFlightReShopDetailResponseBodyModule) GetUserIntentionMemo() *string {
	return s.UserIntentionMemo
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetBaggageRule(v *IntlFlightReShopDetailResponseBodyModuleBaggageRule) *IntlFlightReShopDetailResponseBodyModule {
	s.BaggageRule = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetCloseReason(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.CloseReason = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetCloseTime(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.CloseTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetCloseType(v int32) *IntlFlightReShopDetailResponseBodyModule {
	s.CloseType = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetJourneyList(v []*IntlFlightReShopDetailResponseBodyModuleJourneyList) *IntlFlightReShopDetailResponseBodyModule {
	s.JourneyList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetOrderId(v int64) *IntlFlightReShopDetailResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetOriginJourneyList(v []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) *IntlFlightReShopDetailResponseBodyModule {
	s.OriginJourneyList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetOutOrderId(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetOutReShopApplyId(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetPassengerList(v []*IntlFlightReShopDetailResponseBodyModulePassengerList) *IntlFlightReShopDetailResponseBodyModule {
	s.PassengerList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetPassengerPriceInfoList(v []*IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) *IntlFlightReShopDetailResponseBodyModule {
	s.PassengerPriceInfoList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetPassengerTicketList(v []*IntlFlightReShopDetailResponseBodyModulePassengerTicketList) *IntlFlightReShopDetailResponseBodyModule {
	s.PassengerTicketList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetPayLatestTime(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.PayLatestTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetPayStatus(v int32) *IntlFlightReShopDetailResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetPriceInfo(v *IntlFlightReShopDetailResponseBodyModulePriceInfo) *IntlFlightReShopDetailResponseBodyModule {
	s.PriceInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetReShopApplyId(v int64) *IntlFlightReShopDetailResponseBodyModule {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetReShopReasonCode(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetReShopReasonDesc(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.ReShopReasonDesc = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetRefundChangeRule(v *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule) *IntlFlightReShopDetailResponseBodyModule {
	s.RefundChangeRule = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetStatus(v int32) *IntlFlightReShopDetailResponseBodyModule {
	s.Status = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetSuccessTime(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.SuccessTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) SetUserIntentionMemo(v string) *IntlFlightReShopDetailResponseBodyModule {
	s.UserIntentionMemo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleBaggageRule struct {
	BaggageRuleDesc *string `json:"baggage_rule_desc,omitempty" xml:"baggage_rule_desc,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleBaggageRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleBaggageRule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleBaggageRule) GetBaggageRuleDesc() *string {
	return s.BaggageRuleDesc
}

func (s *IntlFlightReShopDetailResponseBodyModuleBaggageRule) SetBaggageRuleDesc(v string) *IntlFlightReShopDetailResponseBodyModuleBaggageRule {
	s.BaggageRuleDesc = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleBaggageRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyList struct {
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2025-01-01 02:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2025-01-01 01:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 60
	Duration           *int32                                                                   `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetFlightSegmentInfos() []*IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetArrCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetArrCityName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetArrTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetDepCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetDepCityName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetDepTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetDuration(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetFlightSegmentInfos(v []*IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetJourneyIndex(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) SetTransferTime(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyList {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos struct {
	AirlineInfo    *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2025-01-01 02:00
	ArrTime        *string                                                                              `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportInfo *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2025-01-01 01:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 130
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// NS8210
	FlightNo           *string                                                                                    `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                                    `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 737
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 0
	JourneyIndex      *int32                                                                                  `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	LuggageDirectInfo *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer      *string                                                                                 `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc          *string                                                                                 `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 0
	OneMore *int32 `json:"one_more,omitempty" xml:"one_more,omitempty"`
	// example:
	//
	// -
	OneMoreShow *string `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// NS8210XIYHGH0501
	SegmentKey        *string                                                                                 `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	// example:
	//
	// true
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// true
	Stop      *bool   `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetAirlineInfo() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrAirportInfo() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepAirportInfo() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightShareInfo() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetAirlineInfo(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrCityName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetArrTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepCityName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDepTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetDuration(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightNo(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightSize(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetFlightType(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetJourneyIndex(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetManufacturer(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetMealDesc(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetOneMore(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetShare(v bool) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetStop(v bool) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) SetTotalTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo struct {
	// example:
	//
	// NS
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo struct {
	// example:
	//
	// DLC
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T1
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo struct {
	// example:
	//
	// PEK
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T1
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CA1234
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList struct {
	// example:
	//
	// PEK
	StopAirport     *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	// example:
	//
	// T3
	StopArrTerm *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	// example:
	//
	// 2024-01-01 05:00
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// BJS
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	// example:
	//
	// T2
	StopDepTerm *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	// example:
	//
	// 2024-01-01 07:00
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 120
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosFlightStopInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo struct {
	// example:
	//
	// 1
	DepCityLuggageDirect *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	// example:
	//
	// 1
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark struct {
	DepCityVisaRemark *string `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	// example:
	//
	// 1
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleJourneyListFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyList struct {
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2025-01-01 02:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2025-01-01 01:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 60
	Duration           *int32                                                                         `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetFlightSegmentInfos() []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetArrCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetArrCityName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetArrTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetDepCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetDepCityName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetDepTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetDuration(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetFlightSegmentInfos(v []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetJourneyIndex(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) SetTransferTime(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos struct {
	AirlineInfo    *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:25
	ArrTime        *string                                                                                    `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportInfo *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 120
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// NS8210
	FlightNo           *string                                                                                          `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                                          `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 738
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 0
	JourneyIndex      *int32                                                                                        `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	LuggageDirectInfo *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer      *string                                                                                       `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc          *string                                                                                       `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 1
	OneMore     *int32  `json:"one_more,omitempty" xml:"one_more,omitempty"`
	OneMoreShow *string `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// NS8210XIYHGH0501
	SegmentKey        *string                                                                                       `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	// example:
	//
	// true
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// true
	Stop      *bool   `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetAirlineInfo() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetArrAirportInfo() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetDepAirportInfo() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetFlightShareInfo() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetAirlineInfo(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetArrCityName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetArrTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetDepCityName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetDepTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetDuration(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetFlightNo(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetFlightSize(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetFlightType(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetJourneyIndex(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetManufacturer(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetMealDesc(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetOneMore(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetShare(v bool) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetStop(v bool) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) SetTotalTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo struct {
	// example:
	//
	// HKG
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo struct {
	// example:
	//
	// PEK
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T1
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CA0001
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList struct {
	// example:
	//
	// HGH
	StopAirport     *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	// example:
	//
	// T1
	StopArrTerm *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// HGH
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	// example:
	//
	// T1
	StopDepTerm *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	// example:
	//
	// 2023-08-13 09:25
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 20
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosFlightStopInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo struct {
	// example:
	//
	// 1
	DepCityLuggageDirect *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	// example:
	//
	// 1
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark struct {
	DepCityVisaRemark *string `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	// example:
	//
	// 1
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleOriginJourneyListFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePassengerList struct {
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
	// 76230022
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
	// btrip8432002
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetFullName(v string) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetGender(v int32) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.Gender = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetJobNo(v string) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.JobNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetNationality(v string) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.Nationality = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetNationalityCode(v string) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.NationalityCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetPassengerId(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetType(v int32) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetUserId(v string) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.UserId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) SetUserType(v int32) *IntlFlightReShopDetailResponseBodyModulePassengerList {
	s.UserType = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList struct {
	// example:
	//
	// 100001
	PassengerId *int64                                                                   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	PriceInfo   *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo `json:"price_info,omitempty" xml:"price_info,omitempty" type:"Struct"`
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) GetPriceInfo() *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo {
	return s.PriceInfo
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) SetPassengerId(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) SetPriceInfo(v *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList {
	s.PriceInfo = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo struct {
	// example:
	//
	// 4000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// example:
	//
	// 0
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// example:
	//
	// 125000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// example:
	//
	// 2000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) SetHandlingAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo {
	s.HandlingAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) SetTaxDiffAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo {
	s.TaxDiffAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) SetTotalAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo {
	s.TotalAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) SetUpgradeAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo {
	s.UpgradeAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerPriceInfoListPriceInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePassengerTicketList struct {
	// example:
	//
	// 2345678
	PassengerId *int64                                                                   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	TicketList  []*IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList `json:"ticket_list,omitempty" xml:"ticket_list,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerTicketList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerTicketList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketList) GetTicketList() []*IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList {
	return s.TicketList
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketList) SetPassengerId(v int64) *IntlFlightReShopDetailResponseBodyModulePassengerTicketList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketList) SetTicketList(v []*IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) *IntlFlightReShopDetailResponseBodyModulePassengerTicketList {
	s.TicketList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList struct {
	// example:
	//
	// 2025-01-01 00:00:09
	IssueTime *string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// example:
	//
	// P123456
	PnrNo          *string   `json:"pnr_no,omitempty" xml:"pnr_no,omitempty"`
	SegmentKeyList []*string `json:"segment_key_list,omitempty" xml:"segment_key_list,omitempty" type:"Repeated"`
	// example:
	//
	// 781-9574833593
	TicketNo          *string                                                                                   `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketSegmentList []*IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList `json:"ticket_segment_list,omitempty" xml:"ticket_segment_list,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) GetIssueTime() *string {
	return s.IssueTime
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) GetPnrNo() *string {
	return s.PnrNo
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) GetSegmentKeyList() []*string {
	return s.SegmentKeyList
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) GetTicketSegmentList() []*IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	return s.TicketSegmentList
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) SetIssueTime(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList {
	s.IssueTime = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) SetPnrNo(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList {
	s.PnrNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) SetSegmentKeyList(v []*string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList {
	s.SegmentKeyList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) SetTicketNo(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList {
	s.TicketNo = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) SetTicketSegmentList(v []*IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList {
	s.TicketSegmentList = v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList struct {
	// example:
	//
	// G
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	Modified *bool `json:"modified,omitempty" xml:"modified,omitempty"`
	// example:
	//
	// OPEN_FOR_USE
	OpenTicketStatus *string `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
	// example:
	//
	// false
	Refunded *bool `json:"refunded,omitempty" xml:"refunded,omitempty"`
	// example:
	//
	// HX236HKGPVG0509
	SegmentKey *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetModified() *bool {
	return s.Modified
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetOpenTicketStatus() *string {
	return s.OpenTicketStatus
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetRefunded() *bool {
	return s.Refunded
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetCabin(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.Cabin = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetCabinClass(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.CabinClass = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetModified(v bool) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.Modified = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetOpenTicketStatus(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.OpenTicketStatus = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetRefunded(v bool) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.Refunded = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) SetSegmentKey(v string) *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePassengerTicketListTicketListTicketSegmentList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModulePriceInfo struct {
	// example:
	//
	// 4000
	HandlingAmount *int64 `json:"handling_amount,omitempty" xml:"handling_amount,omitempty"`
	// example:
	//
	// 0
	TaxDiffAmount *int64 `json:"tax_diff_amount,omitempty" xml:"tax_diff_amount,omitempty"`
	// example:
	//
	// 125000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// example:
	//
	// 2000
	UpgradeAmount *int64 `json:"upgrade_amount,omitempty" xml:"upgrade_amount,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModulePriceInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModulePriceInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) GetHandlingAmount() *int64 {
	return s.HandlingAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) GetTaxDiffAmount() *int64 {
	return s.TaxDiffAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) GetUpgradeAmount() *int64 {
	return s.UpgradeAmount
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) SetHandlingAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePriceInfo {
	s.HandlingAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) SetTaxDiffAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePriceInfo {
	s.TaxDiffAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) SetTotalAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePriceInfo {
	s.TotalAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) SetUpgradeAmount(v int64) *IntlFlightReShopDetailResponseBodyModulePriceInfo {
	s.UpgradeAmount = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModulePriceInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopDetailResponseBodyModuleRefundChangeRule struct {
	RefundChangeRuleDesc *string `json:"refund_change_rule_desc,omitempty" xml:"refund_change_rule_desc,omitempty"`
}

func (s IntlFlightReShopDetailResponseBodyModuleRefundChangeRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponseBodyModuleRefundChangeRule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule) GetRefundChangeRuleDesc() *string {
	return s.RefundChangeRuleDesc
}

func (s *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule) SetRefundChangeRuleDesc(v string) *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule {
	s.RefundChangeRuleDesc = &v
	return s
}

func (s *IntlFlightReShopDetailResponseBodyModuleRefundChangeRule) Validate() error {
	return dara.Validate(s)
}
