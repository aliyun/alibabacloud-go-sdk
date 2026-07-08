// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventTopMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeThreatEventTopMetricRequest
	GetEventId() *string
	SetInstanceId(v string) *DescribeThreatEventTopMetricRequest
	GetInstanceId() *string
	SetMetric(v string) *DescribeThreatEventTopMetricRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeThreatEventTopMetricRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeThreatEventTopMetricRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeThreatEventTopMetricRequest struct {
	// The ID of the security event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0b7ab137a065aab7656986***11db
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The statistical metric. Different values of Metric correspond to different statistical objects. Valid values:
	//
	// - **time**: aggregates statistics by attack time, sorts them in descending order, and returns the top 5 records.
	//
	// - **src**: aggregates statistics by source IP address of attack requests, sorts them in descending order, and returns the top 5 records.
	//
	// - **target**: aggregates statistics by URL of attack requests (excluding query strings), sorts them in descending order, and returns the top 5 records.
	//
	// - **type**: aggregates statistics by attack type, sorts them in descending order, and returns the top 5 records.
	//
	// - **tools**: aggregates statistics by attack tool, sorts them in descending order, and returns the top 5 records.
	//
	// This parameter is required.
	//
	// example:
	//
	// time
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
}

func (s DescribeThreatEventTopMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventTopMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventTopMetricRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeThreatEventTopMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeThreatEventTopMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeThreatEventTopMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeThreatEventTopMetricRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeThreatEventTopMetricRequest) SetEventId(v string) *DescribeThreatEventTopMetricRequest {
	s.EventId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetInstanceId(v string) *DescribeThreatEventTopMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetMetric(v string) *DescribeThreatEventTopMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetRegionId(v string) *DescribeThreatEventTopMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetResourceManagerResourceGroupId(v string) *DescribeThreatEventTopMetricRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) Validate() error {
	return dara.Validate(s)
}
