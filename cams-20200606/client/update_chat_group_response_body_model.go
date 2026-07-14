// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateChatGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateChatGroupResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateChatGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateChatGroupResponseBody
	GetRequestId() *string
	SetResult(v int64) *UpdateChatGroupResponseBody
	GetResult() *int64
	SetSuccess(v bool) *UpdateChatGroupResponseBody
	GetSuccess() *bool
}

type UpdateChatGroupResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
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
	// The number of affected rows.
	//
	// example:
	//
	// 1
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateChatGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChatGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateChatGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChatGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateChatGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChatGroupResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *UpdateChatGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateChatGroupResponseBody) SetAccessDeniedDetail(v string) *UpdateChatGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateChatGroupResponseBody) SetCode(v string) *UpdateChatGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChatGroupResponseBody) SetMessage(v string) *UpdateChatGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateChatGroupResponseBody) SetRequestId(v string) *UpdateChatGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChatGroupResponseBody) SetResult(v int64) *UpdateChatGroupResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateChatGroupResponseBody) SetSuccess(v bool) *UpdateChatGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateChatGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
