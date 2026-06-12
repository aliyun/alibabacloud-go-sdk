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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If this parameter is empty, no more results exist.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
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
