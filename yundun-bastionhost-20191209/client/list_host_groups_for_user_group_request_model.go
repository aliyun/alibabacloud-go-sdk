// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupName(v string) *ListHostGroupsForUserGroupRequest
	GetHostGroupName() *string
	SetInstanceId(v string) *ListHostGroupsForUserGroupRequest
	GetInstanceId() *string
	SetMode(v string) *ListHostGroupsForUserGroupRequest
	GetMode() *string
	SetPageNumber(v string) *ListHostGroupsForUserGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostGroupsForUserGroupRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostGroupsForUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ListHostGroupsForUserGroupRequest
	GetUserGroupId() *string
}

type ListHostGroupsForUserGroupRequest struct {
	// The name of the host group that you want to query. Only exact match is supported.
	//
	// example:
	//
	// group
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host to which the user group belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the category of the host group that you want to query. Valid values:
	//
	// 	- **Authorized**: queries the host groups that the user group is authorized to manage. This is the default value.
	//
	// 	- **Unauthorized**: queries the host groups that the user group is not authorized to manage.
	//
	// example:
	//
	// Authorized
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host to which the user group belongs.
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

func (s ListHostGroupsForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupRequest) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ListHostGroupsForUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostGroupsForUserGroupRequest) GetMode() *string {
	return s.Mode
}

func (s *ListHostGroupsForUserGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostGroupsForUserGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostGroupsForUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostGroupsForUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListHostGroupsForUserGroupRequest) SetHostGroupName(v string) *ListHostGroupsForUserGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetInstanceId(v string) *ListHostGroupsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetMode(v string) *ListHostGroupsForUserGroupRequest {
	s.Mode = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetPageNumber(v string) *ListHostGroupsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetPageSize(v string) *ListHostGroupsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetRegionId(v string) *ListHostGroupsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetUserGroupId(v string) *ListHostGroupsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
