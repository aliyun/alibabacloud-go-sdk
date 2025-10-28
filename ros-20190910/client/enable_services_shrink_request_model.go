// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServicesShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableServicesShrinkRequest
  GetRegionId() *string 
  SetServiceNamesShrink(v string) *EnableServicesShrinkRequest
  GetServiceNamesShrink() *string 
}

type EnableServicesShrinkRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ServiceNamesShrink *string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty"`
}

func (s EnableServicesShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableServicesShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnableServicesShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableServicesShrinkRequest) GetServiceNamesShrink() *string  {
  return s.ServiceNamesShrink
}

func (s *EnableServicesShrinkRequest) SetRegionId(v string) *EnableServicesShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *EnableServicesShrinkRequest) SetServiceNamesShrink(v string) *EnableServicesShrinkRequest {
  s.ServiceNamesShrink = &v
  return s
}

func (s *EnableServicesShrinkRequest) Validate() error {
  return dara.Validate(s)
}

