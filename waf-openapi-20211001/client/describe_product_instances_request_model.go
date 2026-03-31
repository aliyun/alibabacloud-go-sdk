// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeProductInstancesRequest
	GetInstanceId() *string
	SetOwnerUserId(v string) *DescribeProductInstancesRequest
	GetOwnerUserId() *string
	SetPageNumber(v int64) *DescribeProductInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeProductInstancesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeProductInstancesRequest
	GetRegionId() *string
	SetResourceInstanceAccessStatus(v string) *DescribeProductInstancesRequest
	GetResourceInstanceAccessStatus() *string
	SetResourceInstanceId(v string) *DescribeProductInstancesRequest
	GetResourceInstanceId() *string
	SetResourceInstanceIp(v string) *DescribeProductInstancesRequest
	GetResourceInstanceIp() *string
	SetResourceInstanceName(v string) *DescribeProductInstancesRequest
	GetResourceInstanceName() *string
	SetResourceIp(v string) *DescribeProductInstancesRequest
	GetResourceIp() *string
	SetResourceManagerResourceGroupId(v string) *DescribeProductInstancesRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceName(v string) *DescribeProductInstancesRequest
	GetResourceName() *string
	SetResourceProduct(v string) *DescribeProductInstancesRequest
	GetResourceProduct() *string
	SetResourceRegionId(v string) *DescribeProductInstancesRequest
	GetResourceRegionId() *string
}

type DescribeProductInstancesRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zxu****9d02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 1704********9107
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
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceInstanceAccessStatus *string `json:"ResourceInstanceAccessStatus,omitempty" xml:"ResourceInstanceAccessStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// lb-2zeugkfj81jvo****4tqm
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The IP address of the instance that is added to WAF.
	//
	// example:
	//
	// 1.X.X.1
	ResourceInstanceIp *string `json:"ResourceInstanceIp,omitempty" xml:"ResourceInstanceIp,omitempty"`
	// The name of the instance that is added to WAF.
	//
	// example:
	//
	// demoInstanceName
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// Deprecated
	//
	// The public IP address of the instance.
	//
	// example:
	//
	// 1.X.X.1
	ResourceIp *string `json:"ResourceIp,omitempty" xml:"ResourceIp,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekz6ql****5uzi
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Deprecated
	//
	// The name of the instance.
	//
	// example:
	//
	// exampleResourceName
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The cloud service to which the instance belongs. Valid values:
	//
	// 	- **clb4**: Layer 4 Classic Load Balancer (CLB).
	//
	// 	- **clb7**: Layer 7 CLB.
	//
	// 	- **ecs**: Elastic Compute Service (ECS).
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the instance. Valid values:
	//
	// 	- **cn-chengdu**: China (Chengdu).
	//
	// 	- **cn-beijing**: China (Beijing).
	//
	// 	- **cn-zhangjiakou**: China (Zhangjiakou).
	//
	// 	- **cn-hangzhou**: China (Hangzhou).
	//
	// 	- **cn-shanghai**: China (Shanghai).
	//
	// 	- **cn-shenzhen**: China (Shenzhen).
	//
	// 	- **cn-qingdao**: China (Qingdao).
	//
	// 	- **cn-hongkong**: China (Hong Kong).
	//
	// 	- **ap-southeast-3**: Malaysia (Kuala Lumpur).
	//
	// 	- **ap-southeast-5**: Indonesia (Jakarta).
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeProductInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProductInstancesRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeProductInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeProductInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeProductInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProductInstancesRequest) GetResourceInstanceAccessStatus() *string {
	return s.ResourceInstanceAccessStatus
}

func (s *DescribeProductInstancesRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeProductInstancesRequest) GetResourceInstanceIp() *string {
	return s.ResourceInstanceIp
}

func (s *DescribeProductInstancesRequest) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeProductInstancesRequest) GetResourceIp() *string {
	return s.ResourceIp
}

func (s *DescribeProductInstancesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeProductInstancesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeProductInstancesRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeProductInstancesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeProductInstancesRequest) SetInstanceId(v string) *DescribeProductInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetOwnerUserId(v string) *DescribeProductInstancesRequest {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetPageNumber(v int64) *DescribeProductInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetPageSize(v int64) *DescribeProductInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetRegionId(v string) *DescribeProductInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceInstanceAccessStatus(v string) *DescribeProductInstancesRequest {
	s.ResourceInstanceAccessStatus = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceInstanceId(v string) *DescribeProductInstancesRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceInstanceIp(v string) *DescribeProductInstancesRequest {
	s.ResourceInstanceIp = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceInstanceName(v string) *DescribeProductInstancesRequest {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceIp(v string) *DescribeProductInstancesRequest {
	s.ResourceIp = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceManagerResourceGroupId(v string) *DescribeProductInstancesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceName(v string) *DescribeProductInstancesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceProduct(v string) *DescribeProductInstancesRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceRegionId(v string) *DescribeProductInstancesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeProductInstancesRequest) Validate() error {
	return dara.Validate(s)
}
