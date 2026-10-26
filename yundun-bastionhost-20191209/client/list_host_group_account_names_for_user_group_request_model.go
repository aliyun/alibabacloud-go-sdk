// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupAccountNamesForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *ListHostGroupAccountNamesForUserGroupRequest
	GetInstanceId() *string
	SetProjectId(v int64) *ListHostGroupAccountNamesForUserGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListHostGroupAccountNamesForUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest
	GetUserGroupId() *string
}

type ListHostGroupAccountNamesForUserGroupRequest struct {
	// The ID of the host group.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The instance ID of the Bastionhost instance to which the specified user group belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The region ID of the Bastionhost instance to which the specified user group belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetHostGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetInstanceId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetProjectId(v int64) *ListHostGroupAccountNamesForUserGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetRegionId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetUserGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
