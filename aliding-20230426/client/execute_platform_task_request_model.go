// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePlatformTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppType(v string) *ExecutePlatformTaskRequest
  GetAppType() *string 
  SetFormDataJson(v string) *ExecutePlatformTaskRequest
  GetFormDataJson() *string 
  SetLanguage(v string) *ExecutePlatformTaskRequest
  GetLanguage() *string 
  SetNoExecuteExpressions(v string) *ExecutePlatformTaskRequest
  GetNoExecuteExpressions() *string 
  SetOutResult(v string) *ExecutePlatformTaskRequest
  GetOutResult() *string 
  SetProcessInstanceId(v string) *ExecutePlatformTaskRequest
  GetProcessInstanceId() *string 
  SetRemark(v string) *ExecutePlatformTaskRequest
  GetRemark() *string 
  SetSystemToken(v string) *ExecutePlatformTaskRequest
  GetSystemToken() *string 
}

type ExecutePlatformTaskRequest struct {
  AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
  FormDataJson *string `json:"FormDataJson,omitempty" xml:"FormDataJson,omitempty"`
  Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
  NoExecuteExpressions *string `json:"NoExecuteExpressions,omitempty" xml:"NoExecuteExpressions,omitempty"`
  OutResult *string `json:"OutResult,omitempty" xml:"OutResult,omitempty"`
  ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s ExecutePlatformTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecutePlatformTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecutePlatformTaskRequest) GetAppType() *string  {
  return s.AppType
}

func (s *ExecutePlatformTaskRequest) GetFormDataJson() *string  {
  return s.FormDataJson
}

func (s *ExecutePlatformTaskRequest) GetLanguage() *string  {
  return s.Language
}

func (s *ExecutePlatformTaskRequest) GetNoExecuteExpressions() *string  {
  return s.NoExecuteExpressions
}

func (s *ExecutePlatformTaskRequest) GetOutResult() *string  {
  return s.OutResult
}

func (s *ExecutePlatformTaskRequest) GetProcessInstanceId() *string  {
  return s.ProcessInstanceId
}

func (s *ExecutePlatformTaskRequest) GetRemark() *string  {
  return s.Remark
}

func (s *ExecutePlatformTaskRequest) GetSystemToken() *string  {
  return s.SystemToken
}

func (s *ExecutePlatformTaskRequest) SetAppType(v string) *ExecutePlatformTaskRequest {
  s.AppType = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetFormDataJson(v string) *ExecutePlatformTaskRequest {
  s.FormDataJson = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetLanguage(v string) *ExecutePlatformTaskRequest {
  s.Language = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetNoExecuteExpressions(v string) *ExecutePlatformTaskRequest {
  s.NoExecuteExpressions = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetOutResult(v string) *ExecutePlatformTaskRequest {
  s.OutResult = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetProcessInstanceId(v string) *ExecutePlatformTaskRequest {
  s.ProcessInstanceId = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetRemark(v string) *ExecutePlatformTaskRequest {
  s.Remark = &v
  return s
}

func (s *ExecutePlatformTaskRequest) SetSystemToken(v string) *ExecutePlatformTaskRequest {
  s.SystemToken = &v
  return s
}

func (s *ExecutePlatformTaskRequest) Validate() error {
  return dara.Validate(s)
}

