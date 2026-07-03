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
  // Specifies whether to automatically add alert logs from Security Center, Web Application Firewall (WAF), and Cloud Firewall. By default, alert logs are automatically added.
  // 
  // example:
  // 
  // 1
  AutoSubmit *int32 `json:"AutoSubmit,omitempty" xml:"AutoSubmit,omitempty"`
  // The region of the Data Management center for Threat Analysis. Select the region based on where your assets are located. Valid values:
  // 
  // - cn-hangzhou: Your assets are in the Chinese mainland or Hong Kong (China).
  // 
  // - ap-southeast-1: Your assets are in regions outside China.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The user ID of a member. An administrator can use this parameter to switch to the perspective of the specified member.
  // 
  // example:
  // 
  // 113091674488****
  RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
  // The type of the view.
  // 
  // - 0: The view of the current Alibaba Cloud account.
  // 
  // - 1: The view of all member accounts.
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

