// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEndpointResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableEndpointResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableEndpointResponse
  GetStatusCode() *int32 
  SetBody(v *EnableEndpointResponseBody) *EnableEndpointResponse
  GetBody() *EnableEndpointResponseBody 
}

type EnableEndpointResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableEndpointResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableEndpointResponse) GoString() string {
  return s.String()
}

func (s *EnableEndpointResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableEndpointResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableEndpointResponse) GetBody() *EnableEndpointResponseBody  {
  return s.Body
}

func (s *EnableEndpointResponse) SetHeaders(v map[string]*string) *EnableEndpointResponse {
  s.Headers = v
  return s
}

func (s *EnableEndpointResponse) SetStatusCode(v int32) *EnableEndpointResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableEndpointResponse) SetBody(v *EnableEndpointResponseBody) *EnableEndpointResponse {
  s.Body = v
  return s
}

func (s *EnableEndpointResponse) Validate() error {
  return dara.Validate(s)
}

