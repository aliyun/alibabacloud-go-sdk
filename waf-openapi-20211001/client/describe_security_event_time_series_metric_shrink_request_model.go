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
	// The filter conditions for the query. Multiple filter conditions have a logical AND relationship.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the content of the returned data. Different metrics correspond to different data content. This operation supports the following metrics:
	//
	// - mitigated_requests: Returns the time series statistics of blocked requests.
	//
	// - monitored_requests: Returns the time series statistics of requests that hit only observation-type rules.
	//
	// - mitigated_requests_group_by_defense_scene: Returns data grouped by module. It records a time series graph of the hit count for each module. A single request may hit multiple modules. Therefore, the hit count returned by this metric may not be consistent with the number of requests.
	//
	// - mitigated_requests_group_by_block_defense_scene: Returns data grouped by module. It records a time series graph of the number of blocked requests for each module. A single request is blocked by only one module. Therefore, the count returned by this metric is consistent with the number of requests.
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
