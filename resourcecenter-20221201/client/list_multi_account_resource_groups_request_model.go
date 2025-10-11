// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListMultiAccountResourceGroupsRequest
	GetAccountId() *string
	SetMaxResults(v int32) *ListMultiAccountResourceGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountResourceGroupsRequest
	GetNextToken() *string
	SetResourceGroupIds(v []*string) *ListMultiAccountResourceGroupsRequest
	GetResourceGroupIds() []*string
}

type ListMultiAccountResourceGroupsRequest struct {
	// The ID of the management account or member of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1394339739****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAS2Nboi3t4xGrdlG5/Ks/Q1xPG9jzviYEuZydevXIkgF
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of resource groups.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListMultiAccountResourceGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountResourceGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountResourceGroupsRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *ListMultiAccountResourceGroupsRequest) SetAccountId(v string) *ListMultiAccountResourceGroupsRequest {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetMaxResults(v int32) *ListMultiAccountResourceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetNextToken(v string) *ListMultiAccountResourceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListMultiAccountResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
