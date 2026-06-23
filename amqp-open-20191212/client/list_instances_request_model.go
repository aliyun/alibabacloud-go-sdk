// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
}

type ListInstancesRequest struct {
	// The maximum number of results to return. The recommended value is from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. If this is your first query, leave this parameter empty.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group that contains the instance.
	//
	// example:
	//
	// rg-aekzu74zjgdu4mq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
