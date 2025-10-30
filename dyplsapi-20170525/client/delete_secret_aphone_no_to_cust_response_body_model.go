// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretAPhoneNoToCustResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteSecretAPhoneNoToCustResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteSecretAPhoneNoToCustResponseBody
	GetCode() *string
	SetData(v bool) *DeleteSecretAPhoneNoToCustResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteSecretAPhoneNoToCustResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecretAPhoneNoToCustResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecretAPhoneNoToCustResponseBody
	GetSuccess() *bool
}

type DeleteSecretAPhoneNoToCustResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 删除是否成功
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// 失败错误提示
	//
	// example:
	//
	// 号码组不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecretAPhoneNoToCustResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretAPhoneNoToCustResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetAccessDeniedDetail(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetCode(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetData(v bool) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetMessage(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetRequestId(v string) *DeleteSecretAPhoneNoToCustResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) SetSuccess(v bool) *DeleteSecretAPhoneNoToCustResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponseBody) Validate() error {
	return dara.Validate(s)
}
