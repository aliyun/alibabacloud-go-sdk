// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderCancelResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderCancelResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderCancelResponseBodyModule) *TrainOrderCancelResponseBody
	GetModule() *TrainOrderCancelResponseBodyModule
	SetRequestId(v string) *TrainOrderCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderCancelResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderCancelResponseBody
	GetTraceId() *string
}

type TrainOrderCancelResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainOrderCancelResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainOrderCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCancelResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderCancelResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderCancelResponseBody) GetModule() *TrainOrderCancelResponseBodyModule {
	return s.Module
}

func (s *TrainOrderCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderCancelResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderCancelResponseBody) SetCode(v string) *TrainOrderCancelResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderCancelResponseBody) SetMessage(v string) *TrainOrderCancelResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderCancelResponseBody) SetModule(v *TrainOrderCancelResponseBodyModule) *TrainOrderCancelResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderCancelResponseBody) SetRequestId(v string) *TrainOrderCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderCancelResponseBody) SetSuccess(v bool) *TrainOrderCancelResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderCancelResponseBody) SetTraceId(v string) *TrainOrderCancelResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderCancelResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCancelResponseBodyModule struct {
	// example:
	//
	// 11127278782
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// example:
	//
	// 116019444
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 43534345
	OutChangeOrderId *string `json:"out_change_order_id,omitempty" xml:"out_change_order_id,omitempty"`
	// example:
	//
	// 3702553342926024704
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TrainOrderCancelResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCancelResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderCancelResponseBodyModule) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *TrainOrderCancelResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderCancelResponseBodyModule) GetOutChangeOrderId() *string {
	return s.OutChangeOrderId
}

func (s *TrainOrderCancelResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderCancelResponseBodyModule) GetStatus() *bool {
	return s.Status
}

func (s *TrainOrderCancelResponseBodyModule) SetChangeOrderId(v string) *TrainOrderCancelResponseBodyModule {
	s.ChangeOrderId = &v
	return s
}

func (s *TrainOrderCancelResponseBodyModule) SetOrderId(v string) *TrainOrderCancelResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainOrderCancelResponseBodyModule) SetOutChangeOrderId(v string) *TrainOrderCancelResponseBodyModule {
	s.OutChangeOrderId = &v
	return s
}

func (s *TrainOrderCancelResponseBodyModule) SetOutOrderId(v string) *TrainOrderCancelResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderCancelResponseBodyModule) SetStatus(v bool) *TrainOrderCancelResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainOrderCancelResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
