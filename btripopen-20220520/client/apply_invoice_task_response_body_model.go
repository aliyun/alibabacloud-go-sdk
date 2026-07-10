// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyInvoiceTaskResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyInvoiceTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyInvoiceTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyInvoiceTaskResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyInvoiceTaskResponseBody
	GetTraceId() *string
}

type ApplyInvoiceTaskResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyInvoiceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyInvoiceTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyInvoiceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyInvoiceTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyInvoiceTaskResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyInvoiceTaskResponseBody) SetCode(v string) *ApplyInvoiceTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyInvoiceTaskResponseBody) SetMessage(v string) *ApplyInvoiceTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyInvoiceTaskResponseBody) SetRequestId(v string) *ApplyInvoiceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyInvoiceTaskResponseBody) SetSuccess(v bool) *ApplyInvoiceTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyInvoiceTaskResponseBody) SetTraceId(v string) *ApplyInvoiceTaskResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyInvoiceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
