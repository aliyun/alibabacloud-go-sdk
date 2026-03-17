// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmartAccessGatewayUserRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOwnerAccount(v string) *EnableSmartAccessGatewayUserRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableSmartAccessGatewayUserRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableSmartAccessGatewayUserRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableSmartAccessGatewayUserRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableSmartAccessGatewayUserRequest
  GetResourceOwnerId() *int64 
  SetSmartAGId(v string) *EnableSmartAccessGatewayUserRequest
  GetSmartAGId() *string 
  SetUserName(v string) *EnableSmartAccessGatewayUserRequest
  GetUserName() *string 
}

type EnableSmartAccessGatewayUserRequest struct {
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the region where the SAG APP instance is deployed.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The ID of the SAG APP instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // sag-va03wf4l4idaj*****
  SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
  // The username of the client account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234
  UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s EnableSmartAccessGatewayUserRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSmartAccessGatewayUserRequest) GoString() string {
  return s.String()
}

func (s *EnableSmartAccessGatewayUserRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableSmartAccessGatewayUserRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableSmartAccessGatewayUserRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableSmartAccessGatewayUserRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableSmartAccessGatewayUserRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableSmartAccessGatewayUserRequest) GetSmartAGId() *string  {
  return s.SmartAGId
}

func (s *EnableSmartAccessGatewayUserRequest) GetUserName() *string  {
  return s.UserName
}

func (s *EnableSmartAccessGatewayUserRequest) SetOwnerAccount(v string) *EnableSmartAccessGatewayUserRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) SetOwnerId(v int64) *EnableSmartAccessGatewayUserRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) SetRegionId(v string) *EnableSmartAccessGatewayUserRequest {
  s.RegionId = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) SetResourceOwnerAccount(v string) *EnableSmartAccessGatewayUserRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) SetResourceOwnerId(v int64) *EnableSmartAccessGatewayUserRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) SetSmartAGId(v string) *EnableSmartAccessGatewayUserRequest {
  s.SmartAGId = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) SetUserName(v string) *EnableSmartAccessGatewayUserRequest {
  s.UserName = &v
  return s
}

func (s *EnableSmartAccessGatewayUserRequest) Validate() error {
  return dara.Validate(s)
}

