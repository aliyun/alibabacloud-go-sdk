// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowLogSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateChatFlowLogSettingResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateChatFlowLogSettingResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *UpdateChatFlowLogSettingResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *UpdateChatFlowLogSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateChatFlowLogSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateChatFlowLogSettingResponseBody
	GetSuccess() *bool
}

type UpdateChatFlowLogSettingResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateChatFlowLogSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowLogSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowLogSettingResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateChatFlowLogSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChatFlowLogSettingResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UpdateChatFlowLogSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateChatFlowLogSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChatFlowLogSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateChatFlowLogSettingResponseBody) SetAccessDeniedDetail(v string) *UpdateChatFlowLogSettingResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateChatFlowLogSettingResponseBody) SetCode(v string) *UpdateChatFlowLogSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChatFlowLogSettingResponseBody) SetData(v map[string]interface{}) *UpdateChatFlowLogSettingResponseBody {
	s.Data = v
	return s
}

func (s *UpdateChatFlowLogSettingResponseBody) SetMessage(v string) *UpdateChatFlowLogSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateChatFlowLogSettingResponseBody) SetRequestId(v string) *UpdateChatFlowLogSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChatFlowLogSettingResponseBody) SetSuccess(v bool) *UpdateChatFlowLogSettingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateChatFlowLogSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
