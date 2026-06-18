// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListTaskRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListTaskRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTaskRequest
	GetResourceOwnerId() *int64
	SetRobotName(v string) *ListTaskRequest
	GetRobotName() *string
	SetStatus(v string) *ListTaskRequest
	GetStatus() *string
	SetTaskId(v int64) *ListTaskRequest
	GetTaskId() *int64
	SetTaskName(v string) *ListTaskRequest
	GetTaskName() *string
}

type ListTaskRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of items per page. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The robot name, which is the script name. You can view the names of scripts that have passed Review in the [Script Management](https://aiccs.console.aliyun.com/patter/list) interface.
	//
	// example:
	//
	// 机器人
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// Job Status. Valid values:
	//
	// - **INIT**: Not started.
	//
	// - **RELEASE**: Parsing.
	//
	// - **RUNNING**: Executing.
	//
	// - **STOP**: Paused manually.
	//
	// - **SYSTEM_STOP**: Paused by the system.
	//
	// - **READY**: Pending execution.
	//
	// - **CANCEL**: Stopped manually.
	//
	// - **SYSTEM_CANCEL**: Stopped by the system.
	//
	// - **DONE**: Completed.
	//
	// example:
	//
	// STOP
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique job ID of the robot calling job. You can view it in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface or obtain it by using the [CreateTask](https://help.aliyun.com/document_detail/223556.html) API.
	//
	// example:
	//
	// 12****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The job name. You can view the names of created jobs in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface.
	//
	// example:
	//
	// 任务测试
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskRequest) GoString() string {
	return s.String()
}

func (s *ListTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTaskRequest) GetRobotName() *string {
	return s.RobotName
}

func (s *ListTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskRequest) SetOwnerId(v int64) *ListTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTaskRequest) SetPageNo(v int32) *ListTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListTaskRequest) SetPageSize(v int32) *ListTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskRequest) SetResourceOwnerAccount(v string) *ListTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTaskRequest) SetResourceOwnerId(v int64) *ListTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTaskRequest) SetRobotName(v string) *ListTaskRequest {
	s.RobotName = &v
	return s
}

func (s *ListTaskRequest) SetStatus(v string) *ListTaskRequest {
	s.Status = &v
	return s
}

func (s *ListTaskRequest) SetTaskId(v int64) *ListTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskRequest) SetTaskName(v string) *ListTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskRequest) Validate() error {
	return dara.Validate(s)
}
