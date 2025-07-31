// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteManualNodeShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnv(v string) *ExecuteManualNodeShrinkRequest
  GetEnv() *string 
  SetExecuteCommandShrink(v string) *ExecuteManualNodeShrinkRequest
  GetExecuteCommandShrink() *string 
  SetOpTenantId(v int64) *ExecuteManualNodeShrinkRequest
  GetOpTenantId() *int64 
}

type ExecuteManualNodeShrinkRequest struct {
  // example:
  // 
  // PROD
  Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
  // This parameter is required.
  ExecuteCommandShrink *string `json:"ExecuteCommand,omitempty" xml:"ExecuteCommand,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteManualNodeShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteManualNodeShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteManualNodeShrinkRequest) GetEnv() *string  {
  return s.Env
}

func (s *ExecuteManualNodeShrinkRequest) GetExecuteCommandShrink() *string  {
  return s.ExecuteCommandShrink
}

func (s *ExecuteManualNodeShrinkRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteManualNodeShrinkRequest) SetEnv(v string) *ExecuteManualNodeShrinkRequest {
  s.Env = &v
  return s
}

func (s *ExecuteManualNodeShrinkRequest) SetExecuteCommandShrink(v string) *ExecuteManualNodeShrinkRequest {
  s.ExecuteCommandShrink = &v
  return s
}

func (s *ExecuteManualNodeShrinkRequest) SetOpTenantId(v int64) *ExecuteManualNodeShrinkRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteManualNodeShrinkRequest) Validate() error {
  return dara.Validate(s)
}

