// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationProcessDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOperationProcessDetailRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListOperationProcessDetailRequest
	GetEndTime() *int64
	SetLang(v string) *ListOperationProcessDetailRequest
	GetLang() *string
	SetPageSize(v int32) *ListOperationProcessDetailRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListOperationProcessDetailRequest
	GetStartTime() *int64
	SetStatusCodes(v []*int32) *ListOperationProcessDetailRequest
	GetStatusCodes() []*int32
	SetTaskIds(v []*string) *ListOperationProcessDetailRequest
	GetTaskIds() []*string
}

type ListOperationProcessDetailRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1731555850000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1731469330000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The subtask status codes.
	StatusCodes []*int32 `json:"StatusCodes,omitempty" xml:"StatusCodes,omitempty" type:"Repeated"`
	// The IDs of operation tasks.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s ListOperationProcessDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessDetailRequest) GoString() string {
	return s.String()
}

func (s *ListOperationProcessDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOperationProcessDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationProcessDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *ListOperationProcessDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationProcessDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationProcessDetailRequest) GetStatusCodes() []*int32 {
	return s.StatusCodes
}

func (s *ListOperationProcessDetailRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListOperationProcessDetailRequest) SetCurrentPage(v int32) *ListOperationProcessDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationProcessDetailRequest) SetEndTime(v int64) *ListOperationProcessDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationProcessDetailRequest) SetLang(v string) *ListOperationProcessDetailRequest {
	s.Lang = &v
	return s
}

func (s *ListOperationProcessDetailRequest) SetPageSize(v int32) *ListOperationProcessDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationProcessDetailRequest) SetStartTime(v int64) *ListOperationProcessDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ListOperationProcessDetailRequest) SetStatusCodes(v []*int32) *ListOperationProcessDetailRequest {
	s.StatusCodes = v
	return s
}

func (s *ListOperationProcessDetailRequest) SetTaskIds(v []*string) *ListOperationProcessDetailRequest {
	s.TaskIds = v
	return s
}

func (s *ListOperationProcessDetailRequest) Validate() error {
	return dara.Validate(s)
}
