// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceAccessResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableServiceAccessResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableServiceAccessResponse
  GetStatusCode() *int32 
  SetBody(v *EnableServiceAccessResponseBody) *EnableServiceAccessResponse
  GetBody() *EnableServiceAccessResponseBody 
}

type EnableServiceAccessResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableServiceAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServiceAccessResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceAccessResponse) GoString() string {
  return s.String()
}

func (s *EnableServiceAccessResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableServiceAccessResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableServiceAccessResponse) GetBody() *EnableServiceAccessResponseBody  {
  return s.Body
}

func (s *EnableServiceAccessResponse) SetHeaders(v map[string]*string) *EnableServiceAccessResponse {
  s.Headers = v
  return s
}

func (s *EnableServiceAccessResponse) SetStatusCode(v int32) *EnableServiceAccessResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableServiceAccessResponse) SetBody(v *EnableServiceAccessResponseBody) *EnableServiceAccessResponse {
  s.Body = v
  return s
}

func (s *EnableServiceAccessResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

