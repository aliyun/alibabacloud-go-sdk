// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHostAvailabilityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetId(v []*int64) *EnableHostAvailabilityRequest
  GetId() []*int64 
  SetRegionId(v string) *EnableHostAvailabilityRequest
  GetRegionId() *string 
}

type EnableHostAvailabilityRequest struct {
  // The ID of the availability monitoring task. Valid values of N: 1 to 20.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 123456
  Id []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableHostAvailabilityRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableHostAvailabilityRequest) GoString() string {
  return s.String()
}

func (s *EnableHostAvailabilityRequest) GetId() []*int64  {
  return s.Id
}

func (s *EnableHostAvailabilityRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableHostAvailabilityRequest) SetId(v []*int64) *EnableHostAvailabilityRequest {
  s.Id = v
  return s
}

func (s *EnableHostAvailabilityRequest) SetRegionId(v string) *EnableHostAvailabilityRequest {
  s.RegionId = &v
  return s
}

func (s *EnableHostAvailabilityRequest) Validate() error {
  return dara.Validate(s)
}

