// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableJobRequest interface {
  dara.Model
  String() string
  GoString() string
  SetGroupId(v string) *EnableJobRequest
  GetGroupId() *string 
  SetJobId(v int64) *EnableJobRequest
  GetJobId() *int64 
  SetNamespace(v string) *EnableJobRequest
  GetNamespace() *string 
  SetNamespaceSource(v string) *EnableJobRequest
  GetNamespaceSource() *string 
  SetRegionId(v string) *EnableJobRequest
  GetRegionId() *string 
}

type EnableJobRequest struct {
  // The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
  // 
  // example:
  // 
  // testSchedulerx.defaultGroup
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 92583
  JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
}

func (s EnableJobRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableJobRequest) GoString() string {
  return s.String()
}

func (s *EnableJobRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *EnableJobRequest) GetJobId() *int64  {
  return s.JobId
}

func (s *EnableJobRequest) GetNamespace() *string  {
  return s.Namespace
}

func (s *EnableJobRequest) GetNamespaceSource() *string  {
  return s.NamespaceSource
}

func (s *EnableJobRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableJobRequest) SetGroupId(v string) *EnableJobRequest {
  s.GroupId = &v
  return s
}

func (s *EnableJobRequest) SetJobId(v int64) *EnableJobRequest {
  s.JobId = &v
  return s
}

func (s *EnableJobRequest) SetNamespace(v string) *EnableJobRequest {
  s.Namespace = &v
  return s
}

func (s *EnableJobRequest) SetNamespaceSource(v string) *EnableJobRequest {
  s.NamespaceSource = &v
  return s
}

func (s *EnableJobRequest) SetRegionId(v string) *EnableJobRequest {
  s.RegionId = &v
  return s
}

func (s *EnableJobRequest) Validate() error {
  return dara.Validate(s)
}

