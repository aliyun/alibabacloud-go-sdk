// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *UpdateGroupNameRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateGroupNameRequest
	GetGroupName() *string
	SetOwnerId(v int64) *UpdateGroupNameRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateGroupNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateGroupNameRequest
	GetResourceOwnerId() *int64
}

type UpdateGroupNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateGroupNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupNameRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGroupNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateGroupNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateGroupNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateGroupNameRequest) SetGroupId(v string) *UpdateGroupNameRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupNameRequest) SetGroupName(v string) *UpdateGroupNameRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupNameRequest) SetOwnerId(v int64) *UpdateGroupNameRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateGroupNameRequest) SetResourceOwnerAccount(v string) *UpdateGroupNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateGroupNameRequest) SetResourceOwnerId(v int64) *UpdateGroupNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateGroupNameRequest) Validate() error {
	return dara.Validate(s)
}
