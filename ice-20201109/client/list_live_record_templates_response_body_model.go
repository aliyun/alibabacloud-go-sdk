// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *ListLiveRecordTemplatesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListLiveRecordTemplatesResponseBody
	GetPageSize() *int64
	SetRecordTemplateList(v []*ListLiveRecordTemplatesResponseBodyRecordTemplateList) *ListLiveRecordTemplatesResponseBody
	GetRecordTemplateList() []*ListLiveRecordTemplatesResponseBodyRecordTemplateList
	SetRequestId(v string) *ListLiveRecordTemplatesResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLiveRecordTemplatesResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListLiveRecordTemplatesResponseBody
	GetTotalCount() *int64
}

type ListLiveRecordTemplatesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of recording templates.
	RecordTemplateList []*ListLiveRecordTemplatesResponseBodyRecordTemplateList `json:"RecordTemplateList,omitempty" xml:"RecordTemplateList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BEA98A0C-7870-15FE-B96F-8880BB600A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- asc: sorts the query results in ascending order.
	//
	// 	- desc: sorts the query results in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveRecordTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRecordTemplatesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLiveRecordTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLiveRecordTemplatesResponseBody) GetRecordTemplateList() []*ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	return s.RecordTemplateList
}

func (s *ListLiveRecordTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRecordTemplatesResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLiveRecordTemplatesResponseBody) SetPageNo(v int64) *ListLiveRecordTemplatesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBody) SetPageSize(v int64) *ListLiveRecordTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBody) SetRecordTemplateList(v []*ListLiveRecordTemplatesResponseBodyRecordTemplateList) *ListLiveRecordTemplatesResponseBody {
	s.RecordTemplateList = v
	return s
}

func (s *ListLiveRecordTemplatesResponseBody) SetRequestId(v string) *ListLiveRecordTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBody) SetSortBy(v string) *ListLiveRecordTemplatesResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBody) SetTotalCount(v int64) *ListLiveRecordTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBody) Validate() error {
	if s.RecordTemplateList != nil {
		for _, item := range s.RecordTemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveRecordTemplatesResponseBodyRecordTemplateList struct {
	// The time when the job was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-20T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the template was last modified.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-20T03:26:36Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The template name.
	//
	// example:
	//
	// test template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recording formats.
	RecordFormatList []*ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList `json:"RecordFormatList,omitempty" xml:"RecordFormatList,omitempty" type:"Repeated"`
	// The template ID.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveRecordTemplatesResponseBodyRecordTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordTemplatesResponseBodyRecordTemplateList) GoString() string {
	return s.String()
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) GetLastModified() *string {
	return s.LastModified
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) GetName() *string {
	return s.Name
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) GetRecordFormatList() []*ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList {
	return s.RecordFormatList
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) GetType() *string {
	return s.Type
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) SetCreateTime(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) SetLastModified(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	s.LastModified = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) SetName(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	s.Name = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) SetRecordFormatList(v []*ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) *ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	s.RecordFormatList = v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) SetTemplateId(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) SetType(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateList {
	s.Type = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateList) Validate() error {
	if s.RecordFormatList != nil {
		for _, item := range s.RecordFormatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList struct {
	// The duration of the recording cycle. Unit: seconds.
	//
	// example:
	//
	// 21600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The output file format.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the recording file that is stored in Object Storage Service (OSS).
	//
	// example:
	//
	// record/{JobId}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single segment. Unit: seconds.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The name of the TS segment.
	//
	// example:
	//
	// record/{JobId}/{UnixTimestamp}_{Sequence}
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) GoString() string {
	return s.String()
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) GetFormat() *string {
	return s.Format
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) SetCycleDuration(v int32) *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList {
	s.CycleDuration = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) SetFormat(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList {
	s.Format = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) SetOssObjectPrefix(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList {
	s.OssObjectPrefix = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) SetSliceDuration(v int32) *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList {
	s.SliceDuration = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) SetSliceOssObjectPrefix(v string) *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *ListLiveRecordTemplatesResponseBodyRecordTemplateListRecordFormatList) Validate() error {
	return dara.Validate(s)
}
