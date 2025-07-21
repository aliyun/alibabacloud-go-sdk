// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupParticipantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *ListChatGroupParticipantsRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *ListChatGroupParticipantsRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *ListChatGroupParticipantsRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *ListChatGroupParticipantsRequest
	GetGroupId() *string
	SetOwnerId(v int64) *ListChatGroupParticipantsRequest
	GetOwnerId() *int64
	SetPage(v *ListChatGroupParticipantsRequestPage) *ListChatGroupParticipantsRequest
	GetPage() *ListChatGroupParticipantsRequestPage
	SetResourceOwnerAccount(v string) *ListChatGroupParticipantsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatGroupParticipantsRequest
	GetResourceOwnerId() *int64
}

type ListChatGroupParticipantsRequest struct {
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
	GroupId              *string                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Page                 *ListChatGroupParticipantsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListChatGroupParticipantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsRequest) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatGroupParticipantsRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatGroupParticipantsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatGroupParticipantsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListChatGroupParticipantsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatGroupParticipantsRequest) GetPage() *ListChatGroupParticipantsRequestPage {
	return s.Page
}

func (s *ListChatGroupParticipantsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatGroupParticipantsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatGroupParticipantsRequest) SetBusinessNumber(v string) *ListChatGroupParticipantsRequest {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetChannelType(v string) *ListChatGroupParticipantsRequest {
	s.ChannelType = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetCustSpaceId(v string) *ListChatGroupParticipantsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetGroupId(v string) *ListChatGroupParticipantsRequest {
	s.GroupId = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetOwnerId(v int64) *ListChatGroupParticipantsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetPage(v *ListChatGroupParticipantsRequestPage) *ListChatGroupParticipantsRequest {
	s.Page = v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetResourceOwnerAccount(v string) *ListChatGroupParticipantsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) SetResourceOwnerId(v int64) *ListChatGroupParticipantsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatGroupParticipantsRequest) Validate() error {
	return dara.Validate(s)
}

type ListChatGroupParticipantsRequestPage struct {
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListChatGroupParticipantsRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsRequestPage) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsRequestPage) GetIndex() *int64 {
	return s.Index
}

func (s *ListChatGroupParticipantsRequestPage) GetSize() *int64 {
	return s.Size
}

func (s *ListChatGroupParticipantsRequestPage) SetIndex(v int64) *ListChatGroupParticipantsRequestPage {
	s.Index = &v
	return s
}

func (s *ListChatGroupParticipantsRequestPage) SetSize(v int64) *ListChatGroupParticipantsRequestPage {
	s.Size = &v
	return s
}

func (s *ListChatGroupParticipantsRequestPage) Validate() error {
	return dara.Validate(s)
}
