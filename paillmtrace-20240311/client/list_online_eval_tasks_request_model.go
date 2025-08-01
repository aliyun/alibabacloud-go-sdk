// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineEvalTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListOnlineEvalTasksRequest
	GetKeyword() *string
	SetMaxTime(v string) *ListOnlineEvalTasksRequest
	GetMaxTime() *string
	SetMinTime(v string) *ListOnlineEvalTasksRequest
	GetMinTime() *string
	SetPageNumber(v int32) *ListOnlineEvalTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOnlineEvalTasksRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListOnlineEvalTasksRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListOnlineEvalTasksRequest
	GetSortOrder() *string
}

type ListOnlineEvalTasksRequest struct {
	// Search keyword. It will match on fields such as task name, application name (appName), task description, and evaluation metric name.
	//
	// example:
	//
	// foo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The UTC end time of the search time range
	//
	// example:
	//
	// 2025-04-07 13:24:25
	//
	// 2025-04-07
	MaxTime *string `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	// The UTC start time of the search time range
	//
	// example:
	//
	// 2025-04-05 13:24:25
	//
	// 2025-04-05
	MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	// The current page number. Value range: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 50
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy    *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListOnlineEvalTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineEvalTasksRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOnlineEvalTasksRequest) GetMaxTime() *string {
	return s.MaxTime
}

func (s *ListOnlineEvalTasksRequest) GetMinTime() *string {
	return s.MinTime
}

func (s *ListOnlineEvalTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOnlineEvalTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOnlineEvalTasksRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListOnlineEvalTasksRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListOnlineEvalTasksRequest) SetKeyword(v string) *ListOnlineEvalTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetMaxTime(v string) *ListOnlineEvalTasksRequest {
	s.MaxTime = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetMinTime(v string) *ListOnlineEvalTasksRequest {
	s.MinTime = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetPageNumber(v int32) *ListOnlineEvalTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetPageSize(v int32) *ListOnlineEvalTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetSortBy(v string) *ListOnlineEvalTasksRequest {
	s.SortBy = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetSortOrder(v string) *ListOnlineEvalTasksRequest {
	s.SortOrder = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) Validate() error {
	return dara.Validate(s)
}
