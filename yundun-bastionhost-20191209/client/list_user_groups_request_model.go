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
	SetProjectId(v int64) *ListUserGroupsRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListUserGroupsRequest
	GetRegionId() *string
	SetUserGroupName(v string) *ListUserGroupsRequest
	GetUserGroupName() *string
}

type ListUserGroupsRequest struct {
	// The ID of the Bastionhost instance whose user groups you want to query.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for paging. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging. Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned by default.
	//
	// > Specify a value for this parameter.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user group to query. Only exact match is supported.
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

func (s *ListUserGroupsRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ListUserGroupsRequest) SetProjectId(v int64) *ListUserGroupsRequest {
	s.ProjectId = &v
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
