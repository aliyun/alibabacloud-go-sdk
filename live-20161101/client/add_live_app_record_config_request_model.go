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
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name in the ingest URL. Otherwise, the configuration does not take effect. If you want to match all applications, specify an asterisk (\\*) as the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Duration for stream concatenation. If the live streaming interruption exceeds the set concatenation duration, a new file will be generated. The concatenation duration can be set between 15 to 21600 seconds.
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
	// Recording end time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// > The difference between EndTime and StartTime should not exceed 7 days; if it does, it will be calculated as 7 days. This is only valid for stream-level recording (when StreamName is not empty).
	//
	// example:
	//
	// 2018-04-16T09:57:21Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to enable on-demand recording. Valid values:
	//
	// 	- **0**: disables on-demand recording.
	//
	// 	- **1**: enables on-demand recording by using the HTTP callback method.
	//
	// 	- **2**: enables on-demand recording by parsing the stream ingest parameters.
	//
	// 	- **7**: By default, ApsaraVideo Live does not automatically record live streams. You can call the [RealTimeRecordCommand](https://help.aliyun.com/document_detail/2847882.html) operation to manually start or stop recording.
	//
	// >  If you set the OnDemand parameter to **1**, you need to call the [AddLiveRecordNotifyConfig](https://help.aliyun.com/document_detail/2847891.html) operation to configure the OnDemandUrl parameter. Otherwise, ApsaraVideo Live does not perform on-demand recording.
	//
	// example:
	//
	// 1
	OnDemand *int32 `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	// The name of the OSS bucket where live streaming recording files are stored. To store recorded files in OSS, you need to create an OSS bucket in advance. For creation method, please refer to [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// To store live stream recordings in OSS, you need to create an OSS bucket in advance. For more information, see Configure OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// learn.developer.aliyundoc.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recording details.
	RecordFormat  []*AddLiveAppRecordConfigRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
	SecurityToken *string                                      `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Start time of the recording. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// > The set time must be within 7 days of the actual streaming start time, and is only valid for stream-level recording (when StreamName is not empty).
	//
	// example:
	//
	// 2018-04-10T09:57:21Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Stream broadcast name.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The transcoded stream recording details.
	TranscodeRecordFormat []*AddLiveAppRecordConfigRequestTranscodeRecordFormat `json:"TranscodeRecordFormat,omitempty" xml:"TranscodeRecordFormat,omitempty" type:"Repeated"`
	// Transcoding stream recording template group.
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
	return dara.Validate(s)
}

type AddLiveAppRecordConfigRequestRecordFormat struct {
	// The recording cycle. Unit: seconds. If you do not specify this parameter, the default value 6 hours is used.
	//
	// >
	//
	// 	- If a live stream is interrupted during a recording cycle but is resumed within the interruption duration threshold, the stream is recorded in the same recording before and after the interruption.
	//
	// 	- If a live stream is interrupted for longer than the interruption duration threshold, a new recording is generated.
	//
	// example:
	//
	// 1
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The recording format. Supported formats include M3U8, FLV, MP4, and CMAF. Valid values:
	//
	// >  You need to specify at lease one of the RecordFormat and TranscodeRecordFormat parameters. If you set this parameter to m3u8 or cmaf, you must also specify the RecordFormat.N.SliceOssObjectPrefix and RecordFormat.N.SliceDuration parameters.
	//
	// 	- m3u8
	//
	// 	- flv
	//
	// 	- mp4
	//
	// 	- cmaf
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The naming format of a recording to store in OSS.
	//
	// 	- The name must be less than 256 bytes in length and can contain the {AppName}, {StreamName}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime} variables.
	//
	// 	- The name must contain the {StartTime} and {EndTime} variables or the {EscapedStartTime} and {EscapedEndTime} variables.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single segment. Unit: seconds.
	//
	// >  This parameter takes effect only if you set the RecordFormat.N.Format parameter to m3u8 or cmaf.
	//
	// If you do not specify this parameter, the default value 30 seconds is used. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The naming format of a segment.
	//
	// >  This parameter is required only if you set the RecordFormat.N.Format parameter to m3u8 or cmaf.
	//
	// 	- By default, the duration of a segment is 30 seconds. The segment name must be less than 256 bytes in length and can contain the {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence} variables.
	//
	// 	- The segment name must contain the {UnixTimestamp} and {Sequence} variables.
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
	// The transcoded stream recording cycle. Unit: seconds. If you do not specify this parameter, the default value 6 hours is used.
	//
	// example:
	//
	// 21600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The transcoded stream recording format. Supported formats include M3U8, FLV, MP4, and CMAF. Valid values:
	//
	// >  If you set this parameter to m3u8 or cmaf, you must also specify the TranscodeRecordFormat.N.SliceOssObjectPrefix and TranscodeRecordFormat.N.SliceDuration parameters.
	//
	// 	- m3u8
	//
	// 	- flv
	//
	// 	- mp4
	//
	// 	- cmaf
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The naming format of a transcoded stream recording to store in OSS.
	//
	// 	- The name must be less than 256 bytes in length and can contain the {AppName}, {StreamName}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime} variables.
	//
	// 	- The name must contain the {StartTime} and {EndTime} variables or the {EscapedStartTime} and {EscapedEndTime} variables.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single segment in a transcoded stream recording. Unit: seconds.
	//
	// >  This parameter takes effect only if you set the TranscodeRecordFormat.N.Format parameter to m3u8 or cmaf.
	//
	// If you do not specify this parameter, the default value 30 seconds is used. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The naming format of a segment in a transcoded stream recording.
	//
	// >  This parameter is required only if you set the TranscodeRecordFormat.N.Format parameter to m3u8 or cmaf.
	//
	// 	- By default, the duration of a segment is 30 seconds. The segment name must be less than 256 bytes in length and can contain the {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence} variables.
	//
	// 	- The segment name must contain the {UnixTimestamp} and {Sequence} variables.
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
