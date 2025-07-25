// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTrainingJobEventsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListTrainingJobEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTrainingJobEventsRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListTrainingJobEventsRequest
	GetStartTime() *string
}

type ListTrainingJobEventsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTrainingJobEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobEventsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTrainingJobEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTrainingJobEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobEventsRequest) SetEndTime(v string) *ListTrainingJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobEventsRequest) SetPageNumber(v int64) *ListTrainingJobEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobEventsRequest) SetPageSize(v int64) *ListTrainingJobEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobEventsRequest) SetStartTime(v string) *ListTrainingJobEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobEventsRequest) Validate() error {
	return dara.Validate(s)
}
