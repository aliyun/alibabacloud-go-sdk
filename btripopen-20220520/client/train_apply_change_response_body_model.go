// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyChangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainApplyChangeResponseBody
	GetCode() *string
	SetMessage(v string) *TrainApplyChangeResponseBody
	GetMessage() *string
	SetModule(v *TrainApplyChangeResponseBodyModule) *TrainApplyChangeResponseBody
	GetModule() *TrainApplyChangeResponseBodyModule
	SetRequestId(v string) *TrainApplyChangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainApplyChangeResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainApplyChangeResponseBody
	GetTraceId() *string
}

type TrainApplyChangeResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainApplyChangeResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 2103ad0216854277618591626db2b6
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainApplyChangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeResponseBody) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainApplyChangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainApplyChangeResponseBody) GetModule() *TrainApplyChangeResponseBodyModule {
	return s.Module
}

func (s *TrainApplyChangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainApplyChangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainApplyChangeResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainApplyChangeResponseBody) SetCode(v string) *TrainApplyChangeResponseBody {
	s.Code = &v
	return s
}

func (s *TrainApplyChangeResponseBody) SetMessage(v string) *TrainApplyChangeResponseBody {
	s.Message = &v
	return s
}

func (s *TrainApplyChangeResponseBody) SetModule(v *TrainApplyChangeResponseBodyModule) *TrainApplyChangeResponseBody {
	s.Module = v
	return s
}

func (s *TrainApplyChangeResponseBody) SetRequestId(v string) *TrainApplyChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainApplyChangeResponseBody) SetSuccess(v bool) *TrainApplyChangeResponseBody {
	s.Success = &v
	return s
}

func (s *TrainApplyChangeResponseBody) SetTraceId(v string) *TrainApplyChangeResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainApplyChangeResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainApplyChangeResponseBodyModule struct {
	// example:
	//
	// 1234
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// example:
	//
	// 116019444
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017124195788186048
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1
	PayStatus *string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

func (s TrainApplyChangeResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeResponseBodyModule) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *TrainApplyChangeResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainApplyChangeResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainApplyChangeResponseBodyModule) GetPayStatus() *string {
	return s.PayStatus
}

func (s *TrainApplyChangeResponseBodyModule) SetChangeOrderId(v string) *TrainApplyChangeResponseBodyModule {
	s.ChangeOrderId = &v
	return s
}

func (s *TrainApplyChangeResponseBodyModule) SetOrderId(v string) *TrainApplyChangeResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainApplyChangeResponseBodyModule) SetOutOrderId(v string) *TrainApplyChangeResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainApplyChangeResponseBodyModule) SetPayStatus(v string) *TrainApplyChangeResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *TrainApplyChangeResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
