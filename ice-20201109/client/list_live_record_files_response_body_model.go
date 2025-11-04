// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*ListLiveRecordFilesResponseBodyFiles) *ListLiveRecordFilesResponseBody
	GetFiles() []*ListLiveRecordFilesResponseBodyFiles
	SetPageNo(v int64) *ListLiveRecordFilesResponseBody
	GetPageNo() *int64
	SetPageSize(v string) *ListLiveRecordFilesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListLiveRecordFilesResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLiveRecordFilesResponseBody
	GetSortBy() *string
	SetTotalCount(v string) *ListLiveRecordFilesResponseBody
	GetTotalCount() *string
}

type ListLiveRecordFilesResponseBody struct {
	// The list of index files.
	Files []*ListLiveRecordFilesResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
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
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DE24625C-7C0F-4020-8448-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the index files by creation time.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of files that meet the specified conditions.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveRecordFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRecordFilesResponseBody) GetFiles() []*ListLiveRecordFilesResponseBodyFiles {
	return s.Files
}

func (s *ListLiveRecordFilesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLiveRecordFilesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListLiveRecordFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRecordFilesResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordFilesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListLiveRecordFilesResponseBody) SetFiles(v []*ListLiveRecordFilesResponseBodyFiles) *ListLiveRecordFilesResponseBody {
	s.Files = v
	return s
}

func (s *ListLiveRecordFilesResponseBody) SetPageNo(v int64) *ListLiveRecordFilesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordFilesResponseBody) SetPageSize(v string) *ListLiveRecordFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordFilesResponseBody) SetRequestId(v string) *ListLiveRecordFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRecordFilesResponseBody) SetSortBy(v string) *ListLiveRecordFilesResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordFilesResponseBody) SetTotalCount(v string) *ListLiveRecordFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveRecordFilesResponseBody) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveRecordFilesResponseBodyFiles struct {
	// The time when the file was created in UTC.
	//
	// example:
	//
	// 2016-05-27T09:40:56Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The recording length. Unit: seconds.
	//
	// example:
	//
	// 100.0
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T07:36:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The format of the recording file.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The height of the video.
	//
	// example:
	//
	// 640
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the recording job.
	//
	// example:
	//
	// c4d7f0a4-b506-43f9-8de3-07732c3f****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the recording job.
	//
	// example:
	//
	// LiveRecordJob***
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The ID of the index file.
	//
	// example:
	//
	// c4d7f0a4-b506-43f9-8de3-07732c3f****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The storage information about the recording file.
	//
	// example:
	//
	// { "Type": "oss", "Endpoint":"oss-cn-shanghai.aliyuncs.com", "Bucket": "test-bucket" }
	RecordOutput *string `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty"`
	// The URL of the index file.
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T07:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// LiveStream***
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	// The width of the video.
	//
	// example:
	//
	// 480
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListLiveRecordFilesResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetDuration() *float32 {
	return s.Duration
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetFormat() *string {
	return s.Format
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetHeight() *int32 {
	return s.Height
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetJobId() *string {
	return s.JobId
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetJobName() *string {
	return s.JobName
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetRecordId() *string {
	return s.RecordId
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetRecordOutput() *string {
	return s.RecordOutput
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *ListLiveRecordFilesResponseBodyFiles) GetWidth() *int32 {
	return s.Width
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetCreateTime(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.CreateTime = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetDuration(v float32) *ListLiveRecordFilesResponseBodyFiles {
	s.Duration = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetEndTime(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.EndTime = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetFormat(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.Format = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetHeight(v int32) *ListLiveRecordFilesResponseBodyFiles {
	s.Height = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetJobId(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.JobId = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetJobName(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.JobName = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetRecordId(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.RecordId = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetRecordOutput(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.RecordOutput = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetRecordUrl(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.RecordUrl = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetStartTime(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.StartTime = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetStreamUrl(v string) *ListLiveRecordFilesResponseBodyFiles {
	s.StreamUrl = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) SetWidth(v int32) *ListLiveRecordFilesResponseBodyFiles {
	s.Width = &v
	return s
}

func (s *ListLiveRecordFilesResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
