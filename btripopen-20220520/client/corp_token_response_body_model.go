// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCorpTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CorpTokenResponseBody
	GetCode() *string
	SetData(v *CorpTokenResponseBodyData) *CorpTokenResponseBody
	GetData() *CorpTokenResponseBodyData
	SetMessage(v string) *CorpTokenResponseBody
	GetMessage() *string
	SetModule(v *CorpTokenResponseBodyModule) *CorpTokenResponseBody
	GetModule() *CorpTokenResponseBodyModule
	SetRequestId(v string) *CorpTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CorpTokenResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CorpTokenResponseBody
	GetTraceId() *string
}

type CorpTokenResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CorpTokenResponseBodyData   `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CorpTokenResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否成功
	//
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

func (s CorpTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CorpTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CorpTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CorpTokenResponseBody) GetData() *CorpTokenResponseBodyData {
	return s.Data
}

func (s *CorpTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CorpTokenResponseBody) GetModule() *CorpTokenResponseBodyModule {
	return s.Module
}

func (s *CorpTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CorpTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CorpTokenResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CorpTokenResponseBody) SetCode(v string) *CorpTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CorpTokenResponseBody) SetData(v *CorpTokenResponseBodyData) *CorpTokenResponseBody {
	s.Data = v
	return s
}

func (s *CorpTokenResponseBody) SetMessage(v string) *CorpTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CorpTokenResponseBody) SetModule(v *CorpTokenResponseBodyModule) *CorpTokenResponseBody {
	s.Module = v
	return s
}

func (s *CorpTokenResponseBody) SetRequestId(v string) *CorpTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CorpTokenResponseBody) SetSuccess(v bool) *CorpTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CorpTokenResponseBody) SetTraceId(v string) *CorpTokenResponseBody {
	s.TraceId = &v
	return s
}

func (s *CorpTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type CorpTokenResponseBodyData struct {
	// example:
	//
	// 70000
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// 37j76df
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s CorpTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CorpTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CorpTokenResponseBodyData) GetExpire() *int64 {
	return s.Expire
}

func (s *CorpTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CorpTokenResponseBodyData) SetExpire(v int64) *CorpTokenResponseBodyData {
	s.Expire = &v
	return s
}

func (s *CorpTokenResponseBodyData) SetToken(v string) *CorpTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *CorpTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CorpTokenResponseBodyModule struct {
	// example:
	//
	// 70000
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// 1635744378301
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// example:
	//
	// 37j76df
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s CorpTokenResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CorpTokenResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CorpTokenResponseBodyModule) GetExpire() *int64 {
	return s.Expire
}

func (s *CorpTokenResponseBodyModule) GetStart() *int64 {
	return s.Start
}

func (s *CorpTokenResponseBodyModule) GetToken() *string {
	return s.Token
}

func (s *CorpTokenResponseBodyModule) SetExpire(v int64) *CorpTokenResponseBodyModule {
	s.Expire = &v
	return s
}

func (s *CorpTokenResponseBodyModule) SetStart(v int64) *CorpTokenResponseBodyModule {
	s.Start = &v
	return s
}

func (s *CorpTokenResponseBodyModule) SetToken(v string) *CorpTokenResponseBodyModule {
	s.Token = &v
	return s
}

func (s *CorpTokenResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
