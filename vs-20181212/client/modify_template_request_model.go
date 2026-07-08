// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *ModifyTemplateRequest
	GetCallback() *string
	SetDescription(v string) *ModifyTemplateRequest
	GetDescription() *string
	SetFileFormat(v string) *ModifyTemplateRequest
	GetFileFormat() *string
	SetFlv(v string) *ModifyTemplateRequest
	GetFlv() *string
	SetHlsM3u8(v string) *ModifyTemplateRequest
	GetHlsM3u8() *string
	SetHlsTs(v string) *ModifyTemplateRequest
	GetHlsTs() *string
	SetId(v string) *ModifyTemplateRequest
	GetId() *string
	SetInterval(v int64) *ModifyTemplateRequest
	GetInterval() *int64
	SetJpgOnDemand(v string) *ModifyTemplateRequest
	GetJpgOnDemand() *string
	SetJpgOverwrite(v string) *ModifyTemplateRequest
	GetJpgOverwrite() *string
	SetJpgSequence(v string) *ModifyTemplateRequest
	GetJpgSequence() *string
	SetMp4(v string) *ModifyTemplateRequest
	GetMp4() *string
	SetName(v string) *ModifyTemplateRequest
	GetName() *string
	SetOssBucket(v string) *ModifyTemplateRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *ModifyTemplateRequest
	GetOssEndpoint() *string
	SetOssFilePrefix(v string) *ModifyTemplateRequest
	GetOssFilePrefix() *string
	SetOwnerId(v int64) *ModifyTemplateRequest
	GetOwnerId() *int64
	SetRegion(v string) *ModifyTemplateRequest
	GetRegion() *string
	SetRetention(v int64) *ModifyTemplateRequest
	GetRetention() *int64
	SetTransConfigsJSON(v string) *ModifyTemplateRequest
	GetTransConfigsJSON() *string
	SetTrigger(v string) *ModifyTemplateRequest
	GetTrigger() *string
}

type ModifyTemplateRequest struct {
	// The callback URL that is used after the template is executed.
	//
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// 录制模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file format for storage. Separate multiple values with commas. Valid values:
	//
	// - mp4
	//
	// - flv
	//
	// - hls
	//
	// > Recording in FLV and MP4 formats is not supported in the China (Qingdao) region.
	//
	// example:
	//
	// hls
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The storage path for FLV files. For information about the format, see the description of the Mp4 parameter.
	//
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	Flv *string `json:"Flv,omitempty" xml:"Flv,omitempty"`
	// The storage path for HLS M3U8 files. For information about the format, see the description of the Mp4 parameter.
	//
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	HlsM3u8 *string `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	// The storage path for HLS TS files.
	//
	// - The path supports variables such as {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - You must include the {UnixTimestamp} and {Sequence} variables.
	//
	// example:
	//
	// osspath/record/{StreamName}/{UnixTimestamp}_{Sequence}
	HlsTs *string `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The operation interval, in seconds.
	//
	// example:
	//
	// 3600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The storage path for JPG files. This path is used for on-demand snapshots.
	//
	// - Only JPG images are supported.
	//
	// - The path supports variables such as {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - You must include either {UnixTimestamp} or {Sequence}.
	//
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}/{UnixTimestamp}_ondemand.jpg
	JpgOnDemand *string `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	// The storage path for JPG files. This path is used to overwrite snapshots.
	//
	// - Only JPG images are supported.
	//
	// - The path supports variables such as {AppName} and {StreamName}.
	//
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}.jpg
	JpgOverwrite *string `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	// The storage path for JPG files. This path is used for sequence snapshots.
	//
	// - Only JPG images are supported.
	//
	// - The path supports variables such as {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}.
	//
	// - You must include either {UnixTimestamp} or {Sequence}.
	//
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}/{UnixTimestamp}.jpg
	JpgSequence *string `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	// The storage path for MP4 files.
	//
	// - The path supports variables such as {AppName}, {StreamName}, {Sequence}, {EscapedStartTime}, and {EscapedEndTime}.
	//
	// - You must include {EscapedStartTime} and {EscapedEndTime}.
	//
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	Mp4 *string `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// 录制模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS bucket.
	//
	// example:
	//
	// bucketname
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The domain name of the OSS bucket.
	//
	// example:
	//
	// oss-cn-qingdao.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The prefix of the OSS file.
	//
	// example:
	//
	// oss-prefix
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the Object Storage Service (OSS) bucket is located.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The retention period for time shifting, in days.
	//
	// example:
	//
	// 3
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// An array of transcoding configurations of the TransConfig type, in a JSON-formatted string.
	//
	// example:
	//
	// [{"Fps":25,"Gop":50,"Height":720,"VideoCodec":"h264","Width":1280,"Name":"sd","VideoBitrate":800}]
	TransConfigsJSON *string `json:"TransConfigsJSON,omitempty" xml:"TransConfigsJSON,omitempty"`
	// The trigger type of the template. The default value is auto. Valid values:
	//
	// - auto (automatic)
	//
	// - ondemand (on-demand)
	//
	// example:
	//
	// auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s ModifyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequest) GetCallback() *string {
	return s.Callback
}

func (s *ModifyTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyTemplateRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *ModifyTemplateRequest) GetFlv() *string {
	return s.Flv
}

func (s *ModifyTemplateRequest) GetHlsM3u8() *string {
	return s.HlsM3u8
}

func (s *ModifyTemplateRequest) GetHlsTs() *string {
	return s.HlsTs
}

func (s *ModifyTemplateRequest) GetId() *string {
	return s.Id
}

func (s *ModifyTemplateRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *ModifyTemplateRequest) GetJpgOnDemand() *string {
	return s.JpgOnDemand
}

func (s *ModifyTemplateRequest) GetJpgOverwrite() *string {
	return s.JpgOverwrite
}

func (s *ModifyTemplateRequest) GetJpgSequence() *string {
	return s.JpgSequence
}

func (s *ModifyTemplateRequest) GetMp4() *string {
	return s.Mp4
}

func (s *ModifyTemplateRequest) GetName() *string {
	return s.Name
}

func (s *ModifyTemplateRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *ModifyTemplateRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *ModifyTemplateRequest) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *ModifyTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTemplateRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyTemplateRequest) GetRetention() *int64 {
	return s.Retention
}

func (s *ModifyTemplateRequest) GetTransConfigsJSON() *string {
	return s.TransConfigsJSON
}

func (s *ModifyTemplateRequest) GetTrigger() *string {
	return s.Trigger
}

func (s *ModifyTemplateRequest) SetCallback(v string) *ModifyTemplateRequest {
	s.Callback = &v
	return s
}

func (s *ModifyTemplateRequest) SetDescription(v string) *ModifyTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyTemplateRequest) SetFileFormat(v string) *ModifyTemplateRequest {
	s.FileFormat = &v
	return s
}

func (s *ModifyTemplateRequest) SetFlv(v string) *ModifyTemplateRequest {
	s.Flv = &v
	return s
}

func (s *ModifyTemplateRequest) SetHlsM3u8(v string) *ModifyTemplateRequest {
	s.HlsM3u8 = &v
	return s
}

func (s *ModifyTemplateRequest) SetHlsTs(v string) *ModifyTemplateRequest {
	s.HlsTs = &v
	return s
}

func (s *ModifyTemplateRequest) SetId(v string) *ModifyTemplateRequest {
	s.Id = &v
	return s
}

func (s *ModifyTemplateRequest) SetInterval(v int64) *ModifyTemplateRequest {
	s.Interval = &v
	return s
}

func (s *ModifyTemplateRequest) SetJpgOnDemand(v string) *ModifyTemplateRequest {
	s.JpgOnDemand = &v
	return s
}

func (s *ModifyTemplateRequest) SetJpgOverwrite(v string) *ModifyTemplateRequest {
	s.JpgOverwrite = &v
	return s
}

func (s *ModifyTemplateRequest) SetJpgSequence(v string) *ModifyTemplateRequest {
	s.JpgSequence = &v
	return s
}

func (s *ModifyTemplateRequest) SetMp4(v string) *ModifyTemplateRequest {
	s.Mp4 = &v
	return s
}

func (s *ModifyTemplateRequest) SetName(v string) *ModifyTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyTemplateRequest) SetOssBucket(v string) *ModifyTemplateRequest {
	s.OssBucket = &v
	return s
}

func (s *ModifyTemplateRequest) SetOssEndpoint(v string) *ModifyTemplateRequest {
	s.OssEndpoint = &v
	return s
}

func (s *ModifyTemplateRequest) SetOssFilePrefix(v string) *ModifyTemplateRequest {
	s.OssFilePrefix = &v
	return s
}

func (s *ModifyTemplateRequest) SetOwnerId(v int64) *ModifyTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTemplateRequest) SetRegion(v string) *ModifyTemplateRequest {
	s.Region = &v
	return s
}

func (s *ModifyTemplateRequest) SetRetention(v int64) *ModifyTemplateRequest {
	s.Retention = &v
	return s
}

func (s *ModifyTemplateRequest) SetTransConfigsJSON(v string) *ModifyTemplateRequest {
	s.TransConfigsJSON = &v
	return s
}

func (s *ModifyTemplateRequest) SetTrigger(v string) *ModifyTemplateRequest {
	s.Trigger = &v
	return s
}

func (s *ModifyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
