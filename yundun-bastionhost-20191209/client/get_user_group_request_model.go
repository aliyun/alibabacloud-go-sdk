// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *GetUserGroupRequest
	GetUserGroupId() *string
}

type GetUserGroupRequest struct {
	// The ID of the bastion host in which you want to query the details of the user group.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to query the details of the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetUserGroupRequest) SetInstanceId(v string) *GetUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserGroupRequest) SetRegionId(v string) *GetUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
