// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgCypherRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecCommand(v *ExecKgCypherRequestExecCommand) *ExecKgCypherRequest
  GetExecCommand() *ExecKgCypherRequestExecCommand 
  SetOpTenantId(v int64) *ExecKgCypherRequest
  GetOpTenantId() *int64 
  SetOpUserId(v string) *ExecKgCypherRequest
  GetOpUserId() *string 
  SetWorkspaceId(v string) *ExecKgCypherRequest
  GetWorkspaceId() *string 
}

type ExecKgCypherRequest struct {
  // The custom Cypher query instruction.
  // 
  // This parameter is required.
  ExecCommand *ExecKgCypherRequestExecCommand `json:"ExecCommand,omitempty" xml:"ExecCommand,omitempty" type:"Struct"`
  // The tenant ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s ExecKgCypherRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherRequest) GoString() string {
  return s.String()
}

func (s *ExecKgCypherRequest) GetExecCommand() *ExecKgCypherRequestExecCommand  {
  return s.ExecCommand
}

func (s *ExecKgCypherRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecKgCypherRequest) GetOpUserId() *string  {
  return s.OpUserId
}

func (s *ExecKgCypherRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExecKgCypherRequest) SetExecCommand(v *ExecKgCypherRequestExecCommand) *ExecKgCypherRequest {
  s.ExecCommand = v
  return s
}

func (s *ExecKgCypherRequest) SetOpTenantId(v int64) *ExecKgCypherRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecKgCypherRequest) SetOpUserId(v string) *ExecKgCypherRequest {
  s.OpUserId = &v
  return s
}

func (s *ExecKgCypherRequest) SetWorkspaceId(v string) *ExecKgCypherRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExecKgCypherRequest) Validate() error {
  if s.ExecCommand != nil {
    if err := s.ExecCommand.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecKgCypherRequestExecCommand struct {
  // The maximum number of records to return.
  // 
  // example:
  // 
  // 100
  Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
  // The input parameters of the query statement.
  Params []*ExecKgCypherRequestExecCommandParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // The custom Cypher query statement.
  // 
  // example:
  // 
  // MATCH (n) RETURN n LIMIT 10
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ExecKgCypherRequestExecCommand) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherRequestExecCommand) GoString() string {
  return s.String()
}

func (s *ExecKgCypherRequestExecCommand) GetLimit() *int32  {
  return s.Limit
}

func (s *ExecKgCypherRequestExecCommand) GetParams() []*ExecKgCypherRequestExecCommandParams  {
  return s.Params
}

func (s *ExecKgCypherRequestExecCommand) GetQuery() *string  {
  return s.Query
}

func (s *ExecKgCypherRequestExecCommand) SetLimit(v int32) *ExecKgCypherRequestExecCommand {
  s.Limit = &v
  return s
}

func (s *ExecKgCypherRequestExecCommand) SetParams(v []*ExecKgCypherRequestExecCommandParams) *ExecKgCypherRequestExecCommand {
  s.Params = v
  return s
}

func (s *ExecKgCypherRequestExecCommand) SetQuery(v string) *ExecKgCypherRequestExecCommand {
  s.Query = &v
  return s
}

func (s *ExecKgCypherRequestExecCommand) Validate() error {
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

type ExecKgCypherRequestExecCommandParams struct {
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

func (s ExecKgCypherRequestExecCommandParams) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherRequestExecCommandParams) GoString() string {
  return s.String()
}

func (s *ExecKgCypherRequestExecCommandParams) GetDataType() *string  {
  return s.DataType
}

func (s *ExecKgCypherRequestExecCommandParams) GetKey() *string  {
  return s.Key
}

func (s *ExecKgCypherRequestExecCommandParams) GetValue() *string  {
  return s.Value
}

func (s *ExecKgCypherRequestExecCommandParams) SetDataType(v string) *ExecKgCypherRequestExecCommandParams {
  s.DataType = &v
  return s
}

func (s *ExecKgCypherRequestExecCommandParams) SetKey(v string) *ExecKgCypherRequestExecCommandParams {
  s.Key = &v
  return s
}

func (s *ExecKgCypherRequestExecCommandParams) SetValue(v string) *ExecKgCypherRequestExecCommandParams {
  s.Value = &v
  return s
}

func (s *ExecKgCypherRequestExecCommandParams) Validate() error {
  return dara.Validate(s)
}

