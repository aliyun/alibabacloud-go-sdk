// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAppRecordConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveAppRecordConfigRequest
	GetAppName() *string
	SetDelayTime(v int32) *AddLiveAppRecordConfigRequest
	GetDelayTime() *int32
	SetDomainName(v string) *AddLiveAppRecordConfigRequest
	GetDomainName() *string
	SetEndTime(v string) *AddLiveAppRecordConfigRequest
	GetEndTime() *string
	SetOnDemand(v int32) *AddLiveAppRecordConfigRequest
	GetOnDemand() *int32
	SetOssBucket(v string) *AddLiveAppRecordConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *AddLiveAppRecordConfigRequest
	GetOssEndpoint() *string
	SetOwnerId(v int64) *AddLiveAppRecordConfigRequest
	GetOwnerId() *int64
	SetRecordFormat(v []*AddLiveAppRecordConfigRequestRecordFormat) *AddLiveAppRecordConfigRequest
	GetRecordFormat() []*AddLiveAppRecordConfigRequestRecordFormat
	SetSecurityToken(v string) *AddLiveAppRecordConfigRequest
	GetSecurityToken() *string
	SetStartTime(v string) *AddLiveAppRecordConfigRequest
	GetStartTime() *string
	SetStreamName(v string) *AddLiveAppRecordConfigRequest
	GetStreamName() *string
	SetTranscodeRecordFormat(v []*AddLiveAppRecordConfigRequestTranscodeRecordFormat) *AddLiveAppRecordConfigRequest
	GetTranscodeRecordFormat() []*AddLiveAppRecordConfigRequestTranscodeRecordFormat
	SetTranscodeTemplates(v []*string) *AddLiveAppRecordConfigRequest
	GetTranscodeTemplates() []*string
}

type AddLiveAppRecordConfigRequest struct {
	// The name of the application to which the stream belongs. The template takes effect only when the AppName value matches the AppName in the ingest URL. To match all application names, set this parameter to an asterisk (*).
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The stream discontinuity merging duration. If the live stream is disconnected for longer than the specified merging duration, a new file is generated. Valid values: 15 to 21600. Unit: seconds.
	//
	// example:
	//
	// 180
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The streaming domain of the streamer.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The recording end time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > The difference between EndTime and StartTime cannot exceed 7 days. If it exceeds 7 days, the value is calculated as 7 days. This parameter is valid only for stream-level recording (when StreamName is not empty).
	//
	// example:
	//
	// 2018-04-16T09:57:21Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The on-demand or manual recording mode. Valid values:
	//
	// - **0*	- (default): disabled. Automatic recording is used.
	//
	// - **1**: on-demand recording through HTTP callback. You must first configure OnDemandUrl by calling the [AddLiveRecordNotifyConfig](https://help.aliyun.com/document_detail/2847891.html) operation. Otherwise, recording is not performed by default.
	//
	// - **2**: on-demand recording by parsing stream ingest parameters.
	//
	// - **7**: manual recording. Recording is not performed by default. You can call the [RealTimeRecordCommand](https://help.aliyun.com/document_detail/2847882.html) operation to manually start or stop recording.
	//
	// example:
	//
	// 1
	OnDemand *int32 `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	// The name of the OSS bucket.
	//
	// To store live recordings in OSS, create an OSS bucket in advance. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// To store live recordings in OSS, create an OSS bucket in advance. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-beijing.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recording details.
	RecordFormat  []*AddLiveAppRecordConfigRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
	SecurityToken *string                                      `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The recording start time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > The specified time must be within 7 days of the actual stream ingest start time. This parameter is valid only for stream-level recording (when StreamName is not empty).
	//
	// example:
	//
	// 2018-04-10T09:57:21Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name. The template takes effect only when the StreamName value matches the StreamName in the ingest URL. To match all stream names under the specified AppName, set this parameter to an asterisk (*).
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The transcoded stream recording details.
	TranscodeRecordFormat []*AddLiveAppRecordConfigRequestTranscodeRecordFormat `json:"TranscodeRecordFormat,omitempty" xml:"TranscodeRecordFormat,omitempty" type:"Repeated"`
	// The transcoding template group for transcoded stream recording.
	//
	// example:
	//
	// sd
	TranscodeTemplates []*string `json:"TranscodeTemplates,omitempty" xml:"TranscodeTemplates,omitempty" type:"Repeated"`
}

func (s AddLiveAppRecordConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveAppRecordConfigRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *AddLiveAppRecordConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveAppRecordConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *AddLiveAppRecordConfigRequest) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *AddLiveAppRecordConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *AddLiveAppRecordConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *AddLiveAppRecordConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveAppRecordConfigRequest) GetRecordFormat() []*AddLiveAppRecordConfigRequestRecordFormat {
	return s.RecordFormat
}

func (s *AddLiveAppRecordConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveAppRecordConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddLiveAppRecordConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLiveAppRecordConfigRequest) GetTranscodeRecordFormat() []*AddLiveAppRecordConfigRequestTranscodeRecordFormat {
	return s.TranscodeRecordFormat
}

func (s *AddLiveAppRecordConfigRequest) GetTranscodeTemplates() []*string {
	return s.TranscodeTemplates
}

func (s *AddLiveAppRecordConfigRequest) SetAppName(v string) *AddLiveAppRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetDelayTime(v int32) *AddLiveAppRecordConfigRequest {
	s.DelayTime = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetDomainName(v string) *AddLiveAppRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetEndTime(v string) *AddLiveAppRecordConfigRequest {
	s.EndTime = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOnDemand(v int32) *AddLiveAppRecordConfigRequest {
	s.OnDemand = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOssBucket(v string) *AddLiveAppRecordConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOssEndpoint(v string) *AddLiveAppRecordConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOwnerId(v int64) *AddLiveAppRecordConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetRecordFormat(v []*AddLiveAppRecordConfigRequestRecordFormat) *AddLiveAppRecordConfigRequest {
	s.RecordFormat = v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetSecurityToken(v string) *AddLiveAppRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetStartTime(v string) *AddLiveAppRecordConfigRequest {
	s.StartTime = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetStreamName(v string) *AddLiveAppRecordConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetTranscodeRecordFormat(v []*AddLiveAppRecordConfigRequestTranscodeRecordFormat) *AddLiveAppRecordConfigRequest {
	s.TranscodeRecordFormat = v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetTranscodeTemplates(v []*string) *AddLiveAppRecordConfigRequest {
	s.TranscodeTemplates = v
	return s
}

func (s *AddLiveAppRecordConfigRequest) Validate() error {
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

type AddLiveAppRecordConfigRequestRecordFormat struct {
	// The recording length per epoch. Unit: seconds.
	//
	// > - If this parameter is not specified, the default value varies by recording format: 6 hours for m3u8 and cmaf formats, and 1 hour for flv and mp4 formats.
	//
	// > - If a live stream is disconnected within a recording epoch but resumes stream ingest within the stream discontinuity merging duration, recording continues in the same file. This is Normal behavior.
	//
	// > - A recording file is generated only after the live stream is disconnected for longer than the stream discontinuity merging duration.
	//
	// example:
	//
	// 1
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The format. M3U8, FLV, MP4, and CMAF are supported. Valid values:
	//
	// 	Notice: At least one of RecordFormat and TranscodeRecordFormat must be set. If you select m3u8 or cmaf, you must also set the request parameters RecordFormat.N.SliceOssObjectPrefix and RecordFormat.N.SliceDuration.
	//
	//
	// - m3u8.
	//
	// - flv.
	//
	// - mp4.
	//
	// - cmaf.
	//
	// > Settings for RecordFormat and TranscodeRecordFormat: at least one must be specified.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the recording file stored in OSS.
	//
	// - The file name must be less than 256 bytes and supports variable matching, including {AppName}, {StreamName}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime}.
	//
	// - The value must contain {StartTime} or {EscapedStartTime} and {EndTime} or {EscapedEndTime}.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The segment length of a single segment. Unit: seconds.
	//
	// 	Notice: This parameter takes effect only when RecordFormat.N.Format is set to m3u8 or cmaf.
	//
	//
	// If this parameter is not specified, the default value is 30 seconds. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The segment name.
	//
	// 	Notice: This parameter is required only when RecordFormat.N.Format is set to m3u8 or cmaf.
	//
	//
	// - The default segment length is 30 seconds. The value must be less than 256 bytes and supports variable matching, including {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - The value must contain the {UnixTimestamp} and {Sequence} variables.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{UnixTimestamp}_{Sequence}
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s AddLiveAppRecordConfigRequestRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppRecordConfigRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetCycleDuration(v int32) *AddLiveAppRecordConfigRequestRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetFormat(v string) *AddLiveAppRecordConfigRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetOssObjectPrefix(v string) *AddLiveAppRecordConfigRequestRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetSliceDuration(v int32) *AddLiveAppRecordConfigRequestRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetSliceOssObjectPrefix(v string) *AddLiveAppRecordConfigRequestRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) Validate() error {
	return dara.Validate(s)
}

type AddLiveAppRecordConfigRequestTranscodeRecordFormat struct {
	// The recording length per epoch for transcoding stream recording. Unit: seconds.
	//
	// > If this parameter is not specified, the default value varies by recording format: 6 hours for m3u8 and cmaf formats, and 1 hour for flv and mp4 formats.
	//
	// example:
	//
	// 21600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The transcoding stream recording format. M3U8, FLV, MP4, and CMAF are supported. Valid values:
	//
	// 	Notice: If you select m3u8 or cmaf, you must also set the request parameters TranscodeRecordFormat.N.SliceOssObjectPrefix and TranscodeRecordFormat.N.SliceDuration.
	//
	//
	// - m3u8.
	//
	// - flv.
	//
	// - mp4.
	//
	// - cmaf.
	//
	// > Settings: if you select m3u8 or cmaf format, the corresponding slice parameters must also be configured.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the transcoded stream recording file stored in OSS.
	//
	// - The file name must be less than 256 bytes and supports variable matching, including {AppName}, {StreamName}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime}.
	//
	// - The value must contain {StartTime} or {EscapedStartTime} and {EndTime} or {EscapedEndTime}.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The segment length of a single segment for transcoding stream recording. Unit: seconds.
	//
	// 	Notice: This parameter takes effect only when TranscodeRecordFormat.N.Format (transcoding stream recording format) is set to m3u8 or cmaf.
	//
	//
	// If this parameter is not specified, the default value is 30 seconds. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The segment name for transcoded stream recording.
	//
	// 	Notice: This parameter is required only when TranscodeRecordFormat.N.Format is set to m3u8 or cmaf.
	//
	//
	// - The default segment length is 30 seconds. The value must be less than 256 bytes and supports variable matching, including {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - The value must contain the {UnixTimestamp} and {Sequence} variables.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{UnixTimestamp}_{Sequence}
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s AddLiveAppRecordConfigRequestTranscodeRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppRecordConfigRequestTranscodeRecordFormat) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) SetCycleDuration(v int32) *AddLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) SetFormat(v string) *AddLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.Format = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) SetOssObjectPrefix(v string) *AddLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) SetSliceDuration(v int32) *AddLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) SetSliceOssObjectPrefix(v string) *AddLiveAppRecordConfigRequestTranscodeRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestTranscodeRecordFormat) Validate() error {
	return dara.Validate(s)
}
