// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudAccountRoleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableCloudAccountRoleRequest
  GetClientToken() *string 
  SetCloudAccountId(v string) *EnableCloudAccountRoleRequest
  GetCloudAccountId() *string 
  SetCloudAccountRoleId(v string) *EnableCloudAccountRoleRequest
  GetCloudAccountRoleId() *string 
  SetInstanceId(v string) *EnableCloudAccountRoleRequest
  GetInstanceId() *string 
}

type EnableCloudAccountRoleRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // ca_01kmegjc11qa1txxxxx
  CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
  // 云账号角色ID
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // carole_01kmek49aqxxxx
  CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableCloudAccountRoleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudAccountRoleRequest) GoString() string {
  return s.String()
}

func (s *EnableCloudAccountRoleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableCloudAccountRoleRequest) GetCloudAccountId() *string  {
  return s.CloudAccountId
}

func (s *EnableCloudAccountRoleRequest) GetCloudAccountRoleId() *string  {
  return s.CloudAccountRoleId
}

func (s *EnableCloudAccountRoleRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCloudAccountRoleRequest) SetClientToken(v string) *EnableCloudAccountRoleRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableCloudAccountRoleRequest) SetCloudAccountId(v string) *EnableCloudAccountRoleRequest {
  s.CloudAccountId = &v
  return s
}

func (s *EnableCloudAccountRoleRequest) SetCloudAccountRoleId(v string) *EnableCloudAccountRoleRequest {
  s.CloudAccountRoleId = &v
  return s
}

func (s *EnableCloudAccountRoleRequest) SetInstanceId(v string) *EnableCloudAccountRoleRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCloudAccountRoleRequest) Validate() error {
  return dara.Validate(s)
}

