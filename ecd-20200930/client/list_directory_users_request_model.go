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
	// > This parameter is not publicly available. The value can be 1 or left empty.
	//
	// example:
	//
	// 1
	AssignedInfo *string `json:"AssignedInfo,omitempty" xml:"AssignedInfo,omitempty"`
	// The ID of the AD directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The query string for fuzzy match. If you specify this parameter, the system returns all results that contain the string.
	//
	// example:
	//
	// alice
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return the users with assigned cloud computers only.
	//
	// example:
	//
	// true
	IncludeAssignedUser *bool `json:"IncludeAssignedUser,omitempty" xml:"IncludeAssignedUser,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query. If the value of this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organizational unit (OU) in the specified AD domain.
	//
	// example:
	//
	// example.com/Domain Controllers
	OUPath *string `json:"OUPath,omitempty" xml:"OUPath,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sort type.
	//
	// Valide values:
	//
	// - asc: cloud computers assigned to users on bottom
	//
	// - desc: cloud computers assigned to users on top
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
