// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatFlowMetricResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatFlowMetricResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetChatFlowMetricResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetChatFlowMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatFlowMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChatFlowMetricResponseBody
	GetSuccess() *bool
}

type GetChatFlowMetricResponseBody struct {
	// Details of access denial.
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
	// Returned data object.
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
	// Unique request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChatFlowMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatFlowMetricResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatFlowMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatFlowMetricResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetChatFlowMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatFlowMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatFlowMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChatFlowMetricResponseBody) SetAccessDeniedDetail(v string) *GetChatFlowMetricResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatFlowMetricResponseBody) SetCode(v string) *GetChatFlowMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatFlowMetricResponseBody) SetData(v map[string]interface{}) *GetChatFlowMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetChatFlowMetricResponseBody) SetMessage(v string) *GetChatFlowMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatFlowMetricResponseBody) SetRequestId(v string) *GetChatFlowMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatFlowMetricResponseBody) SetSuccess(v bool) *GetChatFlowMetricResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatFlowMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
