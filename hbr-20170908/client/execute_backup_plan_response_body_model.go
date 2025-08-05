// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBackupPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteBackupPlanResponseBody
  GetCode() *string 
  SetJobId(v string) *ExecuteBackupPlanResponseBody
  GetJobId() *string 
  SetMessage(v string) *ExecuteBackupPlanResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteBackupPlanResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteBackupPlanResponseBody
  GetSuccess() *bool 
}

type ExecuteBackupPlanResponseBody struct {
  // The response code. The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The ID of the backup job.
  // 
  // example:
  // 
  // job-*********************
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteBackupPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBackupPlanResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteBackupPlanResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteBackupPlanResponseBody) GetJobId() *string  {
  return s.JobId
}

func (s *ExecuteBackupPlanResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteBackupPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteBackupPlanResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteBackupPlanResponseBody) SetCode(v string) *ExecuteBackupPlanResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteBackupPlanResponseBody) SetJobId(v string) *ExecuteBackupPlanResponseBody {
  s.JobId = &v
  return s
}

func (s *ExecuteBackupPlanResponseBody) SetMessage(v string) *ExecuteBackupPlanResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteBackupPlanResponseBody) SetRequestId(v string) *ExecuteBackupPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteBackupPlanResponseBody) SetSuccess(v bool) *ExecuteBackupPlanResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteBackupPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

