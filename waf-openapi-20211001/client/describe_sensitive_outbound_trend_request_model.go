// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveOutboundTrendRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveOutboundTrendRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveOutboundTrendRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSensitiveOutboundTrendRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveOutboundTrendRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeSensitiveOutboundTrendRequest
	GetStartTime() *int64
}

type DescribeSensitiveOutboundTrendRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 433
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// >  You can query only data of the previous month, previous 3 months, previous 6 months, previous 12 months, and data generated since January 1 of last year for compliance check. You must specify a valid time range.
	//
	// example:
	//
	// 1725966000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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
	// The beginning of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// >  You can query only data of the previous month, previous 3 months, previous 6 months, previous 12 months, and data generated since January 1 of last year for compliance check. You must specify a valid time range.
	//
	// example:
	//
	// 1672502400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSensitiveOutboundTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundTrendRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveOutboundTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveOutboundTrendRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveOutboundTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveOutboundTrendRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveOutboundTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveOutboundTrendRequest) SetClusterId(v string) *DescribeSensitiveOutboundTrendRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendRequest) SetEndTime(v int64) *DescribeSensitiveOutboundTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendRequest) SetInstanceId(v string) *DescribeSensitiveOutboundTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendRequest) SetRegionId(v string) *DescribeSensitiveOutboundTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveOutboundTrendRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendRequest) SetStartTime(v int64) *DescribeSensitiveOutboundTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendRequest) Validate() error {
	return dara.Validate(s)
}
