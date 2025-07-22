// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableExpressConnectRouterRouteEntriesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableExpressConnectRouterRouteEntriesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponse
  GetStatusCode() *int32 
  SetBody(v *EnableExpressConnectRouterRouteEntriesResponseBody) *EnableExpressConnectRouterRouteEntriesResponse
  GetBody() *EnableExpressConnectRouterRouteEntriesResponseBody 
}

type EnableExpressConnectRouterRouteEntriesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesResponse) GoString() string {
  return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) GetBody() *EnableExpressConnectRouterRouteEntriesResponseBody  {
  return s.Body
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *EnableExpressConnectRouterRouteEntriesResponse {
  s.Headers = v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetBody(v *EnableExpressConnectRouterRouteEntriesResponseBody) *EnableExpressConnectRouterRouteEntriesResponse {
  s.Body = v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) Validate() error {
  return dara.Validate(s)
}

