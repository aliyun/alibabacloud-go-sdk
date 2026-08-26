// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordVodConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveRecordVodConfigRequest
	GetAppName() *string
	SetAutoCompose(v string) *AddLiveRecordVodConfigRequest
	GetAutoCompose() *string
	SetComposeVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest
	GetComposeVodTranscodeGroupId() *string
	SetCycleDuration(v int32) *AddLiveRecordVodConfigRequest
	GetCycleDuration() *int32
	SetDelayTime(v int32) *AddLiveRecordVodConfigRequest
	GetDelayTime() *int32
	SetDomainName(v string) *AddLiveRecordVodConfigRequest
	GetDomainName() *string
	SetOnDemand(v int32) *AddLiveRecordVodConfigRequest
	GetOnDemand() *int32
	SetOwnerId(v int64) *AddLiveRecordVodConfigRequest
	GetOwnerId() *int64
	SetRecordContent(v string) *AddLiveRecordVodConfigRequest
	GetRecordContent() *string
	SetRecordFormat(v []*AddLiveRecordVodConfigRequestRecordFormat) *AddLiveRecordVodConfigRequest
	GetRecordFormat() []*AddLiveRecordVodConfigRequestRecordFormat
	SetRegionId(v string) *AddLiveRecordVodConfigRequest
	GetRegionId() *string
	SetSpaceId(v string) *AddLiveRecordVodConfigRequest
	GetSpaceId() *string
	SetStorageLocation(v string) *AddLiveRecordVodConfigRequest
	GetStorageLocation() *string
	SetStreamName(v string) *AddLiveRecordVodConfigRequest
	GetStreamName() *string
	SetTranscodeTemplates(v []*string) *AddLiveRecordVodConfigRequest
	GetTranscodeTemplates() []*string
	SetVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest
	GetVodTranscodeGroupId() *string
}

type AddLiveRecordVodConfigRequest struct {
	// The name of the application that the stream belongs to. You can find this value on the [stream management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 	Notice: This parameter is ignored if `RecordFormat` is specified.
	//
	// Specifies whether to automatically merge files from multiple recording cycles into a single file after a live stream ends.
	//
	// A value of **ON*	- enables automatic merging. If enabled, you must also specify the `ComposeVodTranscodeGroupId` parameter. By default, automatic merging is disabled.
	//
	// example:
	//
	// ON
	AutoCompose *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// 	Notice: This parameter is ignored if `RecordFormat` is specified.
	//
	// The ID of the ApsaraVideo VOD transcoding template group for transcoding the merged video. This parameter is required if `AutoCompose` is set to `ON`.
	//
	// example:
	//
	// *****
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	// The cycle duration, in seconds. The default value is **3600**. The value must be between **300*	- and **21600**.
	//
	// example:
	//
	// 300
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The stream interruption timeout, in seconds. If a stream interruption is shorter than this duration, recording continues in the same file. If the interruption is longer, a new file is created. Valid values: 15 to 21600.
	//
	// example:
	//
	// 180
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The streaming domain.
	//
	// > Ensure ApsaraVideo VOD is activated in the same region as the streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The recording trigger mode. Valid values:
	//
	// - **0*	- (Default): Automatic recording.
	//
	// - **1**: On-demand recording triggered by an HTTP callback.
	//
	// - **2**: On-demand recording triggered by ingest parameters.
	//
	// - **7**: Manual recording. Allows you to start and stop recording by calling the `RealTimeRecordCommand` operation.
	//
	// example:
	//
	// 0
	OnDemand *int32 `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recorded content. Valid values:
	//
	// - `raw` (Default): Records the source stream.
	//
	// - `transcode`: Records transcoded streams.
	//
	// To record both source and transcoded streams, provide a comma-separated list, for example, `raw,transcode`.
	//
	// > If this parameter is set to include `transcode`, you must specify at least one template in the `TranscodeTemplates` parameter.
	//
	// example:
	//
	// raw
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// A list of format-specific recording configurations.
	RecordFormat []*AddLiveRecordVodConfigRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
	// The region ID. The example value `cn-shanghai` indicates the China (Shanghai) region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VOD application space. You can obtain this ID from the **VOD console*	- or by calling an [API operation to query application information](https://help.aliyun.com/document_detail/454873.html). This parameter applies only when the VOD application space feature is enabled.
	//
	// example:
	//
	// app-1000000
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The storage location.
	//
	// example:
	//
	// ****-tjptr2vatm.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The stream name. You can find this value on the [stream management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// A list of transcoding templates for recording transcoded streams.
	TranscodeTemplates []*string `json:"TranscodeTemplates,omitempty" xml:"TranscodeTemplates,omitempty" type:"Repeated"`
	// 	Notice: This parameter is ignored if `RecordFormat` is specified.
	//
	// The ID of the ApsaraVideo VOD transcoding template group for transcoding recorded videos.
	//
	// example:
	//
	// e2d796d3bb5fd8049d32bff62f94****
	VodTranscodeGroupId *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
}

func (s AddLiveRecordVodConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordVodConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveRecordVodConfigRequest) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *AddLiveRecordVodConfigRequest) GetComposeVodTranscodeGroupId() *string {
	return s.ComposeVodTranscodeGroupId
}

func (s *AddLiveRecordVodConfigRequest) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *AddLiveRecordVodConfigRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *AddLiveRecordVodConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveRecordVodConfigRequest) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *AddLiveRecordVodConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveRecordVodConfigRequest) GetRecordContent() *string {
	return s.RecordContent
}

func (s *AddLiveRecordVodConfigRequest) GetRecordFormat() []*AddLiveRecordVodConfigRequestRecordFormat {
	return s.RecordFormat
}

func (s *AddLiveRecordVodConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveRecordVodConfigRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *AddLiveRecordVodConfigRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AddLiveRecordVodConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLiveRecordVodConfigRequest) GetTranscodeTemplates() []*string {
	return s.TranscodeTemplates
}

func (s *AddLiveRecordVodConfigRequest) GetVodTranscodeGroupId() *string {
	return s.VodTranscodeGroupId
}

func (s *AddLiveRecordVodConfigRequest) SetAppName(v string) *AddLiveRecordVodConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetAutoCompose(v string) *AddLiveRecordVodConfigRequest {
	s.AutoCompose = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetComposeVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetCycleDuration(v int32) *AddLiveRecordVodConfigRequest {
	s.CycleDuration = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetDelayTime(v int32) *AddLiveRecordVodConfigRequest {
	s.DelayTime = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetDomainName(v string) *AddLiveRecordVodConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetOnDemand(v int32) *AddLiveRecordVodConfigRequest {
	s.OnDemand = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetOwnerId(v int64) *AddLiveRecordVodConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetRecordContent(v string) *AddLiveRecordVodConfigRequest {
	s.RecordContent = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetRecordFormat(v []*AddLiveRecordVodConfigRequestRecordFormat) *AddLiveRecordVodConfigRequest {
	s.RecordFormat = v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetRegionId(v string) *AddLiveRecordVodConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetSpaceId(v string) *AddLiveRecordVodConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetStorageLocation(v string) *AddLiveRecordVodConfigRequest {
	s.StorageLocation = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetStreamName(v string) *AddLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetTranscodeTemplates(v []*string) *AddLiveRecordVodConfigRequest {
	s.TranscodeTemplates = v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) Validate() error {
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

type AddLiveRecordVodConfigRequestRecordFormat struct {
	// Specifies whether to automatically merge recording files for this format after the stream ends. Valid values:
	//
	// - `ON`: Enables automatic merging.
	//
	// - `OFF`: Disables automatic merging.
	//
	// example:
	//
	// ON
	AutoCompose *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// The recording format. Valid values:
	//
	// - `m3u8`
	//
	// - `flv`
	//
	// - `mp4`
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The video processing method. Valid values:
	//
	// - `transcode`: Uses a transcoding template group to process the video.
	//
	// - `workflow`: Uses a workflow to process the video.
	//
	// example:
	//
	// transcode
	ProcessMethod *string `json:"ProcessMethod,omitempty" xml:"ProcessMethod,omitempty"`
	// The ID of the transcoding template group or workflow.
	//
	// > The specified ID must match the `ProcessMethod`. For example, provide a transcoding template group ID if `ProcessMethod` is `transcode`, or a workflow ID if `ProcessMethod` is `workflow`.
	//
	// example:
	//
	// e2d796d3bb5fd8049d32bff62f94****
	ProcessTemplateId *string `json:"ProcessTemplateId,omitempty" xml:"ProcessTemplateId,omitempty"`
	// The slice duration, in seconds.
	//
	// This parameter applies only to the `m3u8` format.
	//
	// The value must be between 5 and 30. The default is 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The tags for video categorization.
	//
	// example:
	//
	// sports
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The video source to process. Valid values:
	//
	// - `origin` (Default): The per-cycle recording files.
	//
	// - `compose`: The single video file composed from all cycles.
	//
	// To process both video sources, separate the values with a comma (,), for example, `origin,compose`.
	//
	// example:
	//
	// origin
	VideoProcess *string `json:"VideoProcess,omitempty" xml:"VideoProcess,omitempty"`
}

func (s AddLiveRecordVodConfigRequestRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordVodConfigRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetProcessMethod() *string {
	return s.ProcessMethod
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetProcessTemplateId() *string {
	return s.ProcessTemplateId
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetTags() *string {
	return s.Tags
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) GetVideoProcess() *string {
	return s.VideoProcess
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetAutoCompose(v string) *AddLiveRecordVodConfigRequestRecordFormat {
	s.AutoCompose = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetFormat(v string) *AddLiveRecordVodConfigRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetProcessMethod(v string) *AddLiveRecordVodConfigRequestRecordFormat {
	s.ProcessMethod = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetProcessTemplateId(v string) *AddLiveRecordVodConfigRequestRecordFormat {
	s.ProcessTemplateId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetSliceDuration(v int32) *AddLiveRecordVodConfigRequestRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetTags(v string) *AddLiveRecordVodConfigRequestRecordFormat {
	s.Tags = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) SetVideoProcess(v string) *AddLiveRecordVodConfigRequestRecordFormat {
	s.VideoProcess = &v
	return s
}

func (s *AddLiveRecordVodConfigRequestRecordFormat) Validate() error {
	return dara.Validate(s)
}
