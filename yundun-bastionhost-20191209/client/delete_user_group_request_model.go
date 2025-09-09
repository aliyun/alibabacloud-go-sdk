// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *DeleteUserGroupRequest
	GetUserGroupId() *string
}

type DeleteUserGroupRequest struct {
	// The ID of the bastion host on which you want to delete the user group.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to delete the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to delete.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// １
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DeleteUserGroupRequest) SetInstanceId(v string) *DeleteUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetRegionId(v string) *DeleteUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DeleteUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
