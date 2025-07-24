// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*ClusterSummary) *ListClustersResponseBody
	GetClusters() []*ClusterSummary
	SetMaxResults(v int32) *ListClustersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListClustersResponseBody
	GetTotalCount() *int32
}

type ListClustersResponseBody struct {
	// The clusters.
	Clusters []*ClusterSummary `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the next page returned.
	//
	// example:
	//
	// eyJlY21OZXh0VG9rZW4iOiIxIiwidGFpaGFvTmV4dFRva2VuIjoiNTYiLCJ0YWloYW9OZXh0VG9rZW5JbnQiOjU2LCJlY21OZXh0VG9rZW5JbnQiOjF9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3896A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetClusters() []*ClusterSummary {
	return s.Clusters
}

func (s *ListClustersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClustersResponseBody) SetClusters(v []*ClusterSummary) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetMaxResults(v int32) *ListClustersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListClustersResponseBody) SetNextToken(v string) *ListClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	return dara.Validate(s)
}
