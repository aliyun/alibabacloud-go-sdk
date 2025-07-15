// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcClassicLinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableVpcClassicLinkRequest
  GetClientToken() *string 
  SetOwnerAccount(v string) *EnableVpcClassicLinkRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableVpcClassicLinkRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableVpcClassicLinkRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableVpcClassicLinkRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableVpcClassicLinkRequest
  GetResourceOwnerId() *int64 
  SetVpcId(v string) *EnableVpcClassicLinkRequest
  GetVpcId() *string 
}

type EnableVpcClassicLinkRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
  // 
  // >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-426655440000
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the VPC for which you want to enable ClassicLink.
  // 
  // You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The ID of the VPC.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // vpc-bp1m7v25emi1h5mtc****
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s EnableVpcClassicLinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcClassicLinkRequest) GoString() string {
  return s.String()
}

func (s *EnableVpcClassicLinkRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableVpcClassicLinkRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableVpcClassicLinkRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableVpcClassicLinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableVpcClassicLinkRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableVpcClassicLinkRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableVpcClassicLinkRequest) GetVpcId() *string  {
  return s.VpcId
}

func (s *EnableVpcClassicLinkRequest) SetClientToken(v string) *EnableVpcClassicLinkRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) SetOwnerAccount(v string) *EnableVpcClassicLinkRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) SetOwnerId(v int64) *EnableVpcClassicLinkRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) SetRegionId(v string) *EnableVpcClassicLinkRequest {
  s.RegionId = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) SetResourceOwnerAccount(v string) *EnableVpcClassicLinkRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) SetResourceOwnerId(v int64) *EnableVpcClassicLinkRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) SetVpcId(v string) *EnableVpcClassicLinkRequest {
  s.VpcId = &v
  return s
}

func (s *EnableVpcClassicLinkRequest) Validate() error {
  return dara.Validate(s)
}

