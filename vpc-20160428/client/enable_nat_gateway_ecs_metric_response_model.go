// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNatGatewayEcsMetricResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableNatGatewayEcsMetricResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableNatGatewayEcsMetricResponse
  GetStatusCode() *int32 
  SetBody(v *EnableNatGatewayEcsMetricResponseBody) *EnableNatGatewayEcsMetricResponse
  GetBody() *EnableNatGatewayEcsMetricResponseBody 
}

type EnableNatGatewayEcsMetricResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableNatGatewayEcsMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableNatGatewayEcsMetricResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableNatGatewayEcsMetricResponse) GoString() string {
  return s.String()
}

func (s *EnableNatGatewayEcsMetricResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableNatGatewayEcsMetricResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableNatGatewayEcsMetricResponse) GetBody() *EnableNatGatewayEcsMetricResponseBody  {
  return s.Body
}

func (s *EnableNatGatewayEcsMetricResponse) SetHeaders(v map[string]*string) *EnableNatGatewayEcsMetricResponse {
  s.Headers = v
  return s
}

func (s *EnableNatGatewayEcsMetricResponse) SetStatusCode(v int32) *EnableNatGatewayEcsMetricResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableNatGatewayEcsMetricResponse) SetBody(v *EnableNatGatewayEcsMetricResponseBody) *EnableNatGatewayEcsMetricResponse {
  s.Body = v
  return s
}

func (s *EnableNatGatewayEcsMetricResponse) Validate() error {
  return dara.Validate(s)
}

