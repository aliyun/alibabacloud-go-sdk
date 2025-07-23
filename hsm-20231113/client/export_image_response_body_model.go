// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportImageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetJob(v *ExportImageResponseBodyJob) *ExportImageResponseBody
  GetJob() *ExportImageResponseBodyJob 
  SetRequestId(v string) *ExportImageResponseBody
  GetRequestId() *string 
}

type ExportImageResponseBody struct {
  // The information about the asynchronous task returned.
  Job *ExportImageResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
  // The request ID.
  // 
  // example:
  // 
  // 4C467B38-3910-447D-87BC-AC049166F216
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportImageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportImageResponseBody) GoString() string {
  return s.String()
}

func (s *ExportImageResponseBody) GetJob() *ExportImageResponseBodyJob  {
  return s.Job
}

func (s *ExportImageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportImageResponseBody) SetJob(v *ExportImageResponseBodyJob) *ExportImageResponseBody {
  s.Job = v
  return s
}

func (s *ExportImageResponseBody) SetRequestId(v string) *ExportImageResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportImageResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExportImageResponseBodyJob struct {
  // Indicates whether the task is complete. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
  // The error message returned if the operation is abnormal or fails.
  // 
  // example:
  // 
  // Job.Canceled
  Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
  // The ID of the task.
  // 
  // example:
  // 
  // b1748ca6-6b55-49f4-a6d4-2d694a9f3693
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // The progress of the task. Unit: percent (%).
  // 
  // example:
  // 
  // 100
  Process *int32 `json:"Process,omitempty" xml:"Process,omitempty"`
  // The response returned after the operation succeeds.
  // 
  // example:
  // 
  // success
  Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
  // The task status. Valid values:
  // 
  // 	- running
  // 
  // 	- cancel
  // 
  // 	- fail
  // 
  // 	- success
  // 
  // example:
  // 
  // running
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // The type of the task operation. Valid values:
  // 
  // 	- create
  // 
  // 	- cancel
  // 
  // example:
  // 
  // create
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExportImageResponseBodyJob) String() string {
  return dara.Prettify(s)
}

func (s ExportImageResponseBodyJob) GoString() string {
  return s.String()
}

func (s *ExportImageResponseBodyJob) GetCompleted() *bool  {
  return s.Completed
}

func (s *ExportImageResponseBodyJob) GetError() *string  {
  return s.Error
}

func (s *ExportImageResponseBodyJob) GetJobId() *string  {
  return s.JobId
}

func (s *ExportImageResponseBodyJob) GetProcess() *int32  {
  return s.Process
}

func (s *ExportImageResponseBodyJob) GetResponse() *string  {
  return s.Response
}

func (s *ExportImageResponseBodyJob) GetStatus() *string  {
  return s.Status
}

func (s *ExportImageResponseBodyJob) GetType() *string  {
  return s.Type
}

func (s *ExportImageResponseBodyJob) SetCompleted(v bool) *ExportImageResponseBodyJob {
  s.Completed = &v
  return s
}

func (s *ExportImageResponseBodyJob) SetError(v string) *ExportImageResponseBodyJob {
  s.Error = &v
  return s
}

func (s *ExportImageResponseBodyJob) SetJobId(v string) *ExportImageResponseBodyJob {
  s.JobId = &v
  return s
}

func (s *ExportImageResponseBodyJob) SetProcess(v int32) *ExportImageResponseBodyJob {
  s.Process = &v
  return s
}

func (s *ExportImageResponseBodyJob) SetResponse(v string) *ExportImageResponseBodyJob {
  s.Response = &v
  return s
}

func (s *ExportImageResponseBodyJob) SetStatus(v string) *ExportImageResponseBodyJob {
  s.Status = &v
  return s
}

func (s *ExportImageResponseBodyJob) SetType(v string) *ExportImageResponseBodyJob {
  s.Type = &v
  return s
}

func (s *ExportImageResponseBodyJob) Validate() error {
  return dara.Validate(s)
}

