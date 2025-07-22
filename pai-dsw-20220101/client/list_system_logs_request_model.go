// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtEndTime(v string) *ListSystemLogsRequest
	GetGmtEndTime() *string
	SetGmtStartTime(v string) *ListSystemLogsRequest
	GetGmtStartTime() *string
	SetInstanceId(v string) *ListSystemLogsRequest
	GetInstanceId() *string
	SetLogLevel(v string) *ListSystemLogsRequest
	GetLogLevel() *string
	SetOrder(v string) *ListSystemLogsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListSystemLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSystemLogsRequest
	GetPageSize() *int64
	SetProblemCategory(v string) *ListSystemLogsRequest
	GetProblemCategory() *string
	SetSortBy(v string) *ListSystemLogsRequest
	GetSortBy() *string
	SetSourceRequestId(v string) *ListSystemLogsRequest
	GetSourceRequestId() *string
	SetSourceType(v string) *ListSystemLogsRequest
	GetSourceType() *string
}

type ListSystemLogsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2020-12-08T16:00:00Z
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Error
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// InstanceStartFailed
	ProblemCategory *string `json:"ProblemCategory,omitempty" xml:"ProblemCategory,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 473469C7-******5-B3DB-A3DC0DE3C83E
	SourceRequestId *string `json:"SourceRequestId,omitempty" xml:"SourceRequestId,omitempty"`
	// example:
	//
	// NotebookMainContainerLogs
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListSystemLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemLogsRequest) GoString() string {
	return s.String()
}

func (s *ListSystemLogsRequest) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *ListSystemLogsRequest) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *ListSystemLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSystemLogsRequest) GetLogLevel() *string {
	return s.LogLevel
}

func (s *ListSystemLogsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSystemLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSystemLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSystemLogsRequest) GetProblemCategory() *string {
	return s.ProblemCategory
}

func (s *ListSystemLogsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSystemLogsRequest) GetSourceRequestId() *string {
	return s.SourceRequestId
}

func (s *ListSystemLogsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSystemLogsRequest) SetGmtEndTime(v string) *ListSystemLogsRequest {
	s.GmtEndTime = &v
	return s
}

func (s *ListSystemLogsRequest) SetGmtStartTime(v string) *ListSystemLogsRequest {
	s.GmtStartTime = &v
	return s
}

func (s *ListSystemLogsRequest) SetInstanceId(v string) *ListSystemLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSystemLogsRequest) SetLogLevel(v string) *ListSystemLogsRequest {
	s.LogLevel = &v
	return s
}

func (s *ListSystemLogsRequest) SetOrder(v string) *ListSystemLogsRequest {
	s.Order = &v
	return s
}

func (s *ListSystemLogsRequest) SetPageNumber(v int64) *ListSystemLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSystemLogsRequest) SetPageSize(v int64) *ListSystemLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemLogsRequest) SetProblemCategory(v string) *ListSystemLogsRequest {
	s.ProblemCategory = &v
	return s
}

func (s *ListSystemLogsRequest) SetSortBy(v string) *ListSystemLogsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSystemLogsRequest) SetSourceRequestId(v string) *ListSystemLogsRequest {
	s.SourceRequestId = &v
	return s
}

func (s *ListSystemLogsRequest) SetSourceType(v string) *ListSystemLogsRequest {
	s.SourceType = &v
	return s
}

func (s *ListSystemLogsRequest) Validate() error {
	return dara.Validate(s)
}
