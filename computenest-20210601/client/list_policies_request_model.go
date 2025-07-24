// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPoliciesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListPoliciesRequest
	GetRegionId() *string
}

type ListPoliciesRequest struct {
	// Page size.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next query, an empty nextToken indicates there is no next page.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPoliciesRequest) SetMaxResults(v int32) *ListPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesRequest) SetNextToken(v string) *ListPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesRequest) SetRegionId(v string) *ListPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
