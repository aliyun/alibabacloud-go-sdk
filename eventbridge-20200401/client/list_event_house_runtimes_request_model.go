// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventHouseRuntimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEventHouseRuntimesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEventHouseRuntimesRequest
	GetNextToken() *string
}

type ListEventHouseRuntimesRequest struct {
	// The maximum number of runtimes to return per page. If this parameter is not specified or set to 0, the default value 20 is used. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Not required for the first query. For subsequent queries, use the NextToken returned in the previous response. An empty value indicates that no more pages are available.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListEventHouseRuntimesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventHouseRuntimesRequest) GoString() string {
	return s.String()
}

func (s *ListEventHouseRuntimesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventHouseRuntimesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventHouseRuntimesRequest) SetMaxResults(v int32) *ListEventHouseRuntimesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventHouseRuntimesRequest) SetNextToken(v string) *ListEventHouseRuntimesRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventHouseRuntimesRequest) Validate() error {
	return dara.Validate(s)
}
