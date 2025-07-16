// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppType(v string) *ExecuteTaskRequest
  GetAppType() *string 
  SetDigitalSignUrl(v string) *ExecuteTaskRequest
  GetDigitalSignUrl() *string 
  SetFormDataJson(v string) *ExecuteTaskRequest
  GetFormDataJson() *string 
  SetLanguage(v string) *ExecuteTaskRequest
  GetLanguage() *string 
  SetNoExecuteExpressions(v string) *ExecuteTaskRequest
  GetNoExecuteExpressions() *string 
  SetOutResult(v string) *ExecuteTaskRequest
  GetOutResult() *string 
  SetProcessInstanceId(v string) *ExecuteTaskRequest
  GetProcessInstanceId() *string 
  SetRemark(v string) *ExecuteTaskRequest
  GetRemark() *string 
  SetSystemToken(v string) *ExecuteTaskRequest
  GetSystemToken() *string 
  SetTaskId(v int64) *ExecuteTaskRequest
  GetTaskId() *int64 
}

type ExecuteTaskRequest struct {
  // example:
  // 
  // APP_PBKT0MFBEBTDO8T7SLVP
  AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
  // example:
  // 
  // http://
  DigitalSignUrl *string `json:"DigitalSignUrl,omitempty" xml:"DigitalSignUrl,omitempty"`
  // example:
  // 
  // {}
  FormDataJson *string `json:"FormDataJson,omitempty" xml:"FormDataJson,omitempty"`
  // example:
  // 
  // zh_CN
  Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
  // example:
  // 
  // y
  NoExecuteExpressions *string `json:"NoExecuteExpressions,omitempty" xml:"NoExecuteExpressions,omitempty"`
  // example:
  // 
  // AGREE
  OutResult *string `json:"OutResult,omitempty" xml:"OutResult,omitempty"`
  // example:
  // 
  // f30233fb-72e1-4af4-8cb8-c7e0ea9ee530
  ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  // example:
  // 
  // hexxyy
  SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
  // example:
  // 
  // 12002575L
  TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExecuteTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTaskRequest) GetAppType() *string  {
  return s.AppType
}

func (s *ExecuteTaskRequest) GetDigitalSignUrl() *string  {
  return s.DigitalSignUrl
}

func (s *ExecuteTaskRequest) GetFormDataJson() *string  {
  return s.FormDataJson
}

func (s *ExecuteTaskRequest) GetLanguage() *string  {
  return s.Language
}

func (s *ExecuteTaskRequest) GetNoExecuteExpressions() *string  {
  return s.NoExecuteExpressions
}

func (s *ExecuteTaskRequest) GetOutResult() *string  {
  return s.OutResult
}

func (s *ExecuteTaskRequest) GetProcessInstanceId() *string  {
  return s.ProcessInstanceId
}

func (s *ExecuteTaskRequest) GetRemark() *string  {
  return s.Remark
}

func (s *ExecuteTaskRequest) GetSystemToken() *string  {
  return s.SystemToken
}

func (s *ExecuteTaskRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *ExecuteTaskRequest) SetAppType(v string) *ExecuteTaskRequest {
  s.AppType = &v
  return s
}

func (s *ExecuteTaskRequest) SetDigitalSignUrl(v string) *ExecuteTaskRequest {
  s.DigitalSignUrl = &v
  return s
}

func (s *ExecuteTaskRequest) SetFormDataJson(v string) *ExecuteTaskRequest {
  s.FormDataJson = &v
  return s
}

func (s *ExecuteTaskRequest) SetLanguage(v string) *ExecuteTaskRequest {
  s.Language = &v
  return s
}

func (s *ExecuteTaskRequest) SetNoExecuteExpressions(v string) *ExecuteTaskRequest {
  s.NoExecuteExpressions = &v
  return s
}

func (s *ExecuteTaskRequest) SetOutResult(v string) *ExecuteTaskRequest {
  s.OutResult = &v
  return s
}

func (s *ExecuteTaskRequest) SetProcessInstanceId(v string) *ExecuteTaskRequest {
  s.ProcessInstanceId = &v
  return s
}

func (s *ExecuteTaskRequest) SetRemark(v string) *ExecuteTaskRequest {
  s.Remark = &v
  return s
}

func (s *ExecuteTaskRequest) SetSystemToken(v string) *ExecuteTaskRequest {
  s.SystemToken = &v
  return s
}

func (s *ExecuteTaskRequest) SetTaskId(v int64) *ExecuteTaskRequest {
  s.TaskId = &v
  return s
}

func (s *ExecuteTaskRequest) Validate() error {
  return dara.Validate(s)
}

