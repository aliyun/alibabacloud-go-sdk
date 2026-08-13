// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTokenRequest interface {
  dara.Model
  String() string
  GoString() string
  SetTenantId(v string) *EnableTokenRequest
  GetTenantId() *string 
  SetWnUserId(v string) *EnableTokenRequest
  GetWnUserId() *string 
}

type EnableTokenRequest struct {
  // 租户ID，公共参数，缺省时使用调用方默认租户
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
  // 目标用户 ID（WINNEXO 平台用户ID，空则操作自身，管理员可传入他人 ID 代操作）
  // 
  // example:
  // 
  // 1
  WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s EnableTokenRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableTokenRequest) GoString() string {
  return s.String()
}

func (s *EnableTokenRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *EnableTokenRequest) GetWnUserId() *string  {
  return s.WnUserId
}

func (s *EnableTokenRequest) SetTenantId(v string) *EnableTokenRequest {
  s.TenantId = &v
  return s
}

func (s *EnableTokenRequest) SetWnUserId(v string) *EnableTokenRequest {
  s.WnUserId = &v
  return s
}

func (s *EnableTokenRequest) Validate() error {
  return dara.Validate(s)
}

