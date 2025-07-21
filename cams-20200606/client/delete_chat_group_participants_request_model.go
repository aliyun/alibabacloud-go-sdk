// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupParticipantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *DeleteChatGroupParticipantsRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *DeleteChatGroupParticipantsRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *DeleteChatGroupParticipantsRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *DeleteChatGroupParticipantsRequest
	GetGroupId() *string
	SetList(v []*DeleteChatGroupParticipantsRequestList) *DeleteChatGroupParticipantsRequest
	GetList() []*DeleteChatGroupParticipantsRequestList
	SetOwnerId(v int64) *DeleteChatGroupParticipantsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatGroupParticipantsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatGroupParticipantsRequest
	GetResourceOwnerId() *int64
}

type DeleteChatGroupParticipantsRequest struct {
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
	GroupId              *string                                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	List                 []*DeleteChatGroupParticipantsRequestList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	OwnerId              *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteChatGroupParticipantsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupParticipantsRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupParticipantsRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *DeleteChatGroupParticipantsRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *DeleteChatGroupParticipantsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteChatGroupParticipantsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteChatGroupParticipantsRequest) GetList() []*DeleteChatGroupParticipantsRequestList {
	return s.List
}

func (s *DeleteChatGroupParticipantsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatGroupParticipantsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatGroupParticipantsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatGroupParticipantsRequest) SetBusinessNumber(v string) *DeleteChatGroupParticipantsRequest {
	s.BusinessNumber = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetChannelType(v string) *DeleteChatGroupParticipantsRequest {
	s.ChannelType = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetCustSpaceId(v string) *DeleteChatGroupParticipantsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetGroupId(v string) *DeleteChatGroupParticipantsRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetList(v []*DeleteChatGroupParticipantsRequestList) *DeleteChatGroupParticipantsRequest {
	s.List = v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetOwnerId(v int64) *DeleteChatGroupParticipantsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetResourceOwnerAccount(v string) *DeleteChatGroupParticipantsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) SetResourceOwnerId(v int64) *DeleteChatGroupParticipantsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteChatGroupParticipantsRequestList struct {
	// example:
	//
	// 86138***
	ParticipantNumber *string `json:"ParticipantNumber,omitempty" xml:"ParticipantNumber,omitempty"`
}

func (s DeleteChatGroupParticipantsRequestList) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupParticipantsRequestList) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupParticipantsRequestList) GetParticipantNumber() *string {
	return s.ParticipantNumber
}

func (s *DeleteChatGroupParticipantsRequestList) SetParticipantNumber(v string) *DeleteChatGroupParticipantsRequestList {
	s.ParticipantNumber = &v
	return s
}

func (s *DeleteChatGroupParticipantsRequestList) Validate() error {
	return dara.Validate(s)
}
