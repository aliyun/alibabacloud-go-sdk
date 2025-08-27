// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvoiceRuleDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *InvoiceRuleDeleteResponseBody
	GetMessage() *string
	SetModule(v *InvoiceRuleDeleteResponseBodyModule) *InvoiceRuleDeleteResponseBody
	GetModule() *InvoiceRuleDeleteResponseBodyModule
	SetRequestId(v string) *InvoiceRuleDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvoiceRuleDeleteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InvoiceRuleDeleteResponseBody
	GetTraceId() *string
}

type InvoiceRuleDeleteResponseBody struct {
	// example:
	//
	// 200
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InvoiceRuleDeleteResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s InvoiceRuleDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvoiceRuleDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvoiceRuleDeleteResponseBody) GetModule() *InvoiceRuleDeleteResponseBodyModule {
	return s.Module
}

func (s *InvoiceRuleDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvoiceRuleDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvoiceRuleDeleteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InvoiceRuleDeleteResponseBody) SetCode(v string) *InvoiceRuleDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBody) SetMessage(v string) *InvoiceRuleDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBody) SetModule(v *InvoiceRuleDeleteResponseBodyModule) *InvoiceRuleDeleteResponseBody {
	s.Module = v
	return s
}

func (s *InvoiceRuleDeleteResponseBody) SetRequestId(v string) *InvoiceRuleDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBody) SetSuccess(v bool) *InvoiceRuleDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBody) SetTraceId(v string) *InvoiceRuleDeleteResponseBody {
	s.TraceId = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}

type InvoiceRuleDeleteResponseBodyModule struct {
	// example:
	//
	// 1
	RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	// example:
	//
	// 0
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s InvoiceRuleDeleteResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteResponseBodyModule) GetRemoveNum() *int32 {
	return s.RemoveNum
}

func (s *InvoiceRuleDeleteResponseBodyModule) GetSelectedUserNum() *int32 {
	return s.SelectedUserNum
}

func (s *InvoiceRuleDeleteResponseBodyModule) SetRemoveNum(v int32) *InvoiceRuleDeleteResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBodyModule) SetSelectedUserNum(v int32) *InvoiceRuleDeleteResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

func (s *InvoiceRuleDeleteResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
