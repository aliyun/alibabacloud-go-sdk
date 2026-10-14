// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditTaskState interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *AuditTaskState
	GetEnabled() *bool
	SetRuntimeError(v *AuditTaskStateRuntimeError) *AuditTaskState
	GetRuntimeError() *AuditTaskStateRuntimeError
	SetRuntimeStatus(v string) *AuditTaskState
	GetRuntimeStatus() *string
	SetRuntimeStatusCheckedAt(v string) *AuditTaskState
	GetRuntimeStatusCheckedAt() *string
	SetTask(v map[string]interface{}) *AuditTaskState
	GetTask() map[string]interface{}
	SetTaskId(v string) *AuditTaskState
	GetTaskId() *string
	SetVersion(v int64) *AuditTaskState
	GetVersion() *int64
}

type AuditTaskState struct {
	// The current expected sharing switch for the task.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The runtime error, which is returned only when runtimeStatus is error.
	RuntimeError *AuditTaskStateRuntimeError `json:"runtimeError,omitempty" xml:"runtimeError,omitempty" type:"Struct"`
	// The runtime status, which appears only when includeRuntimeStatus=true is requested. Valid values: pending, starting, stopping, running, stopped, missing, drifted, and error. The list and get operations do not return deleting.
	//
	// example:
	//
	// running
	RuntimeStatus *string `json:"runtimeStatus,omitempty" xml:"runtimeStatus,omitempty"`
	// The time when the remote status is checked. This field appears together with runtimeStatus.
	//
	// example:
	//
	// 1
	RuntimeStatusCheckedAt *string `json:"runtimeStatusCheckedAt,omitempty" xml:"runtimeStatusCheckedAt,omitempty"`
	// The task details in JSON format.
	Task map[string]interface{} `json:"task,omitempty" xml:"task,omitempty"`
	// The stable ID of the task, which must be equal to task.taskId.
	//
	// example:
	//
	// custom_host_intrusion_audit
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The current execution configuration version of the task. The minimum value is 1. Starting or stopping the task does not increment this version.
	//
	// example:
	//
	// 7
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AuditTaskState) String() string {
	return dara.Prettify(s)
}

func (s AuditTaskState) GoString() string {
	return s.String()
}

func (s *AuditTaskState) GetEnabled() *bool {
	return s.Enabled
}

func (s *AuditTaskState) GetRuntimeError() *AuditTaskStateRuntimeError {
	return s.RuntimeError
}

func (s *AuditTaskState) GetRuntimeStatus() *string {
	return s.RuntimeStatus
}

func (s *AuditTaskState) GetRuntimeStatusCheckedAt() *string {
	return s.RuntimeStatusCheckedAt
}

func (s *AuditTaskState) GetTask() map[string]interface{} {
	return s.Task
}

func (s *AuditTaskState) GetTaskId() *string {
	return s.TaskId
}

func (s *AuditTaskState) GetVersion() *int64 {
	return s.Version
}

func (s *AuditTaskState) SetEnabled(v bool) *AuditTaskState {
	s.Enabled = &v
	return s
}

func (s *AuditTaskState) SetRuntimeError(v *AuditTaskStateRuntimeError) *AuditTaskState {
	s.RuntimeError = v
	return s
}

func (s *AuditTaskState) SetRuntimeStatus(v string) *AuditTaskState {
	s.RuntimeStatus = &v
	return s
}

func (s *AuditTaskState) SetRuntimeStatusCheckedAt(v string) *AuditTaskState {
	s.RuntimeStatusCheckedAt = &v
	return s
}

func (s *AuditTaskState) SetTask(v map[string]interface{}) *AuditTaskState {
	s.Task = v
	return s
}

func (s *AuditTaskState) SetTaskId(v string) *AuditTaskState {
	s.TaskId = &v
	return s
}

func (s *AuditTaskState) SetVersion(v int64) *AuditTaskState {
	s.Version = &v
	return s
}

func (s *AuditTaskState) Validate() error {
	if s.RuntimeError != nil {
		if err := s.RuntimeError.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuditTaskStateRuntimeError struct {
	// The error code, such as task_status_query_failed.
	//
	// example:
	//
	// task_status_query_failed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// example:
	//
	// error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the status query request.
	//
	// example:
	//
	// GUID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AuditTaskStateRuntimeError) String() string {
	return dara.Prettify(s)
}

func (s AuditTaskStateRuntimeError) GoString() string {
	return s.String()
}

func (s *AuditTaskStateRuntimeError) GetCode() *string {
	return s.Code
}

func (s *AuditTaskStateRuntimeError) GetMessage() *string {
	return s.Message
}

func (s *AuditTaskStateRuntimeError) GetRequestId() *string {
	return s.RequestId
}

func (s *AuditTaskStateRuntimeError) SetCode(v string) *AuditTaskStateRuntimeError {
	s.Code = &v
	return s
}

func (s *AuditTaskStateRuntimeError) SetMessage(v string) *AuditTaskStateRuntimeError {
	s.Message = &v
	return s
}

func (s *AuditTaskStateRuntimeError) SetRequestId(v string) *AuditTaskStateRuntimeError {
	s.RequestId = &v
	return s
}

func (s *AuditTaskStateRuntimeError) Validate() error {
	return dara.Validate(s)
}
