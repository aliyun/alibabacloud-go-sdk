// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitoringAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *InstallMonitoringAgentRequest
	GetDryRun() *bool
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
	// Specifies whether to perform a dry run on the request. The dry run includes checks on permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. After the request passes the check, the name and description of the route table are directly modified.
	//
	// - **true**: Sends a dry run request. Only checks are performed, and the name and description of the route table are not modified. Check items include whether required parameters are specified and the request format. If the check fails, an error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcibly install the CloudMonitor agent. Valid values:
	//
	// - true (default): Forcibly installs the agent.
	//
	// - false: Does not forcibly install the agent.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Installs the CloudMonitor agent on all Alibaba Cloud hosts within the current Alibaba Cloud account. Valid values:
	//
	// - onlyInstallNotHasAgent: Installs the latest version of the agent only on Alibaba Cloud hosts on which the CloudMonitor agent is not installed.
	//
	// - onlyUpgradeAgent: Upgrades the agent only on Alibaba Cloud hosts on which a non-latest version of the CloudMonitor agent is installed.
	//
	// - installAndUpgrade: Installs the latest version of the agent on Alibaba Cloud hosts on which the CloudMonitor agent is not installed, and upgrades the agent on Alibaba Cloud hosts on which a non-latest version of the CloudMonitor agent is installed.
	//
	// > If you specify this parameter, the InstanceIds parameter becomes invalid.
	//
	// example:
	//
	// onlyInstallNotHasAgent
	InstallCommand *string `json:"InstallCommand,omitempty" xml:"InstallCommand,omitempty"`
	// The ID of the Alibaba Cloud host.
	//
	// Valid values of N: 1 to 10.
	//
	// > You must specify either InstallCommand or InstanceIds.
	//
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

func (s *InstallMonitoringAgentRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *InstallMonitoringAgentRequest) SetDryRun(v bool) *InstallMonitoringAgentRequest {
	s.DryRun = &v
	return s
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
