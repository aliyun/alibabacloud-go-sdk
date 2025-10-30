// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdHocTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecuteCommand(v *ExecuteAdHocTaskRequestExecuteCommand) *ExecuteAdHocTaskRequest
  GetExecuteCommand() *ExecuteAdHocTaskRequestExecuteCommand 
  SetOpTenantId(v int64) *ExecuteAdHocTaskRequest
  GetOpTenantId() *int64 
}

type ExecuteAdHocTaskRequest struct {
  // This parameter is required.
  ExecuteCommand *ExecuteAdHocTaskRequestExecuteCommand `json:"ExecuteCommand,omitempty" xml:"ExecuteCommand,omitempty" type:"Struct"`
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteAdHocTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskRequest) GetExecuteCommand() *ExecuteAdHocTaskRequestExecuteCommand  {
  return s.ExecuteCommand
}

func (s *ExecuteAdHocTaskRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteAdHocTaskRequest) SetExecuteCommand(v *ExecuteAdHocTaskRequestExecuteCommand) *ExecuteAdHocTaskRequest {
  s.ExecuteCommand = v
  return s
}

func (s *ExecuteAdHocTaskRequest) SetOpTenantId(v int64) *ExecuteAdHocTaskRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteAdHocTaskRequest) Validate() error {
  if s.ExecuteCommand != nil {
    if err := s.ExecuteCommand.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAdHocTaskRequestExecuteCommand struct {
  // This parameter is required.
  // 
  // example:
  // 
  // show tables;
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // mysql_catalog
  DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
  // example:
  // 
  // 12131111
  DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
  // example:
  // 
  // erp
  DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // MaxCompute_SQL
  OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
  ParamList []*ExecuteAdHocTaskRequestExecuteCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 123222121
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecuteAdHocTaskRequestExecuteCommand) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskRequestExecuteCommand) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetCode() *string  {
  return s.Code
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetDataSourceCatalog() *string  {
  return s.DataSourceCatalog
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetDataSourceId() *int64  {
  return s.DataSourceId
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetDataSourceSchema() *string  {
  return s.DataSourceSchema
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetOperatorType() *string  {
  return s.OperatorType
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetParamList() []*ExecuteAdHocTaskRequestExecuteCommandParamList  {
  return s.ParamList
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetCode(v string) *ExecuteAdHocTaskRequestExecuteCommand {
  s.Code = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetDataSourceCatalog(v string) *ExecuteAdHocTaskRequestExecuteCommand {
  s.DataSourceCatalog = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetDataSourceId(v int64) *ExecuteAdHocTaskRequestExecuteCommand {
  s.DataSourceId = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetDataSourceSchema(v string) *ExecuteAdHocTaskRequestExecuteCommand {
  s.DataSourceSchema = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetOperatorType(v string) *ExecuteAdHocTaskRequestExecuteCommand {
  s.OperatorType = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetParamList(v []*ExecuteAdHocTaskRequestExecuteCommandParamList) *ExecuteAdHocTaskRequestExecuteCommand {
  s.ParamList = v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) SetProjectId(v int64) *ExecuteAdHocTaskRequestExecuteCommand {
  s.ProjectId = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommand) Validate() error {
  if s.ParamList != nil {
    for _, item := range s.ParamList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteAdHocTaskRequestExecuteCommandParamList struct {
  // This parameter is required.
  // 
  // example:
  // 
  // param1
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // value1
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecuteAdHocTaskRequestExecuteCommandParamList) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskRequestExecuteCommandParamList) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskRequestExecuteCommandParamList) GetKey() *string  {
  return s.Key
}

func (s *ExecuteAdHocTaskRequestExecuteCommandParamList) GetValue() *string  {
  return s.Value
}

func (s *ExecuteAdHocTaskRequestExecuteCommandParamList) SetKey(v string) *ExecuteAdHocTaskRequestExecuteCommandParamList {
  s.Key = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommandParamList) SetValue(v string) *ExecuteAdHocTaskRequestExecuteCommandParamList {
  s.Value = &v
  return s
}

func (s *ExecuteAdHocTaskRequestExecuteCommandParamList) Validate() error {
  return dara.Validate(s)
}

