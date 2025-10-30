// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdHocTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteAdHocTaskResponseBody
  GetCode() *string 
  SetExecuteResult(v *ExecuteAdHocTaskResponseBodyExecuteResult) *ExecuteAdHocTaskResponseBody
  GetExecuteResult() *ExecuteAdHocTaskResponseBodyExecuteResult 
  SetHttpStatusCode(v int32) *ExecuteAdHocTaskResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecuteAdHocTaskResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteAdHocTaskResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAdHocTaskResponseBody
  GetSuccess() *bool 
}

type ExecuteAdHocTaskResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  ExecuteResult *ExecuteAdHocTaskResponseBodyExecuteResult `json:"ExecuteResult,omitempty" xml:"ExecuteResult,omitempty" type:"Struct"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteAdHocTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteAdHocTaskResponseBody) GetExecuteResult() *ExecuteAdHocTaskResponseBodyExecuteResult  {
  return s.ExecuteResult
}

func (s *ExecuteAdHocTaskResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteAdHocTaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteAdHocTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAdHocTaskResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAdHocTaskResponseBody) SetCode(v string) *ExecuteAdHocTaskResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBody) SetExecuteResult(v *ExecuteAdHocTaskResponseBodyExecuteResult) *ExecuteAdHocTaskResponseBody {
  s.ExecuteResult = v
  return s
}

func (s *ExecuteAdHocTaskResponseBody) SetHttpStatusCode(v int32) *ExecuteAdHocTaskResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBody) SetMessage(v string) *ExecuteAdHocTaskResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBody) SetRequestId(v string) *ExecuteAdHocTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBody) SetSuccess(v bool) *ExecuteAdHocTaskResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBody) Validate() error {
  if s.ExecuteResult != nil {
    if err := s.ExecuteResult.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAdHocTaskResponseBodyExecuteResult struct {
  // example:
  // 
  // 1
  SubTaskCount *int32 `json:"SubTaskCount,omitempty" xml:"SubTaskCount,omitempty"`
  // example:
  // 
  // MaxCompute_SQL_300000843_1611548758327
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExecuteAdHocTaskResponseBodyExecuteResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskResponseBodyExecuteResult) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskResponseBodyExecuteResult) GetSubTaskCount() *int32  {
  return s.SubTaskCount
}

func (s *ExecuteAdHocTaskResponseBodyExecuteResult) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExecuteAdHocTaskResponseBodyExecuteResult) SetSubTaskCount(v int32) *ExecuteAdHocTaskResponseBodyExecuteResult {
  s.SubTaskCount = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBodyExecuteResult) SetTaskId(v string) *ExecuteAdHocTaskResponseBodyExecuteResult {
  s.TaskId = &v
  return s
}

func (s *ExecuteAdHocTaskResponseBodyExecuteResult) Validate() error {
  return dara.Validate(s)
}

