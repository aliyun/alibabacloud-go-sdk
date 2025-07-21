// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowLogSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateChatFlowLogSettingResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateChatFlowLogSettingResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateChatFlowLogSettingResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateChatFlowLogSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatFlowLogSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateChatFlowLogSettingResponseBody
	GetSuccess() *bool
}

type CreateChatFlowLogSettingResponseBody struct {
	// Access denied details.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
	//
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// 无
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Values: true for success, false for failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateChatFlowLogSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowLogSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatFlowLogSettingResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateChatFlowLogSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatFlowLogSettingResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateChatFlowLogSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatFlowLogSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatFlowLogSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateChatFlowLogSettingResponseBody) SetAccessDeniedDetail(v string) *CreateChatFlowLogSettingResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateChatFlowLogSettingResponseBody) SetCode(v string) *CreateChatFlowLogSettingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatFlowLogSettingResponseBody) SetData(v map[string]interface{}) *CreateChatFlowLogSettingResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatFlowLogSettingResponseBody) SetMessage(v string) *CreateChatFlowLogSettingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatFlowLogSettingResponseBody) SetRequestId(v string) *CreateChatFlowLogSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatFlowLogSettingResponseBody) SetSuccess(v bool) *CreateChatFlowLogSettingResponseBody {
	s.Success = &v
	return s
}

func (s *CreateChatFlowLogSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
