// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveOutboundDistributionRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveOutboundDistributionRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveOutboundDistributionRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSensitiveOutboundDistributionRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveOutboundDistributionRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeSensitiveOutboundDistributionRequest
	GetStartTime() *int64
}

type DescribeSensitiveOutboundDistributionRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to obtain hybrid cloud cluster information.
	//
	// example:
	//
	// 443
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end time of the query. The time is in the UNIX timestamp format (UTC). Unit: seconds.
	//
	// > Compliance audit currently supports querying data for the last 1 month, last 3 months, last 6 months, last 12 months, and from January 1 of the previous year to the present. The query time range must fall within the supported intervals. **StartTime*	- must not be earlier than the start time of the corresponding interval. If the time range exceeds the supported range, the query fails.
	//
	// > This parameter is optional. The default value is the current time.
	//
	// example:
	//
	// 1725966000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The start time of the query. The time is in the UNIX timestamp format (UTC). Unit: seconds.
	//
	// > Compliance audit currently supports querying data for the last 1 month, last 3 months, last 6 months, last 12 months, and from January 1 of the previous year to the present. The query time range must fall within the supported intervals. **StartTime*	- must not be earlier than the start time of the corresponding interval. If the time range exceeds the supported range, the query fails.
	//
	// > This parameter is optional. The default value is one month before the current time.
	//
	// example:
	//
	// 1672502400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSensitiveOutboundDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundDistributionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundDistributionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveOutboundDistributionRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveOutboundDistributionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveOutboundDistributionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveOutboundDistributionRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveOutboundDistributionRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveOutboundDistributionRequest) SetClusterId(v string) *DescribeSensitiveOutboundDistributionRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionRequest) SetEndTime(v int64) *DescribeSensitiveOutboundDistributionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionRequest) SetInstanceId(v string) *DescribeSensitiveOutboundDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionRequest) SetRegionId(v string) *DescribeSensitiveOutboundDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveOutboundDistributionRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionRequest) SetStartTime(v int64) *DescribeSensitiveOutboundDistributionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionRequest) Validate() error {
	return dara.Validate(s)
}
