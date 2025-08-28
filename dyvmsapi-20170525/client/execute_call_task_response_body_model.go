// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteCallTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteCallTaskResponseBody
  GetCode() *string 
  SetData(v bool) *ExecuteCallTaskResponseBody
  GetData() *bool 
  SetRequestId(v string) *ExecuteCallTaskResponseBody
  GetRequestId() *string 
}

type ExecuteCallTaskResponseBody struct {
  // The response code.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**: The request was successful.
  // 
  // 	- **false**: The request failed.
  // 
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 53D0F0Fe-cbbB-De28-6FCd-DdbBcefA46dD
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteCallTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteCallTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteCallTaskResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteCallTaskResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExecuteCallTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteCallTaskResponseBody) SetCode(v string) *ExecuteCallTaskResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteCallTaskResponseBody) SetData(v bool) *ExecuteCallTaskResponseBody {
  s.Data = &v
  return s
}

func (s *ExecuteCallTaskResponseBody) SetRequestId(v string) *ExecuteCallTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteCallTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

