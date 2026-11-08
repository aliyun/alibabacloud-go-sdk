// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitoringConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoInstall(v bool) *PutMonitoringConfigRequest
	GetAutoInstall() *bool
	SetClientToken(v string) *PutMonitoringConfigRequest
	GetClientToken() *string
	SetDryRun(v bool) *PutMonitoringConfigRequest
	GetDryRun() *bool
	SetEnableInstallAgentNewECS(v bool) *PutMonitoringConfigRequest
	GetEnableInstallAgentNewECS() *bool
	SetRegionId(v string) *PutMonitoringConfigRequest
	GetRegionId() *string
}

type PutMonitoringConfigRequest struct {
	// Ignore this parameter.
	//
	// example:
	//
	// true
	AutoInstall *bool   `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - **true**: Sends a dry run request without deleting the endpoint IPAM pool. The system checks whether all required parameters are specified, the request format is valid, and service limits are met. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request and directly deletes the endpoint IPAM pool if the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to automatically install the CloudMonitor agent on newly purchased ECS instances. Valid values:
	//
	// - true (default): Automatically installs the CloudMonitor agent on newly purchased ECS instances.
	//
	// - false: Does not automatically install the CloudMonitor agent on newly purchased ECS instances.
	//
	// example:
	//
	// true
	EnableInstallAgentNewECS *bool   `json:"EnableInstallAgentNewECS,omitempty" xml:"EnableInstallAgentNewECS,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutMonitoringConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutMonitoringConfigRequest) GoString() string {
	return s.String()
}

func (s *PutMonitoringConfigRequest) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *PutMonitoringConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PutMonitoringConfigRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *PutMonitoringConfigRequest) GetEnableInstallAgentNewECS() *bool {
	return s.EnableInstallAgentNewECS
}

func (s *PutMonitoringConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutMonitoringConfigRequest) SetAutoInstall(v bool) *PutMonitoringConfigRequest {
	s.AutoInstall = &v
	return s
}

func (s *PutMonitoringConfigRequest) SetClientToken(v string) *PutMonitoringConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *PutMonitoringConfigRequest) SetDryRun(v bool) *PutMonitoringConfigRequest {
	s.DryRun = &v
	return s
}

func (s *PutMonitoringConfigRequest) SetEnableInstallAgentNewECS(v bool) *PutMonitoringConfigRequest {
	s.EnableInstallAgentNewECS = &v
	return s
}

func (s *PutMonitoringConfigRequest) SetRegionId(v string) *PutMonitoringConfigRequest {
	s.RegionId = &v
	return s
}

func (s *PutMonitoringConfigRequest) Validate() error {
	return dara.Validate(s)
}
