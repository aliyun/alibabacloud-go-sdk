// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServerlessPublicConnectionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableServerlessPublicConnectionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableServerlessPublicConnectionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableServerlessPublicConnectionResponseBody) *EnableServerlessPublicConnectionResponse
  GetBody() *EnableServerlessPublicConnectionResponseBody 
}

type EnableServerlessPublicConnectionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableServerlessPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServerlessPublicConnectionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableServerlessPublicConnectionResponse) GoString() string {
  return s.String()
}

func (s *EnableServerlessPublicConnectionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableServerlessPublicConnectionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableServerlessPublicConnectionResponse) GetBody() *EnableServerlessPublicConnectionResponseBody  {
  return s.Body
}

func (s *EnableServerlessPublicConnectionResponse) SetHeaders(v map[string]*string) *EnableServerlessPublicConnectionResponse {
  s.Headers = v
  return s
}

func (s *EnableServerlessPublicConnectionResponse) SetStatusCode(v int32) *EnableServerlessPublicConnectionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableServerlessPublicConnectionResponse) SetBody(v *EnableServerlessPublicConnectionResponseBody) *EnableServerlessPublicConnectionResponse {
  s.Body = v
  return s
}

func (s *EnableServerlessPublicConnectionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

