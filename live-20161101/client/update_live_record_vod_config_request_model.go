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
	SetDomainName(v string) *UpdateLiveRecordVodConfigRequest
	GetDomainName() *string
	SetOnDemand(v int32) *UpdateLiveRecordVodConfigRequest
	GetOnDemand() *int32
	SetOwnerId(v int64) *UpdateLiveRecordVodConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveRecordVodConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *UpdateLiveRecordVodConfigRequest
	GetStreamName() *string
	SetVodTranscodeGroupId(v string) *UpdateLiveRecordVodConfigRequest
	GetVodTranscodeGroupId() *string
}

type UpdateLiveRecordVodConfigRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to enable automatic merging. Valid values:
	//
	// 	- **ON**: enables automatic merging. If you set this parameter to ON, you must also specify the ComposeVodTranscodeGroupId parameter.
	//
	// 	- **OFF**: disables automatic merging.
	//
	// example:
	//
	// OFF
	AutoCompose *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// The ID of the transcoding template group in ApsaraVideo VOD that is used to transcode the video file. The video file is generated by merging the VOD files created from live streams.
	//
	// >  To query transcoding template groups, call the [ListTranscodeTemplateGroup](https://help.aliyun.com/document_detail/454928.html) operation.
	//
	// example:
	//
	// *****
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	// The recording cycle. Unit: seconds. Valid values: **300 to 21600**. Default value: **3600**.
	//
	// example:
	//
	// 300
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to enable on-demand recording. Valid values:
	//
	// 	- **0*	- (default): disables on-demand recording.
	//
	// 	- **1**: enables on-demand recording by using the HTTP callback method.
	//
	// example:
	//
	// 0
	OnDemand *int32  `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The ID of the transcoding template group in ApsaraVideo VOD.
	//
	// This parameter is required.
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

func (s *UpdateLiveRecordVodConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveRecordVodConfigRequest) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *UpdateLiveRecordVodConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveRecordVodConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveRecordVodConfigRequest) GetStreamName() *string {
	return s.StreamName
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

func (s *UpdateLiveRecordVodConfigRequest) SetRegionId(v string) *UpdateLiveRecordVodConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetStreamName(v string) *UpdateLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) SetVodTranscodeGroupId(v string) *UpdateLiveRecordVodConfigRequest {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigRequest) Validate() error {
	return dara.Validate(s)
}
