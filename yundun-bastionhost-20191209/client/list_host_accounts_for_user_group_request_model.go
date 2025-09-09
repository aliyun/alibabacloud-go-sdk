// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountName(v string) *ListHostAccountsForUserGroupRequest
	GetHostAccountName() *string
	SetHostId(v string) *ListHostAccountsForUserGroupRequest
	GetHostId() *string
	SetInstanceId(v string) *ListHostAccountsForUserGroupRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListHostAccountsForUserGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostAccountsForUserGroupRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostAccountsForUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ListHostAccountsForUserGroupRequest
	GetUserGroupId() *string
}

type ListHostAccountsForUserGroupRequest struct {
	// The name of the host account that you want to query. Exact match is supported.
	//
	// example:
	//
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host to query.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host on which you want to query the host accounts to be managed by the specified user group on the specified host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The region ID of the bastion host on which you want to query the host accounts to be managed by the specified user group on the specified host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group for which you want to query authorized host accounts.
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

func (s ListHostAccountsForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupRequest) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsForUserGroupRequest) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsForUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostAccountsForUserGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostAccountsForUserGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostAccountsForUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostAccountsForUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListHostAccountsForUserGroupRequest) SetHostAccountName(v string) *ListHostAccountsForUserGroupRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetHostId(v string) *ListHostAccountsForUserGroupRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetInstanceId(v string) *ListHostAccountsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetPageNumber(v string) *ListHostAccountsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetPageSize(v string) *ListHostAccountsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetRegionId(v string) *ListHostAccountsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetUserGroupId(v string) *ListHostAccountsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
