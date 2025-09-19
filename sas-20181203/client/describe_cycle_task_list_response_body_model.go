// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCycleTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCycleScheduleResponseList(v []*DescribeCycleTaskListResponseBodyCycleScheduleResponseList) *DescribeCycleTaskListResponseBody
	GetCycleScheduleResponseList() []*DescribeCycleTaskListResponseBodyCycleScheduleResponseList
	SetPageInfo(v *DescribeCycleTaskListResponseBodyPageInfo) *DescribeCycleTaskListResponseBody
	GetPageInfo() *DescribeCycleTaskListResponseBodyPageInfo
	SetRequestId(v string) *DescribeCycleTaskListResponseBody
	GetRequestId() *string
}

type DescribeCycleTaskListResponseBody struct {
	// An array that consists of periodic scan tasks.
	CycleScheduleResponseList []*DescribeCycleTaskListResponseBodyCycleScheduleResponseList `json:"CycleScheduleResponseList,omitempty" xml:"CycleScheduleResponseList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeCycleTaskListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCycleTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCycleTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCycleTaskListResponseBody) GetCycleScheduleResponseList() []*DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	return s.CycleScheduleResponseList
}

func (s *DescribeCycleTaskListResponseBody) GetPageInfo() *DescribeCycleTaskListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeCycleTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCycleTaskListResponseBody) SetCycleScheduleResponseList(v []*DescribeCycleTaskListResponseBodyCycleScheduleResponseList) *DescribeCycleTaskListResponseBody {
	s.CycleScheduleResponseList = v
	return s
}

func (s *DescribeCycleTaskListResponseBody) SetPageInfo(v *DescribeCycleTaskListResponseBodyPageInfo) *DescribeCycleTaskListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCycleTaskListResponseBody) SetRequestId(v string) *DescribeCycleTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCycleTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCycleTaskListResponseBodyCycleScheduleResponseList struct {
	// The configuration ID.
	//
	// example:
	//
	// 2cdbdeba8dd70586d5814d4cbf21****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Indicates whether the configuration for the task interval was enabled. Valid values:
	//
	// 	- **1**: enabled.
	//
	// 	- **0**: disabled.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the task first started.
	//
	// example:
	//
	// 1667491200000
	FirstDateStr *int64 `json:"FirstDateStr,omitempty" xml:"FirstDateStr,omitempty"`
	// The interval between which two consecutive tasks are run.
	//
	// example:
	//
	// 3
	IntervalPeriod *int32 `json:"IntervalPeriod,omitempty" xml:"IntervalPeriod,omitempty"`
	// The ID of the last task.
	//
	// example:
	//
	// 38730bb078f4a1461d4ed283994c****
	LastTaskId *string `json:"LastTaskId,omitempty" xml:"LastTaskId,omitempty"`
	// The time when the next task starts. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1671184531000
	NextStartTimeStr *int64 `json:"NextStartTimeStr,omitempty" xml:"NextStartTimeStr,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {"userAgreement":"yes","lang":"zh"}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The unit of the scan interval. Valid values:
	//
	// 	- **day**
	//
	// 	- **hour**
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The end time of the task. The time must be a time frame.
	//
	// example:
	//
	// 7
	TargetEndTime *int32 `json:"TargetEndTime,omitempty" xml:"TargetEndTime,omitempty"`
	// The start time of the task. The start time must be a time frame.
	//
	// example:
	//
	// 0
	TargetStartTime *int32 `json:"TargetStartTime,omitempty" xml:"TargetStartTime,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeCycleTaskListResponseBodyCycleScheduleResponseList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GoString() string {
	return s.String()
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetEnable() *int32 {
	return s.Enable
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetFirstDateStr() *int64 {
	return s.FirstDateStr
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetIntervalPeriod() *int32 {
	return s.IntervalPeriod
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetLastTaskId() *string {
	return s.LastTaskId
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetNextStartTimeStr() *int64 {
	return s.NextStartTimeStr
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetParam() *string {
	return s.Param
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetTargetEndTime() *int32 {
	return s.TargetEndTime
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetTargetStartTime() *int32 {
	return s.TargetStartTime
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetConfigId(v string) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.ConfigId = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetEnable(v int32) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.Enable = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetFirstDateStr(v int64) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.FirstDateStr = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetIntervalPeriod(v int32) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.IntervalPeriod = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetLastTaskId(v string) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.LastTaskId = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetNextStartTimeStr(v int64) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.NextStartTimeStr = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetParam(v string) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.Param = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetPeriodUnit(v string) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetTargetEndTime(v int32) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.TargetEndTime = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetTargetStartTime(v int32) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.TargetStartTime = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetTaskName(v string) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.TaskName = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) SetTaskType(v string) *DescribeCycleTaskListResponseBodyCycleScheduleResponseList {
	s.TaskType = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyCycleScheduleResponseList) Validate() error {
	return dara.Validate(s)
}

type DescribeCycleTaskListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCycleTaskListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCycleTaskListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) SetCount(v int32) *DescribeCycleTaskListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeCycleTaskListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) SetPageSize(v int32) *DescribeCycleTaskListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeCycleTaskListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCycleTaskListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
