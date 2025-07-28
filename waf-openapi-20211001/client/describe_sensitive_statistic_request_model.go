// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveStatisticRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveStatisticRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveStatisticRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeSensitiveStatisticRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSensitiveStatisticRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSensitiveStatisticRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveStatisticRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeSensitiveStatisticRequest
	GetStartTime() *int64
	SetStatisticType(v string) *DescribeSensitiveStatisticRequest
	GetStatisticType() *string
}

type DescribeSensitiveStatisticRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1725966000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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
	// The beginning of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1672502400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the statistics. Valid values:
	//
	// 	- **ip**: IP address
	//
	// 	- **host**: domain name
	//
	// 	- **sensitive_code**: sensitive data type
	//
	// 	- **api**: sensitive data-related API
	//
	// example:
	//
	// ip
	StatisticType *string `json:"StatisticType,omitempty" xml:"StatisticType,omitempty"`
}

func (s DescribeSensitiveStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveStatisticRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveStatisticRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSensitiveStatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSensitiveStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveStatisticRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveStatisticRequest) GetStatisticType() *string {
	return s.StatisticType
}

func (s *DescribeSensitiveStatisticRequest) SetClusterId(v string) *DescribeSensitiveStatisticRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetEndTime(v int64) *DescribeSensitiveStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetInstanceId(v string) *DescribeSensitiveStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetPageNumber(v int64) *DescribeSensitiveStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetPageSize(v int64) *DescribeSensitiveStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetRegionId(v string) *DescribeSensitiveStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveStatisticRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetStartTime(v int64) *DescribeSensitiveStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) SetStatisticType(v string) *DescribeSensitiveStatisticRequest {
	s.StatisticType = &v
	return s
}

func (s *DescribeSensitiveStatisticRequest) Validate() error {
	return dara.Validate(s)
}
