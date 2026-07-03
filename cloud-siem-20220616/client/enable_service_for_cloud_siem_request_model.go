// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceForCloudSiemRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableServiceForCloudSiemRequest
  GetRegionId() *string 
}

type EnableServiceForCloudSiemRequest struct {
  // The region of the Data Management center for threat analysis. Select a region based on where your assets are located. Valid values:
  // 
  // - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
  // 
  // - ap-southeast-1: Your assets are in a region outside China.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableServiceForCloudSiemRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceForCloudSiemRequest) GoString() string {
  return s.String()
}

func (s *EnableServiceForCloudSiemRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableServiceForCloudSiemRequest) SetRegionId(v string) *EnableServiceForCloudSiemRequest {
  s.RegionId = &v
  return s
}

func (s *EnableServiceForCloudSiemRequest) Validate() error {
  return dara.Validate(s)
}

