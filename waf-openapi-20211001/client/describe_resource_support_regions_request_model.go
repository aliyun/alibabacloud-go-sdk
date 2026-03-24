// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceSupportRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeResourceSupportRegionsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeResourceSupportRegionsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourceSupportRegionsRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceProduct(v string) *DescribeResourceSupportRegionsRequest
	GetResourceProduct() *string
}

type DescribeResourceSupportRegionsRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to view the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region to which the WAF instance belongs. Valid values:
	//
	// - **cn-hangzhou**: indicates the Chinese mainland.
	//
	// - **ap-southeast-1**: indicates regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzpks****kdjq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The cloud product to which the resource belongs. By default, instances of ALB, MSE, FC, and SAE products are returned. Valid values:
	//
	// - **alb**: indicates the ALB product.
	//
	// - **mse**: indicates the MSE product.
	//
	// - **fc**: indicates the FC product.
	//
	// - **sae**: indicates the SAE product.
	//
	// - **ecs**: indicates the ECS product.
	//
	// - **clb4**: indicates the CLB(TCP) product.
	//
	// - **clb7**: indicates the CLB(HTTP/HTTPS) product.
	//
	// - **apig**: indicates the APIG product.
	//
	// - **nlb**: indicates the NLB product.
	//
	// > Each product supports different regions. When the product filter field has a value, you need to refer to the regions supported by the product. Otherwise, the filtering may fail.
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s DescribeResourceSupportRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceSupportRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceSupportRegionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceSupportRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceSupportRegionsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourceSupportRegionsRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeResourceSupportRegionsRequest) SetInstanceId(v string) *DescribeResourceSupportRegionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceSupportRegionsRequest) SetRegionId(v string) *DescribeResourceSupportRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceSupportRegionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceSupportRegionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceSupportRegionsRequest) SetResourceProduct(v string) *DescribeResourceSupportRegionsRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeResourceSupportRegionsRequest) Validate() error {
	return dara.Validate(s)
}
