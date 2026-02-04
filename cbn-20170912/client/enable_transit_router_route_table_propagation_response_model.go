// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTransitRouterRouteTablePropagationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableTransitRouterRouteTablePropagationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableTransitRouterRouteTablePropagationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableTransitRouterRouteTablePropagationResponseBody) *EnableTransitRouterRouteTablePropagationResponse
  GetBody() *EnableTransitRouterRouteTablePropagationResponseBody 
}

type EnableTransitRouterRouteTablePropagationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableTransitRouterRouteTablePropagationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableTransitRouterRouteTablePropagationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableTransitRouterRouteTablePropagationResponse) GoString() string {
  return s.String()
}

func (s *EnableTransitRouterRouteTablePropagationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableTransitRouterRouteTablePropagationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableTransitRouterRouteTablePropagationResponse) GetBody() *EnableTransitRouterRouteTablePropagationResponseBody  {
  return s.Body
}

func (s *EnableTransitRouterRouteTablePropagationResponse) SetHeaders(v map[string]*string) *EnableTransitRouterRouteTablePropagationResponse {
  s.Headers = v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationResponse) SetStatusCode(v int32) *EnableTransitRouterRouteTablePropagationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationResponse) SetBody(v *EnableTransitRouterRouteTablePropagationResponseBody) *EnableTransitRouterRouteTablePropagationResponse {
  s.Body = v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

