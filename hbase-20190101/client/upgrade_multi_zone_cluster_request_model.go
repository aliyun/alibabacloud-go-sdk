// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMultiZoneClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpgradeMultiZoneClusterRequest
	GetClusterId() *string
	SetComponents(v string) *UpgradeMultiZoneClusterRequest
	GetComponents() *string
	SetRestartComponents(v string) *UpgradeMultiZoneClusterRequest
	GetRestartComponents() *string
	SetRunMode(v string) *UpgradeMultiZoneClusterRequest
	GetRunMode() *string
	SetUpgradeInsName(v string) *UpgradeMultiZoneClusterRequest
	GetUpgradeInsName() *string
	SetVersions(v string) *UpgradeMultiZoneClusterRequest
	GetVersions() *string
}

type UpgradeMultiZoneClusterRequest struct {
	// The ID of the multi-zone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-***************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The component names. You can specify multiple component names separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// LINDORM
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The names of the components that need to be restarted after the upgrade. You can specify multiple component names separated by commas (,).
	//
	// example:
	//
	// LPROXY
	RestartComponents *string `json:"RestartComponents,omitempty" xml:"RestartComponents,omitempty"`
	// The execution mode. If UpgradeInsName is not empty, the mode is forcibly set to single. If UpgradeInsName is empty and RunMode is not specified, the default value is serial. Valid values:
	//
	// - serial: all sub-instances are upgraded.
	//
	// - single: only the specified sub-instance is upgraded.
	//
	// example:
	//
	// serial
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The name of the sub-instance to upgrade. You can obtain this value from the MultiZoneInstanceModels field in the response of the [DescribeMultiZoneCluster](~~DescribeMultiZoneCluster~~) operation. This parameter is optional. If you do not specify this parameter, all sub-instances are upgraded.
	//
	// example:
	//
	// ld-t4n40m3171t4******-az-b
	UpgradeInsName *string `json:"UpgradeInsName,omitempty" xml:"UpgradeInsName,omitempty"`
	// The RPM version to upgrade to. If you do not specify this parameter, the components are upgraded to the latest version. If you specify multiple values for Components, you must also specify the same number of values for Versions, separated by commas (,).
	//
	// example:
	//
	// t-apsara-lindorm-2.1.20-20200518175539.alios7.x86_64
	Versions *string `json:"Versions,omitempty" xml:"Versions,omitempty"`
}

func (s UpgradeMultiZoneClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeMultiZoneClusterRequest) GetComponents() *string {
	return s.Components
}

func (s *UpgradeMultiZoneClusterRequest) GetRestartComponents() *string {
	return s.RestartComponents
}

func (s *UpgradeMultiZoneClusterRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *UpgradeMultiZoneClusterRequest) GetUpgradeInsName() *string {
	return s.UpgradeInsName
}

func (s *UpgradeMultiZoneClusterRequest) GetVersions() *string {
	return s.Versions
}

func (s *UpgradeMultiZoneClusterRequest) SetClusterId(v string) *UpgradeMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetComponents(v string) *UpgradeMultiZoneClusterRequest {
	s.Components = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetRestartComponents(v string) *UpgradeMultiZoneClusterRequest {
	s.RestartComponents = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetRunMode(v string) *UpgradeMultiZoneClusterRequest {
	s.RunMode = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetUpgradeInsName(v string) *UpgradeMultiZoneClusterRequest {
	s.UpgradeInsName = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetVersions(v string) *UpgradeMultiZoneClusterRequest {
	s.Versions = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) Validate() error {
	return dara.Validate(s)
}
