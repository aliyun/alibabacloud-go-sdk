// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTrustedOriginRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableTrustedOriginRequest
  GetInstanceId() *string 
  SetTrustedOriginId(v string) *EnableTrustedOriginRequest
  GetTrustedOriginId() *string 
}

type EnableTrustedOriginRequest struct {
  // The ID of the IDaaS EIAM instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_example
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the trusted origin.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // to_example
  TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
}

func (s EnableTrustedOriginRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableTrustedOriginRequest) GoString() string {
  return s.String()
}

func (s *EnableTrustedOriginRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableTrustedOriginRequest) GetTrustedOriginId() *string  {
  return s.TrustedOriginId
}

func (s *EnableTrustedOriginRequest) SetInstanceId(v string) *EnableTrustedOriginRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableTrustedOriginRequest) SetTrustedOriginId(v string) *EnableTrustedOriginRequest {
  s.TrustedOriginId = &v
  return s
}

func (s *EnableTrustedOriginRequest) Validate() error {
  return dara.Validate(s)
}

