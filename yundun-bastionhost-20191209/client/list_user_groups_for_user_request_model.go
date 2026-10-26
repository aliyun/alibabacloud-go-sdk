// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUserGroupsForUserRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListUserGroupsForUserRequest
	GetKeyword() *string
	SetPageNumber(v string) *ListUserGroupsForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListUserGroupsForUserRequest
	GetPageSize() *string
	SetPagingOff(v string) *ListUserGroupsForUserRequest
	GetPagingOff() *string
	SetProjectId(v int64) *ListUserGroupsForUserRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListUserGroupsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListUserGroupsForUserRequest
	GetUserId() *string
}

type ListUserGroupsForUserRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-m413tuhlo03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the user group to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number for a paged query. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// The maximum value of PageSize is 100. The default number of entries per page is 20. If PageSize is left empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to disable paging. Valid values:
	//
	// - **true**: Disables paging and returns all data.
	//
	// - **false**: Enables paging and returns the number of user groups specified by PageSize.
	//
	// example:
	//
	// true
	PagingOff *string `json:"PagingOff,omitempty" xml:"PagingOff,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserGroupsForUserRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserGroupsForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListUserGroupsForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListUserGroupsForUserRequest) GetPagingOff() *string {
	return s.PagingOff
}

func (s *ListUserGroupsForUserRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListUserGroupsForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserGroupsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserGroupsForUserRequest) SetInstanceId(v string) *ListUserGroupsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetKeyword(v string) *ListUserGroupsForUserRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetPageNumber(v string) *ListUserGroupsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetPageSize(v string) *ListUserGroupsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetPagingOff(v string) *ListUserGroupsForUserRequest {
	s.PagingOff = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetProjectId(v int64) *ListUserGroupsForUserRequest {
	s.ProjectId = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetRegionId(v string) *ListUserGroupsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserGroupsForUserRequest) SetUserId(v string) *ListUserGroupsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListUserGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
