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
	// Specifies filtering conditions. Multiple filter parameters use AND logic.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The Web Application Firewall (WAF) instance ID.
	//
	// > Call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to retrieve the WAF instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the data type to be returned. Valid values:
	//
	// 	- qps: Returns the queries per second (QPS) processed by WAF. This value is calculated using a peak detection method: QPS is measured every 10 seconds, and the highest value within the specified time range is returned.
	//
	// 	- total_requests: Returns the total number of requests processed by WAF.
	//
	// 	- top5_status: Returns the top 5 HTTP status codes returned by the WAF to clients, along with their corresponding time series statistics.
	//
	// 	- top 5_upstream_status: Returns the top 5 HTTP status codes returned by the origin server to clients, along with their corresponding time series data.
	//
	// This parameter is required.
	//
	// example:
	//
	// total_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region ID of WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: The Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
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
