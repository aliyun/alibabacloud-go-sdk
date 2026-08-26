// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSwitchUrgentConfig(v string) *SetCasterConfigRequest
	GetAutoSwitchUrgentConfig() *string
	SetAutoSwitchUrgentOn(v bool) *SetCasterConfigRequest
	GetAutoSwitchUrgentOn() *bool
	SetCallbackUrl(v string) *SetCasterConfigRequest
	GetCallbackUrl() *string
	SetCasterId(v string) *SetCasterConfigRequest
	GetCasterId() *string
	SetCasterName(v string) *SetCasterConfigRequest
	GetCasterName() *string
	SetChannelEnable(v int32) *SetCasterConfigRequest
	GetChannelEnable() *int32
	SetDelay(v float32) *SetCasterConfigRequest
	GetDelay() *float32
	SetDomainName(v string) *SetCasterConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetCasterConfigRequest
	GetOwnerId() *int64
	SetProgramEffect(v int32) *SetCasterConfigRequest
	GetProgramEffect() *int32
	SetProgramName(v string) *SetCasterConfigRequest
	GetProgramName() *string
	SetRecordConfig(v string) *SetCasterConfigRequest
	GetRecordConfig() *string
	SetRegionId(v string) *SetCasterConfigRequest
	GetRegionId() *string
	SetSideOutputUrl(v string) *SetCasterConfigRequest
	GetSideOutputUrl() *string
	SetSideOutputUrlList(v string) *SetCasterConfigRequest
	GetSideOutputUrlList() *string
	SetSyncGroupsConfig(v string) *SetCasterConfigRequest
	GetSyncGroupsConfig() *string
	SetTranscodeConfig(v string) *SetCasterConfigRequest
	GetTranscodeConfig() *string
	SetUrgentImageId(v string) *SetCasterConfigRequest
	GetUrgentImageId() *string
	SetUrgentImageUrl(v string) *SetCasterConfigRequest
	GetUrgentImageUrl() *string
	SetUrgentLiveStreamUrl(v string) *SetCasterConfigRequest
	GetUrgentLiveStreamUrl() *string
	SetUrgentMaterialId(v string) *SetCasterConfigRequest
	GetUrgentMaterialId() *string
}

type SetCasterConfigRequest struct {
	// The automatic standby switchover configuration.
	//
	// `eofThres`: the duration of stream interruption after which the system automatically switches to the standby video, in seconds.
	//
	// example:
	//
	// {"eofThres":3}
	AutoSwitchUrgentConfig *string `json:"AutoSwitchUrgentConfig,omitempty" xml:"AutoSwitchUrgentConfig,omitempty"`
	// Specifies whether to enable automatic switchover to the standby video when the stream is interrupted.
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	AutoSwitchUrgentOn *bool `json:"AutoSwitchUrgentOn,omitempty" xml:"AutoSwitchUrgentOn,omitempty"`
	// The callback URL. To receive callback notifications, enter a valid receiving address that accepts the HTTP protocol. If this parameter is set to empty, callback notifications for the production studio are canceled by default.
	//
	// > For more information about production studio callbacks, see [Cloud production studio callback information](https://help.aliyun.com/document_detail/213633.html).
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page of the ApsaraVideo Live console is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The name of the production studio.
	//
	// example:
	//
	// liveCaster****
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	// Specifies whether to enable Channel. If Channel was previously enabled (ChannelEnable=1), you must explicitly pass ChannelEnable=1 in each call to maintain the channel status. Otherwise, the error InvalidCaster.ChannelDisableUnsupported is returned.
	//
	//
	//
	// - **0*	- (default): disabled.
	//
	// - **1**: enabled.
	//
	// > Channel is disabled by default and cannot be disabled after it is enabled. When Channel is disabled, resources are directly referenced by layouts. To enable Channel for the first time, the production studio must be stopped. Existing layouts are discarded. Resources must first be assigned to a Channel, and new layouts directly reference the Channel. Through Channel, you can adjust the playback progress and status of video sources. In this mode, if the video source, PVW, and PGM areas reference the same resource, the corresponding views remain synchronized.
	//
	// example:
	//
	// 1
	ChannelEnable *int32 `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty"`
	// The stream delay, in seconds.
	//
	// - **0*	- (default): disables stream delay.
	//
	// - Greater than **0**: enables stream delay.
	//
	// - **Empty**: clears the stream delay configuration by default.
	//
	// > The maximum value is 300 seconds.
	//
	// example:
	//
	// 0
	Delay *float32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The primary streaming domain.
	//
	// Complete the domain name configuration before starting the production studio. If this parameter is empty, the domain name configuration of the production studio is cleared by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether the program list takes effect.
	//
	// - **0**: does not take effect.
	//
	// - **1**: takes effect.
	//
	// example:
	//
	// 1
	ProgramEffect *int32 `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty"`
	// The name of the program list. This parameter can be configured when the program list feature is used.
	//
	// example:
	//
	// program_name
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The recording configuration in JSON format. The configuration elements are as follows:
	//
	// - **endpoint**: the API endpoint of the Alibaba Cloud service.
	//
	// - **ossBucket**: the name of the OSS bucket.
	//
	// - **videoFormat**: the video file formats supported for export. Example: `[{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"m3u8\\",\\"CycleDuration\\":21600,\\"SliceOssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{UnixTimestamp}\\"},{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"flv\\",\\"CycleDuration\\":21600}]`.
	//
	// - **interval**: the time interval, in milliseconds (ms).
	//
	// >If this parameter is set to empty, the recording feature is not enabled. If this parameter is set to empty, the recording configuration is cleared by default.
	//
	// example:
	//
	// { "endpoint": "http://oss-cn-********.aliyuncs.com/api",  "ossBucket****": "liveBucket****", "VideoFormat":[{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"m3u8\\",\\"CycleDuration\\":21600,\\"SliceOssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{UnixTimestamp}\\"},{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"flv\\",\\"CycleDuration\\":21600}] "interval": 5 }
	RecordConfig *string `json:"RecordConfig,omitempty" xml:"RecordConfig,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ingest URL that corresponds to the custom bypass output address of the production studio.
	//
	// If this parameter is empty, the ingest URL that corresponds to the output address automatically generated by Alibaba Cloud is used by default.
	//
	// > Currently, SideOutputUrl supports only the RTMP protocol for stream ingest.
	SideOutputUrl *string `json:"SideOutputUrl,omitempty" xml:"SideOutputUrl,omitempty"`
	// The list of multi-destination relay streaming addresses. The addresses can be CDN ingest URLs from Alibaba Cloud or third-party providers. A maximum of 20 RTMP relay addresses can be added to a production studio.
	//
	//
	// > Specify multiple addresses in the array format: ["rtmp://domain/app1/stream1","rtmp://domain/app2/stream2"].
	//
	// example:
	//
	// rtmp://domain/app/stream?***
	SideOutputUrlList *string `json:"SideOutputUrlList,omitempty" xml:"SideOutputUrlList,omitempty"`
	// The multi-view synchronization configuration that synchronizes multiple video sources.
	//
	// Multi-view synchronization has two modes:
	//
	// - mode: 0 (streamer mode. Multiple video sources are synchronized based on the specified mode.)
	//
	// - mode: 1 (conference mode. There is no concept of a streamer video. All video sources are synchronized with each other.)
	//
	//
	//
	// Streamer mode: hostResourceId: the streamer video source in streamer mode.
	//
	// Conference mode: the hostResourceId field is not required. Only the resource IDs in resourceIds need to be provided.
	//
	// example:
	//
	// "[{\\"mode\\":0,\\"resourceIds\\":[\\"5a6c1c33-8424-46f6-813c-c152220a****\\",\\"4e6521dc-a40a-4077-b6bf-1fb12a76****\\"],\\"hostResourceId\\":\\"3aa2b39a-fd0e-4b8c-be73-b7af31c4****\\"}]"
	SyncGroupsConfig *string `json:"SyncGroupsConfig,omitempty" xml:"SyncGroupsConfig,omitempty"`
	// The transcoding configuration.
	//
	// A JSON-formatted string. Use upper camel case for internal fields of the struct. If this parameter is set to empty, the transcoding configuration is cleared by default. If the transcoding template is empty, an error is returned when the production studio starts.
	//
	// example:
	//
	// {"casterTemplate": "lp_ld"}
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty"`
	// The media asset ID of the standby image in the media library.
	//
	// example:
	//
	// a089175eb5f4427684fc0715159a****
	UrgentImageId *string `json:"UrgentImageId,omitempty" xml:"UrgentImageId,omitempty"`
	// The URL of the standby image.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/image.jpg
	UrgentImageUrl *string `json:"UrgentImageUrl,omitempty" xml:"UrgentImageUrl,omitempty"`
	// The URL of the standby live stream.
	//
	// example:
	//
	// rtmp://demo.aliyundoc.com
	UrgentLiveStreamUrl *string `json:"UrgentLiveStreamUrl,omitempty" xml:"UrgentLiveStreamUrl,omitempty"`
	// The media asset ID of the standby video in the media library. If this parameter is set to empty, the standby configuration is cleared by default.
	//
	// example:
	//
	// a2b8e671
	UrgentMaterialId *string `json:"UrgentMaterialId,omitempty" xml:"UrgentMaterialId,omitempty"`
}

func (s SetCasterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCasterConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCasterConfigRequest) GetAutoSwitchUrgentConfig() *string {
	return s.AutoSwitchUrgentConfig
}

func (s *SetCasterConfigRequest) GetAutoSwitchUrgentOn() *bool {
	return s.AutoSwitchUrgentOn
}

func (s *SetCasterConfigRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SetCasterConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *SetCasterConfigRequest) GetCasterName() *string {
	return s.CasterName
}

func (s *SetCasterConfigRequest) GetChannelEnable() *int32 {
	return s.ChannelEnable
}

func (s *SetCasterConfigRequest) GetDelay() *float32 {
	return s.Delay
}

func (s *SetCasterConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetCasterConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCasterConfigRequest) GetProgramEffect() *int32 {
	return s.ProgramEffect
}

func (s *SetCasterConfigRequest) GetProgramName() *string {
	return s.ProgramName
}

func (s *SetCasterConfigRequest) GetRecordConfig() *string {
	return s.RecordConfig
}

func (s *SetCasterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetCasterConfigRequest) GetSideOutputUrl() *string {
	return s.SideOutputUrl
}

func (s *SetCasterConfigRequest) GetSideOutputUrlList() *string {
	return s.SideOutputUrlList
}

func (s *SetCasterConfigRequest) GetSyncGroupsConfig() *string {
	return s.SyncGroupsConfig
}

func (s *SetCasterConfigRequest) GetTranscodeConfig() *string {
	return s.TranscodeConfig
}

func (s *SetCasterConfigRequest) GetUrgentImageId() *string {
	return s.UrgentImageId
}

func (s *SetCasterConfigRequest) GetUrgentImageUrl() *string {
	return s.UrgentImageUrl
}

func (s *SetCasterConfigRequest) GetUrgentLiveStreamUrl() *string {
	return s.UrgentLiveStreamUrl
}

func (s *SetCasterConfigRequest) GetUrgentMaterialId() *string {
	return s.UrgentMaterialId
}

func (s *SetCasterConfigRequest) SetAutoSwitchUrgentConfig(v string) *SetCasterConfigRequest {
	s.AutoSwitchUrgentConfig = &v
	return s
}

func (s *SetCasterConfigRequest) SetAutoSwitchUrgentOn(v bool) *SetCasterConfigRequest {
	s.AutoSwitchUrgentOn = &v
	return s
}

func (s *SetCasterConfigRequest) SetCallbackUrl(v string) *SetCasterConfigRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SetCasterConfigRequest) SetCasterId(v string) *SetCasterConfigRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterConfigRequest) SetCasterName(v string) *SetCasterConfigRequest {
	s.CasterName = &v
	return s
}

func (s *SetCasterConfigRequest) SetChannelEnable(v int32) *SetCasterConfigRequest {
	s.ChannelEnable = &v
	return s
}

func (s *SetCasterConfigRequest) SetDelay(v float32) *SetCasterConfigRequest {
	s.Delay = &v
	return s
}

func (s *SetCasterConfigRequest) SetDomainName(v string) *SetCasterConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetCasterConfigRequest) SetOwnerId(v int64) *SetCasterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCasterConfigRequest) SetProgramEffect(v int32) *SetCasterConfigRequest {
	s.ProgramEffect = &v
	return s
}

func (s *SetCasterConfigRequest) SetProgramName(v string) *SetCasterConfigRequest {
	s.ProgramName = &v
	return s
}

func (s *SetCasterConfigRequest) SetRecordConfig(v string) *SetCasterConfigRequest {
	s.RecordConfig = &v
	return s
}

func (s *SetCasterConfigRequest) SetRegionId(v string) *SetCasterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetCasterConfigRequest) SetSideOutputUrl(v string) *SetCasterConfigRequest {
	s.SideOutputUrl = &v
	return s
}

func (s *SetCasterConfigRequest) SetSideOutputUrlList(v string) *SetCasterConfigRequest {
	s.SideOutputUrlList = &v
	return s
}

func (s *SetCasterConfigRequest) SetSyncGroupsConfig(v string) *SetCasterConfigRequest {
	s.SyncGroupsConfig = &v
	return s
}

func (s *SetCasterConfigRequest) SetTranscodeConfig(v string) *SetCasterConfigRequest {
	s.TranscodeConfig = &v
	return s
}

func (s *SetCasterConfigRequest) SetUrgentImageId(v string) *SetCasterConfigRequest {
	s.UrgentImageId = &v
	return s
}

func (s *SetCasterConfigRequest) SetUrgentImageUrl(v string) *SetCasterConfigRequest {
	s.UrgentImageUrl = &v
	return s
}

func (s *SetCasterConfigRequest) SetUrgentLiveStreamUrl(v string) *SetCasterConfigRequest {
	s.UrgentLiveStreamUrl = &v
	return s
}

func (s *SetCasterConfigRequest) SetUrgentMaterialId(v string) *SetCasterConfigRequest {
	s.UrgentMaterialId = &v
	return s
}

func (s *SetCasterConfigRequest) Validate() error {
	return dara.Validate(s)
}
