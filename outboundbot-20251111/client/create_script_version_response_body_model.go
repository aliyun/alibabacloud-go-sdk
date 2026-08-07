// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScriptVersionResponseBody
	GetCode() *string
	SetData(v string) *CreateScriptVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateScriptVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScriptVersionResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateScriptVersionResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateScriptVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateScriptVersionResponseBody
	GetSuccess() *bool
}

type CreateScriptVersionResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b26
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Instance does not exist. Instance=outb003
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScriptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScriptVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateScriptVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScriptVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScriptVersionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateScriptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScriptVersionResponseBody) SetCode(v string) *CreateScriptVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetData(v string) *CreateScriptVersionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetHttpStatusCode(v int32) *CreateScriptVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetMessage(v string) *CreateScriptVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetParams(v []*string) *CreateScriptVersionResponseBody {
	s.Params = v
	return s
}

func (s *CreateScriptVersionResponseBody) SetRequestId(v string) *CreateScriptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetSuccess(v bool) *CreateScriptVersionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScriptVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
