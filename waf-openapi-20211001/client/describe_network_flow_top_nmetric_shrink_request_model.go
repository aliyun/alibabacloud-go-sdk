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
	// The query filter conditions. Multiple filter conditions are connected by the logical AND operator.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The WAF instance ID.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of data entries to return after the statistics are sorted in descending order. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The metric used to distinguish the returned data content. Different metrics correspond to different data content. The metrics supported by this operation are as follows:
	//
	// - real_client_ip: aggregates and sorts all requests of the current user that pass through WAF by the source IP address of the request in descending order, and returns the top N data entries.
	//
	// - http_user_agent: aggregates and sorts all requests of the current user that pass through WAF by the User-Agent of the request in descending order, and returns the top N data entries.
	//
	// - request_path: aggregates and sorts all requests of the current user that pass through WAF by the URL of the request in descending order, and returns the top N data entries.
	//
	// - matched_host_by_total_requests: returns the top N protected objects with the largest number of requests for the current user and the corresponding number of requests.
	//
	// - matched_host_by_qps: returns the top N protected objects with the highest QPS for the current user and the corresponding QPS data.
	//
	// - matched_host_by_status: when this metric is used, you must specify status in the Conditions field of Filter. WAF aggregates and sorts the data whose HTTP response code returned by WAF is equal to the status specified in Conditions by protected object, and returns the top N data entries. The format for specifying status is as follows:<br>
	//
	// {"Key":"status","OpValue":"eq","Values":"200"}
	//
	// - matched_host_by_upstream_status: when this metric is used, you must specify upstream_status in the Conditions field of Filter. WAF aggregates and sorts the data whose HTTP response code returned by the origin server is equal to the upstream_status specified in Conditions by protected object, and returns the top N data entries. The format for specifying upstream_status is as follows:<br>
	//
	// {"Key":"upstream_status","OpValue":"eq","Values":"200"}
	//
	// This parameter is required.
	//
	// example:
	//
	// matched_host_by_upstream_status
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
