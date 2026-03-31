// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveDetectionResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveDetectionResultRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveDetectionResultRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveDetectionResultRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSensitiveDetectionResultRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveDetectionResultRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeSensitiveDetectionResultRequest
	GetStartTime() *int64
}

type DescribeSensitiveDetectionResultRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
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
	// waf-cn-7pp26f1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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
	// The beginning of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1672502400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSensitiveDetectionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveDetectionResultRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveDetectionResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveDetectionResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveDetectionResultRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveDetectionResultRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveDetectionResultRequest) SetClusterId(v string) *DescribeSensitiveDetectionResultRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveDetectionResultRequest) SetEndTime(v int64) *DescribeSensitiveDetectionResultRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveDetectionResultRequest) SetInstanceId(v string) *DescribeSensitiveDetectionResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveDetectionResultRequest) SetRegionId(v string) *DescribeSensitiveDetectionResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveDetectionResultRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveDetectionResultRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveDetectionResultRequest) SetStartTime(v int64) *DescribeSensitiveDetectionResultRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveDetectionResultRequest) Validate() error {
	return dara.Validate(s)
}
