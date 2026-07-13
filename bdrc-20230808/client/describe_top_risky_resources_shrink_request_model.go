// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopRiskyResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCategoryId(v string) *DescribeTopRiskyResourcesShrinkRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIdsShrink(v string) *DescribeTopRiskyResourcesShrinkRequest
	GetResourceOwnerIdsShrink() *string
	SetResourceType(v string) *DescribeTopRiskyResourcesShrinkRequest
	GetResourceType() *string
}

type DescribeTopRiskyResourcesShrinkRequest struct {
	// The ID of the resource category.
	//
	// example:
	//
	// rc-000***123
	ResourceCategoryId     *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	ResourceOwnerIdsShrink *string `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTopRiskyResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopRiskyResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesShrinkRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeTopRiskyResourcesShrinkRequest) GetResourceOwnerIdsShrink() *string {
	return s.ResourceOwnerIdsShrink
}

func (s *DescribeTopRiskyResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTopRiskyResourcesShrinkRequest) SetResourceCategoryId(v string) *DescribeTopRiskyResourcesShrinkRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeTopRiskyResourcesShrinkRequest) SetResourceOwnerIdsShrink(v string) *DescribeTopRiskyResourcesShrinkRequest {
	s.ResourceOwnerIdsShrink = &v
	return s
}

func (s *DescribeTopRiskyResourcesShrinkRequest) SetResourceType(v string) *DescribeTopRiskyResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTopRiskyResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
