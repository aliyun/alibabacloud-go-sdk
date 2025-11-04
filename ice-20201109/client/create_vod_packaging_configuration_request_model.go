// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationName(v string) *CreateVodPackagingConfigurationRequest
	GetConfigurationName() *string
	SetDescription(v string) *CreateVodPackagingConfigurationRequest
	GetDescription() *string
	SetGroupName(v string) *CreateVodPackagingConfigurationRequest
	GetGroupName() *string
	SetPackageConfig(v *CreateVodPackagingConfigurationRequestPackageConfig) *CreateVodPackagingConfigurationRequest
	GetPackageConfig() *CreateVodPackagingConfigurationRequestPackageConfig
	SetProtocol(v string) *CreateVodPackagingConfigurationRequest
	GetProtocol() *string
}

type CreateVodPackagingConfigurationRequest struct {
	// The name of the packaging configuration. The name must be unique in an account and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// hls_3s
	ConfigurationName *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty"`
	// The description of the packaging configuration.
	//
	// example:
	//
	// HLS 3s vod packaging
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the packaging group. The name can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The packaging configuration.
	PackageConfig *CreateVodPackagingConfigurationRequestPackageConfig `json:"PackageConfig,omitempty" xml:"PackageConfig,omitempty" type:"Struct"`
	// The package type.
	//
	// 	- HLS: packages content into TS segments for delivery over the HLS protocol.
	//
	// 	- HLS_CMAF: packages content into CMAF segments for delivery over the HLS protocol.
	//
	// 	- DASH: packages content for delivery over the DASH protocol.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateVodPackagingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationRequest) GetConfigurationName() *string {
	return s.ConfigurationName
}

func (s *CreateVodPackagingConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVodPackagingConfigurationRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateVodPackagingConfigurationRequest) GetPackageConfig() *CreateVodPackagingConfigurationRequestPackageConfig {
	return s.PackageConfig
}

func (s *CreateVodPackagingConfigurationRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateVodPackagingConfigurationRequest) SetConfigurationName(v string) *CreateVodPackagingConfigurationRequest {
	s.ConfigurationName = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequest) SetDescription(v string) *CreateVodPackagingConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequest) SetGroupName(v string) *CreateVodPackagingConfigurationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequest) SetPackageConfig(v *CreateVodPackagingConfigurationRequestPackageConfig) *CreateVodPackagingConfigurationRequest {
	s.PackageConfig = v
	return s
}

func (s *CreateVodPackagingConfigurationRequest) SetProtocol(v string) *CreateVodPackagingConfigurationRequest {
	s.Protocol = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequest) Validate() error {
	if s.PackageConfig != nil {
		if err := s.PackageConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVodPackagingConfigurationRequestPackageConfig struct {
	// The settings of digital rights management (DRM) encryption.
	//
	// if can be null:
	// true
	DrmProvider *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider `json:"DrmProvider,omitempty" xml:"DrmProvider,omitempty" type:"Struct"`
	// The manifest name. The name can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// index
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The duration of each segment in a packaged stream. Unit: seconds. MediaPackage rounds segments to the nearest multiple of the input segment duration. Valid values: 1 to 30.
	//
	// example:
	//
	// 6
	SegmentDuration *int64 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// The settings of stream selection.
	StreamSelection *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection `json:"StreamSelection,omitempty" xml:"StreamSelection,omitempty" type:"Struct"`
}

func (s CreateVodPackagingConfigurationRequestPackageConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationRequestPackageConfig) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) GetDrmProvider() *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider {
	return s.DrmProvider
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) GetManifestName() *string {
	return s.ManifestName
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) GetSegmentDuration() *int64 {
	return s.SegmentDuration
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) GetStreamSelection() *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection {
	return s.StreamSelection
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) SetDrmProvider(v *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) *CreateVodPackagingConfigurationRequestPackageConfig {
	s.DrmProvider = v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) SetManifestName(v string) *CreateVodPackagingConfigurationRequestPackageConfig {
	s.ManifestName = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) SetSegmentDuration(v int64) *CreateVodPackagingConfigurationRequestPackageConfig {
	s.SegmentDuration = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) SetStreamSelection(v *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) *CreateVodPackagingConfigurationRequestPackageConfig {
	s.StreamSelection = v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfig) Validate() error {
	if s.DrmProvider != nil {
		if err := s.DrmProvider.Validate(); err != nil {
			return err
		}
	}
	if s.StreamSelection != nil {
		if err := s.StreamSelection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVodPackagingConfigurationRequestPackageConfigDrmProvider struct {
	// The encryption method. Valid values:
	//
	// 	- AES_128: Advanced Encryption Standard (AES) with 128-bit key length.
	//
	// 	- SAMPLE_AES: an encryption method that encrypts individual media samples.
	//
	// example:
	//
	// AES_128
	EncryptionMethod *string `json:"EncryptionMethod,omitempty" xml:"EncryptionMethod,omitempty"`
	// A 128-bit, 16-byte hex value represented by a 32-character string that is used with the key for encrypting data blocks. If you leave this parameter empty, MediaPackage creates a constant initialization vector (IV). If it is specified, the value is passed to the DRM service.
	//
	// example:
	//
	// 00001111222233334444555566667777
	IV *string `json:"IV,omitempty" xml:"IV,omitempty"`
	// The ID of the DRM system. The maximum number of system IDs allowed is determined by the protocol type. Limits:
	//
	// 	- DASH: 2
	//
	// 	- HLS: 1
	//
	// 	- HLS_CMAF: 2
	//
	// Apple FairPlay, Google Widevine, and Microsoft PlayReady are supported. Their system IDs are as follows:
	//
	// 	- Apple FairPlay: 94ce86fb-07ff-4f43-adb8-93d2fa968ca2
	//
	// 	- Google Widevine: edef8ba9-79d6-4ace-a3c8-27dcd51d21e
	//
	// 	- Microsoft PlayReady: 9a04f079-9840-4286-ab92-e65be0885f95
	SystemIds []*string `json:"SystemIds,omitempty" xml:"SystemIds,omitempty" type:"Repeated"`
	// The URL of the DRM key provider.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) GetEncryptionMethod() *string {
	return s.EncryptionMethod
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) GetIV() *string {
	return s.IV
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) GetSystemIds() []*string {
	return s.SystemIds
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) GetUrl() *string {
	return s.Url
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) SetEncryptionMethod(v string) *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider {
	s.EncryptionMethod = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) SetIV(v string) *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider {
	s.IV = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) SetSystemIds(v []*string) *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider {
	s.SystemIds = v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) SetUrl(v string) *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider {
	s.Url = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigDrmProvider) Validate() error {
	return dara.Validate(s)
}

type CreateVodPackagingConfigurationRequestPackageConfigStreamSelection struct {
	// The maximum bitrate of the video stream. Unit: bit/s.
	//
	// example:
	//
	// 1000000000
	MaxVideoBitsPerSecond *int64 `json:"MaxVideoBitsPerSecond,omitempty" xml:"MaxVideoBitsPerSecond,omitempty"`
	// The minimum bitrate of the video stream. Unit: bit/s.
	//
	// example:
	//
	// 100000
	MinVideoBitsPerSecond *int64 `json:"MinVideoBitsPerSecond,omitempty" xml:"MinVideoBitsPerSecond,omitempty"`
	// The order of manifest files in the master playlist. Valid values:
	//
	// 	- ORIGINAL: sorts the manifest files in the same order as the source.
	//
	// 	- VIDEO_BITRATE_ASCENDING: sorts the manifest files in ascending order of bitrates, from lowest to highest.
	//
	// 	- VIDEO_BITRATE_DESCENDING: sorts the manifest files in descending order of bitrates, from highest to lowest.
	//
	// example:
	//
	// ORIGINAL
	StreamOrder *string `json:"StreamOrder,omitempty" xml:"StreamOrder,omitempty"`
}

func (s CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) GetMaxVideoBitsPerSecond() *int64 {
	return s.MaxVideoBitsPerSecond
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) GetMinVideoBitsPerSecond() *int64 {
	return s.MinVideoBitsPerSecond
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) GetStreamOrder() *string {
	return s.StreamOrder
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) SetMaxVideoBitsPerSecond(v int64) *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection {
	s.MaxVideoBitsPerSecond = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) SetMinVideoBitsPerSecond(v int64) *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection {
	s.MinVideoBitsPerSecond = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) SetStreamOrder(v string) *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection {
	s.StreamOrder = &v
	return s
}

func (s *CreateVodPackagingConfigurationRequestPackageConfigStreamSelection) Validate() error {
	return dara.Validate(s)
}
