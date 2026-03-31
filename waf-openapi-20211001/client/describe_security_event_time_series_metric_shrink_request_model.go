// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTimeSeriesMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest
	GetInstanceId() *string
	SetMetric(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSecurityEventTimeSeriesMetricShrinkRequest struct {
	// The filter conditions for the query. Multiple conditions are evaluated by using a logical AND.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric whose time series data you want to return. The following metrics are supported:
	//
	// 	- mitigated_requests: The system returns the time series data of requests that are blocked.
	//
	// 	- monitored_requests: The system returns the time series data of requests that match Monitor protection rules.
	//
	// 	- mitigated_requests_group_by_defense_scene: The system returns the number of requests that match each protection module. The returned results are grouped by protection module and can be used to generate time series charts. A request can match multiple protection modules. Therefore, the total number of matched requests is inconsistent with the total number of requests.
	//
	// 	- mitigated_requests_group_by_block_defense_scene: The system returns the number of requests that are blocked by each protection module. The returned results are grouped by protection module and can be used to generate time series charts. A request can be blocked by only one protection module. Therefore, the total number of blocked requests is consistent with the total number of requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// mitigated_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: The Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) SetFilterShrink(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) SetInstanceId(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) SetMetric(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest {
	s.Metric = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) SetRegionId(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTimeSeriesMetricShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
