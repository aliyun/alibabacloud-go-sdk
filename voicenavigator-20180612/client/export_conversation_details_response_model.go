// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportConversationDetailsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportConversationDetailsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportConversationDetailsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportConversationDetailsResponseBody) *ExportConversationDetailsResponse
  GetBody() *ExportConversationDetailsResponseBody 
}

type ExportConversationDetailsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportConversationDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportConversationDetailsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportConversationDetailsResponse) GoString() string {
  return s.String()
}

func (s *ExportConversationDetailsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportConversationDetailsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportConversationDetailsResponse) GetBody() *ExportConversationDetailsResponseBody  {
  return s.Body
}

func (s *ExportConversationDetailsResponse) SetHeaders(v map[string]*string) *ExportConversationDetailsResponse {
  s.Headers = v
  return s
}

func (s *ExportConversationDetailsResponse) SetStatusCode(v int32) *ExportConversationDetailsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportConversationDetailsResponse) SetBody(v *ExportConversationDetailsResponseBody) *ExportConversationDetailsResponse {
  s.Body = v
  return s
}

func (s *ExportConversationDetailsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

