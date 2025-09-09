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
	SetLifecycleId(v string) *ListSystemLogsRequest
	GetLifecycleId() *string
	SetLogLevel(v string) *ListSystemLogsRequest
	GetLogLevel() *string
	SetLogRepository(v string) *ListSystemLogsRequest
	GetLogRepository() *string
	SetOffset(v string) *ListSystemLogsRequest
	GetOffset() *string
	SetOrder(v string) *ListSystemLogsRequest
	GetOrder() *string
	SetProblemCategory(v string) *ListSystemLogsRequest
	GetProblemCategory() *string
	SetSortBy(v string) *ListSystemLogsRequest
	GetSortBy() *string
	SetSourceRequestId(v string) *ListSystemLogsRequest
	GetSourceRequestId() *string
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
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LifecycleId *string `json:"LifecycleId,omitempty" xml:"LifecycleId,omitempty"`
	// example:
	//
	// Error
	LogLevel      *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	LogRepository *string `json:"LogRepository,omitempty" xml:"LogRepository,omitempty"`
	Offset        *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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

func (s *ListSystemLogsRequest) GetLifecycleId() *string {
	return s.LifecycleId
}

func (s *ListSystemLogsRequest) GetLogLevel() *string {
	return s.LogLevel
}

func (s *ListSystemLogsRequest) GetLogRepository() *string {
	return s.LogRepository
}

func (s *ListSystemLogsRequest) GetOffset() *string {
	return s.Offset
}

func (s *ListSystemLogsRequest) GetOrder() *string {
	return s.Order
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

func (s *ListSystemLogsRequest) SetLifecycleId(v string) *ListSystemLogsRequest {
	s.LifecycleId = &v
	return s
}

func (s *ListSystemLogsRequest) SetLogLevel(v string) *ListSystemLogsRequest {
	s.LogLevel = &v
	return s
}

func (s *ListSystemLogsRequest) SetLogRepository(v string) *ListSystemLogsRequest {
	s.LogRepository = &v
	return s
}

func (s *ListSystemLogsRequest) SetOffset(v string) *ListSystemLogsRequest {
	s.Offset = &v
	return s
}

func (s *ListSystemLogsRequest) SetOrder(v string) *ListSystemLogsRequest {
	s.Order = &v
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

func (s *ListSystemLogsRequest) Validate() error {
	return dara.Validate(s)
}
