// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTopNMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *DescribeNetworkFlowTopNMetricShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeNetworkFlowTopNMetricShrinkRequest
	GetInstanceId() *string
	SetLimit(v int64) *DescribeNetworkFlowTopNMetricShrinkRequest
	GetLimit() *int64
	SetMetric(v string) *DescribeNetworkFlowTopNMetricShrinkRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeNetworkFlowTopNMetricShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTopNMetricShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeNetworkFlowTopNMetricShrinkRequest struct {
	// The filter conditions for the query. If you specify multiple filter conditions, all conditions must be met (logical AND).
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Results are sorted in descending order. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The metric by which to query and rank data. Valid values:
	//
	// - real_client_ip: Returns the top N source IP addresses of requests that are sent to WAF. The data is aggregated by source IP address and sorted in descending order.
	//
	// - http_user_agent: Returns the top N user agents of requests that are sent to WAF. The data is aggregated by user agent and sorted in descending order.
	//
	// - request_path: Returns the top N request URLs. The data is aggregated by URL and sorted in descending order.
	//
	// - matched_host_by_total_requests: Returns the top N protected objects by total number of requests received.
	//
	// - matched_host_by_qps: Returns the top N protected objects by queries per second (QPS).
	//
	// - matched_host_by_status: Returns the top N protected objects based on a specific WAF-returned HTTP status code. The data is aggregated by protected object and sorted in descending order. If you set Metric to this value, you must specify the status field in the Conditions parameter of Filter. The format is as follows:<br> {"Key":"status","OpValue":"eq","Values":"200"}<br>
	//
	// - matched_host_by_upstream_status: Returns the top N protected objects based on a specific origin server-returned HTTP status code. The data is aggregated by protected object and sorted in descending order. If you set Metric to this value, you must specify the upstream_status field in the Conditions parameter of Filter. The format is as follows:<br> {"Key":"upstream_status","OpValue":"eq","Values":"200"}<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// matched_host_by_upstream_status
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region in which the WAF instance resides. Valid values:
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

func (s DescribeNetworkFlowTopNMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) SetFilterShrink(v string) *DescribeNetworkFlowTopNMetricShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) SetInstanceId(v string) *DescribeNetworkFlowTopNMetricShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) SetLimit(v int64) *DescribeNetworkFlowTopNMetricShrinkRequest {
	s.Limit = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) SetMetric(v string) *DescribeNetworkFlowTopNMetricShrinkRequest {
	s.Metric = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) SetRegionId(v string) *DescribeNetworkFlowTopNMetricShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTopNMetricShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
