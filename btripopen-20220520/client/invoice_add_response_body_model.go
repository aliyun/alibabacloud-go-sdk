// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceAddResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceAddResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvoiceAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceAddResponseBody
	GetTraceId() *string
}

type InvoiceAddResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceAddResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceAddResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceAddResponseBody) SetCode(v string) *InvoiceAddResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceAddResponseBody) SetMessage(v string) *InvoiceAddResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceAddResponseBody) SetRequestId(v string) *InvoiceAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceAddResponseBody) SetSuccess(v bool) *InvoiceAddResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceAddResponseBody) SetTraceId(v string) *InvoiceAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceAddResponseBody) Validate() error {
	return dara.Validate(s)
}
