// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTrustedOriginResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableTrustedOriginResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableTrustedOriginResponse
  GetStatusCode() *int32 
  SetBody(v *EnableTrustedOriginResponseBody) *EnableTrustedOriginResponse
  GetBody() *EnableTrustedOriginResponseBody 
}

type EnableTrustedOriginResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableTrustedOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableTrustedOriginResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableTrustedOriginResponse) GoString() string {
  return s.String()
}

func (s *EnableTrustedOriginResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableTrustedOriginResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableTrustedOriginResponse) GetBody() *EnableTrustedOriginResponseBody  {
  return s.Body
}

func (s *EnableTrustedOriginResponse) SetHeaders(v map[string]*string) *EnableTrustedOriginResponse {
  s.Headers = v
  return s
}

func (s *EnableTrustedOriginResponse) SetStatusCode(v int32) *EnableTrustedOriginResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableTrustedOriginResponse) SetBody(v *EnableTrustedOriginResponseBody) *EnableTrustedOriginResponse {
  s.Body = v
  return s
}

func (s *EnableTrustedOriginResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

