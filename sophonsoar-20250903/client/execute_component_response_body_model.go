// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteComponentResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteComponentResponseBody
  GetRequestId() *string 
  SetRunResult(v string) *ExecuteComponentResponseBody
  GetRunResult() *string 
}

type ExecuteComponentResponseBody struct {
  // example:
  // 
  // 10B92EE1-4597-593B-A131-7A17D25EF5C9
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // {
  // 
  //     "requestUuid": "fe240b98-27b1-4a36-aec1-550b894318d9",
  // 
  //     "content": {
  // 
  //         "resultData": [],
  // 
  //         "success": true
  // 
  //     }
  // 
  // }
  RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s ExecuteComponentResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteComponentResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteComponentResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteComponentResponseBody) GetRunResult() *string  {
  return s.RunResult
}

func (s *ExecuteComponentResponseBody) SetRequestId(v string) *ExecuteComponentResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteComponentResponseBody) SetRunResult(v string) *ExecuteComponentResponseBody {
  s.RunResult = &v
  return s
}

func (s *ExecuteComponentResponseBody) Validate() error {
  return dara.Validate(s)
}

