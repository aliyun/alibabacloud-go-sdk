// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableServiceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableServiceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableServiceResponseBody) *EnableServiceResponse
  GetBody() *EnableServiceResponseBody 
}

type EnableServiceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServiceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceResponse) GoString() string {
  return s.String()
}

func (s *EnableServiceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableServiceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableServiceResponse) GetBody() *EnableServiceResponseBody  {
  return s.Body
}

func (s *EnableServiceResponse) SetHeaders(v map[string]*string) *EnableServiceResponse {
  s.Headers = v
  return s
}

func (s *EnableServiceResponse) SetStatusCode(v int32) *EnableServiceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableServiceResponse) SetBody(v *EnableServiceResponseBody) *EnableServiceResponse {
  s.Body = v
  return s
}

func (s *EnableServiceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

