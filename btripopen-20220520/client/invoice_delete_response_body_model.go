// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvoiceDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceDeleteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceDeleteResponseBody
	GetTraceId() *string
}

type InvoiceDeleteResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
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

func (s InvoiceDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceDeleteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceDeleteResponseBody) SetCode(v string) *InvoiceDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetMessage(v string) *InvoiceDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetRequestId(v string) *InvoiceDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetSuccess(v bool) *InvoiceDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceDeleteResponseBody) SetTraceId(v string) *InvoiceDeleteResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
