// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainApplyRefundResponseBody
	GetCode() *string
	SetMessage(v string) *TrainApplyRefundResponseBody
	GetMessage() *string
	SetModule(v *TrainApplyRefundResponseBodyModule) *TrainApplyRefundResponseBody
	GetModule() *TrainApplyRefundResponseBodyModule
	SetRequestId(v string) *TrainApplyRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainApplyRefundResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainApplyRefundResponseBody
	GetTraceId() *string
}

type TrainApplyRefundResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainApplyRefundResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s TrainApplyRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundResponseBody) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainApplyRefundResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainApplyRefundResponseBody) GetModule() *TrainApplyRefundResponseBodyModule {
	return s.Module
}

func (s *TrainApplyRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainApplyRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainApplyRefundResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainApplyRefundResponseBody) SetCode(v string) *TrainApplyRefundResponseBody {
	s.Code = &v
	return s
}

func (s *TrainApplyRefundResponseBody) SetMessage(v string) *TrainApplyRefundResponseBody {
	s.Message = &v
	return s
}

func (s *TrainApplyRefundResponseBody) SetModule(v *TrainApplyRefundResponseBodyModule) *TrainApplyRefundResponseBody {
	s.Module = v
	return s
}

func (s *TrainApplyRefundResponseBody) SetRequestId(v string) *TrainApplyRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainApplyRefundResponseBody) SetSuccess(v bool) *TrainApplyRefundResponseBody {
	s.Success = &v
	return s
}

func (s *TrainApplyRefundResponseBody) SetTraceId(v string) *TrainApplyRefundResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainApplyRefundResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainApplyRefundResponseBodyModule struct {
	// example:
	//
	// 116019444
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467200
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s TrainApplyRefundResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainApplyRefundResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainApplyRefundResponseBodyModule) SetOrderId(v string) *TrainApplyRefundResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainApplyRefundResponseBodyModule) SetOutOrderId(v string) *TrainApplyRefundResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainApplyRefundResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
