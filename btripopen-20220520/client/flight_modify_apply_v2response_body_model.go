// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyApplyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightModifyApplyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightModifyApplyV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightModifyApplyV2ResponseBodyModule) *FlightModifyApplyV2ResponseBody
	GetModule() *FlightModifyApplyV2ResponseBodyModule
	SetRequestId(v string) *FlightModifyApplyV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightModifyApplyV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightModifyApplyV2ResponseBody
	GetTraceId() *string
}

type FlightModifyApplyV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightModifyApplyV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 212a8b8216915622178333839e665d
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightModifyApplyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyApplyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightModifyApplyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightModifyApplyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightModifyApplyV2ResponseBody) GetModule() *FlightModifyApplyV2ResponseBodyModule {
	return s.Module
}

func (s *FlightModifyApplyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightModifyApplyV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightModifyApplyV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightModifyApplyV2ResponseBody) SetCode(v string) *FlightModifyApplyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBody) SetMessage(v string) *FlightModifyApplyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBody) SetModule(v *FlightModifyApplyV2ResponseBodyModule) *FlightModifyApplyV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightModifyApplyV2ResponseBody) SetRequestId(v string) *FlightModifyApplyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBody) SetSuccess(v bool) *FlightModifyApplyV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBody) SetTraceId(v string) *FlightModifyApplyV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightModifyApplyV2ResponseBodyModule struct {
	// example:
	//
	// 1017002195370467138
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195786853020
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// example:
	//
	// 1019195786853020
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s FlightModifyApplyV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyApplyV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightModifyApplyV2ResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightModifyApplyV2ResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyApplyV2ResponseBodyModule) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyApplyV2ResponseBodyModule) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *FlightModifyApplyV2ResponseBodyModule) SetOrderId(v string) *FlightModifyApplyV2ResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBodyModule) SetOutOrderId(v string) *FlightModifyApplyV2ResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBodyModule) SetOutSubOrderId(v string) *FlightModifyApplyV2ResponseBodyModule {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBodyModule) SetSubOrderId(v string) *FlightModifyApplyV2ResponseBodyModule {
	s.SubOrderId = &v
	return s
}

func (s *FlightModifyApplyV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
