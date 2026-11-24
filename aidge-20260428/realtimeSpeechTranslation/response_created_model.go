// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponseCreated interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *ResponseCreated
  GetType() *string 
  SetEventId(v string) *ResponseCreated
  GetEventId() *string 
  SetResponse(v *ResponseCreatedResponse) *ResponseCreated
  GetResponse() *ResponseCreatedResponse 
}

type ResponseCreated struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  Response *ResponseCreatedResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
}

func (s ResponseCreated) String() string {
  return dara.Prettify(s)
}

func (s ResponseCreated) GoString() string {
  return s.String()
}

func (s *ResponseCreated) GetType() *string  {
  return s.Type
}

func (s *ResponseCreated) GetEventId() *string  {
  return s.EventId
}

func (s *ResponseCreated) GetResponse() *ResponseCreatedResponse  {
  return s.Response
}

func (s *ResponseCreated) SetType(v string) *ResponseCreated {
  s.Type = &v
  return s
}

func (s *ResponseCreated) SetEventId(v string) *ResponseCreated {
  s.EventId = &v
  return s
}

func (s *ResponseCreated) SetResponse(v *ResponseCreatedResponse) *ResponseCreated {
  s.Response = v
  return s
}

func (s *ResponseCreated) Validate() error {
  if s.Response != nil {
    if err := s.Response.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ResponseCreatedResponse struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  ConversationId *string `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
  OutputAudioFormat *string `json:"output_audio_format,omitempty" xml:"output_audio_format,omitempty"`
  Output []*ResponseCreatedResponseOutput `json:"output,omitempty" xml:"output,omitempty" type:"Repeated"`
}

func (s ResponseCreatedResponse) String() string {
  return dara.Prettify(s)
}

func (s ResponseCreatedResponse) GoString() string {
  return s.String()
}

func (s *ResponseCreatedResponse) GetId() *string  {
  return s.Id
}

func (s *ResponseCreatedResponse) GetObject() *string  {
  return s.Object
}

func (s *ResponseCreatedResponse) GetConversationId() *string  {
  return s.ConversationId
}

func (s *ResponseCreatedResponse) GetStatus() *string  {
  return s.Status
}

func (s *ResponseCreatedResponse) GetModalities() []*string  {
  return s.Modalities
}

func (s *ResponseCreatedResponse) GetVoice() *string  {
  return s.Voice
}

func (s *ResponseCreatedResponse) GetOutputAudioFormat() *string  {
  return s.OutputAudioFormat
}

func (s *ResponseCreatedResponse) GetOutput() []*ResponseCreatedResponseOutput  {
  return s.Output
}

func (s *ResponseCreatedResponse) SetId(v string) *ResponseCreatedResponse {
  s.Id = &v
  return s
}

func (s *ResponseCreatedResponse) SetObject(v string) *ResponseCreatedResponse {
  s.Object = &v
  return s
}

func (s *ResponseCreatedResponse) SetConversationId(v string) *ResponseCreatedResponse {
  s.ConversationId = &v
  return s
}

func (s *ResponseCreatedResponse) SetStatus(v string) *ResponseCreatedResponse {
  s.Status = &v
  return s
}

func (s *ResponseCreatedResponse) SetModalities(v []*string) *ResponseCreatedResponse {
  s.Modalities = v
  return s
}

func (s *ResponseCreatedResponse) SetVoice(v string) *ResponseCreatedResponse {
  s.Voice = &v
  return s
}

func (s *ResponseCreatedResponse) SetOutputAudioFormat(v string) *ResponseCreatedResponse {
  s.OutputAudioFormat = &v
  return s
}

func (s *ResponseCreatedResponse) SetOutput(v []*ResponseCreatedResponseOutput) *ResponseCreatedResponse {
  s.Output = v
  return s
}

func (s *ResponseCreatedResponse) Validate() error {
  if s.Output != nil {
    for _, item := range s.Output {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ResponseCreatedResponseOutput struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
  Content []*ResponseCreatedResponseOutputContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s ResponseCreatedResponseOutput) String() string {
  return dara.Prettify(s)
}

func (s ResponseCreatedResponseOutput) GoString() string {
  return s.String()
}

func (s *ResponseCreatedResponseOutput) GetId() *string  {
  return s.Id
}

func (s *ResponseCreatedResponseOutput) GetObject() *string  {
  return s.Object
}

func (s *ResponseCreatedResponseOutput) GetType() *string  {
  return s.Type
}

func (s *ResponseCreatedResponseOutput) GetStatus() *string  {
  return s.Status
}

func (s *ResponseCreatedResponseOutput) GetRole() *string  {
  return s.Role
}

func (s *ResponseCreatedResponseOutput) GetContent() []*ResponseCreatedResponseOutputContent  {
  return s.Content
}

func (s *ResponseCreatedResponseOutput) SetId(v string) *ResponseCreatedResponseOutput {
  s.Id = &v
  return s
}

func (s *ResponseCreatedResponseOutput) SetObject(v string) *ResponseCreatedResponseOutput {
  s.Object = &v
  return s
}

func (s *ResponseCreatedResponseOutput) SetType(v string) *ResponseCreatedResponseOutput {
  s.Type = &v
  return s
}

func (s *ResponseCreatedResponseOutput) SetStatus(v string) *ResponseCreatedResponseOutput {
  s.Status = &v
  return s
}

func (s *ResponseCreatedResponseOutput) SetRole(v string) *ResponseCreatedResponseOutput {
  s.Role = &v
  return s
}

func (s *ResponseCreatedResponseOutput) SetContent(v []*ResponseCreatedResponseOutputContent) *ResponseCreatedResponseOutput {
  s.Content = v
  return s
}

func (s *ResponseCreatedResponseOutput) Validate() error {
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

type ResponseCreatedResponseOutputContent struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s ResponseCreatedResponseOutputContent) String() string {
  return dara.Prettify(s)
}

func (s ResponseCreatedResponseOutputContent) GoString() string {
  return s.String()
}

func (s *ResponseCreatedResponseOutputContent) GetType() *string  {
  return s.Type
}

func (s *ResponseCreatedResponseOutputContent) GetText() *string  {
  return s.Text
}

func (s *ResponseCreatedResponseOutputContent) GetTranscript() *string  {
  return s.Transcript
}

func (s *ResponseCreatedResponseOutputContent) SetType(v string) *ResponseCreatedResponseOutputContent {
  s.Type = &v
  return s
}

func (s *ResponseCreatedResponseOutputContent) SetText(v string) *ResponseCreatedResponseOutputContent {
  s.Text = &v
  return s
}

func (s *ResponseCreatedResponseOutputContent) SetTranscript(v string) *ResponseCreatedResponseOutputContent {
  s.Transcript = &v
  return s
}

func (s *ResponseCreatedResponseOutputContent) Validate() error {
  return dara.Validate(s)
}

