// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceRuleSaveResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceRuleSaveResponseBody
	GetMessage() *string
	SetModule(v *InvoiceRuleSaveResponseBodyModule) *InvoiceRuleSaveResponseBody
	GetModule() *InvoiceRuleSaveResponseBodyModule
	SetRequestId(v string) *InvoiceRuleSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceRuleSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceRuleSaveResponseBody
	GetTraceId() *string
}

type InvoiceRuleSaveResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InvoiceRuleSaveResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 8465F68D-BC97-5C0F-9161-3E65919D9135
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InvoiceRuleSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceRuleSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceRuleSaveResponseBody) GetModule() *InvoiceRuleSaveResponseBodyModule {
	return s.Module
}

func (s *InvoiceRuleSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceRuleSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceRuleSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceRuleSaveResponseBody) SetCode(v string) *InvoiceRuleSaveResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetMessage(v string) *InvoiceRuleSaveResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetModule(v *InvoiceRuleSaveResponseBodyModule) *InvoiceRuleSaveResponseBody {
	s.Module = v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetRequestId(v string) *InvoiceRuleSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetSuccess(v bool) *InvoiceRuleSaveResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) SetTraceId(v string) *InvoiceRuleSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceRuleSaveResponseBody) Validate() error {
	return dara.Validate(s)
}

type InvoiceRuleSaveResponseBodyModule struct {
	// example:
	//
	// 1
	AddNum *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// example:
	//
	// 1
	RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
}

func (s InvoiceRuleSaveResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveResponseBodyModule) GetAddNum() *int32 {
	return s.AddNum
}

func (s *InvoiceRuleSaveResponseBodyModule) GetRemoveNum() *int32 {
	return s.RemoveNum
}

func (s *InvoiceRuleSaveResponseBodyModule) SetAddNum(v int32) *InvoiceRuleSaveResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *InvoiceRuleSaveResponseBodyModule) SetRemoveNum(v int32) *InvoiceRuleSaveResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *InvoiceRuleSaveResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
