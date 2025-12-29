// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSceneResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteSceneResponseBody
  GetCode() *int32 
  SetMessage(v string) *ExecuteSceneResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteSceneResponseBody
  GetRequestId() *string 
  SetResult(v bool) *ExecuteSceneResponseBody
  GetResult() *bool 
  SetStatusCode(v int32) *ExecuteSceneResponseBody
  GetStatusCode() *int32 
}

type ExecuteSceneResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 191C79AD-F9F9-531E-B8C1-73DF6433B920
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
  // example:
  // 
  // 200
  StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ExecuteSceneResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSceneResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSceneResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteSceneResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteSceneResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSceneResponseBody) GetResult() *bool  {
  return s.Result
}

func (s *ExecuteSceneResponseBody) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSceneResponseBody) SetCode(v int32) *ExecuteSceneResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteSceneResponseBody) SetMessage(v string) *ExecuteSceneResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteSceneResponseBody) SetRequestId(v string) *ExecuteSceneResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSceneResponseBody) SetResult(v bool) *ExecuteSceneResponseBody {
  s.Result = &v
  return s
}

func (s *ExecuteSceneResponseBody) SetStatusCode(v int32) *ExecuteSceneResponseBody {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSceneResponseBody) Validate() error {
  return dara.Validate(s)
}

