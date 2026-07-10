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
	Code      *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                          `json:"message,omitempty" xml:"message,omitempty"`
	Module    *TrainOrderPayResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                            `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                          `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrainOrderPayResponseBodyModule struct {
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	Status     *bool   `json:"status,omitempty" xml:"status,omitempty"`
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
