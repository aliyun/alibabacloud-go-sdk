// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgGremlinRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecCommand(v *ExecKgGremlinRequestExecCommand) *ExecKgGremlinRequest
  GetExecCommand() *ExecKgGremlinRequestExecCommand 
  SetOpTenantId(v int64) *ExecKgGremlinRequest
  GetOpTenantId() *int64 
  SetOpUserId(v string) *ExecKgGremlinRequest
  GetOpUserId() *string 
  SetWorkspaceId(v string) *ExecKgGremlinRequest
  GetWorkspaceId() *string 
}

type ExecKgGremlinRequest struct {
  // The custom Cypher query command.
  // 
  // This parameter is required.
  ExecCommand *ExecKgGremlinRequestExecCommand `json:"ExecCommand,omitempty" xml:"ExecCommand,omitempty" type:"Struct"`
  // The tenant ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
  // The ID of the operator user.
  // 
  // example:
  // 
  // 30001011
  OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
  // The model ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // f1d4559a4db044158305e2d89bccf81f
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExecKgGremlinRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinRequest) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinRequest) GetExecCommand() *ExecKgGremlinRequestExecCommand  {
  return s.ExecCommand
}

func (s *ExecKgGremlinRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecKgGremlinRequest) GetOpUserId() *string  {
  return s.OpUserId
}

func (s *ExecKgGremlinRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExecKgGremlinRequest) SetExecCommand(v *ExecKgGremlinRequestExecCommand) *ExecKgGremlinRequest {
  s.ExecCommand = v
  return s
}

func (s *ExecKgGremlinRequest) SetOpTenantId(v int64) *ExecKgGremlinRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecKgGremlinRequest) SetOpUserId(v string) *ExecKgGremlinRequest {
  s.OpUserId = &v
  return s
}

func (s *ExecKgGremlinRequest) SetWorkspaceId(v string) *ExecKgGremlinRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExecKgGremlinRequest) Validate() error {
  if s.ExecCommand != nil {
    if err := s.ExecCommand.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecKgGremlinRequestExecCommand struct {
  // The maximum number of records to return.
  // 
  // example:
  // 
  // 100
  Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
  // The input parameters of the query statement.
  Params []*ExecKgGremlinRequestExecCommandParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // The custom Cypher query statement.
  // 
  // example:
  // 
  // MATCH (n) RETURN n LIMIT 10
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ExecKgGremlinRequestExecCommand) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinRequestExecCommand) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinRequestExecCommand) GetLimit() *int32  {
  return s.Limit
}

func (s *ExecKgGremlinRequestExecCommand) GetParams() []*ExecKgGremlinRequestExecCommandParams  {
  return s.Params
}

func (s *ExecKgGremlinRequestExecCommand) GetQuery() *string  {
  return s.Query
}

func (s *ExecKgGremlinRequestExecCommand) SetLimit(v int32) *ExecKgGremlinRequestExecCommand {
  s.Limit = &v
  return s
}

func (s *ExecKgGremlinRequestExecCommand) SetParams(v []*ExecKgGremlinRequestExecCommandParams) *ExecKgGremlinRequestExecCommand {
  s.Params = v
  return s
}

func (s *ExecKgGremlinRequestExecCommand) SetQuery(v string) *ExecKgGremlinRequestExecCommand {
  s.Query = &v
  return s
}

func (s *ExecKgGremlinRequestExecCommand) Validate() error {
  if s.Params != nil {
    for _, item := range s.Params {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecKgGremlinRequestExecCommandParams struct {
  // The data type of paramValue.
  // 
  // example:
  // 
  // STRING
  DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
  // paramKey
  // 
  // example:
  // 
  // name
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // paramValue
  // 
  // example:
  // 
  // Alibaba
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecKgGremlinRequestExecCommandParams) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinRequestExecCommandParams) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinRequestExecCommandParams) GetDataType() *string  {
  return s.DataType
}

func (s *ExecKgGremlinRequestExecCommandParams) GetKey() *string  {
  return s.Key
}

func (s *ExecKgGremlinRequestExecCommandParams) GetValue() *string  {
  return s.Value
}

func (s *ExecKgGremlinRequestExecCommandParams) SetDataType(v string) *ExecKgGremlinRequestExecCommandParams {
  s.DataType = &v
  return s
}

func (s *ExecKgGremlinRequestExecCommandParams) SetKey(v string) *ExecKgGremlinRequestExecCommandParams {
  s.Key = &v
  return s
}

func (s *ExecKgGremlinRequestExecCommandParams) SetValue(v string) *ExecKgGremlinRequestExecCommandParams {
  s.Value = &v
  return s
}

func (s *ExecKgGremlinRequestExecCommandParams) Validate() error {
  return dara.Validate(s)
}

