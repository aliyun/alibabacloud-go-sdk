// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAppRecordConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLiveAppRecordConfigRequest
	GetAppName() *string
	SetDelayTime(v int32) *UpdateLiveAppRecordConfigRequest
	GetDelayTime() *int32
	SetDomainName(v string) *UpdateLiveAppRecordConfigRequest
	GetDomainName() *string
	SetEndTime(v string) *UpdateLiveAppRecordConfigRequest
	GetEndTime() *string
	SetOnDemand(v int32) *UpdateLiveAppRecordConfigRequest
	GetOnDemand() *int32
	SetOssEndpoint(v string) *UpdateLiveAppRecordConfigRequest
	GetOssEndpoint() *string
	SetOwnerId(v int64) *UpdateLiveAppRecordConfigRequest
	GetOwnerId() *int64
	SetRecordFormat(v []*UpdateLiveAppRecordConfigRequestRecordFormat) *UpdateLiveAppRecordConfigRequest
	GetRecordFormat() []*UpdateLiveAppRecordConfigRequestRecordFormat
	SetSecurityToken(v string) *UpdateLiveAppRecordConfigRequest
	GetSecurityToken() *string
	SetStartTime(v string) *UpdateLiveAppRecordConfigRequest
	GetStartTime() *string
	SetStreamName(v string) *UpdateLiveAppRecordConfigRequest
	GetStreamName() *string
	SetTranscodeRecordFormat(v []*UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) *UpdateLiveAppRecordConfigRequest
	GetTranscodeRecordFormat() []*UpdateLiveAppRecordConfigRequestTranscodeRecordFormat
	SetTranscodeTemplates(v []*string) *UpdateLiveAppRecordConfigRequest
	GetTranscodeTemplates() []*string
}

type UpdateLiveAppRecordConfigRequest struct {
	// The AppName of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The window in seconds for merging fragmented recording after an interruption. If a stream disconnects and reconnects within this window, the recording will continue in the same file. Valid values: 15 to 21600.
	//
	// example:
	//
	// 180
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The recording end time. Format: *yyyy-MM-dd*T*HH:mm:ss*Z (UTC time).
	//
	// > This parameter is only effective for stream-level recordings. The interval between EndTime and StartTime cannot exceed 7 days.
	//
	// example:
	//
	// 2018-04-16T09:57:21Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies the recording mode. Valid values:
	//
	// - **0**: disables on-demand recording.
	//
	// - **1**: On-demand recording via HTTP callback.
	//
	// - **2**: On-demand recording by parsing parameters in the ingest URL.
	//
	// - **7**: Manual recording. You can call the [RealTimeRecordCommand](https://help.aliyun.com/document_detail/2847882.html) API to manually start or stop recording.
	//
	// > If you set OnDemand to **1**, you need to call the [AddLiveRecordNotifyConfig](https://help.aliyun.com/document_detail/2847891.html) API to configure the OnDemandUrl parameter. Otherwise, ApsaraVideo Live does not perform on-demand recording.
	//
	// example:
	//
	// 1
	OnDemand *int32 `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	// The endpoint for OSS storage. You must create an OSS bucket before using this feature. See [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// learn.developer.aliyundoc.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recording details.
	RecordFormat  []*UpdateLiveAppRecordConfigRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
	SecurityToken *string                                         `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The recording start time. Format: *yyyy-MM-dd*T*HH:mm:ss*Z (UTC time).
	//
	// > This parameter is only effective for stream-level recordings (i.e., when `StreamName` is specified). The time must be within 7 days of the actual stream start time.
	//
	// example:
	//
	// 2018-04-10T09:57:21Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The transcoded stream recording configuration.
	TranscodeRecordFormat []*UpdateLiveAppRecordConfigRequestTranscodeRecordFormat `json:"TranscodeRecordFormat,omitempty" xml:"TranscodeRecordFormat,omitempty" type:"Repeated"`
	// The transcoding template group details.
	TranscodeTemplates []*string `json:"TranscodeTemplates,omitempty" xml:"TranscodeTemplates,omitempty" type:"Repeated"`
}

func (s UpdateLiveAppRecordConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppRecordConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLiveAppRecordConfigRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *UpdateLiveAppRecordConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveAppRecordConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateLiveAppRecordConfigRequest) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *UpdateLiveAppRecordConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *UpdateLiveAppRecordConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAppRecordConfigRequest) GetRecordFormat() []*UpdateLiveAppRecordConfigRequestRecordFormat {
	return s.RecordFormat
}

func (s *UpdateLiveAppRecordConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLiveAppRecordConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateLiveAppRecordConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateLiveAppRecordConfigRequest) GetTranscodeRecordFormat() []*UpdateLiveAppRecordConfigRequestTranscodeRecordFormat {
	return s.TranscodeRecordFormat
}

func (s *UpdateLiveAppRecordConfigRequest) GetTranscodeTemplates() []*string {
	return s.TranscodeTemplates
}

func (s *UpdateLiveAppRecordConfigRequest) SetAppName(v string) *UpdateLiveAppRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetDelayTime(v int32) *UpdateLiveAppRecordConfigRequest {
	s.DelayTime = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetDomainName(v string) *UpdateLiveAppRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetEndTime(v string) *UpdateLiveAppRecordConfigRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetOnDemand(v int32) *UpdateLiveAppRecordConfigRequest {
	s.OnDemand = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetOssEndpoint(v string) *UpdateLiveAppRecordConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetOwnerId(v int64) *UpdateLiveAppRecordConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetRecordFormat(v []*UpdateLiveAppRecordConfigRequestRecordFormat) *UpdateLiveAppRecordConfigRequest {
	s.RecordFormat = v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetSecurityToken(v string) *UpdateLiveAppRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetStartTime(v string) *UpdateLiveAppRecordConfigRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetStreamName(v string) *UpdateLiveAppRecordConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetTranscodeRecordFormat(v []*UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) *UpdateLiveAppRecordConfigRequest {
	s.TranscodeRecordFormat = v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) SetTranscodeTemplates(v []*string) *UpdateLiveAppRecordConfigRequest {
	s.TranscodeTemplates = v
	return s
}

func (s *UpdateLiveAppRecordConfigRequest) Validate() error {
	if s.RecordFormat != nil {
		for _, item := range s.RecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TranscodeRecordFormat != nil {
		for _, item := range s.TranscodeRecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLiveAppRecordConfigRequestRecordFormat struct {
	// The duration of a single recording cycle in seconds. If not specified, the default value is 6 hours
	//
	// > If a live stream is interrupted during a recording cycle but resumes normal streaming within the merge window, recording will continue in the same file. A recording file is generated only when a live stream is interrupted for longer than the merge window.
	//
	// example:
	//
	// 1
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The recording format. Valid values:
	//
	// 	Notice:
	//
	// If you choose m3u8 or cmaf, you must also set SliceOssObjectPrefix and SliceDuration. At least one of RecordFormat or TranscodeRecordFormat must be specified.
	//
	//
	//
	// - m3u8
	//
	// - flv
	//
	// - mp4
	//
	// - cmaf
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The duration of a single segment. Unit: seconds
	//
	// > This parameter takes effect only if you set the RecordFormat.N.Format parameter to m3u8 or cmaf.
	//
	// If you do not specify this parameter, the default value 30 seconds is used. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
}

func (s UpdateLiveAppRecordConfigRequestRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppRecordConfigRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) SetCycleDuration(v int32) *UpdateLiveAppRecordConfigRequestRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) SetFormat(v string) *UpdateLiveAppRecordConfigRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) SetSliceDuration(v int32) *UpdateLiveAppRecordConfigRequestRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequestRecordFormat) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveAppRecordConfigRequestTranscodeRecordFormat struct {
	// The transcoded stream recording cycle. Unit: seconds. If you do not specify this parameter, the default value 6 hours is used.
	//
	// example:
	//
	// 21600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The format of the transcoded stream recording. Valid values:
	//
	// > If you choose m3u8 or cmaf, you must specify the TranscodeRecordFormat.N.SliceOssObjectPrefix and TranscodeRecordFormat.N.SliceDuration parameters.
	//
	// - m3u8
	//
	// - flv
	//
	// - mp4
	//
	// - cmaf
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The duration of a single segment for transcoded stream recording. Unit: seconds.
	//
	// > This parameter takes effect only if you set the TranscodeRecordFormat.N.Format parameter to m3u8 or cmaf.
	//
	// If you do not specify this parameter, the default value 30 seconds is used. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
}

func (s UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) SetCycleDuration(v int32) *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) SetFormat(v string) *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.Format = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) SetSliceDuration(v int32) *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *UpdateLiveAppRecordConfigRequestTranscodeRecordFormat) Validate() error {
	return dara.Validate(s)
}
