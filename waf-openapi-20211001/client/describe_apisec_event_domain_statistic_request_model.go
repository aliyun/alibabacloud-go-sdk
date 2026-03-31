// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventDomainStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecEventDomainStatisticRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeApisecEventDomainStatisticRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeApisecEventDomainStatisticRequest
	GetInstanceId() *string
	SetOrderWay(v string) *DescribeApisecEventDomainStatisticRequest
	GetOrderWay() *string
	SetPageNumber(v int64) *DescribeApisecEventDomainStatisticRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecEventDomainStatisticRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecEventDomainStatisticRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecEventDomainStatisticRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeApisecEventDomainStatisticRequest
	GetStartTime() *int64
}

type DescribeApisecEventDomainStatisticRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1686895256
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sorting order. Valid values:
	//
	// - **asc**: ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **5**.
	//
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmvyknl****fa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1668496310
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApisecEventDomainStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDomainStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDomainStatisticRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecEventDomainStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApisecEventDomainStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecEventDomainStatisticRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeApisecEventDomainStatisticRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecEventDomainStatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecEventDomainStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecEventDomainStatisticRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecEventDomainStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeApisecEventDomainStatisticRequest) SetClusterId(v string) *DescribeApisecEventDomainStatisticRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetEndTime(v int64) *DescribeApisecEventDomainStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetInstanceId(v string) *DescribeApisecEventDomainStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetOrderWay(v string) *DescribeApisecEventDomainStatisticRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetPageNumber(v int64) *DescribeApisecEventDomainStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetPageSize(v int64) *DescribeApisecEventDomainStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetRegionId(v string) *DescribeApisecEventDomainStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecEventDomainStatisticRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) SetStartTime(v int64) *DescribeApisecEventDomainStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticRequest) Validate() error {
	return dara.Validate(s)
}
