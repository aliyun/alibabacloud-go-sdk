// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKibanaPvlNetworkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndpointName(v string) *EnableKibanaPvlNetworkRequest
  GetEndpointName() *string 
  SetSecurityGroups(v []*string) *EnableKibanaPvlNetworkRequest
  GetSecurityGroups() []*string 
  SetVSwitchIdsZone(v []*EnableKibanaPvlNetworkRequestVSwitchIdsZone) *EnableKibanaPvlNetworkRequest
  GetVSwitchIdsZone() []*EnableKibanaPvlNetworkRequestVSwitchIdsZone 
  SetVpcId(v string) *EnableKibanaPvlNetworkRequest
  GetVpcId() *string 
  SetClientToken(v string) *EnableKibanaPvlNetworkRequest
  GetClientToken() *string 
}

type EnableKibanaPvlNetworkRequest struct {
  // The endpoint name.
  // 
  // example:
  // 
  // es-cn-27a3mul6l000xxx-kibana-endpoint
  EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
  // The security groups.
  // 
  // This parameter is required.
  SecurityGroups []*string `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" type:"Repeated"`
  // The vSwitch and zone information.
  VSwitchIdsZone []*EnableKibanaPvlNetworkRequestVSwitchIdsZone `json:"vSwitchIdsZone,omitempty" xml:"vSwitchIdsZone,omitempty" type:"Repeated"`
  // The VPC-connected instance ID.
  // 
  // example:
  // 
  // vpc-bp19ip2ocyv24w0e2****
  VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
  // The client token that is used to ensure the idempotence of the request.
  // 
  // example:
  // 
  // xxxxx
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s EnableKibanaPvlNetworkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableKibanaPvlNetworkRequest) GoString() string {
  return s.String()
}

func (s *EnableKibanaPvlNetworkRequest) GetEndpointName() *string  {
  return s.EndpointName
}

func (s *EnableKibanaPvlNetworkRequest) GetSecurityGroups() []*string  {
  return s.SecurityGroups
}

func (s *EnableKibanaPvlNetworkRequest) GetVSwitchIdsZone() []*EnableKibanaPvlNetworkRequestVSwitchIdsZone  {
  return s.VSwitchIdsZone
}

func (s *EnableKibanaPvlNetworkRequest) GetVpcId() *string  {
  return s.VpcId
}

func (s *EnableKibanaPvlNetworkRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableKibanaPvlNetworkRequest) SetEndpointName(v string) *EnableKibanaPvlNetworkRequest {
  s.EndpointName = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetSecurityGroups(v []*string) *EnableKibanaPvlNetworkRequest {
  s.SecurityGroups = v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetVSwitchIdsZone(v []*EnableKibanaPvlNetworkRequestVSwitchIdsZone) *EnableKibanaPvlNetworkRequest {
  s.VSwitchIdsZone = v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetVpcId(v string) *EnableKibanaPvlNetworkRequest {
  s.VpcId = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetClientToken(v string) *EnableKibanaPvlNetworkRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) Validate() error {
  if s.VSwitchIdsZone != nil {
    for _, item := range s.VSwitchIdsZone {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnableKibanaPvlNetworkRequestVSwitchIdsZone struct {
  // The vSwitch ID.
  // 
  // example:
  // 
  // vsw-bp194pz9iezj6h1n5****
  VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
  // The zone ID.
  // 
  // example:
  // 
  // cn-hangzhou-h
  ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s EnableKibanaPvlNetworkRequestVSwitchIdsZone) String() string {
  return dara.Prettify(s)
}

func (s EnableKibanaPvlNetworkRequestVSwitchIdsZone) GoString() string {
  return s.String()
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) GetVswitchId() *string  {
  return s.VswitchId
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) SetVswitchId(v string) *EnableKibanaPvlNetworkRequestVSwitchIdsZone {
  s.VswitchId = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) SetZoneId(v string) *EnableKibanaPvlNetworkRequestVSwitchIdsZone {
  s.ZoneId = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) Validate() error {
  return dara.Validate(s)
}

