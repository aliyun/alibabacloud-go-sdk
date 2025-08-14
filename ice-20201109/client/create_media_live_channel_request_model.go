// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioSettings(v []*CreateMediaLiveChannelRequestAudioSettings) *CreateMediaLiveChannelRequest
	GetAudioSettings() []*CreateMediaLiveChannelRequestAudioSettings
	SetInputAttachments(v []*CreateMediaLiveChannelRequestInputAttachments) *CreateMediaLiveChannelRequest
	GetInputAttachments() []*CreateMediaLiveChannelRequestInputAttachments
	SetName(v string) *CreateMediaLiveChannelRequest
	GetName() *string
	SetOutputGroups(v []*CreateMediaLiveChannelRequestOutputGroups) *CreateMediaLiveChannelRequest
	GetOutputGroups() []*CreateMediaLiveChannelRequestOutputGroups
	SetVideoSettings(v []*CreateMediaLiveChannelRequestVideoSettings) *CreateMediaLiveChannelRequest
	GetVideoSettings() []*CreateMediaLiveChannelRequestVideoSettings
}

type CreateMediaLiveChannelRequest struct {
	// The audio settings.
	AudioSettings []*CreateMediaLiveChannelRequestAudioSettings `json:"AudioSettings,omitempty" xml:"AudioSettings,omitempty" type:"Repeated"`
	// The associated inputs.
	//
	// This parameter is required.
	InputAttachments []*CreateMediaLiveChannelRequestInputAttachments `json:"InputAttachments,omitempty" xml:"InputAttachments,omitempty" type:"Repeated"`
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
	OutputGroups []*CreateMediaLiveChannelRequestOutputGroups `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	// The video settings.
	VideoSettings []*CreateMediaLiveChannelRequestVideoSettings `json:"VideoSettings,omitempty" xml:"VideoSettings,omitempty" type:"Repeated"`
}

func (s CreateMediaLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequest) GetAudioSettings() []*CreateMediaLiveChannelRequestAudioSettings {
	return s.AudioSettings
}

func (s *CreateMediaLiveChannelRequest) GetInputAttachments() []*CreateMediaLiveChannelRequestInputAttachments {
	return s.InputAttachments
}

func (s *CreateMediaLiveChannelRequest) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelRequest) GetOutputGroups() []*CreateMediaLiveChannelRequestOutputGroups {
	return s.OutputGroups
}

func (s *CreateMediaLiveChannelRequest) GetVideoSettings() []*CreateMediaLiveChannelRequestVideoSettings {
	return s.VideoSettings
}

func (s *CreateMediaLiveChannelRequest) SetAudioSettings(v []*CreateMediaLiveChannelRequestAudioSettings) *CreateMediaLiveChannelRequest {
	s.AudioSettings = v
	return s
}

func (s *CreateMediaLiveChannelRequest) SetInputAttachments(v []*CreateMediaLiveChannelRequestInputAttachments) *CreateMediaLiveChannelRequest {
	s.InputAttachments = v
	return s
}

func (s *CreateMediaLiveChannelRequest) SetName(v string) *CreateMediaLiveChannelRequest {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelRequest) SetOutputGroups(v []*CreateMediaLiveChannelRequestOutputGroups) *CreateMediaLiveChannelRequest {
	s.OutputGroups = v
	return s
}

func (s *CreateMediaLiveChannelRequest) SetVideoSettings(v []*CreateMediaLiveChannelRequestVideoSettings) *CreateMediaLiveChannelRequest {
	s.VideoSettings = v
	return s
}

func (s *CreateMediaLiveChannelRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestAudioSettings struct {
	// The audio codec. If it is not specified, the source specification is used. Valid values: aac and libfdk_aac.
	//
	// example:
	//
	// libfdk_aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio encoding settings.
	AudioCodecSetting *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting `json:"AudioCodecSetting,omitempty" xml:"AudioCodecSetting,omitempty" type:"Struct"`
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

func (s CreateMediaLiveChannelRequestAudioSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestAudioSettings) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestAudioSettings) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *CreateMediaLiveChannelRequestAudioSettings) GetAudioCodecSetting() *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	return s.AudioCodecSetting
}

func (s *CreateMediaLiveChannelRequestAudioSettings) GetAudioSelectorName() *string {
	return s.AudioSelectorName
}

func (s *CreateMediaLiveChannelRequestAudioSettings) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *CreateMediaLiveChannelRequestAudioSettings) GetLanguageName() *string {
	return s.LanguageName
}

func (s *CreateMediaLiveChannelRequestAudioSettings) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelRequestAudioSettings) SetAudioCodec(v string) *CreateMediaLiveChannelRequestAudioSettings {
	s.AudioCodec = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettings) SetAudioCodecSetting(v *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) *CreateMediaLiveChannelRequestAudioSettings {
	s.AudioCodecSetting = v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettings) SetAudioSelectorName(v string) *CreateMediaLiveChannelRequestAudioSettings {
	s.AudioSelectorName = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettings) SetLanguageCode(v string) *CreateMediaLiveChannelRequestAudioSettings {
	s.LanguageCode = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettings) SetLanguageName(v string) *CreateMediaLiveChannelRequestAudioSettings {
	s.LanguageName = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettings) SetName(v string) *CreateMediaLiveChannelRequestAudioSettings {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettings) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting struct {
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

func (s CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GetProfile() *string {
	return s.Profile
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) SetBitrate(v int32) *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	s.Bitrate = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) SetProfile(v string) *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	s.Profile = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) SetSampleRate(v int32) *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting {
	s.SampleRate = &v
	return s
}

func (s *CreateMediaLiveChannelRequestAudioSettingsAudioCodecSetting) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestInputAttachments struct {
	// The audio selectors.
	AudioSelectors []*CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors `json:"AudioSelectors,omitempty" xml:"AudioSelectors,omitempty" type:"Repeated"`
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

func (s CreateMediaLiveChannelRequestInputAttachments) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestInputAttachments) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestInputAttachments) GetAudioSelectors() []*CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	return s.AudioSelectors
}

func (s *CreateMediaLiveChannelRequestInputAttachments) GetInputId() *string {
	return s.InputId
}

func (s *CreateMediaLiveChannelRequestInputAttachments) GetLanguageName() *string {
	return s.LanguageName
}

func (s *CreateMediaLiveChannelRequestInputAttachments) SetAudioSelectors(v []*CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) *CreateMediaLiveChannelRequestInputAttachments {
	s.AudioSelectors = v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachments) SetInputId(v string) *CreateMediaLiveChannelRequestInputAttachments {
	s.InputId = &v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachments) SetLanguageName(v string) *CreateMediaLiveChannelRequestInputAttachments {
	s.LanguageName = &v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachments) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors struct {
	// The audio language selection.
	AudioLanguageSelection *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection `json:"AudioLanguageSelection,omitempty" xml:"AudioLanguageSelection,omitempty" type:"Struct"`
	// The audio PID selection.
	AudioPidSelection *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection `json:"AudioPidSelection,omitempty" xml:"AudioPidSelection,omitempty" type:"Struct"`
	// The audio track selection.
	AudioTrackSelection []*CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection `json:"AudioTrackSelection,omitempty" xml:"AudioTrackSelection,omitempty" type:"Repeated"`
	// The name of the audio selector. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myselector
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetAudioLanguageSelection() *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection {
	return s.AudioLanguageSelection
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetAudioPidSelection() *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection {
	return s.AudioPidSelection
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetAudioTrackSelection() []*CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection {
	return s.AudioTrackSelection
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetAudioLanguageSelection(v *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.AudioLanguageSelection = v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetAudioPidSelection(v *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.AudioPidSelection = v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetAudioTrackSelection(v []*CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.AudioTrackSelection = v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) SetName(v string) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectors) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection struct {
	// Enter a three-letter ISO 639-2 language code from within an audio source.
	//
	// This parameter is required.
	//
	// example:
	//
	// eng
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) SetLanguageCode(v string) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection {
	s.LanguageCode = &v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioLanguageSelection) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection struct {
	// Enter a specific PID from within a source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Pid *int64 `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) GetPid() *int64 {
	return s.Pid
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) SetPid(v int64) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection {
	s.Pid = &v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioPidSelection) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection struct {
	// Specify one or more audio tracks from within a source using Track ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TrackId *int64 `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) GetTrackId() *int64 {
	return s.TrackId
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) SetTrackId(v int64) *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection {
	s.TrackId = &v
	return s
}

func (s *CreateMediaLiveChannelRequestInputAttachmentsAudioSelectorsAudioTrackSelection) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestOutputGroups struct {
	// The MediaPackage destination.
	MediaPackageGroupSetting *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting `json:"MediaPackageGroupSetting,omitempty" xml:"MediaPackageGroupSetting,omitempty" type:"Struct"`
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
	Outputs []*CreateMediaLiveChannelRequestOutputGroupsOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The output group type. Only MediaPackage is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// MediaPackage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMediaLiveChannelRequestOutputGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestOutputGroups) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestOutputGroups) GetMediaPackageGroupSetting() *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting {
	return s.MediaPackageGroupSetting
}

func (s *CreateMediaLiveChannelRequestOutputGroups) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelRequestOutputGroups) GetOutputs() []*CreateMediaLiveChannelRequestOutputGroupsOutputs {
	return s.Outputs
}

func (s *CreateMediaLiveChannelRequestOutputGroups) GetType() *string {
	return s.Type
}

func (s *CreateMediaLiveChannelRequestOutputGroups) SetMediaPackageGroupSetting(v *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) *CreateMediaLiveChannelRequestOutputGroups {
	s.MediaPackageGroupSetting = v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroups) SetName(v string) *CreateMediaLiveChannelRequestOutputGroups {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroups) SetOutputs(v []*CreateMediaLiveChannelRequestOutputGroupsOutputs) *CreateMediaLiveChannelRequestOutputGroups {
	s.Outputs = v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroups) SetType(v string) *CreateMediaLiveChannelRequestOutputGroups {
	s.Type = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroups) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting struct {
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

func (s CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) SetChannelName(v string) *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting {
	s.ChannelName = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) SetGroupName(v string) *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting {
	s.GroupName = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsMediaPackageGroupSetting) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestOutputGroupsOutputs struct {
	// The referenced AudioSettings.
	AudioSettingNames []*string `json:"AudioSettingNames,omitempty" xml:"AudioSettingNames,omitempty" type:"Repeated"`
	// The settings of the output delivered to MediaPackage.
	MediaPackageOutputSetting *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting `json:"MediaPackageOutputSetting,omitempty" xml:"MediaPackageOutputSetting,omitempty" type:"Struct"`
	// The media type of the output. Valid values:
	//
	// 	- 0: Audio and Video.
	//
	// 	- 1: Audio. If you set the value to 1, you cannot reference VideoSettings.
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

func (s CreateMediaLiveChannelRequestOutputGroupsOutputs) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestOutputGroupsOutputs) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) GetAudioSettingNames() []*string {
	return s.AudioSettingNames
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) GetMediaPackageOutputSetting() *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting {
	return s.MediaPackageOutputSetting
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) GetMediaType() *int32 {
	return s.MediaType
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) GetVideoSettingName() *string {
	return s.VideoSettingName
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) SetAudioSettingNames(v []*string) *CreateMediaLiveChannelRequestOutputGroupsOutputs {
	s.AudioSettingNames = v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) SetMediaPackageOutputSetting(v *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) *CreateMediaLiveChannelRequestOutputGroupsOutputs {
	s.MediaPackageOutputSetting = v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) SetMediaType(v int32) *CreateMediaLiveChannelRequestOutputGroupsOutputs {
	s.MediaType = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) SetName(v string) *CreateMediaLiveChannelRequestOutputGroupsOutputs {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) SetVideoSettingName(v string) *CreateMediaLiveChannelRequestOutputGroupsOutputs {
	s.VideoSettingName = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputs) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting struct {
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

func (s CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) GetAudioGroupId() *string {
	return s.AudioGroupId
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) GetNameModifier() *string {
	return s.NameModifier
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) SetAudioGroupId(v string) *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting {
	s.AudioGroupId = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) SetNameModifier(v string) *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting {
	s.NameModifier = &v
	return s
}

func (s *CreateMediaLiveChannelRequestOutputGroupsOutputsMediaPackageOutputSetting) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestVideoSettings struct {
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
	VideoCodecSetting *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting `json:"VideoCodecSetting,omitempty" xml:"VideoCodecSetting,omitempty" type:"Struct"`
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

func (s CreateMediaLiveChannelRequestVideoSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestVideoSettings) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestVideoSettings) GetHeight() *int32 {
	return s.Height
}

func (s *CreateMediaLiveChannelRequestVideoSettings) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelRequestVideoSettings) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *CreateMediaLiveChannelRequestVideoSettings) GetVideoCodecSetting() *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	return s.VideoCodecSetting
}

func (s *CreateMediaLiveChannelRequestVideoSettings) GetVideoCodecType() *string {
	return s.VideoCodecType
}

func (s *CreateMediaLiveChannelRequestVideoSettings) GetWidth() *int32 {
	return s.Width
}

func (s *CreateMediaLiveChannelRequestVideoSettings) SetHeight(v int32) *CreateMediaLiveChannelRequestVideoSettings {
	s.Height = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettings) SetName(v string) *CreateMediaLiveChannelRequestVideoSettings {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettings) SetVideoCodec(v string) *CreateMediaLiveChannelRequestVideoSettings {
	s.VideoCodec = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettings) SetVideoCodecSetting(v *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) *CreateMediaLiveChannelRequestVideoSettings {
	s.VideoCodecSetting = v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettings) SetVideoCodecType(v string) *CreateMediaLiveChannelRequestVideoSettings {
	s.VideoCodecType = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettings) SetWidth(v int32) *CreateMediaLiveChannelRequestVideoSettings {
	s.Width = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettings) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting struct {
	// The video encoding settings.
	CodecDetail *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail `json:"CodecDetail,omitempty" xml:"CodecDetail,omitempty" type:"Struct"`
	// The frame rate. If it is not specified, the source specification is used.
	Framerate *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate `json:"Framerate,omitempty" xml:"Framerate,omitempty" type:"Struct"`
	// The GOP setting. If it is not specified, the source specification is used.
	Gop *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop `json:"Gop,omitempty" xml:"Gop,omitempty" type:"Struct"`
	// The video encoding rate. If it is not specified, the source specification is used.
	Rate *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate `json:"Rate,omitempty" xml:"Rate,omitempty" type:"Struct"`
}

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetCodecDetail() *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail {
	return s.CodecDetail
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetFramerate() *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	return s.Framerate
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetGop() *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	return s.Gop
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) GetRate() *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	return s.Rate
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetCodecDetail(v *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.CodecDetail = v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetFramerate(v *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.Framerate = v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetGop(v *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.Gop = v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) SetRate(v *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting {
	s.Rate = v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSetting) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail struct {
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

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) GetLevel() *string {
	return s.Level
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) GetProfile() *string {
	return s.Profile
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) SetLevel(v string) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail {
	s.Level = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) SetProfile(v string) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail {
	s.Profile = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingCodecDetail) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate struct {
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

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GetFramerateControl() *string {
	return s.FramerateControl
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GetFramerateDenominator() *int32 {
	return s.FramerateDenominator
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) GetFramerateNumerator() *int32 {
	return s.FramerateNumerator
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) SetFramerateControl(v string) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	s.FramerateControl = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) SetFramerateDenominator(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	s.FramerateDenominator = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) SetFramerateNumerator(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate {
	s.FramerateNumerator = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingFramerate) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop struct {
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

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GetBframesNum() *int32 {
	return s.BframesNum
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GetGopSize() *int32 {
	return s.GopSize
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) GetGopSizeUnits() *string {
	return s.GopSizeUnits
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) SetBframesNum(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	s.BframesNum = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) SetGopSize(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	s.GopSize = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) SetGopSizeUnits(v string) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop {
	s.GopSizeUnits = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingGop) Validate() error {
	return dara.Validate(s)
}

type CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate struct {
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

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetBufferSize() *int32 {
	return s.BufferSize
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) GetRateControlMode() *string {
	return s.RateControlMode
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetBitrate(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.Bitrate = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetBufferSize(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.BufferSize = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetMaxBitrate(v int32) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.MaxBitrate = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) SetRateControlMode(v string) *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate {
	s.RateControlMode = &v
	return s
}

func (s *CreateMediaLiveChannelRequestVideoSettingsVideoCodecSettingRate) Validate() error {
	return dara.Validate(s)
}
