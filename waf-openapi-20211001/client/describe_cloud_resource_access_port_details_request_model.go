// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceAccessPortDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeCloudResourceAccessPortDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudResourceAccessPortDetailsRequest
	GetPageSize() *int32
	SetPort(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetPort() *string
	SetProtocol(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetProtocol() *string
	SetRegionId(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceProduct(v string) *DescribeCloudResourceAccessPortDetailsRequest
	GetResourceProduct() *string
}

type DescribeCloudResourceAccessPortDetailsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port of the cloud service that is added to WAF.
	//
	// example:
	//
	// 443
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-2zeugkfj81jvo****4tqm
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The cloud service. Valid values:
	//
	// 	- **clb4**: Layer 4 CLB.
	//
	// 	- **clb7**: Layer 7 CLB.
	//
	// 	- **ecs**: ECS.
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s DescribeCloudResourceAccessPortDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetInstanceId(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetPageNumber(v int32) *DescribeCloudResourceAccessPortDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetPageSize(v int32) *DescribeCloudResourceAccessPortDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetPort(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.Port = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetProtocol(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetRegionId(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetResourceInstanceId(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) SetResourceProduct(v string) *DescribeCloudResourceAccessPortDetailsRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsRequest) Validate() error {
	return dara.Validate(s)
}
