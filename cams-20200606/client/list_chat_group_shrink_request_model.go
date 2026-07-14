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
	SetGroupLink(v string) *ListChatGroupShrinkRequest
	GetGroupLink() *string
	SetGroupStatus(v string) *ListChatGroupShrinkRequest
	GetGroupStatus() *string
	SetGroupType(v string) *ListChatGroupShrinkRequest
	GetGroupType() *string
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
	// The business phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800***
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// The channel type. Valid values: **WHATSAPP**.
	//
	// > Only the WhatsApp channel type is supported.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The SpaceId or instance ID of the ISV sub-customer. This is the channel ID, which can be viewed on the [Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-x***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Telegram group link.
	//
	// example:
	//
	// example
	GroupLink *string `json:"GroupLink,omitempty" xml:"GroupLink,omitempty"`
	// The group status. Valid values:
	//
	// - ACTIVE: In use.
	//
	// - INACTIVE: Not activated.
	//
	// - SUSPENDED: Suspended.
	//
	// - CREATING: Being created.
	//
	// - DELETING: Being deleted.
	//
	// - UPDATING: Being updated.
	//
	// example:
	//
	// ACTIVE
	GroupStatus *string `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	// The Telegram group type.
	//
	// example:
	//
	// example
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pagination information.
	//
	// This parameter is required.
	PageShrink           *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The group subject.
	//
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

func (s *ListChatGroupShrinkRequest) GetGroupLink() *string {
	return s.GroupLink
}

func (s *ListChatGroupShrinkRequest) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *ListChatGroupShrinkRequest) GetGroupType() *string {
	return s.GroupType
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

func (s *ListChatGroupShrinkRequest) SetGroupLink(v string) *ListChatGroupShrinkRequest {
	s.GroupLink = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetGroupStatus(v string) *ListChatGroupShrinkRequest {
	s.GroupStatus = &v
	return s
}

func (s *ListChatGroupShrinkRequest) SetGroupType(v string) *ListChatGroupShrinkRequest {
	s.GroupType = &v
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
