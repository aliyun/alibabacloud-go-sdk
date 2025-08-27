// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceModifyResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceModifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvoiceModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceModifyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceModifyResponseBody
	GetTraceId() *string
}

type InvoiceModifyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s InvoiceModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceModifyResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceModifyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceModifyResponseBody) SetCode(v string) *InvoiceModifyResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetMessage(v string) *InvoiceModifyResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetRequestId(v string) *InvoiceModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetSuccess(v bool) *InvoiceModifyResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceModifyResponseBody) SetTraceId(v string) *InvoiceModifyResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceModifyResponseBody) Validate() error {
	return dara.Validate(s)
}
