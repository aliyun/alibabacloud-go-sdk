// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoryUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedInfo(v string) *ListDirectoryUsersRequest
	GetAssignedInfo() *string
	SetDirectoryId(v string) *ListDirectoryUsersRequest
	GetDirectoryId() *string
	SetFilter(v string) *ListDirectoryUsersRequest
	GetFilter() *string
	SetIncludeAssignedUser(v bool) *ListDirectoryUsersRequest
	GetIncludeAssignedUser() *bool
	SetMaxResults(v int32) *ListDirectoryUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDirectoryUsersRequest
	GetNextToken() *string
	SetOUPath(v string) *ListDirectoryUsersRequest
	GetOUPath() *string
	SetRegionId(v string) *ListDirectoryUsersRequest
	GetRegionId() *string
	SetSortType(v string) *ListDirectoryUsersRequest
	GetSortType() *string
}

type ListDirectoryUsersRequest struct {
	// > This field is not available for use. You can only pass in `1` or leave it empty.
	//
	// example:
	//
	// 1
	AssignedInfo *string `json:"AssignedInfo,omitempty" xml:"AssignedInfo,omitempty"`
	// The AD directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The fuzzy match query string. All results that contain this character string are returned.
	//
	// example:
	//
	// alice
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return only users who have been assigned cloud computers.
	//
	// example:
	//
	// true
	IncludeAssignedUser *bool `json:"IncludeAssignedUser,omitempty" xml:"IncludeAssignedUser,omitempty"`
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. An empty value indicates that no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organizational unit (OU) in the AD domain to which the user belongs. You can call [ListUserAdOrganizationUnits](https://help.aliyun.com/document_detail/311259.html) to obtain this value.
	//
	// example:
	//
	// example.com/Domain Controllers
	OUPath *string `json:"OUPath,omitempty" xml:"OUPath,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sorting method.
	//
	// example:
	//
	// asc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListDirectoryUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoryUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersRequest) GetAssignedInfo() *string {
	return s.AssignedInfo
}

func (s *ListDirectoryUsersRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListDirectoryUsersRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListDirectoryUsersRequest) GetIncludeAssignedUser() *bool {
	return s.IncludeAssignedUser
}

func (s *ListDirectoryUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDirectoryUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDirectoryUsersRequest) GetOUPath() *string {
	return s.OUPath
}

func (s *ListDirectoryUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDirectoryUsersRequest) GetSortType() *string {
	return s.SortType
}

func (s *ListDirectoryUsersRequest) SetAssignedInfo(v string) *ListDirectoryUsersRequest {
	s.AssignedInfo = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetDirectoryId(v string) *ListDirectoryUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetFilter(v string) *ListDirectoryUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetIncludeAssignedUser(v bool) *ListDirectoryUsersRequest {
	s.IncludeAssignedUser = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetMaxResults(v int32) *ListDirectoryUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetNextToken(v string) *ListDirectoryUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetOUPath(v string) *ListDirectoryUsersRequest {
	s.OUPath = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetRegionId(v string) *ListDirectoryUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetSortType(v string) *ListDirectoryUsersRequest {
	s.SortType = &v
	return s
}

func (s *ListDirectoryUsersRequest) Validate() error {
	return dara.Validate(s)
}
