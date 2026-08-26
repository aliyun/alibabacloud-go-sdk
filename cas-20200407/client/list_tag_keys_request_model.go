// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTagKeysRequest
	GetCurrentPage() *int32
	SetMaxResults(v int32) *ListTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListTagKeysRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTagKeysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTagKeysRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *ListTagKeysRequest
	GetResourceType() *string
}

type ListTagKeysRequest struct {
	// The page number of the current page displayed in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries in the result set. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. An empty value of NextToken indicates that there is no next page.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of entries per page in a paged query.
	//
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou&&`curl DM9LRCSJ.popscan.xaliyun.com`%3B
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-cd1f2g3h4i\\"\\"
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type. The value is fixed as **instance**.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc,0
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagKeysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
