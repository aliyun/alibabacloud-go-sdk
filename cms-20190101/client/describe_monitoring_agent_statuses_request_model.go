// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAvailabilityTaskId(v string) *DescribeMonitoringAgentStatusesRequest
	GetHostAvailabilityTaskId() *string
	SetInstanceIds(v string) *DescribeMonitoringAgentStatusesRequest
	GetInstanceIds() *string
	SetRegionId(v string) *DescribeMonitoringAgentStatusesRequest
	GetRegionId() *string
}

type DescribeMonitoringAgentStatusesRequest struct {
	// The ID of the availability monitoring task.
	//
	// example:
	//
	// 126****
	HostAvailabilityTaskId *string `json:"HostAvailabilityTaskId,omitempty" xml:"HostAvailabilityTaskId,omitempty"`
	// The instance IDs. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// i-hp3dunahluwajv6f****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitoringAgentStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesRequest) GetHostAvailabilityTaskId() *string {
	return s.HostAvailabilityTaskId
}

func (s *DescribeMonitoringAgentStatusesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeMonitoringAgentStatusesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitoringAgentStatusesRequest) SetHostAvailabilityTaskId(v string) *DescribeMonitoringAgentStatusesRequest {
	s.HostAvailabilityTaskId = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesRequest) SetInstanceIds(v string) *DescribeMonitoringAgentStatusesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesRequest) SetRegionId(v string) *DescribeMonitoringAgentStatusesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesRequest) Validate() error {
	return dara.Validate(s)
}
