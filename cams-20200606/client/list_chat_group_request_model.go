// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *ListChatGroupRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *ListChatGroupRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *ListChatGroupRequest
	GetCustSpaceId() *string
	SetGroupStatus(v string) *ListChatGroupRequest
	GetGroupStatus() *string
	SetOwnerId(v int64) *ListChatGroupRequest
	GetOwnerId() *int64
	SetPage(v *ListChatGroupRequestPage) *ListChatGroupRequest
	GetPage() *ListChatGroupRequestPage
	SetResourceOwnerAccount(v string) *ListChatGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatGroupRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *ListChatGroupRequest
	GetSubject() *string
}

type ListChatGroupRequest struct {
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
	Page                 *ListChatGroupRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s ListChatGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupRequest) GoString() string {
	return s.String()
}

func (s *ListChatGroupRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatGroupRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatGroupRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatGroupRequest) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *ListChatGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatGroupRequest) GetPage() *ListChatGroupRequestPage {
	return s.Page
}

func (s *ListChatGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatGroupRequest) GetSubject() *string {
	return s.Subject
}

func (s *ListChatGroupRequest) SetBusinessNumber(v string) *ListChatGroupRequest {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatGroupRequest) SetChannelType(v string) *ListChatGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *ListChatGroupRequest) SetCustSpaceId(v string) *ListChatGroupRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatGroupRequest) SetGroupStatus(v string) *ListChatGroupRequest {
	s.GroupStatus = &v
	return s
}

func (s *ListChatGroupRequest) SetOwnerId(v int64) *ListChatGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatGroupRequest) SetPage(v *ListChatGroupRequestPage) *ListChatGroupRequest {
	s.Page = v
	return s
}

func (s *ListChatGroupRequest) SetResourceOwnerAccount(v string) *ListChatGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatGroupRequest) SetResourceOwnerId(v int64) *ListChatGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatGroupRequest) SetSubject(v string) *ListChatGroupRequest {
	s.Subject = &v
	return s
}

func (s *ListChatGroupRequest) Validate() error {
	return dara.Validate(s)
}

type ListChatGroupRequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListChatGroupRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupRequestPage) GoString() string {
	return s.String()
}

func (s *ListChatGroupRequestPage) GetIndex() *int64 {
	return s.Index
}

func (s *ListChatGroupRequestPage) GetSize() *int64 {
	return s.Size
}

func (s *ListChatGroupRequestPage) SetIndex(v int64) *ListChatGroupRequestPage {
	s.Index = &v
	return s
}

func (s *ListChatGroupRequestPage) SetSize(v int64) *ListChatGroupRequestPage {
	s.Size = &v
	return s
}

func (s *ListChatGroupRequestPage) Validate() error {
	return dara.Validate(s)
}
