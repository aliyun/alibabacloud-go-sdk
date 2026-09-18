// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanTasksByTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListScanTasksByTargetResponseBodyData) *ListScanTasksByTargetResponseBody
	GetData() []*ListScanTasksByTargetResponseBodyData
	SetPageNumber(v int64) *ListScanTasksByTargetResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListScanTasksByTargetResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListScanTasksByTargetResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListScanTasksByTargetResponseBody
	GetTotalCount() *int64
}

type ListScanTasksByTargetResponseBody struct {
	// The list of scan tasks on the current page.
	Data []*ListScanTasksByTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The normalized page number that actually takes effect. This value may differ from the input parameter.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The normalized number of entries per page that actually takes effect. This value may differ from the input parameter.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request, used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of scan tasks that match the filter conditions within the last 366-day window.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScanTasksByTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScanTasksByTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListScanTasksByTargetResponseBody) GetData() []*ListScanTasksByTargetResponseBodyData {
	return s.Data
}

func (s *ListScanTasksByTargetResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListScanTasksByTargetResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScanTasksByTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScanTasksByTargetResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListScanTasksByTargetResponseBody) SetData(v []*ListScanTasksByTargetResponseBodyData) *ListScanTasksByTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListScanTasksByTargetResponseBody) SetPageNumber(v int64) *ListScanTasksByTargetResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListScanTasksByTargetResponseBody) SetPageSize(v int64) *ListScanTasksByTargetResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScanTasksByTargetResponseBody) SetRequestId(v string) *ListScanTasksByTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScanTasksByTargetResponseBody) SetTotalCount(v int64) *ListScanTasksByTargetResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScanTasksByTargetResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScanTasksByTargetResponseBodyData struct {
	// The task creation time, in milliseconds (Unix epoch milliseconds).
	//
	// example:
	//
	// 1735689600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The task end time, in milliseconds (Unix epoch milliseconds). This value is null if the task has not ended.
	//
	// example:
	//
	// 1735689600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of samples that the task has executed.
	//
	// example:
	//
	// 80
	ExecuteCaseCount *int64 `json:"ExecuteCaseCount,omitempty" xml:"ExecuteCaseCount,omitempty"`
	// The risk level of the task result. This value is null if the task is not completed or no risk assessment has been generated.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The detection intensity of the task.
	//
	// example:
	//
	// 2
	SampleLevel *int64 `json:"SampleLevel,omitempty" xml:"SampleLevel,omitempty"`
	// The scan type of the task. Historical tasks without a recorded scan type are normalized to attack.
	//
	// example:
	//
	// attack
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The unique identifier of the scan task. You can use this ID for result download and status tracking.
	//
	// example:
	//
	// task-abc123def4567
	ScannerTaskId *string `json:"ScannerTaskId,omitempty" xml:"ScannerTaskId,omitempty"`
	// The task message. This value contains the failure reason if the task failed, or is empty if the task succeeded or no message is available.
	//
	// example:
	//
	// Execution timed out
	ScannerTaskMessage *string `json:"ScannerTaskMessage,omitempty" xml:"ScannerTaskMessage,omitempty"`
	// The current status of the task.
	//
	// example:
	//
	// completed
	ScannerTaskStatus *string `json:"ScannerTaskStatus,omitempty" xml:"ScannerTaskStatus,omitempty"`
	// The task start time, in milliseconds (Unix epoch milliseconds). This value is null if the task has not started.
	//
	// example:
	//
	// 1735689600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task name. If no name is specified during creation, the default value is "Target Scan - target name".
	//
	// example:
	//
	// Target Scan - My Bailian Target
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The total number of samples that the task plans to execute.
	//
	// example:
	//
	// 120
	TotalCaseCount *int64 `json:"TotalCaseCount,omitempty" xml:"TotalCaseCount,omitempty"`
}

func (s ListScanTasksByTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScanTasksByTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScanTasksByTargetResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListScanTasksByTargetResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListScanTasksByTargetResponseBodyData) GetExecuteCaseCount() *int64 {
	return s.ExecuteCaseCount
}

func (s *ListScanTasksByTargetResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListScanTasksByTargetResponseBodyData) GetSampleLevel() *int64 {
	return s.SampleLevel
}

func (s *ListScanTasksByTargetResponseBodyData) GetScanType() *string {
	return s.ScanType
}

func (s *ListScanTasksByTargetResponseBodyData) GetScannerTaskId() *string {
	return s.ScannerTaskId
}

func (s *ListScanTasksByTargetResponseBodyData) GetScannerTaskMessage() *string {
	return s.ScannerTaskMessage
}

func (s *ListScanTasksByTargetResponseBodyData) GetScannerTaskStatus() *string {
	return s.ScannerTaskStatus
}

func (s *ListScanTasksByTargetResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListScanTasksByTargetResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListScanTasksByTargetResponseBodyData) GetTotalCaseCount() *int64 {
	return s.TotalCaseCount
}

func (s *ListScanTasksByTargetResponseBodyData) SetCreateTime(v int64) *ListScanTasksByTargetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetEndTime(v int64) *ListScanTasksByTargetResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetExecuteCaseCount(v int64) *ListScanTasksByTargetResponseBodyData {
	s.ExecuteCaseCount = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetRiskLevel(v string) *ListScanTasksByTargetResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetSampleLevel(v int64) *ListScanTasksByTargetResponseBodyData {
	s.SampleLevel = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetScanType(v string) *ListScanTasksByTargetResponseBodyData {
	s.ScanType = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetScannerTaskId(v string) *ListScanTasksByTargetResponseBodyData {
	s.ScannerTaskId = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetScannerTaskMessage(v string) *ListScanTasksByTargetResponseBodyData {
	s.ScannerTaskMessage = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetScannerTaskStatus(v string) *ListScanTasksByTargetResponseBodyData {
	s.ScannerTaskStatus = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetStartTime(v int64) *ListScanTasksByTargetResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetTaskName(v string) *ListScanTasksByTargetResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) SetTotalCaseCount(v int64) *ListScanTasksByTargetResponseBodyData {
	s.TotalCaseCount = &v
	return s
}

func (s *ListScanTasksByTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
