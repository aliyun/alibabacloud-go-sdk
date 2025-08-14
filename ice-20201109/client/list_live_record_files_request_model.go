// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLiveRecordFilesRequest
	GetEndTime() *string
	SetJobIds(v []*string) *ListLiveRecordFilesRequest
	GetJobIds() []*string
	SetPageNo(v int32) *ListLiveRecordFilesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveRecordFilesRequest
	GetPageSize() *int32
	SetRecordFormat(v string) *ListLiveRecordFilesRequest
	GetRecordFormat() *string
	SetSortBy(v string) *ListLiveRecordFilesRequest
	GetSortBy() *string
	SetStartTime(v string) *ListLiveRecordFilesRequest
	GetStartTime() *string
}

type ListLiveRecordFilesRequest struct {
	// The end of the time range to query. The maximum time range to query is four days. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of job IDs.
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 5 to 30. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The format of the recording file. Valid values:
	//
	// M3U8, FLV, and MP4
	//
	// example:
	//
	// m3u8
	RecordFormat *string `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty"`
	// The sorting order of the index files by creation time. Valid values:
	//
	// asc: The query results are displayed in ascending order. This is the default value.
	//
	// desc: The query results are displayed in descending order.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T08:00:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLiveRecordFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordFilesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRecordFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLiveRecordFilesRequest) GetJobIds() []*string {
	return s.JobIds
}

func (s *ListLiveRecordFilesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveRecordFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveRecordFilesRequest) GetRecordFormat() *string {
	return s.RecordFormat
}

func (s *ListLiveRecordFilesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLiveRecordFilesRequest) SetEndTime(v string) *ListLiveRecordFilesRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveRecordFilesRequest) SetJobIds(v []*string) *ListLiveRecordFilesRequest {
	s.JobIds = v
	return s
}

func (s *ListLiveRecordFilesRequest) SetPageNo(v int32) *ListLiveRecordFilesRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordFilesRequest) SetPageSize(v int32) *ListLiveRecordFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordFilesRequest) SetRecordFormat(v string) *ListLiveRecordFilesRequest {
	s.RecordFormat = &v
	return s
}

func (s *ListLiveRecordFilesRequest) SetSortBy(v string) *ListLiveRecordFilesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordFilesRequest) SetStartTime(v string) *ListLiveRecordFilesRequest {
	s.StartTime = &v
	return s
}

func (s *ListLiveRecordFilesRequest) Validate() error {
	return dara.Validate(s)
}
