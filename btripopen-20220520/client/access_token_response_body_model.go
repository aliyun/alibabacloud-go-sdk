// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccessTokenResponseBody
	GetCode() *string
	SetData(v *AccessTokenResponseBodyData) *AccessTokenResponseBody
	GetData() *AccessTokenResponseBodyData
	SetMessage(v string) *AccessTokenResponseBody
	GetMessage() *string
	SetModule(v *AccessTokenResponseBodyModule) *AccessTokenResponseBody
	GetModule() *AccessTokenResponseBodyModule
	SetRequestId(v string) *AccessTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccessTokenResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AccessTokenResponseBody
	GetTraceId() *string
}

type AccessTokenResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data    *AccessTokenResponseBodyData   `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                        `json:"message,omitempty" xml:"message,omitempty"`
	Module  *AccessTokenResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 成功标识
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

func (s AccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccessTokenResponseBody) GetData() *AccessTokenResponseBodyData {
	return s.Data
}

func (s *AccessTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccessTokenResponseBody) GetModule() *AccessTokenResponseBodyModule {
	return s.Module
}

func (s *AccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccessTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccessTokenResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AccessTokenResponseBody) SetCode(v string) *AccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *AccessTokenResponseBody) SetData(v *AccessTokenResponseBodyData) *AccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *AccessTokenResponseBody) SetMessage(v string) *AccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *AccessTokenResponseBody) SetModule(v *AccessTokenResponseBodyModule) *AccessTokenResponseBody {
	s.Module = v
	return s
}

func (s *AccessTokenResponseBody) SetRequestId(v string) *AccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessTokenResponseBody) SetSuccess(v bool) *AccessTokenResponseBody {
	s.Success = &v
	return s
}

func (s *AccessTokenResponseBody) SetTraceId(v string) *AccessTokenResponseBody {
	s.TraceId = &v
	return s
}

func (s *AccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type AccessTokenResponseBodyData struct {
	// example:
	//
	// 70000
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// 37j76df
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s AccessTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBodyData) GetExpire() *int64 {
	return s.Expire
}

func (s *AccessTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *AccessTokenResponseBodyData) SetExpire(v int64) *AccessTokenResponseBodyData {
	s.Expire = &v
	return s
}

func (s *AccessTokenResponseBodyData) SetToken(v string) *AccessTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *AccessTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AccessTokenResponseBodyModule struct {
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
	// 37j76df
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s AccessTokenResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AccessTokenResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBodyModule) GetExpire() *int64 {
	return s.Expire
}

func (s *AccessTokenResponseBodyModule) GetStart() *int64 {
	return s.Start
}

func (s *AccessTokenResponseBodyModule) GetToken() *string {
	return s.Token
}

func (s *AccessTokenResponseBodyModule) SetExpire(v int64) *AccessTokenResponseBodyModule {
	s.Expire = &v
	return s
}

func (s *AccessTokenResponseBodyModule) SetStart(v int64) *AccessTokenResponseBodyModule {
	s.Start = &v
	return s
}

func (s *AccessTokenResponseBodyModule) SetToken(v string) *AccessTokenResponseBodyModule {
	s.Token = &v
	return s
}

func (s *AccessTokenResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
