// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfficeSiteUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedInfo(v string) *ListOfficeSiteUsersRequest
	GetAssignedInfo() *string
	SetFilter(v string) *ListOfficeSiteUsersRequest
	GetFilter() *string
	SetIncludeAssignedUser(v bool) *ListOfficeSiteUsersRequest
	GetIncludeAssignedUser() *bool
	SetMaxResults(v int32) *ListOfficeSiteUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOfficeSiteUsersRequest
	GetNextToken() *string
	SetOUPath(v string) *ListOfficeSiteUsersRequest
	GetOUPath() *string
	SetOfficeSiteId(v string) *ListOfficeSiteUsersRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ListOfficeSiteUsersRequest
	GetRegionId() *string
	SetSortType(v string) *ListOfficeSiteUsersRequest
	GetSortType() *string
}

type ListOfficeSiteUsersRequest struct {
	// > This parameter is not publicly available. You can only pass in `1` or leave it empty.
	//
	// example:
	//
	// 1
	AssignedInfo *string `json:"AssignedInfo,omitempty" xml:"AssignedInfo,omitempty"`
	// The fuzzy query character string.
	//
	// example:
	//
	// *jin*
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return only users who are assigned cloud computers.
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
	// The pagination token. Leave this parameter empty for the first request or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The specified AD domain organizational unit (OU).
	//
	// example:
	//
	// example.com/Domain Controllers
	OUPath *string `json:"OUPath,omitempty" xml:"OUPath,omitempty"`
	// The office network ID. Only office networks based on enterprise AD accounts are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. Call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
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

func (s ListOfficeSiteUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteUsersRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersRequest) GetAssignedInfo() *string {
	return s.AssignedInfo
}

func (s *ListOfficeSiteUsersRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListOfficeSiteUsersRequest) GetIncludeAssignedUser() *bool {
	return s.IncludeAssignedUser
}

func (s *ListOfficeSiteUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOfficeSiteUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOfficeSiteUsersRequest) GetOUPath() *string {
	return s.OUPath
}

func (s *ListOfficeSiteUsersRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListOfficeSiteUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfficeSiteUsersRequest) GetSortType() *string {
	return s.SortType
}

func (s *ListOfficeSiteUsersRequest) SetAssignedInfo(v string) *ListOfficeSiteUsersRequest {
	s.AssignedInfo = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetFilter(v string) *ListOfficeSiteUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetIncludeAssignedUser(v bool) *ListOfficeSiteUsersRequest {
	s.IncludeAssignedUser = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetMaxResults(v int32) *ListOfficeSiteUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetNextToken(v string) *ListOfficeSiteUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetOUPath(v string) *ListOfficeSiteUsersRequest {
	s.OUPath = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetOfficeSiteId(v string) *ListOfficeSiteUsersRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetRegionId(v string) *ListOfficeSiteUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetSortType(v string) *ListOfficeSiteUsersRequest {
	s.SortType = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) Validate() error {
	return dara.Validate(s)
}
