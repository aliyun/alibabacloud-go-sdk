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
	Code      *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module    *FlightModifyApplyV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlightModifyApplyV2ResponseBodyModule struct {
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId    *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	SubOrderId    *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
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
