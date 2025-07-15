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
	AssignedInfo *string `json:"AssignedInfo,omitempty" xml:"AssignedInfo,omitempty"`
	// The query string for fuzzy query.
	//
	// example:
	//
	// *jin*
	Filter              *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	IncludeAssignedUser *bool   `json:"IncludeAssignedUser,omitempty" xml:"IncludeAssignedUser,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 100.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request or if no next query exists. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organizational unit (OU) of the specified AD domain.
	//
	// example:
	//
	// example.com/Domain Controllers
	OUPath *string `json:"OUPath,omitempty" xml:"OUPath,omitempty"`
	// The office network ID. The office network must be of the enterprise AD account type.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
