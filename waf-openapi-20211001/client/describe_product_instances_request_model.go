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
	SetResourceDomain(v string) *DescribeProductInstancesRequest
	GetResourceDomain() *string
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
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zxu****9d02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UID of the resource ownership user.
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
	// The number of entries per page when paging. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The domain name that is added to WAF.
	//
	// > This parameter is supported only when the cloud service type is ddos.
	//
	// example:
	//
	// www.c**sw.net
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The WAF protection status.
	//
	// example:
	//
	// all
	ResourceInstanceAccessStatus *string `json:"ResourceInstanceAccessStatus,omitempty" xml:"ResourceInstanceAccessStatus,omitempty"`
	// The instance ID of the cloud service.
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
	// The public IP address of the cloud service.
	//
	// example:
	//
	// 1.X.X.1
	ResourceIp *string `json:"ResourceIp,omitempty" xml:"ResourceIp,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-aekz6ql****5uzi
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Deprecated
	//
	// The instance name of the cloud service.
	//
	// example:
	//
	// exampleResourceName
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the cloud service. Valid values:
	//
	// - **clb4**: Layer 4 CLB.
	//
	// - **clb7**: Layer 7 CLB.
	//
	// - **ecs**: ECS.
	//
	// - **nlb**: NLB.
	//
	// - **ddos**: Anti-DDoS.
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the cloud service. Valid values:
	//
	// - **cn-chengdu**: China Southwest 1 (Chengdu).
	//
	// - **cn-beijing**: China North 2 (Beijing).
	//
	// - **cn-zhangjiakou**: China North 3 (Zhangjiakou).
	//
	// - **cn-hangzhou**: China East 1 (Hangzhou).
	//
	// - **cn-shanghai**: China East 2 (Shanghai).
	//
	// - **cn-shenzhen**: China South 1 (Shenzhen).
	//
	// - **cn-qingdao**: China North 1 (Qingdao).
	//
	// - **cn-hongkong**: Hong Kong (China).
	//
	// - **ap-southeast-3**: Malaysia (Kuala Lumpur).
	//
	// - **ap-southeast-5**: Indonesia (Jakarta).
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

func (s *DescribeProductInstancesRequest) GetResourceDomain() *string {
	return s.ResourceDomain
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

func (s *DescribeProductInstancesRequest) SetResourceDomain(v string) *DescribeProductInstancesRequest {
	s.ResourceDomain = &v
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
