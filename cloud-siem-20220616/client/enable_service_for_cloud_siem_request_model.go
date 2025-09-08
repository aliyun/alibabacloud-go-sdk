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
  // The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
  // 
  // 	- cn-hangzhou: Your assets reside in regions in China.
  // 
  // 	- ap-southeast-1: Your assets reside in regions outside China.
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

