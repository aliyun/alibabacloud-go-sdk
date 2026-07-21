// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantRefineByContextResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantRefineByContextResponseBodyData) *ExecuteTextbookAssistantRefineByContextResponseBody
  GetData() *ExecuteTextbookAssistantRefineByContextResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantRefineByContextResponseBody
  GetErrCode() *string 
  SetErrMessage(v int32) *ExecuteTextbookAssistantRefineByContextResponseBody
  GetErrMessage() *int32 
  SetHttpStatusCode(v string) *ExecuteTextbookAssistantRefineByContextResponseBody
  GetHttpStatusCode() *string 
  SetRequestId(v string) *ExecuteTextbookAssistantRefineByContextResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantRefineByContextResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantRefineByContextResponseBody struct {
  // The returned data object.
  Data *ExecuteTextbookAssistantRefineByContextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // The error code returned when the request fails.
  // 
  // example:
  // 
  // 0
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // The error message returned when the request fails.
  // 
  // example:
  // 
  // null
  ErrMessage *int32 `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // The unique request ID.
  // 
  // example:
  // 
  // 6F73C114-A76E-51AD-99E3-BC7B941B69E0
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // A value of `true` indicates that the request succeeded.
  // 
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) GetData() *ExecuteTextbookAssistantRefineByContextResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) GetErrMessage() *int32  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetData(v *ExecuteTextbookAssistantRefineByContextResponseBodyData) *ExecuteTextbookAssistantRefineByContextResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantRefineByContextResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetErrMessage(v int32) *ExecuteTextbookAssistantRefineByContextResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetHttpStatusCode(v string) *ExecuteTextbookAssistantRefineByContextResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantRefineByContextResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantRefineByContextResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantRefineByContextResponseBodyData struct {
  // The result data.
  Result *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyData) GetResult() *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyData) SetResult(v *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) *ExecuteTextbookAssistantRefineByContextResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantRefineByContextResponseBodyDataResult struct {
  // The refined sentence.
  // 
  // example:
  // 
  // Good evening! From the book, how does Mike Black introduce himself?
  Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) GetResult() *string  {
  return s.Result
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) SetResult(v string) *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult {
  s.Result = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

