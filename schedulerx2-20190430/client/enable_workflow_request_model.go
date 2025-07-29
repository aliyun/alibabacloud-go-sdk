// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWorkflowRequest interface {
  dara.Model
  String() string
  GoString() string
  SetGroupId(v string) *EnableWorkflowRequest
  GetGroupId() *string 
  SetNamespace(v string) *EnableWorkflowRequest
  GetNamespace() *string 
  SetNamespaceSource(v string) *EnableWorkflowRequest
  GetNamespaceSource() *string 
  SetRegionId(v string) *EnableWorkflowRequest
  GetRegionId() *string 
  SetWorkflowId(v int64) *EnableWorkflowRequest
  GetWorkflowId() *int64 
}

type EnableWorkflowRequest struct {
  // The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
  // 
  // example:
  // 
  // testSchedulerx.defaultGroup
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
  // The region ID.
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

func (s EnableWorkflowRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableWorkflowRequest) GoString() string {
  return s.String()
}

func (s *EnableWorkflowRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *EnableWorkflowRequest) GetNamespace() *string  {
  return s.Namespace
}

func (s *EnableWorkflowRequest) GetNamespaceSource() *string  {
  return s.NamespaceSource
}

func (s *EnableWorkflowRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableWorkflowRequest) GetWorkflowId() *int64  {
  return s.WorkflowId
}

func (s *EnableWorkflowRequest) SetGroupId(v string) *EnableWorkflowRequest {
  s.GroupId = &v
  return s
}

func (s *EnableWorkflowRequest) SetNamespace(v string) *EnableWorkflowRequest {
  s.Namespace = &v
  return s
}

func (s *EnableWorkflowRequest) SetNamespaceSource(v string) *EnableWorkflowRequest {
  s.NamespaceSource = &v
  return s
}

func (s *EnableWorkflowRequest) SetRegionId(v string) *EnableWorkflowRequest {
  s.RegionId = &v
  return s
}

func (s *EnableWorkflowRequest) SetWorkflowId(v int64) *EnableWorkflowRequest {
  s.WorkflowId = &v
  return s
}

func (s *EnableWorkflowRequest) Validate() error {
  return dara.Validate(s)
}

