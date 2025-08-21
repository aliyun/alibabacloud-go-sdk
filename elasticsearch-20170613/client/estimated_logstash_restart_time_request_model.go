// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedLogstashRestartTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBody(v string) *EstimatedLogstashRestartTimeRequest
  GetBody() *string 
  SetForce(v bool) *EstimatedLogstashRestartTimeRequest
  GetForce() *bool 
}

type EstimatedLogstashRestartTimeRequest struct {
  Body *string `json:"body,omitempty" xml:"body,omitempty"`
  // Specifies whether to forcibly restart the cluster. Default value: false.
  // 
  // example:
  // 
  // false
  Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s EstimatedLogstashRestartTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s EstimatedLogstashRestartTimeRequest) GoString() string {
  return s.String()
}

func (s *EstimatedLogstashRestartTimeRequest) GetBody() *string  {
  return s.Body
}

func (s *EstimatedLogstashRestartTimeRequest) GetForce() *bool  {
  return s.Force
}

func (s *EstimatedLogstashRestartTimeRequest) SetBody(v string) *EstimatedLogstashRestartTimeRequest {
  s.Body = &v
  return s
}

func (s *EstimatedLogstashRestartTimeRequest) SetForce(v bool) *EstimatedLogstashRestartTimeRequest {
  s.Force = &v
  return s
}

func (s *EstimatedLogstashRestartTimeRequest) Validate() error {
  return dara.Validate(s)
}

