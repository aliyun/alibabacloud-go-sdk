// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductDataRedundancyTypeStatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *DescribeProductDataRedundancyTypeStatShrinkRequest
	GetProductType() *string
	SetResourceCategoryId(v string) *DescribeProductDataRedundancyTypeStatShrinkRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIdsShrink(v string) *DescribeProductDataRedundancyTypeStatShrinkRequest
	GetResourceOwnerIdsShrink() *string
}

type DescribeProductDataRedundancyTypeStatShrinkRequest struct {
	// The type of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource category ID.
	//
	// example:
	//
	// rc-000***123
	ResourceCategoryId     *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	ResourceOwnerIdsShrink *string `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty"`
}

func (s DescribeProductDataRedundancyTypeStatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductDataRedundancyTypeStatShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) GetResourceOwnerIdsShrink() *string {
	return s.ResourceOwnerIdsShrink
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) SetProductType(v string) *DescribeProductDataRedundancyTypeStatShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) SetResourceCategoryId(v string) *DescribeProductDataRedundancyTypeStatShrinkRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) SetResourceOwnerIdsShrink(v string) *DescribeProductDataRedundancyTypeStatShrinkRequest {
	s.ResourceOwnerIdsShrink = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
