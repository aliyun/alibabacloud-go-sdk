// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTranscodeJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *ListLiveTranscodeJobsRequest
	GetKeyWord() *string
	SetPageNo(v int32) *ListLiveTranscodeJobsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveTranscodeJobsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListLiveTranscodeJobsRequest
	GetSortBy() *string
	SetStartMode(v int32) *ListLiveTranscodeJobsRequest
	GetStartMode() *int32
	SetStatus(v int32) *ListLiveTranscodeJobsRequest
	GetStatus() *int32
	SetType(v string) *ListLiveTranscodeJobsRequest
	GetType() *string
}

type ListLiveTranscodeJobsRequest struct {
	// The search keyword. You can use the job ID or name as the keyword to search for jobs. If you search for jobs by name, fuzzy match is supported.
	//
	// example:
	//
	// 24ecbb5c-4f98-4194-9400-f17102e27fc5
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order. Valid values:
	//
	// 	- asc
	//
	// 	- desc
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start mode of the transcoding job.
	//
	// 	- 0: The transcoding job immediately starts.
	//
	// 	- 1: The transcoding job starts at the scheduled time.
	//
	// example:
	//
	// 0
	StartMode *int32 `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
	// The state of the job.
	//
	// 0: The job is not started. 1: The job is in progress. 2: The job is stopped.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the template used by the transcoding job.
	//
	// 	- normal
	//
	// 	- narrow-band
	//
	// 	- audio-only
	//
	// 	- origin
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveTranscodeJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListLiveTranscodeJobsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveTranscodeJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveTranscodeJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveTranscodeJobsRequest) GetStartMode() *int32 {
	return s.StartMode
}

func (s *ListLiveTranscodeJobsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListLiveTranscodeJobsRequest) GetType() *string {
	return s.Type
}

func (s *ListLiveTranscodeJobsRequest) SetKeyWord(v string) *ListLiveTranscodeJobsRequest {
	s.KeyWord = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) SetPageNo(v int32) *ListLiveTranscodeJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) SetPageSize(v int32) *ListLiveTranscodeJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) SetSortBy(v string) *ListLiveTranscodeJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) SetStartMode(v int32) *ListLiveTranscodeJobsRequest {
	s.StartMode = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) SetStatus(v int32) *ListLiveTranscodeJobsRequest {
	s.Status = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) SetType(v string) *ListLiveTranscodeJobsRequest {
	s.Type = &v
	return s
}

func (s *ListLiveTranscodeJobsRequest) Validate() error {
	return dara.Validate(s)
}
