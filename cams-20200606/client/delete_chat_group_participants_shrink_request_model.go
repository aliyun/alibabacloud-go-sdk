// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupParticipantsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *DeleteChatGroupParticipantsShrinkRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *DeleteChatGroupParticipantsShrinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *DeleteChatGroupParticipantsShrinkRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *DeleteChatGroupParticipantsShrinkRequest
	GetGroupId() *string
	SetListShrink(v string) *DeleteChatGroupParticipantsShrinkRequest
	GetListShrink() *string
	SetOwnerId(v int64) *DeleteChatGroupParticipantsShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatGroupParticipantsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatGroupParticipantsShrinkRequest
	GetResourceOwnerId() *int64
}

type DeleteChatGroupParticipantsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8613800**
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// example:
	//
	// WHATSAPP。
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EA93UU****
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ListShrink           *string `json:"List,omitempty" xml:"List,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteChatGroupParticipantsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupParticipantsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetListShrink() *string {
	return s.ListShrink
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatGroupParticipantsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetBusinessNumber(v string) *DeleteChatGroupParticipantsShrinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetChannelType(v string) *DeleteChatGroupParticipantsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetCustSpaceId(v string) *DeleteChatGroupParticipantsShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetGroupId(v string) *DeleteChatGroupParticipantsShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetListShrink(v string) *DeleteChatGroupParticipantsShrinkRequest {
	s.ListShrink = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetOwnerId(v int64) *DeleteChatGroupParticipantsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetResourceOwnerAccount(v string) *DeleteChatGroupParticipantsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) SetResourceOwnerId(v int64) *DeleteChatGroupParticipantsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatGroupParticipantsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
