// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTrainingJobLogsRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListTrainingJobLogsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListTrainingJobLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTrainingJobLogsRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListTrainingJobLogsRequest
	GetStartTime() *string
	SetWorkerId(v string) *ListTrainingJobLogsRequest
	GetWorkerId() *string
}

type ListTrainingJobLogsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// train129f212o89d-master-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// example:
	//
	// train129f212o89d-master-0
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s ListTrainingJobLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrainingJobLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTrainingJobLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTrainingJobLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobLogsRequest) GetWorkerId() *string {
	return s.WorkerId
}

func (s *ListTrainingJobLogsRequest) SetEndTime(v string) *ListTrainingJobLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetInstanceId(v string) *ListTrainingJobLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetPageNumber(v int64) *ListTrainingJobLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetPageSize(v int64) *ListTrainingJobLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetStartTime(v string) *ListTrainingJobLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetWorkerId(v string) *ListTrainingJobLogsRequest {
	s.WorkerId = &v
	return s
}

func (s *ListTrainingJobLogsRequest) Validate() error {
	return dara.Validate(s)
}
