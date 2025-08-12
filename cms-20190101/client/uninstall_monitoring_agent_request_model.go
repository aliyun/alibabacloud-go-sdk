// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallMonitoringAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UninstallMonitoringAgentRequest
	GetInstanceId() *string
	SetRegionId(v string) *UninstallMonitoringAgentRequest
	GetRegionId() *string
}

type UninstallMonitoringAgentRequest struct {
	// The ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// host-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UninstallMonitoringAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallMonitoringAgentRequest) GoString() string {
	return s.String()
}

func (s *UninstallMonitoringAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UninstallMonitoringAgentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UninstallMonitoringAgentRequest) SetInstanceId(v string) *UninstallMonitoringAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *UninstallMonitoringAgentRequest) SetRegionId(v string) *UninstallMonitoringAgentRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallMonitoringAgentRequest) Validate() error {
	return dara.Validate(s)
}
