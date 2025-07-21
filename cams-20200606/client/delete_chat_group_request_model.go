// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *DeleteChatGroupRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *DeleteChatGroupRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *DeleteChatGroupRequest
	GetCustSpaceId() *string
	SetGroupId(v string) *DeleteChatGroupRequest
	GetGroupId() *string
	SetOwnerId(v int64) *DeleteChatGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteChatGroupRequest struct {
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
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// E399**
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteChatGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *DeleteChatGroupRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *DeleteChatGroupRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteChatGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteChatGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatGroupRequest) SetBusinessNumber(v string) *DeleteChatGroupRequest {
	s.BusinessNumber = &v
	return s
}

func (s *DeleteChatGroupRequest) SetChannelType(v string) *DeleteChatGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *DeleteChatGroupRequest) SetCustSpaceId(v string) *DeleteChatGroupRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteChatGroupRequest) SetGroupId(v string) *DeleteChatGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteChatGroupRequest) SetOwnerId(v int64) *DeleteChatGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatGroupRequest) SetResourceOwnerAccount(v string) *DeleteChatGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatGroupRequest) SetResourceOwnerId(v int64) *DeleteChatGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatGroupRequest) Validate() error {
	return dara.Validate(s)
}
