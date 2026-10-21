// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteUpgradeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *ExecuteUpgradeRequest
  GetLang() *string 
  SetRegionId(v string) *ExecuteUpgradeRequest
  GetRegionId() *string 
  SetRoleFor(v string) *ExecuteUpgradeRequest
  GetRoleFor() *string 
}

type ExecuteUpgradeRequest struct {
  // The language of the response. Valid values:
  // 
  // - **zh*	- (default): Chinese.
  // 
  // - **en**: English.
  // 
  // Note: This parameter is required in practice. If you do not specify this parameter, the API returns a 400 error.
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The region where the threat detection and response data management center resides. Specify this parameter based on the region where your assets reside. Valid values:
  // 
  // - cn-hangzhou: Your assets reside in the Chinese mainland.
  // 
  // - ap-southeast-1: Your assets reside outside China.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the member to which the administrator switches the view.
  // 
  // example:
  // 
  // 173326*******
  RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ExecuteUpgradeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteUpgradeRequest) GoString() string {
  return s.String()
}

func (s *ExecuteUpgradeRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteUpgradeRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteUpgradeRequest) GetRoleFor() *string  {
  return s.RoleFor
}

func (s *ExecuteUpgradeRequest) SetLang(v string) *ExecuteUpgradeRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteUpgradeRequest) SetRegionId(v string) *ExecuteUpgradeRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteUpgradeRequest) SetRoleFor(v string) *ExecuteUpgradeRequest {
  s.RoleFor = &v
  return s
}

func (s *ExecuteUpgradeRequest) Validate() error {
  return dara.Validate(s)
}

