// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextStartTime(v string) *DescribeRecordsResponseBody
	GetNextStartTime() *string
	SetPageCount(v int64) *DescribeRecordsResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeRecordsResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeRecordsResponseBody
	GetPageSize() *int64
	SetRecords(v []*DescribeRecordsResponseBodyRecords) *DescribeRecordsResponseBody
	GetRecords() []*DescribeRecordsResponseBodyRecords
	SetRequestId(v string) *DescribeRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeRecordsResponseBody
	GetTotalCount() *int64
}

type DescribeRecordsResponseBody struct {
	// example:
	//
	// 2018-12-10T11:00:00Z
	NextStartTime *string `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
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
	PageSize *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*DescribeRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponseBody) GetNextStartTime() *string {
	return s.NextStartTime
}

func (s *DescribeRecordsResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeRecordsResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeRecordsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRecordsResponseBody) GetRecords() []*DescribeRecordsResponseBodyRecords {
	return s.Records
}

func (s *DescribeRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRecordsResponseBody) SetNextStartTime(v string) *DescribeRecordsResponseBody {
	s.NextStartTime = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetPageCount(v int64) *DescribeRecordsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetPageNum(v int64) *DescribeRecordsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetPageSize(v int64) *DescribeRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetRecords(v []*DescribeRecordsResponseBodyRecords) *DescribeRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeRecordsResponseBody) SetRequestId(v string) *DescribeRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordsResponseBody) SetTotalCount(v int64) *DescribeRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRecordsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecordsResponseBodyRecords struct {
	// example:
	//
	// 2021-11-23T18:33:48
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// hls
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// example:
	//
	// 1080
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 2be2a673-6033-4874-b6f2-f2bc0a1*****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// record/live/310*****007/2021-11-23-18-19-38_2021-11-23-18-33-48.m3u8
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	// example:
	//
	// 2021-11-23T18:19:32
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	// example:
	//
	// 388*****204-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// http://my_oss_bucket.oss-cn-qingdao.aliyuncs.com/record/live/310*****007/2021-11-23-18-19-38_2021-11-23-18-33-48.m3u8
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 1920
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponseBodyRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRecordsResponseBodyRecords) GetFileFormat() *string {
	return s.FileFormat
}

func (s *DescribeRecordsResponseBodyRecords) GetHeight() *int64 {
	return s.Height
}

func (s *DescribeRecordsResponseBodyRecords) GetId() *string {
	return s.Id
}

func (s *DescribeRecordsResponseBodyRecords) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeRecordsResponseBodyRecords) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeRecordsResponseBodyRecords) GetOssObject() *string {
	return s.OssObject
}

func (s *DescribeRecordsResponseBodyRecords) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordsResponseBodyRecords) GetStreamId() *string {
	return s.StreamId
}

func (s *DescribeRecordsResponseBodyRecords) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeRecordsResponseBodyRecords) GetType() *string {
	return s.Type
}

func (s *DescribeRecordsResponseBodyRecords) GetUrl() *string {
	return s.Url
}

func (s *DescribeRecordsResponseBodyRecords) GetWidth() *int64 {
	return s.Width
}

func (s *DescribeRecordsResponseBodyRecords) SetEndTime(v string) *DescribeRecordsResponseBodyRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetFileFormat(v string) *DescribeRecordsResponseBodyRecords {
	s.FileFormat = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetHeight(v int64) *DescribeRecordsResponseBodyRecords {
	s.Height = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetId(v string) *DescribeRecordsResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetOssBucket(v string) *DescribeRecordsResponseBodyRecords {
	s.OssBucket = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetOssEndpoint(v string) *DescribeRecordsResponseBodyRecords {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetOssObject(v string) *DescribeRecordsResponseBodyRecords {
	s.OssObject = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetStartTime(v string) *DescribeRecordsResponseBodyRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetStreamId(v string) *DescribeRecordsResponseBodyRecords {
	s.StreamId = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetTemplateId(v string) *DescribeRecordsResponseBodyRecords {
	s.TemplateId = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetType(v string) *DescribeRecordsResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetUrl(v string) *DescribeRecordsResponseBodyRecords {
	s.Url = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) SetWidth(v int64) *DescribeRecordsResponseBodyRecords {
	s.Width = &v
	return s
}

func (s *DescribeRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
