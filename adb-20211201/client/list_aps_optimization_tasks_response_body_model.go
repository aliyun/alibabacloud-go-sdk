// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsOptimizationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApsOptimizationTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApsOptimizationTasksResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListApsOptimizationTasksResponseBodyItems) *ListApsOptimizationTasksResponseBody
	GetItems() []*ListApsOptimizationTasksResponseBodyItems
	SetMessage(v string) *ListApsOptimizationTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListApsOptimizationTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApsOptimizationTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApsOptimizationTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApsOptimizationTasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListApsOptimizationTasksResponseBody
	GetTotalCount() *int64
}

type ListApsOptimizationTasksResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The queried optimization jobs.
	//
	// example:
	//
	// -
	Items []*ListApsOptimizationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApsOptimizationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApsOptimizationTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApsOptimizationTasksResponseBody) GetItems() []*ListApsOptimizationTasksResponseBodyItems {
	return s.Items
}

func (s *ListApsOptimizationTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApsOptimizationTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApsOptimizationTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApsOptimizationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApsOptimizationTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApsOptimizationTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApsOptimizationTasksResponseBody) SetCode(v string) *ListApsOptimizationTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetHttpStatusCode(v int32) *ListApsOptimizationTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetItems(v []*ListApsOptimizationTasksResponseBodyItems) *ListApsOptimizationTasksResponseBody {
	s.Items = v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetMessage(v string) *ListApsOptimizationTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetPageNumber(v int32) *ListApsOptimizationTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetPageSize(v int32) *ListApsOptimizationTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetRequestId(v string) *ListApsOptimizationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetSuccess(v bool) *ListApsOptimizationTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) SetTotalCount(v int64) *ListApsOptimizationTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApsOptimizationTasksResponseBodyItems struct {
	// The computing resources used by the optimization job.
	//
	// example:
	//
	// 2
	ComputeUnit *string `json:"ComputeUnit,omitempty" xml:"ComputeUnit,omitempty"`
	// The time when the optimization job was created.
	//
	// example:
	//
	// 2022-01-23T02:18Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time when the optimization job was modified.
	//
	// example:
	//
	// 2022-09-30T00:15Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The type of the lifecycle management policy.
	//
	// example:
	//
	// StrategyValue
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// The description of the optimization job.
	//
	// example:
	//
	// test
	TaskDesc *string `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	// The execution duration of the optimization job.
	//
	// example:
	//
	// 1000
	TaskDuration *int64 `json:"TaskDuration,omitempty" xml:"TaskDuration,omitempty"`
	// The job ID.
	//
	// example:
	//
	// sj-hz******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The error message.
	//
	// example:
	//
	// -
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// The execution status. Valid values:
	//
	// 1.  NEW
	//
	// 2.  RUNNING
	//
	// 3.  SUCCESS
	//
	// 4.  STOPPED
	//
	// 5.  FAILED
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListApsOptimizationTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetComputeUnit() *string {
	return s.ComputeUnit
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetTaskDesc() *string {
	return s.TaskDesc
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetTaskDuration() *int64 {
	return s.TaskDuration
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetTaskMessage() *string {
	return s.TaskMessage
}

func (s *ListApsOptimizationTasksResponseBodyItems) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetComputeUnit(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.ComputeUnit = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetCreatedTime(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetDBClusterId(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetModifiedTime(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetStrategyType(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.StrategyType = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetTaskDesc(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.TaskDesc = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetTaskDuration(v int64) *ListApsOptimizationTasksResponseBodyItems {
	s.TaskDuration = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetTaskId(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetTaskMessage(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.TaskMessage = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) SetTaskStatus(v string) *ListApsOptimizationTasksResponseBodyItems {
	s.TaskStatus = &v
	return s
}

func (s *ListApsOptimizationTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
