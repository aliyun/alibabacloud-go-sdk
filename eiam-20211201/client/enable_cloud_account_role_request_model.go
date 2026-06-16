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
  // A client token used to ensure the idempotence of the request. Generate a unique value for this parameter. The token can contain only ASCII characters and must be no more than 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The ID of the Alibaba Cloud account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ca_01kmegjc11qa1txxxxx
  CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
  // The ID of the cloud role.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // carole_01kmek49aqxxxx
  CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
  // The instance ID.
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

