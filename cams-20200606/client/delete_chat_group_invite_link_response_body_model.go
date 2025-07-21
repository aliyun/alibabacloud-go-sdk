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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3R938***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
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
