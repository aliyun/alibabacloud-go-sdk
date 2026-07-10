// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceSearchResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceSearchResponseBody
	GetMessage() *string
	SetModule(v []*InvoiceSearchResponseBodyModule) *InvoiceSearchResponseBody
	GetModule() []*InvoiceSearchResponseBodyModule
	SetRequestId(v string) *InvoiceSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceSearchResponseBody
	GetTraceId() *string
}

type InvoiceSearchResponseBody struct {
	Code      *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*InvoiceSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceSearchResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceSearchResponseBody) GetModule() []*InvoiceSearchResponseBodyModule {
	return s.Module
}

func (s *InvoiceSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceSearchResponseBody) SetCode(v string) *InvoiceSearchResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetMessage(v string) *InvoiceSearchResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetModule(v []*InvoiceSearchResponseBodyModule) *InvoiceSearchResponseBody {
	s.Module = v
	return s
}

func (s *InvoiceSearchResponseBody) SetRequestId(v string) *InvoiceSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetSuccess(v bool) *InvoiceSearchResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceSearchResponseBody) SetTraceId(v string) *InvoiceSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceSearchResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InvoiceSearchResponseBodyModule struct {
	Id                 *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s InvoiceSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InvoiceSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InvoiceSearchResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *InvoiceSearchResponseBodyModule) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *InvoiceSearchResponseBodyModule) GetTitle() *string {
	return s.Title
}

func (s *InvoiceSearchResponseBodyModule) SetId(v int64) *InvoiceSearchResponseBodyModule {
	s.Id = &v
	return s
}

func (s *InvoiceSearchResponseBodyModule) SetThirdPartInvoiceId(v string) *InvoiceSearchResponseBodyModule {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *InvoiceSearchResponseBodyModule) SetTitle(v string) *InvoiceSearchResponseBodyModule {
	s.Title = &v
	return s
}

func (s *InvoiceSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
