// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoOpsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAutoOpsTaskResponseBody
	GetRequestId() *string
	SetTask(v *GetAutoOpsTaskResponseBodyTask) *GetAutoOpsTaskResponseBody
	GetTask() *GetAutoOpsTaskResponseBodyTask
}

type GetAutoOpsTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the O&M task.
	Task *GetAutoOpsTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetAutoOpsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoOpsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoOpsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoOpsTaskResponseBody) GetTask() *GetAutoOpsTaskResponseBodyTask {
	return s.Task
}

func (s *GetAutoOpsTaskResponseBody) SetRequestId(v string) *GetAutoOpsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoOpsTaskResponseBody) SetTask(v *GetAutoOpsTaskResponseBodyTask) *GetAutoOpsTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetAutoOpsTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoOpsTaskResponseBodyTask struct {
	// The time when the approval of the O&M task was completed. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1669965908
	AllowedOverTime *int64 `json:"AllowedOverTime,omitempty" xml:"AllowedOverTime,omitempty"`
	// The remarks of the O&M task.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the O&M task.
	//
	// example:
	//
	// taskname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution plan of the O&M task.
	//
	// - **ExecAt**: If the value of ScheduleType is Manual, this parameter is not meaningful. If the value of ScheduleType is FixTime, this parameter indicates the scheduled execution time in seconds as a UNIX timestamp. If the value of ScheduleType is CycleInterval, this parameter indicates the first execution time in seconds as a UNIX timestamp.
	//
	// - **PeriodNum**: If the value of ScheduleType is Manual or FixTime, this parameter is not meaningful. If the value of ScheduleType is CycleInterval, this parameter indicates the interval for periodic execution.
	//
	// - **PeriodUnit**: If the value of ScheduleType is Manual or FixTime, this parameter is not meaningful. If the value of ScheduleType is CycleInterval, this parameter indicates the unit of the periodic execution interval. Valid values: hour and day.
	//
	// example:
	//
	// {"ExecAt":0,"PeriodNum":0,"PeriodUnit":""}
	ScheduleTimeInfo *string `json:"ScheduleTimeInfo,omitempty" xml:"ScheduleTimeInfo,omitempty"`
	// The scheduling type of the task.
	//
	// - **FixTime**: scheduled execution.
	//
	// - **CycleInterval**: periodic execution.
	//
	// - **Manual**: manual execution triggered by the user.
	//
	// example:
	//
	// FixTime
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The content of the script to be executed by the O&M task. The value is Base64-encoded.
	//
	// example:
	//
	// bHM=
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The ID of the script associated with the O&M task. This parameter is returned only when ScriptType is set to SpecificScript.
	//
	// example:
	//
	// 2
	ScriptId *int64 `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the script associated with the O&M task.
	//
	// example:
	//
	// name
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The script type of the O&M task.
	//
	// - **HandInput**: manually entered script.
	//
	// - **SpecificScript**: associated existing script.
	//
	// example:
	//
	// HandInput
	ScriptType *string `json:"ScriptType,omitempty" xml:"ScriptType,omitempty"`
	// The ID of the O&M task.
	//
	// example:
	//
	// 1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status of the O&M task.
	//
	// - **PendingApproval**: pending approval.
	//
	// - **Rejected**: rejected.
	//
	// - **Cancelled**: cancelled.
	//
	// - **PendingExecution**: approved and waiting for execution.
	//
	// - **PrepareRun**: preparing to execute.
	//
	// - **Running**: executing.
	//
	// - **Completed**: execution completed.
	//
	// - **Failed**: execution failed.
	//
	// example:
	//
	// PendingApproval
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
}

func (s GetAutoOpsTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetAutoOpsTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetAutoOpsTaskResponseBodyTask) GetAllowedOverTime() *int64 {
	return s.AllowedOverTime
}

func (s *GetAutoOpsTaskResponseBodyTask) GetComment() *string {
	return s.Comment
}

func (s *GetAutoOpsTaskResponseBodyTask) GetName() *string {
	return s.Name
}

func (s *GetAutoOpsTaskResponseBodyTask) GetScheduleTimeInfo() *string {
	return s.ScheduleTimeInfo
}

func (s *GetAutoOpsTaskResponseBodyTask) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetAutoOpsTaskResponseBodyTask) GetScript() *string {
	return s.Script
}

func (s *GetAutoOpsTaskResponseBodyTask) GetScriptId() *int64 {
	return s.ScriptId
}

func (s *GetAutoOpsTaskResponseBodyTask) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetAutoOpsTaskResponseBodyTask) GetScriptType() *string {
	return s.ScriptType
}

func (s *GetAutoOpsTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAutoOpsTaskResponseBodyTask) GetTaskState() *string {
	return s.TaskState
}

func (s *GetAutoOpsTaskResponseBodyTask) SetAllowedOverTime(v int64) *GetAutoOpsTaskResponseBodyTask {
	s.AllowedOverTime = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetComment(v string) *GetAutoOpsTaskResponseBodyTask {
	s.Comment = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetName(v string) *GetAutoOpsTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetScheduleTimeInfo(v string) *GetAutoOpsTaskResponseBodyTask {
	s.ScheduleTimeInfo = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetScheduleType(v string) *GetAutoOpsTaskResponseBodyTask {
	s.ScheduleType = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetScript(v string) *GetAutoOpsTaskResponseBodyTask {
	s.Script = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetScriptId(v int64) *GetAutoOpsTaskResponseBodyTask {
	s.ScriptId = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetScriptName(v string) *GetAutoOpsTaskResponseBodyTask {
	s.ScriptName = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetScriptType(v string) *GetAutoOpsTaskResponseBodyTask {
	s.ScriptType = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetTaskId(v string) *GetAutoOpsTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) SetTaskState(v string) *GetAutoOpsTaskResponseBodyTask {
	s.TaskState = &v
	return s
}

func (s *GetAutoOpsTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
