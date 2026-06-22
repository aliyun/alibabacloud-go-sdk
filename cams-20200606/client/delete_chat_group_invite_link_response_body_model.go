// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupInviteLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteChatGroupInviteLinkResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteChatGroupInviteLinkResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteChatGroupInviteLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatGroupInviteLinkResponseBody
	GetRequestId() *string
	SetResult(v int64) *DeleteChatGroupInviteLinkResponseBody
	GetResult() *int64
	SetSuccess(v bool) *DeleteChatGroupInviteLinkResponseBody
	GetSuccess() *bool
}

type DeleteChatGroupInviteLinkResponseBody struct {
	// Details about the access denial.
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
	// Indicates if the request was successful. Valid values:
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

func (s DeleteChatGroupInviteLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupInviteLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupInviteLinkResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteChatGroupInviteLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatGroupInviteLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatGroupInviteLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatGroupInviteLinkResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *DeleteChatGroupInviteLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChatGroupInviteLinkResponseBody) SetAccessDeniedDetail(v string) *DeleteChatGroupInviteLinkResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponseBody) SetCode(v string) *DeleteChatGroupInviteLinkResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponseBody) SetMessage(v string) *DeleteChatGroupInviteLinkResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponseBody) SetRequestId(v string) *DeleteChatGroupInviteLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponseBody) SetResult(v int64) *DeleteChatGroupInviteLinkResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponseBody) SetSuccess(v bool) *DeleteChatGroupInviteLinkResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
