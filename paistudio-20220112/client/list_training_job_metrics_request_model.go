// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTrainingJobMetricsRequest
	GetEndTime() *string
	SetName(v string) *ListTrainingJobMetricsRequest
	GetName() *string
	SetOrder(v string) *ListTrainingJobMetricsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListTrainingJobMetricsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTrainingJobMetricsRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListTrainingJobMetricsRequest
	GetStartTime() *string
}

type ListTrainingJobMetricsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// accuracy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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

func (s ListTrainingJobMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobMetricsRequest) GetName() *string {
	return s.Name
}

func (s *ListTrainingJobMetricsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTrainingJobMetricsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTrainingJobMetricsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTrainingJobMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobMetricsRequest) SetEndTime(v string) *ListTrainingJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetName(v string) *ListTrainingJobMetricsRequest {
	s.Name = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetOrder(v string) *ListTrainingJobMetricsRequest {
	s.Order = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetPageNumber(v int64) *ListTrainingJobMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetPageSize(v int64) *ListTrainingJobMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetStartTime(v string) *ListTrainingJobMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) Validate() error {
	return dara.Validate(s)
}
