// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitoringAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *InstallMonitoringAgentRequest
	GetForce() *bool
	SetInstallCommand(v string) *InstallMonitoringAgentRequest
	GetInstallCommand() *string
	SetInstanceIds(v []*string) *InstallMonitoringAgentRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *InstallMonitoringAgentRequest
	GetRegionId() *string
}

type InstallMonitoringAgentRequest struct {
	// Specifies whether to install the CloudMonitor agent. Valid values:
	//
	// 	- true (default value): yes
	//
	// 	- false: no
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Specifies whether to install the CloudMonitor agent on all ECS instances that belong to the current Alibaba Cloud account. Valid values:
	//
	// 	- `onlyInstallNotHasAgent`: installs the latest version of the CloudMonitor agent only on ECS instances on which the agent is not installed.
	//
	// 	- `onlyUpgradeAgent`: upgrades the CloudMonitor agent to the latest version only for ECS instances on which an earlier version of the agent is installed.
	//
	// 	- `installAndUpgrade`: installs the latest version of the CloudMonitor agent on ECS instances on which the agent is not installed, and upgrades the CloudMonitor agent to the latest version for ECS instances on which an earlier version of the agent is installed.
	//
	// >  If you set the InstallCommand parameter, the `InstanceIds` parameter does not take effect.
	//
	// example:
	//
	// onlyInstallNotHasAgent
	InstallCommand *string `json:"InstallCommand,omitempty" xml:"InstallCommand,omitempty"`
	// example:
	//
	// i-m5e0k0bexac8tykr****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallMonitoringAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallMonitoringAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallMonitoringAgentRequest) GetForce() *bool {
	return s.Force
}

func (s *InstallMonitoringAgentRequest) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *InstallMonitoringAgentRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *InstallMonitoringAgentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallMonitoringAgentRequest) SetForce(v bool) *InstallMonitoringAgentRequest {
	s.Force = &v
	return s
}

func (s *InstallMonitoringAgentRequest) SetInstallCommand(v string) *InstallMonitoringAgentRequest {
	s.InstallCommand = &v
	return s
}

func (s *InstallMonitoringAgentRequest) SetInstanceIds(v []*string) *InstallMonitoringAgentRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallMonitoringAgentRequest) SetRegionId(v string) *InstallMonitoringAgentRequest {
	s.RegionId = &v
	return s
}

func (s *InstallMonitoringAgentRequest) Validate() error {
	return dara.Validate(s)
}
