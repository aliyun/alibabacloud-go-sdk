// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLoadBalancerAccessLogResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableLoadBalancerAccessLogResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableLoadBalancerAccessLogResponse
  GetStatusCode() *int32 
  SetBody(v *EnableLoadBalancerAccessLogResponseBody) *EnableLoadBalancerAccessLogResponse
  GetBody() *EnableLoadBalancerAccessLogResponseBody 
}

type EnableLoadBalancerAccessLogResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableLoadBalancerAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableLoadBalancerAccessLogResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableLoadBalancerAccessLogResponse) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerAccessLogResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableLoadBalancerAccessLogResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableLoadBalancerAccessLogResponse) GetBody() *EnableLoadBalancerAccessLogResponseBody  {
  return s.Body
}

func (s *EnableLoadBalancerAccessLogResponse) SetHeaders(v map[string]*string) *EnableLoadBalancerAccessLogResponse {
  s.Headers = v
  return s
}

func (s *EnableLoadBalancerAccessLogResponse) SetStatusCode(v int32) *EnableLoadBalancerAccessLogResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableLoadBalancerAccessLogResponse) SetBody(v *EnableLoadBalancerAccessLogResponseBody) *EnableLoadBalancerAccessLogResponse {
  s.Body = v
  return s
}

func (s *EnableLoadBalancerAccessLogResponse) Validate() error {
  return dara.Validate(s)
}

