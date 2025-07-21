// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *ListChatGroupShrinkRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *ListChatGroupShrinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *ListChatGroupShrinkRequest
	GetCustSpaceId() *string
	SetGroupStatus(v string) *ListChatGroupShrinkRequest
	GetGroupStatus() *string
	SetOwnerId(v int64) *ListChatGroupShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListChatGroupShrinkRequest
	GetPageShrink() *string
	SetResourceOwnerAccount(v string) *ListChatGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatGroupShrinkRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *ListChatGroupShrinkRequest
	GetSubject() *string
}

type ListChatGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8613800***
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-x***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// ACTIVE
	GroupStatus *string `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageShrink           *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s ListChatGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatGroupShrinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatGroupShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatGroupShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatGroupShrinkRequest) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *ListChatGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatGroupShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListChatGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatGroupShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *ListChatGroupShrinkRequest) SetBusinessNumber(v string) *ListChatGroupShrinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetChannelType(v string) *ListChatGroupShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetCustSpaceId(v string) *ListChatGroupShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetGroupStatus(v string) *ListChatGroupShrinkRequest {
	s.GroupStatus = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetOwnerId(v int64) *ListChatGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetPageShrink(v string) *ListChatGroupShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetResourceOwnerAccount(v string) *ListChatGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetResourceOwnerId(v int64) *ListChatGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetSubject(v string) *ListChatGroupShrinkRequest {
	s.Subject = &v
	return s
}

func (s *ListChatGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
