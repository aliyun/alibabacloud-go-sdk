// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatGroupInviteLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *AddChatGroupInviteLinkRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *AddChatGroupInviteLinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *AddChatGroupInviteLinkRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *AddChatGroupInviteLinkRequest
	GetGroupId() *string
	SetOwnerId(v int64) *AddChatGroupInviteLinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddChatGroupInviteLinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddChatGroupInviteLinkRequest
	GetResourceOwnerId() *int64
}

type AddChatGroupInviteLinkRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// E93kdk****
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddChatGroupInviteLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddChatGroupInviteLinkRequest) GoString() string {
	return s.String()
}

func (s *AddChatGroupInviteLinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *AddChatGroupInviteLinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *AddChatGroupInviteLinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddChatGroupInviteLinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddChatGroupInviteLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddChatGroupInviteLinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddChatGroupInviteLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddChatGroupInviteLinkRequest) SetBusinessNumber(v string) *AddChatGroupInviteLinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) SetChannelType(v string) *AddChatGroupInviteLinkRequest {
	s.ChannelType = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) SetCustSpaceId(v string) *AddChatGroupInviteLinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) SetGroupId(v string) *AddChatGroupInviteLinkRequest {
	s.GroupId = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) SetOwnerId(v int64) *AddChatGroupInviteLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) SetResourceOwnerAccount(v string) *AddChatGroupInviteLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) SetResourceOwnerId(v int64) *AddChatGroupInviteLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddChatGroupInviteLinkRequest) Validate() error {
	return dara.Validate(s)
}
