// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEventTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeUserEventTypeRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeUserEventTypeRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeUserEventTypeRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserEventTypeRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserEventTypeRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeUserEventTypeRequest
	GetStartTime() *int64
}

type DescribeUserEventTypeRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 976
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1726113600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-5y***h0t
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
	// rg-aek***ktt3y
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1723435200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserEventTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTypeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeUserEventTypeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUserEventTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserEventTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserEventTypeRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserEventTypeRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUserEventTypeRequest) SetClusterId(v string) *DescribeUserEventTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetEndTime(v int64) *DescribeUserEventTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetInstanceId(v string) *DescribeUserEventTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetRegionId(v string) *DescribeUserEventTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserEventTypeRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetStartTime(v int64) *DescribeUserEventTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserEventTypeRequest) Validate() error {
	return dara.Validate(s)
}
