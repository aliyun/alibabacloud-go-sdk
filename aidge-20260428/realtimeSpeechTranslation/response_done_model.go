// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseDone
  GetType() *string 
  SetEventId(v string) *ResponseDone
  GetEventId() *string 
  SetResponse(v *ResponseDoneResponse) *ResponseDone
  GetResponse() *ResponseDoneResponse 
}

type ResponseDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  Response *ResponseDoneResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
}

func (s ResponseDone) String() string {
  return dara.Prettify(s)
}

func (s ResponseDone) GoString() string {
  return s.String()
}

func (s *ResponseDone) GetType() *string  {
  return s.Type
}

func (s *ResponseDone) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseDone) GetResponse() *ResponseDoneResponse  {
  return s.Response
}

func (s *ResponseDone) SetType(v string) *ResponseDone {
  s.Type = &v
  return s
}

func (s *ResponseDone) SetEventId(v string) *ResponseDone {
  s.EventId = &v
  return s
}

func (s *ResponseDone) SetResponse(v *ResponseDoneResponse) *ResponseDone {
  s.Response = v
  return s
}

func (s *ResponseDone) Validate() error {
  if s.Response != nil {
    if err := s.Response.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseDoneResponse struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  ConversationId *string `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
  OutputAudioFormat *string `json:"output_audio_format,omitempty" xml:"output_audio_format,omitempty"`
  Output []*ResponseDoneResponseOutput `json:"output,omitempty" xml:"output,omitempty" type:"Repeated"`
  Usage *ResponseDoneResponseUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s ResponseDoneResponse) String() string {
  return dara.Prettify(s)
}

func (s ResponseDoneResponse) GoString() string {
  return s.String()
}

func (s *ResponseDoneResponse) GetId() *string  {
  return s.Id
}

func (s *ResponseDoneResponse) GetObject() *string  {
  return s.Object
}

func (s *ResponseDoneResponse) GetConversationId() *string  {
  return s.ConversationId
}

func (s *ResponseDoneResponse) GetStatus() *string  {
  return s.Status
}

func (s *ResponseDoneResponse) GetModalities() []*string  {
  return s.Modalities
}

func (s *ResponseDoneResponse) GetVoice() *string  {
  return s.Voice
}

func (s *ResponseDoneResponse) GetOutputAudioFormat() *string  {
  return s.OutputAudioFormat
}

func (s *ResponseDoneResponse) GetOutput() []*ResponseDoneResponseOutput  {
  return s.Output
}

func (s *ResponseDoneResponse) GetUsage() *ResponseDoneResponseUsage  {
  return s.Usage
}

func (s *ResponseDoneResponse) SetId(v string) *ResponseDoneResponse {
  s.Id = &v
  return s
}

func (s *ResponseDoneResponse) SetObject(v string) *ResponseDoneResponse {
  s.Object = &v
  return s
}

func (s *ResponseDoneResponse) SetConversationId(v string) *ResponseDoneResponse {
  s.ConversationId = &v
  return s
}

func (s *ResponseDoneResponse) SetStatus(v string) *ResponseDoneResponse {
  s.Status = &v
  return s
}

func (s *ResponseDoneResponse) SetModalities(v []*string) *ResponseDoneResponse {
  s.Modalities = v
  return s
}

func (s *ResponseDoneResponse) SetVoice(v string) *ResponseDoneResponse {
  s.Voice = &v
  return s
}

func (s *ResponseDoneResponse) SetOutputAudioFormat(v string) *ResponseDoneResponse {
  s.OutputAudioFormat = &v
  return s
}

func (s *ResponseDoneResponse) SetOutput(v []*ResponseDoneResponseOutput) *ResponseDoneResponse {
  s.Output = v
  return s
}

func (s *ResponseDoneResponse) SetUsage(v *ResponseDoneResponseUsage) *ResponseDoneResponse {
  s.Usage = v
  return s
}

func (s *ResponseDoneResponse) Validate() error {
  if s.Output != nil {
    for _, item := range s.Output {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Usage != nil {
    if err := s.Usage.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseDoneResponseOutput struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  Content []*ResponseDoneResponseOutputContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s ResponseDoneResponseOutput) String() string {
  return dara.Prettify(s)
}

func (s ResponseDoneResponseOutput) GoString() string {
  return s.String()
}

func (s *ResponseDoneResponseOutput) GetId() *string  {
  return s.Id
}

func (s *ResponseDoneResponseOutput) GetObject() *string  {
  return s.Object
}

func (s *ResponseDoneResponseOutput) GetType() *string  {
  return s.Type
}

func (s *ResponseDoneResponseOutput) GetStatus() *string  {
  return s.Status
}

func (s *ResponseDoneResponseOutput) GetRole() *string  {
  return s.Role
}

func (s *ResponseDoneResponseOutput) GetContent() []*ResponseDoneResponseOutputContent  {
  return s.Content
}

func (s *ResponseDoneResponseOutput) SetId(v string) *ResponseDoneResponseOutput {
  s.Id = &v
  return s
}

func (s *ResponseDoneResponseOutput) SetObject(v string) *ResponseDoneResponseOutput {
  s.Object = &v
  return s
}

func (s *ResponseDoneResponseOutput) SetType(v string) *ResponseDoneResponseOutput {
  s.Type = &v
  return s
}

func (s *ResponseDoneResponseOutput) SetStatus(v string) *ResponseDoneResponseOutput {
  s.Status = &v
  return s
}

func (s *ResponseDoneResponseOutput) SetRole(v string) *ResponseDoneResponseOutput {
  s.Role = &v
  return s
}

func (s *ResponseDoneResponseOutput) SetContent(v []*ResponseDoneResponseOutputContent) *ResponseDoneResponseOutput {
  s.Content = v
  return s
}

func (s *ResponseDoneResponseOutput) Validate() error {
  if s.Content != nil {
    for _, item := range s.Content {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ResponseDoneResponseOutputContent struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s ResponseDoneResponseOutputContent) String() string {
  return dara.Prettify(s)
}

func (s ResponseDoneResponseOutputContent) GoString() string {
  return s.String()
}

func (s *ResponseDoneResponseOutputContent) GetType() *string  {
  return s.Type
}

func (s *ResponseDoneResponseOutputContent) GetText() *string  {
  return s.Text
}

func (s *ResponseDoneResponseOutputContent) GetTranscript() *string  {
  return s.Transcript
}

func (s *ResponseDoneResponseOutputContent) SetType(v string) *ResponseDoneResponseOutputContent {
  s.Type = &v
  return s
}

func (s *ResponseDoneResponseOutputContent) SetText(v string) *ResponseDoneResponseOutputContent {
  s.Text = &v
  return s
}

func (s *ResponseDoneResponseOutputContent) SetTranscript(v string) *ResponseDoneResponseOutputContent {
  s.Transcript = &v
  return s
}

func (s *ResponseDoneResponseOutputContent) Validate() error {
  return dara.Validate(s)
}

type ResponseDoneResponseUsage struct {
  TotalTokens *int64 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
  InputTokens *int64 `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
  OutputTokens *int64 `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
  InputTokensDetails *ResponseDoneResponseUsageInputTokensDetails `json:"input_tokens_details,omitempty" xml:"input_tokens_details,omitempty" type:"Struct"`
  OutputTokensDetails *ResponseDoneResponseUsageOutputTokensDetails `json:"output_tokens_details,omitempty" xml:"output_tokens_details,omitempty" type:"Struct"`
}

func (s ResponseDoneResponseUsage) String() string {
  return dara.Prettify(s)
}

func (s ResponseDoneResponseUsage) GoString() string {
  return s.String()
}

func (s *ResponseDoneResponseUsage) GetTotalTokens() *int64  {
  return s.TotalTokens
}

func (s *ResponseDoneResponseUsage) GetInputTokens() *int64  {
  return s.InputTokens
}

func (s *ResponseDoneResponseUsage) GetOutputTokens() *int64  {
  return s.OutputTokens
}

func (s *ResponseDoneResponseUsage) GetInputTokensDetails() *ResponseDoneResponseUsageInputTokensDetails  {
  return s.InputTokensDetails
}

func (s *ResponseDoneResponseUsage) GetOutputTokensDetails() *ResponseDoneResponseUsageOutputTokensDetails  {
  return s.OutputTokensDetails
}

func (s *ResponseDoneResponseUsage) SetTotalTokens(v int64) *ResponseDoneResponseUsage {
  s.TotalTokens = &v
  return s
}

func (s *ResponseDoneResponseUsage) SetInputTokens(v int64) *ResponseDoneResponseUsage {
  s.InputTokens = &v
  return s
}

func (s *ResponseDoneResponseUsage) SetOutputTokens(v int64) *ResponseDoneResponseUsage {
  s.OutputTokens = &v
  return s
}

func (s *ResponseDoneResponseUsage) SetInputTokensDetails(v *ResponseDoneResponseUsageInputTokensDetails) *ResponseDoneResponseUsage {
  s.InputTokensDetails = v
  return s
}

func (s *ResponseDoneResponseUsage) SetOutputTokensDetails(v *ResponseDoneResponseUsageOutputTokensDetails) *ResponseDoneResponseUsage {
  s.OutputTokensDetails = v
  return s
}

func (s *ResponseDoneResponseUsage) Validate() error {
  if s.InputTokensDetails != nil {
    if err := s.InputTokensDetails.Validate(); err != nil {
      return err
    }
  }
  if s.OutputTokensDetails != nil {
    if err := s.OutputTokensDetails.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseDoneResponseUsageInputTokensDetails struct {
  TextTokens *int64 `json:"text_tokens,omitempty" xml:"text_tokens,omitempty"`
  AudioTokens *int64 `json:"audio_tokens,omitempty" xml:"audio_tokens,omitempty"`
}

func (s ResponseDoneResponseUsageInputTokensDetails) String() string {
  return dara.Prettify(s)
}

func (s ResponseDoneResponseUsageInputTokensDetails) GoString() string {
  return s.String()
}

func (s *ResponseDoneResponseUsageInputTokensDetails) GetTextTokens() *int64  {
  return s.TextTokens
}

func (s *ResponseDoneResponseUsageInputTokensDetails) GetAudioTokens() *int64  {
  return s.AudioTokens
}

func (s *ResponseDoneResponseUsageInputTokensDetails) SetTextTokens(v int64) *ResponseDoneResponseUsageInputTokensDetails {
  s.TextTokens = &v
  return s
}

func (s *ResponseDoneResponseUsageInputTokensDetails) SetAudioTokens(v int64) *ResponseDoneResponseUsageInputTokensDetails {
  s.AudioTokens = &v
  return s
}

func (s *ResponseDoneResponseUsageInputTokensDetails) Validate() error {
  return dara.Validate(s)
}

type ResponseDoneResponseUsageOutputTokensDetails struct {
  TextTokens *int64 `json:"text_tokens,omitempty" xml:"text_tokens,omitempty"`
  AudioTokens *int64 `json:"audio_tokens,omitempty" xml:"audio_tokens,omitempty"`
}

func (s ResponseDoneResponseUsageOutputTokensDetails) String() string {
  return dara.Prettify(s)
}

func (s ResponseDoneResponseUsageOutputTokensDetails) GoString() string {
  return s.String()
}

func (s *ResponseDoneResponseUsageOutputTokensDetails) GetTextTokens() *int64  {
  return s.TextTokens
}

func (s *ResponseDoneResponseUsageOutputTokensDetails) GetAudioTokens() *int64  {
  return s.AudioTokens
}

func (s *ResponseDoneResponseUsageOutputTokensDetails) SetTextTokens(v int64) *ResponseDoneResponseUsageOutputTokensDetails {
  s.TextTokens = &v
  return s
}

func (s *ResponseDoneResponseUsageOutputTokensDetails) SetAudioTokens(v int64) *ResponseDoneResponseUsageOutputTokensDetails {
  s.AudioTokens = &v
  return s
}

func (s *ResponseDoneResponseUsageOutputTokensDetails) Validate() error {
  return dara.Validate(s)
}

