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
	// The query filter conditions. Multiple filter conditions are combined by using the logical AND operator.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric used to distinguish the returned data content. Different metrics correspond to different data content. The metrics supported by this operation are as follows:
	//
	// 	- qps: returns the number of requests processed by WAF per second. The peak calculation method is used. A QPS value is calculated every 10 seconds, and the Peak Volume among multiple values is used as the QPS value within the current statistical time granularity.
	//
	// 	- total_requests: returns the total number of requests processed by WAF.
	//
	// 	- top5_status: returns the top 5 response codes returned by WAF to the client and the corresponding time sequence statistics.
	//
	// 	- top 5_upstream_status: returns the top 5 back-to-origin response codes returned by the origin server to the client and the corresponding time sequence statistics.
	//
	// This parameter is required.
	//
	// example:
	//
	// total_requests
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
