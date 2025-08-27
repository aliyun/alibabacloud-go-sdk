// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTBAccountUnbindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TBAccountUnbindResponseBody
	GetCode() *string
	SetMessage(v string) *TBAccountUnbindResponseBody
	GetMessage() *string
	SetModule(v bool) *TBAccountUnbindResponseBody
	GetModule() *bool
	SetRequestId(v string) *TBAccountUnbindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TBAccountUnbindResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TBAccountUnbindResponseBody
	GetTraceId() *string
}

type TBAccountUnbindResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// 456456575656757
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TBAccountUnbindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TBAccountUnbindResponseBody) GoString() string {
	return s.String()
}

func (s *TBAccountUnbindResponseBody) GetCode() *string {
	return s.Code
}

func (s *TBAccountUnbindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TBAccountUnbindResponseBody) GetModule() *bool {
	return s.Module
}

func (s *TBAccountUnbindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TBAccountUnbindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TBAccountUnbindResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TBAccountUnbindResponseBody) SetCode(v string) *TBAccountUnbindResponseBody {
	s.Code = &v
	return s
}

func (s *TBAccountUnbindResponseBody) SetMessage(v string) *TBAccountUnbindResponseBody {
	s.Message = &v
	return s
}

func (s *TBAccountUnbindResponseBody) SetModule(v bool) *TBAccountUnbindResponseBody {
	s.Module = &v
	return s
}

func (s *TBAccountUnbindResponseBody) SetRequestId(v string) *TBAccountUnbindResponseBody {
	s.RequestId = &v
	return s
}

func (s *TBAccountUnbindResponseBody) SetSuccess(v bool) *TBAccountUnbindResponseBody {
	s.Success = &v
	return s
}

func (s *TBAccountUnbindResponseBody) SetTraceId(v string) *TBAccountUnbindResponseBody {
	s.TraceId = &v
	return s
}

func (s *TBAccountUnbindResponseBody) Validate() error {
	return dara.Validate(s)
}
