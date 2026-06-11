// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMembersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMembersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListMembersRequest
	GetRegionId() *string
}

type ListMembersRequest struct {
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token that marks the start of the next page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMembersRequest) SetMaxResults(v int32) *ListMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMembersRequest) SetNextToken(v string) *ListMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListMembersRequest) SetRegionId(v string) *ListMembersRequest {
	s.RegionId = &v
	return s
}

func (s *ListMembersRequest) Validate() error {
	return dara.Validate(s)
}
