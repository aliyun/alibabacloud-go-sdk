// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingSkillJobPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMetaProfilingSkillJobPageResponseBodyData) *ListMetaProfilingSkillJobPageResponseBody
	GetData() []*ListMetaProfilingSkillJobPageResponseBodyData
	SetErrCode(v string) *ListMetaProfilingSkillJobPageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMetaProfilingSkillJobPageResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListMetaProfilingSkillJobPageResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingSkillJobPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMetaProfilingSkillJobPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaProfilingSkillJobPageResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMetaProfilingSkillJobPageResponseBody
	GetTotalCount() *int32
}

type ListMetaProfilingSkillJobPageResponseBody struct {
	// The list of profiling tasks. For the element structure, see the descriptions of the child fields.
	Data []*ListMetaProfilingSkillJobPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// The parameter is invalid or missing.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of records to return per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID used to locate and troubleshoot issues for this call.
	//
	// example:
	//
	// 3F45B8C1-2A6E-4C9B-9E1F-7D2A1B0C5E66
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success, and a value of false indicates failure. If the call fails, troubleshoot the issue based on the errCode and errMessage parameters.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that meet the specified conditions.
	//
	// example:
	//
	// 128
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMetaProfilingSkillJobPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingSkillJobPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetData() []*ListMetaProfilingSkillJobPageResponseBodyData {
	return s.Data
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaProfilingSkillJobPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetData(v []*ListMetaProfilingSkillJobPageResponseBodyData) *ListMetaProfilingSkillJobPageResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetErrCode(v string) *ListMetaProfilingSkillJobPageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetErrMessage(v string) *ListMetaProfilingSkillJobPageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetPageIndex(v int32) *ListMetaProfilingSkillJobPageResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetPageSize(v int32) *ListMetaProfilingSkillJobPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetRequestId(v string) *ListMetaProfilingSkillJobPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetSuccess(v bool) *ListMetaProfilingSkillJobPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) SetTotalCount(v int32) *ListMetaProfilingSkillJobPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBody) Validate() error {
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

type ListMetaProfilingSkillJobPageResponseBodyData struct {
	// The component ID, which is the primary key identifier of the data source component.
	//
	// example:
	//
	// 2001
	ComponentId *int64 `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// The identifier of the current execution status of the task.
	//
	// example:
	//
	// 2
	CurrentTaskStatus *int32 `json:"currentTaskStatus,omitempty" xml:"currentTaskStatus,omitempty"`
	// The data component name.
	//
	// example:
	//
	// hive_cluster_a
	DcName *string `json:"dcName,omitempty" xml:"dcName,omitempty"`
	// The status identifier of the data component.
	//
	// example:
	//
	// 1
	DcStatus *int32 `json:"dcStatus,omitempty" xml:"dcStatus,omitempty"`
	// The type identifier of the data component.
	//
	// example:
	//
	// HIVE
	DcType *string `json:"dcType,omitempty" xml:"dcType,omitempty"`
	// The number of differing nodes.
	//
	// example:
	//
	// 3
	DiffNodeNum *int32 `json:"diffNodeNum,omitempty" xml:"diffNodeNum,omitempty"`
	// The number of differing partitions.
	//
	// example:
	//
	// 15
	DiffPartitionNum *int32 `json:"diffPartitionNum,omitempty" xml:"diffPartitionNum,omitempty"`
	// The number of differing tables.
	//
	// example:
	//
	// 12
	DiffTableNum *int32 `json:"diffTableNum,omitempty" xml:"diffTableNum,omitempty"`
	// The number of differing workflows.
	//
	// example:
	//
	// 2
	DiffWorkflowNum *int32 `json:"diffWorkflowNum,omitempty" xml:"diffWorkflowNum,omitempty"`
	// The category type identifier of the data source.
	//
	// example:
	//
	// SOURCE
	DsCategoryType *string `json:"dsCategoryType,omitempty" xml:"dsCategoryType,omitempty"`
	// The time when the record was created.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Indicates whether incremental changes exist.
	HasIncrementChanges *bool `json:"hasIncrementChanges,omitempty" xml:"hasIncrementChanges,omitempty"`
	// The primary key ID of the object.
	//
	// example:
	//
	// 7001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The description of the profiling task.
	//
	// example:
	//
	// Daily full metadata profiling task
	JobDesc *string `json:"jobDesc,omitempty" xml:"jobDesc,omitempty"`
	// The profiling task name.
	//
	// example:
	//
	// profiling_daily
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The status identifier of the profiling task.
	//
	// example:
	//
	// 1
	JobStatus *int32 `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// The ID of the most recent execution batch.
	//
	// example:
	//
	// 20240719220000
	LastBatchId *string `json:"lastBatchId,omitempty" xml:"lastBatchId,omitempty"`
	// The ID of the most recently executed task.
	//
	// example:
	//
	// 8001
	LastTaskId *int64 `json:"lastTaskId,omitempty" xml:"lastTaskId,omitempty"`
	// The estimated time for the next scheduled execution.
	//
	// example:
	//
	// 2024-07-20 02:00:00
	NextCronTime *string `json:"nextCronTime,omitempty" xml:"nextCronTime,omitempty"`
	// Indicates whether data profiling is enabled.
	//
	// example:
	//
	// 1
	ProfilingEnable *int32 `json:"profilingEnable,omitempty" xml:"profilingEnable,omitempty"`
	// The list of schemas (databases) to be profiled.
	ProfilingSchemaList []*string `json:"profilingSchemaList,omitempty" xml:"profilingSchemaList,omitempty" type:"Repeated"`
	// The identifier of the task execution result.
	//
	// example:
	//
	// 1
	TaskExecResult *int32 `json:"taskExecResult,omitempty" xml:"taskExecResult,omitempty"`
	// The time when the task execution was completed.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	TaskFinishTime *string `json:"taskFinishTime,omitempty" xml:"taskFinishTime,omitempty"`
	// The type identifier of the task.
	//
	// example:
	//
	// PROFILING
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s ListMetaProfilingSkillJobPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingSkillJobPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetComponentId() *int64 {
	return s.ComponentId
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetCurrentTaskStatus() *int32 {
	return s.CurrentTaskStatus
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDcName() *string {
	return s.DcName
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDcStatus() *int32 {
	return s.DcStatus
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDcType() *string {
	return s.DcType
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDiffNodeNum() *int32 {
	return s.DiffNodeNum
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDiffPartitionNum() *int32 {
	return s.DiffPartitionNum
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDiffTableNum() *int32 {
	return s.DiffTableNum
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDiffWorkflowNum() *int32 {
	return s.DiffWorkflowNum
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetDsCategoryType() *string {
	return s.DsCategoryType
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetHasIncrementChanges() *bool {
	return s.HasIncrementChanges
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetJobDesc() *string {
	return s.JobDesc
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetLastBatchId() *string {
	return s.LastBatchId
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetLastTaskId() *int64 {
	return s.LastTaskId
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetNextCronTime() *string {
	return s.NextCronTime
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetProfilingEnable() *int32 {
	return s.ProfilingEnable
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetProfilingSchemaList() []*string {
	return s.ProfilingSchemaList
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetTaskExecResult() *int32 {
	return s.TaskExecResult
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetTaskFinishTime() *string {
	return s.TaskFinishTime
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetComponentId(v int64) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.ComponentId = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetCurrentTaskStatus(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.CurrentTaskStatus = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDcName(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DcName = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDcStatus(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DcStatus = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDcType(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DcType = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDiffNodeNum(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DiffNodeNum = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDiffPartitionNum(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DiffPartitionNum = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDiffTableNum(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DiffTableNum = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDiffWorkflowNum(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DiffWorkflowNum = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetDsCategoryType(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.DsCategoryType = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetGmtCreate(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetHasIncrementChanges(v bool) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.HasIncrementChanges = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetId(v int64) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetJobDesc(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.JobDesc = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetJobName(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.JobName = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetJobStatus(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetLastBatchId(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.LastBatchId = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetLastTaskId(v int64) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.LastTaskId = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetNextCronTime(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.NextCronTime = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetProfilingEnable(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.ProfilingEnable = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetProfilingSchemaList(v []*string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.ProfilingSchemaList = v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetTaskExecResult(v int32) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.TaskExecResult = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetTaskFinishTime(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.TaskFinishTime = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) SetTaskType(v string) *ListMetaProfilingSkillJobPageResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListMetaProfilingSkillJobPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
