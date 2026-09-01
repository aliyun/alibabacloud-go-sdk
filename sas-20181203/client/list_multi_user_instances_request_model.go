// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiUserInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMultiUserInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiUserInstancesRequest
	GetNextToken() *string
}

type ListMultiUserInstancesRequest struct {
	// The number of entries per page for a paged query. Maximum value: 100. Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used for paging. Leave this parameter empty for the first request. For subsequent requests, set this parameter to the NextToken value returned in the previous response.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAB4SwmEAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM0NzY3YTZjNjI3NjZmNmU3MjcxNjk3NDY5MzY3MjY4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMultiUserInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiUserInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiUserInstancesRequest) SetMaxResults(v int32) *ListMultiUserInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiUserInstancesRequest) SetNextToken(v string) *ListMultiUserInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiUserInstancesRequest) Validate() error {
	return dara.Validate(s)
}
