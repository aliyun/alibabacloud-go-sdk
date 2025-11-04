// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLivePackagingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDrmConfig(v *LivePackagingConfigDrmConfig) *LivePackagingConfig
	GetDrmConfig() *LivePackagingConfigDrmConfig
	SetLiveManifestConfigs(v []*LiveManifestConfig) *LivePackagingConfig
	GetLiveManifestConfigs() []*LiveManifestConfig
	SetSegmentDuration(v int32) *LivePackagingConfig
	GetSegmentDuration() *int32
	SetUseAudioRenditionGroups(v bool) *LivePackagingConfig
	GetUseAudioRenditionGroups() *bool
}

type LivePackagingConfig struct {
	DrmConfig               *LivePackagingConfigDrmConfig `json:"DrmConfig,omitempty" xml:"DrmConfig,omitempty" type:"Struct"`
	LiveManifestConfigs     []*LiveManifestConfig         `json:"LiveManifestConfigs,omitempty" xml:"LiveManifestConfigs,omitempty" type:"Repeated"`
	SegmentDuration         *int32                        `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	UseAudioRenditionGroups *bool                         `json:"UseAudioRenditionGroups,omitempty" xml:"UseAudioRenditionGroups,omitempty"`
}

func (s LivePackagingConfig) String() string {
	return dara.Prettify(s)
}

func (s LivePackagingConfig) GoString() string {
	return s.String()
}

func (s *LivePackagingConfig) GetDrmConfig() *LivePackagingConfigDrmConfig {
	return s.DrmConfig
}

func (s *LivePackagingConfig) GetLiveManifestConfigs() []*LiveManifestConfig {
	return s.LiveManifestConfigs
}

func (s *LivePackagingConfig) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *LivePackagingConfig) GetUseAudioRenditionGroups() *bool {
	return s.UseAudioRenditionGroups
}

func (s *LivePackagingConfig) SetDrmConfig(v *LivePackagingConfigDrmConfig) *LivePackagingConfig {
	s.DrmConfig = v
	return s
}

func (s *LivePackagingConfig) SetLiveManifestConfigs(v []*LiveManifestConfig) *LivePackagingConfig {
	s.LiveManifestConfigs = v
	return s
}

func (s *LivePackagingConfig) SetSegmentDuration(v int32) *LivePackagingConfig {
	s.SegmentDuration = &v
	return s
}

func (s *LivePackagingConfig) SetUseAudioRenditionGroups(v bool) *LivePackagingConfig {
	s.UseAudioRenditionGroups = &v
	return s
}

func (s *LivePackagingConfig) Validate() error {
	if s.DrmConfig != nil {
		if err := s.DrmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LiveManifestConfigs != nil {
		for _, item := range s.LiveManifestConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LivePackagingConfigDrmConfig struct {
	ContentId        *string   `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	EncryptionMethod *string   `json:"EncryptionMethod,omitempty" xml:"EncryptionMethod,omitempty"`
	IV               *string   `json:"IV,omitempty" xml:"IV,omitempty"`
	RotatePeriod     *int32    `json:"RotatePeriod,omitempty" xml:"RotatePeriod,omitempty"`
	SystemIds        []*string `json:"SystemIds,omitempty" xml:"SystemIds,omitempty" type:"Repeated"`
	Url              *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s LivePackagingConfigDrmConfig) String() string {
	return dara.Prettify(s)
}

func (s LivePackagingConfigDrmConfig) GoString() string {
	return s.String()
}

func (s *LivePackagingConfigDrmConfig) GetContentId() *string {
	return s.ContentId
}

func (s *LivePackagingConfigDrmConfig) GetEncryptionMethod() *string {
	return s.EncryptionMethod
}

func (s *LivePackagingConfigDrmConfig) GetIV() *string {
	return s.IV
}

func (s *LivePackagingConfigDrmConfig) GetRotatePeriod() *int32 {
	return s.RotatePeriod
}

func (s *LivePackagingConfigDrmConfig) GetSystemIds() []*string {
	return s.SystemIds
}

func (s *LivePackagingConfigDrmConfig) GetUrl() *string {
	return s.Url
}

func (s *LivePackagingConfigDrmConfig) SetContentId(v string) *LivePackagingConfigDrmConfig {
	s.ContentId = &v
	return s
}

func (s *LivePackagingConfigDrmConfig) SetEncryptionMethod(v string) *LivePackagingConfigDrmConfig {
	s.EncryptionMethod = &v
	return s
}

func (s *LivePackagingConfigDrmConfig) SetIV(v string) *LivePackagingConfigDrmConfig {
	s.IV = &v
	return s
}

func (s *LivePackagingConfigDrmConfig) SetRotatePeriod(v int32) *LivePackagingConfigDrmConfig {
	s.RotatePeriod = &v
	return s
}

func (s *LivePackagingConfigDrmConfig) SetSystemIds(v []*string) *LivePackagingConfigDrmConfig {
	s.SystemIds = v
	return s
}

func (s *LivePackagingConfigDrmConfig) SetUrl(v string) *LivePackagingConfigDrmConfig {
	s.Url = &v
	return s
}

func (s *LivePackagingConfigDrmConfig) Validate() error {
	return dara.Validate(s)
}
