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
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The interruption duration for merge. If the stream interruption duration exceeds the specified duration, a new recording is generated. The value of this parameter ranges from 15 to 21600 seconds.
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
	// The recording end time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  The time range that is specified by the EndTime and StartTime parameters must be less than or equal to seven days. If the value exceeds seven days, ApsaraVideo Live considers seven days as the time range. This parameter takes effect only for the live stream specified by the StreamName parameter. If the StreamName parameter is not specified, this parameter does not take effect.
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
	// The endpoint of the Object Storage Service (OSS) bucket.
	//
	// To store live stream recordings in OSS, you need to create an OSS bucket in advance. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
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
	// The recording start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  The start time must be within seven days after the stream ingest starts. This parameter takes effect only for the live stream specified by the StreamName parameter. If the StreamName parameter is not specified, this parameter does not take effect.
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
	// The transcoded stream recording details.
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
	// The recording cycle. Unit: seconds If you do not specify this parameter, the default value 6 hours is used.
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
	// The recording format. Supported formats include M3U8, Flash Video (FLV), MP4, and Common Media Application Format (CMAF). Valid values:
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
	// The duration of a single segment. Unit: seconds
	//
	// >  This parameter takes effect only if you set the RecordFormat.N.Format parameter to m3u8 or cmaf.
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
	// The transcoded stream recording cycle. Unit: seconds If you do not specify this parameter, the default value 6 hours is used.
	//
	// example:
	//
	// 21600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The format of the transcoded stream recording. Supported formats include M3U8, FLV, MP4, and CMAF. Valid values:
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
	// The duration of a single segment in the transcoded stream recording. Unit: seconds.
	//
	// >  This parameter takes effect only if you set the TranscodeRecordFormat.N.Format parameter to m3u8 or cmaf.
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
