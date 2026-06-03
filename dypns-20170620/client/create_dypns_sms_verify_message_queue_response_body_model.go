// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyMessageQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateDypnsSmsVerifyMessageQueueResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDypnsSmsVerifyMessageQueueResponseBody
	GetSuccess() *bool
}

type CreateDypnsSmsVerifyMessageQueueResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//   "RequestId": "C210BE04-8CB4-542A-92E1-44160AB05B01",
	//
	//   "Message": "successful",
	//
	//   "Data": "Alicom-Queue-10********384-DypnsSmsVerifyReport",
	//
	//   "Code": "200",
	//
	//   "Success": true
	//
	// }
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
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

func (s CreateDypnsSmsVerifyMessageQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyMessageQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) SetAccessDeniedDetail(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) SetCode(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) SetData(v map[string]interface{}) *CreateDypnsSmsVerifyMessageQueueResponseBody {
	s.Data = v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) SetMessage(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) SetRequestId(v string) *CreateDypnsSmsVerifyMessageQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) SetSuccess(v bool) *CreateDypnsSmsVerifyMessageQueueResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
