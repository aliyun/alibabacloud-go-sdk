// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeMonitoringAgentProcessesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeMonitoringAgentProcessesRequest
	GetRegionId() *string
}

type DescribeMonitoringAgentProcessesRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-hp3hl3cx1pbahzy8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitoringAgentProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitoringAgentProcessesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitoringAgentProcessesRequest) SetInstanceId(v string) *DescribeMonitoringAgentProcessesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesRequest) SetRegionId(v string) *DescribeMonitoringAgentProcessesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesRequest) Validate() error {
	return dara.Validate(s)
}
