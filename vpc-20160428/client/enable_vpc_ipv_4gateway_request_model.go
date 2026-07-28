// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcIpv4GatewayRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableVpcIpv4GatewayRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableVpcIpv4GatewayRequest
  GetDryRun() *bool 
  SetIpv4GatewayId(v string) *EnableVpcIpv4GatewayRequest
  GetIpv4GatewayId() *string 
  SetOwnerAccount(v string) *EnableVpcIpv4GatewayRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableVpcIpv4GatewayRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableVpcIpv4GatewayRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableVpcIpv4GatewayRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableVpcIpv4GatewayRequest
  GetResourceOwnerId() *int64 
  SetRouteTableList(v []*string) *EnableVpcIpv4GatewayRequest
  GetRouteTableList() []*string 
}

type EnableVpcIpv4GatewayRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
  // 
  // > If you do not specify this parameter, the system automatically uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value may be different for each API request.
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-426655440000
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to perform a dry run. Valid values:
  // 
  // - **true**: performs a dry run. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
  // 
  // - **false*	- (default): sends a normal request, passes the dry run, and returns an HTTP 2xx status code to directly activate IPv4 gateway.
  // 
  // example:
  // 
  // false
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The instance ID of the IPv4 gateway that you want to activate.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ipv4gw-5tsp9lumsxoqizvq2****
  Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the IPv4 gateway.
  // 
  // You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ap-southeast-6
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The list of route tables. The system adds a route entry with the destination CIDR block 0.0.0.0/0 that points to the IPv4 gateway to each route table in the list.
  RouteTableList []*string `json:"RouteTableList,omitempty" xml:"RouteTableList,omitempty" type:"Repeated"`
}

func (s EnableVpcIpv4GatewayRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcIpv4GatewayRequest) GoString() string {
  return s.String()
}

func (s *EnableVpcIpv4GatewayRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableVpcIpv4GatewayRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableVpcIpv4GatewayRequest) GetIpv4GatewayId() *string  {
  return s.Ipv4GatewayId
}

func (s *EnableVpcIpv4GatewayRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableVpcIpv4GatewayRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableVpcIpv4GatewayRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableVpcIpv4GatewayRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableVpcIpv4GatewayRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableVpcIpv4GatewayRequest) GetRouteTableList() []*string  {
  return s.RouteTableList
}

func (s *EnableVpcIpv4GatewayRequest) SetClientToken(v string) *EnableVpcIpv4GatewayRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetDryRun(v bool) *EnableVpcIpv4GatewayRequest {
  s.DryRun = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetIpv4GatewayId(v string) *EnableVpcIpv4GatewayRequest {
  s.Ipv4GatewayId = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetOwnerAccount(v string) *EnableVpcIpv4GatewayRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetOwnerId(v int64) *EnableVpcIpv4GatewayRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetRegionId(v string) *EnableVpcIpv4GatewayRequest {
  s.RegionId = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetResourceOwnerAccount(v string) *EnableVpcIpv4GatewayRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetResourceOwnerId(v int64) *EnableVpcIpv4GatewayRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) SetRouteTableList(v []*string) *EnableVpcIpv4GatewayRequest {
  s.RouteTableList = v
  return s
}

func (s *EnableVpcIpv4GatewayRequest) Validate() error {
  return dara.Validate(s)
}

