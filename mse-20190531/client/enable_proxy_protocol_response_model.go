// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableProxyProtocolResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableProxyProtocolResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableProxyProtocolResponse
  GetStatusCode() *int32 
  SetBody(v *EnableProxyProtocolResponseBody) *EnableProxyProtocolResponse
  GetBody() *EnableProxyProtocolResponseBody 
}

type EnableProxyProtocolResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableProxyProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableProxyProtocolResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableProxyProtocolResponse) GoString() string {
  return s.String()
}

func (s *EnableProxyProtocolResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableProxyProtocolResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableProxyProtocolResponse) GetBody() *EnableProxyProtocolResponseBody  {
  return s.Body
}

func (s *EnableProxyProtocolResponse) SetHeaders(v map[string]*string) *EnableProxyProtocolResponse {
  s.Headers = v
  return s
}

func (s *EnableProxyProtocolResponse) SetStatusCode(v int32) *EnableProxyProtocolResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableProxyProtocolResponse) SetBody(v *EnableProxyProtocolResponseBody) *EnableProxyProtocolResponse {
  s.Body = v
  return s
}

func (s *EnableProxyProtocolResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

