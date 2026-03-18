// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherEnglishParaphraseChatMessageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherEnglishParaphraseChatMessageResponse
  GetStatusCode() *int32 
  SetId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse
  GetId() *string 
  SetEvent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse
  GetEvent() *string 
  SetBody(v *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) *ExecuteAITeacherEnglishParaphraseChatMessageResponse
  GetBody() *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody 
}

type ExecuteAITeacherEnglishParaphraseChatMessageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  Body *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) GetId() *string  {
  return s.Id
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) GetBody() *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetStatusCode(v int32) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
  s.Id = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetEvent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
  s.Event = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetBody(v *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

