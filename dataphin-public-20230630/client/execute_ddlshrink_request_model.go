// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDDLShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetContextShrink(v string) *ExecuteDDLShrinkRequest
  GetContextShrink() *string 
  SetDDLCommandShrink(v string) *ExecuteDDLShrinkRequest
  GetDDLCommandShrink() *string 
  SetOpTenantId(v int64) *ExecuteDDLShrinkRequest
  GetOpTenantId() *int64 
}

type ExecuteDDLShrinkRequest struct {
  // The request context information.
  // 
  // This parameter is required.
  ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
  // The one-click table creation parameters.
  // 
  // This parameter is required.
  DDLCommandShrink *string `json:"DDLCommand,omitempty" xml:"DDLCommand,omitempty"`
  // The tenant ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteDDLShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteDDLShrinkRequest) GetContextShrink() *string  {
  return s.ContextShrink
}

func (s *ExecuteDDLShrinkRequest) GetDDLCommandShrink() *string  {
  return s.DDLCommandShrink
}

func (s *ExecuteDDLShrinkRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteDDLShrinkRequest) SetContextShrink(v string) *ExecuteDDLShrinkRequest {
  s.ContextShrink = &v
  return s
}

func (s *ExecuteDDLShrinkRequest) SetDDLCommandShrink(v string) *ExecuteDDLShrinkRequest {
  s.DDLCommandShrink = &v
  return s
}

func (s *ExecuteDDLShrinkRequest) SetOpTenantId(v int64) *ExecuteDDLShrinkRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteDDLShrinkRequest) Validate() error {
  return dara.Validate(s)
}

