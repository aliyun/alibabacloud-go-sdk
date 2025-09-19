// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeOnceTaskResponseBodyPageInfo) *DescribeOnceTaskResponseBody
	GetPageInfo() *DescribeOnceTaskResponseBodyPageInfo
	SetRequestId(v string) *DescribeOnceTaskResponseBody
	GetRequestId() *string
	SetTaskManageResponseList(v []*DescribeOnceTaskResponseBodyTaskManageResponseList) *DescribeOnceTaskResponseBody
	GetTaskManageResponseList() []*DescribeOnceTaskResponseBodyTaskManageResponseList
}

type DescribeOnceTaskResponseBody struct {
	// The pagination information.
	PageInfo *DescribeOnceTaskResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the tasks.
	TaskManageResponseList []*DescribeOnceTaskResponseBodyTaskManageResponseList `json:"TaskManageResponseList,omitempty" xml:"TaskManageResponseList,omitempty" type:"Repeated"`
}

func (s DescribeOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskResponseBody) GetPageInfo() *DescribeOnceTaskResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOnceTaskResponseBody) GetTaskManageResponseList() []*DescribeOnceTaskResponseBodyTaskManageResponseList {
	return s.TaskManageResponseList
}

func (s *DescribeOnceTaskResponseBody) SetPageInfo(v *DescribeOnceTaskResponseBodyPageInfo) *DescribeOnceTaskResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeOnceTaskResponseBody) SetRequestId(v string) *DescribeOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnceTaskResponseBody) SetTaskManageResponseList(v []*DescribeOnceTaskResponseBodyTaskManageResponseList) *DescribeOnceTaskResponseBody {
	s.TaskManageResponseList = v
	return s
}

func (s *DescribeOnceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOnceTaskResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOnceTaskResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeOnceTaskResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOnceTaskResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOnceTaskResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOnceTaskResponseBodyPageInfo) SetCount(v int32) *DescribeOnceTaskResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeOnceTaskResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyPageInfo) SetPageSize(v int32) *DescribeOnceTaskResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyPageInfo) SetTotalCount(v int32) *DescribeOnceTaskResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeOnceTaskResponseBodyTaskManageResponseList struct {
	// The execution details of the task. The value of this parameter is in the JSON format.
	//
	// 	- **causeCode**: the returned code for the cause.
	//
	// 	- **causeMsg**: the returned message for the cause.
	//
	// 	- **resCode**: the returned code for troubleshooting.
	//
	// 	- **resMsg**: the returned message for troubleshooting.
	//
	// 	- **problemType**: the type of the issue.
	//
	// 	- **dispatchType**: the task delivery method.
	//
	// 	- **uuid**: the UUID of the server.
	//
	// 	- **instanceId**: the instance ID of the server.
	//
	// 	- **internetIp**: the public IP address of the server.
	//
	// 	- **intranetIp**: the private IP address of the server.
	//
	// 	- **instanceName**: the instance name of the server.
	//
	// 	- **url**: the download URL of the troubleshooting log.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "dispatchType": "manual",
	//
	//             "causeMsg": [],
	//
	//             "causeCode": [],
	//
	//             "resCode": [
	//
	//                   "1003"
	//
	//             ],
	//
	//             "resMsg": [
	//
	//                   "powershell -executionpolicy bypass -c \\"(New-Object Net.WebClient).DownloadFile(\\"http://aegis.alicdn.com/download/aegis_client_self_check/win32/aegis_checker.exe\\", $ExecutionContext.SessionState.Path.GetUnresolvedProviderPathFromPSPath(\\".\\\\\\\\aegis_checker.exe\\"))\\"; \\"./aegis_checker.exe -b eyJtb2RlIjoxLCJqc3J2X2RvbWFpbiI6W10sImlzc3VlIjoib2ZmbGluZSIsInVwZGF0ZV9kb21haW4iOltdLCJ1dWlkIjoiaW5ldC1lYWUwNDg2Ny0wMDJmLTQyM2QtYWYwMC1jNzJjZDYyOWIyNDgiLCJjbWRfaWR4IjoiNDRjZThiZWI3ZGYyYTQxMjQ1NGM4ZDc5OTE1ODI1MzMifQ==\\""
	//
	//             ],
	//
	//             "problemType": "offline",
	//
	//             "uuid": "inet-eae04867-002f-423d-af00-c72cd629****"
	//
	//       }
	//
	// ]
	DetailData *string `json:"DetailData,omitempty" xml:"DetailData,omitempty"`
	// The number of tasks that fail to be executed.
	//
	// example:
	//
	// 2
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The progress of the task. Unit: percent (%).
	//
	// example:
	//
	// 10%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution result of the task.
	//
	// example:
	//
	// successful
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The number of tasks that are executed.
	//
	// example:
	//
	// 7
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The timestamp that indicates the time when the task ends. Unit: milliseconds.
	//
	// example:
	//
	// 1650267989000
	TaskEndTime *int64 `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// e900f528f5a6229bb640ca27cb44c98e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The timestamp that indicates the time when the task starts. Unit: milliseconds.
	//
	// example:
	//
	// 1649732012000
	TaskStartTime *int64 `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **1**: The task is started.
	//
	// 	- **2**: The task is complete.
	//
	// 	- **3**: The task fails.
	//
	// 	- **4**: The task times out.
	//
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The text description of the status for the task. Valid values:
	//
	// 	- **INIT**: The task is pending start.
	//
	// 	- **START**: The task is started.
	//
	// 	- **DISPATCH**: The self-check command is issued.
	//
	// 	- **SUCCESS**: The self-check is complete.
	//
	// 	- **FAIL**: The task fails.
	//
	// 	- **TIMEOUT**: The task times out.
	//
	// example:
	//
	// INIT
	TaskStatusText *string `json:"TaskStatusText,omitempty" xml:"TaskStatusText,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **CLIENT_PROBLEM_CHECK**: a task of the Security Center client
	//
	// 	- **CLIENT_DEV_OPS**: an O\\&M task of Cloud Assistant
	//
	// 	- **ASSET_SECURITY_CHECK**: a task for asset information collection
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeOnceTaskResponseBodyTaskManageResponseList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskResponseBodyTaskManageResponseList) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetDetailData() *string {
	return s.DetailData
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetFailCount() *int32 {
	return s.FailCount
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetProgress() *string {
	return s.Progress
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskEndTime() *int64 {
	return s.TaskEndTime
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskStartTime() *int64 {
	return s.TaskStartTime
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskStatusText() *string {
	return s.TaskStatusText
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetDetailData(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.DetailData = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetFailCount(v int32) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.FailCount = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetProgress(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.Progress = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetResultInfo(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.ResultInfo = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetSuccessCount(v int32) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.SuccessCount = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskEndTime(v int64) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskEndTime = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskId(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskId = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskName(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskName = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskStartTime(v int64) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskStartTime = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskStatus(v int32) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskStatus = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskStatusText(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskStatusText = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) SetTaskType(v string) *DescribeOnceTaskResponseBodyTaskManageResponseList {
	s.TaskType = &v
	return s
}

func (s *DescribeOnceTaskResponseBodyTaskManageResponseList) Validate() error {
	return dara.Validate(s)
}
