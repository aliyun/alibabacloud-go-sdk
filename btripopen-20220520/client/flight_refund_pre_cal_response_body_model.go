// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightRefundPreCalResponseBody
	GetCode() *string
	SetMessage(v string) *FlightRefundPreCalResponseBody
	GetMessage() *string
	SetModule(v *FlightRefundPreCalResponseBodyModule) *FlightRefundPreCalResponseBody
	GetModule() *FlightRefundPreCalResponseBodyModule
	SetRequestId(v string) *FlightRefundPreCalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightRefundPreCalResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightRefundPreCalResponseBody
	GetTraceId() *string
}

type FlightRefundPreCalResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightRefundPreCalResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightRefundPreCalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalResponseBody) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightRefundPreCalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightRefundPreCalResponseBody) GetModule() *FlightRefundPreCalResponseBodyModule {
	return s.Module
}

func (s *FlightRefundPreCalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightRefundPreCalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightRefundPreCalResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightRefundPreCalResponseBody) SetCode(v string) *FlightRefundPreCalResponseBody {
	s.Code = &v
	return s
}

func (s *FlightRefundPreCalResponseBody) SetMessage(v string) *FlightRefundPreCalResponseBody {
	s.Message = &v
	return s
}

func (s *FlightRefundPreCalResponseBody) SetModule(v *FlightRefundPreCalResponseBodyModule) *FlightRefundPreCalResponseBody {
	s.Module = v
	return s
}

func (s *FlightRefundPreCalResponseBody) SetRequestId(v string) *FlightRefundPreCalResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightRefundPreCalResponseBody) SetSuccess(v bool) *FlightRefundPreCalResponseBody {
	s.Success = &v
	return s
}

func (s *FlightRefundPreCalResponseBody) SetTraceId(v string) *FlightRefundPreCalResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightRefundPreCalResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalResponseBodyModule struct {
	// example:
	//
	// false
	FlightChange *bool `json:"flight_change,omitempty" xml:"flight_change,omitempty"`
	// example:
	//
	// FlightItem_996677504
	ItemUnitId         *string                                                   `json:"item_unit_id,omitempty" xml:"item_unit_id,omitempty"`
	MultiRefundCalList []*FlightRefundPreCalResponseBodyModuleMultiRefundCalList `json:"multi_refund_cal_list,omitempty" xml:"multi_refund_cal_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	PreRefundMoney *int64 `json:"pre_refund_money,omitempty" xml:"pre_refund_money,omitempty"`
	// example:
	//
	// 1000
	RefundFee    *int64                                              `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	ReturnReason []*FlightRefundPreCalResponseBodyModuleReturnReason `json:"return_reason,omitempty" xml:"return_reason,omitempty" type:"Repeated"`
	// example:
	//
	// 882sudu23s923j9d2
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	Tips      *string `json:"tips,omitempty" xml:"tips,omitempty"`
}

func (s FlightRefundPreCalResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalResponseBodyModule) GetFlightChange() *bool {
	return s.FlightChange
}

func (s *FlightRefundPreCalResponseBodyModule) GetItemUnitId() *string {
	return s.ItemUnitId
}

func (s *FlightRefundPreCalResponseBodyModule) GetMultiRefundCalList() []*FlightRefundPreCalResponseBodyModuleMultiRefundCalList {
	return s.MultiRefundCalList
}

func (s *FlightRefundPreCalResponseBodyModule) GetPreRefundMoney() *int64 {
	return s.PreRefundMoney
}

func (s *FlightRefundPreCalResponseBodyModule) GetRefundFee() *int64 {
	return s.RefundFee
}

func (s *FlightRefundPreCalResponseBodyModule) GetReturnReason() []*FlightRefundPreCalResponseBodyModuleReturnReason {
	return s.ReturnReason
}

func (s *FlightRefundPreCalResponseBodyModule) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightRefundPreCalResponseBodyModule) GetTips() *string {
	return s.Tips
}

func (s *FlightRefundPreCalResponseBodyModule) SetFlightChange(v bool) *FlightRefundPreCalResponseBodyModule {
	s.FlightChange = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetItemUnitId(v string) *FlightRefundPreCalResponseBodyModule {
	s.ItemUnitId = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetMultiRefundCalList(v []*FlightRefundPreCalResponseBodyModuleMultiRefundCalList) *FlightRefundPreCalResponseBodyModule {
	s.MultiRefundCalList = v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetPreRefundMoney(v int64) *FlightRefundPreCalResponseBodyModule {
	s.PreRefundMoney = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetRefundFee(v int64) *FlightRefundPreCalResponseBodyModule {
	s.RefundFee = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetReturnReason(v []*FlightRefundPreCalResponseBodyModuleReturnReason) *FlightRefundPreCalResponseBodyModule {
	s.ReturnReason = v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetSessionId(v string) *FlightRefundPreCalResponseBodyModule {
	s.SessionId = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) SetTips(v string) *FlightRefundPreCalResponseBodyModule {
	s.Tips = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalResponseBodyModuleMultiRefundCalList struct {
	// example:
	//
	// true
	CanApplyRefund *bool   `json:"can_apply_refund,omitempty" xml:"can_apply_refund,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1000
	PreRefundMoney *int64 `json:"pre_refund_money,omitempty" xml:"pre_refund_money,omitempty"`
	// example:
	//
	// 1000
	RefundFee *int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// example:
	//
	// 124
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightRefundPreCalResponseBodyModuleMultiRefundCalList) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalResponseBodyModuleMultiRefundCalList) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) GetCanApplyRefund() *bool {
	return s.CanApplyRefund
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) GetName() *string {
	return s.Name
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) GetPreRefundMoney() *int64 {
	return s.PreRefundMoney
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) GetRefundFee() *int64 {
	return s.RefundFee
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) GetUserId() *string {
	return s.UserId
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) SetCanApplyRefund(v bool) *FlightRefundPreCalResponseBodyModuleMultiRefundCalList {
	s.CanApplyRefund = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) SetName(v string) *FlightRefundPreCalResponseBodyModuleMultiRefundCalList {
	s.Name = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) SetPreRefundMoney(v int64) *FlightRefundPreCalResponseBodyModuleMultiRefundCalList {
	s.PreRefundMoney = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) SetRefundFee(v int64) *FlightRefundPreCalResponseBodyModuleMultiRefundCalList {
	s.RefundFee = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) SetUserId(v string) *FlightRefundPreCalResponseBodyModuleMultiRefundCalList {
	s.UserId = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleMultiRefundCalList) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalResponseBodyModuleReturnReason struct {
	ExtendDesc *string `json:"extend_desc,omitempty" xml:"extend_desc,omitempty"`
	// example:
	//
	// 0
	Person *int32 `json:"person,omitempty" xml:"person,omitempty"`
	// example:
	//
	// 0
	ReasonCode *int32  `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	ReasonShow *string `json:"reason_show,omitempty" xml:"reason_show,omitempty"`
	// example:
	//
	// 1
	ReasonType *int32 `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
	// example:
	//
	// 0
	Volunteer *int32 `json:"volunteer,omitempty" xml:"volunteer,omitempty"`
}

func (s FlightRefundPreCalResponseBodyModuleReturnReason) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalResponseBodyModuleReturnReason) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) GetExtendDesc() *string {
	return s.ExtendDesc
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) GetPerson() *int32 {
	return s.Person
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) GetReasonCode() *int32 {
	return s.ReasonCode
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) GetReasonShow() *string {
	return s.ReasonShow
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) GetReasonType() *int32 {
	return s.ReasonType
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) GetVolunteer() *int32 {
	return s.Volunteer
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) SetExtendDesc(v string) *FlightRefundPreCalResponseBodyModuleReturnReason {
	s.ExtendDesc = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) SetPerson(v int32) *FlightRefundPreCalResponseBodyModuleReturnReason {
	s.Person = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) SetReasonCode(v int32) *FlightRefundPreCalResponseBodyModuleReturnReason {
	s.ReasonCode = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) SetReasonShow(v string) *FlightRefundPreCalResponseBodyModuleReturnReason {
	s.ReasonShow = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) SetReasonType(v int32) *FlightRefundPreCalResponseBodyModuleReturnReason {
	s.ReasonType = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) SetVolunteer(v int32) *FlightRefundPreCalResponseBodyModuleReturnReason {
	s.Volunteer = &v
	return s
}

func (s *FlightRefundPreCalResponseBodyModuleReturnReason) Validate() error {
	return dara.Validate(s)
}
