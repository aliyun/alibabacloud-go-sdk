// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUserGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListUserGroupsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListUserGroupsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListUserGroupsRequest
	GetRegionId() *string
	SetUserGroupName(v string) *ListUserGroupsRequest
	GetUserGroupName() *string
}

type ListUserGroupsRequest struct {
	// The ID of the bastion host on which you want to query user groups.
	//
	//  >You can call the [DescribeInstances ](https://help.aliyun.com/document_detail/462953.html)operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query user groups.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/462924.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user group that you want to query. Only exact match is supported.
	//
	// example:
	//
	// TestGroup01
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserGroupsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListUserGroupsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListUserGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserGroupsRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ListUserGroupsRequest) SetInstanceId(v string) *ListUserGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserGroupsRequest) SetPageNumber(v string) *ListUserGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserGroupsRequest) SetPageSize(v string) *ListUserGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupsRequest) SetRegionId(v string) *ListUserGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserGroupsRequest) SetUserGroupName(v string) *ListUserGroupsRequest {
	s.UserGroupName = &v
	return s
}

func (s *ListUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}
