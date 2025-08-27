// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightRefundPreCalV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightRefundPreCalV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightRefundPreCalV2ResponseBodyModule) *FlightRefundPreCalV2ResponseBody
	GetModule() *FlightRefundPreCalV2ResponseBodyModule
	SetRequestId(v string) *FlightRefundPreCalV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightRefundPreCalV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightRefundPreCalV2ResponseBody
	GetTraceId() *string
}

type FlightRefundPreCalV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightRefundPreCalV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 92359A00-85D8-16C4-AED0-249618DEEC17
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightRefundPreCalV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightRefundPreCalV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightRefundPreCalV2ResponseBody) GetModule() *FlightRefundPreCalV2ResponseBodyModule {
	return s.Module
}

func (s *FlightRefundPreCalV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightRefundPreCalV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightRefundPreCalV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightRefundPreCalV2ResponseBody) SetCode(v string) *FlightRefundPreCalV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBody) SetMessage(v string) *FlightRefundPreCalV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBody) SetModule(v *FlightRefundPreCalV2ResponseBodyModule) *FlightRefundPreCalV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightRefundPreCalV2ResponseBody) SetRequestId(v string) *FlightRefundPreCalV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBody) SetSuccess(v bool) *FlightRefundPreCalV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBody) SetTraceId(v string) *FlightRefundPreCalV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalV2ResponseBodyModule struct {
	MultiRefundFeeDTOS []*FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS `json:"multi_refund_fee_d_t_o_s,omitempty" xml:"multi_refund_fee_d_t_o_s,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	PreRefundMoney *int64 `json:"pre_refund_money,omitempty" xml:"pre_refund_money,omitempty"`
	// example:
	//
	// 100
	RefundChargeFee        *int64                                                          `json:"refund_charge_fee,omitempty" xml:"refund_charge_fee,omitempty"`
	RefundReasonOptionDTOS []*FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS `json:"refund_reason_option_d_t_o_s,omitempty" xml:"refund_reason_option_d_t_o_s,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	ServiceChargeFee *int64 `json:"service_charge_fee,omitempty" xml:"service_charge_fee,omitempty"`
}

func (s FlightRefundPreCalV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2ResponseBodyModule) GetMultiRefundFeeDTOS() []*FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS {
	return s.MultiRefundFeeDTOS
}

func (s *FlightRefundPreCalV2ResponseBodyModule) GetPreRefundMoney() *int64 {
	return s.PreRefundMoney
}

func (s *FlightRefundPreCalV2ResponseBodyModule) GetRefundChargeFee() *int64 {
	return s.RefundChargeFee
}

func (s *FlightRefundPreCalV2ResponseBodyModule) GetRefundReasonOptionDTOS() []*FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS {
	return s.RefundReasonOptionDTOS
}

func (s *FlightRefundPreCalV2ResponseBodyModule) GetServiceChargeFee() *int64 {
	return s.ServiceChargeFee
}

func (s *FlightRefundPreCalV2ResponseBodyModule) SetMultiRefundFeeDTOS(v []*FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) *FlightRefundPreCalV2ResponseBodyModule {
	s.MultiRefundFeeDTOS = v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModule) SetPreRefundMoney(v int64) *FlightRefundPreCalV2ResponseBodyModule {
	s.PreRefundMoney = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModule) SetRefundChargeFee(v int64) *FlightRefundPreCalV2ResponseBodyModule {
	s.RefundChargeFee = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModule) SetRefundReasonOptionDTOS(v []*FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) *FlightRefundPreCalV2ResponseBodyModule {
	s.RefundReasonOptionDTOS = v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModule) SetServiceChargeFee(v int64) *FlightRefundPreCalV2ResponseBodyModule {
	s.ServiceChargeFee = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS struct {
	// example:
	//
	// 3243028
	PassengerId   *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 10000
	PreRefundMoney *int64 `json:"pre_refund_money,omitempty" xml:"pre_refund_money,omitempty"`
	// example:
	//
	// 100
	RefundChargeFee *int64 `json:"refund_charge_fee,omitempty" xml:"refund_charge_fee,omitempty"`
}

func (s FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) GetPreRefundMoney() *int64 {
	return s.PreRefundMoney
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) GetRefundChargeFee() *int64 {
	return s.RefundChargeFee
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) SetPassengerId(v string) *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS {
	s.PassengerId = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) SetPassengerName(v string) *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS {
	s.PassengerName = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) SetPreRefundMoney(v int64) *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS {
	s.PreRefundMoney = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) SetRefundChargeFee(v int64) *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS {
	s.RefundChargeFee = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleMultiRefundFeeDTOS) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS struct {
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// 2
	ReasonType *int32 `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
	// example:
	//
	// true
	Volunteer *bool `json:"volunteer,omitempty" xml:"volunteer,omitempty"`
}

func (s FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) GetReason() *string {
	return s.Reason
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) GetReasonType() *int32 {
	return s.ReasonType
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) GetVolunteer() *bool {
	return s.Volunteer
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) SetReason(v string) *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS {
	s.Reason = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) SetReasonType(v int32) *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS {
	s.ReasonType = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) SetVolunteer(v bool) *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS {
	s.Volunteer = &v
	return s
}

func (s *FlightRefundPreCalV2ResponseBodyModuleRefundReasonOptionDTOS) Validate() error {
	return dara.Validate(s)
}
