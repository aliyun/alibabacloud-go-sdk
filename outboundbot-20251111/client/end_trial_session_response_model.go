// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndTrialSessionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndTrialSessionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndTrialSessionResponse
  GetStatusCode() *int32 
  SetBody(v *EndTrialSessionResponseBody) *EndTrialSessionResponse
  GetBody() *EndTrialSessionResponseBody 
}

type EndTrialSessionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndTrialSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndTrialSessionResponse) String() string {
  return dara.Prettify(s)
}

func (s EndTrialSessionResponse) GoString() string {
  return s.String()
}

func (s *EndTrialSessionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndTrialSessionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndTrialSessionResponse) GetBody() *EndTrialSessionResponseBody  {
  return s.Body
}

func (s *EndTrialSessionResponse) SetHeaders(v map[string]*string) *EndTrialSessionResponse {
  s.Headers = v
  return s
}

func (s *EndTrialSessionResponse) SetStatusCode(v int32) *EndTrialSessionResponse {
  s.StatusCode = &v
  return s
}

func (s *EndTrialSessionResponse) SetBody(v *EndTrialSessionResponseBody) *EndTrialSessionResponse {
  s.Body = v
  return s
}

func (s *EndTrialSessionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

