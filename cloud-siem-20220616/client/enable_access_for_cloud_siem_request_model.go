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
  SetClientToken(v string) *EnableAccessForCloudSiemRequest
  GetClientToken() *string 
  SetRegionId(v string) *EnableAccessForCloudSiemRequest
  GetRegionId() *string 
  SetRoleFor(v int64) *EnableAccessForCloudSiemRequest
  GetRoleFor() *int64 
  SetRoleType(v int32) *EnableAccessForCloudSiemRequest
  GetRoleType() *int32 
}

type EnableAccessForCloudSiemRequest struct {
  // Specifies whether to automatically integrate alert logs from Security Center, Web Application Firewall (WAF), and Cloud Firewall. By default, the logs are automatically integrated.
  // 
  // example:
  // 
  // 1
  AutoSubmit *int32 `json:"AutoSubmit,omitempty" xml:"AutoSubmit,omitempty"`
  // The idempotency token.
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-426614174000
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The region where the threat detection and response data management center resides. Select the management center based on the region of your assets. Valid values:
  // 
  // - cn-hangzhou: assets in the Chinese mainland and Hong Kong (China).
  // 
  // - ap-southeast-1: assets outside China.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the member account to which the administrator switches the view.
  // 
  // example:
  // 
  // 113091674488****
  RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
  // The view type.
  // 
  // - 0: the view of the current Alibaba Cloud account.
  // 
  // - 1: the view of all accounts in the enterprise.
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

func (s *EnableAccessForCloudSiemRequest) GetClientToken() *string  {
  return s.ClientToken
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

func (s *EnableAccessForCloudSiemRequest) SetClientToken(v string) *EnableAccessForCloudSiemRequest {
  s.ClientToken = &v
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

