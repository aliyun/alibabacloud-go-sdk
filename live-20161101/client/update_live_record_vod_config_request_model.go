// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordVodConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLiveRecordVodConfigRequest
	GetAppName() *string
	SetAutoCompose(v string) *UpdateLiveRecordVodConfigRequest
	GetAutoCompose() *string
	SetComposeVodTranscodeGroupId(v string) *UpdateLiveRecordVodConfigRequest
	GetComposeVodTranscodeGroupId() *string
	SetCycleDuration(v int32) *UpdateLiveRecordVodConfigRequest
	GetCycleDuration() *int32
	SetDelayTime(v int32) *UpdateLiveRecordVodConfigRequest
	GetDelayTime() *int32
	SetDomainName(v string) *UpdateLiveRecordVodConfigRequest
	GetDomainName() *string
	SetOnDemand(v int32) *UpdateLiveRecordVodConfigRequest
	GetOnDemand() *int32
	SetOwnerId(v int64) *UpdateLiveRecordVodConfigRequest
	GetOwnerId() *int64
	SetRecordFormat(v []*UpdateLiveRecordVodConfigRequestRecordFormat) *UpdateLiveRecordVodConfigRequest
	GetRecordFormat() []*UpdateLiveRecordVodConfigRequestRecordFormat
	SetRegionId(v string) *UpdateLiveRecordVodConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *UpdateLiveRecordVodConfigRequest
	GetStreamName() *string
	SetTranscodeTemplates(v []*string) *UpdateLiveRecordVodConfigRequest
	GetTranscodeTemplates() []*string
	SetVodTranscodeGroupId(v string) *UpdateLiveRecordVodConfigRequest
	GetVodTranscodeGroupId() *string
}

type UpdateLiveRecordVodConfigRequest struct {
	// The application name. You can view the `AppName` on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 	Notice: This parameter is ignored if the `RecordFormat` parameter is specified.
	//
	// Specifies whether to enable automatic composition. Valid values:
	//
	// - **ON**: Enables automatic composition. If you set this value to ON, you must also specify the `ComposeVodTranscodeGroupId` parameter.
	//
	// example:
	//
	// OFF
	AutoCompose *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// 	Notice: This parameter is ignored if the `RecordFormat` parameter is specified.
	//
	// The ID of the ApsaraVideo for VOD transcoding template group used to transcode the video after automatic composition.
	//
	// > You can get the ID by calling the [Query Transcoding Configuration List](https://help.aliyun.com/document_detail/454928.html) operation.
	//
	// example:
	//
	// *****
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	// The duration of each cyclical recording file, in seconds. Default value: **3600**. Valid values: **300*	- to **21600**.
	//
	// example:
	//
	// 300
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The maximum duration of a stream interruption, in seconds. If a stream interruption exceeds this duration, the system generates a new file. Valid values: 15 to 21600.
	//
	// example:
	//
	// 180
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The on-demand recording mode. Valid values:
	//
	// - **0*	- (default): Disables on-demand recording.
	//
	// - **1**: Enables on-demand recording triggered by an HTTP callback.
	//
	// - **2**: Triggers recording by parsing push streaming parameters.
	//
	// - **7**: Manual recording. Call the [RealTimeRecordCommand](https://help.aliyun.com/document_detail/2847882.html) operation to start or stop recording.
	//
	// example:
	//
	// 0
	OnDemand *int32 `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A list of parameters for each recording format.
	RecordFormat []*UpdateLiveRecordVodConfigRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stream name. You can view the `StreamName` on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// A list of transcoding templates.
	TranscodeTemplates []*string `json:"TranscodeTemplates,omitempty" xml:"TranscodeTemplates,omitempty" type:"Repeated"`
	// 	Notice: This parameter is ignored if the `RecordFormat` parameter is specified. The ID of the ApsaraVideo for VOD transcoding template group.
	//
	// example:
	//
	// e2d796d3bb5fd8049d32bff62f94****
	VodTranscodeGroupId *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
}

func (s UpdateLiveRecordVodConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordVodConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordVodConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLiveRecordVodConfigRequest) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *UpdateLiveRecordVodConfigRequest) GetComposeVodTranscodeGroupId() *string {
	return s.ComposeVodTranscodeGroupId
}

func (s *UpdateLiveRecordVodConfigRequest) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *UpdateLiveRecordVodConfigRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *UpdateLiveRecordVodConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveRecordVodConfigRequest) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *UpdateLiveRecordVodConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveRecordVodConfigRequest) GetRecordFormat() []*UpdateLiveRecordVodConfigRequestRecordFormat {
	return s.RecordFormat
}

func (s *UpdateLiveRecordVodConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveRecordVodConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateLiveRecordVodConfigRequest) GetTranscodeTemplates() []*string {
	return s.TranscodeTemplates
}

func (s *UpdateLiveRecordVodConfigRequest) GetVodTranscodeGroupId() *string {
	return s.VodTranscodeGroupId
}

func (s *UpdateLiveRecordVodConfigRequest) SetAppName(v string) *UpdateLiveRecordVodConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetAutoCompose(v string) *UpdateLiveRecordVodConfigRequest {
	s.AutoCompose = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetComposeVodTranscodeGroupId(v string) *UpdateLiveRecordVodConfigRequest {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetCycleDuration(v int32) *UpdateLiveRecordVodConfigRequest {
	s.CycleDuration = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetDelayTime(v int32) *UpdateLiveRecordVodConfigRequest {
	s.DelayTime = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetDomainName(v string) *UpdateLiveRecordVodConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetOnDemand(v int32) *UpdateLiveRecordVodConfigRequest {
	s.OnDemand = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetOwnerId(v int64) *UpdateLiveRecordVodConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetRecordFormat(v []*UpdateLiveRecordVodConfigRequestRecordFormat) *UpdateLiveRecordVodConfigRequest {
	s.RecordFormat = v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetRegionId(v string) *UpdateLiveRecordVodConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetStreamName(v string) *UpdateLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetTranscodeTemplates(v []*string) *UpdateLiveRecordVodConfigRequest {
	s.TranscodeTemplates = v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetVodTranscodeGroupId(v string) *UpdateLiveRecordVodConfigRequest {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) Validate() error {
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

type UpdateLiveRecordVodConfigRequestRecordFormat struct {
	// Specifies whether to enable automatic composition. Valid values:
	//
	// - `ON`: Enables automatic composition.
	//
	// - `OFF`: Disables automatic composition.
	//
	// example:
	//
	// ON
	AutoCompose *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// The recording storage format.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The video processing method. Valid values:
	//
	// - `transcode`: Processes the video by using a transcoding template group.
	//
	// - `workflow`: Processes the video by using a workflow.
	//
	// example:
	//
	// transcode
	ProcessMethod *string `json:"ProcessMethod,omitempty" xml:"ProcessMethod,omitempty"`
	// The ID of the transcoding template group or workflow.
	//
	// > ## The ID must match the video processing method specified in ProcessMethod. For example, if ProcessMethod is set to transcode, you must use a transcoding template group ID.
	//
	// example:
	//
	// e2d796d3bb5fd8049d32bff62f94****
	ProcessTemplateId *string `json:"ProcessTemplateId,omitempty" xml:"ProcessTemplateId,omitempty"`
	// The duration of each segment, in seconds.
	//
	// 	Notice: This parameter applies only to the `m3u8` format.
	//
	// The default value is 30. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// A tag for video classification.
	//
	// example:
	//
	// sports
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The video source to process. Valid values:
	//
	// - `origin` (default): Processes the source video.
	//
	// - `compose`: Processes the composed video.
	//
	// To process both the source and composed videos, separate the values with a comma. For example, `origin,compose`.
	//
	// example:
	//
	// origin
	VideoProcess *string `json:"VideoProcess,omitempty" xml:"VideoProcess,omitempty"`
}

func (s UpdateLiveRecordVodConfigRequestRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordVodConfigRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetProcessMethod() *string {
	return s.ProcessMethod
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetProcessTemplateId() *string {
	return s.ProcessTemplateId
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetTags() *string {
	return s.Tags
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) GetVideoProcess() *string {
	return s.VideoProcess
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetAutoCompose(v string) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.AutoCompose = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetFormat(v string) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetProcessMethod(v string) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.ProcessMethod = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetProcessTemplateId(v string) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.ProcessTemplateId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetSliceDuration(v int32) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetTags(v string) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.Tags = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) SetVideoProcess(v string) *UpdateLiveRecordVodConfigRequestRecordFormat {
	s.VideoProcess = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequestRecordFormat) Validate() error {
	return dara.Validate(s)
}
