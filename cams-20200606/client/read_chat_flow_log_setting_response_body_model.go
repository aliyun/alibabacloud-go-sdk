// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowLogSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReadChatFlowLogSettingResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ReadChatFlowLogSettingResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *ReadChatFlowLogSettingResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *ReadChatFlowLogSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadChatFlowLogSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadChatFlowLogSettingResponseBody
	GetSuccess() *bool
}

type ReadChatFlowLogSettingResponseBody struct {
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
	// OK
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
	// Whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadChatFlowLogSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowLogSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ReadChatFlowLogSettingResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReadChatFlowLogSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadChatFlowLogSettingResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ReadChatFlowLogSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadChatFlowLogSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadChatFlowLogSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadChatFlowLogSettingResponseBody) SetAccessDeniedDetail(v string) *ReadChatFlowLogSettingResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReadChatFlowLogSettingResponseBody) SetCode(v string) *ReadChatFlowLogSettingResponseBody {
	s.Code = &v
	return s
}

func (s *ReadChatFlowLogSettingResponseBody) SetData(v map[string]interface{}) *ReadChatFlowLogSettingResponseBody {
	s.Data = v
	return s
}

func (s *ReadChatFlowLogSettingResponseBody) SetMessage(v string) *ReadChatFlowLogSettingResponseBody {
	s.Message = &v
	return s
}

func (s *ReadChatFlowLogSettingResponseBody) SetRequestId(v string) *ReadChatFlowLogSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadChatFlowLogSettingResponseBody) SetSuccess(v bool) *ReadChatFlowLogSettingResponseBody {
	s.Success = &v
	return s
}

func (s *ReadChatFlowLogSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
