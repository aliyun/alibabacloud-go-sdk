// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderSCIMCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetSuccess() *bool
}

type DeleteIdentityProviderSCIMCredentialResponseBody struct {
	// 失败时结合 HTTP 状态及 message 处理。
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 本次请求的处理状态，不等同后台任务最终结果。
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// 请求受理或错误消息；异步任务结果需按对应查询接口核对。
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 用于关联本次调用与排障，失败时请保留。
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 表示本次请求处理成功；异步写入仍需继续查询任务结果。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteIdentityProviderSCIMCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderSCIMCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetCode(v string) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetHttpStatusCode(v int32) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetMessage(v string) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetRequestId(v string) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetSuccess(v bool) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
