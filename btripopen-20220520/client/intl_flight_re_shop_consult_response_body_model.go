// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopConsultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopConsultResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopConsultResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopConsultResponseBodyModule) *IntlFlightReShopConsultResponseBody
	GetModule() *IntlFlightReShopConsultResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopConsultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopConsultResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopConsultResponseBody
	GetTraceId() *string
}

type IntlFlightReShopConsultResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Module    *IntlFlightReShopConsultResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightReShopConsultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopConsultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopConsultResponseBody) GetModule() *IntlFlightReShopConsultResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopConsultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopConsultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopConsultResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopConsultResponseBody) SetCode(v string) *IntlFlightReShopConsultResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBody) SetMessage(v string) *IntlFlightReShopConsultResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBody) SetModule(v *IntlFlightReShopConsultResponseBodyModule) *IntlFlightReShopConsultResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopConsultResponseBody) SetRequestId(v string) *IntlFlightReShopConsultResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBody) SetSuccess(v bool) *IntlFlightReShopConsultResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBody) SetTraceId(v string) *IntlFlightReShopConsultResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopConsultResponseBodyModule struct {
	PassengerJourneyGroupInfoList []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList `json:"passenger_journey_group_info_list,omitempty" xml:"passenger_journey_group_info_list,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopConsultResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBodyModule) GetPassengerJourneyGroupInfoList() []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList {
	return s.PassengerJourneyGroupInfoList
}

func (s *IntlFlightReShopConsultResponseBodyModule) SetPassengerJourneyGroupInfoList(v []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) *IntlFlightReShopConsultResponseBodyModule {
	s.PassengerJourneyGroupInfoList = v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList struct {
	PassengerJourneyGroupKey       *string                                                                                                 `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	PassengerList                  []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList                  `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PassengerSegmentStatusInfoList []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList `json:"passenger_segment_status_info_list,omitempty" xml:"passenger_segment_status_info_list,omitempty" type:"Repeated"`
	ReShopReasonInfoList           []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList           `json:"re_shop_reason_info_list,omitempty" xml:"re_shop_reason_info_list,omitempty" type:"Repeated"`
	SegmentList                    []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList                    `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) GetPassengerList() []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList {
	return s.PassengerList
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) GetPassengerSegmentStatusInfoList() []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	return s.PassengerSegmentStatusInfoList
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) GetReShopReasonInfoList() []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList {
	return s.ReShopReasonInfoList
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) GetSegmentList() []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	return s.SegmentList
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) SetPassengerList(v []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.PassengerList = v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) SetPassengerSegmentStatusInfoList(v []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.PassengerSegmentStatusInfoList = v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) SetReShopReasonInfoList(v []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.ReShopReasonInfoList = v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) SetSegmentList(v []*IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList {
	s.SegmentList = v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList struct {
	FullName    *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	PassengerId *int64  `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) SetFullName(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) SetPassengerId(v int64) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList struct {
	CanReShop          *bool   `json:"can_re_shop,omitempty" xml:"can_re_shop,omitempty"`
	PassengerId        *int64  `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	SegmentKey         *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	UnReShopReason     *string `json:"un_re_shop_reason,omitempty" xml:"un_re_shop_reason,omitempty"`
	UnReShopReasonCode *string `json:"un_re_shop_reason_code,omitempty" xml:"un_re_shop_reason_code,omitempty"`
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetCanReShop() *bool {
	return s.CanReShop
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetUnReShopReason() *string {
	return s.UnReShopReason
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) GetUnReShopReasonCode() *string {
	return s.UnReShopReasonCode
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetCanReShop(v bool) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.CanReShop = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetPassengerId(v int64) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetSegmentKey(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetUnReShopReason(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.UnReShopReason = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) SetUnReShopReasonCode(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList {
	s.UnReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListPassengerSegmentStatusInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList struct {
	ReasonCode *string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	ReasonDesc *string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
	Voluntary  *bool   `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) GetReasonDesc() *string {
	return s.ReasonDesc
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) SetReasonCode(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList {
	s.ReasonCode = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) SetReasonDesc(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList {
	s.ReasonDesc = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) SetVoluntary(v bool) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList {
	s.Voluntary = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListReShopReasonInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList struct {
	ArrCityCode  *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	DepCityCode  *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepTime      *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightNo     *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	JourneyIndex *int32  `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32  `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	SegmentKey   *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetArrCityCode(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetDepCityCode(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetDepTime(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetFlightNo(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetJourneyIndex(v int32) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetSegmentIndex(v int32) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) SetSegmentKey(v string) *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopConsultResponseBodyModulePassengerJourneyGroupInfoListSegmentList) Validate() error {
	return dara.Validate(s)
}
