// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgCypherShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecCommandShrink(v string) *ExecKgCypherShrinkRequest
  GetExecCommandShrink() *string 
  SetOpTenantId(v int64) *ExecKgCypherShrinkRequest
  GetOpTenantId() *int64 
  SetOpUserId(v string) *ExecKgCypherShrinkRequest
  GetOpUserId() *string 
  SetWorkspaceId(v string) *ExecKgCypherShrinkRequest
  GetWorkspaceId() *string 
}

type ExecKgCypherShrinkRequest struct {
  // The custom Cypher query instruction.
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

func (s ExecKgCypherShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecKgCypherShrinkRequest) GetExecCommandShrink() *string  {
  return s.ExecCommandShrink
}

func (s *ExecKgCypherShrinkRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecKgCypherShrinkRequest) GetOpUserId() *string  {
  return s.OpUserId
}

func (s *ExecKgCypherShrinkRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExecKgCypherShrinkRequest) SetExecCommandShrink(v string) *ExecKgCypherShrinkRequest {
  s.ExecCommandShrink = &v
  return s
}

func (s *ExecKgCypherShrinkRequest) SetOpTenantId(v int64) *ExecKgCypherShrinkRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecKgCypherShrinkRequest) SetOpUserId(v string) *ExecKgCypherShrinkRequest {
  s.OpUserId = &v
  return s
}

func (s *ExecKgCypherShrinkRequest) SetWorkspaceId(v string) *ExecKgCypherShrinkRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExecKgCypherShrinkRequest) Validate() error {
  return dara.Validate(s)
}

