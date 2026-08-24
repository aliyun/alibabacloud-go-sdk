// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxDownloadSpeed(v int32) *UpdateVulScanGlobalConfigResponseBody
	GetMaxDownloadSpeed() *int32
	SetRequestId(v string) *UpdateVulScanGlobalConfigResponseBody
	GetRequestId() *string
	SetWuyingVulFixConfig(v *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) *UpdateVulScanGlobalConfigResponseBody
	GetWuyingVulFixConfig() *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig
}

type UpdateVulScanGlobalConfigResponseBody struct {
	// The maximum download rate for vulnerability patches on a single user terminal device. Unit: Byte/s. A value of 0 indicates no speed limit.
	//
	// example:
	//
	// 1048576
	MaxDownloadSpeed *int32 `json:"MaxDownloadSpeed,omitempty" xml:"MaxDownloadSpeed,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated
	//
	// The vulnerability fix configuration for WUYING Workspace. This configuration applies only to user terminal devices of the Cloud Desktop type.
	WuyingVulFixConfig *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig `json:"WuyingVulFixConfig,omitempty" xml:"WuyingVulFixConfig,omitempty" type:"Struct"`
}

func (s UpdateVulScanGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVulScanGlobalConfigResponseBody) GetMaxDownloadSpeed() *int32 {
	return s.MaxDownloadSpeed
}

func (s *UpdateVulScanGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVulScanGlobalConfigResponseBody) GetWuyingVulFixConfig() *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig {
	return s.WuyingVulFixConfig
}

func (s *UpdateVulScanGlobalConfigResponseBody) SetMaxDownloadSpeed(v int32) *UpdateVulScanGlobalConfigResponseBody {
	s.MaxDownloadSpeed = &v
	return s
}

func (s *UpdateVulScanGlobalConfigResponseBody) SetRequestId(v string) *UpdateVulScanGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVulScanGlobalConfigResponseBody) SetWuyingVulFixConfig(v *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) *UpdateVulScanGlobalConfigResponseBody {
	s.WuyingVulFixConfig = v
	return s
}

func (s *UpdateVulScanGlobalConfigResponseBody) Validate() error {
	if s.WuyingVulFixConfig != nil {
		if err := s.WuyingVulFixConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig struct {
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

func (s UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) GoString() string {
	return s.String()
}

func (s *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) GetAntiShutdownSwitch() *bool {
	return s.AntiShutdownSwitch
}

func (s *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) GetSnapshotSwitch() *bool {
	return s.SnapshotSwitch
}

func (s *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) SetAntiShutdownSwitch(v bool) *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig {
	s.AntiShutdownSwitch = &v
	return s
}

func (s *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) SetSnapshotSwitch(v bool) *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig {
	s.SnapshotSwitch = &v
	return s
}

func (s *UpdateVulScanGlobalConfigResponseBodyWuyingVulFixConfig) Validate() error {
	return dara.Validate(s)
}
