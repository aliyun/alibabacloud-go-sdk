// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightRefundApplyResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightRefundApplyResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightRefundApplyResponseBodyModule) *IntlFlightRefundApplyResponseBody
	GetModule() *IntlFlightRefundApplyResponseBodyModule
	SetRequestId(v string) *IntlFlightRefundApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightRefundApplyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightRefundApplyResponseBody
	GetTraceId() *string
}

type IntlFlightRefundApplyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightRefundApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightRefundApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightRefundApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightRefundApplyResponseBody) GetModule() *IntlFlightRefundApplyResponseBodyModule {
	return s.Module
}

func (s *IntlFlightRefundApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightRefundApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightRefundApplyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightRefundApplyResponseBody) SetCode(v string) *IntlFlightRefundApplyResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBody) SetMessage(v string) *IntlFlightRefundApplyResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBody) SetModule(v *IntlFlightRefundApplyResponseBodyModule) *IntlFlightRefundApplyResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightRefundApplyResponseBody) SetRequestId(v string) *IntlFlightRefundApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBody) SetSuccess(v bool) *IntlFlightRefundApplyResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBody) SetTraceId(v string) *IntlFlightRefundApplyResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightRefundApplyResponseBodyModule struct {
	// example:
	//
	// 2025011317110900006
	OutRefundApplyId *string `json:"out_refund_apply_id,omitempty" xml:"out_refund_apply_id,omitempty"`
	// example:
	//
	// 1000000005186043
	RefundApplyId *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
}

func (s IntlFlightRefundApplyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyResponseBodyModule) GetOutRefundApplyId() *string {
	return s.OutRefundApplyId
}

func (s *IntlFlightRefundApplyResponseBodyModule) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *IntlFlightRefundApplyResponseBodyModule) SetOutRefundApplyId(v string) *IntlFlightRefundApplyResponseBodyModule {
	s.OutRefundApplyId = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBodyModule) SetRefundApplyId(v string) *IntlFlightRefundApplyResponseBodyModule {
	s.RefundApplyId = &v
	return s
}

func (s *IntlFlightRefundApplyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
