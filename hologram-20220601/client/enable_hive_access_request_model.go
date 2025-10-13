// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHiveAccessRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableHiveAccessRequest
  GetRegionId() *string 
}

type EnableHiveAccessRequest struct {
  // The region ID.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableHiveAccessRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableHiveAccessRequest) GoString() string {
  return s.String()
}

func (s *EnableHiveAccessRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableHiveAccessRequest) SetRegionId(v string) *EnableHiveAccessRequest {
  s.RegionId = &v
  return s
}

func (s *EnableHiveAccessRequest) Validate() error {
  return dara.Validate(s)
}

