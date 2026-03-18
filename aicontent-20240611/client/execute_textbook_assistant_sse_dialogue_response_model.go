// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantSseDialogueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantSseDialogueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantSseDialogueResponse
  GetStatusCode() *int32 
  SetId(v string) *ExecuteTextbookAssistantSseDialogueResponse
  GetId() *string 
  SetEvent(v string) *ExecuteTextbookAssistantSseDialogueResponse
  GetEvent() *string 
  SetBody(v *ExecuteTextbookAssistantSseDialogueResponseBody) *ExecuteTextbookAssistantSseDialogueResponse
  GetBody() *ExecuteTextbookAssistantSseDialogueResponseBody 
}

type ExecuteTextbookAssistantSseDialogueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  Body *ExecuteTextbookAssistantSseDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) GetId() *string  {
  return s.Id
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) GetBody() *ExecuteTextbookAssistantSseDialogueResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantSseDialogueResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantSseDialogueResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetId(v string) *ExecuteTextbookAssistantSseDialogueResponse {
  s.Id = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetEvent(v string) *ExecuteTextbookAssistantSseDialogueResponse {
  s.Event = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetBody(v *ExecuteTextbookAssistantSseDialogueResponseBody) *ExecuteTextbookAssistantSseDialogueResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

