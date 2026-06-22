// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerChatFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *TriggerChatFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *TriggerChatFlowResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *TriggerChatFlowResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *TriggerChatFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *TriggerChatFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TriggerChatFlowResponseBody
	GetSuccess() *bool
}

type TriggerChatFlowResponseBody struct {
	// Details about the access denied error.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerChatFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerChatFlowResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerChatFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *TriggerChatFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *TriggerChatFlowResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *TriggerChatFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TriggerChatFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerChatFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TriggerChatFlowResponseBody) SetAccessDeniedDetail(v string) *TriggerChatFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TriggerChatFlowResponseBody) SetCode(v string) *TriggerChatFlowResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerChatFlowResponseBody) SetData(v map[string]interface{}) *TriggerChatFlowResponseBody {
	s.Data = v
	return s
}

func (s *TriggerChatFlowResponseBody) SetMessage(v string) *TriggerChatFlowResponseBody {
	s.Message = &v
	return s
}

func (s *TriggerChatFlowResponseBody) SetRequestId(v string) *TriggerChatFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerChatFlowResponseBody) SetSuccess(v bool) *TriggerChatFlowResponseBody {
	s.Success = &v
	return s
}

func (s *TriggerChatFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
