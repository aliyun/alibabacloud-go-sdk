// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobInstanceEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTrainingJobInstanceEventsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListTrainingJobInstanceEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTrainingJobInstanceEventsRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListTrainingJobInstanceEventsRequest
	GetStartTime() *string
}

type ListTrainingJobInstanceEventsRequest struct {
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

func (s ListTrainingJobInstanceEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceEventsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobInstanceEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTrainingJobInstanceEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTrainingJobInstanceEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobInstanceEventsRequest) SetEndTime(v string) *ListTrainingJobInstanceEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) SetPageNumber(v int64) *ListTrainingJobInstanceEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) SetPageSize(v int64) *ListTrainingJobInstanceEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) SetStartTime(v string) *ListTrainingJobInstanceEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) Validate() error {
	return dara.Validate(s)
}
