// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAbnormalDomainStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecAbnormalDomainStatisticRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeApisecAbnormalDomainStatisticRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeApisecAbnormalDomainStatisticRequest
	GetInstanceId() *string
	SetOrderWay(v string) *DescribeApisecAbnormalDomainStatisticRequest
	GetOrderWay() *string
	SetPageNumber(v int64) *DescribeApisecAbnormalDomainStatisticRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecAbnormalDomainStatisticRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecAbnormalDomainStatisticRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecAbnormalDomainStatisticRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeApisecAbnormalDomainStatisticRequest
	GetStartTime() *int64
}

type DescribeApisecAbnormalDomainStatisticRequest struct {
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
	// 1687313820
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-45919n***
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
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1682571600
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApisecAbnormalDomainStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalDomainStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetClusterId(v string) *DescribeApisecAbnormalDomainStatisticRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetEndTime(v int64) *DescribeApisecAbnormalDomainStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetInstanceId(v string) *DescribeApisecAbnormalDomainStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetOrderWay(v string) *DescribeApisecAbnormalDomainStatisticRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetPageNumber(v int64) *DescribeApisecAbnormalDomainStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetPageSize(v int64) *DescribeApisecAbnormalDomainStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetRegionId(v string) *DescribeApisecAbnormalDomainStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecAbnormalDomainStatisticRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) SetStartTime(v int64) *DescribeApisecAbnormalDomainStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticRequest) Validate() error {
	return dara.Validate(s)
}
