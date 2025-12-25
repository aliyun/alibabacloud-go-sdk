// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateTicketResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateTicketResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateTicketResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateTicketResponseBody) *EvaluateTicketResponse
  GetBody() *EvaluateTicketResponseBody 
}

type EvaluateTicketResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateTicketResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateTicketResponse) GoString() string {
  return s.String()
}

func (s *EvaluateTicketResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateTicketResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateTicketResponse) GetBody() *EvaluateTicketResponseBody  {
  return s.Body
}

func (s *EvaluateTicketResponse) SetHeaders(v map[string]*string) *EvaluateTicketResponse {
  s.Headers = v
  return s
}

func (s *EvaluateTicketResponse) SetStatusCode(v int32) *EvaluateTicketResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateTicketResponse) SetBody(v *EvaluateTicketResponseBody) *EvaluateTicketResponse {
  s.Body = v
  return s
}

func (s *EvaluateTicketResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

