// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallback(v string) *DescribeTemplateResponseBody
	GetCallback() *string
	SetCreatedTime(v string) *DescribeTemplateResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeTemplateResponseBody
	GetDescription() *string
	SetFileFormat(v string) *DescribeTemplateResponseBody
	GetFileFormat() *string
	SetFlv(v string) *DescribeTemplateResponseBody
	GetFlv() *string
	SetHlsM3u8(v string) *DescribeTemplateResponseBody
	GetHlsM3u8() *string
	SetHlsTs(v string) *DescribeTemplateResponseBody
	GetHlsTs() *string
	SetId(v string) *DescribeTemplateResponseBody
	GetId() *string
	SetInterval(v int64) *DescribeTemplateResponseBody
	GetInterval() *int64
	SetJpgOnDemand(v string) *DescribeTemplateResponseBody
	GetJpgOnDemand() *string
	SetJpgOverwrite(v string) *DescribeTemplateResponseBody
	GetJpgOverwrite() *string
	SetJpgSequence(v string) *DescribeTemplateResponseBody
	GetJpgSequence() *string
	SetMp4(v string) *DescribeTemplateResponseBody
	GetMp4() *string
	SetName(v string) *DescribeTemplateResponseBody
	GetName() *string
	SetOssBucket(v string) *DescribeTemplateResponseBody
	GetOssBucket() *string
	SetOssEndpoint(v string) *DescribeTemplateResponseBody
	GetOssEndpoint() *string
	SetOssFilePrefix(v string) *DescribeTemplateResponseBody
	GetOssFilePrefix() *string
	SetRegion(v string) *DescribeTemplateResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeTemplateResponseBody
	GetRequestId() *string
	SetRetention(v int64) *DescribeTemplateResponseBody
	GetRetention() *int64
	SetTransConfigs(v []*DescribeTemplateResponseBodyTransConfigs) *DescribeTemplateResponseBody
	GetTransConfigs() []*DescribeTemplateResponseBodyTransConfigs
	SetTrigger(v string) *DescribeTemplateResponseBody
	GetTrigger() *string
	SetType(v string) *DescribeTemplateResponseBody
	GetType() *string
}

type DescribeTemplateResponseBody struct {
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// 2020-12-10T10:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hls
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	Flv *string `json:"Flv,omitempty" xml:"Flv,omitempty"`
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	HlsM3u8 *string `json:"HlsM3u8,omitempty" xml:"HlsM3u8,omitempty"`
	// example:
	//
	// osspath/record/{StreamName}/{UnixTimestamp}_{Sequence}
	HlsTs *string `json:"HlsTs,omitempty" xml:"HlsTs,omitempty"`
	// example:
	//
	// 323*****998-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 3600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}/{UnixTimestamp}_ondemand.jpg
	JpgOnDemand *string `json:"JpgOnDemand,omitempty" xml:"JpgOnDemand,omitempty"`
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}.jpg
	JpgOverwrite *string `json:"JpgOverwrite,omitempty" xml:"JpgOverwrite,omitempty"`
	// example:
	//
	// osspath/snapshot/{AppName}/{StreamName}/{UnixTimestamp}.jpg
	JpgSequence *string `json:"JpgSequence,omitempty" xml:"JpgSequence,omitempty"`
	// example:
	//
	// osspath/record/{StreamName}/{EscapedStartTime}_{EscapedEndTime}
	Mp4  *string `json:"Mp4,omitempty" xml:"Mp4,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// my_oss_bucket
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// example:
	//
	// oss-cn-qingdao.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// example:
	//
	// oss-prefix
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	Retention    *int64                                      `json:"Retention,omitempty" xml:"Retention,omitempty"`
	TransConfigs []*DescribeTemplateResponseBodyTransConfigs `json:"TransConfigs,omitempty" xml:"TransConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResponseBody) GetCallback() *string {
	return s.Callback
}

func (s *DescribeTemplateResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeTemplateResponseBody) GetFileFormat() *string {
	return s.FileFormat
}

func (s *DescribeTemplateResponseBody) GetFlv() *string {
	return s.Flv
}

func (s *DescribeTemplateResponseBody) GetHlsM3u8() *string {
	return s.HlsM3u8
}

func (s *DescribeTemplateResponseBody) GetHlsTs() *string {
	return s.HlsTs
}

func (s *DescribeTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeTemplateResponseBody) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeTemplateResponseBody) GetJpgOnDemand() *string {
	return s.JpgOnDemand
}

func (s *DescribeTemplateResponseBody) GetJpgOverwrite() *string {
	return s.JpgOverwrite
}

func (s *DescribeTemplateResponseBody) GetJpgSequence() *string {
	return s.JpgSequence
}

func (s *DescribeTemplateResponseBody) GetMp4() *string {
	return s.Mp4
}

func (s *DescribeTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeTemplateResponseBody) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeTemplateResponseBody) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeTemplateResponseBody) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *DescribeTemplateResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplateResponseBody) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeTemplateResponseBody) GetTransConfigs() []*DescribeTemplateResponseBodyTransConfigs {
	return s.TransConfigs
}

func (s *DescribeTemplateResponseBody) GetTrigger() *string {
	return s.Trigger
}

func (s *DescribeTemplateResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeTemplateResponseBody) SetCallback(v string) *DescribeTemplateResponseBody {
	s.Callback = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetCreatedTime(v string) *DescribeTemplateResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetDescription(v string) *DescribeTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetFileFormat(v string) *DescribeTemplateResponseBody {
	s.FileFormat = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetFlv(v string) *DescribeTemplateResponseBody {
	s.Flv = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetHlsM3u8(v string) *DescribeTemplateResponseBody {
	s.HlsM3u8 = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetHlsTs(v string) *DescribeTemplateResponseBody {
	s.HlsTs = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetId(v string) *DescribeTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetInterval(v int64) *DescribeTemplateResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetJpgOnDemand(v string) *DescribeTemplateResponseBody {
	s.JpgOnDemand = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetJpgOverwrite(v string) *DescribeTemplateResponseBody {
	s.JpgOverwrite = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetJpgSequence(v string) *DescribeTemplateResponseBody {
	s.JpgSequence = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetMp4(v string) *DescribeTemplateResponseBody {
	s.Mp4 = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetName(v string) *DescribeTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetOssBucket(v string) *DescribeTemplateResponseBody {
	s.OssBucket = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetOssEndpoint(v string) *DescribeTemplateResponseBody {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetOssFilePrefix(v string) *DescribeTemplateResponseBody {
	s.OssFilePrefix = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetRegion(v string) *DescribeTemplateResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetRequestId(v string) *DescribeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetRetention(v int64) *DescribeTemplateResponseBody {
	s.Retention = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetTransConfigs(v []*DescribeTemplateResponseBodyTransConfigs) *DescribeTemplateResponseBody {
	s.TransConfigs = v
	return s
}

func (s *DescribeTemplateResponseBody) SetTrigger(v string) *DescribeTemplateResponseBody {
	s.Trigger = &v
	return s
}

func (s *DescribeTemplateResponseBody) SetType(v string) *DescribeTemplateResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplateResponseBodyTransConfigs struct {
	// example:
	//
	// 25
	Fps *int64 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 50
	Gop *int64 `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// example:
	//
	// 720
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 399*****430-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sd
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 800
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// h264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// example:
	//
	// 1280
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeTemplateResponseBodyTransConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResponseBodyTransConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetFps() *int64 {
	return s.Fps
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetGop() *int64 {
	return s.Gop
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetHeight() *int64 {
	return s.Height
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetId() *string {
	return s.Id
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetName() *string {
	return s.Name
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetVideoBitrate() *int64 {
	return s.VideoBitrate
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *DescribeTemplateResponseBodyTransConfigs) GetWidth() *int64 {
	return s.Width
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetFps(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Fps = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetGop(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Gop = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetHeight(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Height = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetId(v string) *DescribeTemplateResponseBodyTransConfigs {
	s.Id = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetName(v string) *DescribeTemplateResponseBodyTransConfigs {
	s.Name = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetVideoBitrate(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetVideoCodec(v string) *DescribeTemplateResponseBodyTransConfigs {
	s.VideoCodec = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) SetWidth(v int64) *DescribeTemplateResponseBodyTransConfigs {
	s.Width = &v
	return s
}

func (s *DescribeTemplateResponseBodyTransConfigs) Validate() error {
	return dara.Validate(s)
}
