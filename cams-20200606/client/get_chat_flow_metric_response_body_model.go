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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
