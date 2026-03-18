// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantDifficultyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantDifficultyResponseBodyData) *ExecuteTextbookAssistantDifficultyResponseBody
  GetData() *ExecuteTextbookAssistantDifficultyResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantDifficultyResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantDifficultyResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteTextbookAssistantDifficultyResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteTextbookAssistantDifficultyResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantDifficultyResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantDifficultyResponseBody struct {
  Data *ExecuteTextbookAssistantDifficultyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // null
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // example:
  // 
  // null
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 0D7D382F-9475-572E-BE83-DDFBF5C5EB24
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) GetData() *ExecuteTextbookAssistantDifficultyResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetData(v *ExecuteTextbookAssistantDifficultyResponseBodyData) *ExecuteTextbookAssistantDifficultyResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantDifficultyResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantDifficultyResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantDifficultyResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantDifficultyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantDifficultyResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantDifficultyResponseBodyData struct {
  Result *ExecuteTextbookAssistantDifficultyResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyData) GetResult() *ExecuteTextbookAssistantDifficultyResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyData) SetResult(v *ExecuteTextbookAssistantDifficultyResponseBodyDataResult) *ExecuteTextbookAssistantDifficultyResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantDifficultyResponseBodyDataResult struct {
  // example:
  // 
  // Let\\"s look at the text again. Mike says, \\"I\\"m Mike Black.\\" Can you try saying it like Mike?
  Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyDataResult) GetResult() *string  {
  return s.Result
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyDataResult) SetResult(v string) *ExecuteTextbookAssistantDifficultyResponseBodyDataResult {
  s.Result = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

