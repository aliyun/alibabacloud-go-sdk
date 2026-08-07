// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishScriptResponseBody
	GetCode() *string
	SetData(v string) *PublishScriptResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *PublishScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *PublishScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *PublishScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishScriptResponseBody
	GetSuccess() *bool
}

type PublishScriptResponseBody struct {
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
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b89
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
	// Instance does not exist. Instance=df408e55-63dc-4c52-9161-74265381b6a4
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

func (s PublishScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptResponseBody) GoString() string {
	return s.String()
}

func (s *PublishScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishScriptResponseBody) GetData() *string {
	return s.Data
}

func (s *PublishScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *PublishScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishScriptResponseBody) SetCode(v string) *PublishScriptResponseBody {
	s.Code = &v
	return s
}

func (s *PublishScriptResponseBody) SetData(v string) *PublishScriptResponseBody {
	s.Data = &v
	return s
}

func (s *PublishScriptResponseBody) SetHttpStatusCode(v int32) *PublishScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishScriptResponseBody) SetMessage(v string) *PublishScriptResponseBody {
	s.Message = &v
	return s
}

func (s *PublishScriptResponseBody) SetParams(v []*string) *PublishScriptResponseBody {
	s.Params = v
	return s
}

func (s *PublishScriptResponseBody) SetRequestId(v string) *PublishScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishScriptResponseBody) SetSuccess(v bool) *PublishScriptResponseBody {
	s.Success = &v
	return s
}

func (s *PublishScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
