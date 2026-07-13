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
	// The product type of the resource. If you omit this parameter, the API returns resources of all product types. For example, specify `oss` to query resources from Object Storage Service.
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
	// A list of resource owner IDs for cross-account resource queries. If you omit this parameter, the API returns resources from the current account.
	//
	// example:
	//
	// [123***7890]
	ResourceOwnerIds []*int64 `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty" type:"Repeated"`
	// The ID of the resource\\"s region. If you omit this parameter, the API returns resources from all regions.
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
