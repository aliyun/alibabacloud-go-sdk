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
	// Configuration for the DRM provider. To disable DRM, leave all fields in this object empty.
	DrmConfig *LivePackagingConfigDrmConfig `json:"DrmConfig,omitempty" xml:"DrmConfig,omitempty" type:"Struct"`
	// Live stream manifest configuration. Only one configuration is supported.
	LiveManifestConfigs []*LiveManifestConfig `json:"LiveManifestConfigs,omitempty" xml:"LiveManifestConfigs,omitempty" type:"Repeated"`
	// The duration of each output segment, in seconds. If not set, this defaults to the channel\\"s configured segment duration. The final segment duration is a multiple of the source segment duration that is closest to and not less than this value. Valid values: 1 to 30.
	//
	// example:
	//
	// 6
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// Specifies whether to create separate audio rendition groups for TS segments.
	//
	// example:
	//
	// true
	UseAudioRenditionGroups *bool `json:"UseAudioRenditionGroups,omitempty" xml:"UseAudioRenditionGroups,omitempty"`
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
	// The content ID in the DRM system. The maximum length is 256 characters. Letters, digits, underscores (_), and hyphens (-) are supported. You must ensure this ID is unique to prevent playback failures.
	//
	// example:
	//
	// live-axb1-9dd2fa123
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// The encryption method. Valid value:
	//
	// 	- SAMPLE_AES
	//
	// If not specified, encryption is disabled.
	//
	// example:
	//
	// SAMPLE_AES
	EncryptionMethod *string `json:"EncryptionMethod,omitempty" xml:"EncryptionMethod,omitempty"`
	// A 128-bit, 16-byte hex value represented by a 32-character string that is used with the key for encrypting data blocks. If you leave this parameter empty, MediaPackage creates a constant initialization vector (IV). If it is specified, the value is passed to the DRM service.
	//
	// example:
	//
	// 00000000000000000000000000000000
	IV *string `json:"IV,omitempty" xml:"IV,omitempty"`
	// The key rotation interval for DRM, in seconds. The default value of 0 disables key rotation.
	//
	// example:
	//
	// 0
	RotatePeriod *int32 `json:"RotatePeriod,omitempty" xml:"RotatePeriod,omitempty"`
	// The ID of the DRM system. The supported systems depend on the protocol.
	//
	// 	- DASH: Supports Google Widevine and Microsoft PlayReady.
	//
	// 	- HLS: DRM is not supported.
	//
	// 	- HLS-CMAF: Supports Apple FairPlay, Google Widevine, and Microsoft PlayReady.
	//
	// The corresponding System IDs are:
	//
	// 	- Apple FairPlay: 94ce86fb-07ff-4f43-adb8-93d2fa968ca2
	//
	// 	- Google Widevine: edef8ba9-79d6-4ace-a3c8-27dcd51d21ed
	//
	// 	- Microsoft PlayReady: 9a04f079-9840-4286-ab92-e65be0885f95
	SystemIds []*string `json:"SystemIds,omitempty" xml:"SystemIds,omitempty" type:"Repeated"`
	// The URL of the DRM key provider.
	//
	// example:
	//
	// https://exampledrm.com/path?arg1=xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
