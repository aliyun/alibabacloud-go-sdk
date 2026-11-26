// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScheduledTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppToken(v string) *ExecuteScheduledTaskRequest
  GetAppToken() *string 
  SetOpsUserCode(v string) *ExecuteScheduledTaskRequest
  GetOpsUserCode() *string 
  SetOpsUserId(v string) *ExecuteScheduledTaskRequest
  GetOpsUserId() *string 
  SetTaskId(v string) *ExecuteScheduledTaskRequest
  GetTaskId() *string 
  SetTenantId(v string) *ExecuteScheduledTaskRequest
  GetTenantId() *string 
}

type ExecuteScheduledTaskRequest struct {
  // 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
  // 
  // example:
  // 
  // wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
  // 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
  // 
  // example:
  // 
  // alice
  OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
  // 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
  // 
  // example:
  // 
  // 10001
  OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
  // 任务 ID
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // exampleTaskId
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // 租户ID，公共参数，缺省时使用调用方默认租户
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ExecuteScheduledTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScheduledTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteScheduledTaskRequest) GetAppToken() *string  {
  return s.AppToken
}

func (s *ExecuteScheduledTaskRequest) GetOpsUserCode() *string  {
  return s.OpsUserCode
}

func (s *ExecuteScheduledTaskRequest) GetOpsUserId() *string  {
  return s.OpsUserId
}

func (s *ExecuteScheduledTaskRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExecuteScheduledTaskRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExecuteScheduledTaskRequest) SetAppToken(v string) *ExecuteScheduledTaskRequest {
  s.AppToken = &v
  return s
}

func (s *ExecuteScheduledTaskRequest) SetOpsUserCode(v string) *ExecuteScheduledTaskRequest {
  s.OpsUserCode = &v
  return s
}

func (s *ExecuteScheduledTaskRequest) SetOpsUserId(v string) *ExecuteScheduledTaskRequest {
  s.OpsUserId = &v
  return s
}

func (s *ExecuteScheduledTaskRequest) SetTaskId(v string) *ExecuteScheduledTaskRequest {
  s.TaskId = &v
  return s
}

func (s *ExecuteScheduledTaskRequest) SetTenantId(v string) *ExecuteScheduledTaskRequest {
  s.TenantId = &v
  return s
}

func (s *ExecuteScheduledTaskRequest) Validate() error {
  return dara.Validate(s)
}

