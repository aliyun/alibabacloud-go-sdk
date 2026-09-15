// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConnectorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableConnectorResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableConnectorResponse
  GetStatusCode() *int32 
  SetBody(v *EnableConnectorResponseBody) *EnableConnectorResponse
  GetBody() *EnableConnectorResponseBody 
}

type EnableConnectorResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableConnectorResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableConnectorResponse) GoString() string {
  return s.String()
}

func (s *EnableConnectorResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableConnectorResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableConnectorResponse) GetBody() *EnableConnectorResponseBody  {
  return s.Body
}

func (s *EnableConnectorResponse) SetHeaders(v map[string]*string) *EnableConnectorResponse {
  s.Headers = v
  return s
}

func (s *EnableConnectorResponse) SetStatusCode(v int32) *EnableConnectorResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableConnectorResponse) SetBody(v *EnableConnectorResponseBody) *EnableConnectorResponse {
  s.Body = v
  return s
}

func (s *EnableConnectorResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

