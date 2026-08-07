// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *DescribeProductsShrinkRequest
	GetProductType() *string
	SetResourceCategoryId(v string) *DescribeProductsShrinkRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIdsShrink(v string) *DescribeProductsShrinkRequest
	GetResourceOwnerIdsShrink() *string
	SetResourceRegionId(v string) *DescribeProductsShrinkRequest
	GetResourceRegionId() *string
}

type DescribeProductsShrinkRequest struct {
	// Filters by the product type of the resource. If this parameter is not specified, all types are queried.
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
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	// The list of resource owner IDs. This parameter is used in cross-account scenarios. If this parameter is left empty, data of the current account is returned by default.
	//
	// example:
	//
	// [123***7890]
	ResourceOwnerIdsShrink *string `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty"`
	// Filters by the region where the resource resides. If this parameter is not specified, all regions are queried.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeProductsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductsShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeProductsShrinkRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeProductsShrinkRequest) GetResourceOwnerIdsShrink() *string {
	return s.ResourceOwnerIdsShrink
}

func (s *DescribeProductsShrinkRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeProductsShrinkRequest) SetProductType(v string) *DescribeProductsShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeProductsShrinkRequest) SetResourceCategoryId(v string) *DescribeProductsShrinkRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeProductsShrinkRequest) SetResourceOwnerIdsShrink(v string) *DescribeProductsShrinkRequest {
	s.ResourceOwnerIdsShrink = &v
	return s
}

func (s *DescribeProductsShrinkRequest) SetResourceRegionId(v string) *DescribeProductsShrinkRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeProductsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
