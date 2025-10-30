// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnrolledAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEnrolledAccountsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEnrolledAccountsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListEnrolledAccountsRequest
	GetRegionId() *string
}

type ListEnrolledAccountsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnrolledAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnrolledAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEnrolledAccountsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEnrolledAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnrolledAccountsRequest) SetMaxResults(v int32) *ListEnrolledAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEnrolledAccountsRequest) SetNextToken(v string) *ListEnrolledAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEnrolledAccountsRequest) SetRegionId(v string) *ListEnrolledAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnrolledAccountsRequest) Validate() error {
	return dara.Validate(s)
}
