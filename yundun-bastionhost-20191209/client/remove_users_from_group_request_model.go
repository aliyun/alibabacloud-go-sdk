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
	SetProjectId(v int64) *RemoveUsersFromGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *RemoveUsersFromGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *RemoveUsersFromGroupRequest
	GetUserGroupId() *string
	SetUserIds(v string) *RemoveUsersFromGroupRequest
	GetUserIds() *string
}

type RemoveUsersFromGroupRequest struct {
	// The ID of the Bastionhost instance to which the user group from which you want to remove users belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the user group from which you want to remove users belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to remove users.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// １
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The IDs of the users that you want to remove. The value of this parameter is a JSON string. You can specify up to 100 user IDs. Separate multiple IDs with commas (,).
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query user IDs.
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

func (s *RemoveUsersFromGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *RemoveUsersFromGroupRequest) SetProjectId(v int64) *RemoveUsersFromGroupRequest {
	s.ProjectId = &v
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
