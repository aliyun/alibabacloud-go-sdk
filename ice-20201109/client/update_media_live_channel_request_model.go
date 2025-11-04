// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioSettings(v []*UpdateMediaLiveChannelRequestAudioSettings) *UpdateMediaLiveChannelRequest
	GetAudioSettings() []*UpdateMediaLiveChannelRequestAudioSettings
	SetChannelId(v string) *UpdateMediaLiveChannelRequest
	GetChannelId() *string
	SetInputAttachments(v []*UpdateMediaLiveChannelRequestInputAttachments) *UpdateMediaLiveChannelRequest
	GetInputAttachments() []*UpdateMediaLiveChannelRequestInputAttachments
	SetName(v string) *UpdateMediaLiveChannelRequest
	GetName() *string
	SetOutputGroups(v []*UpdateMediaLiveChannelRequestOutputGroups) *UpdateMediaLiveChannelRequest
	GetOutputGroups() []*UpdateMediaLiveChannelRequestOutputGroups
	SetVideoSettings(v []*UpdateMediaLiveChannelRequestVideoSettings) *UpdateMediaLiveChannelRequest
	GetVideoSettings() []*UpdateMediaLiveChannelRequestVideoSettings
}

type UpdateMediaLiveChannelRequest struct {
	// The audio settings.
	AudioSettings []*UpdateMediaLiveChannelRequestAudioSettings `json:"AudioSettings,omitempty" xml:"AudioSettings,omitempty" type:"Repeated"`
	// The ID of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The inputs associated with the channel.
	//
	// This parameter is required.
	InputAttachments []*UpdateMediaLiveChannelRequestInputAttachments `json:"InputAttachments,omitempty" xml:"InputAttachments,omitempty" type:"Repeated"`
	// The name of the channel. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// mych
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output groups.
	//
	// This parameter is required.
	OutputGroups []*UpdateMediaLiveChannelRequestOutputGroups `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	// The video settings.
	VideoSettings []*UpdateMediaLiveChannelRequestVideoSettings `json:"VideoSettings,omitempty" xml:"VideoSettings,omitempty" type:"Repeated"`
}

func (s UpdateMediaLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequest) GetAudioSettings() []*UpdateMediaLiveChannelRequestAudioSettings {
	return s.AudioSettings
}

func (s *UpdateMediaLiveChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateMediaLiveChannelRequest) GetInputAttachments() []*UpdateMediaLiveChannelRequestInputAttachments {
	return s.InputAttachments
}

func (s *UpdateMediaLiveChannelRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelRequest) GetOutputGroups() []*UpdateMediaLiveChannelRequestOutputGroups {
	return s.OutputGroups
}

func (s *UpdateMediaLiveChannelRequest) GetVideoSettings() []*UpdateMediaLiveChannelRequestVideoSettings {
	return s.VideoSettings
}

func (s *UpdateMediaLiveChannelRequest) SetAudioSettings(v []*UpdateMediaLiveChannelRequestAudioSettings) *UpdateMediaLiveChannelRequest {
	s.AudioSettings = v
	return s
}

func (s *UpdateMediaLiveChannelRequest) SetChannelId(v string) *UpdateMediaLiveChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateMediaLiveChannelRequest) SetInputAttachments(v []*UpdateMediaLiveChannelRequestInputAttachments) *UpdateMediaLiveChannelRequest {
	s.InputAttachments = v
	return s
}

func (s *UpdateMediaLiveChannelRequest) SetName(v string) *UpdateMediaLiveChannelRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelRequest) SetOutputGroups(v []*UpdateMediaLiveChannelRequestOutputGroups) *UpdateMediaLiveChannelRequest {
	s.OutputGroups = v
	return s
}

func (s *UpdateMediaLiveChannelRequest) SetVideoSettings(v []*UpdateMediaLiveChannelRequestVideoSettings) *UpdateMediaLiveChannelRequest {
	s.VideoSettings = v
	return s
}

func (s *UpdateMediaLiveChannelRequest) Validate() error {
	if s.AudioSettings != nil {
		for _, item := range s.AudioSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InputAttachments != nil {
		for _, item := range s.InputAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputGroups != nil {
		for _, item := range s.OutputGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoSettings != nil {
		for _, item := range s.VideoSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestAudioSettings struct {
	// The audio codec. If it is not specified, the source specification is used. Valid values: aac and libfdk_aac.
	//
	// example:
	//
	// libfdk_aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio encoding settings.
	AudioCodecSetting *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting `json:"AudioCodecSetting,omitempty" xml:"AudioCodecSetting,omitempty" type:"Struct"`
	// The name of the audio selector.
	//
	// example:
	//
	// a1
	AudioSelectorName *string `json:"AudioSelectorName,omitempty" xml:"AudioSelectorName,omitempty"`
	// Enter a three-letter ISO 639-2 language code. If the audio track selected by the audio selector has a language code, the language code specified in the audio selector is used. If the selected audio track does not have a language code, or if the audio selector cannot find a track that matches its criteria, this language code is used.
	//
	// example:
	//
	// eng
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// The tag that identifies the language of the RTMP input. It can be referenced by the output. The maximum length is 32 characters. Supported characters:
	//
	// 	- Unicode letters
	//
	// 	- Digits (0-9)
	//
	// 	- Underscore (_)
	//
	// 	- Hyphen (-)
	//
	// 	- Space (a space cannot be at the beginning or end)
	//
	// example:
	//
	// English
	LanguageName *string `json:"LanguageName,omitempty" xml:"LanguageName,omitempty"`
	// The name of the audio settings. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// audio1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMediaLiveChannelRequestAudioSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestAudioSettings) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) GetAudioCodecSetting() *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	return s.AudioCodecSetting
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) GetAudioSelectorName() *string {
	return s.AudioSelectorName
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) GetLanguageName() *string {
	return s.LanguageName
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) SetAudioCodec(v string) *UpdateMediaLiveChannelRequestAudioSettings {
	s.AudioCodec = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) SetAudioCodecSetting(v *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) *UpdateMediaLiveChannelRequestAudioSettings {
	s.AudioCodecSetting = v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) SetAudioSelectorName(v string) *UpdateMediaLiveChannelRequestAudioSettings {
	s.AudioSelectorName = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) SetLanguageCode(v string) *UpdateMediaLiveChannelRequestAudioSettings {
	s.LanguageCode = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) SetLanguageName(v string) *UpdateMediaLiveChannelRequestAudioSettings {
	s.LanguageName = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) SetName(v string) *UpdateMediaLiveChannelRequestAudioSettings {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettings) Validate() error {
	if s.AudioCodecSetting != nil {
		if err := s.AudioCodecSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting struct {
	// The audio bitrate. Unit: bit/s. Valid values: 8000 to 1000000. The value must be divisible by 1000.
	//
	// example:
	//
	// 200000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The audio codec profile. When AudioCodec is set to aac, AAC-LOW and AAC-MAIN are supported. When AudioCodec is set to libfdk_aac, AAC-LOW, AAC-HE, and AAC-HEV2 are supported.
	//
	// example:
	//
	// AAC-LOW
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sample rate. Unit: Hz. Valid values: 22050, 32000, 44100, 48000, and 96000.
	//
	// example:
	//
	// 44100
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GetProfile() *string {
	return s.Profile
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) SetBitrate(v int32) *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	s.Bitrate = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) SetProfile(v string) *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	s.Profile = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) SetSampleRate(v int32) *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	s.SampleRate = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestInputAttachments struct {
	// The audio selectors.
	AudioSelectors []*UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors `json:"AudioSelectors,omitempty" xml:"AudioSelectors,omitempty" type:"Repeated"`
	// The ID of the associated input.
	//
	// This parameter is required.
	//
	// example:
	//
	// myinput
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// The tag that identifies the language of the RTMP input. It can be referenced by the output. The maximum length is 32 characters. Supported characters:
	//
	// 	- Unicode letters
	//
	// 	- Digits (0-9)
	//
	// 	- Underscore (_)
	//
	// 	- Hyphen (-)
	//
	// 	- Space (a space cannot be at the beginning or end)
	//
	// example:
	//
	// English
	LanguageName *string `json:"LanguageName,omitempty" xml:"LanguageName,omitempty"`
}

func (s UpdateMediaLiveChannelRequestInputAttachments) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestInputAttachments) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) GetAudioSelectors() []*UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	return s.AudioSelectors
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) GetInputId() *string {
	return s.InputId
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) GetLanguageName() *string {
	return s.LanguageName
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) SetAudioSelectors(v []*UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) *UpdateMediaLiveChannelRequestInputAttachments {
	s.AudioSelectors = v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) SetInputId(v string) *UpdateMediaLiveChannelRequestInputAttachments {
	s.InputId = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) SetLanguageName(v string) *UpdateMediaLiveChannelRequestInputAttachments {
	s.LanguageName = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachments) Validate() error {
	if s.AudioSelectors != nil {
		for _, item := range s.AudioSelectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors struct {
	// The audio language selection.
	AudioLanguageSelection *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection `json:"AudioLanguageSelection,omitempty" xml:"AudioLanguageSelection,omitempty" type:"Struct"`
	// The audio PID selection.
	AudioPidSelection *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection `json:"AudioPidSelection,omitempty" xml:"AudioPidSelection,omitempty" type:"Struct"`
	// The audio track selection.
	AudioTrackSelection []*UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection `json:"AudioTrackSelection,omitempty" xml:"AudioTrackSelection,omitempty" type:"Repeated"`
	// The name of the audio selector. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myselector
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetAudioLanguageSelection() *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection {
	return s.AudioLanguageSelection
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetAudioPidSelection() *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection {
	return s.AudioPidSelection
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetAudioTrackSelection() []*UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection {
	return s.AudioTrackSelection
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetAudioLanguageSelection(v *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.AudioLanguageSelection = v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetAudioPidSelection(v *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.AudioPidSelection = v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetAudioTrackSelection(v []*UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.AudioTrackSelection = v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetName(v string) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectors) Validate() error {
	if s.AudioLanguageSelection != nil {
		if err := s.AudioLanguageSelection.Validate(); err != nil {
			return err
		}
	}
	if s.AudioPidSelection != nil {
		if err := s.AudioPidSelection.Validate(); err != nil {
			return err
		}
	}
	if s.AudioTrackSelection != nil {
		for _, item := range s.AudioTrackSelection {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection struct {
	// Enter a three-letter ISO 639-2 language code from within an audio source.
	//
	// This parameter is required.
	//
	// example:
	//
	// eng
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) SetLanguageCode(v string) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection {
	s.LanguageCode = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection struct {
	// Enter a specific PID from within a source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Pid *int64 `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) GetPid() *int64 {
	return s.Pid
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) SetPid(v int64) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection {
	s.Pid = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection struct {
	// Specify one or more audio tracks from within a source using Track ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TrackId *int64 `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) GetTrackId() *int64 {
	return s.TrackId
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) SetTrackId(v int64) *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection {
	s.TrackId = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestOutputGroups struct {
	// The MediaPackage destination.
	MediaPackageGroupSetting *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting `json:"MediaPackageGroupSetting,omitempty" xml:"MediaPackageGroupSetting,omitempty" type:"Struct"`
	// The name of the output group. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The outputs in the output group.
	//
	// This parameter is required.
	Outputs []*UpdateMediaLiveChannelRequestOutputGroupsOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The output group type. Only MediaPackage is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// MediaPackage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateMediaLiveChannelRequestOutputGroups) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestOutputGroups) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) GetMediaPackageGroupSetting() *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting {
	return s.MediaPackageGroupSetting
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) GetOutputs() []*UpdateMediaLiveChannelRequestOutputGroupsOutputs {
	return s.Outputs
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) GetType() *string {
	return s.Type
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) SetMediaPackageGroupSetting(v *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) *UpdateMediaLiveChannelRequestOutputGroups {
	s.MediaPackageGroupSetting = v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) SetName(v string) *UpdateMediaLiveChannelRequestOutputGroups {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) SetOutputs(v []*UpdateMediaLiveChannelRequestOutputGroupsOutputs) *UpdateMediaLiveChannelRequestOutputGroups {
	s.Outputs = v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) SetType(v string) *UpdateMediaLiveChannelRequestOutputGroups {
	s.Type = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroups) Validate() error {
	if s.MediaPackageGroupSetting != nil {
		if err := s.MediaPackageGroupSetting.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting struct {
	// ChannelName in MediaPackage.
	//
	// This parameter is required.
	//
	// example:
	//
	// myPackageChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// GroupName in MediaPackage.
	//
	// This parameter is required.
	//
	// example:
	//
	// myPackageGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) SetChannelName(v string) *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting {
	s.ChannelName = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) SetGroupName(v string) *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting {
	s.GroupName = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestOutputGroupsOutputs struct {
	// The referenced AudioSettings.
	AudioSettingNames []*string `json:"AudioSettingNames,omitempty" xml:"AudioSettingNames,omitempty" type:"Repeated"`
	// The settings of the output delivered to MediaPackage.
	MediaPackageOutputSetting *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting `json:"MediaPackageOutputSetting,omitempty" xml:"MediaPackageOutputSetting,omitempty" type:"Struct"`
	// The media type of the output. Valid values:
	//
	// 	- 0: Audio and Video
	//
	// 	- 1: Audio If you set the value to 1, you cannot reference VideoSettings.
	//
	// 	- 2: Video. If you set the value to 2, you cannot reference AudioSettings.
	//
	// example:
	//
	// 0
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The name of the output. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// output1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the referenced VideoSettings.
	//
	// example:
	//
	// myVideo1
	VideoSettingName *string `json:"VideoSettingName,omitempty" xml:"VideoSettingName,omitempty"`
}

func (s UpdateMediaLiveChannelRequestOutputGroupsOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestOutputGroupsOutputs) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) GetAudioSettingNames() []*string {
	return s.AudioSettingNames
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) GetMediaPackageOutputSetting() *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting {
	return s.MediaPackageOutputSetting
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) GetMediaType() *int32 {
	return s.MediaType
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) GetVideoSettingName() *string {
	return s.VideoSettingName
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) SetAudioSettingNames(v []*string) *UpdateMediaLiveChannelRequestOutputGroupsOutputs {
	s.AudioSettingNames = v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) SetMediaPackageOutputSetting(v *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) *UpdateMediaLiveChannelRequestOutputGroupsOutputs {
	s.MediaPackageOutputSetting = v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) SetMediaType(v int32) *UpdateMediaLiveChannelRequestOutputGroupsOutputs {
	s.MediaType = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) SetName(v string) *UpdateMediaLiveChannelRequestOutputGroupsOutputs {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) SetVideoSettingName(v string) *UpdateMediaLiveChannelRequestOutputGroupsOutputs {
	s.VideoSettingName = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputs) Validate() error {
	if s.MediaPackageOutputSetting != nil {
		if err := s.MediaPackageOutputSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting struct {
	// The manifest audio group ID. To associate several audio tracks into one group, assign the same audio group ID. Viewers can select a track as needed. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 40 characters in length.
	//
	// example:
	//
	// audiogroup
	AudioGroupId *string `json:"AudioGroupId,omitempty" xml:"AudioGroupId,omitempty"`
	// The manifest name modifier. The child manifests include this modifier in their M3U8 file names. Letters, digits, hyphens (-), and underscores (_) are supported. The maximum length is 40 characters.
	//
	// example:
	//
	// 480p
	NameModifier *string `json:"NameModifier,omitempty" xml:"NameModifier,omitempty"`
}

func (s UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) GetAudioGroupId() *string {
	return s.AudioGroupId
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) GetNameModifier() *string {
	return s.NameModifier
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) SetAudioGroupId(v string) *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting {
	s.AudioGroupId = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) SetNameModifier(v string) *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting {
	s.NameModifier = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestVideoSettings struct {
	// The height of the output. If you set it to 0 or leave it empty, the height automatically adapts to the specified width to maintain the original aspect ratio.
	//
	// Valid values:
	//
	// 	- For regular transcoding, the larger dimension cannot exceed 3840 px, and the smaller one cannot exceed 2160 px.
	//
	// 	- For Narrowband HD™ transcoding, the larger dimension cannot exceed 1920 px, and the smaller one cannot exceed 1080 px.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the video settings. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// video1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The video codec. Valid values: H264 and H265.
	//
	// example:
	//
	// H264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// The video encoding settings.
	VideoCodecSetting *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting `json:"VideoCodecSetting,omitempty" xml:"VideoCodecSetting,omitempty" type:"Struct"`
	// The video transcoding method. Valid values:
	//
	// 	- NORMAL: regular transcoding
	//
	// 	- NBHD: Narrowband HD™ transcoding
	//
	// If not specified, regular transcoding is used by default.
	//
	// example:
	//
	// NORMAL
	VideoCodecType *string `json:"VideoCodecType,omitempty" xml:"VideoCodecType,omitempty"`
	// The width of the output. If you set it to 0 or leave it empty, the width automatically adapts to the specified height to maintain the original aspect ratio.
	//
	// Valid values:
	//
	// 	- For regular transcoding, the larger dimension cannot exceed 3840 px, and the smaller one cannot exceed 2160 px.
	//
	// 	- For Narrowband HD™ transcoding, the larger dimension cannot exceed 1920 px, and the smaller one cannot exceed 1080 px.
	//
	// example:
	//
	// 1280
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateMediaLiveChannelRequestVideoSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestVideoSettings) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) GetHeight() *int32 {
	return s.Height
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) GetVideoCodecSetting() *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	return s.VideoCodecSetting
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) GetVideoCodecType() *string {
	return s.VideoCodecType
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) GetWidth() *int32 {
	return s.Width
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) SetHeight(v int32) *UpdateMediaLiveChannelRequestVideoSettings {
	s.Height = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) SetName(v string) *UpdateMediaLiveChannelRequestVideoSettings {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) SetVideoCodec(v string) *UpdateMediaLiveChannelRequestVideoSettings {
	s.VideoCodec = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) SetVideoCodecSetting(v *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) *UpdateMediaLiveChannelRequestVideoSettings {
	s.VideoCodecSetting = v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) SetVideoCodecType(v string) *UpdateMediaLiveChannelRequestVideoSettings {
	s.VideoCodecType = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) SetWidth(v int32) *UpdateMediaLiveChannelRequestVideoSettings {
	s.Width = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettings) Validate() error {
	if s.VideoCodecSetting != nil {
		if err := s.VideoCodecSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting struct {
	// The video encoding settings.
	CodecDetail *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail `json:"CodecDetail,omitempty" xml:"CodecDetail,omitempty" type:"Struct"`
	// The frame rate. If it is not specified, the source specification is used.
	Framerate *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate `json:"Framerate,omitempty" xml:"Framerate,omitempty" type:"Struct"`
	// The GOP setting. If it is not specified, the source specification is used.
	Gop *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop `json:"Gop,omitempty" xml:"Gop,omitempty" type:"Struct"`
	// The video encoding rate. If it is not specified, the source specification is used.
	Rate *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate `json:"Rate,omitempty" xml:"Rate,omitempty" type:"Struct"`
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetCodecDetail() *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail {
	return s.CodecDetail
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetFramerate() *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	return s.Framerate
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetGop() *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	return s.Gop
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetRate() *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	return s.Rate
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetCodecDetail(v *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.CodecDetail = v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetFramerate(v *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.Framerate = v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetGop(v *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.Gop = v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetRate(v *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.Rate = v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) Validate() error {
	if s.CodecDetail != nil {
		if err := s.CodecDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Framerate != nil {
		if err := s.Framerate.Validate(); err != nil {
			return err
		}
	}
	if s.Gop != nil {
		if err := s.Gop.Validate(); err != nil {
			return err
		}
	}
	if s.Rate != nil {
		if err := s.Rate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail struct {
	// The video encoding level. It is not supported yet.
	//
	// example:
	//
	// H264_LEVEL_AUTO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The H.264 profile. Valid values: BASELINE, HIGH, and MAIN. Default value: MAIN. The parameter takes effect only when the codec is H.264.
	//
	// example:
	//
	// MAIN
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) GetLevel() *string {
	return s.Level
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) GetProfile() *string {
	return s.Profile
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) SetLevel(v string) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail {
	s.Level = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) SetProfile(v string) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail {
	s.Profile = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate struct {
	// The frame rate mode. Valid values: SPECIFIED (fixed frame rate) and FROM_SOURCE (use source specification).
	//
	// example:
	//
	// SPECIFIED
	FramerateControl *string `json:"FramerateControl,omitempty" xml:"FramerateControl,omitempty"`
	// The denominator of the fixed frame rate. The parameter is required when FramerateControl is set to SPECIFIED. Valid values: 1 to 60. The numerator must be divisible by the denominator.
	//
	// example:
	//
	// 1
	FramerateDenominator *int32 `json:"FramerateDenominator,omitempty" xml:"FramerateDenominator,omitempty"`
	// The numerator of the fixed frame rate. The parameter is required when FramerateControl is set to SPECIFIED. Valid values: 1 to 60. The numerator must be divisible by the denominator.
	//
	// example:
	//
	// 25
	FramerateNumerator *int32 `json:"FramerateNumerator,omitempty" xml:"FramerateNumerator,omitempty"`
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GetFramerateControl() *string {
	return s.FramerateControl
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GetFramerateDenominator() *int32 {
	return s.FramerateDenominator
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GetFramerateNumerator() *int32 {
	return s.FramerateNumerator
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) SetFramerateControl(v string) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	s.FramerateControl = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) SetFramerateDenominator(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	s.FramerateDenominator = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) SetFramerateNumerator(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	s.FramerateNumerator = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop struct {
	// The number of B frames. Valid values: 1 to 3.
	//
	// example:
	//
	// 3
	BframesNum *int32 `json:"BframesNum,omitempty" xml:"BframesNum,omitempty"`
	// The GOP size. When GopSizeUnits is set to SECONDS, the value range is from 1 to 20. When GopSizeUnits is set to FRAMES, the value range is from 1 to 3000.
	//
	// example:
	//
	// 90
	GopSize *int32 `json:"GopSize,omitempty" xml:"GopSize,omitempty"`
	// The GOP size unit. Valid values: FRAMES and SECONDS.
	//
	// example:
	//
	// FRAMES
	GopSizeUnits *string `json:"GopSizeUnits,omitempty" xml:"GopSizeUnits,omitempty"`
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GetBframesNum() *int32 {
	return s.BframesNum
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GetGopSize() *int32 {
	return s.GopSize
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GetGopSizeUnits() *string {
	return s.GopSizeUnits
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) SetBframesNum(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	s.BframesNum = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) SetGopSize(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	s.GopSize = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) SetGopSizeUnits(v string) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	s.GopSizeUnits = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate struct {
	// The video bitrate. Unit: bit/s. If you set it to 0 or leave it empty, the source specification is used. Valid values: 50000 to 6000000. The value must be divisible by 1000.
	//
	// example:
	//
	// 2500000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video buffer size. Unit: bit/s. Valid values: 100000 to 6000000. The value must be divisible by 1000.
	//
	// example:
	//
	// 6000000
	BufferSize *int32 `json:"BufferSize,omitempty" xml:"BufferSize,omitempty"`
	// The maximum bitrate. Unit: bit/s. Valid values: 100000 to 6000000. The value must be divisible by 1000.
	//
	// example:
	//
	// 6000000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The bitrate control mode. Valid values: CBR, ABR, and VBR.
	//
	// example:
	//
	// ABR
	RateControlMode *string `json:"RateControlMode,omitempty" xml:"RateControlMode,omitempty"`
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetBufferSize() *int32 {
	return s.BufferSize
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetRateControlMode() *string {
	return s.RateControlMode
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetBitrate(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.Bitrate = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetBufferSize(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.BufferSize = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetMaxBitrate(v int32) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.MaxBitrate = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetRateControlMode(v string) *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.RateControlMode = &v
	return s
}

func (s *UpdateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) Validate() error {
	return dara.Validate(s)
}
