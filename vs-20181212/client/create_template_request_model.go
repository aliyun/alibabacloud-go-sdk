// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *CreateTemplateRequest
	GetCallback() *string
	SetDescription(v string) *CreateTemplateRequest
	GetDescription() *string
	SetFileFormat(v string) *CreateTemplateRequest
	GetFileFormat() *string
	SetFlv(v string) *CreateTemplateRequest
	GetFlv() *string
	SetHlsM3u8(v string) *CreateTemplateRequest
	GetHlsM3u8() *string
	SetHlsTs(v string) *CreateTemplateRequest
	GetHlsTs() *string
	SetInterval(v int64) *CreateTemplateRequest
	GetInterval() *int64
	SetJpgOnDemand(v string) *CreateTemplateRequest
	GetJpgOnDemand() *string
	SetJpgOverwrite(v string) *CreateTemplateRequest
	GetJpgOverwrite() *string
	SetJpgSequence(v string) *CreateTemplateRequest
	GetJpgSequence() *string
	SetMp4(v string) *CreateTemplateRequest
	GetMp4() *string
	SetName(v string) *CreateTemplateRequest
	GetName() *string
	SetOssBucket(v string) *CreateTemplateRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *CreateTemplateRequest
	GetOssEndpoint() *string
	SetOssFilePrefix(v string) *CreateTemplateRequest
	GetOssFilePrefix() *string
	SetOwnerId(v int64) *CreateTemplateRequest
	GetOwnerId() *int64
	SetRegion(v string) *CreateTemplateRequest
	GetRegion() *string
	SetRetention(v int64) *CreateTemplateRequest
	GetRetention() *int64
	SetTransConfigsJSON(v string) *CreateTemplateRequest
	GetTransConfigsJSON() *string
	SetTrigger(v string) *CreateTemplateRequest
	GetTrigger() *string
	SetType(v string) *CreateTemplateRequest
	GetType() *string
}

type CreateTemplateRequest struct {
	// Callback URL to be invoked after template execution.
	//
	// > Note: Templates triggered on demand do not support callback parameters.
	//
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// Template description.
	//
	// example:
	//
	// 录制模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Storage file format. Multiple values are separated by commas. Valid values:
	//
	// - mp4
	//
	// - flv
	//
	// - hls
	//
	// > The Qingdao ingest endpoint does not support recording in FLV or MP4 formats.
	//
	// example:
	//
	// hls
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// Storage path for FLV files. For the format, see the description for Mp4.
	//
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	Flv *string `json:"Flv,omitempty" xml:"Flv,omitempty"`
	// Storage path for HLS m3u8 files. For the format, see the description for Mp4.
	//
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	HlsM3u8 *string `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	// Storage path for HLS .ts files.
	//
	// - Variables can be used in the path. Supported variables include {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - The variables {UnixTimestamp} and {Sequence} must both be included.
	//
	// example:
	//
	// osspath/record/{StreamName}/{UnixTimestamp}_{Sequence}
	HlsTs *string `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	// Operation epoch, in seconds.
	//
	// example:
	//
	// 3600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Storage path for on-demand JPG screenshots.
	//
	// - Only JPG images are currently supported.
	//
	// - Variables can be used in the path. Supported variables include {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - Either {UnixTimestamp} or {Sequence} must be included.
	//
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}/{UnixTimestamp}_ondemand.jpg
	JpgOnDemand *string `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	// Storage path for JPG files used in overwrite snapshots.
	//
	// - Only JPG images are currently supported.
	//
	// - Supports variable substitution with {AppName} and {StreamName}.
	//
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}.jpg
	JpgOverwrite *string `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	// Storage path for JPG files used in sequential snapshots.
	//
	// - Only JPG images are currently supported.
	//
	// - Supports variable substitution with {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - Either {UnixTimestamp} or {Sequence} is required.
	//
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}/{UnixTimestamp}.jpg
	JpgSequence *string `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	// Storage path for MP4 files.
	//
	// - The path supports variable substitution. Available variables include {AppName}, {StreamName}, {Sequence}, {EscapedStartTime}, and {EscapedEndTime}.
	//
	// - {EscapedStartTime} and {EscapedEndTime} are required.
	//
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	Mp4 *string `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	// Template Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 录制模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS bucket.
	//
	// example:
	//
	// bucketname
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// Domain name of OSS.
	//
	// example:
	//
	// oss-cn-qingdao.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// OSS file prefix.
	//
	// example:
	//
	// oss-prefix
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region where the OSS bucket resides, that is, the service center.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Time-shift retention period, in days.
	//
	// example:
	//
	// 3
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// An array of TransConfig-type transcoding configurations, formatted as a JSON string.
	//
	// example:
	//
	// [{"Fps":25,"Gop":50,"Height":720,"VideoCodec":"h264","Width":1280,"Name":"sd","VideoBitrate":800}]
	TransConfigsJSON *string `json:"TransConfigsJSON,omitempty" xml:"TransConfigsJSON,omitempty"`
	// Template trigger type. Default value: auto. Valid values:
	//
	// - auto (automatic)
	//
	// - ondemand (on-demand)
	//
	// example:
	//
	// auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// Template type. Valid values:
	//
	// - record (recording)
	//
	// - snapshot (snapshot)
	//
	// - transcode (transcoding)
	//
	// - timeshift (time shifting)
	//
	// This parameter is required.
	//
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetCallback() *string {
	return s.Callback
}

func (s *CreateTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateTemplateRequest) GetFlv() *string {
	return s.Flv
}

func (s *CreateTemplateRequest) GetHlsM3u8() *string {
	return s.HlsM3u8
}

func (s *CreateTemplateRequest) GetHlsTs() *string {
	return s.HlsTs
}

func (s *CreateTemplateRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateTemplateRequest) GetJpgOnDemand() *string {
	return s.JpgOnDemand
}

func (s *CreateTemplateRequest) GetJpgOverwrite() *string {
	return s.JpgOverwrite
}

func (s *CreateTemplateRequest) GetJpgSequence() *string {
	return s.JpgSequence
}

func (s *CreateTemplateRequest) GetMp4() *string {
	return s.Mp4
}

func (s *CreateTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateTemplateRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *CreateTemplateRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *CreateTemplateRequest) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *CreateTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTemplateRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateTemplateRequest) GetRetention() *int64 {
	return s.Retention
}

func (s *CreateTemplateRequest) GetTransConfigsJSON() *string {
	return s.TransConfigsJSON
}

func (s *CreateTemplateRequest) GetTrigger() *string {
	return s.Trigger
}

func (s *CreateTemplateRequest) GetType() *string {
	return s.Type
}

func (s *CreateTemplateRequest) SetCallback(v string) *CreateTemplateRequest {
	s.Callback = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetFileFormat(v string) *CreateTemplateRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateTemplateRequest) SetFlv(v string) *CreateTemplateRequest {
	s.Flv = &v
	return s
}

func (s *CreateTemplateRequest) SetHlsM3u8(v string) *CreateTemplateRequest {
	s.HlsM3u8 = &v
	return s
}

func (s *CreateTemplateRequest) SetHlsTs(v string) *CreateTemplateRequest {
	s.HlsTs = &v
	return s
}

func (s *CreateTemplateRequest) SetInterval(v int64) *CreateTemplateRequest {
	s.Interval = &v
	return s
}

func (s *CreateTemplateRequest) SetJpgOnDemand(v string) *CreateTemplateRequest {
	s.JpgOnDemand = &v
	return s
}

func (s *CreateTemplateRequest) SetJpgOverwrite(v string) *CreateTemplateRequest {
	s.JpgOverwrite = &v
	return s
}

func (s *CreateTemplateRequest) SetJpgSequence(v string) *CreateTemplateRequest {
	s.JpgSequence = &v
	return s
}

func (s *CreateTemplateRequest) SetMp4(v string) *CreateTemplateRequest {
	s.Mp4 = &v
	return s
}

func (s *CreateTemplateRequest) SetName(v string) *CreateTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplateRequest) SetOssBucket(v string) *CreateTemplateRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateTemplateRequest) SetOssEndpoint(v string) *CreateTemplateRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateTemplateRequest) SetOssFilePrefix(v string) *CreateTemplateRequest {
	s.OssFilePrefix = &v
	return s
}

func (s *CreateTemplateRequest) SetOwnerId(v int64) *CreateTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTemplateRequest) SetRegion(v string) *CreateTemplateRequest {
	s.Region = &v
	return s
}

func (s *CreateTemplateRequest) SetRetention(v int64) *CreateTemplateRequest {
	s.Retention = &v
	return s
}

func (s *CreateTemplateRequest) SetTransConfigsJSON(v string) *CreateTemplateRequest {
	s.TransConfigsJSON = &v
	return s
}

func (s *CreateTemplateRequest) SetTrigger(v string) *CreateTemplateRequest {
	s.Trigger = &v
	return s
}

func (s *CreateTemplateRequest) SetType(v string) *CreateTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
