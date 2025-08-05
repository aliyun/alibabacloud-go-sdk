// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePolicyV2ResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecutePolicyV2ResponseBody
  GetCode() *string 
  SetJobId(v string) *ExecutePolicyV2ResponseBody
  GetJobId() *string 
  SetMessage(v string) *ExecutePolicyV2ResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecutePolicyV2ResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecutePolicyV2ResponseBody
  GetSuccess() *bool 
}

type ExecutePolicyV2ResponseBody struct {
  // Return code, 200 indicates success.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Backup job ID.
  // 
  // example:
  // 
  // job-*********************
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // Description of the return message, usually returns \\"successful\\" on success, and corresponding error messages on failure.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Request ID.
  // 
  // example:
  // 
  // F4EEB401-DD21-588D-AE3B-1E835C7655E1
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // - true: Success
  // 
  // - false: Failure
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecutePolicyV2ResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecutePolicyV2ResponseBody) GoString() string {
  return s.String()
}

func (s *ExecutePolicyV2ResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecutePolicyV2ResponseBody) GetJobId() *string  {
  return s.JobId
}

func (s *ExecutePolicyV2ResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecutePolicyV2ResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecutePolicyV2ResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecutePolicyV2ResponseBody) SetCode(v string) *ExecutePolicyV2ResponseBody {
  s.Code = &v
  return s
}

func (s *ExecutePolicyV2ResponseBody) SetJobId(v string) *ExecutePolicyV2ResponseBody {
  s.JobId = &v
  return s
}

func (s *ExecutePolicyV2ResponseBody) SetMessage(v string) *ExecutePolicyV2ResponseBody {
  s.Message = &v
  return s
}

func (s *ExecutePolicyV2ResponseBody) SetRequestId(v string) *ExecutePolicyV2ResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecutePolicyV2ResponseBody) SetSuccess(v bool) *ExecutePolicyV2ResponseBody {
  s.Success = &v
  return s
}

func (s *ExecutePolicyV2ResponseBody) Validate() error {
  return dara.Validate(s)
}

