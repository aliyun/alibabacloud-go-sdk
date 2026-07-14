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
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation is successful.
	//
	// - false: The operation failed.
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
