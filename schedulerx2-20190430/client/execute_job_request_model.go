// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteJobRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCheckJobStatus(v bool) *ExecuteJobRequest
  GetCheckJobStatus() *bool 
  SetDesignateType(v int32) *ExecuteJobRequest
  GetDesignateType() *int32 
  SetGroupId(v string) *ExecuteJobRequest
  GetGroupId() *string 
  SetInstanceParameters(v string) *ExecuteJobRequest
  GetInstanceParameters() *string 
  SetJobId(v int64) *ExecuteJobRequest
  GetJobId() *int64 
  SetLabel(v string) *ExecuteJobRequest
  GetLabel() *string 
  SetNamespace(v string) *ExecuteJobRequest
  GetNamespace() *string 
  SetNamespaceSource(v string) *ExecuteJobRequest
  GetNamespaceSource() *string 
  SetRegionId(v string) *ExecuteJobRequest
  GetRegionId() *string 
  SetWorker(v string) *ExecuteJobRequest
  GetWorker() *string 
}

type ExecuteJobRequest struct {
  // Specifies whether to check the job status. Valid values: -**true**: The job can be run only if the job is enabled. -**false**: The job can be run even if the job is disabled.
  // 
  // example:
  // 
  // true
  CheckJobStatus *bool `json:"CheckJobStatus,omitempty" xml:"CheckJobStatus,omitempty"`
  // The type of the designated machine. Valid values: -**1**: worker. -**2**: label.
  // 
  // example:
  // 
  // 1
  DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
  // The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testSchedulerx.defaultGroup
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The parameters that are passed to trigger the job to run. The input value can be a random string. The parameters that are passed are obtained by calling the `context.getInstanceParameters()` class in the `processor` code. The parameters are different from custom parameters for creating jobs.
  // 
  // example:
  // 
  // test
  InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
  // The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 92583
  JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // The label of the worker.
  // 
  // example:
  // 
  // gray
  Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
  // The worker address of the application. To query the worker address, call the GetWokerList operation.
  // 
  // example:
  // 
  // xxxxxxx@127.0.0.1:222
  Worker *string `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s ExecuteJobRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteJobRequest) GoString() string {
  return s.String()
}

func (s *ExecuteJobRequest) GetCheckJobStatus() *bool  {
  return s.CheckJobStatus
}

func (s *ExecuteJobRequest) GetDesignateType() *int32  {
  return s.DesignateType
}

func (s *ExecuteJobRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *ExecuteJobRequest) GetInstanceParameters() *string  {
  return s.InstanceParameters
}

func (s *ExecuteJobRequest) GetJobId() *int64  {
  return s.JobId
}

func (s *ExecuteJobRequest) GetLabel() *string  {
  return s.Label
}

func (s *ExecuteJobRequest) GetNamespace() *string  {
  return s.Namespace
}

func (s *ExecuteJobRequest) GetNamespaceSource() *string  {
  return s.NamespaceSource
}

func (s *ExecuteJobRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteJobRequest) GetWorker() *string  {
  return s.Worker
}

func (s *ExecuteJobRequest) SetCheckJobStatus(v bool) *ExecuteJobRequest {
  s.CheckJobStatus = &v
  return s
}

func (s *ExecuteJobRequest) SetDesignateType(v int32) *ExecuteJobRequest {
  s.DesignateType = &v
  return s
}

func (s *ExecuteJobRequest) SetGroupId(v string) *ExecuteJobRequest {
  s.GroupId = &v
  return s
}

func (s *ExecuteJobRequest) SetInstanceParameters(v string) *ExecuteJobRequest {
  s.InstanceParameters = &v
  return s
}

func (s *ExecuteJobRequest) SetJobId(v int64) *ExecuteJobRequest {
  s.JobId = &v
  return s
}

func (s *ExecuteJobRequest) SetLabel(v string) *ExecuteJobRequest {
  s.Label = &v
  return s
}

func (s *ExecuteJobRequest) SetNamespace(v string) *ExecuteJobRequest {
  s.Namespace = &v
  return s
}

func (s *ExecuteJobRequest) SetNamespaceSource(v string) *ExecuteJobRequest {
  s.NamespaceSource = &v
  return s
}

func (s *ExecuteJobRequest) SetRegionId(v string) *ExecuteJobRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteJobRequest) SetWorker(v string) *ExecuteJobRequest {
  s.Worker = &v
  return s
}

func (s *ExecuteJobRequest) Validate() error {
  return dara.Validate(s)
}

