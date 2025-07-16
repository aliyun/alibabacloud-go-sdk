// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBatchTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppType(v string) *ExecuteBatchTaskRequest
  GetAppType() *string 
  SetOutResult(v string) *ExecuteBatchTaskRequest
  GetOutResult() *string 
  SetRemark(v string) *ExecuteBatchTaskRequest
  GetRemark() *string 
  SetSystemToken(v string) *ExecuteBatchTaskRequest
  GetSystemToken() *string 
  SetTaskInformationList(v string) *ExecuteBatchTaskRequest
  GetTaskInformationList() *string 
}

type ExecuteBatchTaskRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // APP_PBKTxxx
  AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // disagree
  OutResult *string `json:"OutResult,omitempty" xml:"OutResult,omitempty"`
  // example:
  // 
  // remark
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // hexxxx
  SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // [{"taskId":"2291xxx","formInstId":"d84a79xxx"}, {"taskId":"2291xxx","formInstId":"f8035e2axxx"}]
  TaskInformationList *string `json:"TaskInformationList,omitempty" xml:"TaskInformationList,omitempty"`
}

func (s ExecuteBatchTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBatchTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteBatchTaskRequest) GetAppType() *string  {
  return s.AppType
}

func (s *ExecuteBatchTaskRequest) GetOutResult() *string  {
  return s.OutResult
}

func (s *ExecuteBatchTaskRequest) GetRemark() *string  {
  return s.Remark
}

func (s *ExecuteBatchTaskRequest) GetSystemToken() *string  {
  return s.SystemToken
}

func (s *ExecuteBatchTaskRequest) GetTaskInformationList() *string  {
  return s.TaskInformationList
}

func (s *ExecuteBatchTaskRequest) SetAppType(v string) *ExecuteBatchTaskRequest {
  s.AppType = &v
  return s
}

func (s *ExecuteBatchTaskRequest) SetOutResult(v string) *ExecuteBatchTaskRequest {
  s.OutResult = &v
  return s
}

func (s *ExecuteBatchTaskRequest) SetRemark(v string) *ExecuteBatchTaskRequest {
  s.Remark = &v
  return s
}

func (s *ExecuteBatchTaskRequest) SetSystemToken(v string) *ExecuteBatchTaskRequest {
  s.SystemToken = &v
  return s
}

func (s *ExecuteBatchTaskRequest) SetTaskInformationList(v string) *ExecuteBatchTaskRequest {
  s.TaskInformationList = &v
  return s
}

func (s *ExecuteBatchTaskRequest) Validate() error {
  return dara.Validate(s)
}

