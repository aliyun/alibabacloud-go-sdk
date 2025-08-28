// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteCallTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFireTime(v string) *ExecuteCallTaskRequest
  GetFireTime() *string 
  SetOwnerId(v int64) *ExecuteCallTaskRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *ExecuteCallTaskRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *ExecuteCallTaskRequest
  GetResourceOwnerId() *int64 
  SetStatus(v string) *ExecuteCallTaskRequest
  GetStatus() *string 
  SetTaskId(v int64) *ExecuteCallTaskRequest
  GetTaskId() *int64 
}

type ExecuteCallTaskRequest struct {
  // The time when the call task is executed, in the yyyy-MM-dd HH:mm:ss format.
  // 
  // > You can leave this parameter empty.
  // 
  // example:
  // 
  // 2021-03-09 00:00:00
  FireTime *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The task state. Valid values:
  // 
  // 	- **RUNNING**
  // 
  // 	- **STOP**
  // 
  // 	- **CANCEL**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // RUNNING
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // The task ID. You can call the [CreateCallTask](~~CreateCallTask~~) operation to obtain the task ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2256****
  TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExecuteCallTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteCallTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteCallTaskRequest) GetFireTime() *string  {
  return s.FireTime
}

func (s *ExecuteCallTaskRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExecuteCallTaskRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *ExecuteCallTaskRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *ExecuteCallTaskRequest) GetStatus() *string  {
  return s.Status
}

func (s *ExecuteCallTaskRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *ExecuteCallTaskRequest) SetFireTime(v string) *ExecuteCallTaskRequest {
  s.FireTime = &v
  return s
}

func (s *ExecuteCallTaskRequest) SetOwnerId(v int64) *ExecuteCallTaskRequest {
  s.OwnerId = &v
  return s
}

func (s *ExecuteCallTaskRequest) SetResourceOwnerAccount(v string) *ExecuteCallTaskRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *ExecuteCallTaskRequest) SetResourceOwnerId(v int64) *ExecuteCallTaskRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *ExecuteCallTaskRequest) SetStatus(v string) *ExecuteCallTaskRequest {
  s.Status = &v
  return s
}

func (s *ExecuteCallTaskRequest) SetTaskId(v int64) *ExecuteCallTaskRequest {
  s.TaskId = &v
  return s
}

func (s *ExecuteCallTaskRequest) Validate() error {
  return dara.Validate(s)
}

