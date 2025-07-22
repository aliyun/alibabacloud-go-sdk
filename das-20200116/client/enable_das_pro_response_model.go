// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDasProResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDasProResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDasProResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDasProResponseBody) *EnableDasProResponse
  GetBody() *EnableDasProResponseBody 
}

type EnableDasProResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDasProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDasProResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDasProResponse) GoString() string {
  return s.String()
}

func (s *EnableDasProResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDasProResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDasProResponse) GetBody() *EnableDasProResponseBody  {
  return s.Body
}

func (s *EnableDasProResponse) SetHeaders(v map[string]*string) *EnableDasProResponse {
  s.Headers = v
  return s
}

func (s *EnableDasProResponse) SetStatusCode(v int32) *EnableDasProResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDasProResponse) SetBody(v *EnableDasProResponseBody) *EnableDasProResponse {
  s.Body = v
  return s
}

func (s *EnableDasProResponse) Validate() error {
  return dara.Validate(s)
}

