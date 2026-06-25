// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDisplayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *ModifyDisplayConfigRequest
	GetAndroidInstanceIds() []*string
	SetDisplayConfig(v *ModifyDisplayConfigRequestDisplayConfig) *ModifyDisplayConfigRequest
	GetDisplayConfig() *ModifyDisplayConfigRequestDisplayConfig
}

type ModifyDisplayConfigRequest struct {
	// A list of instance IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// The display settings.
	DisplayConfig *ModifyDisplayConfigRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
}

func (s ModifyDisplayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDisplayConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *ModifyDisplayConfigRequest) GetDisplayConfig() *ModifyDisplayConfigRequestDisplayConfig {
	return s.DisplayConfig
}

func (s *ModifyDisplayConfigRequest) SetAndroidInstanceIds(v []*string) *ModifyDisplayConfigRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ModifyDisplayConfigRequest) SetDisplayConfig(v *ModifyDisplayConfigRequestDisplayConfig) *ModifyDisplayConfigRequest {
	s.DisplayConfig = v
	return s
}

func (s *ModifyDisplayConfigRequest) Validate() error {
	if s.DisplayConfig != nil {
		if err := s.DisplayConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDisplayConfigRequestDisplayConfig struct {
	// The dots per inch (DPI). Valid values: 72 to 600.
	//
	// example:
	//
	// 240
	Dpi *int32 `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The resolution height, in pixels.
	//
	// example:
	//
	// 1920
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution width, in pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s ModifyDisplayConfigRequestDisplayConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDisplayConfigRequestDisplayConfig) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigRequestDisplayConfig) GetDpi() *int32 {
	return s.Dpi
}

func (s *ModifyDisplayConfigRequestDisplayConfig) GetFps() *int32 {
	return s.Fps
}

func (s *ModifyDisplayConfigRequestDisplayConfig) GetLockResolution() *string {
	return s.LockResolution
}

func (s *ModifyDisplayConfigRequestDisplayConfig) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *ModifyDisplayConfigRequestDisplayConfig) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetDpi(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.Dpi = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetFps(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.Fps = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetLockResolution(v string) *ModifyDisplayConfigRequestDisplayConfig {
	s.LockResolution = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetResolutionHeight(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetResolutionWidth(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.ResolutionWidth = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) Validate() error {
	return dara.Validate(s)
}
