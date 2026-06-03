// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyCallBackTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateDypnsSmsVerifyCallBackTestResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDypnsSmsVerifyCallBackTestResponseBody
	GetSuccess() *bool
}

type CreateDypnsSmsVerifyCallBackTestResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDypnsSmsVerifyCallBackTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyCallBackTestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) SetAccessDeniedDetail(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) SetCode(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) SetData(v map[string]interface{}) *CreateDypnsSmsVerifyCallBackTestResponseBody {
	s.Data = v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) SetMessage(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) SetRequestId(v string) *CreateDypnsSmsVerifyCallBackTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) SetSuccess(v bool) *CreateDypnsSmsVerifyCallBackTestResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponseBody) Validate() error {
	return dara.Validate(s)
}
