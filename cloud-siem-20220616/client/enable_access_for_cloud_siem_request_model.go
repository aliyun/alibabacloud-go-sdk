// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAccessForCloudSiemRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAutoSubmit(v int32) *EnableAccessForCloudSiemRequest
  GetAutoSubmit() *int32 
  SetRegionId(v string) *EnableAccessForCloudSiemRequest
  GetRegionId() *string 
  SetRoleFor(v int64) *EnableAccessForCloudSiemRequest
  GetRoleFor() *int64 
  SetRoleType(v int32) *EnableAccessForCloudSiemRequest
  GetRoleType() *int32 
}

type EnableAccessForCloudSiemRequest struct {
  // Whether import the log of SAS alert, the log of WAF alert, the log of CFW alert or not. Valid values:
  // 
  // - 0: not imported automatically
  // 
  // - 1: imported automatically
  // 
  // example:
  // 
  // 1
  AutoSubmit *int32 `json:"AutoSubmit,omitempty" xml:"AutoSubmit,omitempty"`
  // The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
  // 
  // 	- cn-hangzhou: Your assets reside in regions inside China.
  // 
  // 	- ap-southeast-1: Your assets reside in regions outside China.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the account that you switch from the management account.
  // 
  // example:
  // 
  // 113091674488****
  RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
  // The type of the view. Valid values:
  // 
  // - 0: the current Alibaba Cloud account
  // 
  // - 1: the global account
  // 
  // example:
  // 
  // 1
  RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s EnableAccessForCloudSiemRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAccessForCloudSiemRequest) GoString() string {
  return s.String()
}

func (s *EnableAccessForCloudSiemRequest) GetAutoSubmit() *int32  {
  return s.AutoSubmit
}

func (s *EnableAccessForCloudSiemRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAccessForCloudSiemRequest) GetRoleFor() *int64  {
  return s.RoleFor
}

func (s *EnableAccessForCloudSiemRequest) GetRoleType() *int32  {
  return s.RoleType
}

func (s *EnableAccessForCloudSiemRequest) SetAutoSubmit(v int32) *EnableAccessForCloudSiemRequest {
  s.AutoSubmit = &v
  return s
}

func (s *EnableAccessForCloudSiemRequest) SetRegionId(v string) *EnableAccessForCloudSiemRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAccessForCloudSiemRequest) SetRoleFor(v int64) *EnableAccessForCloudSiemRequest {
  s.RoleFor = &v
  return s
}

func (s *EnableAccessForCloudSiemRequest) SetRoleType(v int32) *EnableAccessForCloudSiemRequest {
  s.RoleType = &v
  return s
}

func (s *EnableAccessForCloudSiemRequest) Validate() error {
  return dara.Validate(s)
}

