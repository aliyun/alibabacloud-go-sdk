// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupInviteLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *DeleteChatGroupInviteLinkRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *DeleteChatGroupInviteLinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *DeleteChatGroupInviteLinkRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *DeleteChatGroupInviteLinkRequest
	GetGroupId() *string
	SetOwnerId(v int64) *DeleteChatGroupInviteLinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatGroupInviteLinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatGroupInviteLinkRequest
	GetResourceOwnerId() *int64
}

type DeleteChatGroupInviteLinkRequest struct {
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
	// cams-8***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// E398****
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteChatGroupInviteLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupInviteLinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupInviteLinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *DeleteChatGroupInviteLinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *DeleteChatGroupInviteLinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteChatGroupInviteLinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteChatGroupInviteLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatGroupInviteLinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatGroupInviteLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatGroupInviteLinkRequest) SetBusinessNumber(v string) *DeleteChatGroupInviteLinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) SetChannelType(v string) *DeleteChatGroupInviteLinkRequest {
	s.ChannelType = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) SetCustSpaceId(v string) *DeleteChatGroupInviteLinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) SetGroupId(v string) *DeleteChatGroupInviteLinkRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) SetOwnerId(v int64) *DeleteChatGroupInviteLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) SetResourceOwnerAccount(v string) *DeleteChatGroupInviteLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) SetResourceOwnerId(v int64) *DeleteChatGroupInviteLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatGroupInviteLinkRequest) Validate() error {
	return dara.Validate(s)
}
