// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteChatGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteChatGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteChatGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatGroupResponseBody
	GetRequestId() *string
	SetResult(v int64) *DeleteChatGroupResponseBody
	GetResult() *int64
	SetSuccess(v bool) *DeleteChatGroupResponseBody
	GetSuccess() *bool
}

type DeleteChatGroupResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For information about other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteChatGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteChatGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatGroupResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *DeleteChatGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChatGroupResponseBody) SetAccessDeniedDetail(v string) *DeleteChatGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetCode(v string) *DeleteChatGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetMessage(v string) *DeleteChatGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetRequestId(v string) *DeleteChatGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetResult(v int64) *DeleteChatGroupResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetSuccess(v bool) *DeleteChatGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChatGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
