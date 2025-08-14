// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v []*ListMediaLiveChannelsResponseBodyChannels) *ListMediaLiveChannelsResponseBody
	GetChannels() []*ListMediaLiveChannelsResponseBodyChannels
	SetMaxResults(v int32) *ListMediaLiveChannelsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMediaLiveChannelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMediaLiveChannelsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMediaLiveChannelsResponseBody
	GetTotalCount() *int32
}

type ListMediaLiveChannelsResponseBody struct {
	// The channels.
	Channels []*ListMediaLiveChannelsResponseBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMediaLiveChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBody) GetChannels() []*ListMediaLiveChannelsResponseBodyChannels {
	return s.Channels
}

func (s *ListMediaLiveChannelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaLiveChannelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaLiveChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaLiveChannelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMediaLiveChannelsResponseBody) SetChannels(v []*ListMediaLiveChannelsResponseBodyChannels) *ListMediaLiveChannelsResponseBody {
	s.Channels = v
	return s
}

func (s *ListMediaLiveChannelsResponseBody) SetMaxResults(v int32) *ListMediaLiveChannelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBody) SetNextToken(v string) *ListMediaLiveChannelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBody) SetRequestId(v string) *ListMediaLiveChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBody) SetTotalCount(v int32) *ListMediaLiveChannelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannels struct {
	// The audio settings.
	AudioSettings []*ListMediaLiveChannelsResponseBodyChannelsAudioSettings `json:"AudioSettings,omitempty" xml:"AudioSettings,omitempty" type:"Repeated"`
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
	InputAttachments []*ListMediaLiveChannelsResponseBodyChannelsInputAttachments `json:"InputAttachments,omitempty" xml:"InputAttachments,omitempty" type:"Repeated"`
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
	// The name of the channel.
	//
	// example:
	//
	// mych
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output groups.
	OutputGroups []*ListMediaLiveChannelsResponseBodyChannelsOutputGroups `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	// The state of the channel. Valid values: IDLE, STARTING, RUNNING, RECOVERING, and STOPPING.
	//
	// example:
	//
	// IDLE
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The video settings.
	VideoSettings []*ListMediaLiveChannelsResponseBodyChannelsVideoSettings `json:"VideoSettings,omitempty" xml:"VideoSettings,omitempty" type:"Repeated"`
}

func (s ListMediaLiveChannelsResponseBodyChannels) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetAudioSettings() []*ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	return s.AudioSettings
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetInputAttachments() []*ListMediaLiveChannelsResponseBodyChannelsInputAttachments {
	return s.InputAttachments
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetLastStartTime() *string {
	return s.LastStartTime
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetLastStopTime() *string {
	return s.LastStopTime
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetOutputGroups() []*ListMediaLiveChannelsResponseBodyChannelsOutputGroups {
	return s.OutputGroups
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetState() *string {
	return s.State
}

func (s *ListMediaLiveChannelsResponseBodyChannels) GetVideoSettings() []*ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	return s.VideoSettings
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetAudioSettings(v []*ListMediaLiveChannelsResponseBodyChannelsAudioSettings) *ListMediaLiveChannelsResponseBodyChannels {
	s.AudioSettings = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetChannelId(v string) *ListMediaLiveChannelsResponseBodyChannels {
	s.ChannelId = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetCreateTime(v string) *ListMediaLiveChannelsResponseBodyChannels {
	s.CreateTime = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetInputAttachments(v []*ListMediaLiveChannelsResponseBodyChannelsInputAttachments) *ListMediaLiveChannelsResponseBodyChannels {
	s.InputAttachments = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetLastStartTime(v string) *ListMediaLiveChannelsResponseBodyChannels {
	s.LastStartTime = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetLastStopTime(v string) *ListMediaLiveChannelsResponseBodyChannels {
	s.LastStopTime = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetName(v string) *ListMediaLiveChannelsResponseBodyChannels {
	s.Name = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetOutputGroups(v []*ListMediaLiveChannelsResponseBodyChannelsOutputGroups) *ListMediaLiveChannelsResponseBodyChannels {
	s.OutputGroups = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetState(v string) *ListMediaLiveChannelsResponseBodyChannels {
	s.State = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) SetVideoSettings(v []*ListMediaLiveChannelsResponseBodyChannelsVideoSettings) *ListMediaLiveChannelsResponseBodyChannels {
	s.VideoSettings = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannels) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsAudioSettings struct {
	// The audio codec.
	//
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio encoding settings.
	AudioCodecSetting *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting `json:"AudioCodecSetting,omitempty" xml:"AudioCodecSetting,omitempty" type:"Struct"`
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

func (s ListMediaLiveChannelsResponseBodyChannelsAudioSettings) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GetAudioCodecSetting() *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting {
	return s.AudioCodecSetting
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GetAudioSelectorName() *string {
	return s.AudioSelectorName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GetLanguageName() *string {
	return s.LanguageName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) SetAudioCodec(v string) *ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	s.AudioCodec = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) SetAudioCodecSetting(v *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) *ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	s.AudioCodecSetting = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) SetAudioSelectorName(v string) *ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	s.AudioSelectorName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) SetLanguageCode(v string) *ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	s.LanguageCode = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) SetLanguageName(v string) *ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	s.LanguageName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) SetName(v string) *ListMediaLiveChannelsResponseBodyChannelsAudioSettings {
	s.Name = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettings) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) GetProfile() *string {
	return s.Profile
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) SetBitrate(v int32) *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting {
	s.Bitrate = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) SetProfile(v string) *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting {
	s.Profile = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) SetSampleRate(v int32) *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting {
	s.SampleRate = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsAudioSettingsAudioCodecSetting) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsInputAttachments struct {
	// The audio selectors.
	AudioSelectors []*ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors `json:"AudioSelectors,omitempty" xml:"AudioSelectors,omitempty" type:"Repeated"`
	// The ID of the input.
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
	// The name of the language.
	//
	// example:
	//
	// eng
	LanguageName *string `json:"LanguageName,omitempty" xml:"LanguageName,omitempty"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachments) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) GetAudioSelectors() []*ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors {
	return s.AudioSelectors
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) GetInputId() *string {
	return s.InputId
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) GetInputName() *string {
	return s.InputName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) GetLanguageName() *string {
	return s.LanguageName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) SetAudioSelectors(v []*ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) *ListMediaLiveChannelsResponseBodyChannelsInputAttachments {
	s.AudioSelectors = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) SetInputId(v string) *ListMediaLiveChannelsResponseBodyChannelsInputAttachments {
	s.InputId = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) SetInputName(v string) *ListMediaLiveChannelsResponseBodyChannelsInputAttachments {
	s.InputName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) SetLanguageName(v string) *ListMediaLiveChannelsResponseBodyChannelsInputAttachments {
	s.LanguageName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachments) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors struct {
	// The audio language selection.
	AudioLanguageSelection *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection `json:"AudioLanguageSelection,omitempty" xml:"AudioLanguageSelection,omitempty" type:"Struct"`
	// The audio PID selection.
	AudioPidSelection *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection `json:"AudioPidSelection,omitempty" xml:"AudioPidSelection,omitempty" type:"Struct"`
	// The audio track selection.
	AudioTrackSelection []*ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection `json:"AudioTrackSelection,omitempty" xml:"AudioTrackSelection,omitempty" type:"Repeated"`
	// The name of the audio selector.
	//
	// This parameter is required.
	//
	// example:
	//
	// myselector
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) GetAudioLanguageSelection() *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection {
	return s.AudioLanguageSelection
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) GetAudioPidSelection() *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection {
	return s.AudioPidSelection
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) GetAudioTrackSelection() []*ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection {
	return s.AudioTrackSelection
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) SetAudioLanguageSelection(v *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors {
	s.AudioLanguageSelection = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) SetAudioPidSelection(v *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors {
	s.AudioPidSelection = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) SetAudioTrackSelection(v []*ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors {
	s.AudioTrackSelection = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) SetName(v string) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors {
	s.Name = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectors) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection struct {
	// A three-letter ISO 639-2 language code from within an audio source.
	//
	// This parameter is required.
	//
	// example:
	//
	// eng
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection) SetLanguageCode(v string) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection {
	s.LanguageCode = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioLanguageSelection) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection struct {
	// A PID from within a source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Pid *int64 `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection) GetPid() *int64 {
	return s.Pid
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection) SetPid(v int64) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection {
	s.Pid = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioPidSelection) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection struct {
	// The track ID from within a source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TrackId *int64 `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection) GetTrackId() *int64 {
	return s.TrackId
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection) SetTrackId(v int64) *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection {
	s.TrackId = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsInputAttachmentsAudioSelectorsAudioTrackSelection) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsOutputGroups struct {
	// The MediaPackage destination.
	MediaPackageGroupSetting *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting `json:"MediaPackageGroupSetting,omitempty" xml:"MediaPackageGroupSetting,omitempty" type:"Struct"`
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
	Outputs []*ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The output group type.
	//
	// example:
	//
	// MediaPackage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroups) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroups) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) GetMediaPackageGroupSetting() *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting {
	return s.MediaPackageGroupSetting
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) GetMonitorUrl() *string {
	return s.MonitorUrl
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) GetOutputs() []*ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs {
	return s.Outputs
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) GetType() *string {
	return s.Type
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) SetMediaPackageGroupSetting(v *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) *ListMediaLiveChannelsResponseBodyChannelsOutputGroups {
	s.MediaPackageGroupSetting = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) SetMonitorUrl(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroups {
	s.MonitorUrl = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) SetName(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroups {
	s.Name = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) SetOutputs(v []*ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) *ListMediaLiveChannelsResponseBodyChannelsOutputGroups {
	s.Outputs = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) SetType(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroups {
	s.Type = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroups) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) GetGroupName() *string {
	return s.GroupName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) SetChannelName(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting {
	s.ChannelName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) SetGroupName(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting {
	s.GroupName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsMediaPackageGroupSetting) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs struct {
	// The referenced AudioSettings.
	AudioSettingNames []*string `json:"AudioSettingNames,omitempty" xml:"AudioSettingNames,omitempty" type:"Repeated"`
	// The settings of the output delivered to MediaPackage.
	MediaPackageOutputSetting *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting `json:"MediaPackageOutputSetting,omitempty" xml:"MediaPackageOutputSetting,omitempty" type:"Struct"`
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

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) GetAudioSettingNames() []*string {
	return s.AudioSettingNames
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) GetMediaPackageOutputSetting() *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting {
	return s.MediaPackageOutputSetting
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) GetMediaType() *int32 {
	return s.MediaType
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) GetVideoSettingName() *string {
	return s.VideoSettingName
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) SetAudioSettingNames(v []*string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs {
	s.AudioSettingNames = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) SetMediaPackageOutputSetting(v *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs {
	s.MediaPackageOutputSetting = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) SetMediaType(v int32) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs {
	s.MediaType = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) SetName(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs {
	s.Name = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) SetVideoSettingName(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs {
	s.VideoSettingName = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputs) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) GetAudioGroupId() *string {
	return s.AudioGroupId
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) GetNameModifier() *string {
	return s.NameModifier
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) SetAudioGroupId(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting {
	s.AudioGroupId = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) SetNameModifier(v string) *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting {
	s.NameModifier = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsOutputGroupsOutputsMediaPackageOutputSetting) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsVideoSettings struct {
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
	VideoCodecSetting *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting `json:"VideoCodecSetting,omitempty" xml:"VideoCodecSetting,omitempty" type:"Struct"`
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

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettings) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GetHeight() *int32 {
	return s.Height
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GetName() *string {
	return s.Name
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GetVideoCodecSetting() *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting {
	return s.VideoCodecSetting
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GetVideoCodecType() *string {
	return s.VideoCodecType
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) GetWidth() *int32 {
	return s.Width
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) SetHeight(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	s.Height = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) SetName(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	s.Name = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) SetVideoCodec(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	s.VideoCodec = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) SetVideoCodecSetting(v *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) *ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	s.VideoCodecSetting = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) SetVideoCodecType(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	s.VideoCodecType = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) SetWidth(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettings {
	s.Width = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettings) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting struct {
	// The video encoding settings.
	CodecDetail *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail `json:"CodecDetail,omitempty" xml:"CodecDetail,omitempty" type:"Struct"`
	// The frame rate.
	Framerate *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate `json:"Framerate,omitempty" xml:"Framerate,omitempty" type:"Struct"`
	// The GOP setting.
	Gop *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop `json:"Gop,omitempty" xml:"Gop,omitempty" type:"Struct"`
	// The video encoding rate.
	Rate *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate `json:"Rate,omitempty" xml:"Rate,omitempty" type:"Struct"`
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) GetCodecDetail() *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail {
	return s.CodecDetail
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) GetFramerate() *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate {
	return s.Framerate
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) GetGop() *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop {
	return s.Gop
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) GetRate() *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate {
	return s.Rate
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) SetCodecDetail(v *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting {
	s.CodecDetail = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) SetFramerate(v *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting {
	s.Framerate = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) SetGop(v *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting {
	s.Gop = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) SetRate(v *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting {
	s.Rate = v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSetting) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) GetLevel() *string {
	return s.Level
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) GetProfile() *string {
	return s.Profile
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) SetLevel(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail {
	s.Level = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) SetProfile(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail {
	s.Profile = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingCodecDetail) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) GetFramerateControl() *string {
	return s.FramerateControl
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) GetFramerateDenominator() *int32 {
	return s.FramerateDenominator
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) GetFramerateNumerator() *int32 {
	return s.FramerateNumerator
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) SetFramerateControl(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate {
	s.FramerateControl = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) SetFramerateDenominator(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate {
	s.FramerateDenominator = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) SetFramerateNumerator(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate {
	s.FramerateNumerator = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingFramerate) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) GetBframesNum() *int32 {
	return s.BframesNum
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) GetGopSize() *int32 {
	return s.GopSize
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) GetGopSizeUnits() *string {
	return s.GopSizeUnits
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) SetBframesNum(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop {
	s.BframesNum = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) SetGopSize(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop {
	s.GopSize = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) SetGopSizeUnits(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop {
	s.GopSizeUnits = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingGop) Validate() error {
	return dara.Validate(s)
}

type ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate struct {
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

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) GoString() string {
	return s.String()
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) GetBufferSize() *int32 {
	return s.BufferSize
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) GetRateControlMode() *string {
	return s.RateControlMode
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) SetBitrate(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate {
	s.Bitrate = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) SetBufferSize(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate {
	s.BufferSize = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) SetMaxBitrate(v int32) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate {
	s.MaxBitrate = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) SetRateControlMode(v string) *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate {
	s.RateControlMode = &v
	return s
}

func (s *ListMediaLiveChannelsResponseBodyChannelsVideoSettingsVideoCodecSettingRate) Validate() error {
	return dara.Validate(s)
}
