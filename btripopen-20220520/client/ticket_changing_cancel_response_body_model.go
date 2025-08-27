// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TicketChangingCancelResponseBody
	GetCode() *string
	SetMessage(v string) *TicketChangingCancelResponseBody
	GetMessage() *string
	SetModule(v *TicketChangingCancelResponseBodyModule) *TicketChangingCancelResponseBody
	GetModule() *TicketChangingCancelResponseBodyModule
	SetRequestId(v string) *TicketChangingCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketChangingCancelResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TicketChangingCancelResponseBody
	GetTraceId() *string
}

type TicketChangingCancelResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TicketChangingCancelResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TicketChangingCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingCancelResponseBody) GoString() string {
	return s.String()
}

func (s *TicketChangingCancelResponseBody) GetCode() *string {
	return s.Code
}

func (s *TicketChangingCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TicketChangingCancelResponseBody) GetModule() *TicketChangingCancelResponseBodyModule {
	return s.Module
}

func (s *TicketChangingCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketChangingCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketChangingCancelResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TicketChangingCancelResponseBody) SetCode(v string) *TicketChangingCancelResponseBody {
	s.Code = &v
	return s
}

func (s *TicketChangingCancelResponseBody) SetMessage(v string) *TicketChangingCancelResponseBody {
	s.Message = &v
	return s
}

func (s *TicketChangingCancelResponseBody) SetModule(v *TicketChangingCancelResponseBodyModule) *TicketChangingCancelResponseBody {
	s.Module = v
	return s
}

func (s *TicketChangingCancelResponseBody) SetRequestId(v string) *TicketChangingCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketChangingCancelResponseBody) SetSuccess(v bool) *TicketChangingCancelResponseBody {
	s.Success = &v
	return s
}

func (s *TicketChangingCancelResponseBody) SetTraceId(v string) *TicketChangingCancelResponseBody {
	s.TraceId = &v
	return s
}

func (s *TicketChangingCancelResponseBody) Validate() error {
	return dara.Validate(s)
}

type TicketChangingCancelResponseBodyModule struct {
	// example:
	//
	// 0000-00-00 00:00:00
	CancelTime *string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// example:
	//
	// mid1243
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TicketChangingCancelResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingCancelResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TicketChangingCancelResponseBodyModule) GetCancelTime() *string {
	return s.CancelTime
}

func (s *TicketChangingCancelResponseBodyModule) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingCancelResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *TicketChangingCancelResponseBodyModule) SetCancelTime(v string) *TicketChangingCancelResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *TicketChangingCancelResponseBodyModule) SetDisSubOrderId(v string) *TicketChangingCancelResponseBodyModule {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingCancelResponseBodyModule) SetStatus(v string) *TicketChangingCancelResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TicketChangingCancelResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
