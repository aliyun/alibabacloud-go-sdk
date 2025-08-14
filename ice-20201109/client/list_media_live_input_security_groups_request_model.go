// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveInputSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListMediaLiveInputSecurityGroupsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListMediaLiveInputSecurityGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMediaLiveInputSecurityGroupsRequest
	GetNextToken() *string
	SetSkip(v int32) *ListMediaLiveInputSecurityGroupsRequest
	GetSkip() *int32
	SetSortOrder(v string) *ListMediaLiveInputSecurityGroupsRequest
	GetSortOrder() *string
}

type ListMediaLiveInputSecurityGroupsRequest struct {
	// The keyword of the query. You can perform a fuzzy search on security group ID or name.
	//
	// example:
	//
	// 123
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: If you do not specify this parameter or if you set a value smaller than 10, the default value is 10. If you set a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to be skipped in the query. If the number of entries you attempt to skip exceeds the number of entries that meet the condition, an empty list is returned.
	//
	// example:
	//
	// 20
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The sorting order of the security groups by creation time. Default value: asc. Valid values: desc and asc. asc indicates the ascending order, and desc indicates the descending order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListMediaLiveInputSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputSecurityGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMediaLiveInputSecurityGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaLiveInputSecurityGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaLiveInputSecurityGroupsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListMediaLiveInputSecurityGroupsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListMediaLiveInputSecurityGroupsRequest) SetKeyword(v string) *ListMediaLiveInputSecurityGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsRequest) SetMaxResults(v int32) *ListMediaLiveInputSecurityGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsRequest) SetNextToken(v string) *ListMediaLiveInputSecurityGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsRequest) SetSkip(v int32) *ListMediaLiveInputSecurityGroupsRequest {
	s.Skip = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsRequest) SetSortOrder(v string) *ListMediaLiveInputSecurityGroupsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
