// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightPayOrderV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightPayOrderV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightPayOrderV2ResponseBodyModule) *FlightPayOrderV2ResponseBody
	GetModule() *FlightPayOrderV2ResponseBodyModule
	SetRequestId(v string) *FlightPayOrderV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightPayOrderV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightPayOrderV2ResponseBody
	GetTraceId() *string
}

type FlightPayOrderV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightPayOrderV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s FlightPayOrderV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightPayOrderV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightPayOrderV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightPayOrderV2ResponseBody) GetModule() *FlightPayOrderV2ResponseBodyModule {
	return s.Module
}

func (s *FlightPayOrderV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightPayOrderV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightPayOrderV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightPayOrderV2ResponseBody) SetCode(v string) *FlightPayOrderV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightPayOrderV2ResponseBody) SetMessage(v string) *FlightPayOrderV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightPayOrderV2ResponseBody) SetModule(v *FlightPayOrderV2ResponseBodyModule) *FlightPayOrderV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightPayOrderV2ResponseBody) SetRequestId(v string) *FlightPayOrderV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightPayOrderV2ResponseBody) SetSuccess(v bool) *FlightPayOrderV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightPayOrderV2ResponseBody) SetTraceId(v string) *FlightPayOrderV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightPayOrderV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightPayOrderV2ResponseBodyModule struct {
	// example:
	//
	// 51000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
}

func (s FlightPayOrderV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightPayOrderV2ResponseBodyModule) GetPrice() *int64 {
	return s.Price
}

func (s *FlightPayOrderV2ResponseBodyModule) SetPrice(v int64) *FlightPayOrderV2ResponseBodyModule {
	s.Price = &v
	return s
}

func (s *FlightPayOrderV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
