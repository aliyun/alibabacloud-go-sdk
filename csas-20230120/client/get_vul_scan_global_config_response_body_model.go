// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulScanGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxDownloadSpeed(v int32) *GetVulScanGlobalConfigResponseBody
	GetMaxDownloadSpeed() *int32
	SetRequestId(v string) *GetVulScanGlobalConfigResponseBody
	GetRequestId() *string
	SetWuyingVulFixConfig(v *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) *GetVulScanGlobalConfigResponseBody
	GetWuyingVulFixConfig() *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig
}

type GetVulScanGlobalConfigResponseBody struct {
	// The maximum download speed for vulnerability patches on a single user terminal device. Unit: bytes per second. A value of 0 indicates no speed limit.
	//
	// example:
	//
	// 1048576
	MaxDownloadSpeed *int32 `json:"MaxDownloadSpeed,omitempty" xml:"MaxDownloadSpeed,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The vulnerability fix configuration for WUYING Workspace. This configuration takes effect only on user terminal devices of the Cloud Desktop type.
	WuyingVulFixConfig *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig `json:"WuyingVulFixConfig,omitempty" xml:"WuyingVulFixConfig,omitempty" type:"Struct"`
}

func (s GetVulScanGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulScanGlobalConfigResponseBody) GetMaxDownloadSpeed() *int32 {
	return s.MaxDownloadSpeed
}

func (s *GetVulScanGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulScanGlobalConfigResponseBody) GetWuyingVulFixConfig() *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig {
	return s.WuyingVulFixConfig
}

func (s *GetVulScanGlobalConfigResponseBody) SetMaxDownloadSpeed(v int32) *GetVulScanGlobalConfigResponseBody {
	s.MaxDownloadSpeed = &v
	return s
}

func (s *GetVulScanGlobalConfigResponseBody) SetRequestId(v string) *GetVulScanGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulScanGlobalConfigResponseBody) SetWuyingVulFixConfig(v *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) *GetVulScanGlobalConfigResponseBody {
	s.WuyingVulFixConfig = v
	return s
}

func (s *GetVulScanGlobalConfigResponseBody) Validate() error {
	if s.WuyingVulFixConfig != nil {
		if err := s.WuyingVulFixConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig struct {
	// Specifies whether to prohibit shutdown during the fix process to prevent system exceptions caused by shutting down during patch installation. Valid values:
	//
	// - **true**: Shutdown is prohibited.
	//
	// - **false**: Shutdown is not prohibited.
	//
	// example:
	//
	// true
	AntiShutdownSwitch *bool `json:"AntiShutdownSwitch,omitempty" xml:"AntiShutdownSwitch,omitempty"`
	// Specifies whether to create a snapshot for the cloud desktop before the fix, which can be used for rollback if the fix fails. Valid values:
	//
	// - **true**: A snapshot is created.
	//
	// - **false**: No snapshot is created.
	//
	// example:
	//
	// true
	SnapshotSwitch *bool `json:"SnapshotSwitch,omitempty" xml:"SnapshotSwitch,omitempty"`
}

func (s GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) GoString() string {
	return s.String()
}

func (s *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) GetAntiShutdownSwitch() *bool {
	return s.AntiShutdownSwitch
}

func (s *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) GetSnapshotSwitch() *bool {
	return s.SnapshotSwitch
}

func (s *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) SetAntiShutdownSwitch(v bool) *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig {
	s.AntiShutdownSwitch = &v
	return s
}

func (s *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) SetSnapshotSwitch(v bool) *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig {
	s.SnapshotSwitch = &v
	return s
}

func (s *GetVulScanGlobalConfigResponseBodyWuyingVulFixConfig) Validate() error {
	return dara.Validate(s)
}
