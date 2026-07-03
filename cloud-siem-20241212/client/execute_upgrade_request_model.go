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
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The region of the Data Management center for threat analysis. Select a region for the Management Hub based on the region of your assets. Valid values:
  // 
  // - cn-hangzhou: Your assets are in the Chinese mainland.
  // 
  // - ap-southeast-1: Your assets are in a region outside China.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The user ID of a member. An administrator can switch to the perspective of this member.
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

