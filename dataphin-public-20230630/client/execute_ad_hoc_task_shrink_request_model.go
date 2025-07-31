// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdHocTaskShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecuteCommandShrink(v string) *ExecuteAdHocTaskShrinkRequest
  GetExecuteCommandShrink() *string 
  SetOpTenantId(v int64) *ExecuteAdHocTaskShrinkRequest
  GetOpTenantId() *int64 
}

type ExecuteAdHocTaskShrinkRequest struct {
  // This parameter is required.
  ExecuteCommandShrink *string `json:"ExecuteCommand,omitempty" xml:"ExecuteCommand,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteAdHocTaskShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskShrinkRequest) GetExecuteCommandShrink() *string  {
  return s.ExecuteCommandShrink
}

func (s *ExecuteAdHocTaskShrinkRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteAdHocTaskShrinkRequest) SetExecuteCommandShrink(v string) *ExecuteAdHocTaskShrinkRequest {
  s.ExecuteCommandShrink = &v
  return s
}

func (s *ExecuteAdHocTaskShrinkRequest) SetOpTenantId(v int64) *ExecuteAdHocTaskShrinkRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteAdHocTaskShrinkRequest) Validate() error {
  return dara.Validate(s)
}

