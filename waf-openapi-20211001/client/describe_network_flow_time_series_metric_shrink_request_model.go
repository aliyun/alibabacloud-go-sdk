// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTimeSeriesMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest
	GetInstanceId() *string
	SetMetric(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeNetworkFlowTimeSeriesMetricShrinkRequest struct {
	// The filter conditions for the query. Multiple filter conditions have a logical AND relationship.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric to query. Different metrics return different data. Valid values:
	//
	// - qps: The number of requests that WAF processes per second. This is a peak value. WAF calculates a queries per second (QPS) value every 10 seconds and uses the peak value within the specified time granularity as the QPS.
	//
	// - total_requests: The total number of requests that WAF processes.
	//
	// - top5_status: The top five response codes that WAF returns to the client and the corresponding time series statistics.
	//
	// - top5_upstream_status: The top five response codes that the origin server returns to the client and the corresponding time series statistics.
	//
	// This parameter is required.
	//
	// example:
	//
	// total_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region of the WAF instance. Valid values:
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

func (s DescribeNetworkFlowTimeSeriesMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) SetFilterShrink(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) SetInstanceId(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) SetMetric(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest {
	s.Metric = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) SetRegionId(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTimeSeriesMetricShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
