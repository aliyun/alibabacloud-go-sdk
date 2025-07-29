// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteJobResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteJobResponseBody
  GetCode() *int32 
  SetData(v *ExecuteJobResponseBodyData) *ExecuteJobResponseBody
  GetData() *ExecuteJobResponseBodyData 
  SetMessage(v string) *ExecuteJobResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteJobResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteJobResponseBody
  GetSuccess() *bool 
}

type ExecuteJobResponseBody struct {
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The ID of the job instance that is returned if the request is successful.
  Data *ExecuteJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The error message that is returned only if the corresponding error occurs.
  // 
  // example:
  // 
  // groupid not exist groupId: testSchedulerx.defaultGroup namespace: adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 4F68ABED-AC31-4412-9297-D9A8F0401108****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- `true`
  // 
  // 	- `false`
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteJobResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteJobResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteJobResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteJobResponseBody) GetData() *ExecuteJobResponseBodyData  {
  return s.Data
}

func (s *ExecuteJobResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteJobResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteJobResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteJobResponseBody) SetCode(v int32) *ExecuteJobResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteJobResponseBody) SetData(v *ExecuteJobResponseBodyData) *ExecuteJobResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteJobResponseBody) SetMessage(v string) *ExecuteJobResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteJobResponseBody) SetRequestId(v string) *ExecuteJobResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteJobResponseBody) SetSuccess(v bool) *ExecuteJobResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteJobResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteJobResponseBodyData struct {
  // The job instance ID.
  // 
  // example:
  // 
  // 11111111
  JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
}

func (s ExecuteJobResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteJobResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteJobResponseBodyData) GetJobInstanceId() *int64  {
  return s.JobInstanceId
}

func (s *ExecuteJobResponseBodyData) SetJobInstanceId(v int64) *ExecuteJobResponseBodyData {
  s.JobInstanceId = &v
  return s
}

func (s *ExecuteJobResponseBodyData) Validate() error {
  return dara.Validate(s)
}

