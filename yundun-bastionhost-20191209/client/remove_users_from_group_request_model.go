// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveUsersFromGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *RemoveUsersFromGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *RemoveUsersFromGroupRequest
	GetUserGroupId() *string
	SetUserIds(v string) *RemoveUsersFromGroupRequest
	GetUserIds() *string
}

type RemoveUsersFromGroupRequest struct {
	// The ID of the bastion host for which you want to remove users from the user group.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to remove users from the user group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to remove users.
	//
	// >  You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// １
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user who you want to remove. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the IDs of users.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s RemoveUsersFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveUsersFromGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveUsersFromGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUsersFromGroupRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *RemoveUsersFromGroupRequest) SetInstanceId(v string) *RemoveUsersFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetRegionId(v string) *RemoveUsersFromGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetUserGroupId(v string) *RemoveUsersFromGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetUserIds(v string) *RemoveUsersFromGroupRequest {
	s.UserIds = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
