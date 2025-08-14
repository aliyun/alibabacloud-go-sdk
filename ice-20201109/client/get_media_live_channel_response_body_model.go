// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *GetMediaLiveChannelResponseBodyChannel) *GetMediaLiveChannelResponseBody
	GetChannel() *GetMediaLiveChannelResponseBodyChannel
	SetRequestId(v string) *GetMediaLiveChannelResponseBody
	GetRequestId() *string
}

type GetMediaLiveChannelResponseBody struct {
	// The channel information.
	Channel *GetMediaLiveChannelResponseBodyChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBody) GetChannel() *GetMediaLiveChannelResponseBodyChannel {
	return s.Channel
}

func (s *GetMediaLiveChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaLiveChannelResponseBody) SetChannel(v *GetMediaLiveChannelResponseBodyChannel) *GetMediaLiveChannelResponseBody {
	s.Channel = v
	return s
}

func (s *GetMediaLiveChannelResponseBody) SetRequestId(v string) *GetMediaLiveChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaLiveChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannel struct {
	// The audio settings.
	AudioSettings []*GetMediaLiveChannelResponseBodyChannelAudioSettings `json:"AudioSettings,omitempty" xml:"AudioSettings,omitempty" type:"Repeated"`
	// The ID of the channel.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The time when the channel was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-12-03T06:56:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The inputs associated with the channel.
	InputAttachments []*GetMediaLiveChannelResponseBodyChannelInputAttachments `json:"InputAttachments,omitempty" xml:"InputAttachments,omitempty" type:"Repeated"`
	// The time when the channel was last started. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. If the channel has never been started since it was created, an empty string is returned.
	//
	// example:
	//
	// 2024-12-03T06:56:42Z
	LastStartTime *string `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	// The time when the channel was last stopped. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. If the channel has never stopped since it was created, an empty string is returned.
	//
	// example:
	//
	// 2024-12-03T06:56:42Z
	LastStopTime *string `json:"LastStopTime,omitempty" xml:"LastStopTime,omitempty"`
	// The channel name.
	//
	// example:
	//
	// mych
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output groups.
	OutputGroups []*GetMediaLiveChannelResponseBodyChannelOutputGroups `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	// The state of the channel. Valid values: IDLE, STARTING, RUNNING, RECOVERING, and STOPPING.
	//
	// example:
	//
	// IDLE
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The video settings.
	VideoSettings []*GetMediaLiveChannelResponseBodyChannelVideoSettings `json:"VideoSettings,omitempty" xml:"VideoSettings,omitempty" type:"Repeated"`
}

func (s GetMediaLiveChannelResponseBodyChannel) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannel) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetAudioSettings() []*GetMediaLiveChannelResponseBodyChannelAudioSettings {
	return s.AudioSettings
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetInputAttachments() []*GetMediaLiveChannelResponseBodyChannelInputAttachments {
	return s.InputAttachments
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetLastStartTime() *string {
	return s.LastStartTime
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetLastStopTime() *string {
	return s.LastStopTime
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetOutputGroups() []*GetMediaLiveChannelResponseBodyChannelOutputGroups {
	return s.OutputGroups
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetState() *string {
	return s.State
}

func (s *GetMediaLiveChannelResponseBodyChannel) GetVideoSettings() []*GetMediaLiveChannelResponseBodyChannelVideoSettings {
	return s.VideoSettings
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetAudioSettings(v []*GetMediaLiveChannelResponseBodyChannelAudioSettings) *GetMediaLiveChannelResponseBodyChannel {
	s.AudioSettings = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetChannelId(v string) *GetMediaLiveChannelResponseBodyChannel {
	s.ChannelId = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetCreateTime(v string) *GetMediaLiveChannelResponseBodyChannel {
	s.CreateTime = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetInputAttachments(v []*GetMediaLiveChannelResponseBodyChannelInputAttachments) *GetMediaLiveChannelResponseBodyChannel {
	s.InputAttachments = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetLastStartTime(v string) *GetMediaLiveChannelResponseBodyChannel {
	s.LastStartTime = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetLastStopTime(v string) *GetMediaLiveChannelResponseBodyChannel {
	s.LastStopTime = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetName(v string) *GetMediaLiveChannelResponseBodyChannel {
	s.Name = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetOutputGroups(v []*GetMediaLiveChannelResponseBodyChannelOutputGroups) *GetMediaLiveChannelResponseBodyChannel {
	s.OutputGroups = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetState(v string) *GetMediaLiveChannelResponseBodyChannel {
	s.State = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) SetVideoSettings(v []*GetMediaLiveChannelResponseBodyChannelVideoSettings) *GetMediaLiveChannelResponseBodyChannel {
	s.VideoSettings = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannel) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelAudioSettings struct {
	// The audio codec.
	//
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio encoding settings.
	AudioCodecSetting *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting `json:"AudioCodecSetting,omitempty" xml:"AudioCodecSetting,omitempty" type:"Struct"`
	// The name of the audio selector.
	//
	// example:
	//
	// myselector
	AudioSelectorName *string `json:"AudioSelectorName,omitempty" xml:"AudioSelectorName,omitempty"`
	// A three-letter ISO 639-2 language code.
	//
	// example:
	//
	// eng
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// The name of the language.
	//
	// example:
	//
	// English
	LanguageName *string `json:"LanguageName,omitempty" xml:"LanguageName,omitempty"`
	// The name of the audio settings.
	//
	// example:
	//
	// zhuanfengzhuang
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelAudioSettings) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelAudioSettings) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) GetAudioCodecSetting() *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting {
	return s.AudioCodecSetting
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) GetAudioSelectorName() *string {
	return s.AudioSelectorName
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) GetLanguageName() *string {
	return s.LanguageName
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) SetAudioCodec(v string) *GetMediaLiveChannelResponseBodyChannelAudioSettings {
	s.AudioCodec = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) SetAudioCodecSetting(v *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) *GetMediaLiveChannelResponseBodyChannelAudioSettings {
	s.AudioCodecSetting = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) SetAudioSelectorName(v string) *GetMediaLiveChannelResponseBodyChannelAudioSettings {
	s.AudioSelectorName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) SetLanguageCode(v string) *GetMediaLiveChannelResponseBodyChannelAudioSettings {
	s.LanguageCode = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) SetLanguageName(v string) *GetMediaLiveChannelResponseBodyChannelAudioSettings {
	s.LanguageName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) SetName(v string) *GetMediaLiveChannelResponseBodyChannelAudioSettings {
	s.Name = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettings) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting struct {
	// The audio bitrate. Unit: bit/s.
	//
	// example:
	//
	// 200000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The audio codec profile.
	//
	// example:
	//
	// AAC-LOW
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sample rate. Unit: Hz.
	//
	// example:
	//
	// 44100
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) SetBitrate(v int32) *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting {
	s.Bitrate = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) SetProfile(v string) *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting {
	s.Profile = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) SetSampleRate(v int32) *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting {
	s.SampleRate = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelAudioSettingsAudioCodecSetting) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelInputAttachments struct {
	// The audio selectors.
	AudioSelectors []*GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors `json:"AudioSelectors,omitempty" xml:"AudioSelectors,omitempty" type:"Repeated"`
	// The ID of the associated input.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// The name of the input.
	//
	// example:
	//
	// myinput
	InputName *string `json:"InputName,omitempty" xml:"InputName,omitempty"`
	// The language name.
	//
	// example:
	//
	// eng
	LanguageName *string `json:"LanguageName,omitempty" xml:"LanguageName,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachments) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachments) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) GetAudioSelectors() []*GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors {
	return s.AudioSelectors
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) GetInputId() *string {
	return s.InputId
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) GetInputName() *string {
	return s.InputName
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) GetLanguageName() *string {
	return s.LanguageName
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) SetAudioSelectors(v []*GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) *GetMediaLiveChannelResponseBodyChannelInputAttachments {
	s.AudioSelectors = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) SetInputId(v string) *GetMediaLiveChannelResponseBodyChannelInputAttachments {
	s.InputId = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) SetInputName(v string) *GetMediaLiveChannelResponseBodyChannelInputAttachments {
	s.InputName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) SetLanguageName(v string) *GetMediaLiveChannelResponseBodyChannelInputAttachments {
	s.LanguageName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachments) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors struct {
	// The audio language selection.
	AudioLanguageSelection *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection `json:"AudioLanguageSelection,omitempty" xml:"AudioLanguageSelection,omitempty" type:"Struct"`
	// The audio PID selection.
	AudioPidSelection *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection `json:"AudioPidSelection,omitempty" xml:"AudioPidSelection,omitempty" type:"Struct"`
	// The audio track selection.
	AudioTrackSelection []*GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection `json:"AudioTrackSelection,omitempty" xml:"AudioTrackSelection,omitempty" type:"Repeated"`
	// The name of the audio selector.
	//
	// This parameter is required.
	//
	// example:
	//
	// myselector
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) GetAudioLanguageSelection() *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection {
	return s.AudioLanguageSelection
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) GetAudioPidSelection() *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection {
	return s.AudioPidSelection
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) GetAudioTrackSelection() []*GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection {
	return s.AudioTrackSelection
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) SetAudioLanguageSelection(v *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors {
	s.AudioLanguageSelection = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) SetAudioPidSelection(v *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors {
	s.AudioPidSelection = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) SetAudioTrackSelection(v []*GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors {
	s.AudioTrackSelection = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) SetName(v string) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors {
	s.Name = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectors) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection struct {
	// A three-letter ISO 639-2 language code from within an audio source.
	//
	// This parameter is required.
	//
	// example:
	//
	// eng
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection) SetLanguageCode(v string) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection {
	s.LanguageCode = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioLanguageSelection) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection struct {
	// A PID from within a source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Pid *int64 `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection) GetPid() *int64 {
	return s.Pid
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection) SetPid(v int64) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection {
	s.Pid = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioPidSelection) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection struct {
	// The track ID from within a source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TrackId *int64 `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection) GetTrackId() *int64 {
	return s.TrackId
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection) SetTrackId(v int64) *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection {
	s.TrackId = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelInputAttachmentsAudioSelectorsAudioTrackSelection) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelOutputGroups struct {
	// The MediaPackage destination.
	MediaPackageGroupSetting *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting `json:"MediaPackageGroupSetting,omitempty" xml:"MediaPackageGroupSetting,omitempty" type:"Struct"`
	// The URL for monitoring the output group. The parameter is returned only when the output gourp type is MediaPackage.
	//
	// example:
	//
	// rtmp://xxx
	MonitorUrl *string `json:"MonitorUrl,omitempty" xml:"MonitorUrl,omitempty"`
	// The name of the output group.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The outputs in the output group.
	Outputs []*GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The output group type.
	//
	// example:
	//
	// MediaPackage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroups) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroups) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) GetMediaPackageGroupSetting() *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting {
	return s.MediaPackageGroupSetting
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) GetMonitorUrl() *string {
	return s.MonitorUrl
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) GetOutputs() []*GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs {
	return s.Outputs
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) GetType() *string {
	return s.Type
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) SetMediaPackageGroupSetting(v *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) *GetMediaLiveChannelResponseBodyChannelOutputGroups {
	s.MediaPackageGroupSetting = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) SetMonitorUrl(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroups {
	s.MonitorUrl = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) SetName(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroups {
	s.Name = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) SetOutputs(v []*GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) *GetMediaLiveChannelResponseBodyChannelOutputGroups {
	s.Outputs = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) SetType(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroups {
	s.Type = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroups) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting struct {
	// ChannelName in MediaPackage.
	//
	// example:
	//
	// myPackageChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// GroupName in MediaPackage.
	//
	// example:
	//
	// myPackageGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) GetGroupName() *string {
	return s.GroupName
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) SetChannelName(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting {
	s.ChannelName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) SetGroupName(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting {
	s.GroupName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsMediaPackageGroupSetting) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs struct {
	// The referenced AudioSettings.
	AudioSettingNames []*string `json:"AudioSettingNames,omitempty" xml:"AudioSettingNames,omitempty" type:"Repeated"`
	// The settings of the output delivered to MediaPackage.
	MediaPackageOutputSetting *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting `json:"MediaPackageOutputSetting,omitempty" xml:"MediaPackageOutputSetting,omitempty" type:"Struct"`
	// The media type of the output.
	//
	// example:
	//
	// 0
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The name of the output.
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

func (s GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) GetAudioSettingNames() []*string {
	return s.AudioSettingNames
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) GetMediaPackageOutputSetting() *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting {
	return s.MediaPackageOutputSetting
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) GetVideoSettingName() *string {
	return s.VideoSettingName
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) SetAudioSettingNames(v []*string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs {
	s.AudioSettingNames = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) SetMediaPackageOutputSetting(v *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs {
	s.MediaPackageOutputSetting = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) SetMediaType(v int32) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs {
	s.MediaType = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) SetName(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs {
	s.Name = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) SetVideoSettingName(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs {
	s.VideoSettingName = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputs) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting struct {
	// The manifest audio group ID.
	//
	// example:
	//
	// audiogroup
	AudioGroupId *string `json:"AudioGroupId,omitempty" xml:"AudioGroupId,omitempty"`
	// The manifest name modifier. The child manifests include this modifier in their M3U8 file names.
	//
	// example:
	//
	// 480p
	NameModifier *string `json:"NameModifier,omitempty" xml:"NameModifier,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) GetAudioGroupId() *string {
	return s.AudioGroupId
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) GetNameModifier() *string {
	return s.NameModifier
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) SetAudioGroupId(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting {
	s.AudioGroupId = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) SetNameModifier(v string) *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting {
	s.NameModifier = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelOutputGroupsOutputsMediaPackageOutputSetting) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelVideoSettings struct {
	// The height of the video in pixels.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the video settings.
	//
	// example:
	//
	// video1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The video codec.
	//
	// example:
	//
	// H264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// The video encoding settings.
	VideoCodecSetting *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting `json:"VideoCodecSetting,omitempty" xml:"VideoCodecSetting,omitempty" type:"Struct"`
	// The video transcoding method. Valid values: NORMAL (regular transcoding) and NBHD (Narrowband HD™ transcoding).
	//
	// example:
	//
	// NORMAL
	VideoCodecType *string `json:"VideoCodecType,omitempty" xml:"VideoCodecType,omitempty"`
	// The width of the video in pixels.
	//
	// example:
	//
	// 1280
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettings) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettings) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) GetHeight() *int32 {
	return s.Height
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) GetVideoCodecSetting() *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting {
	return s.VideoCodecSetting
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) GetVideoCodecType() *string {
	return s.VideoCodecType
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) GetWidth() *int32 {
	return s.Width
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) SetHeight(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettings {
	s.Height = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) SetName(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettings {
	s.Name = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) SetVideoCodec(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettings {
	s.VideoCodec = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) SetVideoCodecSetting(v *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) *GetMediaLiveChannelResponseBodyChannelVideoSettings {
	s.VideoCodecSetting = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) SetVideoCodecType(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettings {
	s.VideoCodecType = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) SetWidth(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettings {
	s.Width = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettings) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting struct {
	// The video encoding settings.
	CodecDetail *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail `json:"CodecDetail,omitempty" xml:"CodecDetail,omitempty" type:"Struct"`
	// The frame rate.
	Framerate *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate `json:"Framerate,omitempty" xml:"Framerate,omitempty" type:"Struct"`
	// The GOP setting.
	Gop *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop `json:"Gop,omitempty" xml:"Gop,omitempty" type:"Struct"`
	// The video encoding rate.
	Rate *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate `json:"Rate,omitempty" xml:"Rate,omitempty" type:"Struct"`
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) GetCodecDetail() *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail {
	return s.CodecDetail
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) GetFramerate() *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate {
	return s.Framerate
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) GetGop() *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop {
	return s.Gop
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) GetRate() *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate {
	return s.Rate
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) SetCodecDetail(v *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting {
	s.CodecDetail = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) SetFramerate(v *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting {
	s.Framerate = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) SetGop(v *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting {
	s.Gop = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) SetRate(v *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting {
	s.Rate = v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSetting) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail struct {
	// The video encoding level. It is not supported yet.
	//
	// example:
	//
	// H264_LEVEL_AUTO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The H.264 profile.
	//
	// example:
	//
	// MAIN
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) GetLevel() *string {
	return s.Level
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) SetLevel(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail {
	s.Level = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) SetProfile(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail {
	s.Profile = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingCodecDetail) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate struct {
	// The frame rate mode.
	//
	// example:
	//
	// SPECIFIED
	FramerateControl *string `json:"FramerateControl,omitempty" xml:"FramerateControl,omitempty"`
	// The denominator of the fixed frame rate.
	//
	// example:
	//
	// 1
	FramerateDenominator *int32 `json:"FramerateDenominator,omitempty" xml:"FramerateDenominator,omitempty"`
	// The numerator of the fixed frame rate.
	//
	// example:
	//
	// 25
	FramerateNumerator *int32 `json:"FramerateNumerator,omitempty" xml:"FramerateNumerator,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) GetFramerateControl() *string {
	return s.FramerateControl
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) GetFramerateDenominator() *int32 {
	return s.FramerateDenominator
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) GetFramerateNumerator() *int32 {
	return s.FramerateNumerator
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) SetFramerateControl(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate {
	s.FramerateControl = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) SetFramerateDenominator(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate {
	s.FramerateDenominator = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) SetFramerateNumerator(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate {
	s.FramerateNumerator = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingFramerate) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop struct {
	// The number of B frames.
	//
	// example:
	//
	// 3
	BframesNum *int32 `json:"BframesNum,omitempty" xml:"BframesNum,omitempty"`
	// The GOP size.
	//
	// example:
	//
	// 90
	GopSize *int32 `json:"GopSize,omitempty" xml:"GopSize,omitempty"`
	// The GOP size unit.
	//
	// example:
	//
	// FRAMES
	GopSizeUnits *string `json:"GopSizeUnits,omitempty" xml:"GopSizeUnits,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) GetBframesNum() *int32 {
	return s.BframesNum
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) GetGopSize() *int32 {
	return s.GopSize
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) GetGopSizeUnits() *string {
	return s.GopSizeUnits
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) SetBframesNum(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop {
	s.BframesNum = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) SetGopSize(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop {
	s.GopSize = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) SetGopSizeUnits(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop {
	s.GopSizeUnits = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingGop) Validate() error {
	return dara.Validate(s)
}

type GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate struct {
	// The video bitrate. Unit: bit/s.
	//
	// example:
	//
	// 2500000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video buffer size. Unit: bit/s.
	//
	// example:
	//
	// 6000000
	BufferSize *int32 `json:"BufferSize,omitempty" xml:"BufferSize,omitempty"`
	// The maximum bitrate. Unit: bit/s.
	//
	// example:
	//
	// 6000000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The bitrate control mode.
	//
	// example:
	//
	// ABR
	RateControlMode *string `json:"RateControlMode,omitempty" xml:"RateControlMode,omitempty"`
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) GetBufferSize() *int32 {
	return s.BufferSize
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) GetRateControlMode() *string {
	return s.RateControlMode
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) SetBitrate(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate {
	s.Bitrate = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) SetBufferSize(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate {
	s.BufferSize = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) SetMaxBitrate(v int32) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate {
	s.MaxBitrate = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) SetRateControlMode(v string) *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate {
	s.RateControlMode = &v
	return s
}

func (s *GetMediaLiveChannelResponseBodyChannelVideoSettingsVideoCodecSettingRate) Validate() error {
	return dara.Validate(s)
}
