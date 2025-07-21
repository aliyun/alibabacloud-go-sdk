// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *UpdateChatGroupRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *UpdateChatGroupRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *UpdateChatGroupRequest
	GetCustSpaceId() *string
	SetDescription(v string) *UpdateChatGroupRequest
	GetDescription() *string
	SetGroupId(v string) *UpdateChatGroupRequest
	GetGroupId() *string
	SetOwnerId(v int64) *UpdateChatGroupRequest
	GetOwnerId() *int64
	SetProfilePictureFile(v string) *UpdateChatGroupRequest
	GetProfilePictureFile() *string
	SetResourceOwnerAccount(v string) *UpdateChatGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateChatGroupRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *UpdateChatGroupRequest
	GetSubject() *string
}

type UpdateChatGroupRequest struct {
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
	// cams-***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EA303***
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// https://aliyun.com/img.jpg
	ProfilePictureFile   *string `json:"ProfilePictureFile,omitempty" xml:"ProfilePictureFile,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s UpdateChatGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatGroupRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *UpdateChatGroupRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *UpdateChatGroupRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateChatGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateChatGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateChatGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateChatGroupRequest) GetProfilePictureFile() *string {
	return s.ProfilePictureFile
}

func (s *UpdateChatGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateChatGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateChatGroupRequest) GetSubject() *string {
	return s.Subject
}

func (s *UpdateChatGroupRequest) SetBusinessNumber(v string) *UpdateChatGroupRequest {
	s.BusinessNumber = &v
	return s
}

func (s *UpdateChatGroupRequest) SetChannelType(v string) *UpdateChatGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *UpdateChatGroupRequest) SetCustSpaceId(v string) *UpdateChatGroupRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateChatGroupRequest) SetDescription(v string) *UpdateChatGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateChatGroupRequest) SetGroupId(v string) *UpdateChatGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateChatGroupRequest) SetOwnerId(v int64) *UpdateChatGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateChatGroupRequest) SetProfilePictureFile(v string) *UpdateChatGroupRequest {
	s.ProfilePictureFile = &v
	return s
}

func (s *UpdateChatGroupRequest) SetResourceOwnerAccount(v string) *UpdateChatGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateChatGroupRequest) SetResourceOwnerId(v int64) *UpdateChatGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateChatGroupRequest) SetSubject(v string) *UpdateChatGroupRequest {
	s.Subject = &v
	return s
}

func (s *UpdateChatGroupRequest) Validate() error {
	return dara.Validate(s)
}
