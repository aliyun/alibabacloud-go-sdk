// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMonitorDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListApplicationMonitorDetectResultRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListApplicationMonitorDetectResultRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *ListApplicationMonitorDetectResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationMonitorDetectResultRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListApplicationMonitorDetectResultRequest
	GetRegionId() *string
	SetTaskId(v string) *ListApplicationMonitorDetectResultRequest
	GetTaskId() *string
}

type ListApplicationMonitorDetectResultRequest struct {
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638288000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640164683
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The origin probing task ID.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListApplicationMonitorDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorDetectResultRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListApplicationMonitorDetectResultRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListApplicationMonitorDetectResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationMonitorDetectResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationMonitorDetectResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationMonitorDetectResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListApplicationMonitorDetectResultRequest) SetBeginTime(v int64) *ListApplicationMonitorDetectResultRequest {
	s.BeginTime = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetEndTime(v int64) *ListApplicationMonitorDetectResultRequest {
	s.EndTime = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetPageNumber(v int32) *ListApplicationMonitorDetectResultRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetPageSize(v int32) *ListApplicationMonitorDetectResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetRegionId(v string) *ListApplicationMonitorDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetTaskId(v string) *ListApplicationMonitorDetectResultRequest {
	s.TaskId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
