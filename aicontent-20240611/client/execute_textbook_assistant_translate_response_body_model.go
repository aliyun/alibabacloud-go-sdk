// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantTranslateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantTranslateResponseBodyData) *ExecuteTextbookAssistantTranslateResponseBody
  GetData() *ExecuteTextbookAssistantTranslateResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantTranslateResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantTranslateResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteTextbookAssistantTranslateResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteTextbookAssistantTranslateResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantTranslateResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantTranslateResponseBody struct {
  Data *ExecuteTextbookAssistantTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // 0
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // example:
  // 
  // ""
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 9EB79C1E-36C2-5777-BED6-C23A98DF0637
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) GetData() *ExecuteTextbookAssistantTranslateResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetData(v *ExecuteTextbookAssistantTranslateResponseBodyData) *ExecuteTextbookAssistantTranslateResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantTranslateResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantTranslateResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantTranslateResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantTranslateResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantTranslateResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantTranslateResponseBodyData struct {
  Result *ExecuteTextbookAssistantTranslateResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantTranslateResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyData) GetResult() *ExecuteTextbookAssistantTranslateResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyData) SetResult(v *ExecuteTextbookAssistantTranslateResponseBodyDataResult) *ExecuteTextbookAssistantTranslateResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantTranslateResponseBodyDataResult struct {
  Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyDataResult) GetResult() *string  {
  return s.Result
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyDataResult) SetResult(v string) *ExecuteTextbookAssistantTranslateResponseBodyDataResult {
  s.Result = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

