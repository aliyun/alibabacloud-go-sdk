// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *AddChatGroupRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *AddChatGroupRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *AddChatGroupRequest
	GetCustSpaceId() *string
	SetDescription(v string) *AddChatGroupRequest
	GetDescription() *string
	SetGroupLink(v string) *AddChatGroupRequest
	GetGroupLink() *string
	SetOwnerId(v int64) *AddChatGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddChatGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddChatGroupRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *AddChatGroupRequest
	GetSubject() *string
}

type AddChatGroupRequest struct {
	// The phone number associated with the business account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800***
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// The channel type. Valid value:
	//
	// - **WHATSAPP**
	//
	// > Currently, only the WhatsApp channel type is supported.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The SpaceId or instance ID of the ISV sub-customer, which is the channel ID. Find it on the [Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The group description.
	//
	// example:
	//
	// This is a newly created group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The link to the WhatsApp group.
	//
	// example:
	//
	// https://t.me/testgroup
	GroupLink            *string `json:"GroupLink,omitempty" xml:"GroupLink,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The group title.
	//
	// example:
	//
	// test
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s AddChatGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddChatGroupRequest) GoString() string {
	return s.String()
}

func (s *AddChatGroupRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *AddChatGroupRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *AddChatGroupRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddChatGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *AddChatGroupRequest) GetGroupLink() *string {
	return s.GroupLink
}

func (s *AddChatGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddChatGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddChatGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddChatGroupRequest) GetSubject() *string {
	return s.Subject
}

func (s *AddChatGroupRequest) SetBusinessNumber(v string) *AddChatGroupRequest {
	s.BusinessNumber = &v
	return s
}

func (s *AddChatGroupRequest) SetChannelType(v string) *AddChatGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *AddChatGroupRequest) SetCustSpaceId(v string) *AddChatGroupRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddChatGroupRequest) SetDescription(v string) *AddChatGroupRequest {
	s.Description = &v
	return s
}

func (s *AddChatGroupRequest) SetGroupLink(v string) *AddChatGroupRequest {
	s.GroupLink = &v
	return s
}

func (s *AddChatGroupRequest) SetOwnerId(v int64) *AddChatGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *AddChatGroupRequest) SetResourceOwnerAccount(v string) *AddChatGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddChatGroupRequest) SetResourceOwnerId(v int64) *AddChatGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddChatGroupRequest) SetSubject(v string) *AddChatGroupRequest {
	s.Subject = &v
	return s
}

func (s *AddChatGroupRequest) Validate() error {
	return dara.Validate(s)
}
