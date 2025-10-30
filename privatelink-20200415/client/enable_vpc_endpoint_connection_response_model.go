// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcEndpointConnectionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableVpcEndpointConnectionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableVpcEndpointConnectionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableVpcEndpointConnectionResponseBody) *EnableVpcEndpointConnectionResponse
  GetBody() *EnableVpcEndpointConnectionResponseBody 
}

type EnableVpcEndpointConnectionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableVpcEndpointConnectionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcEndpointConnectionResponse) GoString() string {
  return s.String()
}

func (s *EnableVpcEndpointConnectionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableVpcEndpointConnectionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableVpcEndpointConnectionResponse) GetBody() *EnableVpcEndpointConnectionResponseBody  {
  return s.Body
}

func (s *EnableVpcEndpointConnectionResponse) SetHeaders(v map[string]*string) *EnableVpcEndpointConnectionResponse {
  s.Headers = v
  return s
}

func (s *EnableVpcEndpointConnectionResponse) SetStatusCode(v int32) *EnableVpcEndpointConnectionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableVpcEndpointConnectionResponse) SetBody(v *EnableVpcEndpointConnectionResponseBody) *EnableVpcEndpointConnectionResponse {
  s.Body = v
  return s
}

func (s *EnableVpcEndpointConnectionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

