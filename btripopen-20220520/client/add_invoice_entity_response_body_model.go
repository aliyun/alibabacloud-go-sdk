// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInvoiceEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddInvoiceEntityResponseBody
	GetCode() *string
	SetMessage(v string) *AddInvoiceEntityResponseBody
	GetMessage() *string
	SetModule(v *AddInvoiceEntityResponseBodyModule) *AddInvoiceEntityResponseBody
	GetModule() *AddInvoiceEntityResponseBodyModule
	SetRequestId(v string) *AddInvoiceEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddInvoiceEntityResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AddInvoiceEntityResponseBody
	GetTraceId() *string
}

type AddInvoiceEntityResponseBody struct {
	// example:
	//
	// 200
	Code    *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module  *AddInvoiceEntityResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddInvoiceEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddInvoiceEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddInvoiceEntityResponseBody) GetModule() *AddInvoiceEntityResponseBodyModule {
	return s.Module
}

func (s *AddInvoiceEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInvoiceEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddInvoiceEntityResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AddInvoiceEntityResponseBody) SetCode(v string) *AddInvoiceEntityResponseBody {
	s.Code = &v
	return s
}

func (s *AddInvoiceEntityResponseBody) SetMessage(v string) *AddInvoiceEntityResponseBody {
	s.Message = &v
	return s
}

func (s *AddInvoiceEntityResponseBody) SetModule(v *AddInvoiceEntityResponseBodyModule) *AddInvoiceEntityResponseBody {
	s.Module = v
	return s
}

func (s *AddInvoiceEntityResponseBody) SetRequestId(v string) *AddInvoiceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInvoiceEntityResponseBody) SetSuccess(v bool) *AddInvoiceEntityResponseBody {
	s.Success = &v
	return s
}

func (s *AddInvoiceEntityResponseBody) SetTraceId(v string) *AddInvoiceEntityResponseBody {
	s.TraceId = &v
	return s
}

func (s *AddInvoiceEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddInvoiceEntityResponseBodyModule struct {
	// example:
	//
	// 1
	AddNum *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// example:
	//
	// 2
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s AddInvoiceEntityResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityResponseBodyModule) GetAddNum() *int32 {
	return s.AddNum
}

func (s *AddInvoiceEntityResponseBodyModule) GetSelectedUserNum() *int32 {
	return s.SelectedUserNum
}

func (s *AddInvoiceEntityResponseBodyModule) SetAddNum(v int32) *AddInvoiceEntityResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *AddInvoiceEntityResponseBodyModule) SetSelectedUserNum(v int32) *AddInvoiceEntityResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

func (s *AddInvoiceEntityResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
