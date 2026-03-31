// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSensitiveDomainStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecSensitiveDomainStatisticRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeApisecSensitiveDomainStatisticRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeApisecSensitiveDomainStatisticRequest
	GetInstanceId() *string
	SetOrderWay(v string) *DescribeApisecSensitiveDomainStatisticRequest
	GetOrderWay() *string
	SetPageNumber(v int64) *DescribeApisecSensitiveDomainStatisticRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecSensitiveDomainStatisticRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecSensitiveDomainStatisticRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecSensitiveDomainStatisticRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeApisecSensitiveDomainStatisticRequest
	GetStartTime() *int64
	SetType(v string) *DescribeApisecSensitiveDomainStatisticRequest
	GetType() *string
}

type DescribeApisecSensitiveDomainStatisticRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1686895256
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sorting order. Valid values:
	//
	// -  **asc**: ascending order.
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
	// 	- **cn-hangzhou**: Chinese mainland.
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
	// The beginning of the time range to query. Specify a UNIX timestamp in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1668496310
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The sensitive data type. Valid values:
	//
	// - **request**: sensitive data in requests.
	//
	// - **response**: sensitive data in responses.
	//
	// example:
	//
	// request
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApisecSensitiveDomainStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSensitiveDomainStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetClusterId(v string) *DescribeApisecSensitiveDomainStatisticRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetEndTime(v int64) *DescribeApisecSensitiveDomainStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetInstanceId(v string) *DescribeApisecSensitiveDomainStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetOrderWay(v string) *DescribeApisecSensitiveDomainStatisticRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetPageNumber(v int64) *DescribeApisecSensitiveDomainStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetPageSize(v int64) *DescribeApisecSensitiveDomainStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetRegionId(v string) *DescribeApisecSensitiveDomainStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecSensitiveDomainStatisticRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetStartTime(v int64) *DescribeApisecSensitiveDomainStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) SetType(v string) *DescribeApisecSensitiveDomainStatisticRequest {
	s.Type = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticRequest) Validate() error {
	return dara.Validate(s)
}
