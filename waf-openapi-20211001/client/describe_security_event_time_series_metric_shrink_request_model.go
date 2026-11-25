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
	// The query filter conditions. Multiple filter conditions are connected by the logical AND operator.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric used to distinguish the returned data content. Different metrics correspond to different data content. This operation supports the following metrics:
	//
	// - mitigated_requests: returns the time series statistics of blocked requests.
	//
	// - monitored_requests: returns the time series statistics of requests that only hit monitor rules.
	//
	// - mitigated_requests_group_by_defense_scene: returns data grouped by protection module, recording the time series statistics of hits for each module. A single request may hit multiple modules, so the number of module hits returned by this metric may not match the number of requests.
	//
	// - mitigated_requests_group_by_block_defense_scene: returns data grouped by protection module, recording the time series statistics of blocked requests for each module. A single request is blocked by only one module, so the statistical count returned by this metric is consistent with the number of requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// mitigated_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
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
