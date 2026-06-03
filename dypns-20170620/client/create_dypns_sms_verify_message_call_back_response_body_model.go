// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyMessageCallBackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateDypnsSmsVerifyMessageCallBackResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDypnsSmsVerifyMessageCallBackResponseBody
	GetSuccess() *bool
}

type CreateDypnsSmsVerifyMessageCallBackResponseBody struct {
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
	// 示例值
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

func (s CreateDypnsSmsVerifyMessageCallBackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyMessageCallBackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) SetAccessDeniedDetail(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) SetCode(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) SetData(v map[string]interface{}) *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	s.Data = v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) SetMessage(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) SetRequestId(v string) *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) SetSuccess(v bool) *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponseBody) Validate() error {
	return dara.Validate(s)
}
