// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateChatFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateChatFlowResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateChatFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateChatFlowResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *UpdateChatFlowResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *UpdateChatFlowResponseBody
	GetSuccess() *bool
}

type UpdateChatFlowResponseBody struct {
	// Details of access denial; this field is only returned when RAM verification fails.
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
	// Content of the returned data.
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Indicates whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateChatFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateChatFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChatFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateChatFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChatFlowResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *UpdateChatFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateChatFlowResponseBody) SetAccessDeniedDetail(v string) *UpdateChatFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateChatFlowResponseBody) SetCode(v string) *UpdateChatFlowResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChatFlowResponseBody) SetMessage(v string) *UpdateChatFlowResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateChatFlowResponseBody) SetRequestId(v string) *UpdateChatFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChatFlowResponseBody) SetResponse(v map[string]interface{}) *UpdateChatFlowResponseBody {
	s.Response = v
	return s
}

func (s *UpdateChatFlowResponseBody) SetSuccess(v bool) *UpdateChatFlowResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateChatFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
