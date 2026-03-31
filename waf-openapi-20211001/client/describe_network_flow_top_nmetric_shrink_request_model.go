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
	// An array of filter conditions. Multiple filter parameters use AND logic.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The Web Application Firewall (WAF) instance ID.
	//
	// >  Call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to retrieve the WAF instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Returns up to 10 data entries, sorted in descending order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Specifies the data type to be returned. Valid values:
	//
	// 	- real_client_ip: The top N requests, sorted in descending order by source IP address, aggregated from all the current user\\"s WAF requests.
	//
	// 	- request_path: The top N requests, sorted in descending order by user-agent, aggregated from all the current user\\"s WAF requests.
	//
	// 	- request_path: The top N requests, sorted in descending order by request URL, aggregated from all the current user\\"s WAF requests.
	//
	// 	- matched_host_by_total_requests: The top N protected objects and their request counts for the current user.
	//
	// 	- matched_host_by_qps: The top N protected objects and their queries per second (QPS) values.
	//
	// 	- matched_host_by_status: When using it, you must specify status in the Conditions field of the Filter parameter. If the HTTP response code returned by WAF matches the status specified in the Conditions, then the top N data is returned, sorted in descending order by protected objects. The format for specifying the status is as follows:\\
	//
	//     {"Key":"status","OpValue":"eq","Values":"200"}
	//
	// 	- matched_host_by_upstream_status: When using it, you must specify upstream_status in the Conditions field of the Filter parameter. If the HTTP response code returned by the origin server matches the upstream_status specified, the top N data is returned, sorted in descending order by protected objects. The format for specifying the upstream_status is as follows:\\
	//
	//     {"Key":"upstream_status","OpValue":"eq","Values":"200"}
	//
	// This parameter is required.
	//
	// example:
	//
	// matched_host_by_upstream_status
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
	// The resource group ID.
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
