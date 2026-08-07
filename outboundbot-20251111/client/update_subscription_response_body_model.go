// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSubscriptionResponseBody
	GetCode() *string
	SetData(v string) *UpdateSubscriptionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSubscriptionResponseBody
	GetSuccess() *bool
}

type UpdateSubscriptionResponseBody struct {
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
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
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

func (s UpdateSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSubscriptionResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSubscriptionResponseBody) SetCode(v string) *UpdateSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetData(v string) *UpdateSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetHttpStatusCode(v int32) *UpdateSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetMessage(v string) *UpdateSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetParams(v []*string) *UpdateSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetSuccess(v bool) *UpdateSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
