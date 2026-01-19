// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationResourceServerResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationResourceServerResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationResourceServerResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationResourceServerResponseBody) *EnableApplicationResourceServerResponse
  GetBody() *EnableApplicationResourceServerResponseBody 
}

type EnableApplicationResourceServerResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationResourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationResourceServerResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationResourceServerResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationResourceServerResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationResourceServerResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationResourceServerResponse) GetBody() *EnableApplicationResourceServerResponseBody  {
  return s.Body
}

func (s *EnableApplicationResourceServerResponse) SetHeaders(v map[string]*string) *EnableApplicationResourceServerResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationResourceServerResponse) SetStatusCode(v int32) *EnableApplicationResourceServerResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationResourceServerResponse) SetBody(v *EnableApplicationResourceServerResponseBody) *EnableApplicationResourceServerResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationResourceServerResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

