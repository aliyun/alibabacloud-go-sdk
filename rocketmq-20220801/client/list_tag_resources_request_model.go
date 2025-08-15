// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTagResourcesRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListTagResourcesRequest
	GetResourceId() *string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTag(v string) *ListTagResourcesRequest
	GetTag() *string
}

type ListTagResourcesRequest struct {
	// The position from which the next query starts.
	//
	// example:
	//
	// d09e2b63e1b12d905b7080ff70
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// List of resource IDs, in JSON format.
	//
	// example:
	//
	// ["rmq-cn-pe334n08h08"]
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// List of tags, in JSON format.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTagResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTag() *string {
	return s.Tag
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v string) *ListTagResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v string) *ListTagResourcesRequest {
	s.Tag = &v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
