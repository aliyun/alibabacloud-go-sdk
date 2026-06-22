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
	// The business number. You can obtain the business number by calling the [ListChatGroup](https://help.aliyun.com/document_detail/2932629.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800**
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// The channel type. Valid value:
	//
	// - **WHATSAPP**
	//
	// > Currently, only the WhatsApp channel is supported.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The space ID of the ISV sub-customer or the instance ID. This ID is also the channel ID. You can find the channel ID on the [Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The group ID. You can obtain the group ID by calling the [ListChatGroup](https://help.aliyun.com/document_detail/2932629.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// EA93UU****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// A list of group members to remove.
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
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteChatGroupParticipantsRequestList struct {
	// The participant number of the group member. You can obtain the participant numbers of group members by calling the [ListChatGroupParticipants](https://help.aliyun.com/document_detail/2932628.html) operation.
	//
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
