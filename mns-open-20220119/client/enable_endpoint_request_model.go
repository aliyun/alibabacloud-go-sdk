// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEndpointRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndpointType(v string) *EnableEndpointRequest
  GetEndpointType() *string 
}

type EnableEndpointRequest struct {
  // The type of the endpoint. Valid value:
  // 
  // 	- **public**: indicates public endpoint. (Only the public is supported.)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // public
  EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s EnableEndpointRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableEndpointRequest) GoString() string {
  return s.String()
}

func (s *EnableEndpointRequest) GetEndpointType() *string  {
  return s.EndpointType
}

func (s *EnableEndpointRequest) SetEndpointType(v string) *EnableEndpointRequest {
  s.EndpointType = &v
  return s
}

func (s *EnableEndpointRequest) Validate() error {
  return dara.Validate(s)
}

