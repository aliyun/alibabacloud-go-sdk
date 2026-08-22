// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOpenSearchPublicEndpointResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableOpenSearchPublicEndpointResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableOpenSearchPublicEndpointResponse
  GetStatusCode() *int32 
  SetBody(v *EnableOpenSearchPublicEndpointResponseBody) *EnableOpenSearchPublicEndpointResponse
  GetBody() *EnableOpenSearchPublicEndpointResponseBody 
}

type EnableOpenSearchPublicEndpointResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableOpenSearchPublicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableOpenSearchPublicEndpointResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableOpenSearchPublicEndpointResponse) GoString() string {
  return s.String()
}

func (s *EnableOpenSearchPublicEndpointResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableOpenSearchPublicEndpointResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableOpenSearchPublicEndpointResponse) GetBody() *EnableOpenSearchPublicEndpointResponseBody  {
  return s.Body
}

func (s *EnableOpenSearchPublicEndpointResponse) SetHeaders(v map[string]*string) *EnableOpenSearchPublicEndpointResponse {
  s.Headers = v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponse) SetStatusCode(v int32) *EnableOpenSearchPublicEndpointResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponse) SetBody(v *EnableOpenSearchPublicEndpointResponseBody) *EnableOpenSearchPublicEndpointResponse {
  s.Body = v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

