// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskPopRequest interface {
  dara.Model
  String() string
  GoString() string
  SetConcurrency(v int32) *EditTaskPopRequest
  GetConcurrency() *int32 
  SetDqlTestDatasourceName(v string) *EditTaskPopRequest
  GetDqlTestDatasourceName() *string 
  SetSourceDialect(v string) *EditTaskPopRequest
  GetSourceDialect() *string 
  SetTargetDialect(v string) *EditTaskPopRequest
  GetTargetDialect() *string 
  SetTaskId(v int64) *EditTaskPopRequest
  GetTaskId() *int64 
  SetTaskName(v string) *EditTaskPopRequest
  GetTaskName() *string 
  SetTaskType(v int32) *EditTaskPopRequest
  GetTaskType() *int32 
}

type EditTaskPopRequest struct {
  // The concurrency for controlling the number of concurrent conversion executions.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 5
  Concurrency *int32 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
  // The name of the test data source associated with a DQL task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hive_test_ds
  DqlTestDatasourceName *string `json:"dqlTestDatasourceName,omitempty" xml:"dqlTestDatasourceName,omitempty"`
  // The source SQL dialect.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hive
  SourceDialect *string `json:"sourceDialect,omitempty" xml:"sourceDialect,omitempty"`
  // The target SQL dialect.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hive
  TargetDialect *string `json:"targetDialect,omitempty" xml:"targetDialect,omitempty"`
  // The task ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10001
  TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // The task name.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // data_check_task_demo
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
  // The task type. Valid values:
  // 
  // - 1: DDL
  // 
  // - 2: DQL
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  TaskType *int32 `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s EditTaskPopRequest) String() string {
  return dara.Prettify(s)
}

func (s EditTaskPopRequest) GoString() string {
  return s.String()
}

func (s *EditTaskPopRequest) GetConcurrency() *int32  {
  return s.Concurrency
}

func (s *EditTaskPopRequest) GetDqlTestDatasourceName() *string  {
  return s.DqlTestDatasourceName
}

func (s *EditTaskPopRequest) GetSourceDialect() *string  {
  return s.SourceDialect
}

func (s *EditTaskPopRequest) GetTargetDialect() *string  {
  return s.TargetDialect
}

func (s *EditTaskPopRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *EditTaskPopRequest) GetTaskName() *string  {
  return s.TaskName
}

func (s *EditTaskPopRequest) GetTaskType() *int32  {
  return s.TaskType
}

func (s *EditTaskPopRequest) SetConcurrency(v int32) *EditTaskPopRequest {
  s.Concurrency = &v
  return s
}

func (s *EditTaskPopRequest) SetDqlTestDatasourceName(v string) *EditTaskPopRequest {
  s.DqlTestDatasourceName = &v
  return s
}

func (s *EditTaskPopRequest) SetSourceDialect(v string) *EditTaskPopRequest {
  s.SourceDialect = &v
  return s
}

func (s *EditTaskPopRequest) SetTargetDialect(v string) *EditTaskPopRequest {
  s.TargetDialect = &v
  return s
}

func (s *EditTaskPopRequest) SetTaskId(v int64) *EditTaskPopRequest {
  s.TaskId = &v
  return s
}

func (s *EditTaskPopRequest) SetTaskName(v string) *EditTaskPopRequest {
  s.TaskName = &v
  return s
}

func (s *EditTaskPopRequest) SetTaskType(v int32) *EditTaskPopRequest {
  s.TaskType = &v
  return s
}

func (s *EditTaskPopRequest) Validate() error {
  return dara.Validate(s)
}

