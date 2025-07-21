// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConversationalAutomationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateConversationalAutomationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateConversationalAutomationResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateConversationalAutomationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConversationalAutomationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConversationalAutomationResponseBody
	GetSuccess() *bool
}

type UpdateConversationalAutomationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConversationalAutomationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConversationalAutomationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConversationalAutomationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateConversationalAutomationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConversationalAutomationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConversationalAutomationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConversationalAutomationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConversationalAutomationResponseBody) SetAccessDeniedDetail(v string) *UpdateConversationalAutomationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateConversationalAutomationResponseBody) SetCode(v string) *UpdateConversationalAutomationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConversationalAutomationResponseBody) SetMessage(v string) *UpdateConversationalAutomationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConversationalAutomationResponseBody) SetRequestId(v string) *UpdateConversationalAutomationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConversationalAutomationResponseBody) SetSuccess(v bool) *UpdateConversationalAutomationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConversationalAutomationResponseBody) Validate() error {
	return dara.Validate(s)
}
