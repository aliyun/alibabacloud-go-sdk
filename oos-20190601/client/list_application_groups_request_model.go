// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ListApplicationGroupsRequest
	GetApplicationName() *string
	SetDeployRegionId(v string) *ListApplicationGroupsRequest
	GetDeployRegionId() *string
	SetMaxResults(v int32) *ListApplicationGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationGroupsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListApplicationGroupsRequest
	GetRegionId() *string
	SetResourceId(v string) *ListApplicationGroupsRequest
	GetResourceId() *string
	SetResourceProduct(v string) *ListApplicationGroupsRequest
	GetResourceProduct() *string
	SetResourceType(v string) *ListApplicationGroupsRequest
	GetResourceType() *string
}

type ListApplicationGroupsRequest struct {
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the region in which the related resources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud resource.
	//
	// example:
	//
	// i-2vcj9raxrhxb48zz3whw
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The code of the product to which the cloud resource belongs.
	//
	// example:
	//
	// ecs
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The type of the cloud resource.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListApplicationGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationGroupsRequest) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *ListApplicationGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationGroupsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListApplicationGroupsRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *ListApplicationGroupsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListApplicationGroupsRequest) SetApplicationName(v string) *ListApplicationGroupsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetDeployRegionId(v string) *ListApplicationGroupsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetMaxResults(v int32) *ListApplicationGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetNextToken(v string) *ListApplicationGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetRegionId(v string) *ListApplicationGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetResourceId(v string) *ListApplicationGroupsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetResourceProduct(v string) *ListApplicationGroupsRequest {
	s.ResourceProduct = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetResourceType(v string) *ListApplicationGroupsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListApplicationGroupsRequest) Validate() error {
	return dara.Validate(s)
}
