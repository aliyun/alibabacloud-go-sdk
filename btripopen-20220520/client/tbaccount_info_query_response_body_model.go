// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTBAccountInfoQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TBAccountInfoQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TBAccountInfoQueryResponseBody
	GetMessage() *string
	SetModule(v *TBAccountInfoQueryResponseBodyModule) *TBAccountInfoQueryResponseBody
	GetModule() *TBAccountInfoQueryResponseBodyModule
	SetRequestId(v string) *TBAccountInfoQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TBAccountInfoQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TBAccountInfoQueryResponseBody
	GetTraceId() *string
}

type TBAccountInfoQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TBAccountInfoQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 456456575656757
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f07f316603757445272547d959f
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TBAccountInfoQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TBAccountInfoQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TBAccountInfoQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TBAccountInfoQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TBAccountInfoQueryResponseBody) GetModule() *TBAccountInfoQueryResponseBodyModule {
	return s.Module
}

func (s *TBAccountInfoQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TBAccountInfoQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TBAccountInfoQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TBAccountInfoQueryResponseBody) SetCode(v string) *TBAccountInfoQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TBAccountInfoQueryResponseBody) SetMessage(v string) *TBAccountInfoQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TBAccountInfoQueryResponseBody) SetModule(v *TBAccountInfoQueryResponseBodyModule) *TBAccountInfoQueryResponseBody {
	s.Module = v
	return s
}

func (s *TBAccountInfoQueryResponseBody) SetRequestId(v string) *TBAccountInfoQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TBAccountInfoQueryResponseBody) SetSuccess(v bool) *TBAccountInfoQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TBAccountInfoQueryResponseBody) SetTraceId(v string) *TBAccountInfoQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TBAccountInfoQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TBAccountInfoQueryResponseBodyModule struct {
	TbAccount *string `json:"tb_account,omitempty" xml:"tb_account,omitempty"`
	// example:
	//
	// true
	TbBond *bool `json:"tb_bond,omitempty" xml:"tb_bond,omitempty"`
}

func (s TBAccountInfoQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TBAccountInfoQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TBAccountInfoQueryResponseBodyModule) GetTbAccount() *string {
	return s.TbAccount
}

func (s *TBAccountInfoQueryResponseBodyModule) GetTbBond() *bool {
	return s.TbBond
}

func (s *TBAccountInfoQueryResponseBodyModule) SetTbAccount(v string) *TBAccountInfoQueryResponseBodyModule {
	s.TbAccount = &v
	return s
}

func (s *TBAccountInfoQueryResponseBodyModule) SetTbBond(v bool) *TBAccountInfoQueryResponseBodyModule {
	s.TbBond = &v
	return s
}

func (s *TBAccountInfoQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
