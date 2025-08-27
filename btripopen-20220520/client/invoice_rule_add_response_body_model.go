// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceRuleAddResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceRuleAddResponseBody
	GetMessage() *string
	SetModule(v *InvoiceRuleAddResponseBodyModule) *InvoiceRuleAddResponseBody
	GetModule() *InvoiceRuleAddResponseBodyModule
	SetRequestId(v string) *InvoiceRuleAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceRuleAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceRuleAddResponseBody
	GetTraceId() *string
}

type InvoiceRuleAddResponseBody struct {
	// example:
	//
	// 200
	Code    *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InvoiceRuleAddResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s InvoiceRuleAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceRuleAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceRuleAddResponseBody) GetModule() *InvoiceRuleAddResponseBodyModule {
	return s.Module
}

func (s *InvoiceRuleAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceRuleAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceRuleAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceRuleAddResponseBody) SetCode(v string) *InvoiceRuleAddResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceRuleAddResponseBody) SetMessage(v string) *InvoiceRuleAddResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceRuleAddResponseBody) SetModule(v *InvoiceRuleAddResponseBodyModule) *InvoiceRuleAddResponseBody {
	s.Module = v
	return s
}

func (s *InvoiceRuleAddResponseBody) SetRequestId(v string) *InvoiceRuleAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceRuleAddResponseBody) SetSuccess(v bool) *InvoiceRuleAddResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceRuleAddResponseBody) SetTraceId(v string) *InvoiceRuleAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceRuleAddResponseBody) Validate() error {
	return dara.Validate(s)
}

type InvoiceRuleAddResponseBodyModule struct {
	// example:
	//
	// 1
	AddNum *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// example:
	//
	// 1
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s InvoiceRuleAddResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddResponseBodyModule) GetAddNum() *int32 {
	return s.AddNum
}

func (s *InvoiceRuleAddResponseBodyModule) GetSelectedUserNum() *int32 {
	return s.SelectedUserNum
}

func (s *InvoiceRuleAddResponseBodyModule) SetAddNum(v int32) *InvoiceRuleAddResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *InvoiceRuleAddResponseBodyModule) SetSelectedUserNum(v int32) *InvoiceRuleAddResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

func (s *InvoiceRuleAddResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
