// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupParticipantsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *ListChatGroupParticipantsShrinkRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *ListChatGroupParticipantsShrinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *ListChatGroupParticipantsShrinkRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *ListChatGroupParticipantsShrinkRequest
	GetGroupId() *string
	SetOwnerId(v int64) *ListChatGroupParticipantsShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListChatGroupParticipantsShrinkRequest
	GetPageShrink() *string
	SetResourceOwnerAccount(v string) *ListChatGroupParticipantsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatGroupParticipantsShrinkRequest
	GetResourceOwnerId() *int64
}

type ListChatGroupParticipantsShrinkRequest struct {
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
	// cams-k***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EA939****
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageShrink           *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListChatGroupParticipantsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsShrinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatGroupParticipantsShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatGroupParticipantsShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatGroupParticipantsShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListChatGroupParticipantsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatGroupParticipantsShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListChatGroupParticipantsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatGroupParticipantsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatGroupParticipantsShrinkRequest) SetBusinessNumber(v string) *ListChatGroupParticipantsShrinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetChannelType(v string) *ListChatGroupParticipantsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetCustSpaceId(v string) *ListChatGroupParticipantsShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetGroupId(v string) *ListChatGroupParticipantsShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetOwnerId(v int64) *ListChatGroupParticipantsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetPageShrink(v string) *ListChatGroupParticipantsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetResourceOwnerAccount(v string) *ListChatGroupParticipantsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) SetResourceOwnerId(v int64) *ListChatGroupParticipantsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatGroupParticipantsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
