// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSupabaseProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupabaseProjectsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListSupabaseProjectsRequest
	GetRegionId() *string
}

type ListSupabaseProjectsRequest struct {
	// The maximum number of instances to return per page. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token returned from a previous call. Use it to retrieve the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation get a list of available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSupabaseProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupabaseProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupabaseProjectsRequest) SetMaxResults(v int32) *ListSupabaseProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupabaseProjectsRequest) SetNextToken(v string) *ListSupabaseProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseProjectsRequest) SetRegionId(v string) *ListSupabaseProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupabaseProjectsRequest) Validate() error {
	return dara.Validate(s)
}
