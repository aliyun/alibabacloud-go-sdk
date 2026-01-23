// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTriggerNodeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizDate(v string) *ExecuteTriggerNodeRequest
  GetBizDate() *string 
  SetEnv(v string) *ExecuteTriggerNodeRequest
  GetEnv() *string 
  SetIndex(v int32) *ExecuteTriggerNodeRequest
  GetIndex() *int32 
  SetNodeId(v string) *ExecuteTriggerNodeRequest
  GetNodeId() *string 
  SetOpTenantId(v int64) *ExecuteTriggerNodeRequest
  GetOpTenantId() *int64 
  SetProjectId(v int64) *ExecuteTriggerNodeRequest
  GetProjectId() *int64 
}

type ExecuteTriggerNodeRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 2025-06-01
  BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
  // example:
  // 
  // DEV
  Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
  // example:
  // 
  // 1
  Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // n_12345678
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 30001011
  OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 10110201
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecuteTriggerNodeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTriggerNodeRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTriggerNodeRequest) GetBizDate() *string  {
  return s.BizDate
}

func (s *ExecuteTriggerNodeRequest) GetEnv() *string  {
  return s.Env
}

func (s *ExecuteTriggerNodeRequest) GetIndex() *int32  {
  return s.Index
}

func (s *ExecuteTriggerNodeRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *ExecuteTriggerNodeRequest) GetOpTenantId() *int64  {
  return s.OpTenantId
}

func (s *ExecuteTriggerNodeRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteTriggerNodeRequest) SetBizDate(v string) *ExecuteTriggerNodeRequest {
  s.BizDate = &v
  return s
}

func (s *ExecuteTriggerNodeRequest) SetEnv(v string) *ExecuteTriggerNodeRequest {
  s.Env = &v
  return s
}

func (s *ExecuteTriggerNodeRequest) SetIndex(v int32) *ExecuteTriggerNodeRequest {
  s.Index = &v
  return s
}

func (s *ExecuteTriggerNodeRequest) SetNodeId(v string) *ExecuteTriggerNodeRequest {
  s.NodeId = &v
  return s
}

func (s *ExecuteTriggerNodeRequest) SetOpTenantId(v int64) *ExecuteTriggerNodeRequest {
  s.OpTenantId = &v
  return s
}

func (s *ExecuteTriggerNodeRequest) SetProjectId(v int64) *ExecuteTriggerNodeRequest {
  s.ProjectId = &v
  return s
}

func (s *ExecuteTriggerNodeRequest) Validate() error {
  return dara.Validate(s)
}

