// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantGrammarCheckResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteTextbookAssistantGrammarCheckResponseBodyData) *ExecuteTextbookAssistantGrammarCheckResponseBody
  GetData() *ExecuteTextbookAssistantGrammarCheckResponseBodyData 
  SetErrCode(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteTextbookAssistantGrammarCheckResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTextbookAssistantGrammarCheckResponseBody
  GetSuccess() *bool 
}

type ExecuteTextbookAssistantGrammarCheckResponseBody struct {
  Data *ExecuteTextbookAssistantGrammarCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 0bc1e96d17091734639835114e12c8
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) GetData() *ExecuteTextbookAssistantGrammarCheckResponseBodyData  {
  return s.Data
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetData(v *ExecuteTextbookAssistantGrammarCheckResponseBodyData) *ExecuteTextbookAssistantGrammarCheckResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantGrammarCheckResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantGrammarCheckResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantGrammarCheckResponseBodyData struct {
  Result *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyData) GetResult() *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult  {
  return s.Result
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyData) SetResult(v *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) *ExecuteTextbookAssistantGrammarCheckResponseBodyData {
  s.Result = v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyData) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult struct {
  Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
  // example:
  // 
  // I am you.
  Correction *string `json:"correction,omitempty" xml:"correction,omitempty"`
  // example:
  // 
  // Has_Error
  CorrectionStatus *string `json:"correctionStatus,omitempty" xml:"correctionStatus,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) GetAnalysis() *string  {
  return s.Analysis
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) GetCorrection() *string  {
  return s.Correction
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) GetCorrectionStatus() *string  {
  return s.CorrectionStatus
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) SetAnalysis(v string) *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult {
  s.Analysis = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) SetCorrection(v string) *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult {
  s.Correction = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) SetCorrectionStatus(v string) *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult {
  s.CorrectionStatus = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) Validate() error {
  return dara.Validate(s)
}

