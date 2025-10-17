// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditWorkspaceQueueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnvironments(v []*string) *EditWorkspaceQueueRequest
  GetEnvironments() []*string 
  SetResourceSpec(v *EditWorkspaceQueueRequestResourceSpec) *EditWorkspaceQueueRequest
  GetResourceSpec() *EditWorkspaceQueueRequestResourceSpec 
  SetWorkspaceId(v string) *EditWorkspaceQueueRequest
  GetWorkspaceId() *string 
  SetWorkspaceQueueName(v string) *EditWorkspaceQueueRequest
  GetWorkspaceQueueName() *string 
  SetRegionId(v string) *EditWorkspaceQueueRequest
  GetRegionId() *string 
}

type EditWorkspaceQueueRequest struct {
  Environments []*string `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
  ResourceSpec *EditWorkspaceQueueRequestResourceSpec `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
  // example:
  // 
  // w-975bcfda9625****
  WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
  // example:
  // 
  // dev_queue
  WorkspaceQueueName *string `json:"workspaceQueueName,omitempty" xml:"workspaceQueueName,omitempty"`
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s EditWorkspaceQueueRequest) String() string {
  return dara.Prettify(s)
}

func (s EditWorkspaceQueueRequest) GoString() string {
  return s.String()
}

func (s *EditWorkspaceQueueRequest) GetEnvironments() []*string  {
  return s.Environments
}

func (s *EditWorkspaceQueueRequest) GetResourceSpec() *EditWorkspaceQueueRequestResourceSpec  {
  return s.ResourceSpec
}

func (s *EditWorkspaceQueueRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *EditWorkspaceQueueRequest) GetWorkspaceQueueName() *string  {
  return s.WorkspaceQueueName
}

func (s *EditWorkspaceQueueRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EditWorkspaceQueueRequest) SetEnvironments(v []*string) *EditWorkspaceQueueRequest {
  s.Environments = v
  return s
}

func (s *EditWorkspaceQueueRequest) SetResourceSpec(v *EditWorkspaceQueueRequestResourceSpec) *EditWorkspaceQueueRequest {
  s.ResourceSpec = v
  return s
}

func (s *EditWorkspaceQueueRequest) SetWorkspaceId(v string) *EditWorkspaceQueueRequest {
  s.WorkspaceId = &v
  return s
}

func (s *EditWorkspaceQueueRequest) SetWorkspaceQueueName(v string) *EditWorkspaceQueueRequest {
  s.WorkspaceQueueName = &v
  return s
}

func (s *EditWorkspaceQueueRequest) SetRegionId(v string) *EditWorkspaceQueueRequest {
  s.RegionId = &v
  return s
}

func (s *EditWorkspaceQueueRequest) Validate() error {
  if s.ResourceSpec != nil {
    if err := s.ResourceSpec.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EditWorkspaceQueueRequestResourceSpec struct {
  // example:
  // 
  // 1000
  Cu *int64 `json:"cu,omitempty" xml:"cu,omitempty"`
}

func (s EditWorkspaceQueueRequestResourceSpec) String() string {
  return dara.Prettify(s)
}

func (s EditWorkspaceQueueRequestResourceSpec) GoString() string {
  return s.String()
}

func (s *EditWorkspaceQueueRequestResourceSpec) GetCu() *int64  {
  return s.Cu
}

func (s *EditWorkspaceQueueRequestResourceSpec) SetCu(v int64) *EditWorkspaceQueueRequestResourceSpec {
  s.Cu = &v
  return s
}

func (s *EditWorkspaceQueueRequestResourceSpec) Validate() error {
  return dara.Validate(s)
}

