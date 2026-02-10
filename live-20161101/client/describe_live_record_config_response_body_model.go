// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveAppRecordList(v *DescribeLiveRecordConfigResponseBodyLiveAppRecordList) *DescribeLiveRecordConfigResponseBody
	GetLiveAppRecordList() *DescribeLiveRecordConfigResponseBodyLiveAppRecordList
	SetOrder(v string) *DescribeLiveRecordConfigResponseBody
	GetOrder() *string
	SetPageNum(v int32) *DescribeLiveRecordConfigResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveRecordConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveRecordConfigResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveRecordConfigResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveRecordConfigResponseBody
	GetTotalPage() *int32
}

type DescribeLiveRecordConfigResponseBody struct {
	LiveAppRecordList *DescribeLiveRecordConfigResponseBodyLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" type:"Struct"`
	// The sorting order of recording configurations by creation time.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 5
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5056369B-D337-499E-B8B7-B761BD37B08A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of recording configurations that meet the specified conditions.
	//
	// example:
	//
	// 12
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 20
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveRecordConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBody) GetLiveAppRecordList() *DescribeLiveRecordConfigResponseBodyLiveAppRecordList {
	return s.LiveAppRecordList
}

func (s *DescribeLiveRecordConfigResponseBody) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveRecordConfigResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveRecordConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveRecordConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRecordConfigResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveRecordConfigResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveRecordConfigResponseBody) SetLiveAppRecordList(v *DescribeLiveRecordConfigResponseBodyLiveAppRecordList) *DescribeLiveRecordConfigResponseBody {
	s.LiveAppRecordList = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) SetOrder(v string) *DescribeLiveRecordConfigResponseBody {
	s.Order = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) SetPageNum(v int32) *DescribeLiveRecordConfigResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) SetPageSize(v int32) *DescribeLiveRecordConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) SetRequestId(v string) *DescribeLiveRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) SetTotalNum(v int32) *DescribeLiveRecordConfigResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) SetTotalPage(v int32) *DescribeLiveRecordConfigResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBody) Validate() error {
	if s.LiveAppRecordList != nil {
		if err := s.LiveAppRecordList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordList struct {
	LiveAppRecord []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" type:"Repeated"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordList) GetLiveAppRecord() []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	return s.LiveAppRecord
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordList) SetLiveAppRecord(v []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) *DescribeLiveRecordConfigResponseBodyLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordList) Validate() error {
	if s.LiveAppRecord != nil {
		for _, item := range s.LiveAppRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord struct {
	AppName                   *string                                                                                      `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime                *string                                                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DelayTime                 *int32                                                                                       `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DomainName                *string                                                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                   *string                                                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OnDemond                  *int32                                                                                       `json:"OnDemond,omitempty" xml:"OnDemond,omitempty"`
	OssBucket                 *string                                                                                      `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint               *string                                                                                      `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	RecordFormatList          *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList          `json:"RecordFormatList,omitempty" xml:"RecordFormatList,omitempty" type:"Struct"`
	StartTime                 *string                                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName                *string                                                                                      `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	TranscodeRecordFormatList *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList `json:"TranscodeRecordFormatList,omitempty" xml:"TranscodeRecordFormatList,omitempty" type:"Struct"`
	TranscodeTemplates        *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates        `json:"TranscodeTemplates,omitempty" xml:"TranscodeTemplates,omitempty" type:"Struct"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetOnDemond() *int32 {
	return s.OnDemond
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetRecordFormatList() *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList {
	return s.RecordFormatList
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetTranscodeRecordFormatList() *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList {
	return s.TranscodeRecordFormatList
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) GetTranscodeTemplates() *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates {
	return s.TranscodeTemplates
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetCreateTime(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetDelayTime(v int32) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.DelayTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetOnDemond(v int32) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.OnDemond = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetOssBucket(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetOssEndpoint(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetRecordFormatList(v *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.RecordFormatList = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetTranscodeRecordFormatList(v *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.TranscodeRecordFormatList = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) SetTranscodeTemplates(v *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.TranscodeTemplates = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecord) Validate() error {
	if s.RecordFormatList != nil {
		if err := s.RecordFormatList.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeRecordFormatList != nil {
		if err := s.TranscodeRecordFormatList.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeTemplates != nil {
		if err := s.TranscodeTemplates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList struct {
	RecordFormat []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList) GetRecordFormat() []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	return s.RecordFormat
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList) SetRecordFormat(v []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList {
	s.RecordFormat = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatList) Validate() error {
	if s.RecordFormat != nil {
		for _, item := range s.RecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat struct {
	CycleDuration        *int32  `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty"`
	OssObjectPrefix      *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	SliceDuration        *int32  `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetCycleDuration(v int32) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetFormat(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.Format = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetOssObjectPrefix(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetSliceDuration(v int32) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetSliceOssObjectPrefix(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList struct {
	RecordFormat []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList) GetRecordFormat() []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat {
	return s.RecordFormat
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList) SetRecordFormat(v []*DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList {
	s.RecordFormat = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatList) Validate() error {
	if s.RecordFormat != nil {
		for _, item := range s.RecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat struct {
	CycleDuration        *int32  `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty"`
	OssObjectPrefix      *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	SliceDuration        *int32  `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) SetCycleDuration(v int32) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) SetFormat(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat {
	s.Format = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) SetOssObjectPrefix(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) SetSliceDuration(v int32) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) SetSliceOssObjectPrefix(v string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeRecordFormatListRecordFormat) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates struct {
	Templates []*string `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates) GetTemplates() []*string {
	return s.Templates
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates) SetTemplates(v []*string) *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates {
	s.Templates = v
	return s
}

func (s *DescribeLiveRecordConfigResponseBodyLiveAppRecordListLiveAppRecordTranscodeTemplates) Validate() error {
	return dara.Validate(s)
}
