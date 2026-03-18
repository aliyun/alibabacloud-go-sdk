// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantTranslateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantTranslateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantTranslateResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantTranslateResponseBody) *ExecuteTextbookAssistantTranslateResponse
  GetBody() *ExecuteTextbookAssistantTranslateResponseBody 
}

type ExecuteTextbookAssistantTranslateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantTranslateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantTranslateResponse) GetBody() *ExecuteTextbookAssistantTranslateResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantTranslateResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantTranslateResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantTranslateResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponse) SetBody(v *ExecuteTextbookAssistantTranslateResponseBody) *ExecuteTextbookAssistantTranslateResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantTranslateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

