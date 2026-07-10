// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvoiceEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInvoiceEntityResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteInvoiceEntityResponseBody
	GetMessage() *string
	SetModule(v *DeleteInvoiceEntityResponseBodyModule) *DeleteInvoiceEntityResponseBody
	GetModule() *DeleteInvoiceEntityResponseBodyModule
	SetRequestId(v string) *DeleteInvoiceEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInvoiceEntityResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteInvoiceEntityResponseBody
	GetTraceId() *string
}

type DeleteInvoiceEntityResponseBody struct {
	Code      *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module    *DeleteInvoiceEntityResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteInvoiceEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInvoiceEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInvoiceEntityResponseBody) GetModule() *DeleteInvoiceEntityResponseBodyModule {
	return s.Module
}

func (s *DeleteInvoiceEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInvoiceEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInvoiceEntityResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteInvoiceEntityResponseBody) SetCode(v string) *DeleteInvoiceEntityResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBody) SetMessage(v string) *DeleteInvoiceEntityResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBody) SetModule(v *DeleteInvoiceEntityResponseBodyModule) *DeleteInvoiceEntityResponseBody {
	s.Module = v
	return s
}

func (s *DeleteInvoiceEntityResponseBody) SetRequestId(v string) *DeleteInvoiceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBody) SetSuccess(v bool) *DeleteInvoiceEntityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBody) SetTraceId(v string) *DeleteInvoiceEntityResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteInvoiceEntityResponseBodyModule struct {
	RemoveNum       *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	SelectedUserNum *int32 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

func (s DeleteInvoiceEntityResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityResponseBodyModule) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityResponseBodyModule) GetRemoveNum() *int32 {
	return s.RemoveNum
}

func (s *DeleteInvoiceEntityResponseBodyModule) GetSelectedUserNum() *int32 {
	return s.SelectedUserNum
}

func (s *DeleteInvoiceEntityResponseBodyModule) SetRemoveNum(v int32) *DeleteInvoiceEntityResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBodyModule) SetSelectedUserNum(v int32) *DeleteInvoiceEntityResponseBodyModule {
	s.SelectedUserNum = &v
	return s
}

func (s *DeleteInvoiceEntityResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
