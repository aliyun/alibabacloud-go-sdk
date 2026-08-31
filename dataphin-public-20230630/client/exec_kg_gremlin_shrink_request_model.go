// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgGremlinShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecCommandShrink(v string) *ExecKgGremlinShrinkRequest
  GetExecCommandShrink() *string 
  SetOpTenantId(v int64) *ExecKgGremlinShrinkRequest
  GetOpTenantId() *int64 
  SetOpUserId(v string) *ExecKgGremlinShrinkRequest
  GetOpUserId() *string 
  SetWorkspaceId(v string) *ExecKgGremlinShrinkRequest
  GetWorkspaceId() *string 
}

type ExecKgGremlinShrinkRequest struct {
  // The custom Cypher query command.
  // 
  // This parameter is required.
  ExecCommandShrink *string `json:"ExecCommand,omitempty" xml:"ExecCommand,omitempty"`
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

func (s ExecKgGremlinShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinShrinkRequest) GetExecCommandShrink() *string  {
  return s.ExecCommandShrink
}

func (s *ExecKgGremlinShrinkRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecKgGremlinShrinkRequest) GetOpUserId() *string  {
  return s.OpUserId
}

func (s *ExecKgGremlinShrinkRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExecKgGremlinShrinkRequest) SetExecCommandShrink(v string) *ExecKgGremlinShrinkRequest {
  s.ExecCommandShrink = &v
  return s
}

func (s *ExecKgGremlinShrinkRequest) SetOpTenantId(v int64) *ExecKgGremlinShrinkRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecKgGremlinShrinkRequest) SetOpUserId(v string) *ExecKgGremlinShrinkRequest {
  s.OpUserId = &v
  return s
}

func (s *ExecKgGremlinShrinkRequest) SetWorkspaceId(v string) *ExecKgGremlinShrinkRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExecKgGremlinShrinkRequest) Validate() error {
  return dara.Validate(s)
}

