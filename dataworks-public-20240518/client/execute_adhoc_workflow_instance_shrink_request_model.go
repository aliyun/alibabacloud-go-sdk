// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdhocWorkflowInstanceShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizDate(v int64) *ExecuteAdhocWorkflowInstanceShrinkRequest
  GetBizDate() *int64 
  SetEnvType(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest
  GetEnvType() *string 
  SetName(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest
  GetName() *string 
  SetOwner(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest
  GetOwner() *string 
  SetProjectId(v int64) *ExecuteAdhocWorkflowInstanceShrinkRequest
  GetProjectId() *int64 
  SetTasksShrink(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest
  GetTasksShrink() *string 
}

type ExecuteAdhocWorkflowInstanceShrinkRequest struct {
  // The data timestamp.
  // 
  // example:
  // 
  // 1710239005403
  BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
  // The environment of the workspace. Valid values:
  // 
  // 	- Prod: production environment
  // 
  // 	- Dev: development environment
  // 
  // example:
  // 
  // Prod
  EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
  // The name of the workflow instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // WorkflowInstance1
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The account ID of the owner.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1000
  Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
  // The workspace ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 100
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The tasks.
  // 
  // This parameter is required.
  TasksShrink *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) GetBizDate() *int64  {
  return s.BizDate
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) GetEnvType() *string  {
  return s.EnvType
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) GetName() *string  {
  return s.Name
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) GetOwner() *string  {
  return s.Owner
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) GetTasksShrink() *string  {
  return s.TasksShrink
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) SetBizDate(v int64) *ExecuteAdhocWorkflowInstanceShrinkRequest {
  s.BizDate = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) SetEnvType(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest {
  s.EnvType = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) SetName(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest {
  s.Name = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) SetOwner(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest {
  s.Owner = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) SetProjectId(v int64) *ExecuteAdhocWorkflowInstanceShrinkRequest {
  s.ProjectId = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) SetTasksShrink(v string) *ExecuteAdhocWorkflowInstanceShrinkRequest {
  s.TasksShrink = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceShrinkRequest) Validate() error {
  return dara.Validate(s)
}

