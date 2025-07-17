// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableScalingGroupResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableScalingGroupResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableScalingGroupResponse
  GetStatusCode() *int32 
  SetBody(v *EnableScalingGroupResponseBody) *EnableScalingGroupResponse
  GetBody() *EnableScalingGroupResponseBody 
}

type EnableScalingGroupResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableScalingGroupResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableScalingGroupResponse) GoString() string {
  return s.String()
}

func (s *EnableScalingGroupResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableScalingGroupResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableScalingGroupResponse) GetBody() *EnableScalingGroupResponseBody  {
  return s.Body
}

func (s *EnableScalingGroupResponse) SetHeaders(v map[string]*string) *EnableScalingGroupResponse {
  s.Headers = v
  return s
}

func (s *EnableScalingGroupResponse) SetStatusCode(v int32) *EnableScalingGroupResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableScalingGroupResponse) SetBody(v *EnableScalingGroupResponseBody) *EnableScalingGroupResponse {
  s.Body = v
  return s
}

func (s *EnableScalingGroupResponse) Validate() error {
  return dara.Validate(s)
}

