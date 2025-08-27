// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderPayResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderPayResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderPayResponseBodyModule) *TrainOrderPayResponseBody
	GetModule() *TrainOrderPayResponseBodyModule
	SetRequestId(v string) *TrainOrderPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderPayResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderPayResponseBody
	GetTraceId() *string
}

type TrainOrderPayResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainOrderPayResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s TrainOrderPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderPayResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderPayResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderPayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderPayResponseBody) GetModule() *TrainOrderPayResponseBodyModule {
	return s.Module
}

func (s *TrainOrderPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderPayResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderPayResponseBody) SetCode(v string) *TrainOrderPayResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderPayResponseBody) SetMessage(v string) *TrainOrderPayResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderPayResponseBody) SetModule(v *TrainOrderPayResponseBodyModule) *TrainOrderPayResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderPayResponseBody) SetRequestId(v string) *TrainOrderPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderPayResponseBody) SetSuccess(v bool) *TrainOrderPayResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderPayResponseBody) SetTraceId(v string) *TrainOrderPayResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderPayResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderPayResponseBodyModule struct {
	// example:
	//
	// 1017124195788186048
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TrainOrderPayResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderPayResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderPayResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderPayResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderPayResponseBodyModule) GetStatus() *bool {
	return s.Status
}

func (s *TrainOrderPayResponseBodyModule) SetOrderId(v string) *TrainOrderPayResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainOrderPayResponseBodyModule) SetOutOrderId(v string) *TrainOrderPayResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderPayResponseBodyModule) SetStatus(v bool) *TrainOrderPayResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainOrderPayResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
