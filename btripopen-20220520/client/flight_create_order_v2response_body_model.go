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
	Code      *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module    *FlightCreateOrderV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlightCreateOrderV2ResponseBodyModule struct {
	AsyncCreateOrderKey *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	LatestPayTime       *string `json:"latest_pay_time,omitempty" xml:"latest_pay_time,omitempty"`
	OrderId             *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId          *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	Status              *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TotalPrice          *int64  `json:"total_price,omitempty" xml:"total_price,omitempty"`
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
