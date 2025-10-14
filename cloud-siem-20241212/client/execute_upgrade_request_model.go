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
  // example:
  // 
  // zh。
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // example:
  // 
  // cn-hangzhou。
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // 173326*******。
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

