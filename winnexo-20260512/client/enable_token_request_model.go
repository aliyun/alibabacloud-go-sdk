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
  // The ID of the effective tenant.
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
  // The ID of the target user (WINNEXO platform user ID). If this parameter is left empty, the operation is performed on the caller. Administrators can specify another user\\"s ID to perform the operation on behalf of that user.
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

