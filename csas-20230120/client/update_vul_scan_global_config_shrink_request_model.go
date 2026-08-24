// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanGlobalConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxDownloadSpeed(v int32) *UpdateVulScanGlobalConfigShrinkRequest
	GetMaxDownloadSpeed() *int32
	SetWuyingVulFixConfigShrink(v string) *UpdateVulScanGlobalConfigShrinkRequest
	GetWuyingVulFixConfigShrink() *string
}

type UpdateVulScanGlobalConfigShrinkRequest struct {
	// The maximum download rate for vulnerability patches on a single user terminal device. Unit: Byte/s. A value of 0 indicates no speed limit.
	//
	// example:
	//
	// 1048576
	MaxDownloadSpeed *int32 `json:"MaxDownloadSpeed,omitempty" xml:"MaxDownloadSpeed,omitempty"`
	// Deprecated
	//
	// The vulnerability fix configuration for WUYING Workspace. This configuration applies only to user terminal devices of the Cloud Desktop type.
	WuyingVulFixConfigShrink *string `json:"WuyingVulFixConfig,omitempty" xml:"WuyingVulFixConfig,omitempty"`
}

func (s UpdateVulScanGlobalConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanGlobalConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateVulScanGlobalConfigShrinkRequest) GetMaxDownloadSpeed() *int32 {
	return s.MaxDownloadSpeed
}

func (s *UpdateVulScanGlobalConfigShrinkRequest) GetWuyingVulFixConfigShrink() *string {
	return s.WuyingVulFixConfigShrink
}

func (s *UpdateVulScanGlobalConfigShrinkRequest) SetMaxDownloadSpeed(v int32) *UpdateVulScanGlobalConfigShrinkRequest {
	s.MaxDownloadSpeed = &v
	return s
}

func (s *UpdateVulScanGlobalConfigShrinkRequest) SetWuyingVulFixConfigShrink(v string) *UpdateVulScanGlobalConfigShrinkRequest {
	s.WuyingVulFixConfigShrink = &v
	return s
}

func (s *UpdateVulScanGlobalConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
