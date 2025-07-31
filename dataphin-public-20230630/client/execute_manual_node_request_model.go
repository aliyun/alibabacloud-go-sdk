// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteManualNodeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnv(v string) *ExecuteManualNodeRequest
  GetEnv() *string 
  SetExecuteCommand(v *ExecuteManualNodeRequestExecuteCommand) *ExecuteManualNodeRequest
  GetExecuteCommand() *ExecuteManualNodeRequestExecuteCommand 
  SetOpTenantId(v int64) *ExecuteManualNodeRequest
  GetOpTenantId() *int64 
}

type ExecuteManualNodeRequest struct {
  // example:
  // 
  // PROD
  Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
  // This parameter is required.
  ExecuteCommand *ExecuteManualNodeRequestExecuteCommand `json:"ExecuteCommand,omitempty" xml:"ExecuteCommand,omitempty" type:"Struct"`
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteManualNodeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteManualNodeRequest) GoString() string {
  return s.String()
}

func (s *ExecuteManualNodeRequest) GetEnv() *string  {
  return s.Env
}

func (s *ExecuteManualNodeRequest) GetExecuteCommand() *ExecuteManualNodeRequestExecuteCommand  {
  return s.ExecuteCommand
}

func (s *ExecuteManualNodeRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteManualNodeRequest) SetEnv(v string) *ExecuteManualNodeRequest {
  s.Env = &v
  return s
}

func (s *ExecuteManualNodeRequest) SetExecuteCommand(v *ExecuteManualNodeRequestExecuteCommand) *ExecuteManualNodeRequest {
  s.ExecuteCommand = v
  return s
}

func (s *ExecuteManualNodeRequest) SetOpTenantId(v int64) *ExecuteManualNodeRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteManualNodeRequest) Validate() error {
  return dara.Validate(s)
}

type ExecuteManualNodeRequestExecuteCommand struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 2024-05-07
  EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
  // example:
  // 
  // xx测试
  FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // n_12132
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  ParamList []*ExecuteManualNodeRequestExecuteCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 123324
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2024-05-01
  StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s ExecuteManualNodeRequestExecuteCommand) String() string {
  return dara.Prettify(s)
}

func (s ExecuteManualNodeRequestExecuteCommand) GoString() string {
  return s.String()
}

func (s *ExecuteManualNodeRequestExecuteCommand) GetEndBizDate() *string  {
  return s.EndBizDate
}

func (s *ExecuteManualNodeRequestExecuteCommand) GetFlowName() *string  {
  return s.FlowName
}

func (s *ExecuteManualNodeRequestExecuteCommand) GetNodeId() *string  {
  return s.NodeId
}

func (s *ExecuteManualNodeRequestExecuteCommand) GetParamList() []*ExecuteManualNodeRequestExecuteCommandParamList  {
  return s.ParamList
}

func (s *ExecuteManualNodeRequestExecuteCommand) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteManualNodeRequestExecuteCommand) GetStartBizDate() *string  {
  return s.StartBizDate
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetEndBizDate(v string) *ExecuteManualNodeRequestExecuteCommand {
  s.EndBizDate = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetFlowName(v string) *ExecuteManualNodeRequestExecuteCommand {
  s.FlowName = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetNodeId(v string) *ExecuteManualNodeRequestExecuteCommand {
  s.NodeId = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetParamList(v []*ExecuteManualNodeRequestExecuteCommandParamList) *ExecuteManualNodeRequestExecuteCommand {
  s.ParamList = v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetProjectId(v int64) *ExecuteManualNodeRequestExecuteCommand {
  s.ProjectId = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetStartBizDate(v string) *ExecuteManualNodeRequestExecuteCommand {
  s.StartBizDate = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) Validate() error {
  return dara.Validate(s)
}

type ExecuteManualNodeRequestExecuteCommandParamList struct {
  // example:
  // 
  // param1
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // example:
  // 
  // 1
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecuteManualNodeRequestExecuteCommandParamList) String() string {
  return dara.Prettify(s)
}

func (s ExecuteManualNodeRequestExecuteCommandParamList) GoString() string {
  return s.String()
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) GetKey() *string  {
  return s.Key
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) GetValue() *string  {
  return s.Value
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) SetKey(v string) *ExecuteManualNodeRequestExecuteCommandParamList {
  s.Key = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) SetValue(v string) *ExecuteManualNodeRequestExecuteCommandParamList {
  s.Value = &v
  return s
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) Validate() error {
  return dara.Validate(s)
}

