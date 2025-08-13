// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolicyTypeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOpenType(v string) *EnablePolicyTypeRequest
  GetOpenType() *string 
  SetOwnerAccount(v string) *EnablePolicyTypeRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnablePolicyTypeRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnablePolicyTypeRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnablePolicyTypeRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v string) *EnablePolicyTypeRequest
  GetResourceOwnerId() *string 
  SetUserType(v string) *EnablePolicyTypeRequest
  GetUserType() *string 
}

type EnablePolicyTypeRequest struct {
  OpenType *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s EnablePolicyTypeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnablePolicyTypeRequest) GoString() string {
  return s.String()
}

func (s *EnablePolicyTypeRequest) GetOpenType() *string  {
  return s.OpenType
}

func (s *EnablePolicyTypeRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnablePolicyTypeRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnablePolicyTypeRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnablePolicyTypeRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnablePolicyTypeRequest) GetResourceOwnerId() *string  {
  return s.ResourceOwnerId
}

func (s *EnablePolicyTypeRequest) GetUserType() *string  {
  return s.UserType
}

func (s *EnablePolicyTypeRequest) SetOpenType(v string) *EnablePolicyTypeRequest {
  s.OpenType = &v
  return s
}

func (s *EnablePolicyTypeRequest) SetOwnerAccount(v string) *EnablePolicyTypeRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnablePolicyTypeRequest) SetOwnerId(v int64) *EnablePolicyTypeRequest {
  s.OwnerId = &v
  return s
}

func (s *EnablePolicyTypeRequest) SetRegionId(v string) *EnablePolicyTypeRequest {
  s.RegionId = &v
  return s
}

func (s *EnablePolicyTypeRequest) SetResourceOwnerAccount(v string) *EnablePolicyTypeRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnablePolicyTypeRequest) SetResourceOwnerId(v string) *EnablePolicyTypeRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnablePolicyTypeRequest) SetUserType(v string) *EnablePolicyTypeRequest {
  s.UserType = &v
  return s
}

func (s *EnablePolicyTypeRequest) Validate() error {
  return dara.Validate(s)
}

