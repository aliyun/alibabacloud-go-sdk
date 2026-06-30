// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAgentResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteAgentResponseBody
  GetCode() *string 
  SetData(v *ExecuteAgentResponseBodyData) *ExecuteAgentResponseBody
  GetData() *ExecuteAgentResponseBodyData 
  SetMessage(v string) *ExecuteAgentResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteAgentResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteAgentResponseBody
  GetSuccess() *bool 
}

type ExecuteAgentResponseBody struct {
  // The status code. A value of 200 indicates success.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned result.
  Data *ExecuteAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The error message returned when an error occurs.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // F190ADE9-619A-447D-84E3-7E241A5C428E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. The caller can use this field to determine whether the request was successful. Valid values: **true**: The request was successful. **false/null**: The request failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteAgentResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAgentResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAgentResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteAgentResponseBody) GetData() *ExecuteAgentResponseBodyData  {
  return s.Data
}

func (s *ExecuteAgentResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteAgentResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAgentResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteAgentResponseBody) SetCode(v string) *ExecuteAgentResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteAgentResponseBody) SetData(v *ExecuteAgentResponseBodyData) *ExecuteAgentResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteAgentResponseBody) SetMessage(v string) *ExecuteAgentResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteAgentResponseBody) SetRequestId(v string) *ExecuteAgentResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAgentResponseBody) SetSuccess(v bool) *ExecuteAgentResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteAgentResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAgentResponseBodyData struct {
  // If streaming output is used, this value is null during generation. When generation is complete, the value is stop if the generation ended due to a stop token.
  // 
  // example:
  // 
  // stop
  FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
  // The number of input tokens.
  // 
  // example:
  // 
  // 100
  InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
  // The request ID returned by the large language model service.
  // 
  // example:
  // 
  // 106C6CA0-282D-4AF7-85F0-D2D24***
  LlmRequestId *string `json:"LlmRequestId,omitempty" xml:"LlmRequestId,omitempty"`
  // The number of output tokens.
  // 
  // example:
  // 
  // 200
  OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
  // The result returned by the large language model.
  // 
  // example:
  // 
  // 这段对话似乎是客服与客户之间关于一个服务或产品的讨论，但具体内容难以明确理解，因为对话中的言语比较零散和抽象。
  Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
  // The total number of tokens.
  // 
  // example:
  // 
  // 300
  TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
  // The number of times the plus model was used.
  // 
  // example:
  // 
  // 1
  TyxmPlusCount *string `json:"TyxmPlusCount,omitempty" xml:"TyxmPlusCount,omitempty"`
  // The number of times the turbo model was used.
  // 
  // example:
  // 
  // 1
  TyxmTurboCount *string `json:"TyxmTurboCount,omitempty" xml:"TyxmTurboCount,omitempty"`
}

func (s ExecuteAgentResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAgentResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteAgentResponseBodyData) GetFinishReason() *string  {
  return s.FinishReason
}

func (s *ExecuteAgentResponseBodyData) GetInputTokens() *int64  {
  return s.InputTokens
}

func (s *ExecuteAgentResponseBodyData) GetLlmRequestId() *string  {
  return s.LlmRequestId
}

func (s *ExecuteAgentResponseBodyData) GetOutputTokens() *int64  {
  return s.OutputTokens
}

func (s *ExecuteAgentResponseBodyData) GetText() *string  {
  return s.Text
}

func (s *ExecuteAgentResponseBodyData) GetTotalTokens() *int64  {
  return s.TotalTokens
}

func (s *ExecuteAgentResponseBodyData) GetTyxmPlusCount() *string  {
  return s.TyxmPlusCount
}

func (s *ExecuteAgentResponseBodyData) GetTyxmTurboCount() *string  {
  return s.TyxmTurboCount
}

func (s *ExecuteAgentResponseBodyData) SetFinishReason(v string) *ExecuteAgentResponseBodyData {
  s.FinishReason = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetInputTokens(v int64) *ExecuteAgentResponseBodyData {
  s.InputTokens = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetLlmRequestId(v string) *ExecuteAgentResponseBodyData {
  s.LlmRequestId = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetOutputTokens(v int64) *ExecuteAgentResponseBodyData {
  s.OutputTokens = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetText(v string) *ExecuteAgentResponseBodyData {
  s.Text = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetTotalTokens(v int64) *ExecuteAgentResponseBodyData {
  s.TotalTokens = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetTyxmPlusCount(v string) *ExecuteAgentResponseBodyData {
  s.TyxmPlusCount = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) SetTyxmTurboCount(v string) *ExecuteAgentResponseBodyData {
  s.TyxmTurboCount = &v
  return s
}

func (s *ExecuteAgentResponseBodyData) Validate() error {
  return dara.Validate(s)
}

