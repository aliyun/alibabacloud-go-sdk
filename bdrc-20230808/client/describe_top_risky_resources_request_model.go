// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopRiskyResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCategoryId(v string) *DescribeTopRiskyResourcesRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIds(v []*int64) *DescribeTopRiskyResourcesRequest
	GetResourceOwnerIds() []*int64
	SetResourceType(v string) *DescribeTopRiskyResourcesRequest
	GetResourceType() *string
}

type DescribeTopRiskyResourcesRequest struct {
	// The ID of the resource category.
	//
	// example:
	//
	// rc-000***123
	ResourceCategoryId *string  `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	ResourceOwnerIds   []*int64 `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTopRiskyResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopRiskyResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeTopRiskyResourcesRequest) GetResourceOwnerIds() []*int64 {
	return s.ResourceOwnerIds
}

func (s *DescribeTopRiskyResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTopRiskyResourcesRequest) SetResourceCategoryId(v string) *DescribeTopRiskyResourcesRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeTopRiskyResourcesRequest) SetResourceOwnerIds(v []*int64) *DescribeTopRiskyResourcesRequest {
	s.ResourceOwnerIds = v
	return s
}

func (s *DescribeTopRiskyResourcesRequest) SetResourceType(v string) *DescribeTopRiskyResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTopRiskyResourcesRequest) Validate() error {
	return dara.Validate(s)
}
