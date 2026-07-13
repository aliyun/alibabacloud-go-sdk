// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductDataRedundancyTypeStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *DescribeProductDataRedundancyTypeStatRequest
	GetProductType() *string
	SetResourceCategoryId(v string) *DescribeProductDataRedundancyTypeStatRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIds(v []*int64) *DescribeProductDataRedundancyTypeStatRequest
	GetResourceOwnerIds() []*int64
}

type DescribeProductDataRedundancyTypeStatRequest struct {
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
	ResourceCategoryId *string  `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	ResourceOwnerIds   []*int64 `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty" type:"Repeated"`
}

func (s DescribeProductDataRedundancyTypeStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductDataRedundancyTypeStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductDataRedundancyTypeStatRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeProductDataRedundancyTypeStatRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeProductDataRedundancyTypeStatRequest) GetResourceOwnerIds() []*int64 {
	return s.ResourceOwnerIds
}

func (s *DescribeProductDataRedundancyTypeStatRequest) SetProductType(v string) *DescribeProductDataRedundancyTypeStatRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatRequest) SetResourceCategoryId(v string) *DescribeProductDataRedundancyTypeStatRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatRequest) SetResourceOwnerIds(v []*int64) *DescribeProductDataRedundancyTypeStatRequest {
	s.ResourceOwnerIds = v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatRequest) Validate() error {
	return dara.Validate(s)
}
