// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *DescribeProductsRequest
	GetProductType() *string
	SetResourceCategoryId(v string) *DescribeProductsRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIds(v []*int64) *DescribeProductsRequest
	GetResourceOwnerIds() []*int64
	SetResourceRegionId(v string) *DescribeProductsRequest
	GetResourceRegionId() *string
}

type DescribeProductsRequest struct {
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
	ResourceOwnerIds []*int64 `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty" type:"Repeated"`
	// Filters by the region where the resource resides. If this parameter is not specified, all regions are queried.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeProductsRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeProductsRequest) GetResourceOwnerIds() []*int64 {
	return s.ResourceOwnerIds
}

func (s *DescribeProductsRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeProductsRequest) SetProductType(v string) *DescribeProductsRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeProductsRequest) SetResourceCategoryId(v string) *DescribeProductsRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeProductsRequest) SetResourceOwnerIds(v []*int64) *DescribeProductsRequest {
	s.ResourceOwnerIds = v
	return s
}

func (s *DescribeProductsRequest) SetResourceRegionId(v string) *DescribeProductsRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeProductsRequest) Validate() error {
	return dara.Validate(s)
}
