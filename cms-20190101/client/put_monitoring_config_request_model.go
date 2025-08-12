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
	SetEnableInstallAgentNewECS(v bool) *PutMonitoringConfigRequest
	GetEnableInstallAgentNewECS() *bool
	SetRegionId(v string) *PutMonitoringConfigRequest
	GetRegionId() *string
}

type PutMonitoringConfigRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// Specifies whether to automatically install the CloudMonitor agent on new ECS instances. Valid values:
	//
	// 	- true (default): The CloudMonitor agent is automatically installed on new ECS instances.
	//
	// 	- false: The CloudMonitor agent is not automatically installed on new ECS instances.
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
