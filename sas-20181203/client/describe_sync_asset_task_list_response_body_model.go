// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAssetTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeSyncAssetTaskListResponseBodyPageInfo) *DescribeSyncAssetTaskListResponseBody
	GetPageInfo() *DescribeSyncAssetTaskListResponseBodyPageInfo
	SetRequestId(v string) *DescribeSyncAssetTaskListResponseBody
	GetRequestId() *string
	SetTaskRecords(v []*DescribeSyncAssetTaskListResponseBodyTaskRecords) *DescribeSyncAssetTaskListResponseBody
	GetTaskRecords() []*DescribeSyncAssetTaskListResponseBodyTaskRecords
}

type DescribeSyncAssetTaskListResponseBody struct {
	// The pagination information.
	PageInfo *DescribeSyncAssetTaskListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDC scan tasks.
	TaskRecords []*DescribeSyncAssetTaskListResponseBodyTaskRecords `json:"TaskRecords,omitempty" xml:"TaskRecords,omitempty" type:"Repeated"`
}

func (s DescribeSyncAssetTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskListResponseBody) GetPageInfo() *DescribeSyncAssetTaskListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeSyncAssetTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncAssetTaskListResponseBody) GetTaskRecords() []*DescribeSyncAssetTaskListResponseBodyTaskRecords {
	return s.TaskRecords
}

func (s *DescribeSyncAssetTaskListResponseBody) SetPageInfo(v *DescribeSyncAssetTaskListResponseBodyPageInfo) *DescribeSyncAssetTaskListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBody) SetRequestId(v string) *DescribeSyncAssetTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBody) SetTaskRecords(v []*DescribeSyncAssetTaskListResponseBodyTaskRecords) *DescribeSyncAssetTaskListResponseBody {
	s.TaskRecords = v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSyncAssetTaskListResponseBodyPageInfo struct {
	// The number of IDC scan tasks on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of IDC scan tasks per page. Default value: 20. If you leave this parameter empty, 20 IDC scan tasks are returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of IDC scan tasks returned.
	//
	// example:
	//
	// 110
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSyncAssetTaskListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) SetCount(v int32) *DescribeSyncAssetTaskListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeSyncAssetTaskListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) SetPageSize(v int32) *DescribeSyncAssetTaskListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeSyncAssetTaskListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeSyncAssetTaskListResponseBodyTaskRecords struct {
	// The number of assets that are detected by the task.
	//
	// example:
	//
	// 100
	AssetCount *int32 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The CIDR blocks that are used for scanning. Multiple CIDR blocks are separated by commas (,).
	//
	// example:
	//
	// 1.1.1.1/24,1.1.1.1/24
	IpSegments *string `json:"IpSegments,omitempty" xml:"IpSegments,omitempty"`
	// The progress of the task, in percentage.
	//
	// example:
	//
	// 100
	ProcessRate *int32 `json:"ProcessRate,omitempty" xml:"ProcessRate,omitempty"`
	// The ID of the root task.
	//
	// example:
	//
	// 73c392f9c505129a257472a3f911d65d
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// The timestamp when the task ended.
	//
	// example:
	//
	// 1653965680000
	TaskEndTime *int64 `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// IDC_PROBE_SCAN-1.1.1..124-lse_ubuntu_test1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The timestamp when the task started. Unit: milliseconds.
	//
	// example:
	//
	// 1633746651715
	TaskStartTime *int64 `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// The status of the IDC scan task. Valid Values:
	//
	// 	- **INIT**: The task is not started.
	//
	// 	- **START**: The task is started.
	//
	// 	- **MESSAGE_SEND**: The command is sent.
	//
	// 	- **SUCCESS**: The task is complete.
	//
	// 	- **FAIL**: The task failed.
	//
	// 	- **TIMEOUT**: The task timed out.
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. The value is fixed as **IDC_PROBE_SCAN**, which indicates an IDC scan task.
	//
	// example:
	//
	// IDC_PROBE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeSyncAssetTaskListResponseBodyTaskRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskListResponseBodyTaskRecords) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetAssetCount() *int32 {
	return s.AssetCount
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetIpSegments() *string {
	return s.IpSegments
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetProcessRate() *int32 {
	return s.ProcessRate
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetTaskEndTime() *int64 {
	return s.TaskEndTime
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetTaskStartTime() *int64 {
	return s.TaskStartTime
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetAssetCount(v int32) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.AssetCount = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetIpSegments(v string) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.IpSegments = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetProcessRate(v int32) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.ProcessRate = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetRootTaskId(v string) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.RootTaskId = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetTaskEndTime(v int64) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.TaskEndTime = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetTaskName(v string) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.TaskName = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetTaskStartTime(v int64) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.TaskStartTime = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetTaskStatus(v string) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) SetTaskType(v string) *DescribeSyncAssetTaskListResponseBodyTaskRecords {
	s.TaskType = &v
	return s
}

func (s *DescribeSyncAssetTaskListResponseBodyTaskRecords) Validate() error {
	return dara.Validate(s)
}
