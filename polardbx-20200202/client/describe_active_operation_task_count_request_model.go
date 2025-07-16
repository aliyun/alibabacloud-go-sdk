// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeActiveOperationTaskCountRequest
	GetCategory() *string
	SetProduct(v string) *DescribeActiveOperationTaskCountRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeActiveOperationTaskCountRequest
	GetRegionId() *string
}

type DescribeActiveOperationTaskCountRequest struct {
	// example:
	//
	// Category
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// polarx
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeActiveOperationTaskCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeActiveOperationTaskCountRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeActiveOperationTaskCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActiveOperationTaskCountRequest) SetCategory(v string) *DescribeActiveOperationTaskCountRequest {
	s.Category = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetProduct(v string) *DescribeActiveOperationTaskCountRequest {
	s.Product = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetRegionId(v string) *DescribeActiveOperationTaskCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) Validate() error {
	return dara.Validate(s)
}
