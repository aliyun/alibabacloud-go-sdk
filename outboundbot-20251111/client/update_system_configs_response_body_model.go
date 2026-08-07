// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSystemConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSystemConfigsResponseBody
	GetCode() *string
	SetData(v string) *UpdateSystemConfigsResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateSystemConfigsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSystemConfigsResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateSystemConfigsResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateSystemConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSystemConfigsResponseBody
	GetSuccess() *bool
}

type UpdateSystemConfigsResponseBody struct {
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
	// 暂无使用
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

func (s UpdateSystemConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSystemConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSystemConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSystemConfigsResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateSystemConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSystemConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSystemConfigsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateSystemConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSystemConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSystemConfigsResponseBody) SetCode(v string) *UpdateSystemConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSystemConfigsResponseBody) SetData(v string) *UpdateSystemConfigsResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSystemConfigsResponseBody) SetHttpStatusCode(v int32) *UpdateSystemConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSystemConfigsResponseBody) SetMessage(v string) *UpdateSystemConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSystemConfigsResponseBody) SetParams(v []*string) *UpdateSystemConfigsResponseBody {
	s.Params = v
	return s
}

func (s *UpdateSystemConfigsResponseBody) SetRequestId(v string) *UpdateSystemConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSystemConfigsResponseBody) SetSuccess(v bool) *UpdateSystemConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSystemConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
