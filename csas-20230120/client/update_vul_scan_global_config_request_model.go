// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxDownloadSpeed(v int32) *UpdateVulScanGlobalConfigRequest
	GetMaxDownloadSpeed() *int32
	SetWuyingVulFixConfig(v *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) *UpdateVulScanGlobalConfigRequest
	GetWuyingVulFixConfig() *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig
}

type UpdateVulScanGlobalConfigRequest struct {
	// The maximum download rate for vulnerability patches on a single user terminal device. Unit: Byte/s. A value of 0 indicates no speed limit.
	//
	// example:
	//
	// 1048576
	MaxDownloadSpeed *int32 `json:"MaxDownloadSpeed,omitempty" xml:"MaxDownloadSpeed,omitempty"`
	// Deprecated
	//
	// The vulnerability fix configuration for WUYING Workspace. This configuration applies only to user terminal devices of the Cloud Desktop type.
	WuyingVulFixConfig *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig `json:"WuyingVulFixConfig,omitempty" xml:"WuyingVulFixConfig,omitempty" type:"Struct"`
}

func (s UpdateVulScanGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVulScanGlobalConfigRequest) GetMaxDownloadSpeed() *int32 {
	return s.MaxDownloadSpeed
}

func (s *UpdateVulScanGlobalConfigRequest) GetWuyingVulFixConfig() *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig {
	return s.WuyingVulFixConfig
}

func (s *UpdateVulScanGlobalConfigRequest) SetMaxDownloadSpeed(v int32) *UpdateVulScanGlobalConfigRequest {
	s.MaxDownloadSpeed = &v
	return s
}

func (s *UpdateVulScanGlobalConfigRequest) SetWuyingVulFixConfig(v *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) *UpdateVulScanGlobalConfigRequest {
	s.WuyingVulFixConfig = v
	return s
}

func (s *UpdateVulScanGlobalConfigRequest) Validate() error {
	if s.WuyingVulFixConfig != nil {
		if err := s.WuyingVulFixConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVulScanGlobalConfigRequestWuyingVulFixConfig struct {
	// Deprecated
	//
	// Specifies whether to prohibit shutdown during the fix process to prevent system exceptions caused by shutting down during patch installation. Valid values:
	//
	// - **true**: Prohibit shutdown.
	//
	// - **false**: Do not prohibit shutdown.
	//
	// example:
	//
	// true
	AntiShutdownSwitch *bool `json:"AntiShutdownSwitch,omitempty" xml:"AntiShutdownSwitch,omitempty"`
	// Deprecated
	//
	// Specifies whether to create a snapshot for the cloud desktop before the fix for rollback in case of fix failure. Valid values:
	//
	// - **true**: Create a snapshot.
	//
	// - **false**: Do not create a snapshot.
	//
	// example:
	//
	// true
	SnapshotSwitch *bool `json:"SnapshotSwitch,omitempty" xml:"SnapshotSwitch,omitempty"`
}

func (s UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) GoString() string {
	return s.String()
}

func (s *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) GetAntiShutdownSwitch() *bool {
	return s.AntiShutdownSwitch
}

func (s *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) GetSnapshotSwitch() *bool {
	return s.SnapshotSwitch
}

func (s *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) SetAntiShutdownSwitch(v bool) *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig {
	s.AntiShutdownSwitch = &v
	return s
}

func (s *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) SetSnapshotSwitch(v bool) *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig {
	s.SnapshotSwitch = &v
	return s
}

func (s *UpdateVulScanGlobalConfigRequestWuyingVulFixConfig) Validate() error {
	return dara.Validate(s)
}
