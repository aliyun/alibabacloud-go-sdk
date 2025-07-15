// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceAccessControlRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAclId(v string) *EnableInstanceAccessControlRequest
  GetAclId() *string 
  SetAclType(v string) *EnableInstanceAccessControlRequest
  GetAclType() *string 
  SetAddressIPVersion(v string) *EnableInstanceAccessControlRequest
  GetAddressIPVersion() *string 
  SetInstanceId(v string) *EnableInstanceAccessControlRequest
  GetInstanceId() *string 
  SetSecurityToken(v string) *EnableInstanceAccessControlRequest
  GetSecurityToken() *string 
}

type EnableInstanceAccessControlRequest struct {
  // The ID of the access control policy.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // acl-bp11escro2et2tioscy52
  AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
  // The ACL type. Valid values:
  // 
  // 	- black: blacklist
  // 
  // 	- white: whitelist
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // black
  AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
  // The IP version. Valid values: **ipv4*	- and **ipv6**.
  // 
  // example:
  // 
  // ipv4
  AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
  // The ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // apigateway-cn-v6419k43xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s EnableInstanceAccessControlRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceAccessControlRequest) GoString() string {
  return s.String()
}

func (s *EnableInstanceAccessControlRequest) GetAclId() *string  {
  return s.AclId
}

func (s *EnableInstanceAccessControlRequest) GetAclType() *string  {
  return s.AclType
}

func (s *EnableInstanceAccessControlRequest) GetAddressIPVersion() *string  {
  return s.AddressIPVersion
}

func (s *EnableInstanceAccessControlRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInstanceAccessControlRequest) GetSecurityToken() *string  {
  return s.SecurityToken
}

func (s *EnableInstanceAccessControlRequest) SetAclId(v string) *EnableInstanceAccessControlRequest {
  s.AclId = &v
  return s
}

func (s *EnableInstanceAccessControlRequest) SetAclType(v string) *EnableInstanceAccessControlRequest {
  s.AclType = &v
  return s
}

func (s *EnableInstanceAccessControlRequest) SetAddressIPVersion(v string) *EnableInstanceAccessControlRequest {
  s.AddressIPVersion = &v
  return s
}

func (s *EnableInstanceAccessControlRequest) SetInstanceId(v string) *EnableInstanceAccessControlRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableInstanceAccessControlRequest) SetSecurityToken(v string) *EnableInstanceAccessControlRequest {
  s.SecurityToken = &v
  return s
}

func (s *EnableInstanceAccessControlRequest) Validate() error {
  return dara.Validate(s)
}

