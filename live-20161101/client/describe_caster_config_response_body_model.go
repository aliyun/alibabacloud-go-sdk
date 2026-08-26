// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioMixerMode(v string) *DescribeCasterConfigResponseBody
	GetAudioMixerMode() *string
	SetAutoSwitchUrgentConfig(v string) *DescribeCasterConfigResponseBody
	GetAutoSwitchUrgentConfig() *string
	SetAutoSwitchUrgentOn(v string) *DescribeCasterConfigResponseBody
	GetAutoSwitchUrgentOn() *string
	SetCallbackUrl(v string) *DescribeCasterConfigResponseBody
	GetCallbackUrl() *string
	SetCasterId(v string) *DescribeCasterConfigResponseBody
	GetCasterId() *string
	SetCasterName(v string) *DescribeCasterConfigResponseBody
	GetCasterName() *string
	SetChannelEnable(v int32) *DescribeCasterConfigResponseBody
	GetChannelEnable() *int32
	SetDelay(v float32) *DescribeCasterConfigResponseBody
	GetDelay() *float32
	SetDomainName(v string) *DescribeCasterConfigResponseBody
	GetDomainName() *string
	SetProgramEffect(v int32) *DescribeCasterConfigResponseBody
	GetProgramEffect() *int32
	SetProgramName(v string) *DescribeCasterConfigResponseBody
	GetProgramName() *string
	SetRecordConfig(v *DescribeCasterConfigResponseBodyRecordConfig) *DescribeCasterConfigResponseBody
	GetRecordConfig() *DescribeCasterConfigResponseBodyRecordConfig
	SetRequestId(v string) *DescribeCasterConfigResponseBody
	GetRequestId() *string
	SetSideOutputUrl(v string) *DescribeCasterConfigResponseBody
	GetSideOutputUrl() *string
	SetSideOutputUrlList(v string) *DescribeCasterConfigResponseBody
	GetSideOutputUrlList() *string
	SetSyncGroupsConfig(v *DescribeCasterConfigResponseBodySyncGroupsConfig) *DescribeCasterConfigResponseBody
	GetSyncGroupsConfig() *DescribeCasterConfigResponseBodySyncGroupsConfig
	SetTranscodeConfig(v *DescribeCasterConfigResponseBodyTranscodeConfig) *DescribeCasterConfigResponseBody
	GetTranscodeConfig() *DescribeCasterConfigResponseBodyTranscodeConfig
	SetUrgentImageId(v string) *DescribeCasterConfigResponseBody
	GetUrgentImageId() *string
	SetUrgentImageUrl(v string) *DescribeCasterConfigResponseBody
	GetUrgentImageUrl() *string
	SetUrgentLiveStreamUrl(v string) *DescribeCasterConfigResponseBody
	GetUrgentLiveStreamUrl() *string
	SetUrgentMaterialId(v string) *DescribeCasterConfigResponseBody
	GetUrgentMaterialId() *string
}

type DescribeCasterConfigResponseBody struct {
	AudioMixerMode *string `json:"AudioMixerMode,omitempty" xml:"AudioMixerMode,omitempty"`
	// The automatic standby video switching configuration. `eofThres`: specifies the duration of stream interruption before automatically switching to the standby video. Unit: seconds.
	//
	// example:
	//
	// {"eofThres":3}
	AutoSwitchUrgentConfig *string `json:"AutoSwitchUrgentConfig,omitempty" xml:"AutoSwitchUrgentConfig,omitempty"`
	// Indicates whether automatic switchover to the standby video upon stream interruption is enabled.
	//
	// - **true**: Enabled.
	//
	// - **false**: Shutdown.
	//
	// example:
	//
	// true
	AutoSwitchUrgentOn *string `json:"AutoSwitchUrgentOn,omitempty" xml:"AutoSwitchUrgentOn,omitempty"`
	// The user callback URL.
	//
	// example:
	//
	// http://learn.aliyundoc.com/callBackLive
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The production studio ID.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The production studio name.
	//
	// example:
	//
	// coco-caster10
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	// Indicates whether Channel is enabled. Valid values:
	//
	//
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	ChannelEnable *int32 `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty"`
	// The stream delay. Unit: seconds.
	//
	//
	//
	// - **0**: Stream delay is disabled.
	//
	// - Greater than **0**: Stream delay is enabled.
	//
	// example:
	//
	// 0
	Delay *float32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The primary streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The playlist effective flag. Valid values:
	//
	//
	//
	// - **0**: Not effective.
	//
	// - **1**: Effective.
	//
	// example:
	//
	// 0
	ProgramEffect *int32 `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty"`
	// The playlist name.
	//
	// example:
	//
	// program_name
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The recording configuration. If this parameter is empty, the recording feature is not enabled.
	RecordConfig *DescribeCasterConfigResponseBodyRecordConfig `json:"RecordConfig,omitempty" xml:"RecordConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 97df6b7f-3490-47d2-ac50-8833e1b64597
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The custom side output URL of the production studio.
	//
	// example:
	//
	// rtmp://sophon-developer.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****
	SideOutputUrl *string `json:"SideOutputUrl,omitempty" xml:"SideOutputUrl,omitempty"`
	// The list of custom side output URLs of the production studio.
	//
	// example:
	//
	// ["rtmp://domain/app1/stream1","rtmp://domain/app2/stream2"]
	SideOutputUrlList *string                                           `json:"SideOutputUrlList,omitempty" xml:"SideOutputUrlList,omitempty"`
	SyncGroupsConfig  *DescribeCasterConfigResponseBodySyncGroupsConfig `json:"SyncGroupsConfig,omitempty" xml:"SyncGroupsConfig,omitempty" type:"Struct"`
	// The transcoding configuration.
	TranscodeConfig *DescribeCasterConfigResponseBodyTranscodeConfig `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty" type:"Struct"`
	// The media library asset ID of the standby image.
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
	// The media library asset ID of the standby video.
	//
	// example:
	//
	// 98646538-bcf9-4aef-bd4a-e6bb76588****
	UrgentMaterialId *string `json:"UrgentMaterialId,omitempty" xml:"UrgentMaterialId,omitempty"`
}

func (s DescribeCasterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBody) GetAudioMixerMode() *string {
	return s.AudioMixerMode
}

func (s *DescribeCasterConfigResponseBody) GetAutoSwitchUrgentConfig() *string {
	return s.AutoSwitchUrgentConfig
}

func (s *DescribeCasterConfigResponseBody) GetAutoSwitchUrgentOn() *string {
	return s.AutoSwitchUrgentOn
}

func (s *DescribeCasterConfigResponseBody) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *DescribeCasterConfigResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterConfigResponseBody) GetCasterName() *string {
	return s.CasterName
}

func (s *DescribeCasterConfigResponseBody) GetChannelEnable() *int32 {
	return s.ChannelEnable
}

func (s *DescribeCasterConfigResponseBody) GetDelay() *float32 {
	return s.Delay
}

func (s *DescribeCasterConfigResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCasterConfigResponseBody) GetProgramEffect() *int32 {
	return s.ProgramEffect
}

func (s *DescribeCasterConfigResponseBody) GetProgramName() *string {
	return s.ProgramName
}

func (s *DescribeCasterConfigResponseBody) GetRecordConfig() *DescribeCasterConfigResponseBodyRecordConfig {
	return s.RecordConfig
}

func (s *DescribeCasterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterConfigResponseBody) GetSideOutputUrl() *string {
	return s.SideOutputUrl
}

func (s *DescribeCasterConfigResponseBody) GetSideOutputUrlList() *string {
	return s.SideOutputUrlList
}

func (s *DescribeCasterConfigResponseBody) GetSyncGroupsConfig() *DescribeCasterConfigResponseBodySyncGroupsConfig {
	return s.SyncGroupsConfig
}

func (s *DescribeCasterConfigResponseBody) GetTranscodeConfig() *DescribeCasterConfigResponseBodyTranscodeConfig {
	return s.TranscodeConfig
}

func (s *DescribeCasterConfigResponseBody) GetUrgentImageId() *string {
	return s.UrgentImageId
}

func (s *DescribeCasterConfigResponseBody) GetUrgentImageUrl() *string {
	return s.UrgentImageUrl
}

func (s *DescribeCasterConfigResponseBody) GetUrgentLiveStreamUrl() *string {
	return s.UrgentLiveStreamUrl
}

func (s *DescribeCasterConfigResponseBody) GetUrgentMaterialId() *string {
	return s.UrgentMaterialId
}

func (s *DescribeCasterConfigResponseBody) SetAudioMixerMode(v string) *DescribeCasterConfigResponseBody {
	s.AudioMixerMode = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetAutoSwitchUrgentConfig(v string) *DescribeCasterConfigResponseBody {
	s.AutoSwitchUrgentConfig = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetAutoSwitchUrgentOn(v string) *DescribeCasterConfigResponseBody {
	s.AutoSwitchUrgentOn = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetCallbackUrl(v string) *DescribeCasterConfigResponseBody {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetCasterId(v string) *DescribeCasterConfigResponseBody {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetCasterName(v string) *DescribeCasterConfigResponseBody {
	s.CasterName = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetChannelEnable(v int32) *DescribeCasterConfigResponseBody {
	s.ChannelEnable = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetDelay(v float32) *DescribeCasterConfigResponseBody {
	s.Delay = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetDomainName(v string) *DescribeCasterConfigResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetProgramEffect(v int32) *DescribeCasterConfigResponseBody {
	s.ProgramEffect = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetProgramName(v string) *DescribeCasterConfigResponseBody {
	s.ProgramName = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetRecordConfig(v *DescribeCasterConfigResponseBodyRecordConfig) *DescribeCasterConfigResponseBody {
	s.RecordConfig = v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetRequestId(v string) *DescribeCasterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetSideOutputUrl(v string) *DescribeCasterConfigResponseBody {
	s.SideOutputUrl = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetSideOutputUrlList(v string) *DescribeCasterConfigResponseBody {
	s.SideOutputUrlList = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetSyncGroupsConfig(v *DescribeCasterConfigResponseBodySyncGroupsConfig) *DescribeCasterConfigResponseBody {
	s.SyncGroupsConfig = v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetTranscodeConfig(v *DescribeCasterConfigResponseBodyTranscodeConfig) *DescribeCasterConfigResponseBody {
	s.TranscodeConfig = v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetUrgentImageId(v string) *DescribeCasterConfigResponseBody {
	s.UrgentImageId = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetUrgentImageUrl(v string) *DescribeCasterConfigResponseBody {
	s.UrgentImageUrl = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetUrgentLiveStreamUrl(v string) *DescribeCasterConfigResponseBody {
	s.UrgentLiveStreamUrl = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) SetUrgentMaterialId(v string) *DescribeCasterConfigResponseBody {
	s.UrgentMaterialId = &v
	return s
}

func (s *DescribeCasterConfigResponseBody) Validate() error {
	if s.RecordConfig != nil {
		if err := s.RecordConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SyncGroupsConfig != nil {
		if err := s.SyncGroupsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeConfig != nil {
		if err := s.TranscodeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodyRecordConfig struct {
	// The on-demand recording mode. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: HTTP callback-based.
	//
	// - 2: On-demand recording by parsing stream ingest parameters.
	//
	// - 7: Not recorded by default.
	//
	// example:
	//
	// 0
	OnDemand *int32 `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	// The storage location.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The OSS endpoint of the storage location.
	//
	// example:
	//
	// oss-cn-shanghai.aliyundoc.com
	OssEndpoint  *string                                                   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	RecordFormat *DescribeCasterConfigResponseBodyRecordConfigRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Struct"`
}

func (s DescribeCasterConfigResponseBodyRecordConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyRecordConfig) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) GetRecordFormat() *DescribeCasterConfigResponseBodyRecordConfigRecordFormat {
	return s.RecordFormat
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) SetOnDemand(v int32) *DescribeCasterConfigResponseBodyRecordConfig {
	s.OnDemand = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) SetOssBucket(v string) *DescribeCasterConfigResponseBodyRecordConfig {
	s.OssBucket = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) SetOssEndpoint(v string) *DescribeCasterConfigResponseBodyRecordConfig {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) SetRecordFormat(v *DescribeCasterConfigResponseBodyRecordConfigRecordFormat) *DescribeCasterConfigResponseBodyRecordConfig {
	s.RecordFormat = v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfig) Validate() error {
	if s.RecordFormat != nil {
		if err := s.RecordFormat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodyRecordConfigRecordFormat struct {
	RecordFormat []*DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
}

func (s DescribeCasterConfigResponseBodyRecordConfigRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyRecordConfigRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormat) GetRecordFormat() []*DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat {
	return s.RecordFormat
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormat) SetRecordFormat(v []*DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) *DescribeCasterConfigResponseBodyRecordConfigRecordFormat {
	s.RecordFormat = v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormat) Validate() error {
	if s.RecordFormat != nil {
		for _, item := range s.RecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat struct {
	CycleDuration        *int32  `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty"`
	OssObjectPrefix      *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) SetCycleDuration(v int32) *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) SetFormat(v string) *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat {
	s.Format = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) SetOssObjectPrefix(v string) *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) SetSliceOssObjectPrefix(v string) *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyRecordConfigRecordFormatRecordFormat) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterConfigResponseBodySyncGroupsConfig struct {
	SyncGroup []*DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup `json:"SyncGroup,omitempty" xml:"SyncGroup,omitempty" type:"Repeated"`
}

func (s DescribeCasterConfigResponseBodySyncGroupsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodySyncGroupsConfig) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfig) GetSyncGroup() []*DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup {
	return s.SyncGroup
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfig) SetSyncGroup(v []*DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) *DescribeCasterConfigResponseBodySyncGroupsConfig {
	s.SyncGroup = v
	return s
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfig) Validate() error {
	if s.SyncGroup != nil {
		for _, item := range s.SyncGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup struct {
	HostResourceId *string                                                               `json:"HostResourceId,omitempty" xml:"HostResourceId,omitempty"`
	Mode           *int32                                                                `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ResourceIds    *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Struct"`
}

func (s DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) GetHostResourceId() *string {
	return s.HostResourceId
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) GetMode() *int32 {
	return s.Mode
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) GetResourceIds() *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds {
	return s.ResourceIds
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) SetHostResourceId(v string) *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup {
	s.HostResourceId = &v
	return s
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) SetMode(v int32) *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup {
	s.Mode = &v
	return s
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) SetResourceIds(v *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds) *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup {
	s.ResourceIds = v
	return s
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroup) Validate() error {
	if s.ResourceIds != nil {
		if err := s.ResourceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds struct {
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds) SetResourceId(v []*string) *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds {
	s.ResourceId = v
	return s
}

func (s *DescribeCasterConfigResponseBodySyncGroupsConfigSyncGroupResourceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterConfigResponseBodyTranscodeConfig struct {
	// The production studio transcoding template. Valid values:
	//
	// - **lp_ld**: low definition.
	//
	// - **lp_sd**: standard definition.
	//
	// - **lp_hd**: high definition.
	//
	// - **lp_ud**: ultra-high definition.
	//
	// - **lp_ld_v**: portrait low definition.
	//
	// - **lp_sd_v**: portrait standard definition.
	//
	// - **lp_hd_v**: portrait high definition.
	//
	// - **lp_ud_v**: portrait ultra-high definition.
	//
	// example:
	//
	// lp_hd
	CasterTemplate *string `json:"CasterTemplate,omitempty" xml:"CasterTemplate,omitempty"`
	// The custom configuration.
	CustomParams    *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams    `json:"CustomParams,omitempty" xml:"CustomParams,omitempty" type:"Struct"`
	LiveTemplateIds *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds `json:"LiveTemplateIds,omitempty" xml:"LiveTemplateIds,omitempty" type:"Struct"`
}

func (s DescribeCasterConfigResponseBodyTranscodeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyTranscodeConfig) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) GetCasterTemplate() *string {
	return s.CasterTemplate
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) GetCustomParams() *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams {
	return s.CustomParams
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) GetLiveTemplateIds() *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds {
	return s.LiveTemplateIds
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) SetCasterTemplate(v string) *DescribeCasterConfigResponseBodyTranscodeConfig {
	s.CasterTemplate = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) SetCustomParams(v *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams) *DescribeCasterConfigResponseBodyTranscodeConfig {
	s.CustomParams = v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) SetLiveTemplateIds(v *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds) *DescribeCasterConfigResponseBodyTranscodeConfig {
	s.LiveTemplateIds = v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfig) Validate() error {
	if s.CustomParams != nil {
		if err := s.CustomParams.Validate(); err != nil {
			return err
		}
	}
	if s.LiveTemplateIds != nil {
		if err := s.LiveTemplateIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodyTranscodeConfigCustomParams struct {
	// The video parameters.
	Video *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo `json:"video,omitempty" xml:"video,omitempty" type:"Struct"`
}

func (s DescribeCasterConfigResponseBodyTranscodeConfigCustomParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyTranscodeConfigCustomParams) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams) GetVideo() *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo {
	return s.Video
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams) SetVideo(v *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams {
	s.Video = v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParams) Validate() error {
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo struct {
	// The video bitrate.
	//
	// example:
	//
	// 300
	Bitrate *int32 `json:"bitrate,omitempty" xml:"bitrate,omitempty"`
	// The video frame rate.
	//
	// example:
	//
	// 300
	Fps *int32 `json:"fps,omitempty" xml:"fps,omitempty"`
	// The video height. Unit: pixels (px).
	//
	// example:
	//
	// 720
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// The video width. Unit: pixels (px).
	//
	// example:
	//
	// 1080
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) GetFps() *int32 {
	return s.Fps
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) SetBitrate(v int32) *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) SetFps(v int32) *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo {
	s.Fps = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) SetHeight(v int32) *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo {
	s.Height = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) SetWidth(v int32) *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo {
	s.Width = &v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigCustomParamsVideo) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds struct {
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" type:"Repeated"`
}

func (s DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds) GetLocationId() []*string {
	return s.LocationId
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds) SetLocationId(v []*string) *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds {
	s.LocationId = v
	return s
}

func (s *DescribeCasterConfigResponseBodyTranscodeConfigLiveTemplateIds) Validate() error {
	return dara.Validate(s)
}
