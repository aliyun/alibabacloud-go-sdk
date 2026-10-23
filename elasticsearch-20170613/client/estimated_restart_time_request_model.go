// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedRestartTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBody(v string) *EstimatedRestartTimeRequest
  GetBody() *string 
  SetForce(v bool) *EstimatedRestartTimeRequest
  GetForce() *bool 
}

type EstimatedRestartTimeRequest struct {
  // The restart parameters, such as the restart method, target nodes, and blue-green deployment, represented as a JSON string.
  // 
  // example:
  // 
  // {"restartType":"nodeIp","nodes":["172.16.xx.xx"],"blueGreenDep":true}
  Body *string `json:"body,omitempty" xml:"body,omitempty"`
  // Specifies whether to force a restart. Default value: false.
  // 
  // example:
  // 
  // false
  Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s EstimatedRestartTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s EstimatedRestartTimeRequest) GoString() string {
  return s.String()
}

func (s *EstimatedRestartTimeRequest) GetBody() *string  {
  return s.Body
}

func (s *EstimatedRestartTimeRequest) GetForce() *bool  {
  return s.Force
}

func (s *EstimatedRestartTimeRequest) SetBody(v string) *EstimatedRestartTimeRequest {
  s.Body = &v
  return s
}

func (s *EstimatedRestartTimeRequest) SetForce(v bool) *EstimatedRestartTimeRequest {
  s.Force = &v
  return s
}

func (s *EstimatedRestartTimeRequest) Validate() error {
  return dara.Validate(s)
}

