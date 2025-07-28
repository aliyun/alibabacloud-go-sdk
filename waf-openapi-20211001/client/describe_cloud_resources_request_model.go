// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCloudResourcesRequest
	GetInstanceId() *string
	SetOwnerUserId(v string) *DescribeCloudResourcesRequest
	GetOwnerUserId() *string
	SetPageNumber(v int64) *DescribeCloudResourcesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudResourcesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeCloudResourcesRequest
	GetRegionId() *string
	SetResourceDomain(v string) *DescribeCloudResourcesRequest
	GetResourceDomain() *string
	SetResourceFunction(v string) *DescribeCloudResourcesRequest
	GetResourceFunction() *string
	SetResourceInstanceId(v string) *DescribeCloudResourcesRequest
	GetResourceInstanceId() *string
	SetResourceInstanceName(v string) *DescribeCloudResourcesRequest
	GetResourceInstanceName() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCloudResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceName(v string) *DescribeCloudResourcesRequest
	GetResourceName() *string
	SetResourceProduct(v string) *DescribeCloudResourcesRequest
	GetResourceProduct() *string
	SetResourceRegionId(v string) *DescribeCloudResourcesRequest
	GetResourceRegionId() *string
	SetResourceRouteName(v string) *DescribeCloudResourcesRequest
	GetResourceRouteName() *string
}

type DescribeCloudResourcesRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-lbj****cn0c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 11769793******
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The domain name. You can use this parameter if you set ResourceProduct to fc or sae.
	//
	// example:
	//
	// fc-domain-test
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The function name. You can use this parameter if you set ResourceProduct to fc.
	//
	// example:
	//
	// fc-test
	ResourceFunction *string `json:"ResourceFunction,omitempty" xml:"ResourceFunction,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// alb-43glijk0fr****gths
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The name of the instance that is added to WAF.
	//
	// example:
	//
	// test-name
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm4gh****wela
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Deprecated
	//
	// The name of the resource.
	//
	// example:
	//
	// alb-name
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The cloud service to which the resource belongs. If you do not specify this parameter, the system automatically returns the Application Load Balancer (ALB), Microservices Engine (MSE), Function Compute, and Serverless App Engine (SAE) resources that are added to WAF. Valid values:
	//
	// 	- **alb**: ALB.
	//
	// 	- **mse**: MSE.
	//
	// 	- **fc**: Function Compute.
	//
	// 	- **sae**: SAE.
	//
	// 	- **ecs**: Elastic Compute Service (ECS).
	//
	// 	- **clb4**: Layer 4 Classic Load Balancer (CLB).
	//
	// 	- **clb7**: Layer 7 CLB.
	//
	// 	- **nlb**: Network Load Balancer (NLB).
	//
	// >  Different cloud services are available in different regions. The specified cloud service must be available in the specified region.
	//
	// example:
	//
	// alb
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the resource. For information about region IDs, see the following table.
	//
	// >  Different cloud services are available in different regions. The specified cloud service must be available in the specified region.
	//
	// example:
	//
	// cn-beijing
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The route name. You can use this parameter if you set ResourceProduct to mse.
	//
	// example:
	//
	// mse-default-traffic
	ResourceRouteName *string `json:"ResourceRouteName,omitempty" xml:"ResourceRouteName,omitempty"`
}

func (s DescribeCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudResourcesRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeCloudResourcesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudResourcesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudResourcesRequest) GetResourceDomain() *string {
	return s.ResourceDomain
}

func (s *DescribeCloudResourcesRequest) GetResourceFunction() *string {
	return s.ResourceFunction
}

func (s *DescribeCloudResourcesRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeCloudResourcesRequest) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeCloudResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCloudResourcesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCloudResourcesRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeCloudResourcesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeCloudResourcesRequest) GetResourceRouteName() *string {
	return s.ResourceRouteName
}

func (s *DescribeCloudResourcesRequest) SetInstanceId(v string) *DescribeCloudResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetOwnerUserId(v string) *DescribeCloudResourcesRequest {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetPageNumber(v int64) *DescribeCloudResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetPageSize(v int64) *DescribeCloudResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetRegionId(v string) *DescribeCloudResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceDomain(v string) *DescribeCloudResourcesRequest {
	s.ResourceDomain = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceFunction(v string) *DescribeCloudResourcesRequest {
	s.ResourceFunction = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceInstanceId(v string) *DescribeCloudResourcesRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceInstanceName(v string) *DescribeCloudResourcesRequest {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeCloudResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceName(v string) *DescribeCloudResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceProduct(v string) *DescribeCloudResourcesRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceRegionId(v string) *DescribeCloudResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceRouteName(v string) *DescribeCloudResourcesRequest {
	s.ResourceRouteName = &v
	return s
}

func (s *DescribeCloudResourcesRequest) Validate() error {
	return dara.Validate(s)
}
