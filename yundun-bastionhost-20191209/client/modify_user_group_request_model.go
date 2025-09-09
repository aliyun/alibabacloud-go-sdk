// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyUserGroupRequest
	GetComment() *string
	SetInstanceId(v string) *ModifyUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ModifyUserGroupRequest
	GetUserGroupId() *string
	SetUserGroupName(v string) *ModifyUserGroupRequest
	GetUserGroupName() *string
}

type ModifyUserGroupRequest struct {
	// The new description of the user group. The description can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host in which you want to modify the information about the user group.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to modify the information about the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to modify.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The new name of the user group. This name can be up to 128 characters in length.
	//
	// example:
	//
	// TestUserGroup
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ModifyUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ModifyUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ModifyUserGroupRequest) SetComment(v string) *ModifyUserGroupRequest {
	s.Comment = &v
	return s
}

func (s *ModifyUserGroupRequest) SetInstanceId(v string) *ModifyUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserGroupRequest) SetRegionId(v string) *ModifyUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserGroupRequest) SetUserGroupId(v string) *ModifyUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ModifyUserGroupRequest) SetUserGroupName(v string) *ModifyUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *ModifyUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
