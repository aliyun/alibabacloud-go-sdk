// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundConsultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightRefundConsultResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightRefundConsultResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightRefundConsultResponseBodyModule) *IntlFlightRefundConsultResponseBody
	GetModule() *IntlFlightRefundConsultResponseBodyModule
	SetRequestId(v string) *IntlFlightRefundConsultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightRefundConsultResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightRefundConsultResponseBody
	GetTraceId() *string
}

type IntlFlightRefundConsultResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightRefundConsultResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IntlFlightRefundConsultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightRefundConsultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightRefundConsultResponseBody) GetModule() *IntlFlightRefundConsultResponseBodyModule {
	return s.Module
}

func (s *IntlFlightRefundConsultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightRefundConsultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightRefundConsultResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightRefundConsultResponseBody) SetCode(v string) *IntlFlightRefundConsultResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBody) SetMessage(v string) *IntlFlightRefundConsultResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBody) SetModule(v *IntlFlightRefundConsultResponseBodyModule) *IntlFlightRefundConsultResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightRefundConsultResponseBody) SetRequestId(v string) *IntlFlightRefundConsultResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBody) SetSuccess(v bool) *IntlFlightRefundConsultResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBody) SetTraceId(v string) *IntlFlightRefundConsultResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundConsultResponseBodyModule struct {
	PassengerJourneyGroupInfoList []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList `json:"passenger_journey_group_info_list,omitempty" xml:"passenger_journey_group_info_list,omitempty" type:"Repeated"`
}

func (s IntlFlightRefundConsultResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBodyModule) GetPassengerJourneyGroupInfoList() []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList {
	return s.PassengerJourneyGroupInfoList
}

func (s *IntlFlightRefundConsultResponseBodyModule) SetPassengerJourneyGroupInfoList(v []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) *IntlFlightRefundConsultResponseBodyModule {
	s.PassengerJourneyGroupInfoList = v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList struct {
	// example:
	//
	// edcac4f4c79d40ccb141ddb6da567e65
	PassengerJourneyGroupKey       *string                                                                                                 `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	PassengerList                  []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList                  `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PassengerSegmentStatusInfoList []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList `json:"passenger_segment_status_info_list,omitempty" xml:"passenger_segment_status_info_list,omitempty" type:"Repeated"`
	RefundReasonInfoList           []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList           `json:"refund_reason_info_list,omitempty" xml:"refund_reason_info_list,omitempty" type:"Repeated"`
	SegmentList                    []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList                    `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) GetPassengerList() []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList {
	return s.PassengerList
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) GetPassengerSegmentStatusInfoList() []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	return s.PassengerSegmentStatusInfoList
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) GetRefundReasonInfoList() []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList {
	return s.RefundReasonInfoList
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) GetSegmentList() []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	return s.SegmentList
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) SetPassengerJourneyGroupKey(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) SetPassengerList(v []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.PassengerList = v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) SetPassengerSegmentStatusInfoList(v []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.PassengerSegmentStatusInfoList = v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) SetRefundReasonInfoList(v []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.RefundReasonInfoList = v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) SetSegmentList(v []*IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.SegmentList = v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList struct {
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// example:
	//
	// 1000001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) SetFullName(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) SetPassengerId(v int64) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList struct {
	// example:
	//
	// true
	CanRefund *bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// example:
	//
	// 1000001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// CZ5009PKXHKG0616
	SegmentKey     *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	UnRefundReason *string `json:"un_refund_reason,omitempty" xml:"un_refund_reason,omitempty"`
	// example:
	//
	// 3
	UnRefundReasonCode *string `json:"un_refund_reason_code,omitempty" xml:"un_refund_reason_code,omitempty"`
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetUnRefundReason() *string {
	return s.UnRefundReason
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetUnRefundReasonCode() *string {
	return s.UnRefundReasonCode
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetCanRefund(v bool) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.CanRefund = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetPassengerId(v int64) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetSegmentKey(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetUnRefundReason(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.UnRefundReason = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetUnRefundReasonCode(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.UnRefundReasonCode = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList struct {
	// example:
	//
	// 0
	ReasonCode *string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	ReasonDesc *string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) GetReasonDesc() *string {
	return s.ReasonDesc
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) SetReasonCode(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList {
	s.ReasonCode = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) SetReasonDesc(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList {
	s.ReasonDesc = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) SetVoluntary(v bool) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList {
	s.Voluntary = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListRefundReasonInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList struct {
	// example:
	//
	// HKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2025-06-16 19:20
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// CZ5009
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// CZ5009PKXHKG0616
	SegmentKey *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetArrCityCode(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetDepCityCode(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetDepTime(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.DepTime = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetFlightNo(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetJourneyIndex(v int32) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetSegmentIndex(v int32) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetSegmentKey(v string) *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightRefundConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) Validate() error {
	return dara.Validate(s)
}
