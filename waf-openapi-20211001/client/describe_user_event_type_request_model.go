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
	SetEventScope(v string) *DescribeUserEventTypeRequest
	GetEventScope() *string
	SetInstanceId(v string) *DescribeUserEventTypeRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserEventTypeRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserEventTypeRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeUserEventTypeRequest
	GetStartTime() *int64
	SetUserStatusList(v []*string) *DescribeUserEventTypeRequest
	GetUserStatusList() []*string
}

type DescribeUserEventTypeRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) to obtain hybrid cloud cluster information.
	//
	// example:
	//
	// 976
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp (UTC) in seconds.
	//
	// example:
	//
	// 1726113600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension of the security event.
	//
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-5y***h0t
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-aek***ktt3y
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The start time of the query. The value is a UNIX timestamp (UTC) in seconds.
	//
	// example:
	//
	// 1723435200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of statuses for security event statistics.
	//
	// >By default, security event data in the **toBeConfirmed**, **confirmed**, and **actioned*	- statuses is included in the statistics.
	UserStatusList []*string `json:"UserStatusList,omitempty" xml:"UserStatusList,omitempty" type:"Repeated"`
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

func (s *DescribeUserEventTypeRequest) GetEventScope() *string {
	return s.EventScope
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

func (s *DescribeUserEventTypeRequest) GetUserStatusList() []*string {
	return s.UserStatusList
}

func (s *DescribeUserEventTypeRequest) SetClusterId(v string) *DescribeUserEventTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetEndTime(v int64) *DescribeUserEventTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserEventTypeRequest) SetEventScope(v string) *DescribeUserEventTypeRequest {
	s.EventScope = &v
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

func (s *DescribeUserEventTypeRequest) SetUserStatusList(v []*string) *DescribeUserEventTypeRequest {
	s.UserStatusList = v
	return s
}

func (s *DescribeUserEventTypeRequest) Validate() error {
	return dara.Validate(s)
}
