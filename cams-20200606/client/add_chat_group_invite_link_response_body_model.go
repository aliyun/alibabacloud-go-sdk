// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatGroupInviteLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddChatGroupInviteLinkResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddChatGroupInviteLinkResponseBody
	GetCode() *string
	SetInviteLink(v string) *AddChatGroupInviteLinkResponseBody
	GetInviteLink() *string
	SetMessage(v string) *AddChatGroupInviteLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddChatGroupInviteLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddChatGroupInviteLinkResponseBody
	GetSuccess() *bool
}

type AddChatGroupInviteLinkResponseBody struct {
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
	// The invite link.
	//
	// example:
	//
	// https://chat.whatsapp.com/****
	InviteLink *string `json:"InviteLink,omitempty" xml:"InviteLink,omitempty"`
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

func (s AddChatGroupInviteLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddChatGroupInviteLinkResponseBody) GoString() string {
	return s.String()
}

func (s *AddChatGroupInviteLinkResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddChatGroupInviteLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddChatGroupInviteLinkResponseBody) GetInviteLink() *string {
	return s.InviteLink
}

func (s *AddChatGroupInviteLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddChatGroupInviteLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddChatGroupInviteLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddChatGroupInviteLinkResponseBody) SetAccessDeniedDetail(v string) *AddChatGroupInviteLinkResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddChatGroupInviteLinkResponseBody) SetCode(v string) *AddChatGroupInviteLinkResponseBody {
	s.Code = &v
	return s
}

func (s *AddChatGroupInviteLinkResponseBody) SetInviteLink(v string) *AddChatGroupInviteLinkResponseBody {
	s.InviteLink = &v
	return s
}

func (s *AddChatGroupInviteLinkResponseBody) SetMessage(v string) *AddChatGroupInviteLinkResponseBody {
	s.Message = &v
	return s
}

func (s *AddChatGroupInviteLinkResponseBody) SetRequestId(v string) *AddChatGroupInviteLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddChatGroupInviteLinkResponseBody) SetSuccess(v bool) *AddChatGroupInviteLinkResponseBody {
	s.Success = &v
	return s
}

func (s *AddChatGroupInviteLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
