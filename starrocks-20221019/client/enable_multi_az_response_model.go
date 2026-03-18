// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMultiAzResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableMultiAzResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableMultiAzResponse
  GetStatusCode() *int32 
  SetBody(v *EnableMultiAzResponseBody) *EnableMultiAzResponse
  GetBody() *EnableMultiAzResponseBody 
}

type EnableMultiAzResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableMultiAzResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMultiAzResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAzResponse) GoString() string {
  return s.String()
}

func (s *EnableMultiAzResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableMultiAzResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableMultiAzResponse) GetBody() *EnableMultiAzResponseBody  {
  return s.Body
}

func (s *EnableMultiAzResponse) SetHeaders(v map[string]*string) *EnableMultiAzResponse {
  s.Headers = v
  return s
}

func (s *EnableMultiAzResponse) SetStatusCode(v int32) *EnableMultiAzResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableMultiAzResponse) SetBody(v *EnableMultiAzResponseBody) *EnableMultiAzResponse {
  s.Body = v
  return s
}

func (s *EnableMultiAzResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

