// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteWorkflowRequest interface {
  dara.Model
  String() string
  GoString() string
  SetGroupId(v string) *ExecuteWorkflowRequest
  GetGroupId() *string 
  SetInstanceParameters(v string) *ExecuteWorkflowRequest
  GetInstanceParameters() *string 
  SetNamespace(v string) *ExecuteWorkflowRequest
  GetNamespace() *string 
  SetNamespaceSource(v string) *ExecuteWorkflowRequest
  GetNamespaceSource() *string 
  SetRegionId(v string) *ExecuteWorkflowRequest
  GetRegionId() *string 
  SetWorkflowId(v int64) *ExecuteWorkflowRequest
  GetWorkflowId() *int64 
}

type ExecuteWorkflowRequest struct {
  // The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testSchedulerx.defaultGroup
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The dynamic parameter of the workflow instance. The value of the parameter can be up to 1,000 bytes in length.
  // 
  // example:
  // 
  // test
  InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
  // The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
  Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
  // The source of the namespace. This parameter is required only for a special third party.
  // 
  // example:
  // 
  // schedulerx
  NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
  // The region information.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The workflow ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 111
  WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ExecuteWorkflowRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteWorkflowRequest) GoString() string {
  return s.String()
}

func (s *ExecuteWorkflowRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *ExecuteWorkflowRequest) GetInstanceParameters() *string  {
  return s.InstanceParameters
}

func (s *ExecuteWorkflowRequest) GetNamespace() *string  {
  return s.Namespace
}

func (s *ExecuteWorkflowRequest) GetNamespaceSource() *string  {
  return s.NamespaceSource
}

func (s *ExecuteWorkflowRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteWorkflowRequest) GetWorkflowId() *int64  {
  return s.WorkflowId
}

func (s *ExecuteWorkflowRequest) SetGroupId(v string) *ExecuteWorkflowRequest {
  s.GroupId = &v
  return s
}

func (s *ExecuteWorkflowRequest) SetInstanceParameters(v string) *ExecuteWorkflowRequest {
  s.InstanceParameters = &v
  return s
}

func (s *ExecuteWorkflowRequest) SetNamespace(v string) *ExecuteWorkflowRequest {
  s.Namespace = &v
  return s
}

func (s *ExecuteWorkflowRequest) SetNamespaceSource(v string) *ExecuteWorkflowRequest {
  s.NamespaceSource = &v
  return s
}

func (s *ExecuteWorkflowRequest) SetRegionId(v string) *ExecuteWorkflowRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteWorkflowRequest) SetWorkflowId(v int64) *ExecuteWorkflowRequest {
  s.WorkflowId = &v
  return s
}

func (s *ExecuteWorkflowRequest) Validate() error {
  return dara.Validate(s)
}

