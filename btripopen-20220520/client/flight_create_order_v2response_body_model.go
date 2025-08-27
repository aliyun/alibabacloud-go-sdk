// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightCreateOrderV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightCreateOrderV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightCreateOrderV2ResponseBodyModule) *FlightCreateOrderV2ResponseBody
	GetModule() *FlightCreateOrderV2ResponseBodyModule
	SetRequestId(v string) *FlightCreateOrderV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightCreateOrderV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightCreateOrderV2ResponseBody
	GetTraceId() *string
}

type FlightCreateOrderV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightCreateOrderV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 210bc60a16916593445203790d2a16
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 2103ad0716827336456723986d4bda
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightCreateOrderV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightCreateOrderV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightCreateOrderV2ResponseBody) GetModule() *FlightCreateOrderV2ResponseBodyModule {
	return s.Module
}

func (s *FlightCreateOrderV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightCreateOrderV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightCreateOrderV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightCreateOrderV2ResponseBody) SetCode(v string) *FlightCreateOrderV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBody) SetMessage(v string) *FlightCreateOrderV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBody) SetModule(v *FlightCreateOrderV2ResponseBodyModule) *FlightCreateOrderV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightCreateOrderV2ResponseBody) SetRequestId(v string) *FlightCreateOrderV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBody) SetSuccess(v bool) *FlightCreateOrderV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBody) SetTraceId(v string) *FlightCreateOrderV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightCreateOrderV2ResponseBodyModule struct {
	AsyncCreateOrderKey *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	// example:
	//
	// 2023-08-10 17:42:32
	LatestPayTime *string `json:"latest_pay_time,omitempty" xml:"latest_pay_time,omitempty"`
	// example:
	//
	// 1017002195798359369
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// cheshiapi002kwl
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 32
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 51000
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s FlightCreateOrderV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2ResponseBodyModule) GetAsyncCreateOrderKey() *string {
	return s.AsyncCreateOrderKey
}

func (s *FlightCreateOrderV2ResponseBodyModule) GetLatestPayTime() *string {
	return s.LatestPayTime
}

func (s *FlightCreateOrderV2ResponseBodyModule) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightCreateOrderV2ResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightCreateOrderV2ResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightCreateOrderV2ResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *FlightCreateOrderV2ResponseBodyModule) SetAsyncCreateOrderKey(v string) *FlightCreateOrderV2ResponseBodyModule {
	s.AsyncCreateOrderKey = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBodyModule) SetLatestPayTime(v string) *FlightCreateOrderV2ResponseBodyModule {
	s.LatestPayTime = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBodyModule) SetOrderId(v int64) *FlightCreateOrderV2ResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBodyModule) SetOutOrderId(v string) *FlightCreateOrderV2ResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBodyModule) SetStatus(v int32) *FlightCreateOrderV2ResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBodyModule) SetTotalPrice(v int64) *FlightCreateOrderV2ResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *FlightCreateOrderV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
