// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightCancelOrderV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightCancelOrderV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightCancelOrderV2ResponseBodyModule) *FlightCancelOrderV2ResponseBody
	GetModule() *FlightCancelOrderV2ResponseBodyModule
	SetRequestId(v string) *FlightCancelOrderV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightCancelOrderV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightCancelOrderV2ResponseBody
	GetTraceId() *string
}

type FlightCancelOrderV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightCancelOrderV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s FlightCancelOrderV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightCancelOrderV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightCancelOrderV2ResponseBody) GetModule() *FlightCancelOrderV2ResponseBodyModule {
	return s.Module
}

func (s *FlightCancelOrderV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightCancelOrderV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightCancelOrderV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightCancelOrderV2ResponseBody) SetCode(v string) *FlightCancelOrderV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightCancelOrderV2ResponseBody) SetMessage(v string) *FlightCancelOrderV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightCancelOrderV2ResponseBody) SetModule(v *FlightCancelOrderV2ResponseBodyModule) *FlightCancelOrderV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightCancelOrderV2ResponseBody) SetRequestId(v string) *FlightCancelOrderV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightCancelOrderV2ResponseBody) SetSuccess(v bool) *FlightCancelOrderV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightCancelOrderV2ResponseBody) SetTraceId(v string) *FlightCancelOrderV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightCancelOrderV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightCancelOrderV2ResponseBodyModule struct {
	// example:
	//
	// 2023-08-10 17:45:32
	CancelTime *string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}

func (s FlightCancelOrderV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderV2ResponseBodyModule) GetCancelTime() *string {
	return s.CancelTime
}

func (s *FlightCancelOrderV2ResponseBodyModule) SetCancelTime(v string) *FlightCancelOrderV2ResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *FlightCancelOrderV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
