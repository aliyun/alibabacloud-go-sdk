// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *ListExecutorEventsRequestFilter) *ListExecutorEventsRequest
	GetFilter() *ListExecutorEventsRequestFilter
	SetPageNumber(v int32) *ListExecutorEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExecutorEventsRequest
	GetPageSize() *int32
}

type ListExecutorEventsRequest struct {
	Filter *ListExecutorEventsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExecutorEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorEventsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorEventsRequest) GetFilter() *ListExecutorEventsRequestFilter {
	return s.Filter
}

func (s *ListExecutorEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorEventsRequest) SetFilter(v *ListExecutorEventsRequestFilter) *ListExecutorEventsRequest {
	s.Filter = v
	return s
}

func (s *ListExecutorEventsRequest) SetPageNumber(v int32) *ListExecutorEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorEventsRequest) SetPageSize(v int32) *ListExecutorEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExecutorEventsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExecutorEventsRequestFilter struct {
	ExecutorIds []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 1703820113
	TimeAfter *int64 `json:"TimeAfter,omitempty" xml:"TimeAfter,omitempty"`
	// example:
	//
	// 1703819914
	TimeBefore *int64 `json:"TimeBefore,omitempty" xml:"TimeBefore,omitempty"`
}

func (s ListExecutorEventsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorEventsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListExecutorEventsRequestFilter) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *ListExecutorEventsRequestFilter) GetJobId() *string {
	return s.JobId
}

func (s *ListExecutorEventsRequestFilter) GetLevel() *string {
	return s.Level
}

func (s *ListExecutorEventsRequestFilter) GetTimeAfter() *int64 {
	return s.TimeAfter
}

func (s *ListExecutorEventsRequestFilter) GetTimeBefore() *int64 {
	return s.TimeBefore
}

func (s *ListExecutorEventsRequestFilter) SetExecutorIds(v []*string) *ListExecutorEventsRequestFilter {
	s.ExecutorIds = v
	return s
}

func (s *ListExecutorEventsRequestFilter) SetJobId(v string) *ListExecutorEventsRequestFilter {
	s.JobId = &v
	return s
}

func (s *ListExecutorEventsRequestFilter) SetLevel(v string) *ListExecutorEventsRequestFilter {
	s.Level = &v
	return s
}

func (s *ListExecutorEventsRequestFilter) SetTimeAfter(v int64) *ListExecutorEventsRequestFilter {
	s.TimeAfter = &v
	return s
}

func (s *ListExecutorEventsRequestFilter) SetTimeBefore(v int64) *ListExecutorEventsRequestFilter {
	s.TimeBefore = &v
	return s
}

func (s *ListExecutorEventsRequestFilter) Validate() error {
	return dara.Validate(s)
}
