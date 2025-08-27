// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupCorpTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GroupCorpTokenResponseBody
	GetCode() *string
	SetMessage(v string) *GroupCorpTokenResponseBody
	GetMessage() *string
	SetModule(v *GroupCorpTokenResponseBodyModule) *GroupCorpTokenResponseBody
	GetModule() *GroupCorpTokenResponseBodyModule
	SetRequestId(v string) *GroupCorpTokenResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GroupCorpTokenResponseBody
	GetSuccess() *string
	SetTraceId(v string) *GroupCorpTokenResponseBody
	GetTraceId() *string
}

type GroupCorpTokenResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *GroupCorpTokenResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GroupCorpTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GroupCorpTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCorpTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GroupCorpTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GroupCorpTokenResponseBody) GetModule() *GroupCorpTokenResponseBodyModule {
	return s.Module
}

func (s *GroupCorpTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GroupCorpTokenResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GroupCorpTokenResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GroupCorpTokenResponseBody) SetCode(v string) *GroupCorpTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GroupCorpTokenResponseBody) SetMessage(v string) *GroupCorpTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GroupCorpTokenResponseBody) SetModule(v *GroupCorpTokenResponseBodyModule) *GroupCorpTokenResponseBody {
	s.Module = v
	return s
}

func (s *GroupCorpTokenResponseBody) SetRequestId(v string) *GroupCorpTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GroupCorpTokenResponseBody) SetSuccess(v string) *GroupCorpTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GroupCorpTokenResponseBody) SetTraceId(v string) *GroupCorpTokenResponseBody {
	s.TraceId = &v
	return s
}

func (s *GroupCorpTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GroupCorpTokenResponseBodyModule struct {
	// example:
	//
	// 70000
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// 1652410740914
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// example:
	//
	// *0*37j76df
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCorpTokenResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GroupCorpTokenResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GroupCorpTokenResponseBodyModule) GetExpire() *int64 {
	return s.Expire
}

func (s *GroupCorpTokenResponseBodyModule) GetStart() *int64 {
	return s.Start
}

func (s *GroupCorpTokenResponseBodyModule) GetToken() *string {
	return s.Token
}

func (s *GroupCorpTokenResponseBodyModule) SetExpire(v int64) *GroupCorpTokenResponseBodyModule {
	s.Expire = &v
	return s
}

func (s *GroupCorpTokenResponseBodyModule) SetStart(v int64) *GroupCorpTokenResponseBodyModule {
	s.Start = &v
	return s
}

func (s *GroupCorpTokenResponseBodyModule) SetToken(v string) *GroupCorpTokenResponseBodyModule {
	s.Token = &v
	return s
}

func (s *GroupCorpTokenResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
