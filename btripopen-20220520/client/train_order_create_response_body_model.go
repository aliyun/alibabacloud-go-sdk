// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderCreateResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderCreateResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderCreateResponseBodyModule) *TrainOrderCreateResponseBody
	GetModule() *TrainOrderCreateResponseBodyModule
	SetRequestId(v string) *TrainOrderCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderCreateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderCreateResponseBody
	GetTraceId() *string
}

type TrainOrderCreateResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainOrderCreateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s TrainOrderCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderCreateResponseBody) GetModule() *TrainOrderCreateResponseBodyModule {
	return s.Module
}

func (s *TrainOrderCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderCreateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderCreateResponseBody) SetCode(v string) *TrainOrderCreateResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderCreateResponseBody) SetMessage(v string) *TrainOrderCreateResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderCreateResponseBody) SetModule(v *TrainOrderCreateResponseBodyModule) *TrainOrderCreateResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderCreateResponseBody) SetRequestId(v string) *TrainOrderCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderCreateResponseBody) SetSuccess(v bool) *TrainOrderCreateResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderCreateResponseBody) SetTraceId(v string) *TrainOrderCreateResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderCreateResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateResponseBodyModule struct {
	// example:
	//
	// 1017002195798359369
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1233333
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 9
	PayStatus *string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 5
	ServiceFee *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
}

func (s TrainOrderCreateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderCreateResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderCreateResponseBodyModule) GetPayStatus() *string {
	return s.PayStatus
}

func (s *TrainOrderCreateResponseBodyModule) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *TrainOrderCreateResponseBodyModule) SetOrderId(v string) *TrainOrderCreateResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainOrderCreateResponseBodyModule) SetOutOrderId(v string) *TrainOrderCreateResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderCreateResponseBodyModule) SetPayStatus(v string) *TrainOrderCreateResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *TrainOrderCreateResponseBodyModule) SetServiceFee(v int64) *TrainOrderCreateResponseBodyModule {
	s.ServiceFee = &v
	return s
}

func (s *TrainOrderCreateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
