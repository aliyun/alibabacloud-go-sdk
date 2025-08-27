// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderChangeConfirmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderChangeConfirmResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderChangeConfirmResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderChangeConfirmResponseBodyModule) *TrainOrderChangeConfirmResponseBody
	GetModule() *TrainOrderChangeConfirmResponseBodyModule
	SetRequestId(v string) *TrainOrderChangeConfirmResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderChangeConfirmResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderChangeConfirmResponseBody
	GetTraceId() *string
}

type TrainOrderChangeConfirmResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainOrderChangeConfirmResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s TrainOrderChangeConfirmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderChangeConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderChangeConfirmResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderChangeConfirmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderChangeConfirmResponseBody) GetModule() *TrainOrderChangeConfirmResponseBodyModule {
	return s.Module
}

func (s *TrainOrderChangeConfirmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderChangeConfirmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderChangeConfirmResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderChangeConfirmResponseBody) SetCode(v string) *TrainOrderChangeConfirmResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBody) SetMessage(v string) *TrainOrderChangeConfirmResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBody) SetModule(v *TrainOrderChangeConfirmResponseBodyModule) *TrainOrderChangeConfirmResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderChangeConfirmResponseBody) SetRequestId(v string) *TrainOrderChangeConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBody) SetSuccess(v bool) *TrainOrderChangeConfirmResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBody) SetTraceId(v string) *TrainOrderChangeConfirmResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderChangeConfirmResponseBodyModule struct {
	// example:
	//
	// 123445443444
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// example:
	//
	// 116019444
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1234232
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 0
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TrainOrderChangeConfirmResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderChangeConfirmResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderChangeConfirmResponseBodyModule) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *TrainOrderChangeConfirmResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderChangeConfirmResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderChangeConfirmResponseBodyModule) GetStatus() *bool {
	return s.Status
}

func (s *TrainOrderChangeConfirmResponseBodyModule) SetChangeOrderId(v string) *TrainOrderChangeConfirmResponseBodyModule {
	s.ChangeOrderId = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBodyModule) SetOrderId(v string) *TrainOrderChangeConfirmResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBodyModule) SetOutOrderId(v string) *TrainOrderChangeConfirmResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBodyModule) SetStatus(v bool) *TrainOrderChangeConfirmResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainOrderChangeConfirmResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
