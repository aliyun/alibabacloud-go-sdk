// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitoringAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteMonitoringAgentProcessRequest
	GetInstanceId() *string
	SetProcessId(v string) *DeleteMonitoringAgentProcessRequest
	GetProcessId() *string
	SetProcessName(v string) *DeleteMonitoringAgentProcessRequest
	GetProcessName() *string
	SetRegionId(v string) *DeleteMonitoringAgentProcessRequest
	GetRegionId() *string
}

type DeleteMonitoringAgentProcessRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-KpVny6l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The process ID.
	//
	// > You must configure either `ProcessId` or `ProcessName`.
	//
	// example:
	//
	// 123****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The process name.
	//
	// > You must configure either `ProcessId` or `ProcessName`.
	//
	// example:
	//
	// http
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMonitoringAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitoringAgentProcessRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteMonitoringAgentProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DeleteMonitoringAgentProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *DeleteMonitoringAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMonitoringAgentProcessRequest) SetInstanceId(v string) *DeleteMonitoringAgentProcessRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMonitoringAgentProcessRequest) SetProcessId(v string) *DeleteMonitoringAgentProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *DeleteMonitoringAgentProcessRequest) SetProcessName(v string) *DeleteMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *DeleteMonitoringAgentProcessRequest) SetRegionId(v string) *DeleteMonitoringAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMonitoringAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}
