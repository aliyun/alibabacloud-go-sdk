// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutopilotTuningHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListAutopilotTuningHistoriesRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *ListAutopilotTuningHistoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAutopilotTuningHistoriesRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListAutopilotTuningHistoriesRequest
	GetStartTime() *int64
}

type ListAutopilotTuningHistoriesRequest struct {
	// The query end timestamp in milliseconds. If not specified, the default is the current time. The time span between startTime and endTime cannot exceed 30 days.
	//
	// example:
	//
	// 1689321600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query start timestamp in milliseconds. If not specified, the default is the last 3 days.
	//
	// example:
	//
	// 1689235200000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListAutopilotTuningHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutopilotTuningHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAutopilotTuningHistoriesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAutopilotTuningHistoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAutopilotTuningHistoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutopilotTuningHistoriesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAutopilotTuningHistoriesRequest) SetEndTime(v int64) *ListAutopilotTuningHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAutopilotTuningHistoriesRequest) SetPageNumber(v int32) *ListAutopilotTuningHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAutopilotTuningHistoriesRequest) SetPageSize(v int32) *ListAutopilotTuningHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutopilotTuningHistoriesRequest) SetStartTime(v int64) *ListAutopilotTuningHistoriesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAutopilotTuningHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
