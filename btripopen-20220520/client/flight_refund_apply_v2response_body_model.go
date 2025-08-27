// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightRefundApplyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightRefundApplyV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightRefundApplyV2ResponseBodyModule) *FlightRefundApplyV2ResponseBody
	GetModule() *FlightRefundApplyV2ResponseBodyModule
	SetRequestId(v string) *FlightRefundApplyV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightRefundApplyV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightRefundApplyV2ResponseBody
	GetTraceId() *string
}

type FlightRefundApplyV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightRefundApplyV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s FlightRefundApplyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightRefundApplyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightRefundApplyV2ResponseBody) GetModule() *FlightRefundApplyV2ResponseBodyModule {
	return s.Module
}

func (s *FlightRefundApplyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightRefundApplyV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightRefundApplyV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightRefundApplyV2ResponseBody) SetCode(v string) *FlightRefundApplyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBody) SetMessage(v string) *FlightRefundApplyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBody) SetModule(v *FlightRefundApplyV2ResponseBodyModule) *FlightRefundApplyV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightRefundApplyV2ResponseBody) SetRequestId(v string) *FlightRefundApplyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBody) SetSuccess(v bool) *FlightRefundApplyV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBody) SetTraceId(v string) *FlightRefundApplyV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightRefundApplyV2ResponseBodyModule struct {
	// example:
	//
	// 1683901850297448200
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467200
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195836916039
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// example:
	//
	// 1000000000297003
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s FlightRefundApplyV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2ResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightRefundApplyV2ResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundApplyV2ResponseBodyModule) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightRefundApplyV2ResponseBodyModule) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *FlightRefundApplyV2ResponseBodyModule) SetOrderId(v string) *FlightRefundApplyV2ResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBodyModule) SetOutOrderId(v string) *FlightRefundApplyV2ResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBodyModule) SetOutSubOrderId(v string) *FlightRefundApplyV2ResponseBodyModule {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBodyModule) SetSubOrderId(v string) *FlightRefundApplyV2ResponseBodyModule {
	s.SubOrderId = &v
	return s
}

func (s *FlightRefundApplyV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
