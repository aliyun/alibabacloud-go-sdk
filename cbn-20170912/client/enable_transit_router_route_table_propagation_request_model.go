// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTransitRouterRouteTablePropagationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableTransitRouterRouteTablePropagationRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableTransitRouterRouteTablePropagationRequest
  GetDryRun() *bool 
  SetOwnerAccount(v string) *EnableTransitRouterRouteTablePropagationRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableTransitRouterRouteTablePropagationRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableTransitRouterRouteTablePropagationRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableTransitRouterRouteTablePropagationRequest
  GetResourceOwnerId() *int64 
  SetTransitRouterAttachmentId(v string) *EnableTransitRouterRouteTablePropagationRequest
  GetTransitRouterAttachmentId() *string 
  SetTransitRouterRouteTableId(v string) *EnableTransitRouterRouteTablePropagationRequest
  GetTransitRouterRouteTableId() *string 
}

type EnableTransitRouterRouteTablePropagationRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
  // 
  // > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
  // 
  // example:
  // 
  // 02fb3da4-130e-11e9-8e44-001****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
  // 
  // - **false*	- (default): Sends a normal request and creates the route learning relationship after the request passes the check.
  // 
  // - **true**: Sends a check request. No route learning relationship is created after the request passes the check. The system checks required parameters, request format, and other conditions. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
  // 
  // example:
  // 
  // false
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The ID of the network instance connection.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // tr-attach-nls9fzkfat8934****
  TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
  // The ID of the Enterprise Edition transit router route table.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // vtb-bp1dudbh2d5na6b50****
  TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s EnableTransitRouterRouteTablePropagationRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableTransitRouterRouteTablePropagationRequest) GoString() string {
  return s.String()
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetTransitRouterAttachmentId() *string  {
  return s.TransitRouterAttachmentId
}

func (s *EnableTransitRouterRouteTablePropagationRequest) GetTransitRouterRouteTableId() *string  {
  return s.TransitRouterRouteTableId
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetClientToken(v string) *EnableTransitRouterRouteTablePropagationRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetDryRun(v bool) *EnableTransitRouterRouteTablePropagationRequest {
  s.DryRun = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetOwnerAccount(v string) *EnableTransitRouterRouteTablePropagationRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetOwnerId(v int64) *EnableTransitRouterRouteTablePropagationRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetResourceOwnerAccount(v string) *EnableTransitRouterRouteTablePropagationRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetResourceOwnerId(v int64) *EnableTransitRouterRouteTablePropagationRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetTransitRouterAttachmentId(v string) *EnableTransitRouterRouteTablePropagationRequest {
  s.TransitRouterAttachmentId = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetTransitRouterRouteTableId(v string) *EnableTransitRouterRouteTablePropagationRequest {
  s.TransitRouterRouteTableId = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) Validate() error {
  return dara.Validate(s)
}

