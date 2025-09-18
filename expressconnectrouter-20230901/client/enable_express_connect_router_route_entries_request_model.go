// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableExpressConnectRouterRouteEntriesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableExpressConnectRouterRouteEntriesRequest
  GetClientToken() *string 
  SetDestinationCidrBlock(v string) *EnableExpressConnectRouterRouteEntriesRequest
  GetDestinationCidrBlock() *string 
  SetDryRun(v bool) *EnableExpressConnectRouterRouteEntriesRequest
  GetDryRun() *bool 
  SetEcrId(v string) *EnableExpressConnectRouterRouteEntriesRequest
  GetEcrId() *string 
  SetNexthopInstanceId(v string) *EnableExpressConnectRouterRouteEntriesRequest
  GetNexthopInstanceId() *string 
  SetVersion(v string) *EnableExpressConnectRouterRouteEntriesRequest
  GetVersion() *string 
}

type EnableExpressConnectRouterRouteEntriesRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
  // 
  // >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
  // 
  // example:
  // 
  // FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The destination CIDR block of the route entry in the route table of the ECR.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10.153.32.16/28
  DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
  // Specifies whether to perform only a dry run, without performing the actual request. Valid values:
  // 
  // 	- **true**: performs only a dry run.
  // 
  // 	- **false*	- (default): performs a dry run and performs the actual request.
  // 
  // example:
  // 
  // true
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The ECR ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ecr-mezk2idmsd0vx2****
  EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
  // The ID of the next-hop instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // tr-hp3u4u5f03tfuljis****
  NexthopInstanceId *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
  Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesRequest) GoString() string {
  return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) GetDestinationCidrBlock() *string  {
  return s.DestinationCidrBlock
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) GetEcrId() *string  {
  return s.EcrId
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) GetNexthopInstanceId() *string  {
  return s.NexthopInstanceId
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) GetVersion() *string  {
  return s.Version
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *EnableExpressConnectRouterRouteEntriesRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *EnableExpressConnectRouterRouteEntriesRequest {
  s.DestinationCidrBlock = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *EnableExpressConnectRouterRouteEntriesRequest {
  s.DryRun = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *EnableExpressConnectRouterRouteEntriesRequest {
  s.EcrId = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *EnableExpressConnectRouterRouteEntriesRequest {
  s.NexthopInstanceId = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetVersion(v string) *EnableExpressConnectRouterRouteEntriesRequest {
  s.Version = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) Validate() error {
  return dara.Validate(s)
}

