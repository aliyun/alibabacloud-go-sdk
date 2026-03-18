// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantSuggestionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantSuggestionResponseBodyData) *ExecuteTextbookAssistantSuggestionResponseBody
  GetData() *ExecuteTextbookAssistantSuggestionResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantSuggestionResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantSuggestionResponseBody
  GetErrMessage() *string 
  SetHttpstatusCode(v int32) *ExecuteTextbookAssistantSuggestionResponseBody
  GetHttpstatusCode() *int32 
  SetRequestId(v string) *ExecuteTextbookAssistantSuggestionResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantSuggestionResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantSuggestionResponseBody struct {
  Data *ExecuteTextbookAssistantSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // 0
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // example:
  // 
  // null
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // 200
  HttpstatusCode *int32 `json:"httpstatusCode,omitempty" xml:"httpstatusCode,omitempty"`
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

func (s ExecuteTextbookAssistantSuggestionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) GetData() *ExecuteTextbookAssistantSuggestionResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) GetHttpstatusCode() *int32  {
  return s.HttpstatusCode
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetData(v *ExecuteTextbookAssistantSuggestionResponseBodyData) *ExecuteTextbookAssistantSuggestionResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantSuggestionResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantSuggestionResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetHttpstatusCode(v int32) *ExecuteTextbookAssistantSuggestionResponseBody {
  s.HttpstatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantSuggestionResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantSuggestionResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantSuggestionResponseBodyData struct {
  Result *ExecuteTextbookAssistantSuggestionResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyData) GetResult() *ExecuteTextbookAssistantSuggestionResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyData) SetResult(v *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) *ExecuteTextbookAssistantSuggestionResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantSuggestionResponseBodyDataResult struct {
  ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
  // example:
  // 
  // Good evening! From the book, how does Mike Black introduce himself?
  EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) GetChineseResult() *string  {
  return s.ChineseResult
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) GetEnglishResult() *string  {
  return s.EnglishResult
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantSuggestionResponseBodyDataResult {
  s.ChineseResult = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantSuggestionResponseBodyDataResult {
  s.EnglishResult = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

