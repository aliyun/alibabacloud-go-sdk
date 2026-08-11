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
	// The search keyword. You can search by task ID or name. Name supports fuzzy match.
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
	// The sort order. Sorts by CreateTime. Default value: desc.
	//
	// - asc: ascending order
	//
	// - desc: descending order
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start mode.
	//
	// - 0: start immediately
	//
	// - 1: scheduled start
	//
	// example:
	//
	// 0
	StartMode *int32 `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
	// The task status.
	//
	// - 0: not started
	//
	// - 1: running
	//
	// - 2: stopped
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template type corresponding to the transcoding task.
	//
	// - normal: standard
	//
	// - narrow-band: narrow bandwidth high definition
	//
	// - audio-only: audio only
	//
	// - origin: original quality
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
