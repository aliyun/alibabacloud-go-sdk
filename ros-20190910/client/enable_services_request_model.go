// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServicesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableServicesRequest
  GetRegionId() *string 
  SetServiceNames(v []*string) *EnableServicesRequest
  GetServiceNames() []*string 
}

type EnableServicesRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ServiceNames []*string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty" type:"Repeated"`
}

func (s EnableServicesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableServicesRequest) GoString() string {
  return s.String()
}

func (s *EnableServicesRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableServicesRequest) GetServiceNames() []*string  {
  return s.ServiceNames
}

func (s *EnableServicesRequest) SetRegionId(v string) *EnableServicesRequest {
  s.RegionId = &v
  return s
}

func (s *EnableServicesRequest) SetServiceNames(v []*string) *EnableServicesRequest {
  s.ServiceNames = v
  return s
}

func (s *EnableServicesRequest) Validate() error {
  return dara.Validate(s)
}

