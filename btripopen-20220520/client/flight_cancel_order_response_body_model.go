// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightCancelOrderResponseBody
	GetCode() *string
	SetMessage(v string) *FlightCancelOrderResponseBody
	GetMessage() *string
	SetModule(v *FlightCancelOrderResponseBodyModule) *FlightCancelOrderResponseBody
	GetModule() *FlightCancelOrderResponseBodyModule
	SetRequestId(v string) *FlightCancelOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightCancelOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightCancelOrderResponseBody
	GetTraceId() *string
}

type FlightCancelOrderResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightCancelOrderResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightCancelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightCancelOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightCancelOrderResponseBody) GetModule() *FlightCancelOrderResponseBodyModule {
	return s.Module
}

func (s *FlightCancelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightCancelOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightCancelOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightCancelOrderResponseBody) SetCode(v string) *FlightCancelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *FlightCancelOrderResponseBody) SetMessage(v string) *FlightCancelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *FlightCancelOrderResponseBody) SetModule(v *FlightCancelOrderResponseBodyModule) *FlightCancelOrderResponseBody {
	s.Module = v
	return s
}

func (s *FlightCancelOrderResponseBody) SetRequestId(v string) *FlightCancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightCancelOrderResponseBody) SetSuccess(v bool) *FlightCancelOrderResponseBody {
	s.Success = &v
	return s
}

func (s *FlightCancelOrderResponseBody) SetTraceId(v string) *FlightCancelOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightCancelOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightCancelOrderResponseBodyModule struct {
	// example:
	//
	// 2022-07-04T16:13Z
	CancelTime *string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	FailCode   *string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// example:
	//
	// 0
	OrderStatus *string `json:"order_status,omitempty" xml:"order_status,omitempty"`
}

func (s FlightCancelOrderResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderResponseBodyModule) GetCancelTime() *string {
	return s.CancelTime
}

func (s *FlightCancelOrderResponseBodyModule) GetFailCode() *string {
	return s.FailCode
}

func (s *FlightCancelOrderResponseBodyModule) GetFailReason() *string {
	return s.FailReason
}

func (s *FlightCancelOrderResponseBodyModule) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *FlightCancelOrderResponseBodyModule) SetCancelTime(v string) *FlightCancelOrderResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *FlightCancelOrderResponseBodyModule) SetFailCode(v string) *FlightCancelOrderResponseBodyModule {
	s.FailCode = &v
	return s
}

func (s *FlightCancelOrderResponseBodyModule) SetFailReason(v string) *FlightCancelOrderResponseBodyModule {
	s.FailReason = &v
	return s
}

func (s *FlightCancelOrderResponseBodyModule) SetOrderStatus(v string) *FlightCancelOrderResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *FlightCancelOrderResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
