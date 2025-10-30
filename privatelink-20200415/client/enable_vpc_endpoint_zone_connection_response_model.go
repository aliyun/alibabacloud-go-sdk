// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcEndpointZoneConnectionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableVpcEndpointZoneConnectionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableVpcEndpointZoneConnectionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableVpcEndpointZoneConnectionResponseBody) *EnableVpcEndpointZoneConnectionResponse
  GetBody() *EnableVpcEndpointZoneConnectionResponseBody 
}

type EnableVpcEndpointZoneConnectionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableVpcEndpointZoneConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableVpcEndpointZoneConnectionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcEndpointZoneConnectionResponse) GoString() string {
  return s.String()
}

func (s *EnableVpcEndpointZoneConnectionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableVpcEndpointZoneConnectionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableVpcEndpointZoneConnectionResponse) GetBody() *EnableVpcEndpointZoneConnectionResponseBody  {
  return s.Body
}

func (s *EnableVpcEndpointZoneConnectionResponse) SetHeaders(v map[string]*string) *EnableVpcEndpointZoneConnectionResponse {
  s.Headers = v
  return s
}

func (s *EnableVpcEndpointZoneConnectionResponse) SetStatusCode(v int32) *EnableVpcEndpointZoneConnectionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionResponse) SetBody(v *EnableVpcEndpointZoneConnectionResponseBody) *EnableVpcEndpointZoneConnectionResponse {
  s.Body = v
  return s
}

func (s *EnableVpcEndpointZoneConnectionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

