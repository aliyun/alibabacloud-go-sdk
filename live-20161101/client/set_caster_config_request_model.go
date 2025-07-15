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
	// The configuration for automatic switchover to the standby resource.
	//
	// The `eofThres` field specifies the duration after which the production studio automatically switches to the standby resource if a stream interruption occurs. Unit: seconds.
	//
	// example:
	//
	// {"eofThres":3}
	AutoSwitchUrgentConfig *string `json:"AutoSwitchUrgentConfig,omitempty" xml:"AutoSwitchUrgentConfig,omitempty"`
	// Specifies whether the production studio automatically switches to the standby resource in case of a stream interruption.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoSwitchUrgentOn *bool `json:"AutoSwitchUrgentOn,omitempty" xml:"AutoSwitchUrgentOn,omitempty"`
	// The callback URL. Enter a valid HTTP address for receiving callback notifications. If you do not specify this parameter, the production studio does not send callback notifications.
	//
	// >  For more information about production studio callbacks, see [Production studio callbacks](https://help.aliyun.com/document_detail/213633.html).
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
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
	// Specifies whether to enable channels. Valid values:
	//
	// 	- **0*	- (default): disables channels.
	//
	// 	- **1**: enables channels.
	//
	// > You cannot disable channels after you enable them. If you set this parameter to 0, the production studio references video resources in a layout without using channels. If you enable channels for the first time, make sure that the production studio is in the idle state. After you enable channels, a new layout that references video resources by using channels is generated to replace the original one. Therefore, you must specify video resources for channels. You can use the channels to change the playback progress or status. If the video resource, preview, and program modules of the production studio use the same video source, the three modules display the same content.
	//
	// example:
	//
	// 1
	ChannelEnable *int32 `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty"`
	// Specifies whether to enable stream delay. Unit: seconds. Valid values:
	//
	// 	- **0*	- (default): disables stream delay.
	//
	// 	- **A value greater than 0**: enables stream delay.
	//
	// 	- **Empty**: clears the stream delay configuration.
	//
	//     **
	//
	//     **Note **The maximum value can be 300 seconds.
	//
	// example:
	//
	// 0
	Delay *float32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The main streaming domain.
	//
	// Complete the configuration of the domain name before the production studio is started. If you do not specify this parameter, the domain configuration for the production studio is cleared.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable the carousel playback feature. Valid values:
	//
	// 	- **0**: disables carousel playback.
	//
	// 	- **1**: enables carousel playback.
	//
	// example:
	//
	// 1
	ProgramEffect *int32 `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty"`
	// The name of the playlist for carousel playback. You can specify this parameter if you enable the carousel playback feature.
	//
	// example:
	//
	// program_name
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The recording configuration. The value is a JSON string. You can configure the following fields:
	//
	// 	- **endpoint**: the API server address of an Alibaba Cloud service.
	//
	// 	- **ossBucket**: the name of the Object Storage Service (OSS) bucket.
	//
	// 	- **videoFormat**: the format in which the video file can be exported. Example: `[{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"m3u8\\",\\"CycleDuration\\":21600,\\"SliceOssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{UnixTimestamp}\\"},{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"flv\\",\\"CycleDuration\\":21600}]`.
	//
	// 	- **interval**: the interval between recordings. Unit: milliseconds.
	//
	// > If you do not specify this parameter, the recording feature is disabled and the recording configuration for the production studio is cleared.
	//
	// example:
	//
	// { "endpoint": "http://oss-cn-********.aliyuncs.com/api",  "ossBucket****": "liveBucket****", "VideoFormat":[{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"m3u8\\",\\"CycleDuration\\":21600,\\"SliceOssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{UnixTimestamp}\\"},{\\"OssObjectPrefix\\":\\"record/{AppName}/{StreamName}/{StartTime}_{EndTime}\\",\\"Format\\":\\"flv\\",\\"CycleDuration\\":21600}] "interval": 5 }
	RecordConfig *string `json:"RecordConfig,omitempty" xml:"RecordConfig,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The custom stream redirect URL.
	//
	// If you do not specify this parameter, the production studio uses the redirect URL generated by the system.
	//
	// > Redirect URLs support only the Real-Time Messaging Protocol (RTMP) protocol.
	SideOutputUrl *string `json:"SideOutputUrl,omitempty" xml:"SideOutputUrl,omitempty"`
	// The stream relay URLs. A relay URL can be an Alibaba Cloud URL or a URL from a third-party CDN provider. You can specify up to 20 relay URLs over the RTMP protocol.
	//
	// > Use the following format to specify multiple relay URLs: "rtmp://domain/app1/stream1","rtmp://domain/app2/stream2".
	//
	// example:
	//
	// rtmp://domain/app/stream?***
	SideOutputUrlList *string `json:"SideOutputUrlList,omitempty" xml:"SideOutputUrlList,omitempty"`
	// The multi-view synchronization configuration. You can specify this parameter to synchronize multiple video sources.
	//
	// There are two modes of multi-view synchronization.
	//
	// 	- A value of 0 for the mode field specifies the streamer mode. In this mode, multiple video sources are synchronized based on the settings by the streamer.
	//
	// 	- A value of 1 for the mode field specifies the conference mode. In this mode, all video sources are synchronized.
	//
	// In the streamer mode, the hostResourceId field specifies the video source on the streamer side.
	//
	// In the conference mode, the hostResourceId field is not available. You need to provide only resource IDs that are required.
	//
	// example:
	//
	// "[{\\"mode\\":0,\\"resourceIds\\":[\\"5a6c1c33-8424-46f6-813c-c152220a****\\",\\"4e6521dc-a40a-4077-b6bf-1fb12a76****\\"],\\"hostResourceId\\":\\"3aa2b39a-fd0e-4b8c-be73-b7af31c4****\\"}]"
	SyncGroupsConfig *string `json:"SyncGroupsConfig,omitempty" xml:"SyncGroupsConfig,omitempty"`
	// The transcoding configuration.
	//
	// The value is a JSON string. Use upper camel case for fields of the string. If you do not specify this parameter, the transcoding configuration is cleared. If no transcoding template is available, an error occurs when the production studio is started.
	//
	// example:
	//
	// {"casterTemplate": "lp_ld"}
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty"`
	// The ID of the standby image from the media library.
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
	// The ID of the standby video from the media library. If you do not specify this parameter, the standby video configuration for the production studio is cleared.
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
