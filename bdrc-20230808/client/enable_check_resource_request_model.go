// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCheckResourceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetResourceArn(v string) *EnableCheckResourceRequest
  GetResourceArn() *string 
}

type EnableCheckResourceRequest struct {
  // Unique resource identity.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // acs:ecs:123***890:cn-shanghai:instance/i-001***90
  ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s EnableCheckResourceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCheckResourceRequest) GoString() string {
  return s.String()
}

func (s *EnableCheckResourceRequest) GetResourceArn() *string  {
  return s.ResourceArn
}

func (s *EnableCheckResourceRequest) SetResourceArn(v string) *EnableCheckResourceRequest {
  s.ResourceArn = &v
  return s
}

func (s *EnableCheckResourceRequest) Validate() error {
  return dara.Validate(s)
}

