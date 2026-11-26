// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScheduledTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteScheduledTaskResponseBody
  GetCode() *string 
  SetExecutionId(v string) *ExecuteScheduledTaskResponseBody
  GetExecutionId() *string 
  SetMessage(v string) *ExecuteScheduledTaskResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteScheduledTaskResponseBody
  GetRequestId() *string 
  SetTaskExecution(v *ExecuteScheduledTaskResponseBodyTaskExecution) *ExecuteScheduledTaskResponseBody
  GetTaskExecution() *ExecuteScheduledTaskResponseBodyTaskExecution 
}

type ExecuteScheduledTaskResponseBody struct {
  // 业务状态码：成功为 200，失败为后端错误码
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // 执行记录 ID
  // 
  // example:
  // 
  // exampleExecutionId
  ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
  // 错误描述，成功时为空
  // 
  // example:
  // 
  // ok
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // 请求追踪 ID
  // 
  // example:
  // 
  // 019FF406-1B10-0065-A97D-2D1920C2A03D
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  TaskExecution *ExecuteScheduledTaskResponseBodyTaskExecution `json:"taskExecution,omitempty" xml:"taskExecution,omitempty" type:"Struct"`
}

func (s ExecuteScheduledTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScheduledTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteScheduledTaskResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteScheduledTaskResponseBody) GetExecutionId() *string  {
  return s.ExecutionId
}

func (s *ExecuteScheduledTaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteScheduledTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteScheduledTaskResponseBody) GetTaskExecution() *ExecuteScheduledTaskResponseBodyTaskExecution  {
  return s.TaskExecution
}

func (s *ExecuteScheduledTaskResponseBody) SetCode(v string) *ExecuteScheduledTaskResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBody) SetExecutionId(v string) *ExecuteScheduledTaskResponseBody {
  s.ExecutionId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBody) SetMessage(v string) *ExecuteScheduledTaskResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBody) SetRequestId(v string) *ExecuteScheduledTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBody) SetTaskExecution(v *ExecuteScheduledTaskResponseBodyTaskExecution) *ExecuteScheduledTaskResponseBody {
  s.TaskExecution = v
  return s
}

func (s *ExecuteScheduledTaskResponseBody) Validate() error {
  if s.TaskExecution != nil {
    if err := s.TaskExecution.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteScheduledTaskResponseBodyTaskExecution struct {
  // 任务 AUTO 偏好：none=默认、performance=效果优先、credit=成本优先
  // 
  // example:
  // 
  // none
  AutoPreference *string `json:"autoPreference,omitempty" xml:"autoPreference,omitempty"`
  // example:
  // 
  // BL-10000-0123456789abcdef0123456789abcdef
  BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
  // example:
  // 
  // exampleCollaborationGroupId
  CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
  EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
  // example:
  // 
  // exampleExecutionId
  ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
  // 执行模型档位
  // 
  // example:
  // 
  // auto
  Model *string `json:"model,omitempty" xml:"model,omitempty"`
  OperatingObjectName []*string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty" type:"Repeated"`
  SkillCodes []*string `json:"skillCodes,omitempty" xml:"skillCodes,omitempty" type:"Repeated"`
  // example:
  // 
  // exampleTaskId
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // example:
  // 
  // 每日销售数据汇总
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
  // example:
  // 
  // 每日汇总销售数据并生成报告
  TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
  // example:
  // 
  // 10001
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteScheduledTaskResponseBodyTaskExecution) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScheduledTaskResponseBodyTaskExecution) GoString() string {
  return s.String()
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetAutoPreference() *string  {
  return s.AutoPreference
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetBillingId() *string  {
  return s.BillingId
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetCollaborationGroupId() *string  {
  return s.CollaborationGroupId
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetEnableWebSearch() *bool  {
  return s.EnableWebSearch
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetExecutionId() *string  {
  return s.ExecutionId
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetModel() *string  {
  return s.Model
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetOperatingObjectName() []*string  {
  return s.OperatingObjectName
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetSkillCodes() []*string  {
  return s.SkillCodes
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetTaskName() *string  {
  return s.TaskName
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetTaskUnderstand() *string  {
  return s.TaskUnderstand
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetAutoPreference(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.AutoPreference = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetBillingId(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.BillingId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetCollaborationGroupId(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.CollaborationGroupId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetEnableWebSearch(v bool) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.EnableWebSearch = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetExecutionId(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.ExecutionId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetModel(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.Model = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetOperatingObjectName(v []*string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.OperatingObjectName = v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetSkillCodes(v []*string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.SkillCodes = v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetTaskId(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.TaskId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetTaskName(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.TaskName = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetTaskUnderstand(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.TaskUnderstand = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetTenantId(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.TenantId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) SetUserId(v string) *ExecuteScheduledTaskResponseBodyTaskExecution {
  s.UserId = &v
  return s
}

func (s *ExecuteScheduledTaskResponseBodyTaskExecution) Validate() error {
  return dara.Validate(s)
}

