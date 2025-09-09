// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddUsersToGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *AddUsersToGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *AddUsersToGroupRequest
	GetUserGroupId() *string
	SetUserIds(v string) *AddUsersToGroupRequest
	GetUserIds() *string
}

type AddUsersToGroupRequest struct {
	// The ID of the bastion host for which you want to add users to the user group.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to add users to the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group to which you want to add users.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// １
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user that you want to add to the user group. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s AddUsersToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUsersToGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUsersToGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUsersToGroupRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *AddUsersToGroupRequest) SetInstanceId(v string) *AddUsersToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetRegionId(v string) *AddUsersToGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetUserGroupId(v string) *AddUsersToGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetUserIds(v string) *AddUsersToGroupRequest {
	s.UserIds = &v
	return s
}

func (s *AddUsersToGroupRequest) Validate() error {
	return dara.Validate(s)
}
