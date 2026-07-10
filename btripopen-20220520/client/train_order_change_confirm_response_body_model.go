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
	Code      *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module    *TrainOrderChangeConfirmResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                    `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrainOrderChangeConfirmResponseBodyModule struct {
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId    *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	Status        *bool   `json:"status,omitempty" xml:"status,omitempty"`
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
