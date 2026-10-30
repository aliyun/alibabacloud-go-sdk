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
	SetDryRun(v bool) *PutMonitoringConfigRequest
	GetDryRun() *bool
	SetEnableInstallAgentNewECS(v bool) *PutMonitoringConfigRequest
	GetEnableInstallAgentNewECS() *bool
	SetRegionId(v string) *PutMonitoringConfigRequest
	GetRegionId() *string
}

type PutMonitoringConfigRequest struct {
	// This parameter is not required.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	DryRun      *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to automatically install the CloudMonitor agent on newly purchased ECS instances. Valid values:
	//
	// - true (default): The CloudMonitor agent is automatically installed on newly purchased ECS instances.
	//
	// - false: The CloudMonitor agent is not automatically installed on newly purchased ECS instances.
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
