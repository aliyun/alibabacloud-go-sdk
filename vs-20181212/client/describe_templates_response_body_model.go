// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageCount(v int64) *DescribeTemplatesResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeTemplatesResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody
	GetTemplates() []*DescribeTemplatesResponseBodyTemplates
	SetTotalCount(v int64) *DescribeTemplatesResponseBody
	GetTotalCount() *int64
}

type DescribeTemplatesResponseBody struct {
	// example:
	//
	// 5
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F3F88C96-CA6E-573E-B8F7-5BE83A1A0BCF
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeTemplatesResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplatesResponseBody) GetTemplates() []*DescribeTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTemplatesResponseBody) SetPageCount(v int64) *DescribeTemplatesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageNum(v int64) *DescribeTemplatesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageSize(v int64) *DescribeTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetRequestId(v string) *DescribeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTotalCount(v int64) *DescribeTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTemplatesResponseBodyTemplates struct {
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// 2018-12-10T10:00:00Z
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
	// my_prefix
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 3
	Retention    *int64                                                `json:"Retention,omitempty" xml:"Retention,omitempty"`
	TransConfigs []*DescribeTemplatesResponseBodyTemplatesTransConfigs `json:"TransConfigs,omitempty" xml:"TransConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplates) GetCallback() *string {
	return s.Callback
}

func (s *DescribeTemplatesResponseBodyTemplates) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeTemplatesResponseBodyTemplates) GetFileFormat() *string {
	return s.FileFormat
}

func (s *DescribeTemplatesResponseBodyTemplates) GetFlv() *string {
	return s.Flv
}

func (s *DescribeTemplatesResponseBodyTemplates) GetHlsM3u8() *string {
	return s.HlsM3u8
}

func (s *DescribeTemplatesResponseBodyTemplates) GetHlsTs() *string {
	return s.HlsTs
}

func (s *DescribeTemplatesResponseBodyTemplates) GetId() *string {
	return s.Id
}

func (s *DescribeTemplatesResponseBodyTemplates) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeTemplatesResponseBodyTemplates) GetJpgOnDemand() *string {
	return s.JpgOnDemand
}

func (s *DescribeTemplatesResponseBodyTemplates) GetJpgOverwrite() *string {
	return s.JpgOverwrite
}

func (s *DescribeTemplatesResponseBodyTemplates) GetJpgSequence() *string {
	return s.JpgSequence
}

func (s *DescribeTemplatesResponseBodyTemplates) GetMp4() *string {
	return s.Mp4
}

func (s *DescribeTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeTemplatesResponseBodyTemplates) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeTemplatesResponseBodyTemplates) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeTemplatesResponseBodyTemplates) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *DescribeTemplatesResponseBodyTemplates) GetRegion() *string {
	return s.Region
}

func (s *DescribeTemplatesResponseBodyTemplates) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeTemplatesResponseBodyTemplates) GetTransConfigs() []*DescribeTemplatesResponseBodyTemplatesTransConfigs {
	return s.TransConfigs
}

func (s *DescribeTemplatesResponseBodyTemplates) GetTrigger() *string {
	return s.Trigger
}

func (s *DescribeTemplatesResponseBodyTemplates) GetType() *string {
	return s.Type
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCallback(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Callback = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCreatedTime(v string) *DescribeTemplatesResponseBodyTemplates {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetFileFormat(v string) *DescribeTemplatesResponseBodyTemplates {
	s.FileFormat = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetFlv(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Flv = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetHlsM3u8(v string) *DescribeTemplatesResponseBodyTemplates {
	s.HlsM3u8 = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetHlsTs(v string) *DescribeTemplatesResponseBodyTemplates {
	s.HlsTs = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetInterval(v int64) *DescribeTemplatesResponseBodyTemplates {
	s.Interval = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetJpgOnDemand(v string) *DescribeTemplatesResponseBodyTemplates {
	s.JpgOnDemand = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetJpgOverwrite(v string) *DescribeTemplatesResponseBodyTemplates {
	s.JpgOverwrite = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetJpgSequence(v string) *DescribeTemplatesResponseBodyTemplates {
	s.JpgSequence = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetMp4(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Mp4 = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetName(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetOssBucket(v string) *DescribeTemplatesResponseBodyTemplates {
	s.OssBucket = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetOssEndpoint(v string) *DescribeTemplatesResponseBodyTemplates {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetOssFilePrefix(v string) *DescribeTemplatesResponseBodyTemplates {
	s.OssFilePrefix = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetRegion(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Region = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetRetention(v int64) *DescribeTemplatesResponseBodyTemplates {
	s.Retention = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTransConfigs(v []*DescribeTemplatesResponseBodyTemplatesTransConfigs) *DescribeTemplatesResponseBodyTemplates {
	s.TransConfigs = v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTrigger(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Trigger = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetType(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) Validate() error {
	if s.TransConfigs != nil {
		for _, item := range s.TransConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTemplatesResponseBodyTemplatesTransConfigs struct {
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
	// example:
	//
	// 399788187729597430-cn-qingdao
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DescribeTemplatesResponseBodyTemplatesTransConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplatesTransConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetFps() *int64 {
	return s.Fps
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetGop() *int64 {
	return s.Gop
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetHeight() *int64 {
	return s.Height
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetName() *string {
	return s.Name
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetVideoBitrate() *int64 {
	return s.VideoBitrate
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetWidth() *int64 {
	return s.Width
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) GetId() *string {
	return s.Id
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetFps(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Fps = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetGop(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Gop = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetHeight(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Height = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetName(v string) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetVideoBitrate(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetVideoCodec(v string) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.VideoCodec = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetWidth(v int64) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Width = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) SetId(v string) *DescribeTemplatesResponseBodyTemplatesTransConfigs {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplatesTransConfigs) Validate() error {
	return dara.Validate(s)
}
