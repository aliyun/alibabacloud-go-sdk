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
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of client task query results.
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
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TaskManageResponseList != nil {
		for _, item := range s.TaskManageResponseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOnceTaskResponseBodyPageInfo struct {
	// The number of client tasks displayed on the current page in a paged query.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of client tasks per page in a paged query. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of client tasks returned.
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
	// The task execution details. This parameter is in JSON format.
	//
	// - **causeCode**: the return code of the troubleshooting cause.
	//
	// - **causeMsg**: the return message of the troubleshooting cause.
	//
	// - **resCode**: the troubleshooting return code.
	//
	// - **resMsg**: the troubleshooting return message.
	//
	// - **problemType**: the problem type.
	//
	// - **dispatchType**: the task dispatch method.
	//
	// - **uuid**: the server UUID.
	//
	// - **instanceId**: the ID of the server instance.
	//
	// - **internetIp**: the public IP address of the server.
	//
	// - **intranetIp**: the private IP address of the server.
	//
	// - **instanceName**: the name of the server instance.
	//
	// - **url**: the download URL of the troubleshooting log.
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
	// The number of tasks that failed to be executed.
	//
	// example:
	//
	// 2
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The task progress, in percentage.
	//
	// example:
	//
	// 10%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The task execution result.
	//
	// example:
	//
	// successful
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The number of tasks that are executed successfully.
	//
	// example:
	//
	// 7
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The timestamp when the task actually ends. Unit: milliseconds.
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
	// The task name.
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The timestamp when the task actually starts. Unit: milliseconds.
	//
	// example:
	//
	// 1649732012000
	TaskStartTime *int64 `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// The task status. Valid values:
	//
	// - **1**: Started.
	//
	// - **2**: Completed.
	//
	// - **3**: Failed.
	//
	// - **4**: Timed out.
	//
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The text representation of the task status. Valid values:
	//
	// - **INIT**: Pending.
	//
	// - **START**: Started.
	//
	// - **DISPATCH**: Self-check command dispatched.
	//
	// - **SUCCESS**: Self-check completed.
	//
	// - **FAIL**: Execution failed.
	//
	// - **TIMEOUT**: Timed out.
	//
	// example:
	//
	// INIT
	TaskStatusText *string `json:"TaskStatusText,omitempty" xml:"TaskStatusText,omitempty"`
	// The task type. Valid values:
	//
	// - **CLIENT_PROBLEM_CHECK**: client task
	//
	// - **CLIENT_DEV_OPS**: cloud O&M task
	//
	// - **ASSET_SECURITY_CHECK**: asset information collection task.
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
