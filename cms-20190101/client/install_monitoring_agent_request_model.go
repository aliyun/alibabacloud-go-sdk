// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitoringAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InstallMonitoringAgentRequest
	GetClientToken() *string
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain a maximum of 64 ASCII characters and cannot contain non-ASCII characters.
	//
	// example:
	//
	// 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. The dry run includes permission and instance status checks. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the name and description of the route table are modified.
	//
	// - **true**: sends a dry run request. Only a check is performed, and the name and description of the route table are not modified. Check items include whether required parameters are specified and the request format. If the request fails the check, the corresponding error is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcibly install the CloudMonitor agent. Valid values:
	//
	// - true (default): forcibly installs the agent.
	//
	// - false: does not forcibly install the agent.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Installs the CloudMonitor agent on all Alibaba Cloud hosts within the current Alibaba Cloud account. Valid values:
	//
	// - onlyInstallNotHasAgent: installs the latest version of the agent only on Alibaba Cloud hosts on which the CloudMonitor agent is not installed.
	//
	// - onlyUpgradeAgent: upgrades the agent only on Alibaba Cloud hosts on which a non-latest version of the CloudMonitor agent is installed.
	//
	// - installAndUpgrade: installs the latest version of the agent on Alibaba Cloud hosts on which the CloudMonitor agent is not installed, and upgrades the agent on Alibaba Cloud hosts on which a non-latest version of the CloudMonitor agent is installed.
	//
	// > If you specify this parameter, the InstanceIds parameter is ignored.
	//
	// example:
	//
	// onlyInstallNotHasAgent
	InstallCommand *string `json:"InstallCommand,omitempty" xml:"InstallCommand,omitempty"`
	// The IDs of the Alibaba Cloud hosts.
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

func (s *InstallMonitoringAgentRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *InstallMonitoringAgentRequest) SetClientToken(v string) *InstallMonitoringAgentRequest {
	s.ClientToken = &v
	return s
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
