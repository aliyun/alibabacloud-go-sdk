// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAdOrganizationUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListUserAdOrganizationUnitsRequest
	GetFilter() *string
	SetMaxResults(v int32) *ListUserAdOrganizationUnitsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserAdOrganizationUnitsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *ListUserAdOrganizationUnitsRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ListUserAdOrganizationUnitsRequest
	GetRegionId() *string
}

type ListUserAdOrganizationUnitsRequest struct {
	// The string that you enter for fuzzy search.
	//
	// example:
	//
	// develop
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 500. Default value: 500.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request or if no next request exists. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// CAAAAA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The enterprise AD office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-485361****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserAdOrganizationUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserAdOrganizationUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListUserAdOrganizationUnitsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserAdOrganizationUnitsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAdOrganizationUnitsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListUserAdOrganizationUnitsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserAdOrganizationUnitsRequest) SetFilter(v string) *ListUserAdOrganizationUnitsRequest {
	s.Filter = &v
	return s
}

func (s *ListUserAdOrganizationUnitsRequest) SetMaxResults(v int32) *ListUserAdOrganizationUnitsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserAdOrganizationUnitsRequest) SetNextToken(v string) *ListUserAdOrganizationUnitsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserAdOrganizationUnitsRequest) SetOfficeSiteId(v string) *ListUserAdOrganizationUnitsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListUserAdOrganizationUnitsRequest) SetRegionId(v string) *ListUserAdOrganizationUnitsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserAdOrganizationUnitsRequest) Validate() error {
	return dara.Validate(s)
}
